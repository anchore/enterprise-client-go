# \RBACApi

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddIdp**](RBACApi.md#AddIdp) | **Post** /rbac-manager/saml/idps | 
[**AddIdpUserGroups**](RBACApi.md#AddIdpUserGroups) | **Post** /rbac-manager/saml/idps/{name}/user-group-mappings | 
[**AddRoleUser**](RBACApi.md#AddRoleUser) | **Post** /rbac-manager/roles/{role_name}/members | Add a user to the role
[**DeleteIdp**](RBACApi.md#DeleteIdp) | **Delete** /rbac-manager/saml/idps/{name} | 
[**DeleteIdpUserGroup**](RBACApi.md#DeleteIdpUserGroup) | **Delete** /rbac-manager/saml/idps/{name}/user-group-mappings | 
[**DeleteRoleUser**](RBACApi.md#DeleteRoleUser) | **Delete** /rbac-manager/roles/{role_name}/members | Remove a user from the role
[**GetIdp**](RBACApi.md#GetIdp) | **Get** /rbac-manager/saml/idps/{name} | 
[**GetIdpUserGroups**](RBACApi.md#GetIdpUserGroups) | **Get** /rbac-manager/saml/idps/{name}/user-group-mappings | 
[**GetRole**](RBACApi.md#GetRole) | **Get** /rbac-manager/roles/{role_name} | Get detailed information about a specific role
[**ListIdps**](RBACApi.md#ListIdps) | **Get** /rbac-manager/saml/idps | 
[**ListRoleMembers**](RBACApi.md#ListRoleMembers) | **Get** /rbac-manager/roles/{role_name}/members | Returns a list of objects that have members in the role. The list is filtered by &#39;listRoleMembers&#39; access for the &#39;account&#39; element of each entry.
[**ListRoles**](RBACApi.md#ListRoles) | **Get** /rbac-manager/roles | List roles available in the system
[**ListUserRoles**](RBACApi.md#ListUserRoles) | **Get** /rbac-manager/users/{username}/roles | List the roles for which the requested user is a member
[**MyRoles**](RBACApi.md#MyRoles) | **Get** /rbac-manager/my-roles | List the roles for which the authenticated user is a member
[**SamlLogin**](RBACApi.md#SamlLogin) | **Get** /rbac-manager/saml/login/{idp_name} | 
[**SamlSso**](RBACApi.md#SamlSso) | **Post** /rbac-manager/saml/sso/{idp_name} | 
[**UpdateIdp**](RBACApi.md#UpdateIdp) | **Put** /rbac-manager/saml/idps/{name} | 



## AddIdp

> RbacManagerSamlConfiguration AddIdp(ctx).Configuration(configuration).Execute()





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
    configuration := *openapiclient.NewRbacManagerSamlConfiguration("Name_example", false, "SpEntityId_example", "AcsUrl_example") // RbacManagerSamlConfiguration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RBACApi.AddIdp(context.Background()).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RBACApi.AddIdp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddIdp`: RbacManagerSamlConfiguration
    fmt.Fprintf(os.Stdout, "Response from `RBACApi.AddIdp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddIdpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configuration** | [**RbacManagerSamlConfiguration**](RbacManagerSamlConfiguration.md) |  | 

### Return type

[**RbacManagerSamlConfiguration**](RbacManagerSamlConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddIdpUserGroups

> []RbacManagerIdpUserGroup AddIdpUserGroups(ctx, name).RbacManagerIdpUserGroupPost(rbacManagerIdpUserGroupPost).Execute()





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
    name := "name_example" // string | 
    rbacManagerIdpUserGroupPost := *openapiclient.NewRbacManagerIdpUserGroupPost() // RbacManagerIdpUserGroupPost | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RBACApi.AddIdpUserGroups(context.Background(), name).RbacManagerIdpUserGroupPost(rbacManagerIdpUserGroupPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RBACApi.AddIdpUserGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddIdpUserGroups`: []RbacManagerIdpUserGroup
    fmt.Fprintf(os.Stdout, "Response from `RBACApi.AddIdpUserGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddIdpUserGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rbacManagerIdpUserGroupPost** | [**RbacManagerIdpUserGroupPost**](RbacManagerIdpUserGroupPost.md) |  | 

### Return type

[**[]RbacManagerIdpUserGroup**](RbacManagerIdpUserGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddRoleUser

> RbacManagerRoleMember AddRoleUser(ctx, roleName).Member(member).Execute()

Add a user to the role

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
    roleName := "roleName_example" // string | 
    member := *openapiclient.NewRbacManagerRoleMember("Username_example", "ForAccount_example") // RbacManagerRoleMember | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RBACApi.AddRoleUser(context.Background(), roleName).Member(member).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RBACApi.AddRoleUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRoleUser`: RbacManagerRoleMember
    fmt.Fprintf(os.Stdout, "Response from `RBACApi.AddRoleUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddRoleUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **member** | [**RbacManagerRoleMember**](RbacManagerRoleMember.md) |  | 

### Return type

[**RbacManagerRoleMember**](RbacManagerRoleMember.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdp

> DeleteIdp(ctx, name).Execute()





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
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RBACApi.DeleteIdp(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RBACApi.DeleteIdp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdpRequest struct via the builder pattern


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


## DeleteIdpUserGroup

> DeleteIdpUserGroup(ctx, name).UserGroup(userGroup).Execute()





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
    name := "name_example" // string | 
    userGroup := []string{"Inner_example"} // []string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RBACApi.DeleteIdpUserGroup(context.Background(), name).UserGroup(userGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RBACApi.DeleteIdpUserGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdpUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userGroup** | **[]string** |  | 

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


## DeleteRoleUser

> DeleteRoleUser(ctx, roleName).Username(username).ForAccount(forAccount).Execute()

Remove a user from the role

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
    roleName := "roleName_example" // string | 
    username := "username_example" // string | The username to remove the role for
    forAccount := "forAccount_example" // string | The account that the user has the role to be removed

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RBACApi.DeleteRoleUser(context.Background(), roleName).Username(username).ForAccount(forAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RBACApi.DeleteRoleUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **username** | **string** | The username to remove the role for | 
 **forAccount** | **string** | The account that the user has the role to be removed | 

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


## GetIdp

> RbacManagerSamlConfigurationGet GetIdp(ctx, name).Execute()





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
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RBACApi.GetIdp(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RBACApi.GetIdp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdp`: RbacManagerSamlConfigurationGet
    fmt.Fprintf(os.Stdout, "Response from `RBACApi.GetIdp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RbacManagerSamlConfigurationGet**](RbacManagerSamlConfigurationGet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdpUserGroups

> []RbacManagerIdpUserGroup GetIdpUserGroups(ctx, name).Execute()





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
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RBACApi.GetIdpUserGroups(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RBACApi.GetIdpUserGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdpUserGroups`: []RbacManagerIdpUserGroup
    fmt.Fprintf(os.Stdout, "Response from `RBACApi.GetIdpUserGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdpUserGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]RbacManagerIdpUserGroup**](RbacManagerIdpUserGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRole

> RbacManagerRole GetRole(ctx, roleName).Execute()

Get detailed information about a specific role

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
    roleName := "roleName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RBACApi.GetRole(context.Background(), roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RBACApi.GetRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRole`: RbacManagerRole
    fmt.Fprintf(os.Stdout, "Response from `RBACApi.GetRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RbacManagerRole**](RbacManagerRole.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdps

> []string ListIdps(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RBACApi.ListIdps(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RBACApi.ListIdps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIdps`: []string
    fmt.Fprintf(os.Stdout, "Response from `RBACApi.ListIdps`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListIdpsRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoleMembers

> []RbacManagerRoleMember ListRoleMembers(ctx, roleName).ForAccount(forAccount).Execute()

Returns a list of objects that have members in the role. The list is filtered by 'listRoleMembers' access for the 'account' element of each entry.

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
    roleName := "roleName_example" // string | 
    forAccount := "forAccount_example" // string | Optional filter parameter to limit the set fo returned items to only those with matching account. Will return Access Denied if caller does not have permission to listRoleMembers for that account. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RBACApi.ListRoleMembers(context.Background(), roleName).ForAccount(forAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RBACApi.ListRoleMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoleMembers`: []RbacManagerRoleMember
    fmt.Fprintf(os.Stdout, "Response from `RBACApi.ListRoleMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRoleMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forAccount** | **string** | Optional filter parameter to limit the set fo returned items to only those with matching account. Will return Access Denied if caller does not have permission to listRoleMembers for that account. | 

### Return type

[**[]RbacManagerRoleMember**](RbacManagerRoleMember.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoles

> []RbacManagerRoleSummary ListRoles(ctx).Execute()

List roles available in the system

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RBACApi.ListRoles(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RBACApi.ListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoles`: []RbacManagerRoleSummary
    fmt.Fprintf(os.Stdout, "Response from `RBACApi.ListRoles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRolesRequest struct via the builder pattern


### Return type

[**[]RbacManagerRoleSummary**](RbacManagerRoleSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserRoles

> []RbacManagerRoleMembership ListUserRoles(ctx, username).ForAccount(forAccount).Role(role).Execute()

List the roles for which the requested user is a member

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
    username := "username_example" // string | 
    forAccount := "forAccount_example" // string |  (optional)
    role := "role_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RBACApi.ListUserRoles(context.Background(), username).ForAccount(forAccount).Role(role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RBACApi.ListUserRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserRoles`: []RbacManagerRoleMembership
    fmt.Fprintf(os.Stdout, "Response from `RBACApi.ListUserRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forAccount** | **string** |  | 
 **role** | **string** |  | 

### Return type

[**[]RbacManagerRoleMembership**](RbacManagerRoleMembership.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MyRoles

> []RbacManagerAccountRole MyRoles(ctx).Execute()

List the roles for which the authenticated user is a member

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RBACApi.MyRoles(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RBACApi.MyRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MyRoles`: []RbacManagerAccountRole
    fmt.Fprintf(os.Stdout, "Response from `RBACApi.MyRoles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMyRolesRequest struct via the builder pattern


### Return type

[**[]RbacManagerAccountRole**](RbacManagerAccountRole.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SamlLogin

> RbacManagerTokenResponse SamlLogin(ctx, idpName).Execute()





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
    idpName := "idpName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RBACApi.SamlLogin(context.Background(), idpName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RBACApi.SamlLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SamlLogin`: RbacManagerTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `RBACApi.SamlLogin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSamlLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RbacManagerTokenResponse**](RbacManagerTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SamlSso

> RbacManagerTokenResponse SamlSso(ctx, idpName).Execute()





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
    idpName := "idpName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RBACApi.SamlSso(context.Background(), idpName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RBACApi.SamlSso``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SamlSso`: RbacManagerTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `RBACApi.SamlSso`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSamlSsoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RbacManagerTokenResponse**](RbacManagerTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdp

> RbacManagerSamlConfiguration UpdateIdp(ctx, name).RbacManagerSamlConfiguration(rbacManagerSamlConfiguration).Execute()





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
    name := "name_example" // string | 
    rbacManagerSamlConfiguration := *openapiclient.NewRbacManagerSamlConfiguration("Name_example", false, "SpEntityId_example", "AcsUrl_example") // RbacManagerSamlConfiguration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RBACApi.UpdateIdp(context.Background(), name).RbacManagerSamlConfiguration(rbacManagerSamlConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RBACApi.UpdateIdp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIdp`: RbacManagerSamlConfiguration
    fmt.Fprintf(os.Stdout, "Response from `RBACApi.UpdateIdp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rbacManagerSamlConfiguration** | [**RbacManagerSamlConfiguration**](RbacManagerSamlConfiguration.md) |  | 

### Return type

[**RbacManagerSamlConfiguration**](RbacManagerSamlConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

