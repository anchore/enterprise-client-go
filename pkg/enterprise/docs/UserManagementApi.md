# \UserManagementAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccount**](UserManagementAPI.md#CreateAccount) | **Post** /accounts | Create a new account. Only available to admin user.
[**CreateUser**](UserManagementAPI.md#CreateUser) | **Post** /accounts/{account_name}/users | Create a new user within the specified account.
[**CreateUserCredential**](UserManagementAPI.md#CreateUserCredential) | **Post** /accounts/{account_name}/users/{username}/credentials | add/replace credential
[**DeleteAccount**](UserManagementAPI.md#DeleteAccount) | **Delete** /accounts/{account_name} | Delete the specified account, only allowed if the account is in the disabled state. All users will be deleted along with the account and all resources will be garbage collected
[**DeleteUser**](UserManagementAPI.md#DeleteUser) | **Delete** /accounts/{account_name}/users/{username} | Delete a specific user credential by username of the credential. Cannot be the credential used to authenticate the request.
[**DeleteUserCredential**](UserManagementAPI.md#DeleteUserCredential) | **Delete** /accounts/{account_name}/users/{username}/credentials | Delete a credential by type
[**GetAccount**](UserManagementAPI.md#GetAccount) | **Get** /accounts/{account_name} | Get account info about this specific account.
[**GetAccountUser**](UserManagementAPI.md#GetAccountUser) | **Get** /accounts/{account_name}/users/{username} | Get a specific user in the specified account
[**ListAccounts**](UserManagementAPI.md#ListAccounts) | **Get** /accounts | List account summaries. Only available to the system admin user.
[**ListUserCredentials**](UserManagementAPI.md#ListUserCredentials) | **Get** /accounts/{account_name}/users/{username}/credentials | Get current credential summary
[**ListUsers**](UserManagementAPI.md#ListUsers) | **Get** /accounts/{account_name}/users | List of users found in this account.
[**UpdateAccount**](UserManagementAPI.md#UpdateAccount) | **Put** /accounts/{account_name} | Update the info for this specific account.
[**UpdateAccountState**](UserManagementAPI.md#UpdateAccountState) | **Put** /accounts/{account_name}/state | Update the state of an account to either enabled or disabled. For deletion use the DELETE route



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


## DeleteAccount

> DeleteAccount(ctx, accountName).Execute()

Delete the specified account, only allowed if the account is in the disabled state. All users will be deleted along with the account and all resources will be garbage collected

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

Delete a specific user credential by username of the credential. Cannot be the credential used to authenticate the request.

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


## ListUsers

> []User ListUsers(ctx, accountName).Execute()

List of users found in this account.

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

