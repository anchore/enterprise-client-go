# ContentPackageResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Package** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **interface{}** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Origin** | Pointer to **string** |  | [optional] 
**License** | Pointer to **string** | Deprecated in favor of the &#39;licenses&#39; field\&quot; | [optional] 
**Licenses** | Pointer to **[]string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Cpes** | Pointer to **[]string** | A list of Common Platform Enumerations that may uniquely identify the package | [optional] 

## Methods

### NewContentPackageResponseContent

`func NewContentPackageResponseContent() *ContentPackageResponseContent`

NewContentPackageResponseContent instantiates a new ContentPackageResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentPackageResponseContentWithDefaults

`func NewContentPackageResponseContentWithDefaults() *ContentPackageResponseContent`

NewContentPackageResponseContentWithDefaults instantiates a new ContentPackageResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackage

`func (o *ContentPackageResponseContent) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *ContentPackageResponseContent) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *ContentPackageResponseContent) SetPackage(v string)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *ContentPackageResponseContent) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetVersion

`func (o *ContentPackageResponseContent) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ContentPackageResponseContent) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ContentPackageResponseContent) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ContentPackageResponseContent) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSize

`func (o *ContentPackageResponseContent) GetSize() interface{}`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ContentPackageResponseContent) GetSizeOk() (*interface{}, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ContentPackageResponseContent) SetSize(v interface{})`

SetSize sets Size field to given value.

### HasSize

`func (o *ContentPackageResponseContent) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *ContentPackageResponseContent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentPackageResponseContent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentPackageResponseContent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ContentPackageResponseContent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOrigin

`func (o *ContentPackageResponseContent) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *ContentPackageResponseContent) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *ContentPackageResponseContent) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *ContentPackageResponseContent) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetLicense

`func (o *ContentPackageResponseContent) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *ContentPackageResponseContent) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *ContentPackageResponseContent) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *ContentPackageResponseContent) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetLicenses

`func (o *ContentPackageResponseContent) GetLicenses() []string`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *ContentPackageResponseContent) GetLicensesOk() (*[]string, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *ContentPackageResponseContent) SetLicenses(v []string)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *ContentPackageResponseContent) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### GetLocation

`func (o *ContentPackageResponseContent) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ContentPackageResponseContent) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ContentPackageResponseContent) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ContentPackageResponseContent) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetCpes

`func (o *ContentPackageResponseContent) GetCpes() []string`

GetCpes returns the Cpes field if non-nil, zero value otherwise.

### GetCpesOk

`func (o *ContentPackageResponseContent) GetCpesOk() (*[]string, bool)`

GetCpesOk returns a tuple with the Cpes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpes

`func (o *ContentPackageResponseContent) SetCpes(v []string)`

SetCpes sets Cpes field to given value.

### HasCpes

`func (o *ContentPackageResponseContent) HasCpes() bool`

HasCpes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


