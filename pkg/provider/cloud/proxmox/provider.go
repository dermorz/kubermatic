/*
Copyright 2020 The Kubermatic Kubernetes Platform contributors.

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

package proxmox

import (
	"context"
	"errors"

	kubermaticv1 "k8c.io/kubermatic/v2/pkg/apis/kubermatic/v1"
	"k8c.io/kubermatic/v2/pkg/provider"
)

type Proxmox struct {
}

// NewCloudProvider creates a new fake cloud provider.
func NewCloudProvider(dc *kubermaticv1.Datacenter) (*Proxmox, error) {
	if dc.Spec.Proxmox == nil {
		return nil, errors.New("datacenter is not Proxmox datacenter")
	}

	return &Proxmox{}, nil
}

func (proxmox *Proxmox) InitializeCloudProvider(_ context.Context, _ *kubermaticv1.Cluster, _ provider.ClusterUpdater) (*kubermaticv1.Cluster, error) {
	panic("not implemented") // TODO: Implement
}

func (proxmox *Proxmox) CleanUpCloudProvider(_ context.Context, _ *kubermaticv1.Cluster, _ provider.ClusterUpdater) (*kubermaticv1.Cluster, error) {
	panic("not implemented") // TODO: Implement
}

func (proxmox *Proxmox) DefaultCloudSpec(_ context.Context, _ *kubermaticv1.CloudSpec) error {
	panic("not implemented") // TODO: Implement
}

func (proxmox *Proxmox) ValidateCloudSpec(_ context.Context, _ kubermaticv1.CloudSpec) error {
	panic("not implemented") // TODO: Implement
}

func (proxmox *Proxmox) ValidateCloudSpecUpdate(ctx context.Context, oldSpec kubermaticv1.CloudSpec, newSpec kubermaticv1.CloudSpec) error {
	panic("not implemented") // TODO: Implement
}
