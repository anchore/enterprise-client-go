# LicenseReviewItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Package** | Pointer to [**LicenseReviewPackage**](LicenseReviewPackage.md) |  | [optional] 
**Licenses** | Pointer to [**[]LicenseReviewPackageLicense**](LicenseReviewPackageLicense.md) | List of licenses for the package | [optional] 

## Methods

### NewLicenseReviewItem

`func NewLicenseReviewItem() *LicenseReviewItem`

NewLicenseReviewItem instantiates a new LicenseReviewItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseReviewItemWithDefaults

`func NewLicenseReviewItemWithDefaults() *LicenseReviewItem`

NewLicenseReviewItemWithDefaults instantiates a new LicenseReviewItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackage

`func (o *LicenseReviewItem) GetPackage() LicenseReviewPackage`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *LicenseReviewItem) GetPackageOk() (*LicenseReviewPackage, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *LicenseReviewItem) SetPackage(v LicenseReviewPackage)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *LicenseReviewItem) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetLicenses

`func (o *LicenseReviewItem) GetLicenses() []LicenseReviewPackageLicense`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *LicenseReviewItem) GetLicensesOk() (*[]LicenseReviewPackageLicense, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *LicenseReviewItem) SetLicenses(v []LicenseReviewPackageLicense)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *LicenseReviewItem) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


