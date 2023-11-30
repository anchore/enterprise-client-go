# PackageContent

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

### NewPackageContent

`func NewPackageContent() *PackageContent`

NewPackageContent instantiates a new PackageContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageContentWithDefaults

`func NewPackageContentWithDefaults() *PackageContent`

NewPackageContentWithDefaults instantiates a new PackageContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackage

`func (o *PackageContent) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *PackageContent) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *PackageContent) SetPackage(v string)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *PackageContent) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetVersion

`func (o *PackageContent) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PackageContent) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PackageContent) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PackageContent) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSize

`func (o *PackageContent) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PackageContent) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PackageContent) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *PackageContent) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *PackageContent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PackageContent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PackageContent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PackageContent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOrigin

`func (o *PackageContent) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *PackageContent) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *PackageContent) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *PackageContent) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetLicenses

`func (o *PackageContent) GetLicenses() []string`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *PackageContent) GetLicensesOk() (*[]string, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *PackageContent) SetLicenses(v []string)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *PackageContent) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### GetLocation

`func (o *PackageContent) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PackageContent) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PackageContent) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PackageContent) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetCpes

`func (o *PackageContent) GetCpes() []string`

GetCpes returns the Cpes field if non-nil, zero value otherwise.

### GetCpesOk

`func (o *PackageContent) GetCpesOk() (*[]string, bool)`

GetCpesOk returns a tuple with the Cpes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpes

`func (o *PackageContent) SetCpes(v []string)`

SetCpes sets Cpes field to given value.

### HasCpes

`func (o *PackageContent) HasCpes() bool`

HasCpes returns a boolean if a field has been set.

### GetMetadataType

`func (o *PackageContent) GetMetadataType() string`

GetMetadataType returns the MetadataType field if non-nil, zero value otherwise.

### GetMetadataTypeOk

`func (o *PackageContent) GetMetadataTypeOk() (*string, bool)`

GetMetadataTypeOk returns a tuple with the MetadataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataType

`func (o *PackageContent) SetMetadataType(v string)`

SetMetadataType sets MetadataType field to given value.

### HasMetadataType

`func (o *PackageContent) HasMetadataType() bool`

HasMetadataType returns a boolean if a field has been set.

### GetMetadata

`func (o *PackageContent) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PackageContent) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PackageContent) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PackageContent) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPurl

`func (o *PackageContent) GetPurl() string`

GetPurl returns the Purl field if non-nil, zero value otherwise.

### GetPurlOk

`func (o *PackageContent) GetPurlOk() (*string, bool)`

GetPurlOk returns a tuple with the Purl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurl

`func (o *PackageContent) SetPurl(v string)`

SetPurl sets Purl field to given value.

### HasPurl

`func (o *PackageContent) HasPurl() bool`

HasPurl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


