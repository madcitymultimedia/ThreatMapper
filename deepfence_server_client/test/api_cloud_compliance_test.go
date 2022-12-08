/*
Deepfence ThreatMapper

Testing CloudComplianceApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package deepfence_server_client

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/deepfence/ThreatMapper/deepfence_server_client"
)

func Test_deepfence_server_client_CloudComplianceApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CloudComplianceApiService IngestCloudCompliances", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CloudComplianceApi.IngestCloudCompliances(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
