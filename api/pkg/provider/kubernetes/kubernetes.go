package kubernetes

import (
	"fmt"

	client "github.com/kubermatic/kubermatic/api/pkg/crd/client/master/clientset/versioned"
	kubermaticv1 "github.com/kubermatic/kubermatic/api/pkg/crd/kubermatic/v1"
	"github.com/kubermatic/kubermatic/api/pkg/provider"
	"github.com/kubermatic/kubermatic/api/pkg/util/auth"
	"github.com/kubermatic/kubermatic/api/pkg/util/errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/rand"
)

const (
	userLabelKey = "user"
)

type kubernetesProvider struct {
	crdClient client.Interface

	cps        map[string]provider.CloudProvider
	dcs        map[string]provider.DatacenterMeta
	workerName string
}

// NewKubernetesProvider creates a new kubernetes provider object
func NewKubernetesProvider(
	crdClient client.Interface,
	cps map[string]provider.CloudProvider,
	workerName string,
	dcs map[string]provider.DatacenterMeta,
) provider.ClusterProvider {
	return &kubernetesProvider{
		cps:        cps,
		crdClient:  crdClient,
		workerName: workerName,
		dcs:        dcs,
	}
}

func (p *kubernetesProvider) NewClusterWithCloud(user auth.User, spec *kubermaticv1.ClusterSpec) (*kubermaticv1.Cluster, error) {
	clusters, err := p.Clusters(user)
	if err != nil {
		return nil, err
	}

	// sanity checks for a fresh cluster
	switch {
	case user.ID == "":
		return nil, errors.NewBadRequest("cluster user is required")
	case spec.HumanReadableName == "":
		return nil, errors.NewBadRequest("cluster humanReadableName is required")
	}

	clusterName := rand.String(9)

	for _, c := range clusters {
		if c.Spec.HumanReadableName == spec.HumanReadableName {
			return nil, errors.NewAlreadyExists("cluster", spec.HumanReadableName)
		}
	}

	if spec.SeedDatacenterName == "" {
		dc, found := p.dcs[spec.Cloud.DatacenterName]
		if !found {
			return nil, errors.NewBadRequest("Unknown datacenter")
		}
		spec.SeedDatacenterName = dc.Seed
	}

	c := &kubermaticv1.Cluster{
		ObjectMeta: metav1.ObjectMeta{
			Name: clusterName,
			Labels: map[string]string{
				kubermaticv1.WorkerNameLabelKey: p.workerName,
				userLabelKey:                    user.ID,
			},
		},
		Spec: *spec,
		Status: kubermaticv1.ClusterStatus{
			LastTransitionTime: metav1.Now(),
			Seed:               spec.SeedDatacenterName,
			NamespaceName:      NamespaceName(clusterName),
			UserEmail:          user.Email,
			UserName:           user.Name,
		},
		Address: &kubermaticv1.ClusterAddress{},
	}

	c.Spec.WorkerName = p.workerName
	_, prov, err := provider.ClusterCloudProvider(p.cps, c)
	if err != nil {
		return nil, err
	}

	if err = prov.ValidateCloudSpec(c.Spec.Cloud); err != nil {
		return nil, fmt.Errorf("cloud provider data could not be validated successfully: %v", err)
	}

	c, err = p.crdClient.KubermaticV1().Clusters().Create(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (p *kubernetesProvider) Cluster(user auth.User, cluster string) (*kubermaticv1.Cluster, error) {
	c, err := p.crdClient.KubermaticV1().Clusters().Get(cluster, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	if c.Labels[userLabelKey] != user.ID && !user.IsAdmin() {
		return nil, errors.NewNotAuthorized()
	}
	return c, nil
}

func (p *kubernetesProvider) Clusters(user auth.User) ([]*kubermaticv1.Cluster, error) {
	filter := map[string]string{}
	if !user.IsAdmin() {
		filter[userLabelKey] = user.ID
	}
	selector := labels.SelectorFromSet(labels.Set(filter)).String()
	options := metav1.ListOptions{LabelSelector: selector, FieldSelector: labels.Everything().String()}
	clusterList, err := p.crdClient.KubermaticV1().Clusters().List(options)
	if err != nil {
		return nil, err
	}
	res := []*kubermaticv1.Cluster{}
	for i := range clusterList.Items {
		res = append(res, &clusterList.Items[i])
	}
	return res, nil
}

func (p *kubernetesProvider) DeleteCluster(user auth.User, cluster string) error {
	// check permission by getting the cluster first
	c, err := p.Cluster(user, cluster)
	if err != nil {
		return err
	}

	return p.crdClient.KubermaticV1().Clusters().Delete(c.Name, &metav1.DeleteOptions{})
}

func (p *kubernetesProvider) InitiateClusterUpgrade(user auth.User, name, version string) (*kubermaticv1.Cluster, error) {
	c, err := p.Cluster(user, name)
	if err != nil {
		return nil, err
	}

	c.Spec.MasterVersion = version
	c.Status.Phase = kubermaticv1.UpdatingMasterClusterStatusPhase
	c.Status.LastTransitionTime = metav1.Now()
	c.Status.MasterUpdatePhase = kubermaticv1.StartMasterUpdatePhase

	return p.crdClient.KubermaticV1().Clusters().Update(c)
}
