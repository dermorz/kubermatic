/*
Copyright 2022 The Kubermatic Kubernetes Platform contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package dockercfg

import (
	"k8c.io/kubermatic/v2/pkg/controller/operator/common"
	"k8c.io/kubermatic/v2/pkg/resources/reconciling"

	corev1 "k8s.io/api/core/v1"
)

// SecretCreator returns a function to create a secret in the usercluster that can be used as imagePullSecret.
func SecretCreator(seedDockercfg *corev1.Secret) reconciling.NamedSecretCreatorGetter {
	return func() (string, reconciling.SecretCreator) {
		return common.UserClusterDockercfgSecretName, func(existing *corev1.Secret) (*corev1.Secret, error) {
			existing.Name = common.UserClusterDockercfgSecretName
			existing.Data = seedDockercfg.Data
			existing.Type = corev1.SecretTypeDockerConfigJson

			return existing, nil
		}
	}
}
