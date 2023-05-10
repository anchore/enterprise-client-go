# \ActionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddActionPlan**](ActionsApi.md#AddActionPlan) | **Post** /actions | Submits an Action Plan
[**GetActionPlans**](ActionsApi.md#GetActionPlans) | **Get** /actions | Gets a list of submitted action (remediation) plans



## AddActionPlan

> ActionPlan AddActionPlan(ctx).ActionPlan(actionPlan).Execute()

Submits an Action Plan



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
    actionPlan := *openapiclient.NewActionPlan() // ActionPlan | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ActionsApi.AddActionPlan(context.Background()).ActionPlan(actionPlan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.AddActionPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddActionPlan`: ActionPlan
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.AddActionPlan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddActionPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **actionPlan** | [**ActionPlan**](ActionPlan.md) |  | 

### Return type

[**ActionPlan**](ActionPlan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActionPlans

> []ActionPlan GetActionPlans(ctx).ImageTag(imageTag).ImageDigest(imageDigest).CreatedAfter(createdAfter).XAnchoreAccount(xAnchoreAccount).Execute()

Gets a list of submitted action (remediation) plans



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    imageTag := "imageTag_example" // string |  (optional)
    imageDigest := "imageDigest_example" // string |  (optional)
    createdAfter := time.Now() // time.Time | RFC 3339 formatted UTC timestamp to filter out action plans that were only created after this date (optional)
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ActionsApi.GetActionPlans(context.Background()).ImageTag(imageTag).ImageDigest(imageDigest).CreatedAfter(createdAfter).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.GetActionPlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetActionPlans`: []ActionPlan
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.GetActionPlans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActionPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imageTag** | **string** |  | 
 **imageDigest** | **string** |  | 
 **createdAfter** | **time.Time** | RFC 3339 formatted UTC timestamp to filter out action plans that were only created after this date | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]ActionPlan**](ActionPlan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

