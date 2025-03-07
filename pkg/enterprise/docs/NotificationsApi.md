# \NotificationsAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGithubConfiguration**](NotificationsAPI.md#AddGithubConfiguration) | **Post** /notifications/endpoints/github/configurations | 
[**AddGithubSelector**](NotificationsAPI.md#AddGithubSelector) | **Post** /notifications/endpoints/github/configurations/{uuid}/selectors | 
[**AddJiraConfiguration**](NotificationsAPI.md#AddJiraConfiguration) | **Post** /notifications/endpoints/jira/configurations | 
[**AddJiraSelector**](NotificationsAPI.md#AddJiraSelector) | **Post** /notifications/endpoints/jira/configurations/{uuid}/selectors | 
[**AddSlackConfiguration**](NotificationsAPI.md#AddSlackConfiguration) | **Post** /notifications/endpoints/slack/configurations | 
[**AddSlackSelector**](NotificationsAPI.md#AddSlackSelector) | **Post** /notifications/endpoints/slack/configurations/{uuid}/selectors | 
[**AddSmtpConfiguration**](NotificationsAPI.md#AddSmtpConfiguration) | **Post** /notifications/endpoints/smtp/configurations | 
[**AddSmtpSelector**](NotificationsAPI.md#AddSmtpSelector) | **Post** /notifications/endpoints/smtp/configurations/{uuid}/selectors | 
[**AddTeamsConfiguration**](NotificationsAPI.md#AddTeamsConfiguration) | **Post** /notifications/endpoints/teams/configurations | 
[**AddTeamsSelector**](NotificationsAPI.md#AddTeamsSelector) | **Post** /notifications/endpoints/teams/configurations/{uuid}/selectors | 
[**AddWebhookConfiguration**](NotificationsAPI.md#AddWebhookConfiguration) | **Post** /notifications/endpoints/webhook/configurations | 
[**AddWebhookSelector**](NotificationsAPI.md#AddWebhookSelector) | **Post** /notifications/endpoints/webhook/configurations/{uuid}/selectors | 
[**DeleteGithubConfiguration**](NotificationsAPI.md#DeleteGithubConfiguration) | **Delete** /notifications/endpoints/github/configurations/{uuid} | 
[**DeleteGithubSelector**](NotificationsAPI.md#DeleteGithubSelector) | **Delete** /notifications/endpoints/github/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**DeleteJiraConfiguration**](NotificationsAPI.md#DeleteJiraConfiguration) | **Delete** /notifications/endpoints/jira/configurations/{uuid} | 
[**DeleteJiraSelector**](NotificationsAPI.md#DeleteJiraSelector) | **Delete** /notifications/endpoints/jira/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**DeleteSlackConfiguration**](NotificationsAPI.md#DeleteSlackConfiguration) | **Delete** /notifications/endpoints/slack/configurations/{uuid} | 
[**DeleteSlackSelector**](NotificationsAPI.md#DeleteSlackSelector) | **Delete** /notifications/endpoints/slack/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**DeleteSmtpConfiguration**](NotificationsAPI.md#DeleteSmtpConfiguration) | **Delete** /notifications/endpoints/smtp/configurations/{uuid} | 
[**DeleteSmtpSelector**](NotificationsAPI.md#DeleteSmtpSelector) | **Delete** /notifications/endpoints/smtp/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**DeleteTeamsConfiguration**](NotificationsAPI.md#DeleteTeamsConfiguration) | **Delete** /notifications/endpoints/teams/configurations/{uuid} | 
[**DeleteTeamsSelector**](NotificationsAPI.md#DeleteTeamsSelector) | **Delete** /notifications/endpoints/teams/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**DeleteWebhookConfiguration**](NotificationsAPI.md#DeleteWebhookConfiguration) | **Delete** /notifications/endpoints/webhook/configurations/{uuid} | 
[**DeleteWebhookSelector**](NotificationsAPI.md#DeleteWebhookSelector) | **Delete** /notifications/endpoints/webhook/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**GetGithubConfiguration**](NotificationsAPI.md#GetGithubConfiguration) | **Get** /notifications/endpoints/github/configurations/{uuid} | 
[**GetGithubConfigurationStatus**](NotificationsAPI.md#GetGithubConfigurationStatus) | **Get** /notifications/endpoints/github/configurations/{uuid}/status | 
[**GetGithubSelector**](NotificationsAPI.md#GetGithubSelector) | **Get** /notifications/endpoints/github/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**GetJiraConfiguration**](NotificationsAPI.md#GetJiraConfiguration) | **Get** /notifications/endpoints/jira/configurations/{uuid} | 
[**GetJiraConfigurationStatus**](NotificationsAPI.md#GetJiraConfigurationStatus) | **Get** /notifications/endpoints/jira/configurations/{uuid}/status | 
[**GetJiraSelector**](NotificationsAPI.md#GetJiraSelector) | **Get** /notifications/endpoints/jira/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**GetSlackConfiguration**](NotificationsAPI.md#GetSlackConfiguration) | **Get** /notifications/endpoints/slack/configurations/{uuid} | 
[**GetSlackConfigurationStatus**](NotificationsAPI.md#GetSlackConfigurationStatus) | **Get** /notifications/endpoints/slack/configurations/{uuid}/status | 
[**GetSlackSelector**](NotificationsAPI.md#GetSlackSelector) | **Get** /notifications/endpoints/slack/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**GetSmtpConfiguration**](NotificationsAPI.md#GetSmtpConfiguration) | **Get** /notifications/endpoints/smtp/configurations/{uuid} | 
[**GetSmtpConfigurationStatus**](NotificationsAPI.md#GetSmtpConfigurationStatus) | **Get** /notifications/endpoints/smtp/configurations/{uuid}/status | 
[**GetSmtpSelector**](NotificationsAPI.md#GetSmtpSelector) | **Get** /notifications/endpoints/smtp/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**GetTeamsConfiguration**](NotificationsAPI.md#GetTeamsConfiguration) | **Get** /notifications/endpoints/teams/configurations/{uuid} | 
[**GetTeamsConfigurationStatus**](NotificationsAPI.md#GetTeamsConfigurationStatus) | **Get** /notifications/endpoints/teams/configurations/{uuid}/status | 
[**GetTeamsSelector**](NotificationsAPI.md#GetTeamsSelector) | **Get** /notifications/endpoints/teams/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**GetWebhookConfiguration**](NotificationsAPI.md#GetWebhookConfiguration) | **Get** /notifications/endpoints/webhook/configurations/{uuid} | 
[**GetWebhookConfigurationStatus**](NotificationsAPI.md#GetWebhookConfigurationStatus) | **Get** /notifications/endpoints/webhook/configurations/{uuid}/status | 
[**GetWebhookSelector**](NotificationsAPI.md#GetWebhookSelector) | **Get** /notifications/endpoints/webhook/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**ListEndpoints**](NotificationsAPI.md#ListEndpoints) | **Get** /notifications/endpoints | 
[**ListGithubConfigurations**](NotificationsAPI.md#ListGithubConfigurations) | **Get** /notifications/endpoints/github/configurations | 
[**ListGithubSelectors**](NotificationsAPI.md#ListGithubSelectors) | **Get** /notifications/endpoints/github/configurations/{uuid}/selectors | 
[**ListJiraConfigurations**](NotificationsAPI.md#ListJiraConfigurations) | **Get** /notifications/endpoints/jira/configurations | 
[**ListJiraSelectors**](NotificationsAPI.md#ListJiraSelectors) | **Get** /notifications/endpoints/jira/configurations/{uuid}/selectors | 
[**ListSelectors**](NotificationsAPI.md#ListSelectors) | **Get** /notifications/selectors | 
[**ListSlackConfigurations**](NotificationsAPI.md#ListSlackConfigurations) | **Get** /notifications/endpoints/slack/configurations | 
[**ListSlackSelectors**](NotificationsAPI.md#ListSlackSelectors) | **Get** /notifications/endpoints/slack/configurations/{uuid}/selectors | 
[**ListSmtpConfigurations**](NotificationsAPI.md#ListSmtpConfigurations) | **Get** /notifications/endpoints/smtp/configurations | 
[**ListSmtpSelectors**](NotificationsAPI.md#ListSmtpSelectors) | **Get** /notifications/endpoints/smtp/configurations/{uuid}/selectors | 
[**ListTeamsConfigurations**](NotificationsAPI.md#ListTeamsConfigurations) | **Get** /notifications/endpoints/teams/configurations | 
[**ListTeamsSelectors**](NotificationsAPI.md#ListTeamsSelectors) | **Get** /notifications/endpoints/teams/configurations/{uuid}/selectors | 
[**ListWebhookConfigurations**](NotificationsAPI.md#ListWebhookConfigurations) | **Get** /notifications/endpoints/webhook/configurations | 
[**ListWebhookSelectors**](NotificationsAPI.md#ListWebhookSelectors) | **Get** /notifications/endpoints/webhook/configurations/{uuid}/selectors | 
[**TestGithubConfiguration**](NotificationsAPI.md#TestGithubConfiguration) | **Post** /notifications/endpoints/github/test | 
[**TestJiraConfiguration**](NotificationsAPI.md#TestJiraConfiguration) | **Post** /notifications/endpoints/jira/test | 
[**TestSlackConfiguration**](NotificationsAPI.md#TestSlackConfiguration) | **Post** /notifications/endpoints/slack/test | 
[**TestSmtpConfiguration**](NotificationsAPI.md#TestSmtpConfiguration) | **Post** /notifications/endpoints/smtp/test | 
[**TestStoredGithubConfiguration**](NotificationsAPI.md#TestStoredGithubConfiguration) | **Get** /notifications/endpoints/github/configurations/{uuid}/test | 
[**TestStoredJiraConfiguration**](NotificationsAPI.md#TestStoredJiraConfiguration) | **Get** /notifications/endpoints/jira/configurations/{uuid}/test | 
[**TestStoredSlackConfiguration**](NotificationsAPI.md#TestStoredSlackConfiguration) | **Get** /notifications/endpoints/slack/configurations/{uuid}/test | 
[**TestStoredSmtpConfiguration**](NotificationsAPI.md#TestStoredSmtpConfiguration) | **Get** /notifications/endpoints/smtp/configurations/{uuid}/test | 
[**TestStoredTeamsConfiguration**](NotificationsAPI.md#TestStoredTeamsConfiguration) | **Get** /notifications/endpoints/teams/configurations/{uuid}/test | 
[**TestStoredWebhookConfiguration**](NotificationsAPI.md#TestStoredWebhookConfiguration) | **Get** /notifications/endpoints/webhook/configurations/{uuid}/test | 
[**TestTeamsConfiguration**](NotificationsAPI.md#TestTeamsConfiguration) | **Post** /notifications/endpoints/teams/test | 
[**TestWebhookConfiguration**](NotificationsAPI.md#TestWebhookConfiguration) | **Post** /notifications/endpoints/webhook/test | 
[**UpdateEndpointStatus**](NotificationsAPI.md#UpdateEndpointStatus) | **Put** /notifications/endpoints/{name} | 
[**UpdateGithubConfiguration**](NotificationsAPI.md#UpdateGithubConfiguration) | **Put** /notifications/endpoints/github/configurations/{uuid} | 
[**UpdateGithubSelector**](NotificationsAPI.md#UpdateGithubSelector) | **Put** /notifications/endpoints/github/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**UpdateJiraConfiguration**](NotificationsAPI.md#UpdateJiraConfiguration) | **Put** /notifications/endpoints/jira/configurations/{uuid} | 
[**UpdateJiraSelector**](NotificationsAPI.md#UpdateJiraSelector) | **Put** /notifications/endpoints/jira/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**UpdateSlackConfiguration**](NotificationsAPI.md#UpdateSlackConfiguration) | **Put** /notifications/endpoints/slack/configurations/{uuid} | 
[**UpdateSlackSelector**](NotificationsAPI.md#UpdateSlackSelector) | **Put** /notifications/endpoints/slack/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**UpdateSmtpConfiguration**](NotificationsAPI.md#UpdateSmtpConfiguration) | **Put** /notifications/endpoints/smtp/configurations/{uuid} | 
[**UpdateSmtpSelector**](NotificationsAPI.md#UpdateSmtpSelector) | **Put** /notifications/endpoints/smtp/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**UpdateTeamsConfiguration**](NotificationsAPI.md#UpdateTeamsConfiguration) | **Put** /notifications/endpoints/teams/configurations/{uuid} | 
[**UpdateTeamsSelector**](NotificationsAPI.md#UpdateTeamsSelector) | **Put** /notifications/endpoints/teams/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**UpdateWebhookConfiguration**](NotificationsAPI.md#UpdateWebhookConfiguration) | **Put** /notifications/endpoints/webhook/configurations/{uuid} | 
[**UpdateWebhookSelector**](NotificationsAPI.md#UpdateWebhookSelector) | **Put** /notifications/endpoints/webhook/configurations/{configuration_uuid}/selectors/{selector_uuid} | 



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
	resp, r, err := apiClient.NotificationsAPI.AddGithubConfiguration(context.Background()).Configuration(configuration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.AddGithubConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddGithubConfiguration`: NotificationGitHubEndpointConfigurationBase
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.AddGithubConfiguration`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.AddGithubSelector(context.Background(), uuid).Selector(selector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.AddGithubSelector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddGithubSelector`: NotificationSelector
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.AddGithubSelector`: %v\n", resp)
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
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	configuration := *openapiclient.NewNotificationJiraEndpointConfigurationPost("Url_example", "Username_example", "Password_example", "ProjectKey_example", "IssueType_example") // NotificationJiraEndpointConfigurationPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.AddJiraConfiguration(context.Background()).Configuration(configuration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.AddJiraConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddJiraConfiguration`: NotificationJiraEndpointConfigurationBase
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.AddJiraConfiguration`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.AddJiraSelector(context.Background(), uuid).Selector(selector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.AddJiraSelector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddJiraSelector`: NotificationSelector
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.AddJiraSelector`: %v\n", resp)
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
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	configuration := *openapiclient.NewNotificationSlackEndpointConfiguration("Url_example") // NotificationSlackEndpointConfiguration | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.AddSlackConfiguration(context.Background()).Configuration(configuration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.AddSlackConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddSlackConfiguration`: NotificationSlackEndpointConfiguration
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.AddSlackConfiguration`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.AddSlackSelector(context.Background(), uuid).Selector(selector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.AddSlackSelector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddSlackSelector`: NotificationSelector
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.AddSlackSelector`: %v\n", resp)
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
	configuration := *openapiclient.NewNotificationSMTPEndpointConfiguration("Host_example", int64(123), "From_example", "To_example") // NotificationSMTPEndpointConfiguration | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.AddSmtpConfiguration(context.Background()).Configuration(configuration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.AddSmtpConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddSmtpConfiguration`: NotificationSMTPEndpointConfiguration
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.AddSmtpConfiguration`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.AddSmtpSelector(context.Background(), uuid).Selector(selector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.AddSmtpSelector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddSmtpSelector`: NotificationSelector
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.AddSmtpSelector`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.AddTeamsConfiguration(context.Background()).Configuration(configuration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.AddTeamsConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddTeamsConfiguration`: NotificationTeamsEndpointConfiguration
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.AddTeamsConfiguration`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.AddTeamsSelector(context.Background(), uuid).Selector(selector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.AddTeamsSelector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddTeamsSelector`: NotificationSelector
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.AddTeamsSelector`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.AddWebhookConfiguration(context.Background()).Configuration(configuration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.AddWebhookConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddWebhookConfiguration`: NotificationWebhookEndpointConfiguration
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.AddWebhookConfiguration`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.AddWebhookSelector(context.Background(), uuid).Selector(selector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.AddWebhookSelector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddWebhookSelector`: NotificationSelector
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.AddWebhookSelector`: %v\n", resp)
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
	r, err := apiClient.NotificationsAPI.DeleteGithubConfiguration(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.DeleteGithubConfiguration``: %v\n", err)
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
	r, err := apiClient.NotificationsAPI.DeleteGithubSelector(context.Background(), configurationUuid, selectorUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.DeleteGithubSelector``: %v\n", err)
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
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	uuid := "uuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotificationsAPI.DeleteJiraConfiguration(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.DeleteJiraConfiguration``: %v\n", err)
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
	r, err := apiClient.NotificationsAPI.DeleteJiraSelector(context.Background(), configurationUuid, selectorUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.DeleteJiraSelector``: %v\n", err)
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
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	uuid := "uuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotificationsAPI.DeleteSlackConfiguration(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.DeleteSlackConfiguration``: %v\n", err)
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
	r, err := apiClient.NotificationsAPI.DeleteSlackSelector(context.Background(), configurationUuid, selectorUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.DeleteSlackSelector``: %v\n", err)
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
	r, err := apiClient.NotificationsAPI.DeleteSmtpConfiguration(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.DeleteSmtpConfiguration``: %v\n", err)
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
	r, err := apiClient.NotificationsAPI.DeleteSmtpSelector(context.Background(), configurationUuid, selectorUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.DeleteSmtpSelector``: %v\n", err)
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
	r, err := apiClient.NotificationsAPI.DeleteTeamsConfiguration(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.DeleteTeamsConfiguration``: %v\n", err)
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
	r, err := apiClient.NotificationsAPI.DeleteTeamsSelector(context.Background(), configurationUuid, selectorUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.DeleteTeamsSelector``: %v\n", err)
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
	r, err := apiClient.NotificationsAPI.DeleteWebhookConfiguration(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.DeleteWebhookConfiguration``: %v\n", err)
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
	r, err := apiClient.NotificationsAPI.DeleteWebhookSelector(context.Background(), configurationUuid, selectorUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.DeleteWebhookSelector``: %v\n", err)
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
	resp, r, err := apiClient.NotificationsAPI.GetGithubConfiguration(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetGithubConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGithubConfiguration`: NotificationGitHubEndpointConfigurationBase
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetGithubConfiguration`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.GetGithubConfigurationStatus(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetGithubConfigurationStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGithubConfigurationStatus`: NotificationOperationalStatus
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetGithubConfigurationStatus`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.GetGithubSelector(context.Background(), configurationUuid, selectorUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetGithubSelector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGithubSelector`: NotificationSelector
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetGithubSelector`: %v\n", resp)
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
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	uuid := "uuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.GetJiraConfiguration(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetJiraConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJiraConfiguration`: NotificationJiraEndpointConfigurationBase
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetJiraConfiguration`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.GetJiraConfigurationStatus(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetJiraConfigurationStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJiraConfigurationStatus`: NotificationOperationalStatus
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetJiraConfigurationStatus`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.GetJiraSelector(context.Background(), configurationUuid, selectorUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetJiraSelector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJiraSelector`: NotificationSelector
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetJiraSelector`: %v\n", resp)
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
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	uuid := "uuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.GetSlackConfiguration(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetSlackConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSlackConfiguration`: NotificationSlackEndpointConfiguration
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetSlackConfiguration`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.GetSlackConfigurationStatus(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetSlackConfigurationStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSlackConfigurationStatus`: NotificationOperationalStatus
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetSlackConfigurationStatus`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.GetSlackSelector(context.Background(), configurationUuid, selectorUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetSlackSelector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSlackSelector`: NotificationSelector
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetSlackSelector`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.GetSmtpConfiguration(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetSmtpConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSmtpConfiguration`: NotificationSMTPEndpointConfiguration
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetSmtpConfiguration`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.GetSmtpConfigurationStatus(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetSmtpConfigurationStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSmtpConfigurationStatus`: NotificationOperationalStatus
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetSmtpConfigurationStatus`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.GetSmtpSelector(context.Background(), configurationUuid, selectorUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetSmtpSelector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSmtpSelector`: NotificationSelector
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetSmtpSelector`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.GetTeamsConfiguration(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetTeamsConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamsConfiguration`: NotificationTeamsEndpointConfiguration
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetTeamsConfiguration`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.GetTeamsConfigurationStatus(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetTeamsConfigurationStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamsConfigurationStatus`: NotificationOperationalStatus
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetTeamsConfigurationStatus`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.GetTeamsSelector(context.Background(), configurationUuid, selectorUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetTeamsSelector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamsSelector`: NotificationSelector
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetTeamsSelector`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.GetWebhookConfiguration(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetWebhookConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookConfiguration`: NotificationWebhookEndpointConfiguration
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetWebhookConfiguration`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.GetWebhookConfigurationStatus(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetWebhookConfigurationStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookConfigurationStatus`: NotificationOperationalStatus
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetWebhookConfigurationStatus`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.GetWebhookSelector(context.Background(), configurationUuid, selectorUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetWebhookSelector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookSelector`: NotificationSelector
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetWebhookSelector`: %v\n", resp)
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
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.ListEndpoints(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.ListEndpoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEndpoints`: []NotificationEndpoint
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.ListEndpoints`: %v\n", resp)
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
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.ListGithubConfigurations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.ListGithubConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGithubConfigurations`: []NotificationGitHubEndpointConfigurationBase
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.ListGithubConfigurations`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.ListGithubSelectors(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.ListGithubSelectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGithubSelectors`: []NotificationSelector
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.ListGithubSelectors`: %v\n", resp)
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
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.ListJiraConfigurations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.ListJiraConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListJiraConfigurations`: []NotificationJiraEndpointConfigurationBase
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.ListJiraConfigurations`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.ListJiraSelectors(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.ListJiraSelectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListJiraSelectors`: []NotificationSelector
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.ListJiraSelectors`: %v\n", resp)
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
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.ListSelectors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.ListSelectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSelectors`: []NotificationSelector
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.ListSelectors`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.ListSlackConfigurations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.ListSlackConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSlackConfigurations`: []NotificationSlackEndpointConfiguration
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.ListSlackConfigurations`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.ListSlackSelectors(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.ListSlackSelectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSlackSelectors`: []NotificationSelector
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.ListSlackSelectors`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.ListSmtpConfigurations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.ListSmtpConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSmtpConfigurations`: []NotificationSMTPEndpointConfiguration
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.ListSmtpConfigurations`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.ListSmtpSelectors(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.ListSmtpSelectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSmtpSelectors`: []NotificationSelector
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.ListSmtpSelectors`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.ListTeamsConfigurations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.ListTeamsConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTeamsConfigurations`: []NotificationTeamsEndpointConfiguration
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.ListTeamsConfigurations`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.ListTeamsSelectors(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.ListTeamsSelectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTeamsSelectors`: []NotificationSelector
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.ListTeamsSelectors`: %v\n", resp)
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
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.ListWebhookConfigurations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.ListWebhookConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWebhookConfigurations`: []NotificationWebhookEndpointConfiguration
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.ListWebhookConfigurations`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.ListWebhookSelectors(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.ListWebhookSelectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWebhookSelectors`: []NotificationSelector
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.ListWebhookSelectors`: %v\n", resp)
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
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	configuration := *openapiclient.NewNotificationGitHubEndpointConfigurationPost("Username_example", "AccessToken_example", "Owner_example", "Repository_example") // NotificationGitHubEndpointConfigurationPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.TestGithubConfiguration(context.Background()).Configuration(configuration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.TestGithubConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestGithubConfiguration`: NotificationGitHubTestResult
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.TestGithubConfiguration`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.TestJiraConfiguration(context.Background()).Configuration(configuration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.TestJiraConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestJiraConfiguration`: NotificationJiraTestResult
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.TestJiraConfiguration`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.TestSlackConfiguration(context.Background()).Configuration(configuration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.TestSlackConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestSlackConfiguration`: NotificationSlackTestResult
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.TestSlackConfiguration`: %v\n", resp)
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
	configuration := *openapiclient.NewNotificationSMTPEndpointConfiguration("Host_example", int64(123), "From_example", "To_example") // NotificationSMTPEndpointConfiguration | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.TestSmtpConfiguration(context.Background()).Configuration(configuration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.TestSmtpConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestSmtpConfiguration`: NotificationSMTPTestResult
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.TestSmtpConfiguration`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.TestStoredGithubConfiguration(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.TestStoredGithubConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestStoredGithubConfiguration`: NotificationGitHubTestResult
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.TestStoredGithubConfiguration`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.TestStoredJiraConfiguration(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.TestStoredJiraConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestStoredJiraConfiguration`: NotificationJiraTestResult
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.TestStoredJiraConfiguration`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.TestStoredSlackConfiguration(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.TestStoredSlackConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestStoredSlackConfiguration`: NotificationSlackTestResult
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.TestStoredSlackConfiguration`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.TestStoredSmtpConfiguration(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.TestStoredSmtpConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestStoredSmtpConfiguration`: NotificationSMTPTestResult
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.TestStoredSmtpConfiguration`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.TestStoredTeamsConfiguration(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.TestStoredTeamsConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestStoredTeamsConfiguration`: NotificationTeamsTestResult
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.TestStoredTeamsConfiguration`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.TestStoredWebhookConfiguration(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.TestStoredWebhookConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestStoredWebhookConfiguration`: NotificationWebhookTestResult
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.TestStoredWebhookConfiguration`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.TestTeamsConfiguration(context.Background()).Configuration(configuration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.TestTeamsConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestTeamsConfiguration`: NotificationTeamsTestResult
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.TestTeamsConfiguration`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.TestWebhookConfiguration(context.Background()).Configuration(configuration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.TestWebhookConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestWebhookConfiguration`: NotificationWebhookTestResult
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.TestWebhookConfiguration`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.UpdateEndpointStatus(context.Background(), name).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.UpdateEndpointStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEndpointStatus`: NotificationEndpointEnabledStatus
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.UpdateEndpointStatus`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.UpdateGithubConfiguration(context.Background(), uuid).Configuration(configuration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.UpdateGithubConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGithubConfiguration`: NotificationGitHubEndpointConfigurationBase
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.UpdateGithubConfiguration`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.UpdateGithubSelector(context.Background(), configurationUuid, selectorUuid).Selector(selector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.UpdateGithubSelector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGithubSelector`: NotificationSelector
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.UpdateGithubSelector`: %v\n", resp)
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
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	uuid := "uuid_example" // string | 
	configuration := *openapiclient.NewNotificationJiraEndpointConfigurationPut("Url_example", "Username_example", "ProjectKey_example", "IssueType_example") // NotificationJiraEndpointConfigurationPut | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.UpdateJiraConfiguration(context.Background(), uuid).Configuration(configuration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.UpdateJiraConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateJiraConfiguration`: NotificationJiraEndpointConfigurationBase
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.UpdateJiraConfiguration`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.UpdateJiraSelector(context.Background(), configurationUuid, selectorUuid).Selector(selector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.UpdateJiraSelector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateJiraSelector`: NotificationSelector
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.UpdateJiraSelector`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.UpdateSlackConfiguration(context.Background(), uuid).Configuration(configuration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.UpdateSlackConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSlackConfiguration`: NotificationSlackEndpointConfiguration
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.UpdateSlackConfiguration`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.UpdateSlackSelector(context.Background(), configurationUuid, selectorUuid).Selector(selector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.UpdateSlackSelector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSlackSelector`: NotificationSelector
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.UpdateSlackSelector`: %v\n", resp)
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
	configuration := *openapiclient.NewNotificationSMTPEndpointConfiguration("Host_example", int64(123), "From_example", "To_example") // NotificationSMTPEndpointConfiguration | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.UpdateSmtpConfiguration(context.Background(), uuid).Configuration(configuration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.UpdateSmtpConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSmtpConfiguration`: NotificationSMTPEndpointConfiguration
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.UpdateSmtpConfiguration`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.UpdateSmtpSelector(context.Background(), configurationUuid, selectorUuid).Selector(selector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.UpdateSmtpSelector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSmtpSelector`: NotificationSelector
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.UpdateSmtpSelector`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.UpdateTeamsConfiguration(context.Background(), uuid).Configuration(configuration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.UpdateTeamsConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTeamsConfiguration`: NotificationTeamsEndpointConfiguration
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.UpdateTeamsConfiguration`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.UpdateTeamsSelector(context.Background(), configurationUuid, selectorUuid).Selector(selector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.UpdateTeamsSelector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTeamsSelector`: NotificationSelector
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.UpdateTeamsSelector`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.UpdateWebhookConfiguration(context.Background(), uuid).Configuration(configuration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.UpdateWebhookConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWebhookConfiguration`: NotificationWebhookEndpointConfiguration
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.UpdateWebhookConfiguration`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationsAPI.UpdateWebhookSelector(context.Background(), configurationUuid, selectorUuid).Selector(selector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.UpdateWebhookSelector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWebhookSelector`: NotificationSelector
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.UpdateWebhookSelector`: %v\n", resp)
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

