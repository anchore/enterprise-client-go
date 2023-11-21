# ArtifactLifeCyclePolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] [readonly] 
**Action** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**PolicyConditions** | [**[]ArtifactLifeCyclePolicyConditions**](ArtifactLifeCyclePolicyConditions.md) |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewArtifactLifeCyclePolicyResponse

`func NewArtifactLifeCyclePolicyResponse(action string, name string, policyConditions []ArtifactLifeCyclePolicyConditions, ) *ArtifactLifeCyclePolicyResponse`

NewArtifactLifeCyclePolicyResponse instantiates a new ArtifactLifeCyclePolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactLifeCyclePolicyResponseWithDefaults

`func NewArtifactLifeCyclePolicyResponseWithDefaults() *ArtifactLifeCyclePolicyResponse`

NewArtifactLifeCyclePolicyResponseWithDefaults instantiates a new ArtifactLifeCyclePolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *ArtifactLifeCyclePolicyResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ArtifactLifeCyclePolicyResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ArtifactLifeCyclePolicyResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ArtifactLifeCyclePolicyResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetAction

`func (o *ArtifactLifeCyclePolicyResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ArtifactLifeCyclePolicyResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ArtifactLifeCyclePolicyResponse) SetAction(v string)`

SetAction sets Action field to given value.


### GetName

`func (o *ArtifactLifeCyclePolicyResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArtifactLifeCyclePolicyResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArtifactLifeCyclePolicyResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ArtifactLifeCyclePolicyResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ArtifactLifeCyclePolicyResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ArtifactLifeCyclePolicyResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ArtifactLifeCyclePolicyResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPolicyConditions

`func (o *ArtifactLifeCyclePolicyResponse) GetPolicyConditions() []ArtifactLifeCyclePolicyConditions`

GetPolicyConditions returns the PolicyConditions field if non-nil, zero value otherwise.

### GetPolicyConditionsOk

`func (o *ArtifactLifeCyclePolicyResponse) GetPolicyConditionsOk() (*[]ArtifactLifeCyclePolicyConditions, bool)`

GetPolicyConditionsOk returns a tuple with the PolicyConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyConditions

`func (o *ArtifactLifeCyclePolicyResponse) SetPolicyConditions(v []ArtifactLifeCyclePolicyConditions)`

SetPolicyConditions sets PolicyConditions field to given value.


### GetEnabled

`func (o *ArtifactLifeCyclePolicyResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ArtifactLifeCyclePolicyResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ArtifactLifeCyclePolicyResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ArtifactLifeCyclePolicyResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ArtifactLifeCyclePolicyResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ArtifactLifeCyclePolicyResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ArtifactLifeCyclePolicyResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ArtifactLifeCyclePolicyResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ArtifactLifeCyclePolicyResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArtifactLifeCyclePolicyResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArtifactLifeCyclePolicyResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ArtifactLifeCyclePolicyResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ArtifactLifeCyclePolicyResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ArtifactLifeCyclePolicyResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ArtifactLifeCyclePolicyResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ArtifactLifeCyclePolicyResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


