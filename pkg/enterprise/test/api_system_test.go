/*
Anchore API

Testing SystemAPIService

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

func Test_enterprise_SystemAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SystemAPIService DeleteFeed", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var feed string

		httpRes, err := apiClient.SystemAPI.DeleteFeed(context.Background(), feed).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemAPIService DeleteService", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceName string
		var hostId string

		httpRes, err := apiClient.SystemAPI.DeleteService(context.Background(), serviceName, hostId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemAPIService DescribeErrorCodes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SystemAPI.DescribeErrorCodes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemAPIService DescribePolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SystemAPI.DescribePolicy(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemAPIService GetServiceDetail", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SystemAPI.GetServiceDetail(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemAPIService GetServicesByName", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceName string

		resp, httpRes, err := apiClient.SystemAPI.GetServicesByName(context.Background(), serviceName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemAPIService GetServicesByNameAndHost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceName string
		var hostId string

		resp, httpRes, err := apiClient.SystemAPI.GetServicesByNameAndHost(context.Background(), serviceName, hostId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemAPIService GetStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SystemAPI.GetStatus(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemAPIService GetSystemFeeds", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SystemAPI.GetSystemFeeds(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemAPIService ListServices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SystemAPI.ListServices(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemAPIService PostSystemFeeds", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SystemAPI.PostSystemFeeds(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemAPIService TestWebhook", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var webhookType string

		httpRes, err := apiClient.SystemAPI.TestWebhook(context.Background(), webhookType).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemAPIService ToggleFeedEnabled", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var feed string

		resp, httpRes, err := apiClient.SystemAPI.ToggleFeedEnabled(context.Background(), feed).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
