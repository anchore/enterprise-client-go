# \SBOMManagementAPI

All URIs are relative to */exp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSbom**](SBOMManagementAPI.md#AddSbom) | **Post** /sboms | Add an SBOM to your SBOM library
[**BulkModifySbomGroupMemberships**](SBOMManagementAPI.md#BulkModifySbomGroupMemberships) | **Patch** /sbom-groups/memberships | Modify the set of SBOMs belonging to a set of Groups to simplify bulk updates. To modify the set of SBOMs belonging to a single group, use /sbom-groups/{group_uuid}/memberships. Returns a list of Groups and the SBOM UUIDs belonging to each Group after the requested changes. Limited to a maximum of 10,000 additions and 10,000 removals per group, except when using remove_all.
[**CreateSbomGroup**](SBOMManagementAPI.md#CreateSbomGroup) | **Post** /sbom-groups | Create an SBOM Group
[**DeleteSbom**](SBOMManagementAPI.md#DeleteSbom) | **Delete** /sboms/{sbom_uuid} | Delete an SBOM and all associated document revisions
[**DeleteSbomGroup**](SBOMManagementAPI.md#DeleteSbomGroup) | **Delete** /sbom-groups/{group_uuid} | Delete an SBOM Group
[**DeleteSbomRevision**](SBOMManagementAPI.md#DeleteSbomRevision) | **Delete** /sboms/{sbom_uuid}/revisions/{revision_uuid} | Delete a specific revision of a given SBOM
[**GetLatestSbomByUuid**](SBOMManagementAPI.md#GetLatestSbomByUuid) | **Get** /sboms/{sbom_uuid} | Get the metadata of the latest revision of a given SBOM
[**GetLatestSbomPackagesByUuid**](SBOMManagementAPI.md#GetLatestSbomPackagesByUuid) | **Get** /sboms/{sbom_uuid}/packages | Return the package information for the latest revision of a given SBOM, this is sorted alphabetically by default
[**GetSbomFile**](SBOMManagementAPI.md#GetSbomFile) | **Get** /sboms/{sbom_uuid}/file | Download the content of the latest revision of a given SBOM
[**GetSbomGroupByUuid**](SBOMManagementAPI.md#GetSbomGroupByUuid) | **Get** /sbom-groups/{group_uuid} | Get the details of a specific SBOM Group
[**GetSbomGroupVulnerabilitiesCsv**](SBOMManagementAPI.md#GetSbomGroupVulnerabilitiesCsv) | **Get** /sbom-groups/{group_uuid}/vulnerabilities/detail-csv | Download a csv file containing a list of package vulnerabilities found within the SBOMs belonging to the given SBOM Group during the most recent scans.
[**GetSbomGroupZip**](SBOMManagementAPI.md#GetSbomGroupZip) | **Get** /sbom-groups/{group_uuid}/sbom-bundle | Download a zip file containing all SBOMs from a given group
[**GetSbomRevision**](SBOMManagementAPI.md#GetSbomRevision) | **Get** /sboms/{sbom_uuid}/revisions/{revision_uuid} | Get the metadata for a specific revision of a given SBOM
[**GetSbomRevisionFile**](SBOMManagementAPI.md#GetSbomRevisionFile) | **Get** /sboms/{sbom_uuid}/revisions/{revision_uuid}/file | Download the content of a specific revision of a given SBOM revision
[**GetSbomVulnerabilitiesCsv**](SBOMManagementAPI.md#GetSbomVulnerabilitiesCsv) | **Get** /sboms/{sbom_uuid}/vulnerabilities/detail-csv | Download a csv file containing a list of package  vulnerabilities found within a given SBOM during the most recent scan.
[**ListSbomGroupVulnerabilities**](SBOMManagementAPI.md#ListSbomGroupVulnerabilities) | **Get** /sbom-groups/{group_uuid}/vulnerabilities | Get a list of vulnerabilities found within the SBOMs belonging to the given SBOM Group during the most recent scans.  Returns summary data of the complete result set and a list of the first 100 results.
[**ListSbomGroups**](SBOMManagementAPI.md#ListSbomGroups) | **Get** /sbom-groups | Get a list of SBOM Groups belonging to your account
[**ListSbomVulnerabilities**](SBOMManagementAPI.md#ListSbomVulnerabilities) | **Get** /sboms/{sbom_uuid}/vulnerabilities | Get a list of vulnerabilities found within a given SBOM during the most recent scan. Returns summary data of the complete result set and a list of the first 100 results.
[**ListSboms**](SBOMManagementAPI.md#ListSboms) | **Get** /sboms | Get a list of all SBOMs belonging to your account
[**ModifySbomGroupMemberships**](SBOMManagementAPI.md#ModifySbomGroupMemberships) | **Patch** /sbom-groups/{group_uuid}/memberships | Modify the set of SBOMs belonging to a given Group. Returns a list of SBOM UUIDs belonging to the Group after the requested changes. Limited to a maximum of 10,000 additions and 10,000 removals in a single request, except when using remove_all.
[**PatchSbom**](SBOMManagementAPI.md#PatchSbom) | **Patch** /sboms/{sbom_uuid} | Update the properties of the given SBOM. Existing annotations and group memberships will be updated, not be replaced.
[**PutSbom**](SBOMManagementAPI.md#PutSbom) | **Put** /sboms/{sbom_uuid} | Update the properties of the given SBOM. Replaces all existing annotations and group memberships.
[**UpdateSbomGroup**](SBOMManagementAPI.md#UpdateSbomGroup) | **Put** /sbom-groups/{group_uuid} | Update the details of a specific SBOM Group



## AddSbom

> SBOMDetail AddSbom(ctx).Name(name).Version(version).File(file).XAnchoreAccount(xAnchoreAccount).Annotations(annotations).Groups(groups).Type_(type_).Execute()

Add an SBOM to your SBOM library

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
	name := "name_example" // string | 
	version := "version_example" // string | 
	file := os.NewFile(1234, "some_file") // *os.File | A valid SBOM file
	xAnchoreAccount := "xAnchoreAccount_example" // string | The account name used for the resource scope of this request as user permissions allow (optional)
	annotations := "annotations_example" // string | A json string representing a dict of string type key value pairs (optional)
	groups := []string{"Inner_example"} // []string | A comma separated list of Group UUIDs (optional)
	type_ := "type__example" // string | The type of SBOM being uploaded. (The values are just placeholders for now) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SBOMManagementAPI.AddSbom(context.Background()).Name(name).Version(version).File(file).XAnchoreAccount(xAnchoreAccount).Annotations(annotations).Groups(groups).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SBOMManagementAPI.AddSbom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddSbom`: SBOMDetail
	fmt.Fprintf(os.Stdout, "Response from `SBOMManagementAPI.AddSbom`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddSbomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **version** | **string** |  | 
 **file** | ***os.File** | A valid SBOM file | 
 **xAnchoreAccount** | **string** | The account name used for the resource scope of this request as user permissions allow | 
 **annotations** | **string** | A json string representing a dict of string type key value pairs | 
 **groups** | **[]string** | A comma separated list of Group UUIDs | 
 **type_** | **string** | The type of SBOM being uploaded. (The values are just placeholders for now) | 

### Return type

[**SBOMDetail**](SBOMDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkModifySbomGroupMemberships

> map[string]SBOMGroupMembershipResponse BulkModifySbomGroupMemberships(ctx).SBOMBulkGroupMembershipRequest(sBOMBulkGroupMembershipRequest).XAnchoreAccount(xAnchoreAccount).Execute()

Modify the set of SBOMs belonging to a set of Groups to simplify bulk updates. To modify the set of SBOMs belonging to a single group, use /sbom-groups/{group_uuid}/memberships. Returns a list of Groups and the SBOM UUIDs belonging to each Group after the requested changes. Limited to a maximum of 10,000 additions and 10,000 removals per group, except when using remove_all.

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
	sBOMBulkGroupMembershipRequest := *openapiclient.NewSBOMBulkGroupMembershipRequest([]openapiclient.SBOMBulkGroupMembershipRequestGroupsInner{*openapiclient.NewSBOMBulkGroupMembershipRequestGroupsInner("GroupUuid_example")}) // SBOMBulkGroupMembershipRequest | 
	xAnchoreAccount := "xAnchoreAccount_example" // string | The account name used for the resource scope of this request as user permissions allow (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SBOMManagementAPI.BulkModifySbomGroupMemberships(context.Background()).SBOMBulkGroupMembershipRequest(sBOMBulkGroupMembershipRequest).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SBOMManagementAPI.BulkModifySbomGroupMemberships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkModifySbomGroupMemberships`: map[string]SBOMGroupMembershipResponse
	fmt.Fprintf(os.Stdout, "Response from `SBOMManagementAPI.BulkModifySbomGroupMemberships`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkModifySbomGroupMembershipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sBOMBulkGroupMembershipRequest** | [**SBOMBulkGroupMembershipRequest**](SBOMBulkGroupMembershipRequest.md) |  | 
 **xAnchoreAccount** | **string** | The account name used for the resource scope of this request as user permissions allow | 

### Return type

[**map[string]SBOMGroupMembershipResponse**](SBOMGroupMembershipResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSbomGroup

> SBOMGroupDetail CreateSbomGroup(ctx).CreateSbomGroupRequest(createSbomGroupRequest).XAnchoreAccount(xAnchoreAccount).Execute()

Create an SBOM Group

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
	createSbomGroupRequest := *openapiclient.NewCreateSbomGroupRequest("Name_example", "Version_example") // CreateSbomGroupRequest | Properties of the Group to be created
	xAnchoreAccount := "xAnchoreAccount_example" // string | The account name used for the resource scope of this request as user permissions allow (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SBOMManagementAPI.CreateSbomGroup(context.Background()).CreateSbomGroupRequest(createSbomGroupRequest).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SBOMManagementAPI.CreateSbomGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSbomGroup`: SBOMGroupDetail
	fmt.Fprintf(os.Stdout, "Response from `SBOMManagementAPI.CreateSbomGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSbomGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSbomGroupRequest** | [**CreateSbomGroupRequest**](CreateSbomGroupRequest.md) | Properties of the Group to be created | 
 **xAnchoreAccount** | **string** | The account name used for the resource scope of this request as user permissions allow | 

### Return type

[**SBOMGroupDetail**](SBOMGroupDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSbom

> DeleteSbom(ctx, sbomUuid).XAnchoreAccount(xAnchoreAccount).Execute()

Delete an SBOM and all associated document revisions

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
	sbomUuid := "sbomUuid_example" // string | 
	xAnchoreAccount := "xAnchoreAccount_example" // string | The account name used for the resource scope of this request as user permissions allow (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SBOMManagementAPI.DeleteSbom(context.Background(), sbomUuid).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SBOMManagementAPI.DeleteSbom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sbomUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSbomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | The account name used for the resource scope of this request as user permissions allow | 

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


## DeleteSbomGroup

> DeleteSbomGroup(ctx, groupUuid).XAnchoreAccount(xAnchoreAccount).Execute()

Delete an SBOM Group

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
	groupUuid := "groupUuid_example" // string | 
	xAnchoreAccount := "xAnchoreAccount_example" // string | The account name used for the resource scope of this request as user permissions allow (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SBOMManagementAPI.DeleteSbomGroup(context.Background(), groupUuid).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SBOMManagementAPI.DeleteSbomGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSbomGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | The account name used for the resource scope of this request as user permissions allow | 

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


## DeleteSbomRevision

> DeleteSbomRevision(ctx, sbomUuid, revisionUuid).XAnchoreAccount(xAnchoreAccount).Execute()

Delete a specific revision of a given SBOM

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
	sbomUuid := "sbomUuid_example" // string | 
	revisionUuid := "revisionUuid_example" // string | 
	xAnchoreAccount := "xAnchoreAccount_example" // string | The account name used for the resource scope of this request as user permissions allow (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SBOMManagementAPI.DeleteSbomRevision(context.Background(), sbomUuid, revisionUuid).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SBOMManagementAPI.DeleteSbomRevision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sbomUuid** | **string** |  | 
**revisionUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSbomRevisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAnchoreAccount** | **string** | The account name used for the resource scope of this request as user permissions allow | 

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


## GetLatestSbomByUuid

> SBOMDetail GetLatestSbomByUuid(ctx, sbomUuid).XAnchoreAccount(xAnchoreAccount).Execute()

Get the metadata of the latest revision of a given SBOM

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
	sbomUuid := "sbomUuid_example" // string | 
	xAnchoreAccount := "xAnchoreAccount_example" // string | The account name used for the resource scope of this request as user permissions allow (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SBOMManagementAPI.GetLatestSbomByUuid(context.Background(), sbomUuid).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SBOMManagementAPI.GetLatestSbomByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLatestSbomByUuid`: SBOMDetail
	fmt.Fprintf(os.Stdout, "Response from `SBOMManagementAPI.GetLatestSbomByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sbomUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestSbomByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | The account name used for the resource scope of this request as user permissions allow | 

### Return type

[**SBOMDetail**](SBOMDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestSbomPackagesByUuid

> SBOMPackageDetail GetLatestSbomPackagesByUuid(ctx, sbomUuid).XAnchoreAccount(xAnchoreAccount).Limit(limit).Page(page).OrderBy(orderBy).OrderByDescending(orderByDescending).Execute()

Return the package information for the latest revision of a given SBOM, this is sorted alphabetically by default

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
	sbomUuid := "sbomUuid_example" // string | 
	xAnchoreAccount := "xAnchoreAccount_example" // string | The account name used for the resource scope of this request as user permissions allow (optional)
	limit := int32(56) // int32 | Maximum number of rows to return per page (optional) (default to 1000)
	page := int32(56) // int32 | Page number to return (optional) (default to 1)
	orderBy := "orderBy_example" // string | Field name to order by, ascending by default (optional)
	orderByDescending := true // bool | Configures the sort order of the specified order_by value to be descending (true) instead of ascending (false) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SBOMManagementAPI.GetLatestSbomPackagesByUuid(context.Background(), sbomUuid).XAnchoreAccount(xAnchoreAccount).Limit(limit).Page(page).OrderBy(orderBy).OrderByDescending(orderByDescending).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SBOMManagementAPI.GetLatestSbomPackagesByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLatestSbomPackagesByUuid`: SBOMPackageDetail
	fmt.Fprintf(os.Stdout, "Response from `SBOMManagementAPI.GetLatestSbomPackagesByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sbomUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestSbomPackagesByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | The account name used for the resource scope of this request as user permissions allow | 
 **limit** | **int32** | Maximum number of rows to return per page | [default to 1000]
 **page** | **int32** | Page number to return | [default to 1]
 **orderBy** | **string** | Field name to order by, ascending by default | 
 **orderByDescending** | **bool** | Configures the sort order of the specified order_by value to be descending (true) instead of ascending (false) | 

### Return type

[**SBOMPackageDetail**](SBOMPackageDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSbomFile

> *os.File GetSbomFile(ctx, sbomUuid).XAnchoreAccount(xAnchoreAccount).Execute()

Download the content of the latest revision of a given SBOM

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
	sbomUuid := "sbomUuid_example" // string | 
	xAnchoreAccount := "xAnchoreAccount_example" // string | The account name used for the resource scope of this request as user permissions allow (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SBOMManagementAPI.GetSbomFile(context.Background(), sbomUuid).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SBOMManagementAPI.GetSbomFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSbomFile`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `SBOMManagementAPI.GetSbomFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sbomUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSbomFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | The account name used for the resource scope of this request as user permissions allow | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSbomGroupByUuid

> SBOMGroupDetail GetSbomGroupByUuid(ctx, groupUuid).XAnchoreAccount(xAnchoreAccount).Execute()

Get the details of a specific SBOM Group

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
	groupUuid := "groupUuid_example" // string | 
	xAnchoreAccount := "xAnchoreAccount_example" // string | The account name used for the resource scope of this request as user permissions allow (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SBOMManagementAPI.GetSbomGroupByUuid(context.Background(), groupUuid).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SBOMManagementAPI.GetSbomGroupByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSbomGroupByUuid`: SBOMGroupDetail
	fmt.Fprintf(os.Stdout, "Response from `SBOMManagementAPI.GetSbomGroupByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSbomGroupByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | The account name used for the resource scope of this request as user permissions allow | 

### Return type

[**SBOMGroupDetail**](SBOMGroupDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSbomGroupVulnerabilitiesCsv

> *os.File GetSbomGroupVulnerabilitiesCsv(ctx, groupUuid).XAnchoreAccount(xAnchoreAccount).AnchoreScore(anchoreScore).CvssScore(cvssScore).EpssScore(epssScore).EpssPercentile(epssPercentile).IsKev(isKev).Severity(severity).LastFoundAfter(lastFoundAfter).LastFoundBefore(lastFoundBefore).OrderBy(orderBy).OrderByDescending(orderByDescending).Execute()

Download a csv file containing a list of package vulnerabilities found within the SBOMs belonging to the given SBOM Group during the most recent scans.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	groupUuid := "groupUuid_example" // string | 
	xAnchoreAccount := "xAnchoreAccount_example" // string | The account name used for the resource scope of this request as user permissions allow (optional)
	anchoreScore := float32(3.4) // float32 | The minimum Anchore Score to be included in the results (optional)
	cvssScore := float32(3.4) // float32 | The minimum CVSS Score to be included in the results (optional)
	epssScore := float32(3.4) // float32 | The minimum EPSS Score to be included in the results (optional)
	epssPercentile := float32(3.4) // float32 | The minimum EPSS Percentile to be included in the results (optional)
	isKev := true // bool | Filter results based on presence in the CISA Known Exploitable Vulnerabilities list (optional)
	severity := "severity_example" // string | The minimum severity to be included in the results (optional)
	lastFoundAfter := time.Now() // time.Time | The earliest last_found_at date to be included in the results (optional)
	lastFoundBefore := time.Now() // time.Time | The latest last_found_at date to be included in the results (optional)
	orderBy := []string{"OrderBy_example"} // []string | List of field name(s) to order by, ascending by default (optional)
	orderByDescending := []bool{false} // []bool | Configures the sort order of each specified order_by column to be descending (true) instead of ascending (false) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SBOMManagementAPI.GetSbomGroupVulnerabilitiesCsv(context.Background(), groupUuid).XAnchoreAccount(xAnchoreAccount).AnchoreScore(anchoreScore).CvssScore(cvssScore).EpssScore(epssScore).EpssPercentile(epssPercentile).IsKev(isKev).Severity(severity).LastFoundAfter(lastFoundAfter).LastFoundBefore(lastFoundBefore).OrderBy(orderBy).OrderByDescending(orderByDescending).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SBOMManagementAPI.GetSbomGroupVulnerabilitiesCsv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSbomGroupVulnerabilitiesCsv`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `SBOMManagementAPI.GetSbomGroupVulnerabilitiesCsv`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSbomGroupVulnerabilitiesCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | The account name used for the resource scope of this request as user permissions allow | 
 **anchoreScore** | **float32** | The minimum Anchore Score to be included in the results | 
 **cvssScore** | **float32** | The minimum CVSS Score to be included in the results | 
 **epssScore** | **float32** | The minimum EPSS Score to be included in the results | 
 **epssPercentile** | **float32** | The minimum EPSS Percentile to be included in the results | 
 **isKev** | **bool** | Filter results based on presence in the CISA Known Exploitable Vulnerabilities list | 
 **severity** | **string** | The minimum severity to be included in the results | 
 **lastFoundAfter** | **time.Time** | The earliest last_found_at date to be included in the results | 
 **lastFoundBefore** | **time.Time** | The latest last_found_at date to be included in the results | 
 **orderBy** | **[]string** | List of field name(s) to order by, ascending by default | 
 **orderByDescending** | **[]bool** | Configures the sort order of each specified order_by column to be descending (true) instead of ascending (false) | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSbomGroupZip

> *os.File GetSbomGroupZip(ctx, groupUuid).XAnchoreAccount(xAnchoreAccount).Execute()

Download a zip file containing all SBOMs from a given group

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
	groupUuid := "groupUuid_example" // string | 
	xAnchoreAccount := "xAnchoreAccount_example" // string | The account name used for the resource scope of this request as user permissions allow (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SBOMManagementAPI.GetSbomGroupZip(context.Background(), groupUuid).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SBOMManagementAPI.GetSbomGroupZip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSbomGroupZip`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `SBOMManagementAPI.GetSbomGroupZip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSbomGroupZipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | The account name used for the resource scope of this request as user permissions allow | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/zip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSbomRevision

> SBOMDetail GetSbomRevision(ctx, sbomUuid, revisionUuid).XAnchoreAccount(xAnchoreAccount).Execute()

Get the metadata for a specific revision of a given SBOM

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
	sbomUuid := "sbomUuid_example" // string | 
	revisionUuid := "revisionUuid_example" // string | 
	xAnchoreAccount := "xAnchoreAccount_example" // string | The account name used for the resource scope of this request as user permissions allow (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SBOMManagementAPI.GetSbomRevision(context.Background(), sbomUuid, revisionUuid).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SBOMManagementAPI.GetSbomRevision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSbomRevision`: SBOMDetail
	fmt.Fprintf(os.Stdout, "Response from `SBOMManagementAPI.GetSbomRevision`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sbomUuid** | **string** |  | 
**revisionUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSbomRevisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAnchoreAccount** | **string** | The account name used for the resource scope of this request as user permissions allow | 

### Return type

[**SBOMDetail**](SBOMDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSbomRevisionFile

> *os.File GetSbomRevisionFile(ctx, sbomUuid, revisionUuid).XAnchoreAccount(xAnchoreAccount).Execute()

Download the content of a specific revision of a given SBOM revision

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
	sbomUuid := "sbomUuid_example" // string | 
	revisionUuid := "revisionUuid_example" // string | 
	xAnchoreAccount := "xAnchoreAccount_example" // string | The account name used for the resource scope of this request as user permissions allow (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SBOMManagementAPI.GetSbomRevisionFile(context.Background(), sbomUuid, revisionUuid).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SBOMManagementAPI.GetSbomRevisionFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSbomRevisionFile`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `SBOMManagementAPI.GetSbomRevisionFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sbomUuid** | **string** |  | 
**revisionUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSbomRevisionFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAnchoreAccount** | **string** | The account name used for the resource scope of this request as user permissions allow | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSbomVulnerabilitiesCsv

> *os.File GetSbomVulnerabilitiesCsv(ctx, sbomUuid).XAnchoreAccount(xAnchoreAccount).ForceRefresh(forceRefresh).AnchoreScore(anchoreScore).CvssScore(cvssScore).EpssScore(epssScore).EpssPercentile(epssPercentile).IsKev(isKev).Severity(severity).LastFoundAfter(lastFoundAfter).LastFoundBefore(lastFoundBefore).OrderBy(orderBy).OrderByDescending(orderByDescending).Execute()

Download a csv file containing a list of package  vulnerabilities found within a given SBOM during the most recent scan.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	sbomUuid := "sbomUuid_example" // string | 
	xAnchoreAccount := "xAnchoreAccount_example" // string | The account name used for the resource scope of this request as user permissions allow (optional)
	forceRefresh := true // bool | Force a re-scan of SBOM vulnerabilities before returning results (optional) (default to false)
	anchoreScore := float32(3.4) // float32 | The minimum Anchore Score to be included in the results (optional)
	cvssScore := float32(3.4) // float32 | The minimum CVSS Score to be included in the results (optional)
	epssScore := float32(3.4) // float32 | The minimum EPSS Score to be included in the results (optional)
	epssPercentile := float32(3.4) // float32 | The minimum EPSS Percentile to be included in the results (optional)
	isKev := true // bool | Filter results based on presence in the CISA Known Exploitable Vulnerabilities list (optional)
	severity := "severity_example" // string | The minimum severity to be included in the results (optional)
	lastFoundAfter := time.Now() // time.Time | The earliest last_found_at date to be included in the results (optional)
	lastFoundBefore := time.Now() // time.Time | The latest last_found_at date to be included in the results (optional)
	orderBy := []string{"OrderBy_example"} // []string | List of field name(s) to order by, ascending by default (optional)
	orderByDescending := []bool{false} // []bool | Configures the sort order of each specified order_by column to be descending (true) instead of ascending (false) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SBOMManagementAPI.GetSbomVulnerabilitiesCsv(context.Background(), sbomUuid).XAnchoreAccount(xAnchoreAccount).ForceRefresh(forceRefresh).AnchoreScore(anchoreScore).CvssScore(cvssScore).EpssScore(epssScore).EpssPercentile(epssPercentile).IsKev(isKev).Severity(severity).LastFoundAfter(lastFoundAfter).LastFoundBefore(lastFoundBefore).OrderBy(orderBy).OrderByDescending(orderByDescending).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SBOMManagementAPI.GetSbomVulnerabilitiesCsv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSbomVulnerabilitiesCsv`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `SBOMManagementAPI.GetSbomVulnerabilitiesCsv`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sbomUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSbomVulnerabilitiesCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | The account name used for the resource scope of this request as user permissions allow | 
 **forceRefresh** | **bool** | Force a re-scan of SBOM vulnerabilities before returning results | [default to false]
 **anchoreScore** | **float32** | The minimum Anchore Score to be included in the results | 
 **cvssScore** | **float32** | The minimum CVSS Score to be included in the results | 
 **epssScore** | **float32** | The minimum EPSS Score to be included in the results | 
 **epssPercentile** | **float32** | The minimum EPSS Percentile to be included in the results | 
 **isKev** | **bool** | Filter results based on presence in the CISA Known Exploitable Vulnerabilities list | 
 **severity** | **string** | The minimum severity to be included in the results | 
 **lastFoundAfter** | **time.Time** | The earliest last_found_at date to be included in the results | 
 **lastFoundBefore** | **time.Time** | The latest last_found_at date to be included in the results | 
 **orderBy** | **[]string** | List of field name(s) to order by, ascending by default | 
 **orderByDescending** | **[]bool** | Configures the sort order of each specified order_by column to be descending (true) instead of ascending (false) | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSbomGroupVulnerabilities

> SBOMVulnerabilityList ListSbomGroupVulnerabilities(ctx, groupUuid).XAnchoreAccount(xAnchoreAccount).AnchoreScore(anchoreScore).CvssScore(cvssScore).EpssScore(epssScore).EpssPercentile(epssPercentile).IsKev(isKev).Severity(severity).LastFoundAfter(lastFoundAfter).LastFoundBefore(lastFoundBefore).OrderBy(orderBy).OrderByDescending(orderByDescending).MockData(mockData).Execute()

Get a list of vulnerabilities found within the SBOMs belonging to the given SBOM Group during the most recent scans.  Returns summary data of the complete result set and a list of the first 100 results.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	groupUuid := "groupUuid_example" // string | 
	xAnchoreAccount := "xAnchoreAccount_example" // string | The account name used for the resource scope of this request as user permissions allow (optional)
	anchoreScore := float32(3.4) // float32 | The minimum Anchore Score to be included in the results (optional)
	cvssScore := float32(3.4) // float32 | The minimum CVSS Score to be included in the results (optional)
	epssScore := float32(3.4) // float32 | The minimum EPSS Score to be included in the results (optional)
	epssPercentile := float32(3.4) // float32 | The minimum EPSS Percentile to be included in the results (optional)
	isKev := true // bool | Filter results based on presence in the CISA Known Exploitable Vulnerabilities list (optional)
	severity := "severity_example" // string | The minimum severity to be included in the results (optional)
	lastFoundAfter := time.Now() // time.Time | The earliest last_found_at date to be included in the results (optional)
	lastFoundBefore := time.Now() // time.Time | The latest last_found_at date to be included in the results (optional)
	orderBy := []string{"OrderBy_example"} // []string | List of field name(s) to order by, ascending by default (optional)
	orderByDescending := []bool{false} // []bool | Configures the sort order of each specified order_by column to be descending (true) instead of ascending (false) (optional)
	mockData := true // bool | Returns mock data if true. For testing purposes only. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SBOMManagementAPI.ListSbomGroupVulnerabilities(context.Background(), groupUuid).XAnchoreAccount(xAnchoreAccount).AnchoreScore(anchoreScore).CvssScore(cvssScore).EpssScore(epssScore).EpssPercentile(epssPercentile).IsKev(isKev).Severity(severity).LastFoundAfter(lastFoundAfter).LastFoundBefore(lastFoundBefore).OrderBy(orderBy).OrderByDescending(orderByDescending).MockData(mockData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SBOMManagementAPI.ListSbomGroupVulnerabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSbomGroupVulnerabilities`: SBOMVulnerabilityList
	fmt.Fprintf(os.Stdout, "Response from `SBOMManagementAPI.ListSbomGroupVulnerabilities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSbomGroupVulnerabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | The account name used for the resource scope of this request as user permissions allow | 
 **anchoreScore** | **float32** | The minimum Anchore Score to be included in the results | 
 **cvssScore** | **float32** | The minimum CVSS Score to be included in the results | 
 **epssScore** | **float32** | The minimum EPSS Score to be included in the results | 
 **epssPercentile** | **float32** | The minimum EPSS Percentile to be included in the results | 
 **isKev** | **bool** | Filter results based on presence in the CISA Known Exploitable Vulnerabilities list | 
 **severity** | **string** | The minimum severity to be included in the results | 
 **lastFoundAfter** | **time.Time** | The earliest last_found_at date to be included in the results | 
 **lastFoundBefore** | **time.Time** | The latest last_found_at date to be included in the results | 
 **orderBy** | **[]string** | List of field name(s) to order by, ascending by default | 
 **orderByDescending** | **[]bool** | Configures the sort order of each specified order_by column to be descending (true) instead of ascending (false) | 
 **mockData** | **bool** | Returns mock data if true. For testing purposes only. | [default to false]

### Return type

[**SBOMVulnerabilityList**](SBOMVulnerabilityList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSbomGroups

> SBOMGroupDetailList ListSbomGroups(ctx).XAnchoreAccount(xAnchoreAccount).Name(name).Version(version).OrderBy(orderBy).OrderByDescending(orderByDescending).Limit(limit).Page(page).Execute()

Get a list of SBOM Groups belonging to your account

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
	xAnchoreAccount := "xAnchoreAccount_example" // string | The account name used for the resource scope of this request as user permissions allow (optional)
	name := "name_example" // string | group name, not case sensitive (optional)
	version := "version_example" // string | group version (optional)
	orderBy := []string{"OrderBy_example"} // []string | List of field name(s) to order by, ascending by default (optional)
	orderByDescending := []bool{false} // []bool | Configures the sort order of each specified order_by column to be descending (true) instead of ascending (false) (optional)
	limit := int32(56) // int32 | Maximum number of rows to return per page (optional) (default to 1000)
	page := int32(56) // int32 | Page number to return (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SBOMManagementAPI.ListSbomGroups(context.Background()).XAnchoreAccount(xAnchoreAccount).Name(name).Version(version).OrderBy(orderBy).OrderByDescending(orderByDescending).Limit(limit).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SBOMManagementAPI.ListSbomGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSbomGroups`: SBOMGroupDetailList
	fmt.Fprintf(os.Stdout, "Response from `SBOMManagementAPI.ListSbomGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSbomGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAnchoreAccount** | **string** | The account name used for the resource scope of this request as user permissions allow | 
 **name** | **string** | group name, not case sensitive | 
 **version** | **string** | group version | 
 **orderBy** | **[]string** | List of field name(s) to order by, ascending by default | 
 **orderByDescending** | **[]bool** | Configures the sort order of each specified order_by column to be descending (true) instead of ascending (false) | 
 **limit** | **int32** | Maximum number of rows to return per page | [default to 1000]
 **page** | **int32** | Page number to return | [default to 1]

### Return type

[**SBOMGroupDetailList**](SBOMGroupDetailList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSbomVulnerabilities

> SBOMVulnerabilityList ListSbomVulnerabilities(ctx, sbomUuid).XAnchoreAccount(xAnchoreAccount).ForceRefresh(forceRefresh).AnchoreScore(anchoreScore).CvssScore(cvssScore).EpssScore(epssScore).EpssPercentile(epssPercentile).IsKev(isKev).Severity(severity).LastFoundAfter(lastFoundAfter).LastFoundBefore(lastFoundBefore).OrderBy(orderBy).OrderByDescending(orderByDescending).Execute()

Get a list of vulnerabilities found within a given SBOM during the most recent scan. Returns summary data of the complete result set and a list of the first 100 results.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	sbomUuid := "sbomUuid_example" // string | 
	xAnchoreAccount := "xAnchoreAccount_example" // string | The account name used for the resource scope of this request as user permissions allow (optional)
	forceRefresh := true // bool | Force a re-scan of SBOM vulnerabilities before returning results (optional) (default to false)
	anchoreScore := float32(3.4) // float32 | The minimum Anchore Score to be included in the results (optional)
	cvssScore := float32(3.4) // float32 | The minimum CVSS Score to be included in the results (optional)
	epssScore := float32(3.4) // float32 | The minimum EPSS Score to be included in the results (optional)
	epssPercentile := float32(3.4) // float32 | The minimum EPSS Percentile to be included in the results (optional)
	isKev := true // bool | Filter results based on presence in the CISA Known Exploitable Vulnerabilities list (optional)
	severity := "severity_example" // string | The minimum severity to be included in the results (optional)
	lastFoundAfter := time.Now() // time.Time | The earliest last_found_at date to be included in the results (optional)
	lastFoundBefore := time.Now() // time.Time | The latest last_found_at date to be included in the results (optional)
	orderBy := []string{"OrderBy_example"} // []string | List of field name(s) to order by, ascending by default (optional)
	orderByDescending := []bool{false} // []bool | Configures the sort order of each specified order_by column to be descending (true) instead of ascending (false) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SBOMManagementAPI.ListSbomVulnerabilities(context.Background(), sbomUuid).XAnchoreAccount(xAnchoreAccount).ForceRefresh(forceRefresh).AnchoreScore(anchoreScore).CvssScore(cvssScore).EpssScore(epssScore).EpssPercentile(epssPercentile).IsKev(isKev).Severity(severity).LastFoundAfter(lastFoundAfter).LastFoundBefore(lastFoundBefore).OrderBy(orderBy).OrderByDescending(orderByDescending).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SBOMManagementAPI.ListSbomVulnerabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSbomVulnerabilities`: SBOMVulnerabilityList
	fmt.Fprintf(os.Stdout, "Response from `SBOMManagementAPI.ListSbomVulnerabilities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sbomUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSbomVulnerabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | The account name used for the resource scope of this request as user permissions allow | 
 **forceRefresh** | **bool** | Force a re-scan of SBOM vulnerabilities before returning results | [default to false]
 **anchoreScore** | **float32** | The minimum Anchore Score to be included in the results | 
 **cvssScore** | **float32** | The minimum CVSS Score to be included in the results | 
 **epssScore** | **float32** | The minimum EPSS Score to be included in the results | 
 **epssPercentile** | **float32** | The minimum EPSS Percentile to be included in the results | 
 **isKev** | **bool** | Filter results based on presence in the CISA Known Exploitable Vulnerabilities list | 
 **severity** | **string** | The minimum severity to be included in the results | 
 **lastFoundAfter** | **time.Time** | The earliest last_found_at date to be included in the results | 
 **lastFoundBefore** | **time.Time** | The latest last_found_at date to be included in the results | 
 **orderBy** | **[]string** | List of field name(s) to order by, ascending by default | 
 **orderByDescending** | **[]bool** | Configures the sort order of each specified order_by column to be descending (true) instead of ascending (false) | 

### Return type

[**SBOMVulnerabilityList**](SBOMVulnerabilityList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSboms

> SBOMDetailList ListSboms(ctx).XAnchoreAccount(xAnchoreAccount).Name(name).Version(version).Groups(groups).OrderBy(orderBy).OrderByDescending(orderByDescending).Limit(limit).Page(page).Execute()

Get a list of all SBOMs belonging to your account

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
	xAnchoreAccount := "xAnchoreAccount_example" // string | The account name used for the resource scope of this request as user permissions allow (optional)
	name := "name_example" // string | sbom name, not case sensitive (optional)
	version := "version_example" // string | sbom version (optional)
	groups := []string{"Inner_example"} // []string | A list of Group UUIDs (optional)
	orderBy := []string{"OrderBy_example"} // []string | List of field name(s) to order by, ascending by default (optional)
	orderByDescending := []bool{false} // []bool | Configures the sort order of each specified order_by column to be descending (true) instead of ascending (false) (optional)
	limit := int32(56) // int32 | Maximum number of rows to return per page (optional) (default to 1000)
	page := int32(56) // int32 | Page number to return (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SBOMManagementAPI.ListSboms(context.Background()).XAnchoreAccount(xAnchoreAccount).Name(name).Version(version).Groups(groups).OrderBy(orderBy).OrderByDescending(orderByDescending).Limit(limit).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SBOMManagementAPI.ListSboms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSboms`: SBOMDetailList
	fmt.Fprintf(os.Stdout, "Response from `SBOMManagementAPI.ListSboms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSbomsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAnchoreAccount** | **string** | The account name used for the resource scope of this request as user permissions allow | 
 **name** | **string** | sbom name, not case sensitive | 
 **version** | **string** | sbom version | 
 **groups** | **[]string** | A list of Group UUIDs | 
 **orderBy** | **[]string** | List of field name(s) to order by, ascending by default | 
 **orderByDescending** | **[]bool** | Configures the sort order of each specified order_by column to be descending (true) instead of ascending (false) | 
 **limit** | **int32** | Maximum number of rows to return per page | [default to 1000]
 **page** | **int32** | Page number to return | [default to 1]

### Return type

[**SBOMDetailList**](SBOMDetailList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifySbomGroupMemberships

> SBOMGroupMembershipResponse ModifySbomGroupMemberships(ctx, groupUuid).SBOMGroupMembershipRequest(sBOMGroupMembershipRequest).XAnchoreAccount(xAnchoreAccount).Execute()

Modify the set of SBOMs belonging to a given Group. Returns a list of SBOM UUIDs belonging to the Group after the requested changes. Limited to a maximum of 10,000 additions and 10,000 removals in a single request, except when using remove_all.

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
	groupUuid := "groupUuid_example" // string | 
	sBOMGroupMembershipRequest := *openapiclient.NewSBOMGroupMembershipRequest() // SBOMGroupMembershipRequest | 
	xAnchoreAccount := "xAnchoreAccount_example" // string | The account name used for the resource scope of this request as user permissions allow (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SBOMManagementAPI.ModifySbomGroupMemberships(context.Background(), groupUuid).SBOMGroupMembershipRequest(sBOMGroupMembershipRequest).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SBOMManagementAPI.ModifySbomGroupMemberships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySbomGroupMemberships`: SBOMGroupMembershipResponse
	fmt.Fprintf(os.Stdout, "Response from `SBOMManagementAPI.ModifySbomGroupMemberships`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySbomGroupMembershipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sBOMGroupMembershipRequest** | [**SBOMGroupMembershipRequest**](SBOMGroupMembershipRequest.md) |  | 
 **xAnchoreAccount** | **string** | The account name used for the resource scope of this request as user permissions allow | 

### Return type

[**SBOMGroupMembershipResponse**](SBOMGroupMembershipResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSbom

> SBOMDetail PatchSbom(ctx, sbomUuid).XAnchoreAccount(xAnchoreAccount).Name(name).Version(version).Annotations(annotations).Groups(groups).File(file).Execute()

Update the properties of the given SBOM. Existing annotations and group memberships will be updated, not be replaced.

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
	sbomUuid := "sbomUuid_example" // string | 
	xAnchoreAccount := "xAnchoreAccount_example" // string | The account name used for the resource scope of this request as user permissions allow (optional)
	name := "name_example" // string |  (optional)
	version := "version_example" // string |  (optional)
	annotations := "annotations_example" // string | A json string representing a dict of string type key value pairs (optional)
	groups := []string{"Inner_example"} // []string | A comma separated list of Group UUIDs (optional)
	file := os.NewFile(1234, "some_file") // *os.File | A valid SBOM file (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SBOMManagementAPI.PatchSbom(context.Background(), sbomUuid).XAnchoreAccount(xAnchoreAccount).Name(name).Version(version).Annotations(annotations).Groups(groups).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SBOMManagementAPI.PatchSbom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSbom`: SBOMDetail
	fmt.Fprintf(os.Stdout, "Response from `SBOMManagementAPI.PatchSbom`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sbomUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSbomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | The account name used for the resource scope of this request as user permissions allow | 
 **name** | **string** |  | 
 **version** | **string** |  | 
 **annotations** | **string** | A json string representing a dict of string type key value pairs | 
 **groups** | **[]string** | A comma separated list of Group UUIDs | 
 **file** | ***os.File** | A valid SBOM file | 

### Return type

[**SBOMDetail**](SBOMDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSbom

> SBOMDetail PutSbom(ctx, sbomUuid).XAnchoreAccount(xAnchoreAccount).Name(name).Version(version).Annotations(annotations).Groups(groups).File(file).Execute()

Update the properties of the given SBOM. Replaces all existing annotations and group memberships.

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
	sbomUuid := "sbomUuid_example" // string | 
	xAnchoreAccount := "xAnchoreAccount_example" // string | The account name used for the resource scope of this request as user permissions allow (optional)
	name := "name_example" // string |  (optional)
	version := "version_example" // string |  (optional)
	annotations := "annotations_example" // string | A json string representing a dict of string type key value pairs (optional)
	groups := []string{"Inner_example"} // []string | A comma separated list of Group UUIDs (optional)
	file := os.NewFile(1234, "some_file") // *os.File | A valid SBOM file (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SBOMManagementAPI.PutSbom(context.Background(), sbomUuid).XAnchoreAccount(xAnchoreAccount).Name(name).Version(version).Annotations(annotations).Groups(groups).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SBOMManagementAPI.PutSbom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSbom`: SBOMDetail
	fmt.Fprintf(os.Stdout, "Response from `SBOMManagementAPI.PutSbom`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sbomUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSbomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | The account name used for the resource scope of this request as user permissions allow | 
 **name** | **string** |  | 
 **version** | **string** |  | 
 **annotations** | **string** | A json string representing a dict of string type key value pairs | 
 **groups** | **[]string** | A comma separated list of Group UUIDs | 
 **file** | ***os.File** | A valid SBOM file | 

### Return type

[**SBOMDetail**](SBOMDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSbomGroup

> SBOMGroupDetail UpdateSbomGroup(ctx, groupUuid).UpdateSbomGroupRequest(updateSbomGroupRequest).XAnchoreAccount(xAnchoreAccount).Execute()

Update the details of a specific SBOM Group

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
	groupUuid := "groupUuid_example" // string | 
	updateSbomGroupRequest := *openapiclient.NewUpdateSbomGroupRequest() // UpdateSbomGroupRequest | Properties of the Group to be updated
	xAnchoreAccount := "xAnchoreAccount_example" // string | The account name used for the resource scope of this request as user permissions allow (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SBOMManagementAPI.UpdateSbomGroup(context.Background(), groupUuid).UpdateSbomGroupRequest(updateSbomGroupRequest).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SBOMManagementAPI.UpdateSbomGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSbomGroup`: SBOMGroupDetail
	fmt.Fprintf(os.Stdout, "Response from `SBOMManagementAPI.UpdateSbomGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSbomGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSbomGroupRequest** | [**UpdateSbomGroupRequest**](UpdateSbomGroupRequest.md) | Properties of the Group to be updated | 
 **xAnchoreAccount** | **string** | The account name used for the resource scope of this request as user permissions allow | 

### Return type

[**SBOMGroupDetail**](SBOMGroupDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

