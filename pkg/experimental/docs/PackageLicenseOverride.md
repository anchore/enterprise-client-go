# PackageLicenseOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Purl** | **string** | The Package URL (PURL) of the package should match a format like \&quot;pkg:pkg_type/pkg_namespace/pkg_name@pkg_version\&quot; | 
**LicenseOverrides** | [**[]LicenseOverride**](LicenseOverride.md) |  | 

## Methods

### NewPackageLicenseOverride

`func NewPackageLicenseOverride(purl string, licenseOverrides []LicenseOverride, ) *PackageLicenseOverride`

NewPackageLicenseOverride instantiates a new PackageLicenseOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageLicenseOverrideWithDefaults

`func NewPackageLicenseOverrideWithDefaults() *PackageLicenseOverride`

NewPackageLicenseOverrideWithDefaults instantiates a new PackageLicenseOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPurl

`func (o *PackageLicenseOverride) GetPurl() string`

GetPurl returns the Purl field if non-nil, zero value otherwise.

### GetPurlOk

`func (o *PackageLicenseOverride) GetPurlOk() (*string, bool)`

GetPurlOk returns a tuple with the Purl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurl

`func (o *PackageLicenseOverride) SetPurl(v string)`

SetPurl sets Purl field to given value.


### GetLicenseOverrides

`func (o *PackageLicenseOverride) GetLicenseOverrides() []LicenseOverride`

GetLicenseOverrides returns the LicenseOverrides field if non-nil, zero value otherwise.

### GetLicenseOverridesOk

`func (o *PackageLicenseOverride) GetLicenseOverridesOk() (*[]LicenseOverride, bool)`

GetLicenseOverridesOk returns a tuple with the LicenseOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseOverrides

`func (o *PackageLicenseOverride) SetLicenseOverrides(v []LicenseOverride)`

SetLicenseOverrides sets LicenseOverrides field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


