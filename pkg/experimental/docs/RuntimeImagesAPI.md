# \RuntimeImagesAPI

All URIs are relative to */exp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetKubernetesImages**](RuntimeImagesAPI.md#GetKubernetesImages) | **Get** /kubernetes-images-summary | Summary of Kubernetes runtime images for this account



## GetKubernetesImages

> AnchoreSummaryRuntimeImages GetKubernetesImages(ctx).OrderBy(orderBy).OrderByDescending(orderByDescending).Context(context).Evaluated(evaluated).Compliant(compliant).Severity(severity).Filter(filter).Limit(limit).Page(page).Execute()

Summary of Kubernetes runtime images for this account



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
	orderBy := []string{"OrderBy_example"} // []string | List of field name(s) to order by, ascending by default (optional)
	orderByDescending := []bool{false} // []bool | Configures the sort order of each specified order_by column to be descending (true) instead of ascending (false) (optional)
	context := "context_example" // string | Matches on an exact cluster name or exact cluster/namespace combination (optional)
	evaluated := true // bool | Filter by evaluated (optional)
	compliant := true // bool | Filter by compliance (optional)
	severity := []string{"Severity_example"} // []string | Results will include images with at least one vulnerability matching any of the given severities (optional)
	filter := "filter_example" // string | Filter against 'image_tag', 'reported_digest', or 'context' fields, using partial or full string match (optional)
	limit := int32(56) // int32 | Maximum number of rows to return (optional)
	page := int32(56) // int32 | Page number to return, one's based (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RuntimeImagesAPI.GetKubernetesImages(context.Background()).OrderBy(orderBy).OrderByDescending(orderByDescending).Context(context).Evaluated(evaluated).Compliant(compliant).Severity(severity).Filter(filter).Limit(limit).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RuntimeImagesAPI.GetKubernetesImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKubernetesImages`: AnchoreSummaryRuntimeImages
	fmt.Fprintf(os.Stdout, "Response from `RuntimeImagesAPI.GetKubernetesImages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderBy** | **[]string** | List of field name(s) to order by, ascending by default | 
 **orderByDescending** | **[]bool** | Configures the sort order of each specified order_by column to be descending (true) instead of ascending (false) | 
 **context** | **string** | Matches on an exact cluster name or exact cluster/namespace combination | 
 **evaluated** | **bool** | Filter by evaluated | 
 **compliant** | **bool** | Filter by compliance | 
 **severity** | **[]string** | Results will include images with at least one vulnerability matching any of the given severities | 
 **filter** | **string** | Filter against &#39;image_tag&#39;, &#39;reported_digest&#39;, or &#39;context&#39; fields, using partial or full string match | 
 **limit** | **int32** | Maximum number of rows to return | 
 **page** | **int32** | Page number to return, one&#39;s based | [default to 1]

### Return type

[**AnchoreSummaryRuntimeImages**](AnchoreSummaryRuntimeImages.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

