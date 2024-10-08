# \ArchivesAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveImageAnalysis**](ArchivesAPI.md#ArchiveImageAnalysis) | **Post** /archives/images | 
[**CreateAnalysisArchiveRule**](ArchivesAPI.md#CreateAnalysisArchiveRule) | **Post** /archives/rules | 
[**DeleteAnalysisArchiveRule**](ArchivesAPI.md#DeleteAnalysisArchiveRule) | **Delete** /archives/rules/{rule_id} | 
[**DeleteArchivedAnalysis**](ArchivesAPI.md#DeleteArchivedAnalysis) | **Delete** /archives/images/{image_digest} | 
[**GetAnalysisArchiveRule**](ArchivesAPI.md#GetAnalysisArchiveRule) | **Get** /archives/rules/{rule_id} | 
[**GetArchivedAnalysis**](ArchivesAPI.md#GetArchivedAnalysis) | **Get** /archives/images/{image_digest} | 
[**ListAnalysisArchive**](ArchivesAPI.md#ListAnalysisArchive) | **Get** /archives/images | 
[**ListAnalysisArchiveRules**](ArchivesAPI.md#ListAnalysisArchiveRules) | **Get** /archives/rules | 
[**ListArchives**](ArchivesAPI.md#ListArchives) | **Get** /archives | 



## ArchiveImageAnalysis

> []AnalysisArchiveAddResult ArchiveImageAnalysis(ctx).ImageReferences(imageReferences).Execute()



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
	imageReferences := []string{"Property_example"} // []string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArchivesAPI.ArchiveImageAnalysis(context.Background()).ImageReferences(imageReferences).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArchivesAPI.ArchiveImageAnalysis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArchiveImageAnalysis`: []AnalysisArchiveAddResult
	fmt.Fprintf(os.Stdout, "Response from `ArchivesAPI.ArchiveImageAnalysis`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveImageAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imageReferences** | **[]string** |  | 

### Return type

[**[]AnalysisArchiveAddResult**](AnalysisArchiveAddResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAnalysisArchiveRule

> AnalysisArchiveTransitionRule CreateAnalysisArchiveRule(ctx).Rule(rule).Execute()



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
	rule := *openapiclient.NewAnalysisArchiveTransitionRule("Transition_example") // AnalysisArchiveTransitionRule | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArchivesAPI.CreateAnalysisArchiveRule(context.Background()).Rule(rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArchivesAPI.CreateAnalysisArchiveRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAnalysisArchiveRule`: AnalysisArchiveTransitionRule
	fmt.Fprintf(os.Stdout, "Response from `ArchivesAPI.CreateAnalysisArchiveRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAnalysisArchiveRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rule** | [**AnalysisArchiveTransitionRule**](AnalysisArchiveTransitionRule.md) |  | 

### Return type

[**AnalysisArchiveTransitionRule**](AnalysisArchiveTransitionRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAnalysisArchiveRule

> DeleteAnalysisArchiveRule(ctx, ruleId).Execute()



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
	ruleId := "ruleId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ArchivesAPI.DeleteAnalysisArchiveRule(context.Background(), ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArchivesAPI.DeleteAnalysisArchiveRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAnalysisArchiveRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteArchivedAnalysis

> DeleteArchivedAnalysis(ctx, imageDigest).Force(force).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ArchivesAPI.DeleteArchivedAnalysis(context.Background(), imageDigest).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArchivesAPI.DeleteArchivedAnalysis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteArchivedAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalysisArchiveRule

> AnalysisArchiveTransitionRule GetAnalysisArchiveRule(ctx, ruleId).Execute()



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
	ruleId := "ruleId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArchivesAPI.GetAnalysisArchiveRule(context.Background(), ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArchivesAPI.GetAnalysisArchiveRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnalysisArchiveRule`: AnalysisArchiveTransitionRule
	fmt.Fprintf(os.Stdout, "Response from `ArchivesAPI.GetAnalysisArchiveRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalysisArchiveRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AnalysisArchiveTransitionRule**](AnalysisArchiveTransitionRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArchivedAnalysis

> ArchivedAnalysis GetArchivedAnalysis(ctx, imageDigest).Execute()





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
	imageDigest := "imageDigest_example" // string | The image digest to identify the image analysis

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArchivesAPI.GetArchivedAnalysis(context.Background(), imageDigest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArchivesAPI.GetArchivedAnalysis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetArchivedAnalysis`: ArchivedAnalysis
	fmt.Fprintf(os.Stdout, "Response from `ArchivesAPI.GetArchivedAnalysis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string** | The image digest to identify the image analysis | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArchivedAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ArchivedAnalysis**](ArchivedAnalysis.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAnalysisArchive

> []ArchivedAnalysis ListAnalysisArchive(ctx).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArchivesAPI.ListAnalysisArchive(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArchivesAPI.ListAnalysisArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAnalysisArchive`: []ArchivedAnalysis
	fmt.Fprintf(os.Stdout, "Response from `ArchivesAPI.ListAnalysisArchive`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAnalysisArchiveRequest struct via the builder pattern


### Return type

[**[]ArchivedAnalysis**](ArchivedAnalysis.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAnalysisArchiveRules

> []AnalysisArchiveTransitionRule ListAnalysisArchiveRules(ctx).SystemGlobal(systemGlobal).Execute()



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
	systemGlobal := true // bool | If true include system global rules (owned by admin) even for non-admin users. Defaults to true if not set. Can be set to false to exclude globals (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArchivesAPI.ListAnalysisArchiveRules(context.Background()).SystemGlobal(systemGlobal).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArchivesAPI.ListAnalysisArchiveRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAnalysisArchiveRules`: []AnalysisArchiveTransitionRule
	fmt.Fprintf(os.Stdout, "Response from `ArchivesAPI.ListAnalysisArchiveRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAnalysisArchiveRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemGlobal** | **bool** | If true include system global rules (owned by admin) even for non-admin users. Defaults to true if not set. Can be set to false to exclude globals | 

### Return type

[**[]AnalysisArchiveTransitionRule**](AnalysisArchiveTransitionRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListArchives

> ArchiveSummary ListArchives(ctx).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArchivesAPI.ListArchives(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArchivesAPI.ListArchives``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListArchives`: ArchiveSummary
	fmt.Fprintf(os.Stdout, "Response from `ArchivesAPI.ListArchives`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListArchivesRequest struct via the builder pattern


### Return type

[**ArchiveSummary**](ArchiveSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

