# \QueryAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryImagesByPackage**](QueryAPI.md#QueryImagesByPackage) | **Get** /query/images/by-package | List of images containing given package
[**QueryVulnerabilities**](QueryAPI.md#QueryVulnerabilities) | **Get** /query/vulnerabilities | Listing information about given vulnerability



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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    name := "name_example" // string | Name of package to search for (e.g. sed)
    packageType := "packageType_example" // string | Type of package to filter on (e.g. dpkg) (optional)
    version := "version_example" // string | Version of named package to filter on (e.g. 4.4-1) (optional)
    page := "page_example" // string | The page of results to fetch. Pages start at 1 (optional)
    limit := int32(56) // int32 | Limit the number of records for the requested page. If omitted or set to 0, return all results in a single page (optional)
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueryAPI.QueryImagesByPackage(context.Background()).Name(name).PackageType(packageType).Version(version).Page(page).Limit(limit).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryAPI.QueryImagesByPackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryImagesByPackage`: PaginatedImageList
    fmt.Fprintf(os.Stdout, "Response from `QueryAPI.QueryImagesByPackage`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    id := []string{"Inner_example"} // []string | The ID of the vulnerability (e.g. CVE-1999-0001)
    affectedPackage := "affectedPackage_example" // string | Filter results by specified package name (e.g. sed) (optional)
    affectedPackageVersion := "affectedPackageVersion_example" // string | Filter results by specified package version (e.g. 4.4-1) (optional)
    page := "page_example" // string | The page of results to fetch. Pages start at 1 (optional) (default to "1")
    limit := int32(56) // int32 | Limit the number of records for the requested page. If omitted or set to 0, return all results in a single page (optional)
    namespace := []string{"Inner_example"} // []string | Namespace(s) to filter vulnerability records by (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueryAPI.QueryVulnerabilities(context.Background()).Id(id).AffectedPackage(affectedPackage).AffectedPackageVersion(affectedPackageVersion).Page(page).Limit(limit).Namespace(namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryAPI.QueryVulnerabilities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryVulnerabilities`: PaginatedVulnerabilityList
    fmt.Fprintf(os.Stdout, "Response from `QueryAPI.QueryVulnerabilities`: %v\n", resp)
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

