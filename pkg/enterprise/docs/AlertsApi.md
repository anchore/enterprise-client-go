# \AlertsAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAlertSummaries**](AlertsAPI.md#GetAlertSummaries) | **Get** /alerts/summaries | List all alert summaries scoped to the account
[**GetComplianceViolationAlert**](AlertsAPI.md#GetComplianceViolationAlert) | **Get** /alerts/compliance-violations/{uuid} | Get compliance violation alert by id
[**GetComplianceViolationAlerts**](AlertsAPI.md#GetComplianceViolationAlerts) | **Get** /alerts/compliance-violations | List all compliance violation alerts scoped to the account
[**UpdateComplianceViolationAlertState**](AlertsAPI.md#UpdateComplianceViolationAlertState) | **Patch** /alerts/compliance-violations/{uuid} | Open or close a compliance violation alert



## GetAlertSummaries

> []AlertSummary GetAlertSummaries(ctx).Page(page).Limit(limit).Type_(type_).State(state).CreatedAfter(createdAfter).CreatedBefore(createdBefore).ResourceLabel(resourceLabel).XAnchoreAccount(xAnchoreAccount).Execute()

List all alert summaries scoped to the account



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
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	type_ := "type__example" // string | Filter for alerts based on the type such as compliance violation (optional) (default to "all")
	state := "state_example" // string | Filter for alerts by current state, defaults to open alerts unless specified (optional) (default to "open")
	createdAfter := time.Now() // time.Time | Filter for alerts generated after the timestamp (optional)
	createdBefore := time.Now() // time.Time | Filter for alerts generated before the timestamp (optional)
	resourceLabel := []string{"Inner_example"} // []string | Filter for alerts associated with a resource where the label in key=value format such as tag=docker.io/library/alpine:latest or repository=library/alpine (optional)
	xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.GetAlertSummaries(context.Background()).Page(page).Limit(limit).Type_(type_).State(state).CreatedAfter(createdAfter).CreatedBefore(createdBefore).ResourceLabel(resourceLabel).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.GetAlertSummaries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertSummaries`: []AlertSummary
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.GetAlertSummaries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertSummariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]
 **type_** | **string** | Filter for alerts based on the type such as compliance violation | [default to &quot;all&quot;]
 **state** | **string** | Filter for alerts by current state, defaults to open alerts unless specified | [default to &quot;open&quot;]
 **createdAfter** | **time.Time** | Filter for alerts generated after the timestamp | 
 **createdBefore** | **time.Time** | Filter for alerts generated before the timestamp | 
 **resourceLabel** | **[]string** | Filter for alerts associated with a resource where the label in key&#x3D;value format such as tag&#x3D;docker.io/library/alpine:latest or repository&#x3D;library/alpine | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]AlertSummary**](AlertSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComplianceViolationAlert

> ComplianceViolationAlert GetComplianceViolationAlert(ctx, uuid).XAnchoreAccount(xAnchoreAccount).Execute()

Get compliance violation alert by id



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
	uuid := "uuid_example" // string | Identifier for the alert
	xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.GetComplianceViolationAlert(context.Background(), uuid).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.GetComplianceViolationAlert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComplianceViolationAlert`: ComplianceViolationAlert
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.GetComplianceViolationAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Identifier for the alert | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComplianceViolationAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**ComplianceViolationAlert**](ComplianceViolationAlert.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComplianceViolationAlerts

> []ComplianceViolationAlert GetComplianceViolationAlerts(ctx).Page(page).Limit(limit).State(state).CreatedAfter(createdAfter).CreatedBefore(createdBefore).ResourceImageDigest(resourceImageDigest).ResourceImageTag(resourceImageTag).ResourceRegistry(resourceRegistry).ResourceRepository(resourceRepository).XAnchoreAccount(xAnchoreAccount).Execute()

List all compliance violation alerts scoped to the account



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
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	state := "state_example" // string | Filter for alerts by current state, defaults to open alerts unless specified (optional) (default to "open")
	createdAfter := time.Now() // time.Time | Filter for alerts generated after the timestamp (optional)
	createdBefore := time.Now() // time.Time | Filter for alerts generated before the timestamp (optional)
	resourceImageDigest := "resourceImageDigest_example" // string | Filter for alerts associated with image digest (optional)
	resourceImageTag := "resourceImageTag_example" // string | Filter for alerts generated for the tag (optional)
	resourceRegistry := "resourceRegistry_example" // string | Filter for alerts associated with registry (optional)
	resourceRepository := "resourceRepository_example" // string | Filter for alerts associated with repository (optional)
	xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.GetComplianceViolationAlerts(context.Background()).Page(page).Limit(limit).State(state).CreatedAfter(createdAfter).CreatedBefore(createdBefore).ResourceImageDigest(resourceImageDigest).ResourceImageTag(resourceImageTag).ResourceRegistry(resourceRegistry).ResourceRepository(resourceRepository).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.GetComplianceViolationAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComplianceViolationAlerts`: []ComplianceViolationAlert
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.GetComplianceViolationAlerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetComplianceViolationAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]
 **state** | **string** | Filter for alerts by current state, defaults to open alerts unless specified | [default to &quot;open&quot;]
 **createdAfter** | **time.Time** | Filter for alerts generated after the timestamp | 
 **createdBefore** | **time.Time** | Filter for alerts generated before the timestamp | 
 **resourceImageDigest** | **string** | Filter for alerts associated with image digest | 
 **resourceImageTag** | **string** | Filter for alerts generated for the tag | 
 **resourceRegistry** | **string** | Filter for alerts associated with registry | 
 **resourceRepository** | **string** | Filter for alerts associated with repository | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]ComplianceViolationAlert**](ComplianceViolationAlert.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateComplianceViolationAlertState

> ComplianceViolationAlert UpdateComplianceViolationAlertState(ctx, uuid).Body(body).XAnchoreAccount(xAnchoreAccount).Execute()

Open or close a compliance violation alert



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
	uuid := "uuid_example" // string | Identifier for the alert
	body := *openapiclient.NewComplianceViolationAlertState("State_example") // ComplianceViolationAlertState | 
	xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.UpdateComplianceViolationAlertState(context.Background(), uuid).Body(body).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.UpdateComplianceViolationAlertState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateComplianceViolationAlertState`: ComplianceViolationAlert
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.UpdateComplianceViolationAlertState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Identifier for the alert | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateComplianceViolationAlertStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ComplianceViolationAlertState**](ComplianceViolationAlertState.md) |  | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**ComplianceViolationAlert**](ComplianceViolationAlert.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

