# \StatelessAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VulnerabilityScanSbom**](StatelessAPI.md#VulnerabilityScanSbom) | **Post** /vulnerability-scan | Return a vulnerability scan for the uploaded SBOM without storing the SBOM and without any side-effects in the system.



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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    sbom := interface{}{ ... } // interface{} | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatelessAPI.VulnerabilityScanSbom(context.Background()).Sbom(sbom).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatelessAPI.VulnerabilityScanSbom``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VulnerabilityScanSbom`: SBOMVulnerabilitiesResponse
    fmt.Fprintf(os.Stdout, "Response from `StatelessAPI.VulnerabilityScanSbom`: %v\n", resp)
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

