package provider_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kubermatic/kubermatic/api/pkg/credentials"
	"github.com/kubermatic/kubermatic/api/pkg/handler/test"
	"github.com/kubermatic/kubermatic/api/pkg/handler/test/hack"
)

func TestGCPCredentialEndpoint(t *testing.T) {
	t.Parallel()
	testcases := []struct {
		name             string
		credentials      []credentials.GCPCredentials
		httpStatus       int
		expectedResponse string
	}{
		{
			name:             "test no credentials for GCP",
			httpStatus:       http.StatusOK,
			expectedResponse: "{}",
		},
		{
			name: "test list of credential names for GCP",
			credentials: []credentials.GCPCredentials{
				{Name: "first"},
				{Name: "second"},
			},
			httpStatus:       http.StatusOK,
			expectedResponse: `{"names":["first","second"]}`,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {

			req := httptest.NewRequest("GET", "/api/v1/providers/gcp/credentials", strings.NewReader(""))

			credentialsManager := credentials.New()
			cred := credentialsManager.GetCredentials()
			cred.GCP = tc.credentials

			res := httptest.NewRecorder()
			router, err := test.CreateCredentialTestEndpoint(credentialsManager, hack.NewTestRouting)
			if err != nil {
				t.Fatalf("failed to create test endpoint due to %v\n", err)
			}
			router.ServeHTTP(res, req)

			// validate
			assert.Equal(t, tc.httpStatus, res.Code)

			if res.Code == http.StatusOK {
				compareJSON(t, res, tc.expectedResponse)
			}
		})
	}
}
