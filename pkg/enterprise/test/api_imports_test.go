/*
Anchore API

Testing ImportsAPIService

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

func Test_enterprise_ImportsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ImportsAPIService CreateOperation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ImportsAPI.CreateOperation(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportsAPIService CreateSourcesOperation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ImportsAPI.CreateSourcesOperation(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportsAPIService FinalizeOperation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var operationId string

		resp, httpRes, err := apiClient.ImportsAPI.FinalizeOperation(context.Background(), operationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportsAPIService GetImportSourcesSbom", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var operationId string

		resp, httpRes, err := apiClient.ImportsAPI.GetImportSourcesSbom(context.Background(), operationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportsAPIService GetOperation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var operationId string

		resp, httpRes, err := apiClient.ImportsAPI.GetOperation(context.Background(), operationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportsAPIService GetSourcesOperation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var operationId string

		resp, httpRes, err := apiClient.ImportsAPI.GetSourcesOperation(context.Background(), operationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportsAPIService ImportContentSearches", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var operationId string

		resp, httpRes, err := apiClient.ImportsAPI.ImportContentSearches(context.Background(), operationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportsAPIService ImportFileContents", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var operationId string

		resp, httpRes, err := apiClient.ImportsAPI.ImportFileContents(context.Background(), operationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportsAPIService ImportImageConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var operationId string

		resp, httpRes, err := apiClient.ImportsAPI.ImportImageConfig(context.Background(), operationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportsAPIService ImportImageDockerfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var operationId string

		resp, httpRes, err := apiClient.ImportsAPI.ImportImageDockerfile(context.Background(), operationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportsAPIService ImportImageManifest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var operationId string

		resp, httpRes, err := apiClient.ImportsAPI.ImportImageManifest(context.Background(), operationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportsAPIService ImportImagePackages", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var operationId string

		resp, httpRes, err := apiClient.ImportsAPI.ImportImagePackages(context.Background(), operationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportsAPIService ImportImageParentManifest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var operationId string

		resp, httpRes, err := apiClient.ImportsAPI.ImportImageParentManifest(context.Background(), operationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportsAPIService ImportSecretSearches", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var operationId string

		resp, httpRes, err := apiClient.ImportsAPI.ImportSecretSearches(context.Background(), operationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportsAPIService InvalidateOperation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var operationId string

		resp, httpRes, err := apiClient.ImportsAPI.InvalidateOperation(context.Background(), operationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportsAPIService InvalidateSourcesOperation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var operationId string

		resp, httpRes, err := apiClient.ImportsAPI.InvalidateSourcesOperation(context.Background(), operationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportsAPIService ListImportContentSearches", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var operationId string

		resp, httpRes, err := apiClient.ImportsAPI.ListImportContentSearches(context.Background(), operationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportsAPIService ListImportDockerfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var operationId string

		resp, httpRes, err := apiClient.ImportsAPI.ListImportDockerfiles(context.Background(), operationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportsAPIService ListImportFileContents", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var operationId string

		resp, httpRes, err := apiClient.ImportsAPI.ListImportFileContents(context.Background(), operationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportsAPIService ListImportImageConfigs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var operationId string

		resp, httpRes, err := apiClient.ImportsAPI.ListImportImageConfigs(context.Background(), operationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportsAPIService ListImportImageManifests", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var operationId string

		resp, httpRes, err := apiClient.ImportsAPI.ListImportImageManifests(context.Background(), operationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportsAPIService ListImportPackages", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var operationId string

		resp, httpRes, err := apiClient.ImportsAPI.ListImportPackages(context.Background(), operationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportsAPIService ListImportParentManifests", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var operationId string

		resp, httpRes, err := apiClient.ImportsAPI.ListImportParentManifests(context.Background(), operationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportsAPIService ListImportSecretSearches", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var operationId string

		resp, httpRes, err := apiClient.ImportsAPI.ListImportSecretSearches(context.Background(), operationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportsAPIService ListOperations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ImportsAPI.ListOperations(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportsAPIService ListSourcesOperations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ImportsAPI.ListSourcesOperations(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportsAPIService UploadImportSourcesSbom", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var operationId string

		resp, httpRes, err := apiClient.ImportsAPI.UploadImportSourcesSbom(context.Background(), operationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}