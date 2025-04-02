# \ImagesAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddImage**](ImagesAPI.md#AddImage) | **Post** /images | Submit a new image for analysis by the engine
[**DeleteImage**](ImagesAPI.md#DeleteImage) | **Delete** /images/{image_digest} | Delete an image analysis
[**DeleteImagesAsync**](ImagesAPI.md#DeleteImagesAsync) | **Delete** /images | Bulk mark images for deletion
[**GetImage**](ImagesAPI.md#GetImage) | **Get** /images/{image_digest} | Get image metadata
[**GetImageAncestors**](ImagesAPI.md#GetImageAncestors) | **Get** /images/{image_digest}/ancestors | Return the list of ancestor images for the given image
[**GetImageContentByType**](ImagesAPI.md#GetImageContentByType) | **Get** /images/{image_digest}/content/{content_type} | Get the content of an image by type
[**GetImageContentByTypeFiles**](ImagesAPI.md#GetImageContentByTypeFiles) | **Get** /images/{image_digest}/content/files | Get the content of an image by type files
[**GetImageContentByTypeJavaPackage**](ImagesAPI.md#GetImageContentByTypeJavaPackage) | **Get** /images/{image_digest}/content/java | Get the content of an image by type java
[**GetImageContentByTypeLicense**](ImagesAPI.md#GetImageContentByTypeLicense) | **Get** /images/{image_digest}/content/licenses | Get package licenses for an image
[**GetImageContentByTypeMalware**](ImagesAPI.md#GetImageContentByTypeMalware) | **Get** /images/{image_digest}/content/malware | Get the content of an image by type malware
[**GetImageContentSummary**](ImagesAPI.md#GetImageContentSummary) | **Get** /images/{image_digest}/content-summary | Get image content summary
[**GetImageMetadataByType**](ImagesAPI.md#GetImageMetadataByType) | **Get** /images/{image_digest}/metadata/{metadata_type} | Get the metadata of an image by type
[**GetImagePolicyCheckByDigest**](ImagesAPI.md#GetImagePolicyCheckByDigest) | **Get** /images/{image_digest}/check | Check policy evaluation status for image
[**GetImageSbomCyclonedxJson**](ImagesAPI.md#GetImageSbomCyclonedxJson) | **Get** /images/{image_digest}/sboms/cyclonedx-json | Get image sbom in the CycloneDX format
[**GetImageSbomNativeJson**](ImagesAPI.md#GetImageSbomNativeJson) | **Get** /images/{image_digest}/sboms/native-json | Get image sbom in the native Anchore format
[**GetImageSbomSpdxJson**](ImagesAPI.md#GetImageSbomSpdxJson) | **Get** /images/{image_digest}/sboms/spdx-json | Get image sbom in the SPDX format
[**GetImageVulnerabilitiesByDigest**](ImagesAPI.md#GetImageVulnerabilitiesByDigest) | **Get** /images/{image_digest}/vuln/{vuln_type} | Get vulnerabilities by type
[**GetImageVulnerabilityTypes**](ImagesAPI.md#GetImageVulnerabilityTypes) | **Get** /images/{image_digest}/vuln | Get vulnerability types
[**ListFileContentSearchResults**](ImagesAPI.md#ListFileContentSearchResults) | **Get** /images/{image_digest}/artifacts/file-content-search | Return a list of analyzer artifacts of the specified type
[**ListImageContent**](ImagesAPI.md#ListImageContent) | **Get** /images/{image_digest}/content | List image content types
[**ListImageMetadata**](ImagesAPI.md#ListImageMetadata) | **Get** /images/{image_digest}/metadata | List image metadata types
[**ListImages**](ImagesAPI.md#ListImages) | **Get** /images | List all visible images
[**ListRetrievedFiles**](ImagesAPI.md#ListRetrievedFiles) | **Get** /images/{image_digest}/artifacts/retrieved-files | Return a list of analyzer artifacts of the specified type
[**ListSecretSearchResults**](ImagesAPI.md#ListSecretSearchResults) | **Get** /images/{image_digest}/artifacts/secret-search | Return a list of analyzer artifacts of the specified type
[**SummaryImageCounts**](ImagesAPI.md#SummaryImageCounts) | **Get** /summaries/image-counts | Image summary counts
[**SummaryImageTags**](ImagesAPI.md#SummaryImageTags) | **Get** /summaries/image-tags | Summarize image tags



## AddImage

> AnchoreImage AddImage(ctx).Image(image).Force(force).AutoSubscribe(autoSubscribe).XAnchoreAccount(xAnchoreAccount).Execute()

Submit a new image for analysis by the engine



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
	image := *openapiclient.NewImageAnalysisRequest() // ImageAnalysisRequest | 
	force := true // bool | Override any existing entry in the system (optional)
	autoSubscribe := true // bool | Indicates if tag will be subscribed for registry updates monitoring (optional)
	xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.AddImage(context.Background()).Image(image).Force(force).AutoSubscribe(autoSubscribe).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.AddImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddImage`: AnchoreImage
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.AddImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **image** | [**ImageAnalysisRequest**](ImageAnalysisRequest.md) |  | 
 **force** | **bool** | Override any existing entry in the system | 
 **autoSubscribe** | **bool** | Indicates if tag will be subscribed for registry updates monitoring | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**AnchoreImage**](AnchoreImage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteImage

> DeleteImageResponse DeleteImage(ctx, imageDigest).Force(force).XAnchoreAccount(xAnchoreAccount).Execute()

Delete an image analysis

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
	force := true // bool |  (optional)
	xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.DeleteImage(context.Background(), imageDigest).Force(force).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.DeleteImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteImage`: DeleteImageResponse
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.DeleteImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** |  | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**DeleteImageResponse**](DeleteImageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteImagesAsync

> []DeleteImageResponse DeleteImagesAsync(ctx).ImageDigests(imageDigests).Force(force).XAnchoreAccount(xAnchoreAccount).Execute()

Bulk mark images for deletion



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
	imageDigests := []string{"Inner_example"} // []string | 
	force := true // bool |  (optional)
	xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.DeleteImagesAsync(context.Background()).ImageDigests(imageDigests).Force(force).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.DeleteImagesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteImagesAsync`: []DeleteImageResponse
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.DeleteImagesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteImagesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imageDigests** | **[]string** |  | 
 **force** | **bool** |  | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]DeleteImageResponse**](DeleteImageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImage

> AnchoreImage GetImage(ctx, imageDigest).XAnchoreAccount(xAnchoreAccount).Execute()

Get image metadata

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
	xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.GetImage(context.Background(), imageDigest).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.GetImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImage`: AnchoreImage
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.GetImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**AnchoreImage**](AnchoreImage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageAncestors

> []ImageAncestor GetImageAncestors(ctx, imageDigest).XAnchoreAccount(xAnchoreAccount).Execute()

Return the list of ancestor images for the given image



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
	xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.GetImageAncestors(context.Background(), imageDigest).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.GetImageAncestors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImageAncestors`: []ImageAncestor
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.GetImageAncestors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageAncestorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]ImageAncestor**](ImageAncestor.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageContentByType

> ContentPackageResponse GetImageContentByType(ctx, imageDigest, contentType).XAnchoreAccount(xAnchoreAccount).Execute()

Get the content of an image by type

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
	contentType := "contentType_example" // string | 
	xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.GetImageContentByType(context.Background(), imageDigest, contentType).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.GetImageContentByType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImageContentByType`: ContentPackageResponse
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.GetImageContentByType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string** |  | 
**contentType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageContentByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**ContentPackageResponse**](ContentPackageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageContentByTypeFiles

> ContentFilesResponse GetImageContentByTypeFiles(ctx, imageDigest).XAnchoreAccount(xAnchoreAccount).Execute()

Get the content of an image by type files

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
	xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.GetImageContentByTypeFiles(context.Background(), imageDigest).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.GetImageContentByTypeFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImageContentByTypeFiles`: ContentFilesResponse
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.GetImageContentByTypeFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageContentByTypeFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**ContentFilesResponse**](ContentFilesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageContentByTypeJavaPackage

> ContentJavaPackageResponse GetImageContentByTypeJavaPackage(ctx, imageDigest).XAnchoreAccount(xAnchoreAccount).Execute()

Get the content of an image by type java

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
	xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.GetImageContentByTypeJavaPackage(context.Background(), imageDigest).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.GetImageContentByTypeJavaPackage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImageContentByTypeJavaPackage`: ContentJavaPackageResponse
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.GetImageContentByTypeJavaPackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageContentByTypeJavaPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**ContentJavaPackageResponse**](ContentJavaPackageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageContentByTypeLicense

> LicenseReviewResponse GetImageContentByTypeLicense(ctx, imageDigest).XAnchoreAccount(xAnchoreAccount).Execute()

Get package licenses for an image



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
	xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.GetImageContentByTypeLicense(context.Background(), imageDigest).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.GetImageContentByTypeLicense``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImageContentByTypeLicense`: LicenseReviewResponse
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.GetImageContentByTypeLicense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageContentByTypeLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**LicenseReviewResponse**](LicenseReviewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageContentByTypeMalware

> ContentMalwareResponse GetImageContentByTypeMalware(ctx, imageDigest).XAnchoreAccount(xAnchoreAccount).Execute()

Get the content of an image by type malware

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
	xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.GetImageContentByTypeMalware(context.Background(), imageDigest).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.GetImageContentByTypeMalware``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImageContentByTypeMalware`: ContentMalwareResponse
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.GetImageContentByTypeMalware`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageContentByTypeMalwareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**ContentMalwareResponse**](ContentMalwareResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageContentSummary

> ImageContentSummary GetImageContentSummary(ctx, imageDigest).XAnchoreAccount(xAnchoreAccount).Execute()

Get image content summary

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
	xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.GetImageContentSummary(context.Background(), imageDigest).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.GetImageContentSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImageContentSummary`: ImageContentSummary
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.GetImageContentSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageContentSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**ImageContentSummary**](ImageContentSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageMetadataByType

> MetadataResponse GetImageMetadataByType(ctx, imageDigest, metadataType).XAnchoreAccount(xAnchoreAccount).Execute()

Get the metadata of an image by type

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
	metadataType := "metadataType_example" // string | 
	xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.GetImageMetadataByType(context.Background(), imageDigest, metadataType).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.GetImageMetadataByType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImageMetadataByType`: MetadataResponse
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.GetImageMetadataByType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string** |  | 
**metadataType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageMetadataByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**MetadataResponse**](MetadataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImagePolicyCheckByDigest

> PolicyEvaluation GetImagePolicyCheckByDigest(ctx, imageDigest).Tag(tag).PolicyId(policyId).Detail(detail).History(history).Interactive(interactive).BaseDigest(baseDigest).XAnchoreAccount(xAnchoreAccount).Execute()

Check policy evaluation status for image



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
	tag := "tag_example" // string | 
	policyId := "policyId_example" // string |  (optional)
	detail := true // bool |  (optional) (default to true)
	history := true // bool |  (optional) (default to false)
	interactive := true // bool |  (optional) (default to false)
	baseDigest := "baseDigest_example" // string | Digest of a base image. If specified the evaluation will indicate results inherited from the base image. Can specify \"auto\" to have the base image automatically calculated. (optional)
	xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.GetImagePolicyCheckByDigest(context.Background(), imageDigest).Tag(tag).PolicyId(policyId).Detail(detail).History(history).Interactive(interactive).BaseDigest(baseDigest).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.GetImagePolicyCheckByDigest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImagePolicyCheckByDigest`: PolicyEvaluation
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.GetImagePolicyCheckByDigest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImagePolicyCheckByDigestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tag** | **string** |  | 
 **policyId** | **string** |  | 
 **detail** | **bool** |  | [default to true]
 **history** | **bool** |  | [default to false]
 **interactive** | **bool** |  | [default to false]
 **baseDigest** | **string** | Digest of a base image. If specified the evaluation will indicate results inherited from the base image. Can specify \&quot;auto\&quot; to have the base image automatically calculated. | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**PolicyEvaluation**](PolicyEvaluation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageSbomCyclonedxJson

> string GetImageSbomCyclonedxJson(ctx, imageDigest).XAnchoreAccount(xAnchoreAccount).Execute()

Get image sbom in the CycloneDX format

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
	xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.GetImageSbomCyclonedxJson(context.Background(), imageDigest).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.GetImageSbomCyclonedxJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImageSbomCyclonedxJson`: string
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.GetImageSbomCyclonedxJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageSbomCyclonedxJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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


## GetImageSbomNativeJson

> string GetImageSbomNativeJson(ctx, imageDigest).XAnchoreAccount(xAnchoreAccount).Execute()

Get image sbom in the native Anchore format

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
	xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.GetImageSbomNativeJson(context.Background(), imageDigest).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.GetImageSbomNativeJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImageSbomNativeJson`: string
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.GetImageSbomNativeJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageSbomNativeJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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


## GetImageSbomSpdxJson

> string GetImageSbomSpdxJson(ctx, imageDigest).XAnchoreAccount(xAnchoreAccount).Execute()

Get image sbom in the SPDX format

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
	xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.GetImageSbomSpdxJson(context.Background(), imageDigest).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.GetImageSbomSpdxJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImageSbomSpdxJson`: string
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.GetImageSbomSpdxJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageSbomSpdxJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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


## GetImageVulnerabilitiesByDigest

> ImagePackageVulnerabilityResponse GetImageVulnerabilitiesByDigest(ctx, imageDigest, vulnType).ForceRefresh(forceRefresh).IncludeVulnDescription(includeVulnDescription).VendorOnly(vendorOnly).BaseDigest(baseDigest).XAnchoreAccount(xAnchoreAccount).Execute()

Get vulnerabilities by type

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
	vulnType := "vulnType_example" // string | 
	forceRefresh := true // bool |  (optional) (default to false)
	includeVulnDescription := true // bool |  (optional) (default to false)
	vendorOnly := true // bool | Filter results to include only vulnerabilities that are not marked as invalid by upstream OS vendor data. When set to true, it will filter out all vulnerabilities where `will_not_fix` is False. If false all vulnerabilities are returned regardless of `will_not_fix` (optional) (default to true)
	baseDigest := "baseDigest_example" // string | Digest of a base image. If specified the vulnerabilities will indicate inheritance from the base image. Can specify \"auto\" to have the base image automatically calculated. (optional)
	xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.GetImageVulnerabilitiesByDigest(context.Background(), imageDigest, vulnType).ForceRefresh(forceRefresh).IncludeVulnDescription(includeVulnDescription).VendorOnly(vendorOnly).BaseDigest(baseDigest).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.GetImageVulnerabilitiesByDigest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImageVulnerabilitiesByDigest`: ImagePackageVulnerabilityResponse
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.GetImageVulnerabilitiesByDigest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string** |  | 
**vulnType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageVulnerabilitiesByDigestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **forceRefresh** | **bool** |  | [default to false]
 **includeVulnDescription** | **bool** |  | [default to false]
 **vendorOnly** | **bool** | Filter results to include only vulnerabilities that are not marked as invalid by upstream OS vendor data. When set to true, it will filter out all vulnerabilities where &#x60;will_not_fix&#x60; is False. If false all vulnerabilities are returned regardless of &#x60;will_not_fix&#x60; | [default to true]
 **baseDigest** | **string** | Digest of a base image. If specified the vulnerabilities will indicate inheritance from the base image. Can specify \&quot;auto\&quot; to have the base image automatically calculated. | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**ImagePackageVulnerabilityResponse**](ImagePackageVulnerabilityResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageVulnerabilityTypes

> []string GetImageVulnerabilityTypes(ctx, imageDigest).XAnchoreAccount(xAnchoreAccount).Execute()

Get vulnerability types

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
	xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.GetImageVulnerabilityTypes(context.Background(), imageDigest).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.GetImageVulnerabilityTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImageVulnerabilityTypes`: []string
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.GetImageVulnerabilityTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageVulnerabilityTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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
	resp, r, err := apiClient.ImagesAPI.ListFileContentSearchResults(context.Background(), imageDigest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.ListFileContentSearchResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFileContentSearchResults`: []FileContentSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.ListFileContentSearchResults`: %v\n", resp)
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


## ListImageContent

> []string ListImageContent(ctx, imageDigest).XAnchoreAccount(xAnchoreAccount).Execute()

List image content types

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
	xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.ListImageContent(context.Background(), imageDigest).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.ListImageContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListImageContent`: []string
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.ListImageContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListImageContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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


## ListImageMetadata

> []string ListImageMetadata(ctx, imageDigest).XAnchoreAccount(xAnchoreAccount).Execute()

List image metadata types

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
	xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.ListImageMetadata(context.Background(), imageDigest).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.ListImageMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListImageMetadata`: []string
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.ListImageMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListImageMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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


## ListImages

> AnchoreImageList ListImages(ctx).ImageId(imageId).History(history).FullTag(fullTag).ImageStatus(imageStatus).AnalysisStatus(analysisStatus).AnalyzedSince(analyzedSince).XAnchoreAccount(xAnchoreAccount).Execute()

List all visible images



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
	imageId := "imageId_example" // string | Filter results matching image ID (optional)
	history := true // bool | Include image history in the response (optional)
	fullTag := "fullTag_example" // string | Full docker-pull string to filter results by (e.g. docker.io/library/nginx:latest, or myhost.com:5000/testimages:v1.1.1) (optional)
	imageStatus := "imageStatus_example" // string | Filter by image_status value on the record. Default if omitted is 'active'. (optional) (default to "active")
	analysisStatus := "analysisStatus_example" // string | Filter by analysis_status value on the record. (optional)
	analyzedSince := "analyzedSince_example" // string | Filter by images analyzed on or after the specified datetime (optional)
	xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.ListImages(context.Background()).ImageId(imageId).History(history).FullTag(fullTag).ImageStatus(imageStatus).AnalysisStatus(analysisStatus).AnalyzedSince(analyzedSince).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.ListImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListImages`: AnchoreImageList
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.ListImages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imageId** | **string** | Filter results matching image ID | 
 **history** | **bool** | Include image history in the response | 
 **fullTag** | **string** | Full docker-pull string to filter results by (e.g. docker.io/library/nginx:latest, or myhost.com:5000/testimages:v1.1.1) | 
 **imageStatus** | **string** | Filter by image_status value on the record. Default if omitted is &#39;active&#39;. | [default to &quot;active&quot;]
 **analysisStatus** | **string** | Filter by analysis_status value on the record. | 
 **analyzedSince** | **string** | Filter by images analyzed on or after the specified datetime | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**AnchoreImageList**](AnchoreImageList.md)

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
	resp, r, err := apiClient.ImagesAPI.ListRetrievedFiles(context.Background(), imageDigest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.ListRetrievedFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRetrievedFiles`: []RetrievedFile
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.ListRetrievedFiles`: %v\n", resp)
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
	resp, r, err := apiClient.ImagesAPI.ListSecretSearchResults(context.Background(), imageDigest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.ListSecretSearchResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSecretSearchResults`: []SecretSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.ListSecretSearchResults`: %v\n", resp)
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


## SummaryImageCounts

> AnchoreImageSummaryCounts SummaryImageCounts(ctx).ImageStatus(imageStatus).Registry(registry).Repo(repo).XAnchoreAccount(xAnchoreAccount).Execute()

Image summary counts



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
	imageStatus := []string{"ImageStatus_example"} // []string | Filter images in one or more states such as active, deleting. Defaults to active images only if unspecified (optional) (default to ["active"])
	registry := "registry_example" // string | Filter by registry (optional)
	repo := "repo_example" // string | Filter by repo (optional)
	xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.SummaryImageCounts(context.Background()).ImageStatus(imageStatus).Registry(registry).Repo(repo).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.SummaryImageCounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SummaryImageCounts`: AnchoreImageSummaryCounts
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.SummaryImageCounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSummaryImageCountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imageStatus** | **[]string** | Filter images in one or more states such as active, deleting. Defaults to active images only if unspecified | [default to [&quot;active&quot;]]
 **registry** | **string** | Filter by registry | 
 **repo** | **string** | Filter by repo | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**AnchoreImageSummaryCounts**](AnchoreImageSummaryCounts.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SummaryImageTags

> AnchoreImageTagSummaryList SummaryImageTags(ctx).ImageStatus(imageStatus).AnalysisStatus(analysisStatus).AnalyzedSince(analyzedSince).Registry(registry).Repository(repository).Tag(tag).Runtime(runtime).OrderBy(orderBy).OrderByDescending(orderByDescending).Filter(filter).Limit(limit).Page(page).XAnchoreAccount(xAnchoreAccount).Execute()

Summarize image tags



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
	imageStatus := []string{"ImageStatus_example"} // []string | Filter images in one or more states such as active, deleting. Defaults to active images only if unspecified (optional) (default to ["active"])
	analysisStatus := []string{"AnalysisStatus_example"} // []string | Filter images in one or more analysis_status such as analyzed, not_analyzed, analysis_failed. Defaults to unfiltered if unspecified (optional) (default to ["all"])
	analyzedSince := "analyzedSince_example" // string | Filter images analyzed on or after the specified datetime (optional)
	registry := "registry_example" // string | A registry name to filter result by (e.g. \"docker.io\") (optional)
	repository := "repository_example" // string | A repository name to filter results by (e.g. \"jboss/keycloak\") (optional)
	tag := "tag_example" // string | A tag value to filter results by (e.g. \"latest\", or \"v1.2.0\") (optional)
	runtime := true // bool | Filter by images with runtime inventory (optional)
	orderBy := []string{"OrderBy_example"} // []string | List of field name(s) to order by, ascending by default (optional)
	orderByDescending := []bool{false} // []bool | Configures the sort order of each specified order_by column to be descending (true) instead of ascending (false) (optional)
	filter := "filter_example" // string | Filter by 'image_digest' or 'full_tag' fields, using partial or full string match (optional)
	limit := int32(56) // int32 | Maximum number of rows to return (optional)
	page := int32(56) // int32 | Page number to return, one's based (optional) (default to 1)
	xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.SummaryImageTags(context.Background()).ImageStatus(imageStatus).AnalysisStatus(analysisStatus).AnalyzedSince(analyzedSince).Registry(registry).Repository(repository).Tag(tag).Runtime(runtime).OrderBy(orderBy).OrderByDescending(orderByDescending).Filter(filter).Limit(limit).Page(page).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.SummaryImageTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SummaryImageTags`: AnchoreImageTagSummaryList
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.SummaryImageTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSummaryImageTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imageStatus** | **[]string** | Filter images in one or more states such as active, deleting. Defaults to active images only if unspecified | [default to [&quot;active&quot;]]
 **analysisStatus** | **[]string** | Filter images in one or more analysis_status such as analyzed, not_analyzed, analysis_failed. Defaults to unfiltered if unspecified | [default to [&quot;all&quot;]]
 **analyzedSince** | **string** | Filter images analyzed on or after the specified datetime | 
 **registry** | **string** | A registry name to filter result by (e.g. \&quot;docker.io\&quot;) | 
 **repository** | **string** | A repository name to filter results by (e.g. \&quot;jboss/keycloak\&quot;) | 
 **tag** | **string** | A tag value to filter results by (e.g. \&quot;latest\&quot;, or \&quot;v1.2.0\&quot;) | 
 **runtime** | **bool** | Filter by images with runtime inventory | 
 **orderBy** | **[]string** | List of field name(s) to order by, ascending by default | 
 **orderByDescending** | **[]bool** | Configures the sort order of each specified order_by column to be descending (true) instead of ascending (false) | 
 **filter** | **string** | Filter by &#39;image_digest&#39; or &#39;full_tag&#39; fields, using partial or full string match | 
 **limit** | **int32** | Maximum number of rows to return | 
 **page** | **int32** | Page number to return, one&#39;s based | [default to 1]
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**AnchoreImageTagSummaryList**](AnchoreImageTagSummaryList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

