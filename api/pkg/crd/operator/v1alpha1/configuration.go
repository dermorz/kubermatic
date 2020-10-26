package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
)

// ExposeStrategy is the strategy to expose the cluster with.
type ExposeStrategy string

const (
	// NodePortStrategy creates a NodePort with a "nodeport-proxy.k8s.io/expose": "true" annotation to expose
	// all clusters on one central Service of type LoadBalancer via the NodePort proxy.
	NodePortStrategy ExposeStrategy = "NodePort"
	// LoadBalancerStrategy creates a LoadBalancer service per cluster.
	LoadBalancerStrategy ExposeStrategy = "LoadBalancer"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubermaticConfiguration is the configuration required for running Kubermatic.
type KubermaticConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec KubermaticConfigurationSpec `json:"spec"`
}

// KubermaticConfigurationSpec is the spec for a Kubermatic installation.
type KubermaticConfigurationSpec struct {
	// ImagePullSecret is used to authenticate against Docker registries.
	ImagePullSecret string `json:"imagePullSecret,omitempty"`
	// Auth defines keys and URLs for Dex.
	Auth KubermaticAuthConfiguration `json:"auth"`
	// FeatureGates are used to optionally enable certain features.
	FeatureGates sets.String `json:"featureGates,omitempty"`
	// UI configures the dashboard.
	UI KubermaticUIConfiguration `json:"ui,omitempty"`
	// API configures the frontend REST API used by the dashboard.
	API KubermaticAPIConfiguration `json:"api,omitempty"`
	// SeedController configures the seed-controller-manager.
	SeedController KubermaticSeedControllerConfiguration `json:"seedController,omitempty"`
	// MasterController configures the master-controller-manager.
	MasterController KubermaticMasterControllerConfiguration `json:"masterController,omitempty"`
	// UserCluster configures various aspects of the user-created clusters.
	UserCluster KubermaticUserClusterConfiguration `json:"userCluster,omitempty"`
	// MasterFiles is a map of additional files to mount into each master component.
	MasterFiles map[string]string `json:"masterFiles,omitempty"`
	// ExposeStrategy is the strategy to expose the cluster with.
	// Note: The `seed_dns_overwrite` setting of a Seed's datacenter doesn't have any effect
	// if this is set to LoadBalancerStrategy.
	ExposeStrategy ExposeStrategy `json:"exposeStrategy,omitempty"`
	// Ingress contains settings for making the API and UI accessible remotely.
	Ingress KubermaticIngressConfiguration `json:"ingress,omitempty"`
}

// KubermaticAuthConfiguration defines keys and URLs for Dex.
type KubermaticAuthConfiguration struct {
	ClientID                 string `json:"clientID,omitempty"`
	TokenIssuer              string `json:"tokenIssuer,omitempty"`
	IssuerRedirectURL        string `json:"issuerRedirectURL,omitempty"`
	IssuerClientID           string `json:"issuerClientID,omitempty"`
	IssuerClientSecret       string `json:"issuerClientSecret,omitempty"`
	IssuerCookieKey          string `json:"issuerCookieKey,omitempty"`
	CABundle                 string `json:"caBundle,omitempty"`
	ServiceAccountKey        string `json:"serviceAccountKey,omitempty"`
	SkipTokenIssuerTLSVerify bool   `json:"skipTokenIssuerTLSVerify,omitempty"`
}

// KubermaticAPIConfiguration configures the dashboard.
type KubermaticAPIConfiguration struct {
	// DockerRepository is the repository containing the Kubermatic REST API image.
	DockerRepository string `json:"dockerRepository,omitempty"`
	// AccessibleAddons is a list of addons that should be enabled in the API.
	AccessibleAddons []string `json:"accessibleAddons,omitempty"`
	// PProfEndpoint controls the port the API should listen on to provide pprof
	// data. This port is never exposed from the container and only available via port-forwardings.
	PProfEndpoint *string `json:"pprofEndpoint,omitempty"`
	// Resources describes the requested and maximum allowed CPU/memory usage.
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`
	// DebugLog enables more verbose logging.
	DebugLog bool `json:"debugLog,omitempty"`
	// Replicas sets the number of pod replicas for the API deployment.
	Replicas *int32 `json:"replicas,omitempty"`
}

// KubermaticUIConfiguration configures the dashboard.
type KubermaticUIConfiguration struct {
	// DockerRepository is the repository containing the Kubermatic dashboard image.
	DockerRepository string `json:"dockerRepository,omitempty"`
	// Config sets flags for various dashboard features.
	Config string `json:"config,omitempty"`
	// Resources describes the requested and maximum allowed CPU/memory usage.
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`
	// Replicas sets the number of pod replicas for the UI deployment.
	Replicas *int32 `json:"replicas,omitempty"`
}

