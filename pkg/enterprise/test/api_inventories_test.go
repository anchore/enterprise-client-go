/*
Anchore API

Testing InventoriesAPIService

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

func Test_enterprise_InventoriesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test InventoriesAPIService DeleteInventory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.InventoriesAPI.DeleteInventory(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoriesAPIService DeleteKubernetesNamespaces", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.InventoriesAPI.DeleteKubernetesNamespaces(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoriesAPIService GetEcsContainers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.InventoriesAPI.GetEcsContainers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoriesAPIService GetEcsServices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.InventoriesAPI.GetEcsServices(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoriesAPIService GetEcsTasks", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.InventoriesAPI.GetEcsTasks(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoriesAPIService GetImageInventory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.InventoriesAPI.GetImageInventory(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoriesAPIService GetKubernetesContainers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.InventoriesAPI.GetKubernetesContainers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoriesAPIService GetKubernetesNamespace", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespaceId string

		resp, httpRes, err := apiClient.InventoriesAPI.GetKubernetesNamespace(context.Background(), namespaceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoriesAPIService GetKubernetesNamespaces", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.InventoriesAPI.GetKubernetesNamespaces(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoriesAPIService GetKubernetesNode", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeId string

		resp, httpRes, err := apiClient.InventoriesAPI.GetKubernetesNode(context.Background(), nodeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoriesAPIService GetKubernetesNodes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.InventoriesAPI.GetKubernetesNodes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoriesAPIService GetKubernetesPod", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var podId string

		resp, httpRes, err := apiClient.InventoriesAPI.GetKubernetesPod(context.Background(), podId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoriesAPIService GetKubernetesPods", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.InventoriesAPI.GetKubernetesPods(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoriesAPIService PostEcsInventory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.InventoriesAPI.PostEcsInventory(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoriesAPIService PostKubernetesInventory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.InventoriesAPI.PostKubernetesInventory(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
