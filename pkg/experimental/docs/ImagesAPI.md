# \ImagesAPI

All URIs are relative to */exp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SummaryImages**](ImagesAPI.md#SummaryImages) | **Get** /images | Summarize images by registry and repository
[**SummaryImagesForRegistry**](ImagesAPI.md#SummaryImagesForRegistry) | **Get** /images/registry | Summarize image registry
[**SummaryImagesForRepo**](ImagesAPI.md#SummaryImagesForRepo) | **Get** /images/repo | Summarize image repository
[**SummaryImagesForTag**](ImagesAPI.md#SummaryImagesForTag) | **Get** /images/tag | Summary image digests for a registry/repo:tag



## SummaryImages

> AnchoreSummaryImages SummaryImages(ctx).OrderBy(orderBy).OrderByDescending(orderByDescending).Filter(filter).Limit(limit).Page(page).Execute()

Summarize images by registry and repository



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
	filter := "filter_example" // string | Filter against 'registry', 'repo', or 'tag_count' fields, using partial or full string match (optional)
	limit := int32(56) // int32 | Maximum number of rows to return (optional)
	page := int32(56) // int32 | Page number to return, one's based (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.SummaryImages(context.Background()).OrderBy(orderBy).OrderByDescending(orderByDescending).Filter(filter).Limit(limit).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.SummaryImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SummaryImages`: AnchoreSummaryImages
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.SummaryImages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSummaryImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderBy** | **[]string** | List of field name(s) to order by, ascending by default | 
 **orderByDescending** | **[]bool** | Configures the sort order of each specified order_by column to be descending (true) instead of ascending (false) | 
 **filter** | **string** | Filter against &#39;registry&#39;, &#39;repo&#39;, or &#39;tag_count&#39; fields, using partial or full string match | 
 **limit** | **int32** | Maximum number of rows to return | 
 **page** | **int32** | Page number to return, one&#39;s based | [default to 1]

### Return type

[**AnchoreSummaryImages**](AnchoreSummaryImages.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SummaryImagesForRegistry

> AnchoreSummaryImageRegistry SummaryImagesForRegistry(ctx).Registry(registry).OrderBy(orderBy).OrderByDescending(orderByDescending).Filter(filter).Limit(limit).Page(page).Execute()

Summarize image registry



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
	registry := "registry_example" // string | Registry name
	orderBy := []string{"OrderBy_example"} // []string | List of field name(s) to order by, ascending by default (optional)
	orderByDescending := []bool{false} // []bool | Configures the sort order of each specified order_by column to be descending (true) instead of ascending (false) (optional)
	filter := "filter_example" // string | Filter against 'repo' field, using partial or full string match (optional)
	limit := int32(56) // int32 | Maximum number of rows to return (optional)
	page := int32(56) // int32 | Page number to return, one's based (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.SummaryImagesForRegistry(context.Background()).Registry(registry).OrderBy(orderBy).OrderByDescending(orderByDescending).Filter(filter).Limit(limit).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.SummaryImagesForRegistry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SummaryImagesForRegistry`: AnchoreSummaryImageRegistry
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.SummaryImagesForRegistry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSummaryImagesForRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registry** | **string** | Registry name | 
 **orderBy** | **[]string** | List of field name(s) to order by, ascending by default | 
 **orderByDescending** | **[]bool** | Configures the sort order of each specified order_by column to be descending (true) instead of ascending (false) | 
 **filter** | **string** | Filter against &#39;repo&#39; field, using partial or full string match | 
 **limit** | **int32** | Maximum number of rows to return | 
 **page** | **int32** | Page number to return, one&#39;s based | [default to 1]

### Return type

[**AnchoreSummaryImageRegistry**](AnchoreSummaryImageRegistry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SummaryImagesForRepo

> AnchoreSummaryImageRepo SummaryImagesForRepo(ctx).Registry(registry).Repo(repo).OrderBy(orderBy).OrderByDescending(orderByDescending).Filter(filter).Limit(limit).Page(page).Execute()

Summarize image repository



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
	registry := "registry_example" // string | Registry name
	repo := "repo_example" // string | Repository name
	orderBy := []string{"OrderBy_example"} // []string | List of field name(s) to order by, ascending by default (optional)
	orderByDescending := []bool{false} // []bool | Configures the sort order of each specified order_by column to be descending (true) instead of ascending (false) (optional)
	filter := "filter_example" // string | Filter against 'tag', 'analysis_status', or 'image_digest' fields, using partial or full string match (optional)
	limit := int32(56) // int32 | Maximum number of rows to return (optional)
	page := int32(56) // int32 | Page number to return, one's based (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.SummaryImagesForRepo(context.Background()).Registry(registry).Repo(repo).OrderBy(orderBy).OrderByDescending(orderByDescending).Filter(filter).Limit(limit).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.SummaryImagesForRepo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SummaryImagesForRepo`: AnchoreSummaryImageRepo
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.SummaryImagesForRepo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSummaryImagesForRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registry** | **string** | Registry name | 
 **repo** | **string** | Repository name | 
 **orderBy** | **[]string** | List of field name(s) to order by, ascending by default | 
 **orderByDescending** | **[]bool** | Configures the sort order of each specified order_by column to be descending (true) instead of ascending (false) | 
 **filter** | **string** | Filter against &#39;tag&#39;, &#39;analysis_status&#39;, or &#39;image_digest&#39; fields, using partial or full string match | 
 **limit** | **int32** | Maximum number of rows to return | 
 **page** | **int32** | Page number to return, one&#39;s based | [default to 1]

### Return type

[**AnchoreSummaryImageRepo**](AnchoreSummaryImageRepo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SummaryImagesForTag

> AnchoreSummaryImageTag SummaryImagesForTag(ctx).Registry(registry).Repo(repo).Tag(tag).OrderBy(orderBy).OrderByDescending(orderByDescending).Filter(filter).Limit(limit).Page(page).Execute()

Summary image digests for a registry/repo:tag



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
	registry := "registry_example" // string | Registry name
	repo := "repo_example" // string | Repository name
	tag := "tag_example" // string | Tag name
	orderBy := []string{"OrderBy_example"} // []string | List of field name(s) to order by, ascending by default (optional)
	orderByDescending := []bool{false} // []bool | Configures the sort order of each specified order_by column to be descending (true) instead of ascending (false) (optional)
	filter := "filter_example" // string | Filter against 'image_digest', 'image_id', or 'analysis_status' fields, using partial or full string match (optional)
	limit := int32(56) // int32 | Maximum number of rows to return (optional)
	page := int32(56) // int32 | Page number to return, one's based (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.SummaryImagesForTag(context.Background()).Registry(registry).Repo(repo).Tag(tag).OrderBy(orderBy).OrderByDescending(orderByDescending).Filter(filter).Limit(limit).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.SummaryImagesForTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SummaryImagesForTag`: AnchoreSummaryImageTag
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.SummaryImagesForTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSummaryImagesForTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registry** | **string** | Registry name | 
 **repo** | **string** | Repository name | 
 **tag** | **string** | Tag name | 
 **orderBy** | **[]string** | List of field name(s) to order by, ascending by default | 
 **orderByDescending** | **[]bool** | Configures the sort order of each specified order_by column to be descending (true) instead of ascending (false) | 
 **filter** | **string** | Filter against &#39;image_digest&#39;, &#39;image_id&#39;, or &#39;analysis_status&#39; fields, using partial or full string match | 
 **limit** | **int32** | Maximum number of rows to return | 
 **page** | **int32** | Page number to return, one&#39;s based | [default to 1]

### Return type

[**AnchoreSummaryImageTag**](AnchoreSummaryImageTag.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