// KubermaticSeedControllerConfiguration configures the Kubermatic seed controller-manager.
type KubermaticSeedControllerConfiguration struct {
	// DockerRepository is the repository containing the Kubermatic seed-controller-manager image.
	DockerRepository string `json:"dockerRepository,omitempty"`
	// BackupStoreContainer is the container used for shipping etcd snapshots to a backup location.
	BackupStoreContainer string `json:"backupStoreContainer,omitempty"`
	// BackupCleanupContainer is the container used for removing expired backups from the storage location.
	BackupCleanupContainer string `json:"backupCleanupContainer,omitempty"`
	// PProfEndpoint controls the port the seed-controller-manager should listen on to provide pprof
	// data. This port is never exposed from the container and only available via port-forwardings.
	PProfEndpoint *string `json:"pprofEndpoint,omitempty"`
	// Resources describes the requested and maximum allowed CPU/memory usage.
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`
	// DebugLog enables more verbose logging.
	DebugLog bool `json:"debugLog,omitempty"`
	// Replicas sets the number of pod replicas for the seed-controller-manager.
	Replicas *int32 `json:"replicas,omitempty"`
}

// KubermaticUserClusterConfiguration controls various aspects of the user-created clusters.
type KubermaticUserClusterConfiguration struct {
	// KubermaticDockerRepository is the repository containing the Kubermatic user-cluster-controller-manager image.
	KubermaticDockerRepository string `json:"kubermaticDockerRepository,omitempty"`
	// DNATControllerDockerRepository is the repository containing the Kubermatic user-cluster-controller-manager image.
	DNATControllerDockerRepository string `json:"dnatControllerDockerRepository,omitempty"`
	// OverwriteRegistry specifies a custom Docker registry which will be used for all images
	// used inside user clusters (user cluster control plane + addons). This also applies to
	// the KubermaticDockerRepository and DNATControllerDockerRepository fields.
	OverwriteRegistry string `json:"overwriteRegistry,omitempty"`
	// Addons controls the optional additions installed into each user cluster.
	Addons KubermaticAddonsConfiguration `json:"addons,omitempty"`
	// NodePortRange is the port range for customer clusters - this must match the NodePort
	// range of the seed cluster.
	NodePortRange string `json:"nodePortRange,omitempty"`
	// Monitoring can be used to fine-tune to in-cluster Prometheus.
	Monitoring KubermaticUserClusterMonitoringConfiguration `json:"monitoring,omitempty"`
	// DisableAPIServerEndpointReconciling can be used to toggle the `--endpoint-reconciler-type` flag for
	// the Kubernetes API server.
	DisableAPIServerEndpointReconciling bool `json:"disableApiserverEndpointReconciling,omitempty"`
	// EtcdVolumeSize configures the volume size to use for each etcd pod inside user clusters.
	EtcdVolumeSize string `json:"etcdVolumeSize,omitempty"`
	// APIServerReplicas configures the replica count for the API-Server deployment inside user clusters.
	APIServerReplicas *int32 `json:"apiserverReplicas,omitempty"`
}

// KubermaticAddonsConfiguration controls the optional additions installed into each user cluster.
type KubermaticAddonsConfiguration struct {
	// Kubernetes controls the addons for Kubernetes-based clusters.
	Kubernetes KubermaticAddonConfiguration `json:"kubernetes,omitempty"`
	// Openshift controls the addons for Openshift-based clusters.
	Openshift KubermaticAddonConfiguration `json:"openshift,omitempty"`
}

// KubermaticUserClusterMonitoringConfiguration can be used to fine-tune to in-cluster Prometheus.
type KubermaticUserClusterMonitoringConfiguration struct {
	// DisableDefaultRules disables the recording and alerting rules.
	DisableDefaultRules bool `json:"disableDefaultRules,omitempty"`
	// DisableDefaultScrapingConfigs disables the default scraping targets.
	DisableDefaultScrapingConfigs bool `json:"disableDefaultScrapingConfigs,omitempty"`
	// CustomRules can be used to inject custom recording and alerting rules. This field
	// must be a YAML-formatted string with a `group` element at its root, as documented
	// on https://prometheus.io/docs/prometheus/2.14/configuration/alerting_rules/.
	CustomRules string `json:"customRules,omitempty"`
	// CustomScrapingConfigs can be used to inject custom scraping rules. This must be a
	// YAML-formatted string containing an array of scrape configurations as documented
	// on https://prometheus.io/docs/prometheus/2.14/configuration/configuration/#scrape_config.
	CustomScrapingConfigs string `json:"customScrapingConfigs,omitempty"`
	// ScrapeAnnotationPrefix (if set) is used to make the in-cluster Prometheus scrape pods
	// inside the user clusters.
	ScrapeAnnotationPrefix string `json:"scrapeAnnotationPrefix,omitempty"`
}

// KubermaticAddonConfiguration describes the addons for a given cluster runtime.
type KubermaticAddonConfiguration struct {
	// Default is the list of addons to be installed by default into each cluster.
	// Mutually exclusive with "defaultManifests".
	Default []string `json:"default,omitempty"`
	// DefaultManifests is a list of addon manifests to install into all clusters.
	// Mutually exclusive with "default".
	DefaultManifests string `json:"defaultManifests,omitempty"`
	// DockerRepository is the repository containing the Docker image containing
	// the possible addon manifests.
	DockerRepository string `json:"dockerRepository,omitempty"`
}

type KubermaticIngressConfiguration struct {
	// Domain is the base domain where the dashboard shall be available. Even with
	// a disabled Ingress, this must always be a valid hostname.
	Domain string `json:"domain"`

	// ClassName is the Ingress resource's class name, used for selecting the appropriate
	// ingress controller.
	ClassName string `json:"className,omitempty"`

	// Disable will prevent an Ingress from being created at all. This is mostly useful
	// during testing. If the Ingress is disabled, the CertificateIssuer setting can also
	// be left empty, as no Certificate resource will be created.
	Disable bool `json:"disable,omitempty"`

	// CertificateIssuer is the name of a cert-manager Issuer or ClusterIssuer (default)
	// that will be used to acquire the certificate for the configured domain.
	// To use a namespaced Issuer, set the Kind to "Issuer" and manually create the
	// matching Issuer in Kubermatic's namespace.
	CertificateIssuer corev1.TypedLocalObjectReference `json:"certificateIssuer,omitempty"`
}

// KubermaticMasterControllerConfiguration configures the Kubermatic master controller-manager.
type KubermaticMasterControllerConfiguration struct {
	// DockerRepository is the repository containing the Kubermatic master-controller-manager image.
	DockerRepository string `json:"dockerRepository,omitempty"`
	// ProjectsMigrator configures the migrator for user projects.
	ProjectsMigrator KubermaticProjectsMigratorConfiguration `json:"projectsMigrator,omitempty"`
	// PProfEndpoint controls the port the master-controller-manager should listen on to provide pprof
	// data. This port is never exposed from the container and only available via port-forwardings.
	PProfEndpoint *string `json:"pprofEndpoint,omitempty"`
	// Resources describes the requested and maximum allowed CPU/memory usage.
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`
	// DebugLog enables more verbose logging.
	DebugLog bool `json:"debugLog,omitempty"`
	// Replicas sets the number of pod replicas for the master-controller-manager.
	Replicas *int32 `json:"replicas,omitempty"`
}

// KubermaticProjectsMigratorConfiguration configures the Kubermatic master controller-manager.
type KubermaticProjectsMigratorConfiguration struct {
	// DryRun makes the migrator only log the actions it would take.
	DryRun bool `json:"dryRun,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubermaticConfigurationList is a collection of KubermaticConfigurations.
type KubermaticConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []KubermaticConfiguration `json:"items"`
}