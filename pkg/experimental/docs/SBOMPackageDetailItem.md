# SBOMPackageDetailItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Purl** | Pointer to **string** |  | [optional] 
**Licenses** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSBOMPackageDetailItem

`func NewSBOMPackageDetailItem() *SBOMPackageDetailItem`

NewSBOMPackageDetailItem instantiates a new SBOMPackageDetailItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSBOMPackageDetailItemWithDefaults

`func NewSBOMPackageDetailItemWithDefaults() *SBOMPackageDetailItem`

NewSBOMPackageDetailItemWithDefaults instantiates a new SBOMPackageDetailItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SBOMPackageDetailItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SBOMPackageDetailItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SBOMPackageDetailItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SBOMPackageDetailItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *SBOMPackageDetailItem) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SBOMPackageDetailItem) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SBOMPackageDetailItem) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SBOMPackageDetailItem) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetPurl

`func (o *SBOMPackageDetailItem) GetPurl() string`

GetPurl returns the Purl field if non-nil, zero value otherwise.

### GetPurlOk

`func (o *SBOMPackageDetailItem) GetPurlOk() (*string, bool)`

GetPurlOk returns a tuple with the Purl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurl

`func (o *SBOMPackageDetailItem) SetPurl(v string)`

SetPurl sets Purl field to given value.

### HasPurl

`func (o *SBOMPackageDetailItem) HasPurl() bool`

HasPurl returns a boolean if a field has been set.

### GetLicenses

`func (o *SBOMPackageDetailItem) GetLicenses() []string`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *SBOMPackageDetailItem) GetLicensesOk() (*[]string, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *SBOMPackageDetailItem) SetLicenses(v []string)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *SBOMPackageDetailItem) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


