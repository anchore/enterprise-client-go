# \NotificationsApi

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGithubConfiguration**](NotificationsApi.md#AddGithubConfiguration) | **Post** /notifications/endpoints/github/configurations | 
[**AddGithubSelector**](NotificationsApi.md#AddGithubSelector) | **Post** /notifications/endpoints/github/configurations/{uuid}/selectors | 
[**AddJiraConfiguration**](NotificationsApi.md#AddJiraConfiguration) | **Post** /notifications/endpoints/jira/configurations | 
[**AddJiraSelector**](NotificationsApi.md#AddJiraSelector) | **Post** /notifications/endpoints/jira/configurations/{uuid}/selectors | 
[**AddSlackConfiguration**](NotificationsApi.md#AddSlackConfiguration) | **Post** /notifications/endpoints/slack/configurations | 
[**AddSlackSelector**](NotificationsApi.md#AddSlackSelector) | **Post** /notifications/endpoints/slack/configurations/{uuid}/selectors | 
[**AddSmtpConfiguration**](NotificationsApi.md#AddSmtpConfiguration) | **Post** /notifications/endpoints/smtp/configurations | 
[**AddSmtpSelector**](NotificationsApi.md#AddSmtpSelector) | **Post** /notifications/endpoints/smtp/configurations/{uuid}/selectors | 
[**AddTeamsConfiguration**](NotificationsApi.md#AddTeamsConfiguration) | **Post** /notifications/endpoints/teams/configurations | 
[**AddTeamsSelector**](NotificationsApi.md#AddTeamsSelector) | **Post** /notifications/endpoints/teams/configurations/{uuid}/selectors | 
[**AddWebhookConfiguration**](NotificationsApi.md#AddWebhookConfiguration) | **Post** /notifications/endpoints/webhook/configurations | 
[**AddWebhookSelector**](NotificationsApi.md#AddWebhookSelector) | **Post** /notifications/endpoints/webhook/configurations/{uuid}/selectors | 
[**DeleteGithubConfiguration**](NotificationsApi.md#DeleteGithubConfiguration) | **Delete** /notifications/endpoints/github/configurations/{uuid} | 
[**DeleteGithubSelector**](NotificationsApi.md#DeleteGithubSelector) | **Delete** /notifications/endpoints/github/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**DeleteJiraConfiguration**](NotificationsApi.md#DeleteJiraConfiguration) | **Delete** /notifications/endpoints/jira/configurations/{uuid} | 
[**DeleteJiraSelector**](NotificationsApi.md#DeleteJiraSelector) | **Delete** /notifications/endpoints/jira/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**DeleteSlackConfiguration**](NotificationsApi.md#DeleteSlackConfiguration) | **Delete** /notifications/endpoints/slack/configurations/{uuid} | 
[**DeleteSlackSelector**](NotificationsApi.md#DeleteSlackSelector) | **Delete** /notifications/endpoints/slack/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**DeleteSmtpConfiguration**](NotificationsApi.md#DeleteSmtpConfiguration) | **Delete** /notifications/endpoints/smtp/configurations/{uuid} | 
[**DeleteSmtpSelector**](NotificationsApi.md#DeleteSmtpSelector) | **Delete** /notifications/endpoints/smtp/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**DeleteTeamsConfiguration**](NotificationsApi.md#DeleteTeamsConfiguration) | **Delete** /notifications/endpoints/teams/configurations/{uuid} | 
[**DeleteTeamsSelector**](NotificationsApi.md#DeleteTeamsSelector) | **Delete** /notifications/endpoints/teams/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**DeleteWebhookConfiguration**](NotificationsApi.md#DeleteWebhookConfiguration) | **Delete** /notifications/endpoints/webhook/configurations/{uuid} | 
[**DeleteWebhookSelector**](NotificationsApi.md#DeleteWebhookSelector) | **Delete** /notifications/endpoints/webhook/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**GetGithubConfiguration**](NotificationsApi.md#GetGithubConfiguration) | **Get** /notifications/endpoints/github/configurations/{uuid} | 
[**GetGithubConfigurationStatus**](NotificationsApi.md#GetGithubConfigurationStatus) | **Get** /notifications/endpoints/github/configurations/{uuid}/status | 
[**GetGithubSelector**](NotificationsApi.md#GetGithubSelector) | **Get** /notifications/endpoints/github/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**GetJiraConfiguration**](NotificationsApi.md#GetJiraConfiguration) | **Get** /notifications/endpoints/jira/configurations/{uuid} | 
[**GetJiraConfigurationStatus**](NotificationsApi.md#GetJiraConfigurationStatus) | **Get** /notifications/endpoints/jira/configurations/{uuid}/status | 
[**GetJiraSelector**](NotificationsApi.md#GetJiraSelector) | **Get** /notifications/endpoints/jira/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**GetSlackConfiguration**](NotificationsApi.md#GetSlackConfiguration) | **Get** /notifications/endpoints/slack/configurations/{uuid} | 
[**GetSlackConfigurationStatus**](NotificationsApi.md#GetSlackConfigurationStatus) | **Get** /notifications/endpoints/slack/configurations/{uuid}/status | 
[**GetSlackSelector**](NotificationsApi.md#GetSlackSelector) | **Get** /notifications/endpoints/slack/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**GetSmtpConfiguration**](NotificationsApi.md#GetSmtpConfiguration) | **Get** /notifications/endpoints/smtp/configurations/{uuid} | 
[**GetSmtpConfigurationStatus**](NotificationsApi.md#GetSmtpConfigurationStatus) | **Get** /notifications/endpoints/smtp/configurations/{uuid}/status | 
[**GetSmtpSelector**](NotificationsApi.md#GetSmtpSelector) | **Get** /notifications/endpoints/smtp/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**GetTeamsConfiguration**](NotificationsApi.md#GetTeamsConfiguration) | **Get** /notifications/endpoints/teams/configurations/{uuid} | 
[**GetTeamsConfigurationStatus**](NotificationsApi.md#GetTeamsConfigurationStatus) | **Get** /notifications/endpoints/teams/configurations/{uuid}/status | 
[**GetTeamsSelector**](NotificationsApi.md#GetTeamsSelector) | **Get** /notifications/endpoints/teams/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**GetWebhookConfiguration**](NotificationsApi.md#GetWebhookConfiguration) | **Get** /notifications/endpoints/webhook/configurations/{uuid} | 
[**GetWebhookConfigurationStatus**](NotificationsApi.md#GetWebhookConfigurationStatus) | **Get** /notifications/endpoints/webhook/configurations/{uuid}/status | 
[**GetWebhookSelector**](NotificationsApi.md#GetWebhookSelector) | **Get** /notifications/endpoints/webhook/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**ListEndpoints**](NotificationsApi.md#ListEndpoints) | **Get** /notifications/endpoints | 
[**ListGithubConfigurations**](NotificationsApi.md#ListGithubConfigurations) | **Get** /notifications/endpoints/github/configurations | 
[**ListGithubSelectors**](NotificationsApi.md#ListGithubSelectors) | **Get** /notifications/endpoints/github/configurations/{uuid}/selectors | 
[**ListJiraConfigurations**](NotificationsApi.md#ListJiraConfigurations) | **Get** /notifications/endpoints/jira/configurations | 
[**ListJiraSelectors**](NotificationsApi.md#ListJiraSelectors) | **Get** /notifications/endpoints/jira/configurations/{uuid}/selectors | 
[**ListSelectors**](NotificationsApi.md#ListSelectors) | **Get** /notifications/selectors | 
[**ListSlackConfigurations**](NotificationsApi.md#ListSlackConfigurations) | **Get** /notifications/endpoints/slack/configurations | 
[**ListSlackSelectors**](NotificationsApi.md#ListSlackSelectors) | **Get** /notifications/endpoints/slack/configurations/{uuid}/selectors | 
[**ListSmtpConfigurations**](NotificationsApi.md#ListSmtpConfigurations) | **Get** /notifications/endpoints/smtp/configurations | 
[**ListSmtpSelectors**](NotificationsApi.md#ListSmtpSelectors) | **Get** /notifications/endpoints/smtp/configurations/{uuid}/selectors | 
[**ListTeamsConfigurations**](NotificationsApi.md#ListTeamsConfigurations) | **Get** /notifications/endpoints/teams/configurations | 
[**ListTeamsSelectors**](NotificationsApi.md#ListTeamsSelectors) | **Get** /notifications/endpoints/teams/configurations/{uuid}/selectors | 
[**ListWebhookConfigurations**](NotificationsApi.md#ListWebhookConfigurations) | **Get** /notifications/endpoints/webhook/configurations | 
[**ListWebhookSelectors**](NotificationsApi.md#ListWebhookSelectors) | **Get** /notifications/endpoints/webhook/configurations/{uuid}/selectors | 
[**TestGithubConfiguration**](NotificationsApi.md#TestGithubConfiguration) | **Post** /notifications/endpoints/github/test | 
[**TestJiraConfiguration**](NotificationsApi.md#TestJiraConfiguration) | **Post** /notifications/endpoints/jira/test | 
[**TestSlackConfiguration**](NotificationsApi.md#TestSlackConfiguration) | **Post** /notifications/endpoints/slack/test | 
[**TestSmtpConfiguration**](NotificationsApi.md#TestSmtpConfiguration) | **Post** /notifications/endpoints/smtp/test | 
[**TestStoredGithubConfiguration**](NotificationsApi.md#TestStoredGithubConfiguration) | **Get** /notifications/endpoints/github/configurations/{uuid}/test | 
[**TestStoredJiraConfiguration**](NotificationsApi.md#TestStoredJiraConfiguration) | **Get** /notifications/endpoints/jira/configurations/{uuid}/test | 
[**TestStoredSlackConfiguration**](NotificationsApi.md#TestStoredSlackConfiguration) | **Get** /notifications/endpoints/slack/configurations/{uuid}/test | 
[**TestStoredSmtpConfiguration**](NotificationsApi.md#TestStoredSmtpConfiguration) | **Get** /notifications/endpoints/smtp/configurations/{uuid}/test | 
[**TestStoredTeamsConfiguration**](NotificationsApi.md#TestStoredTeamsConfiguration) | **Get** /notifications/endpoints/teams/configurations/{uuid}/test | 
[**TestStoredWebhookConfiguration**](NotificationsApi.md#TestStoredWebhookConfiguration) | **Get** /notifications/endpoints/webhook/configurations/{uuid}/test | 
[**TestTeamsConfiguration**](NotificationsApi.md#TestTeamsConfiguration) | **Post** /notifications/endpoints/teams/test | 
[**TestWebhookConfiguration**](NotificationsApi.md#TestWebhookConfiguration) | **Post** /notifications/endpoints/webhook/test | 
[**UpdateEndpointStatus**](NotificationsApi.md#UpdateEndpointStatus) | **Put** /notifications/endpoints/{name} | 
[**UpdateGithubConfiguration**](NotificationsApi.md#UpdateGithubConfiguration) | **Put** /notifications/endpoints/github/configurations/{uuid} | 
[**UpdateGithubSelector**](NotificationsApi.md#UpdateGithubSelector) | **Put** /notifications/endpoints/github/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**UpdateJiraConfiguration**](NotificationsApi.md#UpdateJiraConfiguration) | **Put** /notifications/endpoints/jira/configurations/{uuid} | 
[**UpdateJiraSelector**](NotificationsApi.md#UpdateJiraSelector) | **Put** /notifications/endpoints/jira/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**UpdateSlackConfiguration**](NotificationsApi.md#UpdateSlackConfiguration) | **Put** /notifications/endpoints/slack/configurations/{uuid} | 
[**UpdateSlackSelector**](NotificationsApi.md#UpdateSlackSelector) | **Put** /notifications/endpoints/slack/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**UpdateSmtpConfiguration**](NotificationsApi.md#UpdateSmtpConfiguration) | **Put** /notifications/endpoints/smtp/configurations/{uuid} | 
[**UpdateSmtpSelector**](NotificationsApi.md#UpdateSmtpSelector) | **Put** /notifications/endpoints/smtp/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**UpdateTeamsConfiguration**](NotificationsApi.md#UpdateTeamsConfiguration) | **Put** /notifications/endpoints/teams/configurations/{uuid} | 
[**UpdateTeamsSelector**](NotificationsApi.md#UpdateTeamsSelector) | **Put** /notifications/endpoints/teams/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**UpdateWebhookConfiguration**](NotificationsApi.md#UpdateWebhookConfiguration) | **Put** /notifications/endpoints/webhook/configurations/{uuid} | 
[**UpdateWebhookSelector**](NotificationsApi.md#UpdateWebhookSelector) | **Put** /notifications/endpoints/webhook/configurations/{configuration_uuid}/selectors/{selector_uuid} | 



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.AddGithubConfiguration(context.Background()).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.AddGithubConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddGithubConfiguration`: NotificationGitHubEndpointConfigurationBase
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.AddGithubConfiguration`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.AddGithubSelector(context.Background(), uuid).Selector(selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.AddGithubSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddGithubSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.AddGithubSelector`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.AddJiraConfiguration(context.Background()).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.AddJiraConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddJiraConfiguration`: NotificationJiraEndpointConfigurationBase
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.AddJiraConfiguration`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.AddJiraSelector(context.Background(), uuid).Selector(selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.AddJiraSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddJiraSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.AddJiraSelector`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.AddSlackConfiguration(context.Background()).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.AddSlackConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSlackConfiguration`: NotificationSlackEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.AddSlackConfiguration`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.AddSlackSelector(context.Background(), uuid).Selector(selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.AddSlackSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSlackSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.AddSlackSelector`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.AddSmtpConfiguration(context.Background()).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.AddSmtpConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSmtpConfiguration`: NotificationSMTPEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.AddSmtpConfiguration`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.AddSmtpSelector(context.Background(), uuid).Selector(selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.AddSmtpSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSmtpSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.AddSmtpSelector`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.AddTeamsConfiguration(context.Background()).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.AddTeamsConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTeamsConfiguration`: NotificationTeamsEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.AddTeamsConfiguration`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.AddTeamsSelector(context.Background(), uuid).Selector(selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.AddTeamsSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTeamsSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.AddTeamsSelector`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.AddWebhookConfiguration(context.Background()).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.AddWebhookConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddWebhookConfiguration`: NotificationWebhookEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.AddWebhookConfiguration`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.AddWebhookSelector(context.Background(), uuid).Selector(selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.AddWebhookSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddWebhookSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.AddWebhookSelector`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.DeleteGithubConfiguration(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.DeleteGithubConfiguration``: %v\n", err)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.DeleteGithubSelector(context.Background(), configurationUuid, selectorUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.DeleteGithubSelector``: %v\n", err)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.DeleteJiraConfiguration(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.DeleteJiraConfiguration``: %v\n", err)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.DeleteJiraSelector(context.Background(), configurationUuid, selectorUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.DeleteJiraSelector``: %v\n", err)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.DeleteSlackConfiguration(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.DeleteSlackConfiguration``: %v\n", err)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.DeleteSlackSelector(context.Background(), configurationUuid, selectorUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.DeleteSlackSelector``: %v\n", err)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.DeleteSmtpConfiguration(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.DeleteSmtpConfiguration``: %v\n", err)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.DeleteSmtpSelector(context.Background(), configurationUuid, selectorUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.DeleteSmtpSelector``: %v\n", err)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.DeleteTeamsConfiguration(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.DeleteTeamsConfiguration``: %v\n", err)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.DeleteTeamsSelector(context.Background(), configurationUuid, selectorUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.DeleteTeamsSelector``: %v\n", err)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.DeleteWebhookConfiguration(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.DeleteWebhookConfiguration``: %v\n", err)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.DeleteWebhookSelector(context.Background(), configurationUuid, selectorUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.DeleteWebhookSelector``: %v\n", err)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.GetGithubConfiguration(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.GetGithubConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGithubConfiguration`: NotificationGitHubEndpointConfigurationBase
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.GetGithubConfiguration`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.GetGithubConfigurationStatus(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.GetGithubConfigurationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGithubConfigurationStatus`: NotificationOperationalStatus
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.GetGithubConfigurationStatus`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.GetGithubSelector(context.Background(), configurationUuid, selectorUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.GetGithubSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGithubSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.GetGithubSelector`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.GetJiraConfiguration(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.GetJiraConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJiraConfiguration`: NotificationJiraEndpointConfigurationBase
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.GetJiraConfiguration`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.GetJiraConfigurationStatus(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.GetJiraConfigurationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJiraConfigurationStatus`: NotificationOperationalStatus
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.GetJiraConfigurationStatus`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.GetJiraSelector(context.Background(), configurationUuid, selectorUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.GetJiraSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJiraSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.GetJiraSelector`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.GetSlackConfiguration(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.GetSlackConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSlackConfiguration`: NotificationSlackEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.GetSlackConfiguration`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.GetSlackConfigurationStatus(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.GetSlackConfigurationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSlackConfigurationStatus`: NotificationOperationalStatus
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.GetSlackConfigurationStatus`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.GetSlackSelector(context.Background(), configurationUuid, selectorUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.GetSlackSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSlackSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.GetSlackSelector`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.GetSmtpConfiguration(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.GetSmtpConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSmtpConfiguration`: NotificationSMTPEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.GetSmtpConfiguration`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.GetSmtpConfigurationStatus(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.GetSmtpConfigurationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSmtpConfigurationStatus`: NotificationOperationalStatus
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.GetSmtpConfigurationStatus`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.GetSmtpSelector(context.Background(), configurationUuid, selectorUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.GetSmtpSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSmtpSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.GetSmtpSelector`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.GetTeamsConfiguration(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.GetTeamsConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTeamsConfiguration`: NotificationTeamsEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.GetTeamsConfiguration`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.GetTeamsConfigurationStatus(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.GetTeamsConfigurationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTeamsConfigurationStatus`: NotificationOperationalStatus
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.GetTeamsConfigurationStatus`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.GetTeamsSelector(context.Background(), configurationUuid, selectorUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.GetTeamsSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTeamsSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.GetTeamsSelector`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.GetWebhookConfiguration(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.GetWebhookConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhookConfiguration`: NotificationWebhookEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.GetWebhookConfiguration`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.GetWebhookConfigurationStatus(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.GetWebhookConfigurationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhookConfigurationStatus`: NotificationOperationalStatus
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.GetWebhookConfigurationStatus`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.GetWebhookSelector(context.Background(), configurationUuid, selectorUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.GetWebhookSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhookSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.GetWebhookSelector`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.ListEndpoints(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.ListEndpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEndpoints`: []NotificationEndpoint
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.ListEndpoints`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.ListGithubConfigurations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.ListGithubConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGithubConfigurations`: []NotificationGitHubEndpointConfigurationBase
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.ListGithubConfigurations`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.ListGithubSelectors(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.ListGithubSelectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGithubSelectors`: []NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.ListGithubSelectors`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.ListJiraConfigurations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.ListJiraConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListJiraConfigurations`: []NotificationJiraEndpointConfigurationBase
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.ListJiraConfigurations`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.ListJiraSelectors(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.ListJiraSelectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListJiraSelectors`: []NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.ListJiraSelectors`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.ListSelectors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.ListSelectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSelectors`: []NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.ListSelectors`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.ListSlackConfigurations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.ListSlackConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSlackConfigurations`: []NotificationSlackEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.ListSlackConfigurations`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.ListSlackSelectors(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.ListSlackSelectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSlackSelectors`: []NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.ListSlackSelectors`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.ListSmtpConfigurations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.ListSmtpConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSmtpConfigurations`: []NotificationSMTPEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.ListSmtpConfigurations`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.ListSmtpSelectors(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.ListSmtpSelectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSmtpSelectors`: []NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.ListSmtpSelectors`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.ListTeamsConfigurations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.ListTeamsConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTeamsConfigurations`: []NotificationTeamsEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.ListTeamsConfigurations`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.ListTeamsSelectors(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.ListTeamsSelectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTeamsSelectors`: []NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.ListTeamsSelectors`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.ListWebhookConfigurations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.ListWebhookConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWebhookConfigurations`: []NotificationWebhookEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.ListWebhookConfigurations`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.ListWebhookSelectors(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.ListWebhookSelectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWebhookSelectors`: []NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.ListWebhookSelectors`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.TestGithubConfiguration(context.Background()).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.TestGithubConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestGithubConfiguration`: NotificationGitHubTestResult
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.TestGithubConfiguration`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.TestJiraConfiguration(context.Background()).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.TestJiraConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestJiraConfiguration`: NotificationJiraTestResult
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.TestJiraConfiguration`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.TestSlackConfiguration(context.Background()).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.TestSlackConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestSlackConfiguration`: NotificationSlackTestResult
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.TestSlackConfiguration`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.TestSmtpConfiguration(context.Background()).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.TestSmtpConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestSmtpConfiguration`: NotificationSMTPTestResult
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.TestSmtpConfiguration`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.TestStoredGithubConfiguration(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.TestStoredGithubConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestStoredGithubConfiguration`: NotificationGitHubTestResult
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.TestStoredGithubConfiguration`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.TestStoredJiraConfiguration(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.TestStoredJiraConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestStoredJiraConfiguration`: NotificationJiraTestResult
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.TestStoredJiraConfiguration`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.TestStoredSlackConfiguration(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.TestStoredSlackConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestStoredSlackConfiguration`: NotificationSlackTestResult
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.TestStoredSlackConfiguration`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.TestStoredSmtpConfiguration(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.TestStoredSmtpConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestStoredSmtpConfiguration`: NotificationSMTPTestResult
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.TestStoredSmtpConfiguration`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.TestStoredTeamsConfiguration(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.TestStoredTeamsConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestStoredTeamsConfiguration`: NotificationTeamsTestResult
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.TestStoredTeamsConfiguration`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.TestStoredWebhookConfiguration(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.TestStoredWebhookConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestStoredWebhookConfiguration`: NotificationWebhookTestResult
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.TestStoredWebhookConfiguration`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.TestTeamsConfiguration(context.Background()).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.TestTeamsConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestTeamsConfiguration`: NotificationTeamsTestResult
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.TestTeamsConfiguration`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.TestWebhookConfiguration(context.Background()).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.TestWebhookConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestWebhookConfiguration`: NotificationWebhookTestResult
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.TestWebhookConfiguration`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.UpdateEndpointStatus(context.Background(), name).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.UpdateEndpointStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEndpointStatus`: NotificationEndpointEnabledStatus
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.UpdateEndpointStatus`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.UpdateGithubConfiguration(context.Background(), uuid).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.UpdateGithubConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGithubConfiguration`: NotificationGitHubEndpointConfigurationBase
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.UpdateGithubConfiguration`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.UpdateGithubSelector(context.Background(), configurationUuid, selectorUuid).Selector(selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.UpdateGithubSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGithubSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.UpdateGithubSelector`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.UpdateJiraConfiguration(context.Background(), uuid).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.UpdateJiraConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateJiraConfiguration`: NotificationJiraEndpointConfigurationBase
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.UpdateJiraConfiguration`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.UpdateJiraSelector(context.Background(), configurationUuid, selectorUuid).Selector(selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.UpdateJiraSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateJiraSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.UpdateJiraSelector`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.UpdateSlackConfiguration(context.Background(), uuid).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.UpdateSlackConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSlackConfiguration`: NotificationSlackEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.UpdateSlackConfiguration`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.UpdateSlackSelector(context.Background(), configurationUuid, selectorUuid).Selector(selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.UpdateSlackSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSlackSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.UpdateSlackSelector`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.UpdateSmtpConfiguration(context.Background(), uuid).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.UpdateSmtpConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSmtpConfiguration`: NotificationSMTPEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.UpdateSmtpConfiguration`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.UpdateSmtpSelector(context.Background(), configurationUuid, selectorUuid).Selector(selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.UpdateSmtpSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSmtpSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.UpdateSmtpSelector`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.UpdateTeamsConfiguration(context.Background(), uuid).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.UpdateTeamsConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTeamsConfiguration`: NotificationTeamsEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.UpdateTeamsConfiguration`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.UpdateTeamsSelector(context.Background(), configurationUuid, selectorUuid).Selector(selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.UpdateTeamsSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTeamsSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.UpdateTeamsSelector`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.UpdateWebhookConfiguration(context.Background(), uuid).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.UpdateWebhookConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWebhookConfiguration`: NotificationWebhookEndpointConfiguration
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.UpdateWebhookConfiguration`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.UpdateWebhookSelector(context.Background(), configurationUuid, selectorUuid).Selector(selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.UpdateWebhookSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWebhookSelector`: NotificationSelector
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.UpdateWebhookSelector`: %v\n", resp)
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

