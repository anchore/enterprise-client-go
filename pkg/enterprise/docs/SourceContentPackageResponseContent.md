# SourceContentPackageResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Package** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Origin** | Pointer to **string** |  | [optional] 
**License** | Pointer to **string** | Deprecated in favor of the &#39;licenses&#39; field\&quot; | [optional] 
**Licenses** | Pointer to **[]string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Cpes** | Pointer to **[]string** | A list of Common Platform Enumerations that may uniquely identify the package | [optional] 
**MetadataType** | Pointer to **string** | The type of the metadata entry | [optional] 
**Metadata** | Pointer to **interface{}** | Package type specific metadata | [optional] 

## Methods

### NewSourceContentPackageResponseContent

`func NewSourceContentPackageResponseContent() *SourceContentPackageResponseContent`

NewSourceContentPackageResponseContent instantiates a new SourceContentPackageResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceContentPackageResponseContentWithDefaults

`func NewSourceContentPackageResponseContentWithDefaults() *SourceContentPackageResponseContent`

NewSourceContentPackageResponseContentWithDefaults instantiates a new SourceContentPackageResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackage

`func (o *SourceContentPackageResponseContent) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *SourceContentPackageResponseContent) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *SourceContentPackageResponseContent) SetPackage(v string)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *SourceContentPackageResponseContent) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetVersion

`func (o *SourceContentPackageResponseContent) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SourceContentPackageResponseContent) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SourceContentPackageResponseContent) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SourceContentPackageResponseContent) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSize

`func (o *SourceContentPackageResponseContent) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SourceContentPackageResponseContent) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SourceContentPackageResponseContent) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *SourceContentPackageResponseContent) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *SourceContentPackageResponseContent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SourceContentPackageResponseContent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SourceContentPackageResponseContent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SourceContentPackageResponseContent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOrigin

`func (o *SourceContentPackageResponseContent) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *SourceContentPackageResponseContent) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *SourceContentPackageResponseContent) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *SourceContentPackageResponseContent) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetLicense

`func (o *SourceContentPackageResponseContent) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *SourceContentPackageResponseContent) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *SourceContentPackageResponseContent) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *SourceContentPackageResponseContent) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetLicenses

`func (o *SourceContentPackageResponseContent) GetLicenses() []string`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *SourceContentPackageResponseContent) GetLicensesOk() (*[]string, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *SourceContentPackageResponseContent) SetLicenses(v []string)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *SourceContentPackageResponseContent) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### GetLocation

`func (o *SourceContentPackageResponseContent) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SourceContentPackageResponseContent) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SourceContentPackageResponseContent) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SourceContentPackageResponseContent) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetCpes

`func (o *SourceContentPackageResponseContent) GetCpes() []string`

GetCpes returns the Cpes field if non-nil, zero value otherwise.

### GetCpesOk

`func (o *SourceContentPackageResponseContent) GetCpesOk() (*[]string, bool)`

GetCpesOk returns a tuple with the Cpes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpes

`func (o *SourceContentPackageResponseContent) SetCpes(v []string)`

SetCpes sets Cpes field to given value.

### HasCpes

`func (o *SourceContentPackageResponseContent) HasCpes() bool`

HasCpes returns a boolean if a field has been set.

### GetMetadataType

`func (o *SourceContentPackageResponseContent) GetMetadataType() string`

GetMetadataType returns the MetadataType field if non-nil, zero value otherwise.

### GetMetadataTypeOk

`func (o *SourceContentPackageResponseContent) GetMetadataTypeOk() (*string, bool)`

GetMetadataTypeOk returns a tuple with the MetadataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataType

`func (o *SourceContentPackageResponseContent) SetMetadataType(v string)`

SetMetadataType sets MetadataType field to given value.

### HasMetadataType

`func (o *SourceContentPackageResponseContent) HasMetadataType() bool`

HasMetadataType returns a boolean if a field has been set.

### GetMetadata

`func (o *SourceContentPackageResponseContent) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SourceContentPackageResponseContent) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SourceContentPackageResponseContent) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SourceContentPackageResponseContent) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


