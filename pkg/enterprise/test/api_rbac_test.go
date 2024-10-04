/*
Anchore API

Testing RBACAPIService

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

func Test_enterprise_RBACAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RBACAPIService AddIdp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RBACAPI.AddIdp(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RBACAPIService AddIdpUserGroups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.RBACAPI.AddIdpUserGroups(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RBACAPIService AddRoleUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var roleName string

		resp, httpRes, err := apiClient.RBACAPI.AddRoleUser(context.Background(), roleName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RBACAPIService DeleteIdp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		httpRes, err := apiClient.RBACAPI.DeleteIdp(context.Background(), name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RBACAPIService DeleteIdpUserGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		httpRes, err := apiClient.RBACAPI.DeleteIdpUserGroup(context.Background(), name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RBACAPIService DeleteRoleUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var roleName string

		httpRes, err := apiClient.RBACAPI.DeleteRoleUser(context.Background(), roleName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RBACAPIService GetIdp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.RBACAPI.GetIdp(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RBACAPIService GetIdpUserGroups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.RBACAPI.GetIdpUserGroups(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RBACAPIService GetRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var roleName string

		resp, httpRes, err := apiClient.RBACAPI.GetRole(context.Background(), roleName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RBACAPIService ListIdps", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RBACAPI.ListIdps(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RBACAPIService ListRoleMembers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var roleName string

		resp, httpRes, err := apiClient.RBACAPI.ListRoleMembers(context.Background(), roleName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RBACAPIService ListRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RBACAPI.ListRoles(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RBACAPIService ListUserRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var username string

		resp, httpRes, err := apiClient.RBACAPI.ListUserRoles(context.Background(), username).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RBACAPIService MyRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RBACAPI.MyRoles(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RBACAPIService SamlLogin", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var idpName string

		resp, httpRes, err := apiClient.RBACAPI.SamlLogin(context.Background(), idpName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RBACAPIService SamlSso", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var idpName string

		resp, httpRes, err := apiClient.RBACAPI.SamlSso(context.Background(), idpName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RBACAPIService UpdateIdp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.RBACAPI.UpdateIdp(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
