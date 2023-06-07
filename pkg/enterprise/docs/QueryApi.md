# \QueryApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryImagesByPackage**](QueryApi.md#QueryImagesByPackage) | **Get** /query/images/by-package | List of images containing given package
[**QueryImagesByVulnerability**](QueryApi.md#QueryImagesByVulnerability) | **Get** /query/images/by-vulnerability | List images vulnerable to the specific vulnerability ID.
[**QueryVulnerabilities**](QueryApi.md#QueryVulnerabilities) | **Get** /query/vulnerabilities | Listing information about given vulnerability



## QueryImagesByPackage

> PaginatedImageList QueryImagesByPackage(ctx).Name(name).PackageType(packageType).Version(version).Page(page).Limit(limit).XAnchoreAccount(xAnchoreAccount).Execute()

List of images containing given package



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
    name := "name_example" // string | Name of package to search for (e.g. sed)
    packageType := "packageType_example" // string | Type of package to filter on (e.g. dpkg) (optional)
    version := "version_example" // string | Version of named package to filter on (e.g. 4.4-1) (optional)
    page := "page_example" // string | The page of results to fetch. Pages start at 1 (optional)
    limit := int32(56) // int32 | Limit the number of records for the requested page. If omitted or set to 0, return all results in a single page (optional)
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QueryApi.QueryImagesByPackage(context.Background()).Name(name).PackageType(packageType).Version(version).Page(page).Limit(limit).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryApi.QueryImagesByPackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryImagesByPackage`: PaginatedImageList
    fmt.Fprintf(os.Stdout, "Response from `QueryApi.QueryImagesByPackage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryImagesByPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Name of package to search for (e.g. sed) | 
 **packageType** | **string** | Type of package to filter on (e.g. dpkg) | 
 **version** | **string** | Version of named package to filter on (e.g. 4.4-1) | 
 **page** | **string** | The page of results to fetch. Pages start at 1 | 
 **limit** | **int32** | Limit the number of records for the requested page. If omitted or set to 0, return all results in a single page | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**PaginatedImageList**](PaginatedImageList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryImagesByVulnerability

> PaginatedVulnerableImageList QueryImagesByVulnerability(ctx).VulnerabilityId(vulnerabilityId).Namespace(namespace).AffectedPackage(affectedPackage).Severity(severity).VendorOnly(vendorOnly).Page(page).Limit(limit).XAnchoreAccount(xAnchoreAccount).Execute()

List images vulnerable to the specific vulnerability ID.



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
    vulnerabilityId := "vulnerabilityId_example" // string | The ID of the vulnerability to search for within all images stored in anchore-engine (e.g. CVE-1999-0001)
    namespace := "namespace_example" // string | Filter results to images within the given vulnerability namespace (e.g. debian:8, ubuntu:14.04) (optional)
    affectedPackage := "affectedPackage_example" // string | Filter results to images with vulnerable packages with the given package name (e.g. libssl) (optional)
    severity := "severity_example" // string | Filter results to vulnerable package/vulnerability with the given severity (optional)
    vendorOnly := true // bool | Filter results to include only vulnerabilities that are not marked as invalid by upstream OS vendor data (optional) (default to true)
    page := int32(56) // int32 | The page of results to fetch. Pages start at 1 (optional)
    limit := int32(56) // int32 | Limit the number of records for the requested page. If omitted or set to 0, return all results in a single page (optional)
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QueryApi.QueryImagesByVulnerability(context.Background()).VulnerabilityId(vulnerabilityId).Namespace(namespace).AffectedPackage(affectedPackage).Severity(severity).VendorOnly(vendorOnly).Page(page).Limit(limit).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryApi.QueryImagesByVulnerability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryImagesByVulnerability`: PaginatedVulnerableImageList
    fmt.Fprintf(os.Stdout, "Response from `QueryApi.QueryImagesByVulnerability`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryImagesByVulnerabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vulnerabilityId** | **string** | The ID of the vulnerability to search for within all images stored in anchore-engine (e.g. CVE-1999-0001) | 
 **namespace** | **string** | Filter results to images within the given vulnerability namespace (e.g. debian:8, ubuntu:14.04) | 
 **affectedPackage** | **string** | Filter results to images with vulnerable packages with the given package name (e.g. libssl) | 
 **severity** | **string** | Filter results to vulnerable package/vulnerability with the given severity | 
 **vendorOnly** | **bool** | Filter results to include only vulnerabilities that are not marked as invalid by upstream OS vendor data | [default to true]
 **page** | **int32** | The page of results to fetch. Pages start at 1 | 
 **limit** | **int32** | Limit the number of records for the requested page. If omitted or set to 0, return all results in a single page | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**PaginatedVulnerableImageList**](PaginatedVulnerableImageList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryVulnerabilities

> PaginatedVulnerabilityList QueryVulnerabilities(ctx).Id(id).AffectedPackage(affectedPackage).AffectedPackageVersion(affectedPackageVersion).Page(page).Limit(limit).Namespace(namespace).Execute()

Listing information about given vulnerability



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
    id := []string{"Inner_example"} // []string | The ID of the vulnerability (e.g. CVE-1999-0001)
    affectedPackage := "affectedPackage_example" // string | Filter results by specified package name (e.g. sed) (optional)
    affectedPackageVersion := "affectedPackageVersion_example" // string | Filter results by specified package version (e.g. 4.4-1) (optional)
    page := "page_example" // string | The page of results to fetch. Pages start at 1 (optional) (default to "1")
    limit := int32(56) // int32 | Limit the number of records for the requested page. If omitted or set to 0, return all results in a single page (optional)
    namespace := []string{"Inner_example"} // []string | Namespace(s) to filter vulnerability records by (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QueryApi.QueryVulnerabilities(context.Background()).Id(id).AffectedPackage(affectedPackage).AffectedPackageVersion(affectedPackageVersion).Page(page).Limit(limit).Namespace(namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryApi.QueryVulnerabilities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryVulnerabilities`: PaginatedVulnerabilityList
    fmt.Fprintf(os.Stdout, "Response from `QueryApi.QueryVulnerabilities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryVulnerabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** | The ID of the vulnerability (e.g. CVE-1999-0001) | 
 **affectedPackage** | **string** | Filter results by specified package name (e.g. sed) | 
 **affectedPackageVersion** | **string** | Filter results by specified package version (e.g. 4.4-1) | 
 **page** | **string** | The page of results to fetch. Pages start at 1 | [default to &quot;1&quot;]
 **limit** | **int32** | Limit the number of records for the requested page. If omitted or set to 0, return all results in a single page | 
 **namespace** | **[]string** | Namespace(s) to filter vulnerability records by | 

### Return type

[**PaginatedVulnerabilityList**](PaginatedVulnerabilityList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

