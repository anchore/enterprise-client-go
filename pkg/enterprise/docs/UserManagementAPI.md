# \UserManagementAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUserGroupRoles**](UserManagementAPI.md#AddUserGroupRoles) | **Post** /system/user-groups/{group_uuid}/roles | Add account role(s) to this user group
[**AddUserGroupUsers**](UserManagementAPI.md#AddUserGroupUsers) | **Post** /system/user-groups/{group_uuid}/users | Add user(s) to a user group
[**CreateAccount**](UserManagementAPI.md#CreateAccount) | **Post** /accounts | Create a new account. Only available to admin user.
[**CreateUser**](UserManagementAPI.md#CreateUser) | **Post** /accounts/{account_name}/users | Create a new user within the specified account.
[**CreateUserApiKey**](UserManagementAPI.md#CreateUserApiKey) | **Post** /accounts/{account_name}/users/{username}/api-keys | Add a new API key
[**CreateUserCredential**](UserManagementAPI.md#CreateUserCredential) | **Post** /accounts/{account_name}/users/{username}/credentials | add/replace credential
[**CreateUserGroup**](UserManagementAPI.md#CreateUserGroup) | **Post** /system/user-groups | Create a new user group
[**DeleteAccount**](UserManagementAPI.md#DeleteAccount) | **Delete** /accounts/{account_name} | Delete the specified account, only allowed if the account is in the disabled state. All users will be deleted along with the account and all resources will be garbage collected. The deleted account name will not be available for reuse immediately until all resources are garbage collected async.
[**DeleteUser**](UserManagementAPI.md#DeleteUser) | **Delete** /accounts/{account_name}/users/{username} | Delete a specific user by username.
[**DeleteUserApiKey**](UserManagementAPI.md#DeleteUserApiKey) | **Delete** /accounts/{account_name}/users/{username}/api-keys/{key_name} | Delete a user API key
[**DeleteUserCredential**](UserManagementAPI.md#DeleteUserCredential) | **Delete** /accounts/{account_name}/users/{username}/credentials | Delete a credential by type
[**DeleteUserGroup**](UserManagementAPI.md#DeleteUserGroup) | **Delete** /system/user-groups/{group_uuid} | Delete a user group
[**DeleteUserGroupRole**](UserManagementAPI.md#DeleteUserGroupRole) | **Delete** /system/user-groups/{group_uuid}/roles | Remove account role(s) from this user group
[**DeleteUserGroupUsers**](UserManagementAPI.md#DeleteUserGroupUsers) | **Delete** /system/user-groups/{group_uuid}/users | Remove a user from a user group
[**GetAccount**](UserManagementAPI.md#GetAccount) | **Get** /accounts/{account_name} | Get account info about this specific account.
[**GetAccountUser**](UserManagementAPI.md#GetAccountUser) | **Get** /accounts/{account_name}/users/{username} | Get a specific user in the specified account
[**GetOauthToken**](UserManagementAPI.md#GetOauthToken) | **Post** /oauth/token | 
[**GetUserApiKey**](UserManagementAPI.md#GetUserApiKey) | **Get** /accounts/{account_name}/users/{username}/api-keys/{key_name} | Get a user API key
[**GetUserGroup**](UserManagementAPI.md#GetUserGroup) | **Get** /system/user-groups/{group_uuid} | Get a user group
[**ListAccounts**](UserManagementAPI.md#ListAccounts) | **Get** /accounts | List account summaries. Only available to the system admin user.
[**ListAllSystemUsers**](UserManagementAPI.md#ListAllSystemUsers) | **Get** /accounts/users | List all system account users. Only available to &#x60;admin&#x60; users.
[**ListUserApiKeys**](UserManagementAPI.md#ListUserApiKeys) | **Get** /accounts/{account_name}/users/{username}/api-keys | Get a list of API keys
[**ListUserCredentials**](UserManagementAPI.md#ListUserCredentials) | **Get** /accounts/{account_name}/users/{username}/credentials | Get current credential summary
[**ListUserGroupRoles**](UserManagementAPI.md#ListUserGroupRoles) | **Get** /system/user-groups/{group_uuid}/roles | Get a list of all user group roles
[**ListUserGroupUsers**](UserManagementAPI.md#ListUserGroupUsers) | **Get** /system/user-groups/{group_uuid}/users | Get a list of user group users
[**ListUserGroups**](UserManagementAPI.md#ListUserGroups) | **Get** /system/user-groups | List user groups
[**ListUsers**](UserManagementAPI.md#ListUsers) | **Get** /accounts/{account_name}/users | List of users that are primary within this account. The response object will only contain roles for this account as well as any system roles.
[**ListUsersWithRoles**](UserManagementAPI.md#ListUsersWithRoles) | **Get** /accounts/{account_name}/users-with-roles | List of users who have rbac roles in this account.
[**PatchUserApiKey**](UserManagementAPI.md#PatchUserApiKey) | **Patch** /accounts/{account_name}/users/{username}/api-keys/{key_name} | Patch a user API key
[**RevokeOauthToken**](UserManagementAPI.md#RevokeOauthToken) | **Post** /oauth/revoke | 
[**UpdateAccount**](UserManagementAPI.md#UpdateAccount) | **Put** /accounts/{account_name} | Update the info for this specific account.
[**UpdateAccountState**](UserManagementAPI.md#UpdateAccountState) | **Put** /accounts/{account_name}/state | Update the state of an account to either enabled or disabled. For deletion use the DELETE route
[**UpdateUserGroup**](UserManagementAPI.md#UpdateUserGroup) | **Patch** /system/user-groups/{group_uuid} | Update a user group



## AddUserGroupRoles

> UserGroupRoles AddUserGroupRoles(ctx, groupUuid).UserGroupRolePost(userGroupRolePost).Execute()

Add account role(s) to this user group

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	groupUuid := "groupUuid_example" // string | 
	userGroupRolePost := *openapiclient.NewUserGroupRolePost([]openapiclient.UserGroupRolePostRolesInner{*openapiclient.NewUserGroupRolePostRolesInner()}) // UserGroupRolePost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.AddUserGroupRoles(context.Background(), groupUuid).UserGroupRolePost(userGroupRolePost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.AddUserGroupRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddUserGroupRoles`: UserGroupRoles
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.AddUserGroupRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddUserGroupRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userGroupRolePost** | [**UserGroupRolePost**](UserGroupRolePost.md) |  | 

### Return type

[**UserGroupRoles**](UserGroupRoles.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddUserGroupUsers

> UserGroupUsers AddUserGroupUsers(ctx, groupUuid).UserGroupUsersPost(userGroupUsersPost).Execute()

Add user(s) to a user group

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	groupUuid := "groupUuid_example" // string | 
	userGroupUsersPost := *openapiclient.NewUserGroupUsersPost([]openapiclient.UserGroupUsersPostUsernamesInner{*openapiclient.NewUserGroupUsersPostUsernamesInner()}) // UserGroupUsersPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.AddUserGroupUsers(context.Background(), groupUuid).UserGroupUsersPost(userGroupUsersPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.AddUserGroupUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddUserGroupUsers`: UserGroupUsers
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.AddUserGroupUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddUserGroupUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userGroupUsersPost** | [**UserGroupUsersPost**](UserGroupUsersPost.md) |  | 

### Return type

[**UserGroupUsers**](UserGroupUsers.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccount

> Account CreateAccount(ctx).Account(account).Execute()

Create a new account. Only available to admin user.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	account := *openapiclient.NewAccountCreationRequest("Name_example") // AccountCreationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.CreateAccount(context.Background()).Account(account).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.CreateAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccount`: Account
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.CreateAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | [**AccountCreationRequest**](AccountCreationRequest.md) |  | 

### Return type

[**Account**](Account.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> User CreateUser(ctx, accountName).User(user).Execute()

Create a new user within the specified account.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	accountName := "accountName_example" // string | 
	user := *openapiclient.NewUserCreationRequest("Username_example") // UserCreationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.CreateUser(context.Background(), accountName).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.CreateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUser`: User
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.CreateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**UserCreationRequest**](UserCreationRequest.md) |  | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserApiKey

> UserApiKey CreateUserApiKey(ctx, accountName, username).Apikey(apikey).Execute()

Add a new API key

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	accountName := "accountName_example" // string | 
	username := "username_example" // string | 
	apikey := *openapiclient.NewUserApiKey("Name_example", time.Now()) // UserApiKey | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.CreateUserApiKey(context.Background(), accountName, username).Apikey(apikey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.CreateUserApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUserApiKey`: UserApiKey
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.CreateUserApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** |  | 
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apikey** | [**UserApiKey**](UserApiKey.md) |  | 

### Return type

[**UserApiKey**](UserApiKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserCredential

> AccessCredential CreateUserCredential(ctx, accountName, username).Credential(credential).Execute()

add/replace credential

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	accountName := "accountName_example" // string | 
	username := "username_example" // string | 
	credential := *openapiclient.NewAccessCredential("Type_example", "Value_example") // AccessCredential | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.CreateUserCredential(context.Background(), accountName, username).Credential(credential).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.CreateUserCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUserCredential`: AccessCredential
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.CreateUserCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** |  | 
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **credential** | [**AccessCredential**](AccessCredential.md) |  | 

### Return type

[**AccessCredential**](AccessCredential.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserGroup

> UserGroup CreateUserGroup(ctx).UserGroupPost(userGroupPost).Execute()

Create a new user group

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	userGroupPost := *openapiclient.NewUserGroupPost("Name_example") // UserGroupPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.CreateUserGroup(context.Background()).UserGroupPost(userGroupPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.CreateUserGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUserGroup`: UserGroup
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.CreateUserGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userGroupPost** | [**UserGroupPost**](UserGroupPost.md) |  | 

### Return type

[**UserGroup**](UserGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccount

> DeleteAccount(ctx, accountName).Execute()

Delete the specified account, only allowed if the account is in the disabled state. All users will be deleted along with the account and all resources will be garbage collected. The deleted account name will not be available for reuse immediately until all resources are garbage collected async.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	accountName := "accountName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserManagementAPI.DeleteAccount(context.Background(), accountName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.DeleteAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> DeleteUser(ctx, accountName, username).Execute()

Delete a specific user by username.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	accountName := "accountName_example" // string | 
	username := "username_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserManagementAPI.DeleteUser(context.Background(), accountName, username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.DeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** |  | 
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserApiKey

> DeleteUserApiKey(ctx, accountName, username, keyName).Execute()

Delete a user API key

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	accountName := "accountName_example" // string | 
	username := "username_example" // string | 
	keyName := "keyName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserManagementAPI.DeleteUserApiKey(context.Background(), accountName, username, keyName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.DeleteUserApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** |  | 
**username** | **string** |  | 
**keyName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserCredential

> DeleteUserCredential(ctx, accountName, username).CredentialType(credentialType).Execute()

Delete a credential by type

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	accountName := "accountName_example" // string | 
	username := "username_example" // string | 
	credentialType := "credentialType_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserManagementAPI.DeleteUserCredential(context.Background(), accountName, username).CredentialType(credentialType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.DeleteUserCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** |  | 
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **credentialType** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserGroup

> DeleteUserGroup(ctx, groupUuid).Execute()

Delete a user group

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	groupUuid := "groupUuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserManagementAPI.DeleteUserGroup(context.Background(), groupUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.DeleteUserGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserGroupRole

> DeleteUserGroupRole(ctx, groupUuid).MembershipId(membershipId).AllRolesForAccount(allRolesForAccount).Execute()

Remove account role(s) from this user group

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	groupUuid := "groupUuid_example" // string | 
	membershipId := []string{"Inner_example"} // []string | A list of membership ids to remove from the user group in the format of membership_id=1&membership_id=2 (optional)
	allRolesForAccount := []string{"Inner_example"} // []string | A list of accounts to remove all roles from the user group in the format of all_roles_for_account=account1&all_roles_for_account=account2 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserManagementAPI.DeleteUserGroupRole(context.Background(), groupUuid).MembershipId(membershipId).AllRolesForAccount(allRolesForAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.DeleteUserGroupRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserGroupRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **membershipId** | **[]string** | A list of membership ids to remove from the user group in the format of membership_id&#x3D;1&amp;membership_id&#x3D;2 | 
 **allRolesForAccount** | **[]string** | A list of accounts to remove all roles from the user group in the format of all_roles_for_account&#x3D;account1&amp;all_roles_for_account&#x3D;account2 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserGroupUsers

> DeleteUserGroupUsers(ctx, groupUuid).Username(username).Execute()

Remove a user from a user group

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	groupUuid := "groupUuid_example" // string | 
	username := []string{"Inner_example"} // []string | A list of usernames to remove from the user group in the format of username=user1&username=user2

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserManagementAPI.DeleteUserGroupUsers(context.Background(), groupUuid).Username(username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.DeleteUserGroupUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserGroupUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **username** | **[]string** | A list of usernames to remove from the user group in the format of username&#x3D;user1&amp;username&#x3D;user2 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccount

> Account GetAccount(ctx, accountName).Execute()

Get account info about this specific account.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	accountName := "accountName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.GetAccount(context.Background(), accountName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.GetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccount`: Account
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Account**](Account.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountUser

> User GetAccountUser(ctx, accountName, username).Execute()

Get a specific user in the specified account

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	accountName := "accountName_example" // string | 
	username := "username_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.GetAccountUser(context.Background(), accountName, username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.GetAccountUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountUser`: User
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.GetAccountUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** |  | 
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOauthToken

> TokenResponse GetOauthToken(ctx).GrantType(grantType).Username(username).Password(password).ClientId(clientId).RefreshToken(refreshToken).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	grantType := "grantType_example" // string | OAuth Grant type for token (optional) (default to "password")
	username := "username_example" // string | User to assign OAuth token to (optional)
	password := "password_example" // string | Password for corresponding user (optional)
	clientId := "clientId_example" // string | The type of client used for the OAuth token (optional) (default to "anonymous")
	refreshToken := "refreshToken_example" // string | The refresh token from a previous password grant request, used to get a new access_token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.GetOauthToken(context.Background()).GrantType(grantType).Username(username).Password(password).ClientId(clientId).RefreshToken(refreshToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.GetOauthToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOauthToken`: TokenResponse
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.GetOauthToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOauthTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **string** | OAuth Grant type for token | [default to &quot;password&quot;]
 **username** | **string** | User to assign OAuth token to | 
 **password** | **string** | Password for corresponding user | 
 **clientId** | **string** | The type of client used for the OAuth token | [default to &quot;anonymous&quot;]
 **refreshToken** | **string** | The refresh token from a previous password grant request, used to get a new access_token | 

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserApiKey

> UserApiKey GetUserApiKey(ctx, accountName, username, keyName).Execute()

Get a user API key

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	accountName := "accountName_example" // string | 
	username := "username_example" // string | 
	keyName := "keyName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.GetUserApiKey(context.Background(), accountName, username, keyName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.GetUserApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserApiKey`: UserApiKey
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.GetUserApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** |  | 
**username** | **string** |  | 
**keyName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**UserApiKey**](UserApiKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserGroup

> UserGroup GetUserGroup(ctx, groupUuid).Execute()

Get a user group

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	groupUuid := "groupUuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.GetUserGroup(context.Background(), groupUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.GetUserGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserGroup`: UserGroup
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.GetUserGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserGroup**](UserGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccounts

> []Account ListAccounts(ctx).State(state).Execute()

List account summaries. Only available to the system admin user.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	state := "state_example" // string | Filter accounts by state (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.ListAccounts(context.Background()).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.ListAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAccounts`: []Account
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.ListAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **string** | Filter accounts by state | 

### Return type

[**[]Account**](Account.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllSystemUsers

> Users ListAllSystemUsers(ctx).Execute()

List all system account users. Only available to `admin` users.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.ListAllSystemUsers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.ListAllSystemUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAllSystemUsers`: Users
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.ListAllSystemUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAllSystemUsersRequest struct via the builder pattern


### Return type

[**Users**](Users.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserApiKeys

> ApiKeyList ListUserApiKeys(ctx, accountName, username).Execute()

Get a list of API keys

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	accountName := "accountName_example" // string | 
	username := "username_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.ListUserApiKeys(context.Background(), accountName, username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.ListUserApiKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserApiKeys`: ApiKeyList
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.ListUserApiKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** |  | 
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserApiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiKeyList**](ApiKeyList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserCredentials

> []AccessCredential ListUserCredentials(ctx, accountName, username).Execute()

Get current credential summary

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	accountName := "accountName_example" // string | 
	username := "username_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.ListUserCredentials(context.Background(), accountName, username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.ListUserCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserCredentials`: []AccessCredential
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.ListUserCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** |  | 
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]AccessCredential**](AccessCredential.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserGroupRoles

> UserGroupRoles ListUserGroupRoles(ctx, groupUuid).Execute()

Get a list of all user group roles

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	groupUuid := "groupUuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.ListUserGroupRoles(context.Background(), groupUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.ListUserGroupRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserGroupRoles`: UserGroupRoles
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.ListUserGroupRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserGroupRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserGroupRoles**](UserGroupRoles.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserGroupUsers

> UserGroupUsers ListUserGroupUsers(ctx, groupUuid).Execute()

Get a list of user group users

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	groupUuid := "groupUuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.ListUserGroupUsers(context.Background(), groupUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.ListUserGroupUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserGroupUsers`: UserGroupUsers
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.ListUserGroupUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserGroupUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserGroupUsers**](UserGroupUsers.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserGroups

> []UserGroup ListUserGroups(ctx).ContainsUser(containsUser).UserGroupName(userGroupName).ContainsAccount(containsAccount).Execute()

List user groups

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	containsUser := "containsUser_example" // string | Filter the user groups to only those that contain the specified user (optional)
	userGroupName := "userGroupName_example" // string | Filter results to match the specified user group name (optional)
	containsAccount := "containsAccount_example" // string | Filter the results to only those that have roles in the specified account (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.ListUserGroups(context.Background()).ContainsUser(containsUser).UserGroupName(userGroupName).ContainsAccount(containsAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.ListUserGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserGroups`: []UserGroup
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.ListUserGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUserGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **containsUser** | **string** | Filter the user groups to only those that contain the specified user | 
 **userGroupName** | **string** | Filter results to match the specified user group name | 
 **containsAccount** | **string** | Filter the results to only those that have roles in the specified account | 

### Return type

[**[]UserGroup**](UserGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsers

> []User ListUsers(ctx, accountName).Execute()

List of users that are primary within this account. The response object will only contain roles for this account as well as any system roles.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	accountName := "accountName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.ListUsers(context.Background(), accountName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.ListUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUsers`: []User
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.ListUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsersWithRoles

> Users ListUsersWithRoles(ctx, accountName).Execute()

List of users who have rbac roles in this account.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	accountName := "accountName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.ListUsersWithRoles(context.Background(), accountName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.ListUsersWithRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUsersWithRoles`: Users
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.ListUsersWithRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUsersWithRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Users**](Users.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchUserApiKey

> UserApiKey PatchUserApiKey(ctx, accountName, username, keyName).PatchUserApiKeyRequest(patchUserApiKeyRequest).Execute()

Patch a user API key

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	accountName := "accountName_example" // string | 
	username := "username_example" // string | 
	keyName := "keyName_example" // string | 
	patchUserApiKeyRequest := *openapiclient.NewPatchUserApiKeyRequest() // PatchUserApiKeyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.PatchUserApiKey(context.Background(), accountName, username, keyName).PatchUserApiKeyRequest(patchUserApiKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.PatchUserApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchUserApiKey`: UserApiKey
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.PatchUserApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** |  | 
**username** | **string** |  | 
**keyName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchUserApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **patchUserApiKeyRequest** | [**PatchUserApiKeyRequest**](PatchUserApiKeyRequest.md) |  | 

### Return type

[**UserApiKey**](UserApiKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeOauthToken

> RevokeOauthToken(ctx).Token(token).TokenTypeHint(tokenTypeHint).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	token := "token_example" // string | The token to be revoked (optional)
	tokenTypeHint := "tokenTypeHint_example" // string | A hint about the type of token to be revoked (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserManagementAPI.RevokeOauthToken(context.Background()).Token(token).TokenTypeHint(tokenTypeHint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.RevokeOauthToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRevokeOauthTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | The token to be revoked | 
 **tokenTypeHint** | **string** | A hint about the type of token to be revoked | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccount

> Account UpdateAccount(ctx, accountName).Info(info).XAnchoreAccount(xAnchoreAccount).Execute()

Update the info for this specific account.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	accountName := "accountName_example" // string | 
	info := *openapiclient.NewAccountInfo() // AccountInfo | 
	xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.UpdateAccount(context.Background(), accountName).Info(info).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.UpdateAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccount`: Account
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.UpdateAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**AccountInfo**](AccountInfo.md) |  | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**Account**](Account.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountState

> AccountStatus UpdateAccountState(ctx, accountName).DesiredState(desiredState).Execute()

Update the state of an account to either enabled or disabled. For deletion use the DELETE route

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	accountName := "accountName_example" // string | 
	desiredState := *openapiclient.NewAccountStatus() // AccountStatus | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.UpdateAccountState(context.Background(), accountName).DesiredState(desiredState).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.UpdateAccountState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccountState`: AccountStatus
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.UpdateAccountState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **desiredState** | [**AccountStatus**](AccountStatus.md) |  | 

### Return type

[**AccountStatus**](AccountStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserGroup

> UserGroup UpdateUserGroup(ctx, groupUuid).UserGroupPatch(userGroupPatch).Execute()

Update a user group

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	groupUuid := "groupUuid_example" // string | 
	userGroupPatch := *openapiclient.NewUserGroupPatch("Description_example") // UserGroupPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.UpdateUserGroup(context.Background(), groupUuid).UserGroupPatch(userGroupPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.UpdateUserGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserGroup`: UserGroup
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.UpdateUserGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userGroupPatch** | [**UserGroupPatch**](UserGroupPatch.md) |  | 

### Return type

[**UserGroup**](UserGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

