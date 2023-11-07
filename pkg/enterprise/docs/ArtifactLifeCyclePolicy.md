# ArtifactLifeCyclePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**Action** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**PolicyConditions** | [**ArtifactLifeCyclePolicyConditions**](ArtifactLifeCyclePolicyConditions.md) |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewArtifactLifeCyclePolicy

`func NewArtifactLifeCyclePolicy(action string, name string, policyConditions ArtifactLifeCyclePolicyConditions, ) *ArtifactLifeCyclePolicy`

NewArtifactLifeCyclePolicy instantiates a new ArtifactLifeCyclePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactLifeCyclePolicyWithDefaults

`func NewArtifactLifeCyclePolicyWithDefaults() *ArtifactLifeCyclePolicy`

NewArtifactLifeCyclePolicyWithDefaults instantiates a new ArtifactLifeCyclePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *ArtifactLifeCyclePolicy) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ArtifactLifeCyclePolicy) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ArtifactLifeCyclePolicy) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ArtifactLifeCyclePolicy) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetAction

`func (o *ArtifactLifeCyclePolicy) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ArtifactLifeCyclePolicy) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ArtifactLifeCyclePolicy) SetAction(v string)`

SetAction sets Action field to given value.


### GetName

`func (o *ArtifactLifeCyclePolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArtifactLifeCyclePolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArtifactLifeCyclePolicy) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ArtifactLifeCyclePolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ArtifactLifeCyclePolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ArtifactLifeCyclePolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ArtifactLifeCyclePolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPolicyConditions

`func (o *ArtifactLifeCyclePolicy) GetPolicyConditions() ArtifactLifeCyclePolicyConditions`

GetPolicyConditions returns the PolicyConditions field if non-nil, zero value otherwise.

### GetPolicyConditionsOk

`func (o *ArtifactLifeCyclePolicy) GetPolicyConditionsOk() (*ArtifactLifeCyclePolicyConditions, bool)`

GetPolicyConditionsOk returns a tuple with the PolicyConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyConditions

`func (o *ArtifactLifeCyclePolicy) SetPolicyConditions(v ArtifactLifeCyclePolicyConditions)`

SetPolicyConditions sets PolicyConditions field to given value.


### GetEnabled

`func (o *ArtifactLifeCyclePolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ArtifactLifeCyclePolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ArtifactLifeCyclePolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ArtifactLifeCyclePolicy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ArtifactLifeCyclePolicy) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ArtifactLifeCyclePolicy) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ArtifactLifeCyclePolicy) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ArtifactLifeCyclePolicy) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ArtifactLifeCyclePolicy) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArtifactLifeCyclePolicy) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArtifactLifeCyclePolicy) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ArtifactLifeCyclePolicy) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


