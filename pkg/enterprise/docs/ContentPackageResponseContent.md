# ContentPackageResponseContentInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Package** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Origin** | Pointer to **string** |  | [optional] 
**Licenses** | Pointer to **[]string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Cpes** | Pointer to **[]string** | A list of Common Platform Enumerations that may uniquely identify the package | [optional] 
**MetadataType** | Pointer to **string** | The type of the metadata entry | [optional] 
**Metadata** | Pointer to **interface{}** | Package type specific metadata | [optional] 
**Purl** | Pointer to **string** |  | [optional] 

## Methods

### NewContentPackageResponseContentInner

`func NewContentPackageResponseContentInner() *ContentPackageResponseContentInner`

NewContentPackageResponseContentInner instantiates a new ContentPackageResponseContentInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentPackageResponseContentInnerWithDefaults

`func NewContentPackageResponseContentInnerWithDefaults() *ContentPackageResponseContentInner`

NewContentPackageResponseContentInnerWithDefaults instantiates a new ContentPackageResponseContentInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackage

`func (o *ContentPackageResponseContentInner) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *ContentPackageResponseContentInner) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *ContentPackageResponseContentInner) SetPackage(v string)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *ContentPackageResponseContentInner) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetVersion

`func (o *ContentPackageResponseContentInner) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ContentPackageResponseContentInner) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ContentPackageResponseContentInner) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ContentPackageResponseContentInner) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSize

`func (o *ContentPackageResponseContentInner) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ContentPackageResponseContentInner) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ContentPackageResponseContentInner) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *ContentPackageResponseContentInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *ContentPackageResponseContentInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentPackageResponseContentInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentPackageResponseContentInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ContentPackageResponseContentInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOrigin

`func (o *ContentPackageResponseContentInner) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *ContentPackageResponseContentInner) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *ContentPackageResponseContentInner) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *ContentPackageResponseContentInner) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetLicenses

`func (o *ContentPackageResponseContentInner) GetLicenses() []string`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *ContentPackageResponseContentInner) GetLicensesOk() (*[]string, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *ContentPackageResponseContentInner) SetLicenses(v []string)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *ContentPackageResponseContentInner) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### GetLocation

`func (o *ContentPackageResponseContentInner) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ContentPackageResponseContentInner) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ContentPackageResponseContentInner) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ContentPackageResponseContentInner) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetCpes

`func (o *ContentPackageResponseContentInner) GetCpes() []string`

GetCpes returns the Cpes field if non-nil, zero value otherwise.

### GetCpesOk

`func (o *ContentPackageResponseContentInner) GetCpesOk() (*[]string, bool)`

GetCpesOk returns a tuple with the Cpes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpes

`func (o *ContentPackageResponseContentInner) SetCpes(v []string)`

SetCpes sets Cpes field to given value.

### HasCpes

`func (o *ContentPackageResponseContentInner) HasCpes() bool`

HasCpes returns a boolean if a field has been set.

### GetMetadataType

`func (o *ContentPackageResponseContentInner) GetMetadataType() string`

GetMetadataType returns the MetadataType field if non-nil, zero value otherwise.

### GetMetadataTypeOk

`func (o *ContentPackageResponseContentInner) GetMetadataTypeOk() (*string, bool)`

GetMetadataTypeOk returns a tuple with the MetadataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataType

`func (o *ContentPackageResponseContentInner) SetMetadataType(v string)`

SetMetadataType sets MetadataType field to given value.

### HasMetadataType

`func (o *ContentPackageResponseContentInner) HasMetadataType() bool`

HasMetadataType returns a boolean if a field has been set.

### GetMetadata

`func (o *ContentPackageResponseContentInner) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ContentPackageResponseContentInner) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ContentPackageResponseContentInner) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ContentPackageResponseContentInner) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPurl

<<<<<<<< HEAD:pkg/enterprise/docs/ContentPackageResponseContentInner.md
`func (o *ContentPackageResponseContentInner) GetPurl() string`
========
`func (o *ContentPackageResponseContent) GetPurl() string`
>>>>>>>> main:pkg/enterprise/docs/ContentPackageResponseContent.md

GetPurl returns the Purl field if non-nil, zero value otherwise.

### GetPurlOk

<<<<<<<< HEAD:pkg/enterprise/docs/ContentPackageResponseContentInner.md
`func (o *ContentPackageResponseContentInner) GetPurlOk() (*string, bool)`
========
`func (o *ContentPackageResponseContent) GetPurlOk() (*string, bool)`
>>>>>>>> main:pkg/enterprise/docs/ContentPackageResponseContent.md

GetPurlOk returns a tuple with the Purl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurl

<<<<<<<< HEAD:pkg/enterprise/docs/ContentPackageResponseContentInner.md
`func (o *ContentPackageResponseContentInner) SetPurl(v string)`
========
`func (o *ContentPackageResponseContent) SetPurl(v string)`
>>>>>>>> main:pkg/enterprise/docs/ContentPackageResponseContent.md

SetPurl sets Purl field to given value.

### HasPurl

<<<<<<<< HEAD:pkg/enterprise/docs/ContentPackageResponseContentInner.md
`func (o *ContentPackageResponseContentInner) HasPurl() bool`
========
`func (o *ContentPackageResponseContent) HasPurl() bool`
>>>>>>>> main:pkg/enterprise/docs/ContentPackageResponseContent.md

HasPurl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


