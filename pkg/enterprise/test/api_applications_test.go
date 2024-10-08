/*
Anchore API

Testing ApplicationsAPIService

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

func Test_enterprise_ApplicationsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ApplicationsAPIService AddApplication", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ApplicationsAPI.AddApplication(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService AddApplicationVersion", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationId string

		resp, httpRes, err := apiClient.ApplicationsAPI.AddApplicationVersion(context.Background(), applicationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService AddArtifactToApplicationVersion", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationId string
		var applicationVersionId string

		resp, httpRes, err := apiClient.ApplicationsAPI.AddArtifactToApplicationVersion(context.Background(), applicationId, applicationVersionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService DeleteApplication", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationId string

		httpRes, err := apiClient.ApplicationsAPI.DeleteApplication(context.Background(), applicationId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService DeleteApplicationVersion", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationId string
		var applicationVersionId string

		httpRes, err := apiClient.ApplicationsAPI.DeleteApplicationVersion(context.Background(), applicationId, applicationVersionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService GetApplication", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationId string

		resp, httpRes, err := apiClient.ApplicationsAPI.GetApplication(context.Background(), applicationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService GetApplicationVersion", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationId string
		var applicationVersionId string

		resp, httpRes, err := apiClient.ApplicationsAPI.GetApplicationVersion(context.Background(), applicationId, applicationVersionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService GetApplicationVersionSbom", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationId string
		var applicationVersionId string

		resp, httpRes, err := apiClient.ApplicationsAPI.GetApplicationVersionSbom(context.Background(), applicationId, applicationVersionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService GetApplicationVersionVulnerabilities", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationId string
		var applicationVersionId string

		resp, httpRes, err := apiClient.ApplicationsAPI.GetApplicationVersionVulnerabilities(context.Background(), applicationId, applicationVersionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService GetApplicationVersions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationId string

		resp, httpRes, err := apiClient.ApplicationsAPI.GetApplicationVersions(context.Background(), applicationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService GetApplications", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ApplicationsAPI.GetApplications(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService ListArtifacts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationId string
		var applicationVersionId string

		resp, httpRes, err := apiClient.ApplicationsAPI.ListArtifacts(context.Background(), applicationId, applicationVersionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService RemoveArtifactFromApplicationVersion", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationId string
		var applicationVersionId string
		var associationId string

		httpRes, err := apiClient.ApplicationsAPI.RemoveArtifactFromApplicationVersion(context.Background(), applicationId, applicationVersionId, associationId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService UpdateApplication", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationId string

		resp, httpRes, err := apiClient.ApplicationsAPI.UpdateApplication(context.Background(), applicationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService UpdateApplicationVersion", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationId string
		var applicationVersionId string

		resp, httpRes, err := apiClient.ApplicationsAPI.UpdateApplicationVersion(context.Background(), applicationId, applicationVersionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
