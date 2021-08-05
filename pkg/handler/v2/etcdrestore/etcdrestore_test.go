/*
Copyright 2021 The Kubermatic Kubernetes Platform contributors.

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

package etcdrestore_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	apiv1 "k8c.io/kubermatic/v2/pkg/api/v1"
	apiv2 "k8c.io/kubermatic/v2/pkg/api/v2"
	"k8c.io/kubermatic/v2/pkg/handler/test"
	"k8c.io/kubermatic/v2/pkg/handler/test/hack"

	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

func TestCreateEndpoint(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		Name                      string
		ProjectID                 string
		ClusterID                 string
		ExistingKubermaticObjects []ctrlruntimeclient.Object
		ExistingAPIUser           *apiv1.User
		EtcdRestore               *apiv2.EtcdRestore
		ExpectedHTTPStatusCode    int
		ExpectedResponse          *apiv2.EtcdRestore
	}{
		{
			Name:      "create etcd restore in the given cluster",
			ProjectID: test.GenDefaultProject().Name,
			ClusterID: test.GenDefaultCluster().Name,
			ExistingKubermaticObjects: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				test.GenDefaultCluster(),
			),
			ExistingAPIUser:        test.GenDefaultAPIUser(),
			EtcdRestore:            test.GenAPIEtcdRestore("test-er", test.GenDefaultCluster().Name),
			ExpectedHTTPStatusCode: http.StatusCreated,
			ExpectedResponse:       test.GenAPIEtcdRestore("test-er", test.GenDefaultCluster().Name),
		},
		{
			Name:      "user john cannot create etcd restore that belongs to bob's cluster",
			ProjectID: test.GenDefaultProject().Name,
			ClusterID: test.GenDefaultCluster().Name,
			ExistingKubermaticObjects: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				test.GenDefaultCluster(),
				test.GenAdminUser("John", "john@acme.com", false),
			),
			ExistingAPIUser:        test.GenAPIUser("John", "john@acme.com"),
			EtcdRestore:            test.GenAPIEtcdRestore("test-er", test.GenDefaultCluster().Name),
			ExpectedHTTPStatusCode: http.StatusForbidden,
			ExpectedResponse:       nil,
		},
		{
			Name:      "admin user john can create etcd restore that belongs to bob's cluster",
			ProjectID: test.GenDefaultProject().Name,
			ClusterID: test.GenDefaultCluster().Name,
			ExistingKubermaticObjects: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				test.GenDefaultCluster(),
				test.GenAdminUser("John", "john@acme.com", true),
			),
			ExistingAPIUser:        test.GenAPIUser("John", "john@acme.com"),
			EtcdRestore:            test.GenAPIEtcdRestore("test-er", test.GenDefaultCluster().Name),
			ExpectedHTTPStatusCode: http.StatusCreated,
			ExpectedResponse:       test.GenAPIEtcdRestore("test-er", test.GenDefaultCluster().Name),
		},
		{
			Name:      "validation fails when backup name is empty",
			ProjectID: test.GenDefaultProject().Name,
			ClusterID: test.GenDefaultCluster().Name,
			ExistingKubermaticObjects: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				test.GenDefaultCluster(),
			),
			ExistingAPIUser: test.GenDefaultAPIUser(),
			EtcdRestore: func() *apiv2.EtcdRestore {
				er := test.GenAPIEtcdRestore("test-er", test.GenDefaultCluster().Name)
				er.Spec.BackupName = ""
				return er
			}(),
			ExpectedHTTPStatusCode: http.StatusBadRequest,
			ExpectedResponse:       nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			requestURL := fmt.Sprintf("/api/v2/projects/%s/clusters/%s/etcdrestores", tc.ProjectID, tc.ClusterID)
			body, err := json.Marshal(tc.EtcdRestore)
			if err != nil {
				t.Fatalf("failed marshalling etcdrestore: %v", err)
			}
			req := httptest.NewRequest(http.MethodPost, requestURL, bytes.NewBuffer(body))
			resp := httptest.NewRecorder()

			ep, err := test.CreateTestEndpoint(*tc.ExistingAPIUser, nil, tc.ExistingKubermaticObjects, nil, nil, hack.NewTestRouting)
			if err != nil {
				t.Fatalf("failed to create test endpoint due to: %v", err)
			}
			ep.ServeHTTP(resp, req)

			if resp.Code != tc.ExpectedHTTPStatusCode {
				t.Fatalf("Expected HTTP status code %d, got %d: %s", tc.ExpectedHTTPStatusCode, resp.Code, resp.Body.String())
			}
			if resp.Code == http.StatusCreated {
				b, err := json.Marshal(tc.ExpectedResponse)
				if err != nil {
					t.Fatalf("failed to marshal expected response %v", err)
				}
				test.CompareWithResult(t, resp, string(b))
			}
		})
	}
}

func TestGetEndpoint(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		Name                      string
		EtcdRestoreName           string
		ProjectID                 string
		ClusterID                 string
		ExistingKubermaticObjects []ctrlruntimeclient.Object
		ExistingAPIUser           *apiv1.User
		ExpectedResponse          *apiv2.EtcdRestore
		ExpectedHTTPStatusCode    int
	}{
		{
			Name:            "get etcdrestore that belongs to the given cluster",
			EtcdRestoreName: "test-1",
			ProjectID:       test.GenDefaultProject().Name,
			ClusterID:       test.GenDefaultCluster().Name,
			ExistingKubermaticObjects: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				test.GenDefaultCluster(),
				test.GenEtcdRestore("test-1", test.GenDefaultCluster()),
			),
			ExistingAPIUser:        test.GenDefaultAPIUser(),
			ExpectedHTTPStatusCode: http.StatusOK,
			ExpectedResponse:       test.GenAPIEtcdRestore("test-1", test.GenDefaultCluster().Name),
		},
		{
			Name:            "get etcdrestore which doesn't exist",
			EtcdRestoreName: "test-1",
			ProjectID:       test.GenDefaultProject().Name,
			ClusterID:       test.GenDefaultCluster().Name,
			ExistingKubermaticObjects: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				test.GenDefaultCluster(),
			),
			ExistingAPIUser:        test.GenDefaultAPIUser(),
			ExpectedHTTPStatusCode: http.StatusNotFound,
		},
		{
			Name:            "user john cannot get etcdrestore that belongs to bob's cluster",
			EtcdRestoreName: "test-1",
			ProjectID:       test.GenDefaultProject().Name,
			ClusterID:       test.GenDefaultCluster().Name,
			ExistingKubermaticObjects: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				test.GenDefaultCluster(),
				test.GenAdminUser("John", "john@acme.com", false),
				test.GenEtcdRestore("test-1", test.GenDefaultCluster()),
			),
			ExistingAPIUser:        test.GenAPIUser("John", "john@acme.com"),
			ExpectedHTTPStatusCode: http.StatusForbidden,
		},
		{
			Name:            "admin user john can get etcdrestore that belongs to bob's cluster",
			EtcdRestoreName: "test-1",
			ProjectID:       test.GenDefaultProject().Name,
			ClusterID:       test.GenDefaultCluster().Name,
			ExistingKubermaticObjects: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				test.GenDefaultCluster(),
				test.GenAdminUser("John", "john@acme.com", true),
				test.GenEtcdRestore("test-1", test.GenDefaultCluster()),
			),
			ExistingAPIUser:        test.GenAPIUser("John", "john@acme.com"),
			ExpectedHTTPStatusCode: http.StatusOK,
			ExpectedResponse:       test.GenAPIEtcdRestore("test-1", test.GenDefaultCluster().Name),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			requestURL := fmt.Sprintf("/api/v2/projects/%s/clusters/%s/etcdrestores/%s", tc.ProjectID, tc.ClusterID, tc.EtcdRestoreName)
			req := httptest.NewRequest(http.MethodGet, requestURL, nil)
			resp := httptest.NewRecorder()

			ep, err := test.CreateTestEndpoint(*tc.ExistingAPIUser, nil, tc.ExistingKubermaticObjects, nil, nil, hack.NewTestRouting)
			if err != nil {
				t.Fatalf("failed to create test endpoint due to: %v", err)
			}
			ep.ServeHTTP(resp, req)

			if resp.Code != tc.ExpectedHTTPStatusCode {
				t.Fatalf("Expected HTTP status code %d, got %d: %s", tc.ExpectedHTTPStatusCode, resp.Code, resp.Body.String())
			}
			if resp.Code == http.StatusOK {
				b, err := json.Marshal(tc.ExpectedResponse)
				if err != nil {
					t.Fatalf("failed to marshal expected response: %v", err)
				}

				test.CompareWithResult(t, resp, string(b))
			}

		})
	}
}

func TestListEndpoint(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		Name                      string
		ProjectID                 string
		ClusterID                 string
		ExistingKubermaticObjects []ctrlruntimeclient.Object
		ExistingAPIUser           *apiv1.User
		ExpectedResponse          []*apiv2.EtcdRestore
		ExpectedHTTPStatusCode    int
	}{
		{
			Name:      "list etcdrestores that belongs to the given cluster",
			ProjectID: test.GenDefaultProject().Name,
			ClusterID: test.GenDefaultCluster().Name,
			ExistingKubermaticObjects: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				test.GenDefaultCluster(),
				test.GenEtcdRestore("test-1", test.GenDefaultCluster()),
				test.GenEtcdRestore("test-2", test.GenDefaultCluster()),
			),
			ExistingAPIUser:        test.GenDefaultAPIUser(),
			ExpectedHTTPStatusCode: http.StatusOK,
			ExpectedResponse: []*apiv2.EtcdRestore{
				test.GenAPIEtcdRestore("test-1", test.GenDefaultCluster().Name),
				test.GenAPIEtcdRestore("test-2", test.GenDefaultCluster().Name),
			},
		},
		{
			Name:      "user john cannot list etcdrestores that belongs to bob's cluster",
			ProjectID: test.GenDefaultProject().Name,
			ClusterID: test.GenDefaultCluster().Name,
			ExistingKubermaticObjects: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				test.GenDefaultCluster(),
				test.GenAdminUser("John", "john@acme.com", false),
				test.GenEtcdRestore("test-1", test.GenDefaultCluster()),
				test.GenEtcdRestore("test-2", test.GenDefaultCluster()),
			),
			ExistingAPIUser:        test.GenAPIUser("John", "john@acme.com"),
			ExpectedHTTPStatusCode: http.StatusForbidden,
		},
		{
			Name:      "admin user john can get etcdrestore that belongs to bob's cluster",
			ProjectID: test.GenDefaultProject().Name,
			ClusterID: test.GenDefaultCluster().Name,
			ExistingKubermaticObjects: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				test.GenDefaultCluster(),
				test.GenAdminUser("John", "john@acme.com", true),
				test.GenEtcdRestore("test-1", test.GenDefaultCluster()),
				test.GenEtcdRestore("test-2", test.GenDefaultCluster()),
			),
			ExistingAPIUser:        test.GenAPIUser("John", "john@acme.com"),
			ExpectedHTTPStatusCode: http.StatusOK,
			ExpectedResponse: []*apiv2.EtcdRestore{
				test.GenAPIEtcdRestore("test-1", test.GenDefaultCluster().Name),
				test.GenAPIEtcdRestore("test-2", test.GenDefaultCluster().Name),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			requestURL := fmt.Sprintf("/api/v2/projects/%s/clusters/%s/etcdrestores", tc.ProjectID, tc.ClusterID)
			req := httptest.NewRequest(http.MethodGet, requestURL, nil)
			resp := httptest.NewRecorder()

			ep, err := test.CreateTestEndpoint(*tc.ExistingAPIUser, nil, tc.ExistingKubermaticObjects, nil, nil, hack.NewTestRouting)
			if err != nil {
				t.Fatalf("failed to create test endpoint due to: %v", err)
			}
			ep.ServeHTTP(resp, req)

			if resp.Code != tc.ExpectedHTTPStatusCode {
				t.Fatalf("Expected HTTP status code %d, got %d: %s", tc.ExpectedHTTPStatusCode, resp.Code, resp.Body.String())
			}
			if resp.Code == http.StatusOK {
				etcdRestores := test.NewEtcdRestoreSliceWrapper{}
				etcdRestores.DecodeOrDie(resp.Body, t).Sort()

				expectedEtcdRestores := test.NewEtcdRestoreSliceWrapper(tc.ExpectedResponse)
				expectedEtcdRestores.Sort()

				etcdRestores.EqualOrDie(expectedEtcdRestores, t)
			}

		})
	}
}