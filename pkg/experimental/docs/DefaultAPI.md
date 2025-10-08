# \DefaultAPI

All URIs are relative to */exp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePackageLicenseOverrides**](DefaultAPI.md#CreatePackageLicenseOverrides) | **Post** /system/package-overrides/licenses | Create new license overrides for a package
[**GetPackageLicenseOverrides**](DefaultAPI.md#GetPackageLicenseOverrides) | **Get** /system/package-overrides/licenses | List all license overrides by package purl
[**UpdatePackageLicenseOverrides**](DefaultAPI.md#UpdatePackageLicenseOverrides) | **Put** /system/package-overrides/licenses | Update existing license overrides



## CreatePackageLicenseOverrides

> PackageLicenseOverride CreatePackageLicenseOverrides(ctx).PackageLicenseOverride(packageLicenseOverride).Execute()

Create new license overrides for a package

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
	packageLicenseOverride := *openapiclient.NewPackageLicenseOverride("Purl_example", []openapiclient.LicenseOverride{*openapiclient.NewLicenseOverride("LicenseId_example")}) // PackageLicenseOverride | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreatePackageLicenseOverrides(context.Background()).PackageLicenseOverride(packageLicenseOverride).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreatePackageLicenseOverrides``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePackageLicenseOverrides`: PackageLicenseOverride
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreatePackageLicenseOverrides`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePackageLicenseOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **packageLicenseOverride** | [**PackageLicenseOverride**](PackageLicenseOverride.md) |  | 

### Return type

[**PackageLicenseOverride**](PackageLicenseOverride.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPackageLicenseOverrides

> GetPackageLicenseOverrides200Response GetPackageLicenseOverrides(ctx).Purl(purl).Execute()

List all license overrides by package purl

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
	purl := "purl_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetPackageLicenseOverrides(context.Background()).Purl(purl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetPackageLicenseOverrides``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPackageLicenseOverrides`: GetPackageLicenseOverrides200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetPackageLicenseOverrides`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPackageLicenseOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **purl** | **string** |  | 

### Return type

[**GetPackageLicenseOverrides200Response**](GetPackageLicenseOverrides200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePackageLicenseOverrides

> PackageLicenseOverride UpdatePackageLicenseOverrides(ctx).PackageLicenseOverride(packageLicenseOverride).Execute()

Update existing license overrides

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
	packageLicenseOverride := *openapiclient.NewPackageLicenseOverride("Purl_example", []openapiclient.LicenseOverride{*openapiclient.NewLicenseOverride("LicenseId_example")}) // PackageLicenseOverride | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdatePackageLicenseOverrides(context.Background()).PackageLicenseOverride(packageLicenseOverride).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdatePackageLicenseOverrides``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePackageLicenseOverrides`: PackageLicenseOverride
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdatePackageLicenseOverrides`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePackageLicenseOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **packageLicenseOverride** | [**PackageLicenseOverride**](PackageLicenseOverride.md) |  | 

### Return type

[**PackageLicenseOverride**](PackageLicenseOverride.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

