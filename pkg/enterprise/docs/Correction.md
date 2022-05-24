# Correction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Identifier for the correction | [optional] 
**Type** | **string** | Type of correction | 
**Description** | Pointer to **string** |  | [optional] 
**Match** | [**CorrectionMatch**](CorrectionMatch.md) |  | 
**Replace** | [**[]CorrectionFieldMatch**](CorrectionFieldMatch.md) |  | 
**CreatedAt** | Pointer to **time.Time** | RFC 3339 formatted UTC timestamp when the correction was generated | [optional] 
**LastUpdated** | Pointer to **time.Time** | RFC 3339 formatted UTC timestamp when the correction was last modified | [optional] 

## Methods

### NewCorrection

`func NewCorrection(type_ string, match CorrectionMatch, replace []CorrectionFieldMatch, ) *Correction`

NewCorrection instantiates a new Correction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorrectionWithDefaults

`func NewCorrectionWithDefaults() *Correction`

NewCorrectionWithDefaults instantiates a new Correction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *Correction) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Correction) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Correction) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Correction) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *Correction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Correction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Correction) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *Correction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Correction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Correction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Correction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMatch

`func (o *Correction) GetMatch() CorrectionMatch`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *Correction) GetMatchOk() (*CorrectionMatch, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *Correction) SetMatch(v CorrectionMatch)`

SetMatch sets Match field to given value.


### GetReplace

`func (o *Correction) GetReplace() []CorrectionFieldMatch`

GetReplace returns the Replace field if non-nil, zero value otherwise.

### GetReplaceOk

`func (o *Correction) GetReplaceOk() (*[]CorrectionFieldMatch, bool)`

GetReplaceOk returns a tuple with the Replace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplace

`func (o *Correction) SetReplace(v []CorrectionFieldMatch)`

SetReplace sets Replace field to given value.


### GetCreatedAt

`func (o *Correction) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Correction) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Correction) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Correction) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Correction) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Correction) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Correction) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Correction) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


