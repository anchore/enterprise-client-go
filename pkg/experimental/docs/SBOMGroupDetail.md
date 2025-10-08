# SBOMGroupDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**AccountName** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Sboms** | Pointer to [**[]SBOMDetailGroupsInner**](SBOMDetailGroupsInner.md) |  | [optional] 
**HasSupportedSboms** | Pointer to **bool** | Indicates whether the SBOM Group contains at least one SBOM that is supported for vulnerability scanning and content listing. | [optional] 
**HasUnsupportedSboms** | Pointer to **bool** | Indicates whether the SBOM Group contains at least one SBOM that is NOT supported for vulnerability scanning and content listing. | [optional] 
**QualityScore** | Pointer to **float32** |  | [optional] 

## Methods

### NewSBOMGroupDetail

`func NewSBOMGroupDetail() *SBOMGroupDetail`

NewSBOMGroupDetail instantiates a new SBOMGroupDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSBOMGroupDetailWithDefaults

`func NewSBOMGroupDetailWithDefaults() *SBOMGroupDetail`

NewSBOMGroupDetailWithDefaults instantiates a new SBOMGroupDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *SBOMGroupDetail) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SBOMGroupDetail) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SBOMGroupDetail) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *SBOMGroupDetail) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *SBOMGroupDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SBOMGroupDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SBOMGroupDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SBOMGroupDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccountName

`func (o *SBOMGroupDetail) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *SBOMGroupDetail) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *SBOMGroupDetail) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *SBOMGroupDetail) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetVersion

`func (o *SBOMGroupDetail) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SBOMGroupDetail) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SBOMGroupDetail) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SBOMGroupDetail) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDescription

`func (o *SBOMGroupDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SBOMGroupDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SBOMGroupDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SBOMGroupDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SBOMGroupDetail) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SBOMGroupDetail) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SBOMGroupDetail) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SBOMGroupDetail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SBOMGroupDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SBOMGroupDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SBOMGroupDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SBOMGroupDetail) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetSboms

`func (o *SBOMGroupDetail) GetSboms() []SBOMDetailGroupsInner`

GetSboms returns the Sboms field if non-nil, zero value otherwise.

### GetSbomsOk

`func (o *SBOMGroupDetail) GetSbomsOk() (*[]SBOMDetailGroupsInner, bool)`

GetSbomsOk returns a tuple with the Sboms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSboms

`func (o *SBOMGroupDetail) SetSboms(v []SBOMDetailGroupsInner)`

SetSboms sets Sboms field to given value.

### HasSboms

`func (o *SBOMGroupDetail) HasSboms() bool`

HasSboms returns a boolean if a field has been set.

### GetHasSupportedSboms

`func (o *SBOMGroupDetail) GetHasSupportedSboms() bool`

GetHasSupportedSboms returns the HasSupportedSboms field if non-nil, zero value otherwise.

### GetHasSupportedSbomsOk

`func (o *SBOMGroupDetail) GetHasSupportedSbomsOk() (*bool, bool)`

GetHasSupportedSbomsOk returns a tuple with the HasSupportedSboms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSupportedSboms

`func (o *SBOMGroupDetail) SetHasSupportedSboms(v bool)`

SetHasSupportedSboms sets HasSupportedSboms field to given value.

### HasHasSupportedSboms

`func (o *SBOMGroupDetail) HasHasSupportedSboms() bool`

HasHasSupportedSboms returns a boolean if a field has been set.

### GetHasUnsupportedSboms

`func (o *SBOMGroupDetail) GetHasUnsupportedSboms() bool`

GetHasUnsupportedSboms returns the HasUnsupportedSboms field if non-nil, zero value otherwise.

### GetHasUnsupportedSbomsOk

`func (o *SBOMGroupDetail) GetHasUnsupportedSbomsOk() (*bool, bool)`

GetHasUnsupportedSbomsOk returns a tuple with the HasUnsupportedSboms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUnsupportedSboms

`func (o *SBOMGroupDetail) SetHasUnsupportedSboms(v bool)`

SetHasUnsupportedSboms sets HasUnsupportedSboms field to given value.

### HasHasUnsupportedSboms

`func (o *SBOMGroupDetail) HasHasUnsupportedSboms() bool`

HasHasUnsupportedSboms returns a boolean if a field has been set.

### GetQualityScore

`func (o *SBOMGroupDetail) GetQualityScore() float32`

GetQualityScore returns the QualityScore field if non-nil, zero value otherwise.

### GetQualityScoreOk

`func (o *SBOMGroupDetail) GetQualityScoreOk() (*float32, bool)`

GetQualityScoreOk returns a tuple with the QualityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityScore

`func (o *SBOMGroupDetail) SetQualityScore(v float32)`

SetQualityScore sets QualityScore field to given value.

### HasQualityScore

`func (o *SBOMGroupDetail) HasQualityScore() bool`

HasQualityScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


