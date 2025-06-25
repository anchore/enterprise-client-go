# \OnetimeAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StatelessScan**](OnetimeAPI.md#StatelessScan) | **Post** /scan | Return a onetime evaluation of the provided data.



## StatelessScan

> ScanResponse StatelessScan(ctx).Tag(tag).ScanRequest(scanRequest).XAnchoreAccount(xAnchoreAccount).PolicyId(policyId).Execute()

Return a onetime evaluation of the provided data.



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
	tag := "tag_example" // string | A valid docker tag reference (e.g. docker.io/nginx:latest) that will be used as part of the policy evaluation.
	scanRequest := *openapiclient.NewScanRequest() // ScanRequest | 
	xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)
	policyId := "policyId_example" // string | The ID of the policy used to evaluate the image (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnetimeAPI.StatelessScan(context.Background()).Tag(tag).ScanRequest(scanRequest).XAnchoreAccount(xAnchoreAccount).PolicyId(policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnetimeAPI.StatelessScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StatelessScan`: ScanResponse
	fmt.Fprintf(os.Stdout, "Response from `OnetimeAPI.StatelessScan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStatelessScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tag** | **string** | A valid docker tag reference (e.g. docker.io/nginx:latest) that will be used as part of the policy evaluation. | 
 **scanRequest** | [**ScanRequest**](ScanRequest.md) |  | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 
 **policyId** | **string** | The ID of the policy used to evaluate the image | 

### Return type

[**ScanResponse**](ScanResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

