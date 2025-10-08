# LicenseOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LicenseId** | **string** | The ID of the license override e.g. GPL-2.0-only, Apache-2.0 | 
**LicenseName** | Pointer to **string** | The human readable name of the license e.g. GNU General Public License v2.0 | [optional] 
**LicenseText** | Pointer to **string** | The full text of the license | [optional] 
**LicenseUrl** | Pointer to **string** | The URL to the license definition | [optional] 
**LicenseHeader** | Pointer to **string** | The standard license header for the package’s license | [optional] 
**UpdatedByUser** | Pointer to **string** | The user who last updated the license override | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time the license override was created | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The time the license override was last updated | [optional] 

## Methods

### NewLicenseOverride

`func NewLicenseOverride(licenseId string, ) *LicenseOverride`

NewLicenseOverride instantiates a new LicenseOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseOverrideWithDefaults

`func NewLicenseOverrideWithDefaults() *LicenseOverride`

NewLicenseOverrideWithDefaults instantiates a new LicenseOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicenseId

`func (o *LicenseOverride) GetLicenseId() string`

GetLicenseId returns the LicenseId field if non-nil, zero value otherwise.

### GetLicenseIdOk

`func (o *LicenseOverride) GetLicenseIdOk() (*string, bool)`

GetLicenseIdOk returns a tuple with the LicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseId

`func (o *LicenseOverride) SetLicenseId(v string)`

SetLicenseId sets LicenseId field to given value.


### GetLicenseName

`func (o *LicenseOverride) GetLicenseName() string`

GetLicenseName returns the LicenseName field if non-nil, zero value otherwise.

### GetLicenseNameOk

`func (o *LicenseOverride) GetLicenseNameOk() (*string, bool)`

GetLicenseNameOk returns a tuple with the LicenseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseName

`func (o *LicenseOverride) SetLicenseName(v string)`

SetLicenseName sets LicenseName field to given value.

### HasLicenseName

`func (o *LicenseOverride) HasLicenseName() bool`

HasLicenseName returns a boolean if a field has been set.

### GetLicenseText

`func (o *LicenseOverride) GetLicenseText() string`

GetLicenseText returns the LicenseText field if non-nil, zero value otherwise.

### GetLicenseTextOk

`func (o *LicenseOverride) GetLicenseTextOk() (*string, bool)`

GetLicenseTextOk returns a tuple with the LicenseText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseText

`func (o *LicenseOverride) SetLicenseText(v string)`

SetLicenseText sets LicenseText field to given value.

### HasLicenseText

`func (o *LicenseOverride) HasLicenseText() bool`

HasLicenseText returns a boolean if a field has been set.

### GetLicenseUrl

`func (o *LicenseOverride) GetLicenseUrl() string`

GetLicenseUrl returns the LicenseUrl field if non-nil, zero value otherwise.

### GetLicenseUrlOk

`func (o *LicenseOverride) GetLicenseUrlOk() (*string, bool)`

GetLicenseUrlOk returns a tuple with the LicenseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseUrl

`func (o *LicenseOverride) SetLicenseUrl(v string)`

SetLicenseUrl sets LicenseUrl field to given value.

### HasLicenseUrl

`func (o *LicenseOverride) HasLicenseUrl() bool`

HasLicenseUrl returns a boolean if a field has been set.

### GetLicenseHeader

`func (o *LicenseOverride) GetLicenseHeader() string`

GetLicenseHeader returns the LicenseHeader field if non-nil, zero value otherwise.

### GetLicenseHeaderOk

`func (o *LicenseOverride) GetLicenseHeaderOk() (*string, bool)`

GetLicenseHeaderOk returns a tuple with the LicenseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseHeader

`func (o *LicenseOverride) SetLicenseHeader(v string)`

SetLicenseHeader sets LicenseHeader field to given value.

### HasLicenseHeader

`func (o *LicenseOverride) HasLicenseHeader() bool`

HasLicenseHeader returns a boolean if a field has been set.

### GetUpdatedByUser

`func (o *LicenseOverride) GetUpdatedByUser() string`

GetUpdatedByUser returns the UpdatedByUser field if non-nil, zero value otherwise.

### GetUpdatedByUserOk

`func (o *LicenseOverride) GetUpdatedByUserOk() (*string, bool)`

GetUpdatedByUserOk returns a tuple with the UpdatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByUser

`func (o *LicenseOverride) SetUpdatedByUser(v string)`

SetUpdatedByUser sets UpdatedByUser field to given value.

### HasUpdatedByUser

`func (o *LicenseOverride) HasUpdatedByUser() bool`

HasUpdatedByUser returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LicenseOverride) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LicenseOverride) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LicenseOverride) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LicenseOverride) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *LicenseOverride) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LicenseOverride) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LicenseOverride) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LicenseOverride) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


