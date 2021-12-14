// dockercfg TODO: Find good package description
package dockercfg

import (
	"k8c.io/kubermatic/v2/pkg/controller/operator/common"
	"k8c.io/kubermatic/v2/pkg/resources/reconciling"

	corev1 "k8s.io/api/core/v1"
)

func SecretCreator(seedDockercfg *corev1.Secret) reconciling.NamedSecretCreatorGetter {
	return func() (string, reconciling.SecretCreator) {
		return common.UserClusterDockercfgSecretName, func(existing *corev1.Secret) (*corev1.Secret, error) {
			existing.Name = common.UserClusterDockercfgSecretName
			existing.Data = seedDockercfg.Data
			// foo
			existing.Type = corev1.SecretTypeDockerConfigJson

			return existing, nil
		}
	}
}
