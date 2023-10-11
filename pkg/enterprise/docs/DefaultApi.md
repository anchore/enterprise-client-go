# \DefaultAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGithubConfiguration**](DefaultAPI.md#AddGithubConfiguration) | **Post** /notifications/endpoints/github/configurations | 
[**AddGithubSelector**](DefaultAPI.md#AddGithubSelector) | **Post** /notifications/endpoints/github/configurations/{uuid}/selectors | 
[**AddIdp**](DefaultAPI.md#AddIdp) | **Post** /rbac-manager/saml/idps | 
[**AddJiraConfiguration**](DefaultAPI.md#AddJiraConfiguration) | **Post** /notifications/endpoints/jira/configurations | 
[**AddJiraSelector**](DefaultAPI.md#AddJiraSelector) | **Post** /notifications/endpoints/jira/configurations/{uuid}/selectors | 
[**AddRoleUser**](DefaultAPI.md#AddRoleUser) | **Post** /rbac-manager/roles/{role_name}/members | Add a user to the role
[**AddSlackConfiguration**](DefaultAPI.md#AddSlackConfiguration) | **Post** /notifications/endpoints/slack/configurations | 
[**AddSlackSelector**](DefaultAPI.md#AddSlackSelector) | **Post** /notifications/endpoints/slack/configurations/{uuid}/selectors | 
[**AddSmtpConfiguration**](DefaultAPI.md#AddSmtpConfiguration) | **Post** /notifications/endpoints/smtp/configurations | 
[**AddSmtpSelector**](DefaultAPI.md#AddSmtpSelector) | **Post** /notifications/endpoints/smtp/configurations/{uuid}/selectors | 
[**AddTeamsConfiguration**](DefaultAPI.md#AddTeamsConfiguration) | **Post** /notifications/endpoints/teams/configurations | 
[**AddTeamsSelector**](DefaultAPI.md#AddTeamsSelector) | **Post** /notifications/endpoints/teams/configurations/{uuid}/selectors | 
[**AddWebhookConfiguration**](DefaultAPI.md#AddWebhookConfiguration) | **Post** /notifications/endpoints/webhook/configurations | 
[**AddWebhookSelector**](DefaultAPI.md#AddWebhookSelector) | **Post** /notifications/endpoints/webhook/configurations/{uuid}/selectors | 
[**DeleteGithubConfiguration**](DefaultAPI.md#DeleteGithubConfiguration) | **Delete** /notifications/endpoints/github/configurations/{uuid} | 
[**DeleteGithubSelector**](DefaultAPI.md#DeleteGithubSelector) | **Delete** /notifications/endpoints/github/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**DeleteIdp**](DefaultAPI.md#DeleteIdp) | **Delete** /rbac-manager/saml/idps/{name} | 
[**DeleteJiraConfiguration**](DefaultAPI.md#DeleteJiraConfiguration) | **Delete** /notifications/endpoints/jira/configurations/{uuid} | 
[**DeleteJiraSelector**](DefaultAPI.md#DeleteJiraSelector) | **Delete** /notifications/endpoints/jira/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**DeleteRoleUser**](DefaultAPI.md#DeleteRoleUser) | **Delete** /rbac-manager/roles/{role_name}/members | Remove a user from the role
[**DeleteSlackConfiguration**](DefaultAPI.md#DeleteSlackConfiguration) | **Delete** /notifications/endpoints/slack/configurations/{uuid} | 
[**DeleteSlackSelector**](DefaultAPI.md#DeleteSlackSelector) | **Delete** /notifications/endpoints/slack/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**DeleteSmtpConfiguration**](DefaultAPI.md#DeleteSmtpConfiguration) | **Delete** /notifications/endpoints/smtp/configurations/{uuid} | 
[**DeleteSmtpSelector**](DefaultAPI.md#DeleteSmtpSelector) | **Delete** /notifications/endpoints/smtp/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**DeleteTeamsConfiguration**](DefaultAPI.md#DeleteTeamsConfiguration) | **Delete** /notifications/endpoints/teams/configurations/{uuid} | 
[**DeleteTeamsSelector**](DefaultAPI.md#DeleteTeamsSelector) | **Delete** /notifications/endpoints/teams/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**DeleteWebhookConfiguration**](DefaultAPI.md#DeleteWebhookConfiguration) | **Delete** /notifications/endpoints/webhook/configurations/{uuid} | 
[**DeleteWebhookSelector**](DefaultAPI.md#DeleteWebhookSelector) | **Delete** /notifications/endpoints/webhook/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**GetGithubConfiguration**](DefaultAPI.md#GetGithubConfiguration) | **Get** /notifications/endpoints/github/configurations/{uuid} | 
[**GetGithubConfigurationStatus**](DefaultAPI.md#GetGithubConfigurationStatus) | **Get** /notifications/endpoints/github/configurations/{uuid}/status | 
[**GetGithubSelector**](DefaultAPI.md#GetGithubSelector) | **Get** /notifications/endpoints/github/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**GetGlobalQueryResult**](DefaultAPI.md#GetGlobalQueryResult) | **Get** /reporting/reports/global/scheduled-query-results/{result_uuid} | 
[**GetIdp**](DefaultAPI.md#GetIdp) | **Get** /rbac-manager/saml/idps/{name} | 
[**GetJiraConfiguration**](DefaultAPI.md#GetJiraConfiguration) | **Get** /notifications/endpoints/jira/configurations/{uuid} | 
[**GetJiraConfigurationStatus**](DefaultAPI.md#GetJiraConfigurationStatus) | **Get** /notifications/endpoints/jira/configurations/{uuid}/status | 
[**GetJiraSelector**](DefaultAPI.md#GetJiraSelector) | **Get** /notifications/endpoints/jira/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**GetOauthToken**](DefaultAPI.md#GetOauthToken) | **Post** /oauth/token | 
[**GetQueryResult**](DefaultAPI.md#GetQueryResult) | **Get** /reporting/scheduled-query-results/{result_uuid} | 
[**GetRole**](DefaultAPI.md#GetRole) | **Get** /rbac-manager/roles/{role_name} | Get detailed information about a specific role
[**GetSlackConfiguration**](DefaultAPI.md#GetSlackConfiguration) | **Get** /notifications/endpoints/slack/configurations/{uuid} | 
[**GetSlackConfigurationStatus**](DefaultAPI.md#GetSlackConfigurationStatus) | **Get** /notifications/endpoints/slack/configurations/{uuid}/status | 
[**GetSlackSelector**](DefaultAPI.md#GetSlackSelector) | **Get** /notifications/endpoints/slack/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**GetSmtpConfiguration**](DefaultAPI.md#GetSmtpConfiguration) | **Get** /notifications/endpoints/smtp/configurations/{uuid} | 
[**GetSmtpConfigurationStatus**](DefaultAPI.md#GetSmtpConfigurationStatus) | **Get** /notifications/endpoints/smtp/configurations/{uuid}/status | 
[**GetSmtpSelector**](DefaultAPI.md#GetSmtpSelector) | **Get** /notifications/endpoints/smtp/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**GetTeamsConfiguration**](DefaultAPI.md#GetTeamsConfiguration) | **Get** /notifications/endpoints/teams/configurations/{uuid} | 
[**GetTeamsConfigurationStatus**](DefaultAPI.md#GetTeamsConfigurationStatus) | **Get** /notifications/endpoints/teams/configurations/{uuid}/status | 
[**GetTeamsSelector**](DefaultAPI.md#GetTeamsSelector) | **Get** /notifications/endpoints/teams/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**GetWebhookConfiguration**](DefaultAPI.md#GetWebhookConfiguration) | **Get** /notifications/endpoints/webhook/configurations/{uuid} | 
[**GetWebhookConfigurationStatus**](DefaultAPI.md#GetWebhookConfigurationStatus) | **Get** /notifications/endpoints/webhook/configurations/{uuid}/status | 
[**GetWebhookSelector**](DefaultAPI.md#GetWebhookSelector) | **Get** /notifications/endpoints/webhook/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**HealthCheck**](DefaultAPI.md#HealthCheck) | **Get** /health | 
[**ListEndpoints**](DefaultAPI.md#ListEndpoints) | **Get** /notifications/endpoints | 
[**ListFileContentSearchResults**](DefaultAPI.md#ListFileContentSearchResults) | **Get** /images/{image_digest}/artifacts/file-content-search | Return a list of analyzer artifacts of the specified type
[**ListGithubConfigurations**](DefaultAPI.md#ListGithubConfigurations) | **Get** /notifications/endpoints/github/configurations | 
[**ListGithubSelectors**](DefaultAPI.md#ListGithubSelectors) | **Get** /notifications/endpoints/github/configurations/{uuid}/selectors | 
[**ListIdps**](DefaultAPI.md#ListIdps) | **Get** /rbac-manager/saml/idps | 
[**ListJiraConfigurations**](DefaultAPI.md#ListJiraConfigurations) | **Get** /notifications/endpoints/jira/configurations | 
[**ListJiraSelectors**](DefaultAPI.md#ListJiraSelectors) | **Get** /notifications/endpoints/jira/configurations/{uuid}/selectors | 
[**ListRetrievedFiles**](DefaultAPI.md#ListRetrievedFiles) | **Get** /images/{image_digest}/artifacts/retrieved-files | Return a list of analyzer artifacts of the specified type
[**ListRoleMembers**](DefaultAPI.md#ListRoleMembers) | **Get** /rbac-manager/roles/{role_name}/members | Returns a list of objects that have members in the role. The list is filtered by &#39;listRoleMembers&#39; access for the &#39;account&#39; element of each entry.
[**ListRoles**](DefaultAPI.md#ListRoles) | **Get** /rbac-manager/roles | List roles available in the system
[**ListSecretSearchResults**](DefaultAPI.md#ListSecretSearchResults) | **Get** /images/{image_digest}/artifacts/secret-search | Return a list of analyzer artifacts of the specified type
[**ListSelectors**](DefaultAPI.md#ListSelectors) | **Get** /notifications/selectors | 
[**ListSlackConfigurations**](DefaultAPI.md#ListSlackConfigurations) | **Get** /notifications/endpoints/slack/configurations | 
[**ListSlackSelectors**](DefaultAPI.md#ListSlackSelectors) | **Get** /notifications/endpoints/slack/configurations/{uuid}/selectors | 
[**ListSmtpConfigurations**](DefaultAPI.md#ListSmtpConfigurations) | **Get** /notifications/endpoints/smtp/configurations | 
[**ListSmtpSelectors**](DefaultAPI.md#ListSmtpSelectors) | **Get** /notifications/endpoints/smtp/configurations/{uuid}/selectors | 
[**ListTeamsConfigurations**](DefaultAPI.md#ListTeamsConfigurations) | **Get** /notifications/endpoints/teams/configurations | 
[**ListTeamsSelectors**](DefaultAPI.md#ListTeamsSelectors) | **Get** /notifications/endpoints/teams/configurations/{uuid}/selectors | 
[**ListUserRoles**](DefaultAPI.md#ListUserRoles) | **Get** /rbac-manager/users/{username}/roles | List the roles for which the requested user is a member
[**ListWebhookConfigurations**](DefaultAPI.md#ListWebhookConfigurations) | **Get** /notifications/endpoints/webhook/configurations | 
[**ListWebhookSelectors**](DefaultAPI.md#ListWebhookSelectors) | **Get** /notifications/endpoints/webhook/configurations/{uuid}/selectors | 
[**MyRoles**](DefaultAPI.md#MyRoles) | **Get** /rbac-manager/my-roles | List the roles for which the authenticated user is a member
[**Ping**](DefaultAPI.md#Ping) | **Get** / | 
[**RevokeOauthToken**](DefaultAPI.md#RevokeOauthToken) | **Post** /oauth/revoke | 
[**SamlLogin**](DefaultAPI.md#SamlLogin) | **Get** /rbac-manager/saml/login/{idp_name} | 
[**SamlSso**](DefaultAPI.md#SamlSso) | **Post** /rbac-manager/saml/sso/{idp_name} | 
[**TestGithubConfiguration**](DefaultAPI.md#TestGithubConfiguration) | **Post** /notifications/endpoints/github/test | 
[**TestJiraConfiguration**](DefaultAPI.md#TestJiraConfiguration) | **Post** /notifications/endpoints/jira/test | 
[**TestSlackConfiguration**](DefaultAPI.md#TestSlackConfiguration) | **Post** /notifications/endpoints/slack/test | 
[**TestSmtpConfiguration**](DefaultAPI.md#TestSmtpConfiguration) | **Post** /notifications/endpoints/smtp/test | 
[**TestStoredGithubConfiguration**](DefaultAPI.md#TestStoredGithubConfiguration) | **Get** /notifications/endpoints/github/configurations/{uuid}/test | 
[**TestStoredJiraConfiguration**](DefaultAPI.md#TestStoredJiraConfiguration) | **Get** /notifications/endpoints/jira/configurations/{uuid}/test | 
[**TestStoredSlackConfiguration**](DefaultAPI.md#TestStoredSlackConfiguration) | **Get** /notifications/endpoints/slack/configurations/{uuid}/test | 
[**TestStoredSmtpConfiguration**](DefaultAPI.md#TestStoredSmtpConfiguration) | **Get** /notifications/endpoints/smtp/configurations/{uuid}/test | 
[**TestStoredTeamsConfiguration**](DefaultAPI.md#TestStoredTeamsConfiguration) | **Get** /notifications/endpoints/teams/configurations/{uuid}/test | 
[**TestStoredWebhookConfiguration**](DefaultAPI.md#TestStoredWebhookConfiguration) | **Get** /notifications/endpoints/webhook/configurations/{uuid}/test | 
[**TestTeamsConfiguration**](DefaultAPI.md#TestTeamsConfiguration) | **Post** /notifications/endpoints/teams/test | 
[**TestWebhookConfiguration**](DefaultAPI.md#TestWebhookConfiguration) | **Post** /notifications/endpoints/webhook/test | 
[**UpdateEndpointStatus**](DefaultAPI.md#UpdateEndpointStatus) | **Put** /notifications/endpoints/{name} | 
[**UpdateGithubConfiguration**](DefaultAPI.md#UpdateGithubConfiguration) | **Put** /notifications/endpoints/github/configurations/{uuid} | 
[**UpdateGithubSelector**](DefaultAPI.md#UpdateGithubSelector) | **Put** /notifications/endpoints/github/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**UpdateIdp**](DefaultAPI.md#UpdateIdp) | **Put** /rbac-manager/saml/idps/{name} | 
[**UpdateJiraConfiguration**](DefaultAPI.md#UpdateJiraConfiguration) | **Put** /notifications/endpoints/jira/configurations/{uuid} | 
[**UpdateJiraSelector**](DefaultAPI.md#UpdateJiraSelector) | **Put** /notifications/endpoints/jira/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**UpdateSlackConfiguration**](DefaultAPI.md#UpdateSlackConfiguration) | **Put** /notifications/endpoints/slack/configurations/{uuid} | 
[**UpdateSlackSelector**](DefaultAPI.md#UpdateSlackSelector) | **Put** /notifications/endpoints/slack/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**UpdateSmtpConfiguration**](DefaultAPI.md#UpdateSmtpConfiguration) | **Put** /notifications/endpoints/smtp/configurations/{uuid} | 
[**UpdateSmtpSelector**](DefaultAPI.md#UpdateSmtpSelector) | **Put** /notifications/endpoints/smtp/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**UpdateTeamsConfiguration**](DefaultAPI.md#UpdateTeamsConfiguration) | **Put** /notifications/endpoints/teams/configurations/{uuid} | 
[**UpdateTeamsSelector**](DefaultAPI.md#UpdateTeamsSelector) | **Put** /notifications/endpoints/teams/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**UpdateWebhookConfiguration**](DefaultAPI.md#UpdateWebhookConfiguration) | **Put** /notifications/endpoints/webhook/configurations/{uuid} | 
[**UpdateWebhookSelector**](DefaultAPI.md#UpdateWebhookSelector) | **Put** /notifications/endpoints/webhook/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**VersionCheck**](DefaultAPI.md#VersionCheck) | **Get** /version | 



## AddGithubConfiguration

> NotificationGitHubEndpointConfigurationBase AddGithubConfiguration(ctx).Configuration(configuration).Execute()





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
    configuration := *openapiclient.NewNotificationGitHubEndpointConfigurationPost("Username_example", "AccessToken_example", "Owner_example", "Repository_example") // NotificationGitHubEndpointConfigurationPost | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.AddGithubConfiguration(context.Background()).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddGithubConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddGithubConfiguration`: NotificationGitHubEndpointConfigurationBase
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddGithubConfiguration`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 
    selector := *openapiclient.NewNotificationSelector("Scope_example", *openapiclient.NewNotificationEventSelector("Level_example", "ResourceType_example", "Type_example")) // NotificationSelector | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.AddGithubSelector(context.Background(), uuid).Selector(selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddGithubSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddGithubSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddGithubSelector`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    configuration := *openapiclient.NewRbacManagerSamlConfiguration("Name_example", false, "SpEntityId_example", "AcsUrl_example") // RbacManagerSamlConfiguration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.AddIdp(context.Background()).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddIdp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddIdp`: RbacManagerSamlConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddIdp`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    configuration := *openapiclient.NewNotificationJiraEndpointConfigurationPost("Url_example", "Username_example", "Password_example", "ProjectKey_example", "IssueType_example") // NotificationJiraEndpointConfigurationPost | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.AddJiraConfiguration(context.Background()).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddJiraConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddJiraConfiguration`: NotificationJiraEndpointConfigurationBase
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddJiraConfiguration`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 
    selector := *openapiclient.NewNotificationSelector("Scope_example", *openapiclient.NewNotificationEventSelector("Level_example", "ResourceType_example", "Type_example")) // NotificationSelector | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.AddJiraSelector(context.Background(), uuid).Selector(selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddJiraSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddJiraSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddJiraSelector`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    roleName := "roleName_example" // string | 
    member := *openapiclient.NewRbacManagerRoleMember("Username_example", "ForAccount_example") // RbacManagerRoleMember | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.AddRoleUser(context.Background(), roleName).Member(member).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddRoleUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRoleUser`: RbacManagerRoleMember
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddRoleUser`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    configuration := *openapiclient.NewNotificationSlackEndpointConfiguration("Url_example") // NotificationSlackEndpointConfiguration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.AddSlackConfiguration(context.Background()).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddSlackConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSlackConfiguration`: NotificationSlackEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddSlackConfiguration`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 
    selector := *openapiclient.NewNotificationSelector("Scope_example", *openapiclient.NewNotificationEventSelector("Level_example", "ResourceType_example", "Type_example")) // NotificationSelector | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.AddSlackSelector(context.Background(), uuid).Selector(selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddSlackSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSlackSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddSlackSelector`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    configuration := *openapiclient.NewNotificationSMTPEndpointConfiguration("Host_example", int32(123), "From_example", "To_example") // NotificationSMTPEndpointConfiguration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.AddSmtpConfiguration(context.Background()).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddSmtpConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSmtpConfiguration`: NotificationSMTPEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddSmtpConfiguration`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 
    selector := *openapiclient.NewNotificationSelector("Scope_example", *openapiclient.NewNotificationEventSelector("Level_example", "ResourceType_example", "Type_example")) // NotificationSelector | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.AddSmtpSelector(context.Background(), uuid).Selector(selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddSmtpSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSmtpSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddSmtpSelector`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    configuration := *openapiclient.NewNotificationTeamsEndpointConfiguration("Url_example") // NotificationTeamsEndpointConfiguration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.AddTeamsConfiguration(context.Background()).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddTeamsConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTeamsConfiguration`: NotificationTeamsEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddTeamsConfiguration`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 
    selector := *openapiclient.NewNotificationSelector("Scope_example", *openapiclient.NewNotificationEventSelector("Level_example", "ResourceType_example", "Type_example")) // NotificationSelector | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.AddTeamsSelector(context.Background(), uuid).Selector(selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddTeamsSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTeamsSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddTeamsSelector`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    configuration := *openapiclient.NewNotificationWebhookEndpointConfiguration("Url_example") // NotificationWebhookEndpointConfiguration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.AddWebhookConfiguration(context.Background()).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddWebhookConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddWebhookConfiguration`: NotificationWebhookEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddWebhookConfiguration`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 
    selector := *openapiclient.NewNotificationSelector("Scope_example", *openapiclient.NewNotificationEventSelector("Level_example", "ResourceType_example", "Type_example")) // NotificationSelector | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.AddWebhookSelector(context.Background(), uuid).Selector(selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddWebhookSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddWebhookSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddWebhookSelector`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.DeleteGithubConfiguration(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteGithubConfiguration``: %v\n", err)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    configurationUuid := "configurationUuid_example" // string | 
    selectorUuid := "selectorUuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.DeleteGithubSelector(context.Background(), configurationUuid, selectorUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteGithubSelector``: %v\n", err)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.DeleteIdp(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteIdp``: %v\n", err)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.DeleteJiraConfiguration(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteJiraConfiguration``: %v\n", err)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    configurationUuid := "configurationUuid_example" // string | 
    selectorUuid := "selectorUuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.DeleteJiraSelector(context.Background(), configurationUuid, selectorUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteJiraSelector``: %v\n", err)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    roleName := "roleName_example" // string | 
    username := "username_example" // string | The username to remove the role for
    forAccount := "forAccount_example" // string | The account that the user has the role to be removed

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.DeleteRoleUser(context.Background(), roleName).Username(username).ForAccount(forAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteRoleUser``: %v\n", err)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.DeleteSlackConfiguration(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteSlackConfiguration``: %v\n", err)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    configurationUuid := "configurationUuid_example" // string | 
    selectorUuid := "selectorUuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.DeleteSlackSelector(context.Background(), configurationUuid, selectorUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteSlackSelector``: %v\n", err)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.DeleteSmtpConfiguration(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteSmtpConfiguration``: %v\n", err)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    configurationUuid := "configurationUuid_example" // string | 
    selectorUuid := "selectorUuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.DeleteSmtpSelector(context.Background(), configurationUuid, selectorUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteSmtpSelector``: %v\n", err)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.DeleteTeamsConfiguration(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteTeamsConfiguration``: %v\n", err)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    configurationUuid := "configurationUuid_example" // string | 
    selectorUuid := "selectorUuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.DeleteTeamsSelector(context.Background(), configurationUuid, selectorUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteTeamsSelector``: %v\n", err)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.DeleteWebhookConfiguration(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteWebhookConfiguration``: %v\n", err)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    configurationUuid := "configurationUuid_example" // string | 
    selectorUuid := "selectorUuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.DeleteWebhookSelector(context.Background(), configurationUuid, selectorUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteWebhookSelector``: %v\n", err)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetGithubConfiguration(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetGithubConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGithubConfiguration`: NotificationGitHubEndpointConfigurationBase
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetGithubConfiguration`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetGithubConfigurationStatus(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetGithubConfigurationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGithubConfigurationStatus`: NotificationOperationalStatus
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetGithubConfigurationStatus`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    configurationUuid := "configurationUuid_example" // string | 
    selectorUuid := "selectorUuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetGithubSelector(context.Background(), configurationUuid, selectorUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetGithubSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGithubSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetGithubSelector`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    resultUuid := "resultUuid_example" // string | 
    page := int32(56) // int32 | Page number to fetch. If omitted, '1' is default. Page numbers start at 1 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.GetGlobalQueryResult(context.Background(), resultUuid).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetGlobalQueryResult``: %v\n", err)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetIdp(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetIdp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdp`: RbacManagerSamlConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetIdp`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetJiraConfiguration(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetJiraConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJiraConfiguration`: NotificationJiraEndpointConfigurationBase
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetJiraConfiguration`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetJiraConfigurationStatus(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetJiraConfigurationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJiraConfigurationStatus`: NotificationOperationalStatus
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetJiraConfigurationStatus`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    configurationUuid := "configurationUuid_example" // string | 
    selectorUuid := "selectorUuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetJiraSelector(context.Background(), configurationUuid, selectorUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetJiraSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJiraSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetJiraSelector`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultAPI.GetOauthToken(context.Background()).GrantType(grantType).Username(username).Password(password).ClientId(clientId).RefreshToken(refreshToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetOauthToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOauthToken`: TokenResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetOauthToken`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    resultUuid := "resultUuid_example" // string | 
    page := int32(56) // int32 | Page number to fetch. If omitted, '1' is default. Page numbers start at 1 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.GetQueryResult(context.Background(), resultUuid).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetQueryResult``: %v\n", err)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    roleName := "roleName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetRole(context.Background(), roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRole`: RbacManagerRole
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRole`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetSlackConfiguration(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSlackConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSlackConfiguration`: NotificationSlackEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSlackConfiguration`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetSlackConfigurationStatus(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSlackConfigurationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSlackConfigurationStatus`: NotificationOperationalStatus
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSlackConfigurationStatus`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    configurationUuid := "configurationUuid_example" // string | 
    selectorUuid := "selectorUuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetSlackSelector(context.Background(), configurationUuid, selectorUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSlackSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSlackSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSlackSelector`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetSmtpConfiguration(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSmtpConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSmtpConfiguration`: NotificationSMTPEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSmtpConfiguration`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetSmtpConfigurationStatus(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSmtpConfigurationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSmtpConfigurationStatus`: NotificationOperationalStatus
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSmtpConfigurationStatus`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    configurationUuid := "configurationUuid_example" // string | 
    selectorUuid := "selectorUuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetSmtpSelector(context.Background(), configurationUuid, selectorUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSmtpSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSmtpSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSmtpSelector`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetTeamsConfiguration(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetTeamsConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTeamsConfiguration`: NotificationTeamsEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetTeamsConfiguration`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetTeamsConfigurationStatus(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetTeamsConfigurationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTeamsConfigurationStatus`: NotificationOperationalStatus
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetTeamsConfigurationStatus`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    configurationUuid := "configurationUuid_example" // string | 
    selectorUuid := "selectorUuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetTeamsSelector(context.Background(), configurationUuid, selectorUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetTeamsSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTeamsSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetTeamsSelector`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetWebhookConfiguration(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetWebhookConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhookConfiguration`: NotificationWebhookEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetWebhookConfiguration`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetWebhookConfigurationStatus(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetWebhookConfigurationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhookConfigurationStatus`: NotificationOperationalStatus
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetWebhookConfigurationStatus`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    configurationUuid := "configurationUuid_example" // string | 
    selectorUuid := "selectorUuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetWebhookSelector(context.Background(), configurationUuid, selectorUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetWebhookSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhookSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetWebhookSelector`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.HealthCheck(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.HealthCheck``: %v\n", err)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ListEndpoints(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListEndpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEndpoints`: []NotificationEndpoint
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListEndpoints`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    imageDigest := "imageDigest_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ListFileContentSearchResults(context.Background(), imageDigest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListFileContentSearchResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFileContentSearchResults`: []FileContentSearchResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListFileContentSearchResults`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ListGithubConfigurations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListGithubConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGithubConfigurations`: []NotificationGitHubEndpointConfigurationBase
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListGithubConfigurations`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ListGithubSelectors(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListGithubSelectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGithubSelectors`: []NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListGithubSelectors`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ListIdps(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListIdps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIdps`: []string
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListIdps`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ListJiraConfigurations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListJiraConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListJiraConfigurations`: []NotificationJiraEndpointConfigurationBase
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListJiraConfigurations`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ListJiraSelectors(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListJiraSelectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListJiraSelectors`: []NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListJiraSelectors`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    imageDigest := "imageDigest_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ListRetrievedFiles(context.Background(), imageDigest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListRetrievedFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRetrievedFiles`: []RetrievedFile
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListRetrievedFiles`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    roleName := "roleName_example" // string | 
    forAccount := "forAccount_example" // string | Optional filter parameter to limit the set fo returned items to only those with matching account. Will return Access Denied if caller does not have permission to listRoleMembers for that account. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ListRoleMembers(context.Background(), roleName).ForAccount(forAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListRoleMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoleMembers`: []RbacManagerRoleMember
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListRoleMembers`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ListRoles(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoles`: []RbacManagerRoleSummary
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListRoles`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    imageDigest := "imageDigest_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ListSecretSearchResults(context.Background(), imageDigest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListSecretSearchResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSecretSearchResults`: []SecretSearchResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListSecretSearchResults`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ListSelectors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListSelectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSelectors`: []NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListSelectors`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ListSlackConfigurations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListSlackConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSlackConfigurations`: []NotificationSlackEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListSlackConfigurations`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ListSlackSelectors(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListSlackSelectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSlackSelectors`: []NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListSlackSelectors`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ListSmtpConfigurations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListSmtpConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSmtpConfigurations`: []NotificationSMTPEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListSmtpConfigurations`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ListSmtpSelectors(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListSmtpSelectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSmtpSelectors`: []NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListSmtpSelectors`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ListTeamsConfigurations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListTeamsConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTeamsConfigurations`: []NotificationTeamsEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListTeamsConfigurations`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ListTeamsSelectors(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListTeamsSelectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTeamsSelectors`: []NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListTeamsSelectors`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    username := "username_example" // string | 
    forAccount := "forAccount_example" // string |  (optional)
    role := "role_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ListUserRoles(context.Background(), username).ForAccount(forAccount).Role(role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListUserRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserRoles`: []RbacManagerRoleMembership
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListUserRoles`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ListWebhookConfigurations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListWebhookConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWebhookConfigurations`: []NotificationWebhookEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListWebhookConfigurations`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ListWebhookSelectors(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListWebhookSelectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWebhookSelectors`: []NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListWebhookSelectors`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.MyRoles(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.MyRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MyRoles`: []RbacManagerAccountRole
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.MyRoles`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.Ping(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Ping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Ping`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Ping`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    token := "token_example" // string | The token to be revoked (optional)
    tokenTypeHint := "tokenTypeHint_example" // string | A hint about the type of token to be revoked (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.RevokeOauthToken(context.Background()).Token(token).TokenTypeHint(tokenTypeHint).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RevokeOauthToken``: %v\n", err)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    idpName := "idpName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.SamlLogin(context.Background(), idpName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SamlLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SamlLogin`: RbacManagerTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.SamlLogin`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    idpName := "idpName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.SamlSso(context.Background(), idpName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SamlSso``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SamlSso`: RbacManagerTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.SamlSso`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    configuration := *openapiclient.NewNotificationGitHubEndpointConfigurationPost("Username_example", "AccessToken_example", "Owner_example", "Repository_example") // NotificationGitHubEndpointConfigurationPost | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.TestGithubConfiguration(context.Background()).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TestGithubConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestGithubConfiguration`: NotificationGitHubTestResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.TestGithubConfiguration`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    configuration := *openapiclient.NewNotificationJiraEndpointConfigurationPost("Url_example", "Username_example", "Password_example", "ProjectKey_example", "IssueType_example") // NotificationJiraEndpointConfigurationPost | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.TestJiraConfiguration(context.Background()).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TestJiraConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestJiraConfiguration`: NotificationJiraTestResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.TestJiraConfiguration`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    configuration := *openapiclient.NewNotificationSlackEndpointConfiguration("Url_example") // NotificationSlackEndpointConfiguration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.TestSlackConfiguration(context.Background()).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TestSlackConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestSlackConfiguration`: NotificationSlackTestResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.TestSlackConfiguration`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    configuration := *openapiclient.NewNotificationSMTPEndpointConfiguration("Host_example", int32(123), "From_example", "To_example") // NotificationSMTPEndpointConfiguration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.TestSmtpConfiguration(context.Background()).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TestSmtpConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestSmtpConfiguration`: NotificationSMTPTestResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.TestSmtpConfiguration`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.TestStoredGithubConfiguration(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TestStoredGithubConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestStoredGithubConfiguration`: NotificationGitHubTestResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.TestStoredGithubConfiguration`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.TestStoredJiraConfiguration(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TestStoredJiraConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestStoredJiraConfiguration`: NotificationJiraTestResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.TestStoredJiraConfiguration`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.TestStoredSlackConfiguration(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TestStoredSlackConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestStoredSlackConfiguration`: NotificationSlackTestResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.TestStoredSlackConfiguration`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.TestStoredSmtpConfiguration(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TestStoredSmtpConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestStoredSmtpConfiguration`: NotificationSMTPTestResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.TestStoredSmtpConfiguration`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.TestStoredTeamsConfiguration(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TestStoredTeamsConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestStoredTeamsConfiguration`: NotificationTeamsTestResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.TestStoredTeamsConfiguration`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.TestStoredWebhookConfiguration(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TestStoredWebhookConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestStoredWebhookConfiguration`: NotificationWebhookTestResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.TestStoredWebhookConfiguration`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    configuration := *openapiclient.NewNotificationTeamsEndpointConfiguration("Url_example") // NotificationTeamsEndpointConfiguration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.TestTeamsConfiguration(context.Background()).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TestTeamsConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestTeamsConfiguration`: NotificationTeamsTestResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.TestTeamsConfiguration`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    configuration := *openapiclient.NewNotificationWebhookEndpointConfiguration("Url_example") // NotificationWebhookEndpointConfiguration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.TestWebhookConfiguration(context.Background()).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TestWebhookConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestWebhookConfiguration`: NotificationWebhookTestResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.TestWebhookConfiguration`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    name := "name_example" // string | 
    status := *openapiclient.NewNotificationEndpointEnabledStatus() // NotificationEndpointEnabledStatus | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateEndpointStatus(context.Background(), name).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateEndpointStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEndpointStatus`: NotificationEndpointEnabledStatus
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateEndpointStatus`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 
    configuration := *openapiclient.NewNotificationGitHubEndpointConfigurationPut("Username_example", "Owner_example", "Repository_example") // NotificationGitHubEndpointConfigurationPut | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateGithubConfiguration(context.Background(), uuid).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateGithubConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGithubConfiguration`: NotificationGitHubEndpointConfigurationBase
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateGithubConfiguration`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    configurationUuid := "configurationUuid_example" // string | 
    selectorUuid := "selectorUuid_example" // string | 
    selector := *openapiclient.NewNotificationSelector("Scope_example", *openapiclient.NewNotificationEventSelector("Level_example", "ResourceType_example", "Type_example")) // NotificationSelector | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateGithubSelector(context.Background(), configurationUuid, selectorUuid).Selector(selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateGithubSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGithubSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateGithubSelector`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    name := "name_example" // string | 
    configuration := *openapiclient.NewRbacManagerSamlConfiguration("Name_example", false, "SpEntityId_example", "AcsUrl_example") // RbacManagerSamlConfiguration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateIdp(context.Background(), name).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateIdp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIdp`: RbacManagerSamlConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateIdp`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 
    configuration := *openapiclient.NewNotificationJiraEndpointConfigurationPut("Url_example", "Username_example", "ProjectKey_example", "IssueType_example") // NotificationJiraEndpointConfigurationPut | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateJiraConfiguration(context.Background(), uuid).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateJiraConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateJiraConfiguration`: NotificationJiraEndpointConfigurationBase
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateJiraConfiguration`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    configurationUuid := "configurationUuid_example" // string | 
    selectorUuid := "selectorUuid_example" // string | 
    selector := *openapiclient.NewNotificationSelector("Scope_example", *openapiclient.NewNotificationEventSelector("Level_example", "ResourceType_example", "Type_example")) // NotificationSelector | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateJiraSelector(context.Background(), configurationUuid, selectorUuid).Selector(selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateJiraSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateJiraSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateJiraSelector`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 
    configuration := *openapiclient.NewNotificationSlackEndpointConfiguration("Url_example") // NotificationSlackEndpointConfiguration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateSlackConfiguration(context.Background(), uuid).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateSlackConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSlackConfiguration`: NotificationSlackEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateSlackConfiguration`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    configurationUuid := "configurationUuid_example" // string | 
    selectorUuid := "selectorUuid_example" // string | 
    selector := *openapiclient.NewNotificationSelector("Scope_example", *openapiclient.NewNotificationEventSelector("Level_example", "ResourceType_example", "Type_example")) // NotificationSelector | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateSlackSelector(context.Background(), configurationUuid, selectorUuid).Selector(selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateSlackSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSlackSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateSlackSelector`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 
    configuration := *openapiclient.NewNotificationSMTPEndpointConfiguration("Host_example", int32(123), "From_example", "To_example") // NotificationSMTPEndpointConfiguration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateSmtpConfiguration(context.Background(), uuid).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateSmtpConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSmtpConfiguration`: NotificationSMTPEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateSmtpConfiguration`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    configurationUuid := "configurationUuid_example" // string | 
    selectorUuid := "selectorUuid_example" // string | 
    selector := *openapiclient.NewNotificationSelector("Scope_example", *openapiclient.NewNotificationEventSelector("Level_example", "ResourceType_example", "Type_example")) // NotificationSelector | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateSmtpSelector(context.Background(), configurationUuid, selectorUuid).Selector(selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateSmtpSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSmtpSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateSmtpSelector`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 
    configuration := *openapiclient.NewNotificationTeamsEndpointConfiguration("Url_example") // NotificationTeamsEndpointConfiguration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateTeamsConfiguration(context.Background(), uuid).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateTeamsConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTeamsConfiguration`: NotificationTeamsEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateTeamsConfiguration`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    configurationUuid := "configurationUuid_example" // string | 
    selectorUuid := "selectorUuid_example" // string | 
    selector := *openapiclient.NewNotificationSelector("Scope_example", *openapiclient.NewNotificationEventSelector("Level_example", "ResourceType_example", "Type_example")) // NotificationSelector | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateTeamsSelector(context.Background(), configurationUuid, selectorUuid).Selector(selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateTeamsSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTeamsSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateTeamsSelector`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    uuid := "uuid_example" // string | 
    configuration := *openapiclient.NewNotificationWebhookEndpointConfiguration("Url_example") // NotificationWebhookEndpointConfiguration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateWebhookConfiguration(context.Background(), uuid).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateWebhookConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWebhookConfiguration`: NotificationWebhookEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateWebhookConfiguration`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    configurationUuid := "configurationUuid_example" // string | 
    selectorUuid := "selectorUuid_example" // string | 
    selector := *openapiclient.NewNotificationSelector("Scope_example", *openapiclient.NewNotificationEventSelector("Level_example", "ResourceType_example", "Type_example")) // NotificationSelector | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateWebhookSelector(context.Background(), configurationUuid, selectorUuid).Selector(selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateWebhookSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWebhookSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateWebhookSelector`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.VersionCheck(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.VersionCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VersionCheck`: ServiceVersion
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.VersionCheck`: %v\n", resp)
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

