#!/usr/bin/env bash

# Copyright 2020 The Kubermatic Kubernetes Platform contributors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -euo pipefail

cd $(dirname $0)/..
source hack/lib.sh

echodate "Checking Kubermatic CE licenses..."
boilerplate \
  -boilerplates hack/boilerplate/ce \
  -exclude addons/canal/canal.yaml \
  -exclude pkg/controller/seed-controller-manager/addon/testdata/istio \
  -exclude containers/startup-script/manage-startup-script.sh \
  -exclude pkg/resources/certificates/triple/triple.go \
  -exclude pkg/resources/etcd/testdata \
  -exclude pkg/ee

echodate "Checking Kubermatic EE licenses..."
boilerplate \
  -boilerplates hack/boilerplate/ee \
  pkg/ee