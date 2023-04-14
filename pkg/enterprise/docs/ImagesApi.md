# \ImagesApi

All URIs are relative to *http://localhost/enterprise*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetImageAncestors**](ImagesApi.md#GetImageAncestors) | **Get** /images/{image_digest}/ancestors | Return the list of ancestor images for the given image
[**GetImagePolicyCheckByDigest**](ImagesApi.md#GetImagePolicyCheckByDigest) | **Get** /images/{image_digest}/check | Check policy evaluation status for image
[**GetImageVulnerabilitiesByDigest**](ImagesApi.md#GetImageVulnerabilitiesByDigest) | **Get** /images/{image_digest}/vuln/{vtype} | Get vulnerabilities by type



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
    openapiclient "./openapi"
)

func main() {
    imageDigest := "imageDigest_example" // string | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImagesApi.GetImageAncestors(context.Background(), imageDigest).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.GetImageAncestors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetImageAncestors`: []ImageAncestor
    fmt.Fprintf(os.Stdout, "Response from `ImagesApi.GetImageAncestors`: %v\n", resp)
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


## GetImagePolicyCheckByDigest

> []interface{} GetImagePolicyCheckByDigest(ctx, imageDigest).Tag(tag).PolicyId(policyId).Detail(detail).History(history).Interactive(interactive).BaseDigest(baseDigest).XAnchoreAccount(xAnchoreAccount).Execute()

Check policy evaluation status for image



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
    tag := "tag_example" // string | 
    policyId := "policyId_example" // string |  (optional)
    detail := true // bool |  (optional) (default to true)
    history := true // bool |  (optional) (default to false)
    interactive := true // bool |  (optional) (default to false)
    baseDigest := "baseDigest_example" // string | Digest of a base image. If specified the evaluation will indicate results inherited from the base image (optional)
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImagesApi.GetImagePolicyCheckByDigest(context.Background(), imageDigest).Tag(tag).PolicyId(policyId).Detail(detail).History(history).Interactive(interactive).BaseDigest(baseDigest).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.GetImagePolicyCheckByDigest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetImagePolicyCheckByDigest`: []interface{}
    fmt.Fprintf(os.Stdout, "Response from `ImagesApi.GetImagePolicyCheckByDigest`: %v\n", resp)
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
 **baseDigest** | **string** | Digest of a base image. If specified the evaluation will indicate results inherited from the base image | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageVulnerabilitiesByDigest

> EnterpriseVulnerabilityResponse GetImageVulnerabilitiesByDigest(ctx, imageDigest, vtype).ForceRefresh(forceRefresh).VendorOnly(vendorOnly).BaseDigest(baseDigest).XAnchoreAccount(xAnchoreAccount).Execute()

Get vulnerabilities by type

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
    vtype := "vtype_example" // string | 
    forceRefresh := true // bool |  (optional) (default to false)
    vendorOnly := true // bool | Filter results to include only vulnerabilities that are not marked as invalid by upstream OS vendor data. When set to true, it will filter out all vulnerabilities where `will_not_fix` is False. If false all vulnerabilities are returned regardless of `will_not_fix` (optional) (default to true)
    baseDigest := "baseDigest_example" // string | Digest of a base image. If specified the vulnerabilities will indicate inheritance from the base image (optional)
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImagesApi.GetImageVulnerabilitiesByDigest(context.Background(), imageDigest, vtype).ForceRefresh(forceRefresh).VendorOnly(vendorOnly).BaseDigest(baseDigest).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.GetImageVulnerabilitiesByDigest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetImageVulnerabilitiesByDigest`: EnterpriseVulnerabilityResponse
    fmt.Fprintf(os.Stdout, "Response from `ImagesApi.GetImageVulnerabilitiesByDigest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string** |  | 
**vtype** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageVulnerabilitiesByDigestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **forceRefresh** | **bool** |  | [default to false]
 **vendorOnly** | **bool** | Filter results to include only vulnerabilities that are not marked as invalid by upstream OS vendor data. When set to true, it will filter out all vulnerabilities where &#x60;will_not_fix&#x60; is False. If false all vulnerabilities are returned regardless of &#x60;will_not_fix&#x60; | [default to true]
 **baseDigest** | **string** | Digest of a base image. If specified the vulnerabilities will indicate inheritance from the base image | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**EnterpriseVulnerabilityResponse**](EnterpriseVulnerabilityResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

