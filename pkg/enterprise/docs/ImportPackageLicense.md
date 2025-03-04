# ImportPackageLicense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | The raw license value reported in the packages metadata | [optional] 
**SpdxExpression** | Pointer to **string** | The SPDX expression for the license | [optional] 
**Type** | Pointer to **string** | Where the license information came from | [optional] 
**Urls** | Pointer to **[]string** | URLs associated with the license | [optional] 
**Locations** | Pointer to [**[]ImportPackageLocation**](ImportPackageLocation.md) |  | [optional] 
**Contents** | Pointer to **string** | The contents of the license file | [optional] 

## Methods

### NewImportPackageLicense

`func NewImportPackageLicense() *ImportPackageLicense`

NewImportPackageLicense instantiates a new ImportPackageLicense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportPackageLicenseWithDefaults

`func NewImportPackageLicenseWithDefaults() *ImportPackageLicense`

NewImportPackageLicenseWithDefaults instantiates a new ImportPackageLicense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ImportPackageLicense) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ImportPackageLicense) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ImportPackageLicense) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ImportPackageLicense) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetSpdxExpression

`func (o *ImportPackageLicense) GetSpdxExpression() string`

GetSpdxExpression returns the SpdxExpression field if non-nil, zero value otherwise.

### GetSpdxExpressionOk

`func (o *ImportPackageLicense) GetSpdxExpressionOk() (*string, bool)`

GetSpdxExpressionOk returns a tuple with the SpdxExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpdxExpression

`func (o *ImportPackageLicense) SetSpdxExpression(v string)`

SetSpdxExpression sets SpdxExpression field to given value.

### HasSpdxExpression

`func (o *ImportPackageLicense) HasSpdxExpression() bool`

HasSpdxExpression returns a boolean if a field has been set.

### GetType

`func (o *ImportPackageLicense) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ImportPackageLicense) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ImportPackageLicense) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ImportPackageLicense) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrls

`func (o *ImportPackageLicense) GetUrls() []string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *ImportPackageLicense) GetUrlsOk() (*[]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *ImportPackageLicense) SetUrls(v []string)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *ImportPackageLicense) HasUrls() bool`

HasUrls returns a boolean if a field has been set.

### GetLocations

`func (o *ImportPackageLicense) GetLocations() []ImportPackageLocation`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *ImportPackageLicense) GetLocationsOk() (*[]ImportPackageLocation, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *ImportPackageLicense) SetLocations(v []ImportPackageLocation)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *ImportPackageLicense) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetContents

`func (o *ImportPackageLicense) GetContents() string`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *ImportPackageLicense) GetContentsOk() (*string, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *ImportPackageLicense) SetContents(v string)`

SetContents sets Contents field to given value.

### HasContents

`func (o *ImportPackageLicense) HasContents() bool`

HasContents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


