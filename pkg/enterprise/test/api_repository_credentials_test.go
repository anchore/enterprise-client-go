/*
Anchore API

Testing RepositoryCredentialsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package enterprise

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func Test_enterprise_RepositoryCredentialsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RepositoryCredentialsAPIService AddRepository", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RepositoryCredentialsAPI.AddRepository(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
