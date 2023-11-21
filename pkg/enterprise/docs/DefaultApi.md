# \DefaultApi

<<<<<<< HEAD
All URIs are relative to */v2*
=======
All URIs are relative to *http://localhost/v2*
>>>>>>> main

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGithubConfiguration**](DefaultApi.md#AddGithubConfiguration) | **Post** /notifications/endpoints/github/configurations | 
[**AddGithubSelector**](DefaultApi.md#AddGithubSelector) | **Post** /notifications/endpoints/github/configurations/{uuid}/selectors | 
[**AddIdp**](DefaultApi.md#AddIdp) | **Post** /rbac-manager/saml/idps | 
[**AddJiraConfiguration**](DefaultApi.md#AddJiraConfiguration) | **Post** /notifications/endpoints/jira/configurations | 
[**AddJiraSelector**](DefaultApi.md#AddJiraSelector) | **Post** /notifications/endpoints/jira/configurations/{uuid}/selectors | 
[**AddRoleUser**](DefaultApi.md#AddRoleUser) | **Post** /rbac-manager/roles/{role_name}/members | Add a user to the role
[**AddSlackConfiguration**](DefaultApi.md#AddSlackConfiguration) | **Post** /notifications/endpoints/slack/configurations | 
[**AddSlackSelector**](DefaultApi.md#AddSlackSelector) | **Post** /notifications/endpoints/slack/configurations/{uuid}/selectors | 
[**AddSmtpConfiguration**](DefaultApi.md#AddSmtpConfiguration) | **Post** /notifications/endpoints/smtp/configurations | 
[**AddSmtpSelector**](DefaultApi.md#AddSmtpSelector) | **Post** /notifications/endpoints/smtp/configurations/{uuid}/selectors | 
[**AddTeamsConfiguration**](DefaultApi.md#AddTeamsConfiguration) | **Post** /notifications/endpoints/teams/configurations | 
[**AddTeamsSelector**](DefaultApi.md#AddTeamsSelector) | **Post** /notifications/endpoints/teams/configurations/{uuid}/selectors | 
[**AddWebhookConfiguration**](DefaultApi.md#AddWebhookConfiguration) | **Post** /notifications/endpoints/webhook/configurations | 
[**AddWebhookSelector**](DefaultApi.md#AddWebhookSelector) | **Post** /notifications/endpoints/webhook/configurations/{uuid}/selectors | 
[**DeleteGithubConfiguration**](DefaultApi.md#DeleteGithubConfiguration) | **Delete** /notifications/endpoints/github/configurations/{uuid} | 
[**DeleteGithubSelector**](DefaultApi.md#DeleteGithubSelector) | **Delete** /notifications/endpoints/github/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**DeleteIdp**](DefaultApi.md#DeleteIdp) | **Delete** /rbac-manager/saml/idps/{name} | 
[**DeleteJiraConfiguration**](DefaultApi.md#DeleteJiraConfiguration) | **Delete** /notifications/endpoints/jira/configurations/{uuid} | 
[**DeleteJiraSelector**](DefaultApi.md#DeleteJiraSelector) | **Delete** /notifications/endpoints/jira/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**DeleteRoleUser**](DefaultApi.md#DeleteRoleUser) | **Delete** /rbac-manager/roles/{role_name}/members | Remove a user from the role
[**DeleteSlackConfiguration**](DefaultApi.md#DeleteSlackConfiguration) | **Delete** /notifications/endpoints/slack/configurations/{uuid} | 
[**DeleteSlackSelector**](DefaultApi.md#DeleteSlackSelector) | **Delete** /notifications/endpoints/slack/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**DeleteSmtpConfiguration**](DefaultApi.md#DeleteSmtpConfiguration) | **Delete** /notifications/endpoints/smtp/configurations/{uuid} | 
[**DeleteSmtpSelector**](DefaultApi.md#DeleteSmtpSelector) | **Delete** /notifications/endpoints/smtp/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**DeleteTeamsConfiguration**](DefaultApi.md#DeleteTeamsConfiguration) | **Delete** /notifications/endpoints/teams/configurations/{uuid} | 
[**DeleteTeamsSelector**](DefaultApi.md#DeleteTeamsSelector) | **Delete** /notifications/endpoints/teams/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**DeleteWebhookConfiguration**](DefaultApi.md#DeleteWebhookConfiguration) | **Delete** /notifications/endpoints/webhook/configurations/{uuid} | 
[**DeleteWebhookSelector**](DefaultApi.md#DeleteWebhookSelector) | **Delete** /notifications/endpoints/webhook/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**GetGithubConfiguration**](DefaultApi.md#GetGithubConfiguration) | **Get** /notifications/endpoints/github/configurations/{uuid} | 
[**GetGithubConfigurationStatus**](DefaultApi.md#GetGithubConfigurationStatus) | **Get** /notifications/endpoints/github/configurations/{uuid}/status | 
[**GetGithubSelector**](DefaultApi.md#GetGithubSelector) | **Get** /notifications/endpoints/github/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**GetGlobalQueryResult**](DefaultApi.md#GetGlobalQueryResult) | **Get** /reporting/reports/global/scheduled-query-results/{result_uuid} | 
[**GetIdp**](DefaultApi.md#GetIdp) | **Get** /rbac-manager/saml/idps/{name} | 
[**GetJiraConfiguration**](DefaultApi.md#GetJiraConfiguration) | **Get** /notifications/endpoints/jira/configurations/{uuid} | 
[**GetJiraConfigurationStatus**](DefaultApi.md#GetJiraConfigurationStatus) | **Get** /notifications/endpoints/jira/configurations/{uuid}/status | 
[**GetJiraSelector**](DefaultApi.md#GetJiraSelector) | **Get** /notifications/endpoints/jira/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**GetOauthToken**](DefaultApi.md#GetOauthToken) | **Post** /oauth/token | 
[**GetQueryResult**](DefaultApi.md#GetQueryResult) | **Get** /reporting/scheduled-query-results/{result_uuid} | 
[**GetRole**](DefaultApi.md#GetRole) | **Get** /rbac-manager/roles/{role_name} | Get detailed information about a specific role
[**GetSlackConfiguration**](DefaultApi.md#GetSlackConfiguration) | **Get** /notifications/endpoints/slack/configurations/{uuid} | 
[**GetSlackConfigurationStatus**](DefaultApi.md#GetSlackConfigurationStatus) | **Get** /notifications/endpoints/slack/configurations/{uuid}/status | 
[**GetSlackSelector**](DefaultApi.md#GetSlackSelector) | **Get** /notifications/endpoints/slack/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**GetSmtpConfiguration**](DefaultApi.md#GetSmtpConfiguration) | **Get** /notifications/endpoints/smtp/configurations/{uuid} | 
[**GetSmtpConfigurationStatus**](DefaultApi.md#GetSmtpConfigurationStatus) | **Get** /notifications/endpoints/smtp/configurations/{uuid}/status | 
[**GetSmtpSelector**](DefaultApi.md#GetSmtpSelector) | **Get** /notifications/endpoints/smtp/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**GetTeamsConfiguration**](DefaultApi.md#GetTeamsConfiguration) | **Get** /notifications/endpoints/teams/configurations/{uuid} | 
[**GetTeamsConfigurationStatus**](DefaultApi.md#GetTeamsConfigurationStatus) | **Get** /notifications/endpoints/teams/configurations/{uuid}/status | 
[**GetTeamsSelector**](DefaultApi.md#GetTeamsSelector) | **Get** /notifications/endpoints/teams/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**GetWebhookConfiguration**](DefaultApi.md#GetWebhookConfiguration) | **Get** /notifications/endpoints/webhook/configurations/{uuid} | 
[**GetWebhookConfigurationStatus**](DefaultApi.md#GetWebhookConfigurationStatus) | **Get** /notifications/endpoints/webhook/configurations/{uuid}/status | 
[**GetWebhookSelector**](DefaultApi.md#GetWebhookSelector) | **Get** /notifications/endpoints/webhook/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**HealthCheck**](DefaultApi.md#HealthCheck) | **Get** /health | 
[**ListEndpoints**](DefaultApi.md#ListEndpoints) | **Get** /notifications/endpoints | 
[**ListFileContentSearchResults**](DefaultApi.md#ListFileContentSearchResults) | **Get** /images/{image_digest}/artifacts/file-content-search | Return a list of analyzer artifacts of the specified type
[**ListGithubConfigurations**](DefaultApi.md#ListGithubConfigurations) | **Get** /notifications/endpoints/github/configurations | 
[**ListGithubSelectors**](DefaultApi.md#ListGithubSelectors) | **Get** /notifications/endpoints/github/configurations/{uuid}/selectors | 
[**ListIdps**](DefaultApi.md#ListIdps) | **Get** /rbac-manager/saml/idps | 
[**ListJiraConfigurations**](DefaultApi.md#ListJiraConfigurations) | **Get** /notifications/endpoints/jira/configurations | 
[**ListJiraSelectors**](DefaultApi.md#ListJiraSelectors) | **Get** /notifications/endpoints/jira/configurations/{uuid}/selectors | 
[**ListRetrievedFiles**](DefaultApi.md#ListRetrievedFiles) | **Get** /images/{image_digest}/artifacts/retrieved-files | Return a list of analyzer artifacts of the specified type
[**ListRoleMembers**](DefaultApi.md#ListRoleMembers) | **Get** /rbac-manager/roles/{role_name}/members | Returns a list of objects that have members in the role. The list is filtered by &#39;listRoleMembers&#39; access for the &#39;account&#39; element of each entry.
[**ListRoles**](DefaultApi.md#ListRoles) | **Get** /rbac-manager/roles | List roles available in the system
[**ListSecretSearchResults**](DefaultApi.md#ListSecretSearchResults) | **Get** /images/{image_digest}/artifacts/secret-search | Return a list of analyzer artifacts of the specified type
[**ListSelectors**](DefaultApi.md#ListSelectors) | **Get** /notifications/selectors | 
[**ListSlackConfigurations**](DefaultApi.md#ListSlackConfigurations) | **Get** /notifications/endpoints/slack/configurations | 
[**ListSlackSelectors**](DefaultApi.md#ListSlackSelectors) | **Get** /notifications/endpoints/slack/configurations/{uuid}/selectors | 
[**ListSmtpConfigurations**](DefaultApi.md#ListSmtpConfigurations) | **Get** /notifications/endpoints/smtp/configurations | 
[**ListSmtpSelectors**](DefaultApi.md#ListSmtpSelectors) | **Get** /notifications/endpoints/smtp/configurations/{uuid}/selectors | 
[**ListTeamsConfigurations**](DefaultApi.md#ListTeamsConfigurations) | **Get** /notifications/endpoints/teams/configurations | 
[**ListTeamsSelectors**](DefaultApi.md#ListTeamsSelectors) | **Get** /notifications/endpoints/teams/configurations/{uuid}/selectors | 
[**ListUserRoles**](DefaultApi.md#ListUserRoles) | **Get** /rbac-manager/users/{username}/roles | List the roles for which the requested user is a member
[**ListWebhookConfigurations**](DefaultApi.md#ListWebhookConfigurations) | **Get** /notifications/endpoints/webhook/configurations | 
[**ListWebhookSelectors**](DefaultApi.md#ListWebhookSelectors) | **Get** /notifications/endpoints/webhook/configurations/{uuid}/selectors | 
[**MyRoles**](DefaultApi.md#MyRoles) | **Get** /rbac-manager/my-roles | List the roles for which the authenticated user is a member
[**Ping**](DefaultApi.md#Ping) | **Get** / | 
[**RevokeOauthToken**](DefaultApi.md#RevokeOauthToken) | **Post** /oauth/revoke | 
[**SamlLogin**](DefaultApi.md#SamlLogin) | **Get** /rbac-manager/saml/login/{idp_name} | 
[**SamlSso**](DefaultApi.md#SamlSso) | **Post** /rbac-manager/saml/sso/{idp_name} | 
[**TestGithubConfiguration**](DefaultApi.md#TestGithubConfiguration) | **Post** /notifications/endpoints/github/test | 
[**TestJiraConfiguration**](DefaultApi.md#TestJiraConfiguration) | **Post** /notifications/endpoints/jira/test | 
[**TestSlackConfiguration**](DefaultApi.md#TestSlackConfiguration) | **Post** /notifications/endpoints/slack/test | 
[**TestSmtpConfiguration**](DefaultApi.md#TestSmtpConfiguration) | **Post** /notifications/endpoints/smtp/test | 
[**TestStoredGithubConfiguration**](DefaultApi.md#TestStoredGithubConfiguration) | **Get** /notifications/endpoints/github/configurations/{uuid}/test | 
[**TestStoredJiraConfiguration**](DefaultApi.md#TestStoredJiraConfiguration) | **Get** /notifications/endpoints/jira/configurations/{uuid}/test | 
[**TestStoredSlackConfiguration**](DefaultApi.md#TestStoredSlackConfiguration) | **Get** /notifications/endpoints/slack/configurations/{uuid}/test | 
[**TestStoredSmtpConfiguration**](DefaultApi.md#TestStoredSmtpConfiguration) | **Get** /notifications/endpoints/smtp/configurations/{uuid}/test | 
[**TestStoredTeamsConfiguration**](DefaultApi.md#TestStoredTeamsConfiguration) | **Get** /notifications/endpoints/teams/configurations/{uuid}/test | 
[**TestStoredWebhookConfiguration**](DefaultApi.md#TestStoredWebhookConfiguration) | **Get** /notifications/endpoints/webhook/configurations/{uuid}/test | 
[**TestTeamsConfiguration**](DefaultApi.md#TestTeamsConfiguration) | **Post** /notifications/endpoints/teams/test | 
[**TestWebhookConfiguration**](DefaultApi.md#TestWebhookConfiguration) | **Post** /notifications/endpoints/webhook/test | 
[**UpdateEndpointStatus**](DefaultApi.md#UpdateEndpointStatus) | **Put** /notifications/endpoints/{name} | 
[**UpdateGithubConfiguration**](DefaultApi.md#UpdateGithubConfiguration) | **Put** /notifications/endpoints/github/configurations/{uuid} | 
[**UpdateGithubSelector**](DefaultApi.md#UpdateGithubSelector) | **Put** /notifications/endpoints/github/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**UpdateIdp**](DefaultApi.md#UpdateIdp) | **Put** /rbac-manager/saml/idps/{name} | 
[**UpdateJiraConfiguration**](DefaultApi.md#UpdateJiraConfiguration) | **Put** /notifications/endpoints/jira/configurations/{uuid} | 
[**UpdateJiraSelector**](DefaultApi.md#UpdateJiraSelector) | **Put** /notifications/endpoints/jira/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**UpdateSlackConfiguration**](DefaultApi.md#UpdateSlackConfiguration) | **Put** /notifications/endpoints/slack/configurations/{uuid} | 
[**UpdateSlackSelector**](DefaultApi.md#UpdateSlackSelector) | **Put** /notifications/endpoints/slack/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**UpdateSmtpConfiguration**](DefaultApi.md#UpdateSmtpConfiguration) | **Put** /notifications/endpoints/smtp/configurations/{uuid} | 
[**UpdateSmtpSelector**](DefaultApi.md#UpdateSmtpSelector) | **Put** /notifications/endpoints/smtp/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**UpdateTeamsConfiguration**](DefaultApi.md#UpdateTeamsConfiguration) | **Put** /notifications/endpoints/teams/configurations/{uuid} | 
[**UpdateTeamsSelector**](DefaultApi.md#UpdateTeamsSelector) | **Put** /notifications/endpoints/teams/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**UpdateWebhookConfiguration**](DefaultApi.md#UpdateWebhookConfiguration) | **Put** /notifications/endpoints/webhook/configurations/{uuid} | 
[**UpdateWebhookSelector**](DefaultApi.md#UpdateWebhookSelector) | **Put** /notifications/endpoints/webhook/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**VersionCheck**](DefaultApi.md#VersionCheck) | **Get** /version | 



## AddGithubConfiguration

> NotificationGitHubEndpointConfigurationBase AddGithubConfiguration(ctx).Configuration(configuration).Execute()





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
    configuration := *openapiclient.NewNotificationGitHubEndpointConfigurationPost("Username_example", "AccessToken_example", "Owner_example", "Repository_example") // NotificationGitHubEndpointConfigurationPost | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AddGithubConfiguration(context.Background()).Configuration(configuration).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AddGithubConfiguration(context.Background()).Configuration(configuration).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddGithubConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddGithubConfiguration`: NotificationGitHubEndpointConfigurationBase
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddGithubConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddGithubConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configuration** | [**NotificationGitHubEndpointConfigurationPost**](NotificationGitHubEndpointConfigurationPost.md) |  | 

### Return type

[**NotificationGitHubEndpointConfigurationBase**](NotificationGitHubEndpointConfigurationBase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddGithubSelector

> NotificationSelector AddGithubSelector(ctx, uuid).Selector(selector).Execute()





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
    uuid := "uuid_example" // string | 
    selector := *openapiclient.NewNotificationSelector("Scope_example", *openapiclient.NewNotificationEventSelector("Level_example", "ResourceType_example", "Type_example")) // NotificationSelector | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AddGithubSelector(context.Background(), uuid).Selector(selector).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AddGithubSelector(context.Background(), uuid).Selector(selector).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddGithubSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddGithubSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddGithubSelector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddGithubSelectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **selector** | [**NotificationSelector**](NotificationSelector.md) |  | 

### Return type

[**NotificationSelector**](NotificationSelector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AddIdp(context.Background()).Configuration(configuration).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AddIdp(context.Background()).Configuration(configuration).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddIdp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddIdp`: RbacManagerSamlConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddIdp`: %v\n", resp)
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


## AddJiraConfiguration

> NotificationJiraEndpointConfigurationBase AddJiraConfiguration(ctx).Configuration(configuration).Execute()





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
    configuration := *openapiclient.NewNotificationJiraEndpointConfigurationPost("Url_example", "Username_example", "Password_example", "ProjectKey_example", "IssueType_example") // NotificationJiraEndpointConfigurationPost | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AddJiraConfiguration(context.Background()).Configuration(configuration).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AddJiraConfiguration(context.Background()).Configuration(configuration).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddJiraConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddJiraConfiguration`: NotificationJiraEndpointConfigurationBase
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddJiraConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddJiraConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configuration** | [**NotificationJiraEndpointConfigurationPost**](NotificationJiraEndpointConfigurationPost.md) |  | 

### Return type

[**NotificationJiraEndpointConfigurationBase**](NotificationJiraEndpointConfigurationBase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddJiraSelector

> NotificationSelector AddJiraSelector(ctx, uuid).Selector(selector).Execute()





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
    uuid := "uuid_example" // string | 
    selector := *openapiclient.NewNotificationSelector("Scope_example", *openapiclient.NewNotificationEventSelector("Level_example", "ResourceType_example", "Type_example")) // NotificationSelector | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AddJiraSelector(context.Background(), uuid).Selector(selector).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AddJiraSelector(context.Background(), uuid).Selector(selector).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddJiraSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddJiraSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddJiraSelector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddJiraSelectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **selector** | [**NotificationSelector**](NotificationSelector.md) |  | 

### Return type

[**NotificationSelector**](NotificationSelector.md)

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
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AddRoleUser(context.Background(), roleName).Member(member).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AddRoleUser(context.Background(), roleName).Member(member).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddRoleUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRoleUser`: RbacManagerRoleMember
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddRoleUser`: %v\n", resp)
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


## AddSlackConfiguration

> NotificationSlackEndpointConfiguration AddSlackConfiguration(ctx).Configuration(configuration).Execute()





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
    configuration := *openapiclient.NewNotificationSlackEndpointConfiguration("Url_example") // NotificationSlackEndpointConfiguration | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AddSlackConfiguration(context.Background()).Configuration(configuration).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AddSlackConfiguration(context.Background()).Configuration(configuration).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddSlackConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSlackConfiguration`: NotificationSlackEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddSlackConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddSlackConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configuration** | [**NotificationSlackEndpointConfiguration**](NotificationSlackEndpointConfiguration.md) |  | 

### Return type

[**NotificationSlackEndpointConfiguration**](NotificationSlackEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddSlackSelector

> NotificationSelector AddSlackSelector(ctx, uuid).Selector(selector).Execute()





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
    uuid := "uuid_example" // string | 
    selector := *openapiclient.NewNotificationSelector("Scope_example", *openapiclient.NewNotificationEventSelector("Level_example", "ResourceType_example", "Type_example")) // NotificationSelector | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AddSlackSelector(context.Background(), uuid).Selector(selector).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AddSlackSelector(context.Background(), uuid).Selector(selector).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddSlackSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSlackSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddSlackSelector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSlackSelectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **selector** | [**NotificationSelector**](NotificationSelector.md) |  | 

### Return type

[**NotificationSelector**](NotificationSelector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddSmtpConfiguration

> NotificationSMTPEndpointConfiguration AddSmtpConfiguration(ctx).Configuration(configuration).Execute()





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
    configuration := *openapiclient.NewNotificationSMTPEndpointConfiguration("Host_example", int32(123), "From_example", "To_example") // NotificationSMTPEndpointConfiguration | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AddSmtpConfiguration(context.Background()).Configuration(configuration).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AddSmtpConfiguration(context.Background()).Configuration(configuration).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddSmtpConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSmtpConfiguration`: NotificationSMTPEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddSmtpConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddSmtpConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configuration** | [**NotificationSMTPEndpointConfiguration**](NotificationSMTPEndpointConfiguration.md) |  | 

### Return type

[**NotificationSMTPEndpointConfiguration**](NotificationSMTPEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddSmtpSelector

> NotificationSelector AddSmtpSelector(ctx, uuid).Selector(selector).Execute()





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
    uuid := "uuid_example" // string | 
    selector := *openapiclient.NewNotificationSelector("Scope_example", *openapiclient.NewNotificationEventSelector("Level_example", "ResourceType_example", "Type_example")) // NotificationSelector | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AddSmtpSelector(context.Background(), uuid).Selector(selector).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AddSmtpSelector(context.Background(), uuid).Selector(selector).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddSmtpSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSmtpSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddSmtpSelector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSmtpSelectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **selector** | [**NotificationSelector**](NotificationSelector.md) |  | 

### Return type

[**NotificationSelector**](NotificationSelector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddTeamsConfiguration

> NotificationTeamsEndpointConfiguration AddTeamsConfiguration(ctx).Configuration(configuration).Execute()





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
    configuration := *openapiclient.NewNotificationTeamsEndpointConfiguration("Url_example") // NotificationTeamsEndpointConfiguration | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AddTeamsConfiguration(context.Background()).Configuration(configuration).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AddTeamsConfiguration(context.Background()).Configuration(configuration).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddTeamsConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTeamsConfiguration`: NotificationTeamsEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddTeamsConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddTeamsConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configuration** | [**NotificationTeamsEndpointConfiguration**](NotificationTeamsEndpointConfiguration.md) |  | 

### Return type

[**NotificationTeamsEndpointConfiguration**](NotificationTeamsEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddTeamsSelector

> NotificationSelector AddTeamsSelector(ctx, uuid).Selector(selector).Execute()





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
    uuid := "uuid_example" // string | 
    selector := *openapiclient.NewNotificationSelector("Scope_example", *openapiclient.NewNotificationEventSelector("Level_example", "ResourceType_example", "Type_example")) // NotificationSelector | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AddTeamsSelector(context.Background(), uuid).Selector(selector).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AddTeamsSelector(context.Background(), uuid).Selector(selector).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddTeamsSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTeamsSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddTeamsSelector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTeamsSelectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **selector** | [**NotificationSelector**](NotificationSelector.md) |  | 

### Return type

[**NotificationSelector**](NotificationSelector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddWebhookConfiguration

> NotificationWebhookEndpointConfiguration AddWebhookConfiguration(ctx).Configuration(configuration).Execute()





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
    configuration := *openapiclient.NewNotificationWebhookEndpointConfiguration("Url_example") // NotificationWebhookEndpointConfiguration | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AddWebhookConfiguration(context.Background()).Configuration(configuration).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AddWebhookConfiguration(context.Background()).Configuration(configuration).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddWebhookConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddWebhookConfiguration`: NotificationWebhookEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddWebhookConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddWebhookConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configuration** | [**NotificationWebhookEndpointConfiguration**](NotificationWebhookEndpointConfiguration.md) |  | 

### Return type

[**NotificationWebhookEndpointConfiguration**](NotificationWebhookEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddWebhookSelector

> NotificationSelector AddWebhookSelector(ctx, uuid).Selector(selector).Execute()





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
    uuid := "uuid_example" // string | 
    selector := *openapiclient.NewNotificationSelector("Scope_example", *openapiclient.NewNotificationEventSelector("Level_example", "ResourceType_example", "Type_example")) // NotificationSelector | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AddWebhookSelector(context.Background(), uuid).Selector(selector).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AddWebhookSelector(context.Background(), uuid).Selector(selector).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddWebhookSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddWebhookSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddWebhookSelector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddWebhookSelectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **selector** | [**NotificationSelector**](NotificationSelector.md) |  | 

### Return type

[**NotificationSelector**](NotificationSelector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGithubConfiguration

> DeleteGithubConfiguration(ctx, uuid).Execute()





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
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteGithubConfiguration(context.Background(), uuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteGithubConfiguration(context.Background(), uuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteGithubConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGithubConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGithubSelector

> DeleteGithubSelector(ctx, configurationUuid, selectorUuid).Execute()





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
    configurationUuid := "configurationUuid_example" // string | 
    selectorUuid := "selectorUuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteGithubSelector(context.Background(), configurationUuid, selectorUuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteGithubSelector(context.Background(), configurationUuid, selectorUuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteGithubSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationUuid** | **string** |  | 
**selectorUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGithubSelectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

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
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteIdp(context.Background(), name).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteIdp(context.Background(), name).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteIdp``: %v\n", err)
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


## DeleteJiraConfiguration

> DeleteJiraConfiguration(ctx, uuid).Execute()





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
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteJiraConfiguration(context.Background(), uuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteJiraConfiguration(context.Background(), uuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteJiraConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteJiraConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteJiraSelector

> DeleteJiraSelector(ctx, configurationUuid, selectorUuid).Execute()





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
    configurationUuid := "configurationUuid_example" // string | 
    selectorUuid := "selectorUuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteJiraSelector(context.Background(), configurationUuid, selectorUuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteJiraSelector(context.Background(), configurationUuid, selectorUuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteJiraSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationUuid** | **string** |  | 
**selectorUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteJiraSelectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

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
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteRoleUser(context.Background(), roleName).Username(username).ForAccount(forAccount).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteRoleUser(context.Background(), roleName).Username(username).ForAccount(forAccount).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteRoleUser``: %v\n", err)
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


## DeleteSlackConfiguration

> DeleteSlackConfiguration(ctx, uuid).Execute()





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
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteSlackConfiguration(context.Background(), uuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSlackConfiguration(context.Background(), uuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSlackConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSlackConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSlackSelector

> DeleteSlackSelector(ctx, configurationUuid, selectorUuid).Execute()





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
    configurationUuid := "configurationUuid_example" // string | 
    selectorUuid := "selectorUuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteSlackSelector(context.Background(), configurationUuid, selectorUuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSlackSelector(context.Background(), configurationUuid, selectorUuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSlackSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationUuid** | **string** |  | 
**selectorUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSlackSelectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSmtpConfiguration

> DeleteSmtpConfiguration(ctx, uuid).Execute()





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
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteSmtpConfiguration(context.Background(), uuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSmtpConfiguration(context.Background(), uuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSmtpConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSmtpConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSmtpSelector

> DeleteSmtpSelector(ctx, configurationUuid, selectorUuid).Execute()





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
    configurationUuid := "configurationUuid_example" // string | 
    selectorUuid := "selectorUuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteSmtpSelector(context.Background(), configurationUuid, selectorUuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSmtpSelector(context.Background(), configurationUuid, selectorUuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSmtpSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationUuid** | **string** |  | 
**selectorUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSmtpSelectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTeamsConfiguration

> DeleteTeamsConfiguration(ctx, uuid).Execute()





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
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteTeamsConfiguration(context.Background(), uuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteTeamsConfiguration(context.Background(), uuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteTeamsConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTeamsConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTeamsSelector

> DeleteTeamsSelector(ctx, configurationUuid, selectorUuid).Execute()





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
    configurationUuid := "configurationUuid_example" // string | 
    selectorUuid := "selectorUuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteTeamsSelector(context.Background(), configurationUuid, selectorUuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteTeamsSelector(context.Background(), configurationUuid, selectorUuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteTeamsSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationUuid** | **string** |  | 
**selectorUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTeamsSelectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebhookConfiguration

> DeleteWebhookConfiguration(ctx, uuid).Execute()





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
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteWebhookConfiguration(context.Background(), uuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteWebhookConfiguration(context.Background(), uuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteWebhookConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebhookSelector

> DeleteWebhookSelector(ctx, configurationUuid, selectorUuid).Execute()





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
    configurationUuid := "configurationUuid_example" // string | 
    selectorUuid := "selectorUuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteWebhookSelector(context.Background(), configurationUuid, selectorUuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteWebhookSelector(context.Background(), configurationUuid, selectorUuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteWebhookSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationUuid** | **string** |  | 
**selectorUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookSelectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGithubConfiguration

> NotificationGitHubEndpointConfigurationBase GetGithubConfiguration(ctx, uuid).Execute()





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
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetGithubConfiguration(context.Background(), uuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetGithubConfiguration(context.Background(), uuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetGithubConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGithubConfiguration`: NotificationGitHubEndpointConfigurationBase
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetGithubConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGithubConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationGitHubEndpointConfigurationBase**](NotificationGitHubEndpointConfigurationBase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGithubConfigurationStatus

> NotificationOperationalStatus GetGithubConfigurationStatus(ctx, uuid).Execute()





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
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetGithubConfigurationStatus(context.Background(), uuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetGithubConfigurationStatus(context.Background(), uuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetGithubConfigurationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGithubConfigurationStatus`: NotificationOperationalStatus
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetGithubConfigurationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGithubConfigurationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationOperationalStatus**](NotificationOperationalStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGithubSelector

> NotificationSelector GetGithubSelector(ctx, configurationUuid, selectorUuid).Execute()





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
    configurationUuid := "configurationUuid_example" // string | 
    selectorUuid := "selectorUuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetGithubSelector(context.Background(), configurationUuid, selectorUuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetGithubSelector(context.Background(), configurationUuid, selectorUuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetGithubSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGithubSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetGithubSelector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationUuid** | **string** |  | 
**selectorUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGithubSelectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NotificationSelector**](NotificationSelector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalQueryResult

> GetGlobalQueryResult(ctx, resultUuid).Page(page).Execute()





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
    resultUuid := "resultUuid_example" // string | 
    page := int32(56) // int32 | Page number to fetch. If omitted, '1' is default. Page numbers start at 1 (optional)

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetGlobalQueryResult(context.Background(), resultUuid).Page(page).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetGlobalQueryResult(context.Background(), resultUuid).Page(page).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetGlobalQueryResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resultUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalQueryResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number to fetch. If omitted, &#39;1&#39; is default. Page numbers start at 1 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdp

> RbacManagerSamlConfiguration GetIdp(ctx, name).Execute()





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
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetIdp(context.Background(), name).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetIdp(context.Background(), name).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetIdp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdp`: RbacManagerSamlConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetIdp`: %v\n", resp)
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

[**RbacManagerSamlConfiguration**](RbacManagerSamlConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJiraConfiguration

> NotificationJiraEndpointConfigurationBase GetJiraConfiguration(ctx, uuid).Execute()





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
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetJiraConfiguration(context.Background(), uuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetJiraConfiguration(context.Background(), uuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetJiraConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJiraConfiguration`: NotificationJiraEndpointConfigurationBase
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetJiraConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJiraConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationJiraEndpointConfigurationBase**](NotificationJiraEndpointConfigurationBase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJiraConfigurationStatus

> NotificationOperationalStatus GetJiraConfigurationStatus(ctx, uuid).Execute()





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
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetJiraConfigurationStatus(context.Background(), uuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetJiraConfigurationStatus(context.Background(), uuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetJiraConfigurationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJiraConfigurationStatus`: NotificationOperationalStatus
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetJiraConfigurationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJiraConfigurationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationOperationalStatus**](NotificationOperationalStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJiraSelector

> NotificationSelector GetJiraSelector(ctx, configurationUuid, selectorUuid).Execute()





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
    configurationUuid := "configurationUuid_example" // string | 
    selectorUuid := "selectorUuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetJiraSelector(context.Background(), configurationUuid, selectorUuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetJiraSelector(context.Background(), configurationUuid, selectorUuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetJiraSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJiraSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetJiraSelector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationUuid** | **string** |  | 
**selectorUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJiraSelectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NotificationSelector**](NotificationSelector.md)

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
    openapiclient "./openapi"
)

func main() {
    grantType := "grantType_example" // string | OAuth Grant type for token (optional) (default to "password")
    username := "username_example" // string | User to assign OAuth token to (optional)
    password := "password_example" // string | Password for corresponding user (optional)
    clientId := "clientId_example" // string | The type of client used for the OAuth token (optional) (default to "anonymous")
    refreshToken := "refreshToken_example" // string | The refresh token from a previous password grant request, used to get a new access_token (optional)

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetOauthToken(context.Background()).GrantType(grantType).Username(username).Password(password).ClientId(clientId).RefreshToken(refreshToken).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetOauthToken(context.Background()).GrantType(grantType).Username(username).Password(password).ClientId(clientId).RefreshToken(refreshToken).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetOauthToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOauthToken`: TokenResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetOauthToken`: %v\n", resp)
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


## GetQueryResult

> GetQueryResult(ctx, resultUuid).Page(page).Execute()





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
    resultUuid := "resultUuid_example" // string | 
    page := int32(56) // int32 | Page number to fetch. If omitted, '1' is default. Page numbers start at 1 (optional)

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetQueryResult(context.Background(), resultUuid).Page(page).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetQueryResult(context.Background(), resultUuid).Page(page).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetQueryResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resultUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueryResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number to fetch. If omitted, &#39;1&#39; is default. Page numbers start at 1 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

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
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetRole(context.Background(), roleName).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetRole(context.Background(), roleName).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRole`: RbacManagerRole
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetRole`: %v\n", resp)
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


## GetSlackConfiguration

> NotificationSlackEndpointConfiguration GetSlackConfiguration(ctx, uuid).Execute()





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
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSlackConfiguration(context.Background(), uuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetSlackConfiguration(context.Background(), uuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSlackConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSlackConfiguration`: NotificationSlackEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSlackConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSlackConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationSlackEndpointConfiguration**](NotificationSlackEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSlackConfigurationStatus

> NotificationOperationalStatus GetSlackConfigurationStatus(ctx, uuid).Execute()





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
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSlackConfigurationStatus(context.Background(), uuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetSlackConfigurationStatus(context.Background(), uuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSlackConfigurationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSlackConfigurationStatus`: NotificationOperationalStatus
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSlackConfigurationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSlackConfigurationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationOperationalStatus**](NotificationOperationalStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSlackSelector

> NotificationSelector GetSlackSelector(ctx, configurationUuid, selectorUuid).Execute()





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
    configurationUuid := "configurationUuid_example" // string | 
    selectorUuid := "selectorUuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSlackSelector(context.Background(), configurationUuid, selectorUuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetSlackSelector(context.Background(), configurationUuid, selectorUuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSlackSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSlackSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSlackSelector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationUuid** | **string** |  | 
**selectorUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSlackSelectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NotificationSelector**](NotificationSelector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSmtpConfiguration

> NotificationSMTPEndpointConfiguration GetSmtpConfiguration(ctx, uuid).Execute()





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
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSmtpConfiguration(context.Background(), uuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetSmtpConfiguration(context.Background(), uuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSmtpConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSmtpConfiguration`: NotificationSMTPEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSmtpConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSmtpConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationSMTPEndpointConfiguration**](NotificationSMTPEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSmtpConfigurationStatus

> NotificationOperationalStatus GetSmtpConfigurationStatus(ctx, uuid).Execute()





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
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSmtpConfigurationStatus(context.Background(), uuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetSmtpConfigurationStatus(context.Background(), uuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSmtpConfigurationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSmtpConfigurationStatus`: NotificationOperationalStatus
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSmtpConfigurationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSmtpConfigurationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationOperationalStatus**](NotificationOperationalStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSmtpSelector

> NotificationSelector GetSmtpSelector(ctx, configurationUuid, selectorUuid).Execute()





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
    configurationUuid := "configurationUuid_example" // string | 
    selectorUuid := "selectorUuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSmtpSelector(context.Background(), configurationUuid, selectorUuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetSmtpSelector(context.Background(), configurationUuid, selectorUuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSmtpSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSmtpSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSmtpSelector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationUuid** | **string** |  | 
**selectorUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSmtpSelectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NotificationSelector**](NotificationSelector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamsConfiguration

> NotificationTeamsEndpointConfiguration GetTeamsConfiguration(ctx, uuid).Execute()





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
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetTeamsConfiguration(context.Background(), uuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetTeamsConfiguration(context.Background(), uuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetTeamsConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTeamsConfiguration`: NotificationTeamsEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetTeamsConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamsConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationTeamsEndpointConfiguration**](NotificationTeamsEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamsConfigurationStatus

> NotificationOperationalStatus GetTeamsConfigurationStatus(ctx, uuid).Execute()





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
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetTeamsConfigurationStatus(context.Background(), uuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetTeamsConfigurationStatus(context.Background(), uuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetTeamsConfigurationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTeamsConfigurationStatus`: NotificationOperationalStatus
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetTeamsConfigurationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamsConfigurationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationOperationalStatus**](NotificationOperationalStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamsSelector

> NotificationSelector GetTeamsSelector(ctx, configurationUuid, selectorUuid).Execute()





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
    configurationUuid := "configurationUuid_example" // string | 
    selectorUuid := "selectorUuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetTeamsSelector(context.Background(), configurationUuid, selectorUuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetTeamsSelector(context.Background(), configurationUuid, selectorUuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetTeamsSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTeamsSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetTeamsSelector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationUuid** | **string** |  | 
**selectorUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamsSelectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NotificationSelector**](NotificationSelector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookConfiguration

> NotificationWebhookEndpointConfiguration GetWebhookConfiguration(ctx, uuid).Execute()





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
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetWebhookConfiguration(context.Background(), uuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetWebhookConfiguration(context.Background(), uuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetWebhookConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhookConfiguration`: NotificationWebhookEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetWebhookConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationWebhookEndpointConfiguration**](NotificationWebhookEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookConfigurationStatus

> NotificationOperationalStatus GetWebhookConfigurationStatus(ctx, uuid).Execute()





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
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetWebhookConfigurationStatus(context.Background(), uuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetWebhookConfigurationStatus(context.Background(), uuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetWebhookConfigurationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhookConfigurationStatus`: NotificationOperationalStatus
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetWebhookConfigurationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookConfigurationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationOperationalStatus**](NotificationOperationalStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookSelector

> NotificationSelector GetWebhookSelector(ctx, configurationUuid, selectorUuid).Execute()





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
    configurationUuid := "configurationUuid_example" // string | 
    selectorUuid := "selectorUuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetWebhookSelector(context.Background(), configurationUuid, selectorUuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetWebhookSelector(context.Background(), configurationUuid, selectorUuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetWebhookSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhookSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetWebhookSelector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationUuid** | **string** |  | 
**selectorUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookSelectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NotificationSelector**](NotificationSelector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HealthCheck

> HealthCheck(ctx).Execute()





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
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.HealthCheck(context.Background()).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.HealthCheck(context.Background()).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.HealthCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHealthCheckRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEndpoints

> []NotificationEndpoint ListEndpoints(ctx).Execute()





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
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListEndpoints(context.Background()).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListEndpoints(context.Background()).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListEndpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEndpoints`: []NotificationEndpoint
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListEndpoints`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListEndpointsRequest struct via the builder pattern


### Return type

[**[]NotificationEndpoint**](NotificationEndpoint.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFileContentSearchResults

> []FileContentSearchResult ListFileContentSearchResults(ctx, imageDigest).Execute()

Return a list of analyzer artifacts of the specified type

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
    imageDigest := "imageDigest_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListFileContentSearchResults(context.Background(), imageDigest).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListFileContentSearchResults(context.Background(), imageDigest).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListFileContentSearchResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFileContentSearchResults`: []FileContentSearchResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListFileContentSearchResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFileContentSearchResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]FileContentSearchResult**](FileContentSearchResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGithubConfigurations

> []NotificationGitHubEndpointConfigurationBase ListGithubConfigurations(ctx).Execute()





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
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListGithubConfigurations(context.Background()).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListGithubConfigurations(context.Background()).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListGithubConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGithubConfigurations`: []NotificationGitHubEndpointConfigurationBase
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListGithubConfigurations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListGithubConfigurationsRequest struct via the builder pattern


### Return type

[**[]NotificationGitHubEndpointConfigurationBase**](NotificationGitHubEndpointConfigurationBase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGithubSelectors

> []NotificationSelector ListGithubSelectors(ctx, uuid).Execute()





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
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListGithubSelectors(context.Background(), uuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListGithubSelectors(context.Background(), uuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListGithubSelectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGithubSelectors`: []NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListGithubSelectors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGithubSelectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]NotificationSelector**](NotificationSelector.md)

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
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListIdps(context.Background()).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListIdps(context.Background()).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListIdps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIdps`: []string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListIdps`: %v\n", resp)
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


## ListJiraConfigurations

> []NotificationJiraEndpointConfigurationBase ListJiraConfigurations(ctx).Execute()





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
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListJiraConfigurations(context.Background()).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListJiraConfigurations(context.Background()).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListJiraConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListJiraConfigurations`: []NotificationJiraEndpointConfigurationBase
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListJiraConfigurations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListJiraConfigurationsRequest struct via the builder pattern


### Return type

[**[]NotificationJiraEndpointConfigurationBase**](NotificationJiraEndpointConfigurationBase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListJiraSelectors

> []NotificationSelector ListJiraSelectors(ctx, uuid).Execute()





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
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListJiraSelectors(context.Background(), uuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListJiraSelectors(context.Background(), uuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListJiraSelectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListJiraSelectors`: []NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListJiraSelectors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListJiraSelectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]NotificationSelector**](NotificationSelector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRetrievedFiles

> []RetrievedFile ListRetrievedFiles(ctx, imageDigest).Execute()

Return a list of analyzer artifacts of the specified type

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
    imageDigest := "imageDigest_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListRetrievedFiles(context.Background(), imageDigest).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListRetrievedFiles(context.Background(), imageDigest).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListRetrievedFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRetrievedFiles`: []RetrievedFile
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListRetrievedFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRetrievedFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]RetrievedFile**](RetrievedFile.md)

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
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListRoleMembers(context.Background(), roleName).ForAccount(forAccount).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListRoleMembers(context.Background(), roleName).ForAccount(forAccount).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListRoleMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoleMembers`: []RbacManagerRoleMember
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListRoleMembers`: %v\n", resp)
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
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListRoles(context.Background()).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListRoles(context.Background()).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoles`: []RbacManagerRoleSummary
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListRoles`: %v\n", resp)
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


## ListSecretSearchResults

> []SecretSearchResult ListSecretSearchResults(ctx, imageDigest).Execute()

Return a list of analyzer artifacts of the specified type

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
    imageDigest := "imageDigest_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListSecretSearchResults(context.Background(), imageDigest).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSecretSearchResults(context.Background(), imageDigest).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSecretSearchResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSecretSearchResults`: []SecretSearchResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSecretSearchResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSecretSearchResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]SecretSearchResult**](SecretSearchResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSelectors

> []NotificationSelector ListSelectors(ctx).Execute()





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
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListSelectors(context.Background()).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSelectors(context.Background()).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSelectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSelectors`: []NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSelectors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSelectorsRequest struct via the builder pattern


### Return type

[**[]NotificationSelector**](NotificationSelector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSlackConfigurations

> []NotificationSlackEndpointConfiguration ListSlackConfigurations(ctx).Execute()





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
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListSlackConfigurations(context.Background()).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSlackConfigurations(context.Background()).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSlackConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSlackConfigurations`: []NotificationSlackEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSlackConfigurations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSlackConfigurationsRequest struct via the builder pattern


### Return type

[**[]NotificationSlackEndpointConfiguration**](NotificationSlackEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSlackSelectors

> []NotificationSelector ListSlackSelectors(ctx, uuid).Execute()





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
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListSlackSelectors(context.Background(), uuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSlackSelectors(context.Background(), uuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSlackSelectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSlackSelectors`: []NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSlackSelectors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSlackSelectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]NotificationSelector**](NotificationSelector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSmtpConfigurations

> []NotificationSMTPEndpointConfiguration ListSmtpConfigurations(ctx).Execute()





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
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListSmtpConfigurations(context.Background()).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSmtpConfigurations(context.Background()).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSmtpConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSmtpConfigurations`: []NotificationSMTPEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSmtpConfigurations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSmtpConfigurationsRequest struct via the builder pattern


### Return type

[**[]NotificationSMTPEndpointConfiguration**](NotificationSMTPEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSmtpSelectors

> []NotificationSelector ListSmtpSelectors(ctx, uuid).Execute()





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
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListSmtpSelectors(context.Background(), uuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSmtpSelectors(context.Background(), uuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSmtpSelectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSmtpSelectors`: []NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSmtpSelectors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSmtpSelectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]NotificationSelector**](NotificationSelector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTeamsConfigurations

> []NotificationTeamsEndpointConfiguration ListTeamsConfigurations(ctx).Execute()





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
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListTeamsConfigurations(context.Background()).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListTeamsConfigurations(context.Background()).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListTeamsConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTeamsConfigurations`: []NotificationTeamsEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListTeamsConfigurations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTeamsConfigurationsRequest struct via the builder pattern


### Return type

[**[]NotificationTeamsEndpointConfiguration**](NotificationTeamsEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTeamsSelectors

> []NotificationSelector ListTeamsSelectors(ctx, uuid).Execute()





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
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListTeamsSelectors(context.Background(), uuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListTeamsSelectors(context.Background(), uuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListTeamsSelectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTeamsSelectors`: []NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListTeamsSelectors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTeamsSelectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]NotificationSelector**](NotificationSelector.md)

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
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListUserRoles(context.Background(), username).ForAccount(forAccount).Role(role).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListUserRoles(context.Background(), username).ForAccount(forAccount).Role(role).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListUserRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserRoles`: []RbacManagerRoleMembership
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListUserRoles`: %v\n", resp)
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


## ListWebhookConfigurations

> []NotificationWebhookEndpointConfiguration ListWebhookConfigurations(ctx).Execute()





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
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListWebhookConfigurations(context.Background()).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListWebhookConfigurations(context.Background()).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListWebhookConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWebhookConfigurations`: []NotificationWebhookEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListWebhookConfigurations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWebhookConfigurationsRequest struct via the builder pattern


### Return type

[**[]NotificationWebhookEndpointConfiguration**](NotificationWebhookEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebhookSelectors

> []NotificationSelector ListWebhookSelectors(ctx, uuid).Execute()





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
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListWebhookSelectors(context.Background(), uuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListWebhookSelectors(context.Background(), uuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListWebhookSelectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWebhookSelectors`: []NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListWebhookSelectors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWebhookSelectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]NotificationSelector**](NotificationSelector.md)

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
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.MyRoles(context.Background()).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.MyRoles(context.Background()).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.MyRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MyRoles`: []RbacManagerAccountRole
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.MyRoles`: %v\n", resp)
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


## Ping

> string Ping(ctx).Execute()





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
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Ping(context.Background()).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.Ping(context.Background()).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Ping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Ping`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Ping`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPingRequest struct via the builder pattern


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
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
    openapiclient "./openapi"
)

func main() {
    token := "token_example" // string | The token to be revoked (optional)
    tokenTypeHint := "tokenTypeHint_example" // string | A hint about the type of token to be revoked (optional)

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RevokeOauthToken(context.Background()).Token(token).TokenTypeHint(tokenTypeHint).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.RevokeOauthToken(context.Background()).Token(token).TokenTypeHint(tokenTypeHint).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RevokeOauthToken``: %v\n", err)
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
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SamlLogin(context.Background(), idpName).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.SamlLogin(context.Background(), idpName).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SamlLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SamlLogin`: RbacManagerTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SamlLogin`: %v\n", resp)
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
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SamlSso(context.Background(), idpName).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.SamlSso(context.Background(), idpName).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SamlSso``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SamlSso`: RbacManagerTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SamlSso`: %v\n", resp)
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


## TestGithubConfiguration

> NotificationGitHubTestResult TestGithubConfiguration(ctx).Configuration(configuration).Execute()





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
    configuration := *openapiclient.NewNotificationGitHubEndpointConfigurationPost("Username_example", "AccessToken_example", "Owner_example", "Repository_example") // NotificationGitHubEndpointConfigurationPost | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.TestGithubConfiguration(context.Background()).Configuration(configuration).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.TestGithubConfiguration(context.Background()).Configuration(configuration).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TestGithubConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestGithubConfiguration`: NotificationGitHubTestResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TestGithubConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestGithubConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configuration** | [**NotificationGitHubEndpointConfigurationPost**](NotificationGitHubEndpointConfigurationPost.md) |  | 

### Return type

[**NotificationGitHubTestResult**](NotificationGitHubTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestJiraConfiguration

> NotificationJiraTestResult TestJiraConfiguration(ctx).Configuration(configuration).Execute()





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
    configuration := *openapiclient.NewNotificationJiraEndpointConfigurationPost("Url_example", "Username_example", "Password_example", "ProjectKey_example", "IssueType_example") // NotificationJiraEndpointConfigurationPost | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.TestJiraConfiguration(context.Background()).Configuration(configuration).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.TestJiraConfiguration(context.Background()).Configuration(configuration).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TestJiraConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestJiraConfiguration`: NotificationJiraTestResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TestJiraConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestJiraConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configuration** | [**NotificationJiraEndpointConfigurationPost**](NotificationJiraEndpointConfigurationPost.md) |  | 

### Return type

[**NotificationJiraTestResult**](NotificationJiraTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestSlackConfiguration

> NotificationSlackTestResult TestSlackConfiguration(ctx).Configuration(configuration).Execute()





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
    configuration := *openapiclient.NewNotificationSlackEndpointConfiguration("Url_example") // NotificationSlackEndpointConfiguration | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.TestSlackConfiguration(context.Background()).Configuration(configuration).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.TestSlackConfiguration(context.Background()).Configuration(configuration).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TestSlackConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestSlackConfiguration`: NotificationSlackTestResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TestSlackConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestSlackConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configuration** | [**NotificationSlackEndpointConfiguration**](NotificationSlackEndpointConfiguration.md) |  | 

### Return type

[**NotificationSlackTestResult**](NotificationSlackTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestSmtpConfiguration

> NotificationSMTPTestResult TestSmtpConfiguration(ctx).Configuration(configuration).Execute()





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
    configuration := *openapiclient.NewNotificationSMTPEndpointConfiguration("Host_example", int32(123), "From_example", "To_example") // NotificationSMTPEndpointConfiguration | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.TestSmtpConfiguration(context.Background()).Configuration(configuration).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.TestSmtpConfiguration(context.Background()).Configuration(configuration).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TestSmtpConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestSmtpConfiguration`: NotificationSMTPTestResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TestSmtpConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestSmtpConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configuration** | [**NotificationSMTPEndpointConfiguration**](NotificationSMTPEndpointConfiguration.md) |  | 

### Return type

[**NotificationSMTPTestResult**](NotificationSMTPTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestStoredGithubConfiguration

> NotificationGitHubTestResult TestStoredGithubConfiguration(ctx, uuid).Execute()





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
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.TestStoredGithubConfiguration(context.Background(), uuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.TestStoredGithubConfiguration(context.Background(), uuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TestStoredGithubConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestStoredGithubConfiguration`: NotificationGitHubTestResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TestStoredGithubConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestStoredGithubConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationGitHubTestResult**](NotificationGitHubTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestStoredJiraConfiguration

> NotificationJiraTestResult TestStoredJiraConfiguration(ctx, uuid).Execute()





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
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.TestStoredJiraConfiguration(context.Background(), uuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.TestStoredJiraConfiguration(context.Background(), uuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TestStoredJiraConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestStoredJiraConfiguration`: NotificationJiraTestResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TestStoredJiraConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestStoredJiraConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationJiraTestResult**](NotificationJiraTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestStoredSlackConfiguration

> NotificationSlackTestResult TestStoredSlackConfiguration(ctx, uuid).Execute()





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
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.TestStoredSlackConfiguration(context.Background(), uuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.TestStoredSlackConfiguration(context.Background(), uuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TestStoredSlackConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestStoredSlackConfiguration`: NotificationSlackTestResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TestStoredSlackConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestStoredSlackConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationSlackTestResult**](NotificationSlackTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestStoredSmtpConfiguration

> NotificationSMTPTestResult TestStoredSmtpConfiguration(ctx, uuid).Execute()





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
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.TestStoredSmtpConfiguration(context.Background(), uuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.TestStoredSmtpConfiguration(context.Background(), uuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TestStoredSmtpConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestStoredSmtpConfiguration`: NotificationSMTPTestResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TestStoredSmtpConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestStoredSmtpConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationSMTPTestResult**](NotificationSMTPTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestStoredTeamsConfiguration

> NotificationTeamsTestResult TestStoredTeamsConfiguration(ctx, uuid).Execute()





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
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.TestStoredTeamsConfiguration(context.Background(), uuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.TestStoredTeamsConfiguration(context.Background(), uuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TestStoredTeamsConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestStoredTeamsConfiguration`: NotificationTeamsTestResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TestStoredTeamsConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestStoredTeamsConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationTeamsTestResult**](NotificationTeamsTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestStoredWebhookConfiguration

> NotificationWebhookTestResult TestStoredWebhookConfiguration(ctx, uuid).Execute()





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
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.TestStoredWebhookConfiguration(context.Background(), uuid).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.TestStoredWebhookConfiguration(context.Background(), uuid).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TestStoredWebhookConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestStoredWebhookConfiguration`: NotificationWebhookTestResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TestStoredWebhookConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestStoredWebhookConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationWebhookTestResult**](NotificationWebhookTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestTeamsConfiguration

> NotificationTeamsTestResult TestTeamsConfiguration(ctx).Configuration(configuration).Execute()





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
    configuration := *openapiclient.NewNotificationTeamsEndpointConfiguration("Url_example") // NotificationTeamsEndpointConfiguration | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.TestTeamsConfiguration(context.Background()).Configuration(configuration).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.TestTeamsConfiguration(context.Background()).Configuration(configuration).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TestTeamsConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestTeamsConfiguration`: NotificationTeamsTestResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TestTeamsConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestTeamsConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configuration** | [**NotificationTeamsEndpointConfiguration**](NotificationTeamsEndpointConfiguration.md) |  | 

### Return type

[**NotificationTeamsTestResult**](NotificationTeamsTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestWebhookConfiguration

> NotificationWebhookTestResult TestWebhookConfiguration(ctx).Configuration(configuration).Execute()





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
    configuration := *openapiclient.NewNotificationWebhookEndpointConfiguration("Url_example") // NotificationWebhookEndpointConfiguration | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.TestWebhookConfiguration(context.Background()).Configuration(configuration).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.TestWebhookConfiguration(context.Background()).Configuration(configuration).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TestWebhookConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestWebhookConfiguration`: NotificationWebhookTestResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TestWebhookConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestWebhookConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configuration** | [**NotificationWebhookEndpointConfiguration**](NotificationWebhookEndpointConfiguration.md) |  | 

### Return type

[**NotificationWebhookTestResult**](NotificationWebhookTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEndpointStatus

> NotificationEndpointEnabledStatus UpdateEndpointStatus(ctx, name).Status(status).Execute()





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
    status := *openapiclient.NewNotificationEndpointEnabledStatus() // NotificationEndpointEnabledStatus | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateEndpointStatus(context.Background(), name).Status(status).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateEndpointStatus(context.Background(), name).Status(status).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateEndpointStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEndpointStatus`: NotificationEndpointEnabledStatus
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateEndpointStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEndpointStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | [**NotificationEndpointEnabledStatus**](NotificationEndpointEnabledStatus.md) |  | 

### Return type

[**NotificationEndpointEnabledStatus**](NotificationEndpointEnabledStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGithubConfiguration

> NotificationGitHubEndpointConfigurationBase UpdateGithubConfiguration(ctx, uuid).Configuration(configuration).Execute()





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
    uuid := "uuid_example" // string | 
    configuration := *openapiclient.NewNotificationGitHubEndpointConfigurationPut("Username_example", "Owner_example", "Repository_example") // NotificationGitHubEndpointConfigurationPut | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateGithubConfiguration(context.Background(), uuid).Configuration(configuration).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateGithubConfiguration(context.Background(), uuid).Configuration(configuration).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateGithubConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGithubConfiguration`: NotificationGitHubEndpointConfigurationBase
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateGithubConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGithubConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **configuration** | [**NotificationGitHubEndpointConfigurationPut**](NotificationGitHubEndpointConfigurationPut.md) |  | 

### Return type

[**NotificationGitHubEndpointConfigurationBase**](NotificationGitHubEndpointConfigurationBase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGithubSelector

> NotificationSelector UpdateGithubSelector(ctx, configurationUuid, selectorUuid).Selector(selector).Execute()





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
    configurationUuid := "configurationUuid_example" // string | 
    selectorUuid := "selectorUuid_example" // string | 
    selector := *openapiclient.NewNotificationSelector("Scope_example", *openapiclient.NewNotificationEventSelector("Level_example", "ResourceType_example", "Type_example")) // NotificationSelector | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateGithubSelector(context.Background(), configurationUuid, selectorUuid).Selector(selector).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateGithubSelector(context.Background(), configurationUuid, selectorUuid).Selector(selector).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateGithubSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGithubSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateGithubSelector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationUuid** | **string** |  | 
**selectorUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGithubSelectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **selector** | [**NotificationSelector**](NotificationSelector.md) |  | 

### Return type

[**NotificationSelector**](NotificationSelector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdp

> RbacManagerSamlConfiguration UpdateIdp(ctx, name).Configuration(configuration).Execute()





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
    configuration := *openapiclient.NewRbacManagerSamlConfiguration("Name_example", false, "SpEntityId_example", "AcsUrl_example") // RbacManagerSamlConfiguration | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateIdp(context.Background(), name).Configuration(configuration).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateIdp(context.Background(), name).Configuration(configuration).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateIdp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIdp`: RbacManagerSamlConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateIdp`: %v\n", resp)
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


## UpdateJiraConfiguration

> NotificationJiraEndpointConfigurationBase UpdateJiraConfiguration(ctx, uuid).Configuration(configuration).Execute()





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
    uuid := "uuid_example" // string | 
    configuration := *openapiclient.NewNotificationJiraEndpointConfigurationPut("Url_example", "Username_example", "ProjectKey_example", "IssueType_example") // NotificationJiraEndpointConfigurationPut | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateJiraConfiguration(context.Background(), uuid).Configuration(configuration).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateJiraConfiguration(context.Background(), uuid).Configuration(configuration).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateJiraConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateJiraConfiguration`: NotificationJiraEndpointConfigurationBase
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateJiraConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateJiraConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **configuration** | [**NotificationJiraEndpointConfigurationPut**](NotificationJiraEndpointConfigurationPut.md) |  | 

### Return type

[**NotificationJiraEndpointConfigurationBase**](NotificationJiraEndpointConfigurationBase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateJiraSelector

> NotificationSelector UpdateJiraSelector(ctx, configurationUuid, selectorUuid).Selector(selector).Execute()





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
    configurationUuid := "configurationUuid_example" // string | 
    selectorUuid := "selectorUuid_example" // string | 
    selector := *openapiclient.NewNotificationSelector("Scope_example", *openapiclient.NewNotificationEventSelector("Level_example", "ResourceType_example", "Type_example")) // NotificationSelector | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateJiraSelector(context.Background(), configurationUuid, selectorUuid).Selector(selector).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateJiraSelector(context.Background(), configurationUuid, selectorUuid).Selector(selector).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateJiraSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateJiraSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateJiraSelector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationUuid** | **string** |  | 
**selectorUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateJiraSelectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **selector** | [**NotificationSelector**](NotificationSelector.md) |  | 

### Return type

[**NotificationSelector**](NotificationSelector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSlackConfiguration

> NotificationSlackEndpointConfiguration UpdateSlackConfiguration(ctx, uuid).Configuration(configuration).Execute()





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
    uuid := "uuid_example" // string | 
    configuration := *openapiclient.NewNotificationSlackEndpointConfiguration("Url_example") // NotificationSlackEndpointConfiguration | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateSlackConfiguration(context.Background(), uuid).Configuration(configuration).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateSlackConfiguration(context.Background(), uuid).Configuration(configuration).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSlackConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSlackConfiguration`: NotificationSlackEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSlackConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSlackConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **configuration** | [**NotificationSlackEndpointConfiguration**](NotificationSlackEndpointConfiguration.md) |  | 

### Return type

[**NotificationSlackEndpointConfiguration**](NotificationSlackEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSlackSelector

> NotificationSelector UpdateSlackSelector(ctx, configurationUuid, selectorUuid).Selector(selector).Execute()





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
    configurationUuid := "configurationUuid_example" // string | 
    selectorUuid := "selectorUuid_example" // string | 
    selector := *openapiclient.NewNotificationSelector("Scope_example", *openapiclient.NewNotificationEventSelector("Level_example", "ResourceType_example", "Type_example")) // NotificationSelector | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateSlackSelector(context.Background(), configurationUuid, selectorUuid).Selector(selector).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateSlackSelector(context.Background(), configurationUuid, selectorUuid).Selector(selector).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSlackSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSlackSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSlackSelector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationUuid** | **string** |  | 
**selectorUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSlackSelectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **selector** | [**NotificationSelector**](NotificationSelector.md) |  | 

### Return type

[**NotificationSelector**](NotificationSelector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSmtpConfiguration

> NotificationSMTPEndpointConfiguration UpdateSmtpConfiguration(ctx, uuid).Configuration(configuration).Execute()





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
    uuid := "uuid_example" // string | 
    configuration := *openapiclient.NewNotificationSMTPEndpointConfiguration("Host_example", int32(123), "From_example", "To_example") // NotificationSMTPEndpointConfiguration | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateSmtpConfiguration(context.Background(), uuid).Configuration(configuration).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateSmtpConfiguration(context.Background(), uuid).Configuration(configuration).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSmtpConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSmtpConfiguration`: NotificationSMTPEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSmtpConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSmtpConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **configuration** | [**NotificationSMTPEndpointConfiguration**](NotificationSMTPEndpointConfiguration.md) |  | 

### Return type

[**NotificationSMTPEndpointConfiguration**](NotificationSMTPEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSmtpSelector

> NotificationSelector UpdateSmtpSelector(ctx, configurationUuid, selectorUuid).Selector(selector).Execute()





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
    configurationUuid := "configurationUuid_example" // string | 
    selectorUuid := "selectorUuid_example" // string | 
    selector := *openapiclient.NewNotificationSelector("Scope_example", *openapiclient.NewNotificationEventSelector("Level_example", "ResourceType_example", "Type_example")) // NotificationSelector | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateSmtpSelector(context.Background(), configurationUuid, selectorUuid).Selector(selector).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateSmtpSelector(context.Background(), configurationUuid, selectorUuid).Selector(selector).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSmtpSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSmtpSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSmtpSelector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationUuid** | **string** |  | 
**selectorUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSmtpSelectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **selector** | [**NotificationSelector**](NotificationSelector.md) |  | 

### Return type

[**NotificationSelector**](NotificationSelector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTeamsConfiguration

> NotificationTeamsEndpointConfiguration UpdateTeamsConfiguration(ctx, uuid).Configuration(configuration).Execute()





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
    uuid := "uuid_example" // string | 
    configuration := *openapiclient.NewNotificationTeamsEndpointConfiguration("Url_example") // NotificationTeamsEndpointConfiguration | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateTeamsConfiguration(context.Background(), uuid).Configuration(configuration).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateTeamsConfiguration(context.Background(), uuid).Configuration(configuration).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateTeamsConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTeamsConfiguration`: NotificationTeamsEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateTeamsConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTeamsConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **configuration** | [**NotificationTeamsEndpointConfiguration**](NotificationTeamsEndpointConfiguration.md) |  | 

### Return type

[**NotificationTeamsEndpointConfiguration**](NotificationTeamsEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTeamsSelector

> NotificationSelector UpdateTeamsSelector(ctx, configurationUuid, selectorUuid).Selector(selector).Execute()





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
    configurationUuid := "configurationUuid_example" // string | 
    selectorUuid := "selectorUuid_example" // string | 
    selector := *openapiclient.NewNotificationSelector("Scope_example", *openapiclient.NewNotificationEventSelector("Level_example", "ResourceType_example", "Type_example")) // NotificationSelector | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateTeamsSelector(context.Background(), configurationUuid, selectorUuid).Selector(selector).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateTeamsSelector(context.Background(), configurationUuid, selectorUuid).Selector(selector).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateTeamsSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTeamsSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateTeamsSelector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationUuid** | **string** |  | 
**selectorUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTeamsSelectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **selector** | [**NotificationSelector**](NotificationSelector.md) |  | 

### Return type

[**NotificationSelector**](NotificationSelector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhookConfiguration

> NotificationWebhookEndpointConfiguration UpdateWebhookConfiguration(ctx, uuid).Configuration(configuration).Execute()





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
    uuid := "uuid_example" // string | 
    configuration := *openapiclient.NewNotificationWebhookEndpointConfiguration("Url_example") // NotificationWebhookEndpointConfiguration | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateWebhookConfiguration(context.Background(), uuid).Configuration(configuration).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateWebhookConfiguration(context.Background(), uuid).Configuration(configuration).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateWebhookConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWebhookConfiguration`: NotificationWebhookEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateWebhookConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **configuration** | [**NotificationWebhookEndpointConfiguration**](NotificationWebhookEndpointConfiguration.md) |  | 

### Return type

[**NotificationWebhookEndpointConfiguration**](NotificationWebhookEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhookSelector

> NotificationSelector UpdateWebhookSelector(ctx, configurationUuid, selectorUuid).Selector(selector).Execute()





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
    configurationUuid := "configurationUuid_example" // string | 
    selectorUuid := "selectorUuid_example" // string | 
    selector := *openapiclient.NewNotificationSelector("Scope_example", *openapiclient.NewNotificationEventSelector("Level_example", "ResourceType_example", "Type_example")) // NotificationSelector | 

    configuration := openapiclient.NewConfiguration()
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateWebhookSelector(context.Background(), configurationUuid, selectorUuid).Selector(selector).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateWebhookSelector(context.Background(), configurationUuid, selectorUuid).Selector(selector).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateWebhookSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWebhookSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateWebhookSelector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationUuid** | **string** |  | 
**selectorUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookSelectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **selector** | [**NotificationSelector**](NotificationSelector.md) |  | 

### Return type

[**NotificationSelector**](NotificationSelector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VersionCheck

> ServiceVersion VersionCheck(ctx).Execute()





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
<<<<<<< HEAD
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.VersionCheck(context.Background()).Execute()
=======
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.VersionCheck(context.Background()).Execute()
>>>>>>> main
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.VersionCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VersionCheck`: ServiceVersion
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.VersionCheck`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVersionCheckRequest struct via the builder pattern


### Return type

[**ServiceVersion**](ServiceVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

