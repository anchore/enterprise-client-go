# \OnetimeAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OnetimeEvaluation**](OnetimeAPI.md#OnetimeEvaluation) | **Post** /onetime_evaluation | Return a onetime evaluation of the provided data.



## OnetimeEvaluation

> OnetimeEvaluationResponse OnetimeEvaluation(ctx).Tag(tag).OnetimeEvaluationRequest(onetimeEvaluationRequest).XAnchoreAccount(xAnchoreAccount).PolicyId(policyId).Execute()

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
	tag := "tag_example" // string | 
	onetimeEvaluationRequest := *openapiclient.NewOnetimeEvaluationRequest() // OnetimeEvaluationRequest | 
	xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)
	policyId := "policyId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnetimeAPI.OnetimeEvaluation(context.Background()).Tag(tag).OnetimeEvaluationRequest(onetimeEvaluationRequest).XAnchoreAccount(xAnchoreAccount).PolicyId(policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnetimeAPI.OnetimeEvaluation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OnetimeEvaluation`: OnetimeEvaluationResponse
	fmt.Fprintf(os.Stdout, "Response from `OnetimeAPI.OnetimeEvaluation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnetimeEvaluationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tag** | **string** |  | 
 **onetimeEvaluationRequest** | [**OnetimeEvaluationRequest**](OnetimeEvaluationRequest.md) |  | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 
 **policyId** | **string** |  | 

### Return type

[**OnetimeEvaluationResponse**](OnetimeEvaluationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

