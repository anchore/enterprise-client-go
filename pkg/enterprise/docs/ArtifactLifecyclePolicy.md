# ArtifactLifecyclePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | A system defined unique identifier. | [optional] [readonly] 
**Action** | **string** | The action that should be taken when the rule parameters are met. | 
**Name** | **string** | A user provided name for the policy. This name must be unique for an Artifact Lifecycle Policy. | 
**Description** | Pointer to **string** | A user provided description for the policy. | [optional] 
**PolicyConditions** | [**ArtifactLifecyclePolicyConditions**](ArtifactLifecyclePolicyConditions.md) |  | 
**Enabled** | Pointer to **bool** | Indicates if the policy should be active or not.  Defaulted to false. | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewArtifactLifecyclePolicy

`func NewArtifactLifecyclePolicy(action string, name string, policyConditions ArtifactLifecyclePolicyConditions, ) *ArtifactLifecyclePolicy`

NewArtifactLifecyclePolicy instantiates a new ArtifactLifecyclePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactLifecyclePolicyWithDefaults

`func NewArtifactLifecyclePolicyWithDefaults() *ArtifactLifecyclePolicy`

NewArtifactLifecyclePolicyWithDefaults instantiates a new ArtifactLifecyclePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *ArtifactLifecyclePolicy) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ArtifactLifecyclePolicy) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ArtifactLifecyclePolicy) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ArtifactLifecyclePolicy) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetAction

`func (o *ArtifactLifecyclePolicy) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ArtifactLifecyclePolicy) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ArtifactLifecyclePolicy) SetAction(v string)`

SetAction sets Action field to given value.


### GetName

`func (o *ArtifactLifecyclePolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArtifactLifecyclePolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArtifactLifecyclePolicy) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ArtifactLifecyclePolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ArtifactLifecyclePolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ArtifactLifecyclePolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ArtifactLifecyclePolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPolicyConditions

`func (o *ArtifactLifecyclePolicy) GetPolicyConditions() ArtifactLifecyclePolicyConditions`

GetPolicyConditions returns the PolicyConditions field if non-nil, zero value otherwise.

### GetPolicyConditionsOk

`func (o *ArtifactLifecyclePolicy) GetPolicyConditionsOk() (*ArtifactLifecyclePolicyConditions, bool)`

GetPolicyConditionsOk returns a tuple with the PolicyConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyConditions

`func (o *ArtifactLifecyclePolicy) SetPolicyConditions(v ArtifactLifecyclePolicyConditions)`

SetPolicyConditions sets PolicyConditions field to given value.


### GetEnabled

`func (o *ArtifactLifecyclePolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ArtifactLifecyclePolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ArtifactLifecyclePolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ArtifactLifecyclePolicy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ArtifactLifecyclePolicy) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ArtifactLifecyclePolicy) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ArtifactLifecyclePolicy) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ArtifactLifecyclePolicy) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ArtifactLifecyclePolicy) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArtifactLifecyclePolicy) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArtifactLifecyclePolicy) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ArtifactLifecyclePolicy) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ArtifactLifecyclePolicy) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ArtifactLifecyclePolicy) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ArtifactLifecyclePolicy) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ArtifactLifecyclePolicy) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


