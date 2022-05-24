# \UserManagementApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccount**](UserManagementApi.md#CreateAccount) | **Post** /accounts | Create a new user. Only avaialble to admin user.
[**CreateUser**](UserManagementApi.md#CreateUser) | **Post** /accounts/{accountname}/users | Create a new user
[**CreateUserCredential**](UserManagementApi.md#CreateUserCredential) | **Post** /accounts/{accountname}/users/{username}/credentials | add/replace credential
[**DeleteAccount**](UserManagementApi.md#DeleteAccount) | **Delete** /accounts/{accountname} | Delete the specified account, only allowed if the account is in the disabled state. All users will be deleted along with the account and all resources will be garbage collected
[**DeleteUser**](UserManagementApi.md#DeleteUser) | **Delete** /accounts/{accountname}/users/{username} | Delete a specific user credential by username of the credential. Cannot be the credential used to authenticate the request.
[**DeleteUserCredential**](UserManagementApi.md#DeleteUserCredential) | **Delete** /accounts/{accountname}/users/{username}/credentials | Delete a credential by type
[**GetAccount**](UserManagementApi.md#GetAccount) | **Get** /accounts/{accountname} | Get info about an user. Only available to admin user. Uses the main user Id, not a username.
[**GetAccountUser**](UserManagementApi.md#GetAccountUser) | **Get** /accounts/{accountname}/users/{username} | Get a specific user in the specified account
[**ListAccounts**](UserManagementApi.md#ListAccounts) | **Get** /accounts | List user summaries. Only available to the system admin user.
[**ListUserCredentials**](UserManagementApi.md#ListUserCredentials) | **Get** /accounts/{accountname}/users/{username}/credentials | Get current credential summary
[**ListUsers**](UserManagementApi.md#ListUsers) | **Get** /accounts/{accountname}/users | List accounts for the user
[**UpdateAccountState**](UserManagementApi.md#UpdateAccountState) | **Put** /accounts/{accountname}/state | Update the state of an account to either enabled or disabled. For deletion use the DELETE route



## CreateAccount

> Account CreateAccount(ctx).Account(account).Execute()

Create a new user. Only avaialble to admin user.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    account := *openapiclient.NewAccountCreationRequest("Name_example") // AccountCreationRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserManagementApi.CreateAccount(context.Background()).Account(account).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementApi.CreateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccount`: Account
    fmt.Fprintf(os.Stdout, "Response from `UserManagementApi.CreateAccount`: %v\n", resp)
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

> User CreateUser(ctx, accountname).User(user).Execute()

Create a new user

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountname := "accountname_example" // string | 
    user := *openapiclient.NewUserCreationRequest("Username_example", "Password_example") // UserCreationRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserManagementApi.CreateUser(context.Background(), accountname).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementApi.CreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUser`: User
    fmt.Fprintf(os.Stdout, "Response from `UserManagementApi.CreateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountname** | **string** |  | 

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


## CreateUserCredential

> User CreateUserCredential(ctx, accountname, username).Credential(credential).Execute()

add/replace credential

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountname := "accountname_example" // string | 
    username := "username_example" // string | 
    credential := *openapiclient.NewAccessCredential("Type_example", "Value_example") // AccessCredential | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserManagementApi.CreateUserCredential(context.Background(), accountname, username).Credential(credential).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementApi.CreateUserCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUserCredential`: User
    fmt.Fprintf(os.Stdout, "Response from `UserManagementApi.CreateUserCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountname** | **string** |  | 
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **credential** | [**AccessCredential**](AccessCredential.md) |  | 

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


## DeleteAccount

> DeleteAccount(ctx, accountname).Execute()

Delete the specified account, only allowed if the account is in the disabled state. All users will be deleted along with the account and all resources will be garbage collected

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountname := "accountname_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserManagementApi.DeleteAccount(context.Background(), accountname).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementApi.DeleteAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountname** | **string** |  | 

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

> DeleteUser(ctx, accountname, username).Execute()

Delete a specific user credential by username of the credential. Cannot be the credential used to authenticate the request.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountname := "accountname_example" // string | 
    username := "username_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserManagementApi.DeleteUser(context.Background(), accountname, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementApi.DeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountname** | **string** |  | 
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


## DeleteUserCredential

> DeleteUserCredential(ctx, accountname, username).CredentialType(credentialType).Execute()

Delete a credential by type

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountname := "accountname_example" // string | 
    username := "username_example" // string | 
    credentialType := "credentialType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserManagementApi.DeleteUserCredential(context.Background(), accountname, username).CredentialType(credentialType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementApi.DeleteUserCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountname** | **string** |  | 
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


## GetAccount

> Account GetAccount(ctx, accountname).Execute()

Get info about an user. Only available to admin user. Uses the main user Id, not a username.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountname := "accountname_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserManagementApi.GetAccount(context.Background(), accountname).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementApi.GetAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccount`: Account
    fmt.Fprintf(os.Stdout, "Response from `UserManagementApi.GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountname** | **string** |  | 

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

> User GetAccountUser(ctx, accountname, username).Execute()

Get a specific user in the specified account

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountname := "accountname_example" // string | 
    username := "username_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserManagementApi.GetAccountUser(context.Background(), accountname, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementApi.GetAccountUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountUser`: User
    fmt.Fprintf(os.Stdout, "Response from `UserManagementApi.GetAccountUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountname** | **string** |  | 
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


## ListAccounts

> []Account ListAccounts(ctx).State(state).Execute()

List user summaries. Only available to the system admin user.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    state := "state_example" // string | Filter accounts by state (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserManagementApi.ListAccounts(context.Background()).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementApi.ListAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccounts`: []Account
    fmt.Fprintf(os.Stdout, "Response from `UserManagementApi.ListAccounts`: %v\n", resp)
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


## ListUserCredentials

> []AccessCredential ListUserCredentials(ctx, accountname, username).Execute()

Get current credential summary

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountname := "accountname_example" // string | 
    username := "username_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserManagementApi.ListUserCredentials(context.Background(), accountname, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementApi.ListUserCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserCredentials`: []AccessCredential
    fmt.Fprintf(os.Stdout, "Response from `UserManagementApi.ListUserCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountname** | **string** |  | 
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


## ListUsers

> []User ListUsers(ctx, accountname).Execute()

List accounts for the user

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountname := "accountname_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserManagementApi.ListUsers(context.Background(), accountname).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementApi.ListUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsers`: []User
    fmt.Fprintf(os.Stdout, "Response from `UserManagementApi.ListUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountname** | **string** |  | 

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


## UpdateAccountState

> AccountStatus UpdateAccountState(ctx, accountname).DesiredState(desiredState).Execute()

Update the state of an account to either enabled or disabled. For deletion use the DELETE route

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountname := "accountname_example" // string | 
    desiredState := *openapiclient.NewAccountStatus() // AccountStatus | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserManagementApi.UpdateAccountState(context.Background(), accountname).DesiredState(desiredState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementApi.UpdateAccountState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccountState`: AccountStatus
    fmt.Fprintf(os.Stdout, "Response from `UserManagementApi.UpdateAccountState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountname** | **string** |  | 

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

