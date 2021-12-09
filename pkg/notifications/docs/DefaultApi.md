# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGithubConfiguration**](DefaultApi.md#AddGithubConfiguration) | **Post** /endpoints/github/configurations | 
[**AddGithubSelector**](DefaultApi.md#AddGithubSelector) | **Post** /endpoints/github/configurations/{uuid}/selectors | 
[**AddJiraConfiguration**](DefaultApi.md#AddJiraConfiguration) | **Post** /endpoints/jira/configurations | 
[**AddJiraSelector**](DefaultApi.md#AddJiraSelector) | **Post** /endpoints/jira/configurations/{uuid}/selectors | 
[**AddSlackConfiguration**](DefaultApi.md#AddSlackConfiguration) | **Post** /endpoints/slack/configurations | 
[**AddSlackSelector**](DefaultApi.md#AddSlackSelector) | **Post** /endpoints/slack/configurations/{uuid}/selectors | 
[**AddSmtpConfiguration**](DefaultApi.md#AddSmtpConfiguration) | **Post** /endpoints/smtp/configurations | 
[**AddSmtpSelector**](DefaultApi.md#AddSmtpSelector) | **Post** /endpoints/smtp/configurations/{uuid}/selectors | 
[**AddTeamsConfiguration**](DefaultApi.md#AddTeamsConfiguration) | **Post** /endpoints/teams/configurations | 
[**AddTeamsSelector**](DefaultApi.md#AddTeamsSelector) | **Post** /endpoints/teams/configurations/{uuid}/selectors | 
[**AddWebhookConfiguration**](DefaultApi.md#AddWebhookConfiguration) | **Post** /endpoints/webhook/configurations | 
[**AddWebhookSelector**](DefaultApi.md#AddWebhookSelector) | **Post** /endpoints/webhook/configurations/{uuid}/selectors | 
[**DeleteGithubConfiguration**](DefaultApi.md#DeleteGithubConfiguration) | **Delete** /endpoints/github/configurations/{uuid} | 
[**DeleteGithubSelector**](DefaultApi.md#DeleteGithubSelector) | **Delete** /endpoints/github/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**DeleteJiraConfiguration**](DefaultApi.md#DeleteJiraConfiguration) | **Delete** /endpoints/jira/configurations/{uuid} | 
[**DeleteJiraSelector**](DefaultApi.md#DeleteJiraSelector) | **Delete** /endpoints/jira/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**DeleteSlackConfiguration**](DefaultApi.md#DeleteSlackConfiguration) | **Delete** /endpoints/slack/configurations/{uuid} | 
[**DeleteSlackSelector**](DefaultApi.md#DeleteSlackSelector) | **Delete** /endpoints/slack/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**DeleteSmtpConfiguration**](DefaultApi.md#DeleteSmtpConfiguration) | **Delete** /endpoints/smtp/configurations/{uuid} | 
[**DeleteSmtpSelector**](DefaultApi.md#DeleteSmtpSelector) | **Delete** /endpoints/smtp/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**DeleteTeamsConfiguration**](DefaultApi.md#DeleteTeamsConfiguration) | **Delete** /endpoints/teams/configurations/{uuid} | 
[**DeleteTeamsSelector**](DefaultApi.md#DeleteTeamsSelector) | **Delete** /endpoints/teams/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**DeleteWebhookConfiguration**](DefaultApi.md#DeleteWebhookConfiguration) | **Delete** /endpoints/webhook/configurations/{uuid} | 
[**DeleteWebhookSelector**](DefaultApi.md#DeleteWebhookSelector) | **Delete** /endpoints/webhook/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**GetGithubConfiguration**](DefaultApi.md#GetGithubConfiguration) | **Get** /endpoints/github/configurations/{uuid} | 
[**GetGithubConfigurationStatus**](DefaultApi.md#GetGithubConfigurationStatus) | **Get** /endpoints/github/configurations/{uuid}/status | 
[**GetGithubSelector**](DefaultApi.md#GetGithubSelector) | **Get** /endpoints/github/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**GetJiraConfiguration**](DefaultApi.md#GetJiraConfiguration) | **Get** /endpoints/jira/configurations/{uuid} | 
[**GetJiraConfigurationStatus**](DefaultApi.md#GetJiraConfigurationStatus) | **Get** /endpoints/jira/configurations/{uuid}/status | 
[**GetJiraSelector**](DefaultApi.md#GetJiraSelector) | **Get** /endpoints/jira/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**GetSlackConfiguration**](DefaultApi.md#GetSlackConfiguration) | **Get** /endpoints/slack/configurations/{uuid} | 
[**GetSlackConfigurationStatus**](DefaultApi.md#GetSlackConfigurationStatus) | **Get** /endpoints/slack/configurations/{uuid}/status | 
[**GetSlackSelector**](DefaultApi.md#GetSlackSelector) | **Get** /endpoints/slack/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**GetSmtpConfiguration**](DefaultApi.md#GetSmtpConfiguration) | **Get** /endpoints/smtp/configurations/{uuid} | 
[**GetSmtpConfigurationStatus**](DefaultApi.md#GetSmtpConfigurationStatus) | **Get** /endpoints/smtp/configurations/{uuid}/status | 
[**GetSmtpSelector**](DefaultApi.md#GetSmtpSelector) | **Get** /endpoints/smtp/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**GetStatus**](DefaultApi.md#GetStatus) | **Get** /status | Service status
[**GetTeamsConfiguration**](DefaultApi.md#GetTeamsConfiguration) | **Get** /endpoints/teams/configurations/{uuid} | 
[**GetTeamsConfigurationStatus**](DefaultApi.md#GetTeamsConfigurationStatus) | **Get** /endpoints/teams/configurations/{uuid}/status | 
[**GetTeamsSelector**](DefaultApi.md#GetTeamsSelector) | **Get** /endpoints/teams/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**GetWebhookConfiguration**](DefaultApi.md#GetWebhookConfiguration) | **Get** /endpoints/webhook/configurations/{uuid} | 
[**GetWebhookConfigurationStatus**](DefaultApi.md#GetWebhookConfigurationStatus) | **Get** /endpoints/webhook/configurations/{uuid}/status | 
[**GetWebhookSelector**](DefaultApi.md#GetWebhookSelector) | **Get** /endpoints/webhook/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**HealthCheck**](DefaultApi.md#HealthCheck) | **Get** /health | 
[**ListEndpoints**](DefaultApi.md#ListEndpoints) | **Get** /endpoints | 
[**ListGithubConfigurations**](DefaultApi.md#ListGithubConfigurations) | **Get** /endpoints/github/configurations | 
[**ListGithubSelectors**](DefaultApi.md#ListGithubSelectors) | **Get** /endpoints/github/configurations/{uuid}/selectors | 
[**ListJiraConfigurations**](DefaultApi.md#ListJiraConfigurations) | **Get** /endpoints/jira/configurations | 
[**ListJiraSelectors**](DefaultApi.md#ListJiraSelectors) | **Get** /endpoints/jira/configurations/{uuid}/selectors | 
[**ListSelectors**](DefaultApi.md#ListSelectors) | **Get** /selectors | 
[**ListSlackConfigurations**](DefaultApi.md#ListSlackConfigurations) | **Get** /endpoints/slack/configurations | 
[**ListSlackSelectors**](DefaultApi.md#ListSlackSelectors) | **Get** /endpoints/slack/configurations/{uuid}/selectors | 
[**ListSmtpConfigurations**](DefaultApi.md#ListSmtpConfigurations) | **Get** /endpoints/smtp/configurations | 
[**ListSmtpSelectors**](DefaultApi.md#ListSmtpSelectors) | **Get** /endpoints/smtp/configurations/{uuid}/selectors | 
[**ListTeamsConfigurations**](DefaultApi.md#ListTeamsConfigurations) | **Get** /endpoints/teams/configurations | 
[**ListTeamsSelectors**](DefaultApi.md#ListTeamsSelectors) | **Get** /endpoints/teams/configurations/{uuid}/selectors | 
[**ListWebhookConfigurations**](DefaultApi.md#ListWebhookConfigurations) | **Get** /endpoints/webhook/configurations | 
[**ListWebhookSelectors**](DefaultApi.md#ListWebhookSelectors) | **Get** /endpoints/webhook/configurations/{uuid}/selectors | 
[**Notify**](DefaultApi.md#Notify) | **Post** /internal/endpoints/{name}/configurations/{uuid}/notify | 
[**TestGithubConfiguration**](DefaultApi.md#TestGithubConfiguration) | **Post** /endpoints/github/test | 
[**TestJiraConfiguration**](DefaultApi.md#TestJiraConfiguration) | **Post** /endpoints/jira/test | 
[**TestSlackConfiguration**](DefaultApi.md#TestSlackConfiguration) | **Post** /endpoints/slack/test | 
[**TestSmtpConfiguration**](DefaultApi.md#TestSmtpConfiguration) | **Post** /endpoints/smtp/test | 
[**TestTeamsConfiguration**](DefaultApi.md#TestTeamsConfiguration) | **Post** /endpoints/teams/test | 
[**TestWebhookConfiguration**](DefaultApi.md#TestWebhookConfiguration) | **Post** /endpoints/webhook/test | 
[**UpdateEndpointStatus**](DefaultApi.md#UpdateEndpointStatus) | **Put** /endpoints/{name} | 
[**UpdateGithubConfiguration**](DefaultApi.md#UpdateGithubConfiguration) | **Put** /endpoints/github/configurations/{uuid} | 
[**UpdateGithubSelector**](DefaultApi.md#UpdateGithubSelector) | **Put** /endpoints/github/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**UpdateJiraConfiguration**](DefaultApi.md#UpdateJiraConfiguration) | **Put** /endpoints/jira/configurations/{uuid} | 
[**UpdateJiraSelector**](DefaultApi.md#UpdateJiraSelector) | **Put** /endpoints/jira/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**UpdateSlackConfiguration**](DefaultApi.md#UpdateSlackConfiguration) | **Put** /endpoints/slack/configurations/{uuid} | 
[**UpdateSlackSelector**](DefaultApi.md#UpdateSlackSelector) | **Put** /endpoints/slack/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**UpdateSmtpConfiguration**](DefaultApi.md#UpdateSmtpConfiguration) | **Put** /endpoints/smtp/configurations/{uuid} | 
[**UpdateSmtpSelector**](DefaultApi.md#UpdateSmtpSelector) | **Put** /endpoints/smtp/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**UpdateTeamsConfiguration**](DefaultApi.md#UpdateTeamsConfiguration) | **Put** /endpoints/teams/configurations/{uuid} | 
[**UpdateTeamsSelector**](DefaultApi.md#UpdateTeamsSelector) | **Put** /endpoints/teams/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**UpdateWebhookConfiguration**](DefaultApi.md#UpdateWebhookConfiguration) | **Put** /endpoints/webhook/configurations/{uuid} | 
[**UpdateWebhookSelector**](DefaultApi.md#UpdateWebhookSelector) | **Put** /endpoints/webhook/configurations/{configuration_uuid}/selectors/{selector_uuid} | 
[**VersionCheck**](DefaultApi.md#VersionCheck) | **Get** /version | 



## AddGithubConfiguration

> GitHubEndpointConfiguration AddGithubConfiguration(ctx, configuration)



Create a new GitHub endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configuration** | [**GitHubEndpointConfiguration**](GitHubEndpointConfiguration.md)|  | 

### Return type

[**GitHubEndpointConfiguration**](GitHubEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddGithubSelector

> Selector AddGithubSelector(ctx, uuid, selector)



Add selector for mapping events for delievery to a GitHub endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 
**selector** | [**Selector**](Selector.md)|  | 

### Return type

[**Selector**](Selector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddJiraConfiguration

> JiraEndpointConfiguration AddJiraConfiguration(ctx, configuration)



Create a new Jira endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configuration** | [**JiraEndpointConfiguration**](JiraEndpointConfiguration.md)|  | 

### Return type

[**JiraEndpointConfiguration**](JiraEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddJiraSelector

> Selector AddJiraSelector(ctx, uuid, selector)



Add selector for mapping events for delievery to a Jira endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 
**selector** | [**Selector**](Selector.md)|  | 

### Return type

[**Selector**](Selector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddSlackConfiguration

> SlackEndpointConfiguration AddSlackConfiguration(ctx, configuration)



Create a new Slack endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configuration** | [**SlackEndpointConfiguration**](SlackEndpointConfiguration.md)|  | 

### Return type

[**SlackEndpointConfiguration**](SlackEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddSlackSelector

> Selector AddSlackSelector(ctx, uuid, selector)



Add selector for mapping events for delievery to a Slack endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 
**selector** | [**Selector**](Selector.md)|  | 

### Return type

[**Selector**](Selector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddSmtpConfiguration

> SmtpEndpointConfiguration AddSmtpConfiguration(ctx, configuration)



Create a new SMTP endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configuration** | [**SmtpEndpointConfiguration**](SmtpEndpointConfiguration.md)|  | 

### Return type

[**SmtpEndpointConfiguration**](SMTPEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddSmtpSelector

> Selector AddSmtpSelector(ctx, uuid, selector)



Add selector for mapping events for delievery to a SMTP endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 
**selector** | [**Selector**](Selector.md)|  | 

### Return type

[**Selector**](Selector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddTeamsConfiguration

> TeamsEndpointConfiguration AddTeamsConfiguration(ctx, configuration)



Create a new Teams endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configuration** | [**TeamsEndpointConfiguration**](TeamsEndpointConfiguration.md)|  | 

### Return type

[**TeamsEndpointConfiguration**](TeamsEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddTeamsSelector

> Selector AddTeamsSelector(ctx, uuid, selector)



Add selector for mapping events for delievery to a Teams endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 
**selector** | [**Selector**](Selector.md)|  | 

### Return type

[**Selector**](Selector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddWebhookConfiguration

> WebhookEndpointConfiguration AddWebhookConfiguration(ctx, configuration)



Create a new Webhook endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configuration** | [**WebhookEndpointConfiguration**](WebhookEndpointConfiguration.md)|  | 

### Return type

[**WebhookEndpointConfiguration**](WebhookEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddWebhookSelector

> Selector AddWebhookSelector(ctx, uuid, selector)



Add selector for mapping events for delievery to a Webhook endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 
**selector** | [**Selector**](Selector.md)|  | 

### Return type

[**Selector**](Selector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGithubConfiguration

> DeleteGithubConfiguration(ctx, uuid)



Delete a GitHub endpoint configuration by it's UUID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 

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

> DeleteGithubSelector(ctx, configurationUuid, selectorUuid)



Delete a selector mapped to a GitHub endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationUuid** | **string**|  | 
**selectorUuid** | **string**|  | 

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

> DeleteJiraConfiguration(ctx, uuid)



Delete a Jira endpoint configuration by it's UUID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 

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

> DeleteJiraSelector(ctx, configurationUuid, selectorUuid)



Delete a selector mapped to a Jira endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationUuid** | **string**|  | 
**selectorUuid** | **string**|  | 

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

> DeleteSlackConfiguration(ctx, uuid)



Delete a Slack endpoint configuration by it's UUID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 

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

> DeleteSlackSelector(ctx, configurationUuid, selectorUuid)



Delete a selector mapped to a Slack endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationUuid** | **string**|  | 
**selectorUuid** | **string**|  | 

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

> DeleteSmtpConfiguration(ctx, uuid)



Delete a SMTP endpoint configuration by it's UUID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 

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

> DeleteSmtpSelector(ctx, configurationUuid, selectorUuid)



Delete a selector mapped to a SMTP endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationUuid** | **string**|  | 
**selectorUuid** | **string**|  | 

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

> DeleteTeamsConfiguration(ctx, uuid)



Delete a Teams endpoint configuration by it's UUID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 

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

> DeleteTeamsSelector(ctx, configurationUuid, selectorUuid)



Delete a selector mapped to a Teams endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationUuid** | **string**|  | 
**selectorUuid** | **string**|  | 

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

> DeleteWebhookConfiguration(ctx, uuid)



Delete a Webhook endpoint configuration by it's UUID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 

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

> DeleteWebhookSelector(ctx, configurationUuid, selectorUuid)



Delete a selector mapped to a Webhook endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationUuid** | **string**|  | 
**selectorUuid** | **string**|  | 

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

> GitHubEndpointConfiguration GetGithubConfiguration(ctx, uuid)



Get a GitHub endpoint configuration by it's UUID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 

### Return type

[**GitHubEndpointConfiguration**](GitHubEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGithubConfigurationStatus

> OperationalStatus GetGithubConfigurationStatus(ctx, uuid)



Get operational status for a GitHub endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 

### Return type

[**OperationalStatus**](OperationalStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGithubSelector

> Selector GetGithubSelector(ctx, configurationUuid, selectorUuid)



Get a selector mapped to a GitHub endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationUuid** | **string**|  | 
**selectorUuid** | **string**|  | 

### Return type

[**Selector**](Selector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJiraConfiguration

> JiraEndpointConfiguration GetJiraConfiguration(ctx, uuid)



Get a Jira endpoint configuration by it's UUID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 

### Return type

[**JiraEndpointConfiguration**](JiraEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJiraConfigurationStatus

> OperationalStatus GetJiraConfigurationStatus(ctx, uuid)



Get operational status for a Jira endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 

### Return type

[**OperationalStatus**](OperationalStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJiraSelector

> Selector GetJiraSelector(ctx, configurationUuid, selectorUuid)



Get a selector mapped to a Jira endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationUuid** | **string**|  | 
**selectorUuid** | **string**|  | 

### Return type

[**Selector**](Selector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSlackConfiguration

> SlackEndpointConfiguration GetSlackConfiguration(ctx, uuid)



Get a Slack endpoint configuration by it's UUID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 

### Return type

[**SlackEndpointConfiguration**](SlackEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSlackConfigurationStatus

> OperationalStatus GetSlackConfigurationStatus(ctx, uuid)



Get operational status for a Slack endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 

### Return type

[**OperationalStatus**](OperationalStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSlackSelector

> Selector GetSlackSelector(ctx, configurationUuid, selectorUuid)



Get a selector mapped to a Slack endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationUuid** | **string**|  | 
**selectorUuid** | **string**|  | 

### Return type

[**Selector**](Selector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSmtpConfiguration

> SmtpEndpointConfiguration GetSmtpConfiguration(ctx, uuid)



Get a SMTP endpoint configuration by it's UUID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 

### Return type

[**SmtpEndpointConfiguration**](SMTPEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSmtpConfigurationStatus

> OperationalStatus GetSmtpConfigurationStatus(ctx, uuid)



Get operational status for a SMTP endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 

### Return type

[**OperationalStatus**](OperationalStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSmtpSelector

> Selector GetSmtpSelector(ctx, configurationUuid, selectorUuid)



Get a selector mapped to a SMTP endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationUuid** | **string**|  | 
**selectorUuid** | **string**|  | 

### Return type

[**Selector**](Selector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatus

> StatusResponse GetStatus(ctx, )

Service status

Get the API service status

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamsConfiguration

> TeamsEndpointConfiguration GetTeamsConfiguration(ctx, uuid)



Get a Teams endpoint configuration by it's UUID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 

### Return type

[**TeamsEndpointConfiguration**](TeamsEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamsConfigurationStatus

> OperationalStatus GetTeamsConfigurationStatus(ctx, uuid)



Get operational status for a Teams endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 

### Return type

[**OperationalStatus**](OperationalStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamsSelector

> Selector GetTeamsSelector(ctx, configurationUuid, selectorUuid)



Get a selector mapped to a Teams endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationUuid** | **string**|  | 
**selectorUuid** | **string**|  | 

### Return type

[**Selector**](Selector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookConfiguration

> WebhookEndpointConfiguration GetWebhookConfiguration(ctx, uuid)



Get a Webhook endpoint configuration by it's UUID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 

### Return type

[**WebhookEndpointConfiguration**](WebhookEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookConfigurationStatus

> OperationalStatus GetWebhookConfigurationStatus(ctx, uuid)



Get operational status for a Webhook endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 

### Return type

[**OperationalStatus**](OperationalStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookSelector

> Selector GetWebhookSelector(ctx, configurationUuid, selectorUuid)



Get a selector mapped to a Webhook endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationUuid** | **string**|  | 
**selectorUuid** | **string**|  | 

### Return type

[**Selector**](Selector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HealthCheck

> HealthCheck(ctx, )



Health check, returns 200 and no body if service is running

### Required Parameters

This endpoint does not need any parameter.

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

> []Endpoint ListEndpoints(ctx, )



List the system installed notification endpoints

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Endpoint**](Endpoint.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGithubConfigurations

> []GitHubEndpointConfiguration ListGithubConfigurations(ctx, )



List GitHub endpoint configurations

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]GitHubEndpointConfiguration**](GitHubEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGithubSelectors

> []Selector ListGithubSelectors(ctx, uuid)



List selectors mapping events for delivery to a GitHub endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 

### Return type

[**[]Selector**](Selector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListJiraConfigurations

> []JiraEndpointConfiguration ListJiraConfigurations(ctx, )



List Jira endpoint configurations

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]JiraEndpointConfiguration**](JiraEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListJiraSelectors

> []Selector ListJiraSelectors(ctx, uuid)



List selectors mapping events for delivery to a Jira endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 

### Return type

[**[]Selector**](Selector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSelectors

> []Selector ListSelectors(ctx, )



List all selectors mapped to endpoint configurations for the account

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Selector**](Selector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSlackConfigurations

> []SlackEndpointConfiguration ListSlackConfigurations(ctx, )



List Slack endpoint configurations

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]SlackEndpointConfiguration**](SlackEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSlackSelectors

> []Selector ListSlackSelectors(ctx, uuid)



List selectors mapping events for delivery to a Slack endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 

### Return type

[**[]Selector**](Selector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSmtpConfigurations

> []SmtpEndpointConfiguration ListSmtpConfigurations(ctx, )



List SMTP endpoint configurations

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]SmtpEndpointConfiguration**](SMTPEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSmtpSelectors

> []Selector ListSmtpSelectors(ctx, uuid)



List selectors mapping events for delivery to a SMTP endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 

### Return type

[**[]Selector**](Selector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTeamsConfigurations

> []TeamsEndpointConfiguration ListTeamsConfigurations(ctx, )



List Teams endpoint configurations

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]TeamsEndpointConfiguration**](TeamsEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTeamsSelectors

> []Selector ListTeamsSelectors(ctx, uuid)



List selectors mapping events for delivery to a Teams endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 

### Return type

[**[]Selector**](Selector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebhookConfigurations

> []WebhookEndpointConfiguration ListWebhookConfigurations(ctx, )



List Webhook endpoint configurations

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]WebhookEndpointConfiguration**](WebhookEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebhookSelectors

> []Selector ListWebhookSelectors(ctx, uuid)



List selectors mapping events for delivery to a Webhook endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 

### Return type

[**[]Selector**](Selector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Notify

> DispatchResult Notify(ctx, name, uuid, payload)



notify the configuration for the specified endpoint with the provided payload. Currently only internal services can call this endpoint

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**|  | 
**uuid** | **string**|  | 
**payload** | [**SynchronousNotificationPayload**](SynchronousNotificationPayload.md)|  | 

### Return type

[**DispatchResult**](DispatchResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestGithubConfiguration

> GitHubTestResult TestGithubConfiguration(ctx, configuration)



Test GitHub endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configuration** | [**GitHubEndpointConfiguration**](GitHubEndpointConfiguration.md)|  | 

### Return type

[**GitHubTestResult**](GitHubTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestJiraConfiguration

> JiraTestResult TestJiraConfiguration(ctx, configuration)



Test Jira endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configuration** | [**JiraEndpointConfiguration**](JiraEndpointConfiguration.md)|  | 

### Return type

[**JiraTestResult**](JiraTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestSlackConfiguration

> SlackTestResult TestSlackConfiguration(ctx, configuration)



Test Slack endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configuration** | [**SlackEndpointConfiguration**](SlackEndpointConfiguration.md)|  | 

### Return type

[**SlackTestResult**](SlackTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestSmtpConfiguration

> SmtpTestResult TestSmtpConfiguration(ctx, configuration)



Test SMTP endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configuration** | [**SmtpEndpointConfiguration**](SmtpEndpointConfiguration.md)|  | 

### Return type

[**SmtpTestResult**](SMTPTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestTeamsConfiguration

> TeamsTestResult TestTeamsConfiguration(ctx, configuration)



Test Teams endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configuration** | [**TeamsEndpointConfiguration**](TeamsEndpointConfiguration.md)|  | 

### Return type

[**TeamsTestResult**](TeamsTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestWebhookConfiguration

> WebhookTestResult TestWebhookConfiguration(ctx, configuration)



Test Webhook endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configuration** | [**WebhookEndpointConfiguration**](WebhookEndpointConfiguration.md)|  | 

### Return type

[**WebhookTestResult**](WebhookTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEndpointStatus

> EndpointEnabledStatus UpdateEndpointStatus(ctx, name, status)



Update enabled status of an endpoint

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**|  | 
**status** | [**EndpointEnabledStatus**](EndpointEnabledStatus.md)|  | 

### Return type

[**EndpointEnabledStatus**](EndpointEnabledStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGithubConfiguration

> GitHubEndpointConfiguration UpdateGithubConfiguration(ctx, uuid, configuration)



Update a GitHub endpoint configuration by it's UUID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 
**configuration** | [**GitHubEndpointConfiguration**](GitHubEndpointConfiguration.md)|  | 

### Return type

[**GitHubEndpointConfiguration**](GitHubEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGithubSelector

> Selector UpdateGithubSelector(ctx, configurationUuid, selectorUuid, selector)



Update a selector mapped to a GitHub endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationUuid** | **string**|  | 
**selectorUuid** | **string**|  | 
**selector** | [**Selector**](Selector.md)|  | 

### Return type

[**Selector**](Selector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateJiraConfiguration

> JiraEndpointConfiguration UpdateJiraConfiguration(ctx, uuid, configuration)



Update a Jira endpoint configuration by it's UUID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 
**configuration** | [**JiraEndpointConfiguration**](JiraEndpointConfiguration.md)|  | 

### Return type

[**JiraEndpointConfiguration**](JiraEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateJiraSelector

> Selector UpdateJiraSelector(ctx, configurationUuid, selectorUuid, selector)



Update a selector mapped to a Jira endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationUuid** | **string**|  | 
**selectorUuid** | **string**|  | 
**selector** | [**Selector**](Selector.md)|  | 

### Return type

[**Selector**](Selector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSlackConfiguration

> SlackEndpointConfiguration UpdateSlackConfiguration(ctx, uuid, configuration)



Update a Slack endpoint configuration by it's UUID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 
**configuration** | [**SlackEndpointConfiguration**](SlackEndpointConfiguration.md)|  | 

### Return type

[**SlackEndpointConfiguration**](SlackEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSlackSelector

> Selector UpdateSlackSelector(ctx, configurationUuid, selectorUuid, selector)



Update a selector mapped to a Slack endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationUuid** | **string**|  | 
**selectorUuid** | **string**|  | 
**selector** | [**Selector**](Selector.md)|  | 

### Return type

[**Selector**](Selector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSmtpConfiguration

> SmtpEndpointConfiguration UpdateSmtpConfiguration(ctx, uuid, configuration)



Update a SMTP endpoint configuration by it's UUID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 
**configuration** | [**SmtpEndpointConfiguration**](SmtpEndpointConfiguration.md)|  | 

### Return type

[**SmtpEndpointConfiguration**](SMTPEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSmtpSelector

> Selector UpdateSmtpSelector(ctx, configurationUuid, selectorUuid, selector)



Update a selector mapped to a SMTP endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationUuid** | **string**|  | 
**selectorUuid** | **string**|  | 
**selector** | [**Selector**](Selector.md)|  | 

### Return type

[**Selector**](Selector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTeamsConfiguration

> TeamsEndpointConfiguration UpdateTeamsConfiguration(ctx, uuid, configuration)



Update a Teams endpoint configuration by it's UUID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 
**configuration** | [**TeamsEndpointConfiguration**](TeamsEndpointConfiguration.md)|  | 

### Return type

[**TeamsEndpointConfiguration**](TeamsEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTeamsSelector

> Selector UpdateTeamsSelector(ctx, configurationUuid, selectorUuid, selector)



Update a selector mapped to a Teams endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationUuid** | **string**|  | 
**selectorUuid** | **string**|  | 
**selector** | [**Selector**](Selector.md)|  | 

### Return type

[**Selector**](Selector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhookConfiguration

> WebhookEndpointConfiguration UpdateWebhookConfiguration(ctx, uuid, configuration)



Update a Webhook endpoint configuration by it's UUID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 
**configuration** | [**WebhookEndpointConfiguration**](WebhookEndpointConfiguration.md)|  | 

### Return type

[**WebhookEndpointConfiguration**](WebhookEndpointConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhookSelector

> Selector UpdateWebhookSelector(ctx, configurationUuid, selectorUuid, selector)



Update a selector mapped to a Webhook endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationUuid** | **string**|  | 
**selectorUuid** | **string**|  | 
**selector** | [**Selector**](Selector.md)|  | 

### Return type

[**Selector**](Selector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VersionCheck

> ServiceVersion VersionCheck(ctx, )



Returns the version object for the service, including db schema version info

### Required Parameters

This endpoint does not need any parameter.

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

