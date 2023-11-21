# PolicyRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Gate** | **string** |  | 
**Trigger** | **string** |  | 
**Action** | **string** |  | 
**Description** | Pointer to **string** | Description of the policy rule, human readable | [optional] 
<<<<<<< HEAD
**Params** | [**[]PolicyRuleParamsInner**](PolicyRuleParamsInner.md) |  | 
=======
**Params** | [**[]PolicyRuleParams**](PolicyRuleParams.md) |  | 
>>>>>>> main
**Recommendation** | Pointer to **string** |  | [optional] 

## Methods

### NewPolicyRule

<<<<<<< HEAD
`func NewPolicyRule(id string, gate string, trigger string, action string, params []PolicyRuleParamsInner, ) *PolicyRule`
=======
`func NewPolicyRule(id string, gate string, trigger string, action string, params []PolicyRuleParams, ) *PolicyRule`
>>>>>>> main

NewPolicyRule instantiates a new PolicyRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyRuleWithDefaults

`func NewPolicyRuleWithDefaults() *PolicyRule`

NewPolicyRuleWithDefaults instantiates a new PolicyRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PolicyRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyRule) SetId(v string)`

SetId sets Id field to given value.


### GetGate

`func (o *PolicyRule) GetGate() string`

GetGate returns the Gate field if non-nil, zero value otherwise.

### GetGateOk

`func (o *PolicyRule) GetGateOk() (*string, bool)`

GetGateOk returns a tuple with the Gate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGate

`func (o *PolicyRule) SetGate(v string)`

SetGate sets Gate field to given value.


### GetTrigger

`func (o *PolicyRule) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *PolicyRule) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *PolicyRule) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.


### GetAction

`func (o *PolicyRule) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PolicyRule) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PolicyRule) SetAction(v string)`

SetAction sets Action field to given value.


### GetDescription

`func (o *PolicyRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParams

`func (o *PolicyRule) GetParams() []PolicyRuleParamsInner`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *PolicyRule) GetParamsOk() (*[]PolicyRuleParamsInner, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *PolicyRule) SetParams(v []PolicyRuleParamsInner)`

SetParams sets Params field to given value.


### GetRecommendation

`func (o *PolicyRule) GetRecommendation() string`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *PolicyRule) GetRecommendationOk() (*string, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *PolicyRule) SetRecommendation(v string)`

SetRecommendation sets Recommendation field to given value.

### HasRecommendation

`func (o *PolicyRule) HasRecommendation() bool`

HasRecommendation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


