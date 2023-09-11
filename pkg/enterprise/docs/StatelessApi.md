# \StatelessApi

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VulnerabilityScanSbom**](StatelessApi.md#VulnerabilityScanSbom) | **Post** /vulnerability-scan | Return a vulnerability scan for the uploaded SBOM without storing the SBOM and without any side-effects in the system.



## VulnerabilityScanSbom

> SBOMVulnerabilitiesResponse VulnerabilityScanSbom(ctx).Sbom(sbom).XAnchoreAccount(xAnchoreAccount).Execute()

Return a vulnerability scan for the uploaded SBOM without storing the SBOM and without any side-effects in the system.



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
    sbom := interface{}(Object) // interface{} | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StatelessApi.VulnerabilityScanSbom(context.Background()).Sbom(sbom).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatelessApi.VulnerabilityScanSbom``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VulnerabilityScanSbom`: SBOMVulnerabilitiesResponse
    fmt.Fprintf(os.Stdout, "Response from `StatelessApi.VulnerabilityScanSbom`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVulnerabilityScanSbomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sbom** | **interface{}** |  | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**SBOMVulnerabilitiesResponse**](SBOMVulnerabilitiesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

