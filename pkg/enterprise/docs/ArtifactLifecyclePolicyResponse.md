# ArtifactLifecyclePolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | A system defined unique identifier. | [optional] [readonly] 
**Action** | **string** | The action that should be taken when the rule parameters are met. | 
**Name** | **string** | A user provided name for the policy. This name must be unique for an Artifact Lifecycle Policy. | 
**Description** | Pointer to **string** | A user provided description for the policy. | [optional] 
**PolicyConditions** | [**[]ArtifactLifecyclePolicyConditions**](ArtifactLifecyclePolicyConditions.md) |  | 
**Enabled** | Pointer to **bool** | Indicates if the policy should be active or not.  Defaulted to false. | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewArtifactLifecyclePolicyResponse

`func NewArtifactLifecyclePolicyResponse(action string, name string, policyConditions []ArtifactLifecyclePolicyConditions, ) *ArtifactLifecyclePolicyResponse`

NewArtifactLifecyclePolicyResponse instantiates a new ArtifactLifecyclePolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactLifecyclePolicyResponseWithDefaults

`func NewArtifactLifecyclePolicyResponseWithDefaults() *ArtifactLifecyclePolicyResponse`

NewArtifactLifecyclePolicyResponseWithDefaults instantiates a new ArtifactLifecyclePolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *ArtifactLifecyclePolicyResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ArtifactLifecyclePolicyResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ArtifactLifecyclePolicyResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ArtifactLifecyclePolicyResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetAction

`func (o *ArtifactLifecyclePolicyResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ArtifactLifecyclePolicyResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ArtifactLifecyclePolicyResponse) SetAction(v string)`

SetAction sets Action field to given value.


### GetName

`func (o *ArtifactLifecyclePolicyResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArtifactLifecyclePolicyResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArtifactLifecyclePolicyResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ArtifactLifecyclePolicyResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ArtifactLifecyclePolicyResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ArtifactLifecyclePolicyResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ArtifactLifecyclePolicyResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPolicyConditions

`func (o *ArtifactLifecyclePolicyResponse) GetPolicyConditions() []ArtifactLifecyclePolicyConditions`

GetPolicyConditions returns the PolicyConditions field if non-nil, zero value otherwise.

### GetPolicyConditionsOk

`func (o *ArtifactLifecyclePolicyResponse) GetPolicyConditionsOk() (*[]ArtifactLifecyclePolicyConditions, bool)`

GetPolicyConditionsOk returns a tuple with the PolicyConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyConditions

`func (o *ArtifactLifecyclePolicyResponse) SetPolicyConditions(v []ArtifactLifecyclePolicyConditions)`

SetPolicyConditions sets PolicyConditions field to given value.


### GetEnabled

`func (o *ArtifactLifecyclePolicyResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ArtifactLifecyclePolicyResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ArtifactLifecyclePolicyResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ArtifactLifecyclePolicyResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ArtifactLifecyclePolicyResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ArtifactLifecyclePolicyResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ArtifactLifecyclePolicyResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ArtifactLifecyclePolicyResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ArtifactLifecyclePolicyResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArtifactLifecyclePolicyResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArtifactLifecyclePolicyResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ArtifactLifecyclePolicyResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ArtifactLifecyclePolicyResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ArtifactLifecyclePolicyResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ArtifactLifecyclePolicyResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ArtifactLifecyclePolicyResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


