# PolicyEvaluationFinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TriggerId** | **string** | ID of this policy trigger finding (can be used to allowlist this finding) | 
**Gate** | **string** | Name of the gate that generated this finding | 
**Trigger** | **string** | Name of the trigger that generated this finding | 
**Message** | **string** | Description of the finding | 
**Action** | **string** | The action associated with this finding | 
**PolicyId** | **string** | ID of the policy that this gate and trigger are a part of | 
**Recommendation** | **string** | User provided details for resolving this finding | 
**RuleId** | **string** | ID of the policy rule that that generated this finding | 
**Allowlisted** | **bool** | Indicates if this finding was allowlisted or not | 
**AllowlistMatch** | Pointer to [**NullablePolicyFindingAllowlistMatch**](PolicyFindingAllowlistMatch.md) |  | [optional] 
**InheritedFromBase** | Pointer to **NullableBool** | Indicates if this finding was found in the base image | [optional] 

## Methods

### NewPolicyEvaluationFinding

`func NewPolicyEvaluationFinding(triggerId string, gate string, trigger string, message string, action string, policyId string, recommendation string, ruleId string, allowlisted bool, ) *PolicyEvaluationFinding`

NewPolicyEvaluationFinding instantiates a new PolicyEvaluationFinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyEvaluationFindingWithDefaults

`func NewPolicyEvaluationFindingWithDefaults() *PolicyEvaluationFinding`

NewPolicyEvaluationFindingWithDefaults instantiates a new PolicyEvaluationFinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTriggerId

`func (o *PolicyEvaluationFinding) GetTriggerId() string`

GetTriggerId returns the TriggerId field if non-nil, zero value otherwise.

### GetTriggerIdOk

`func (o *PolicyEvaluationFinding) GetTriggerIdOk() (*string, bool)`

GetTriggerIdOk returns a tuple with the TriggerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerId

`func (o *PolicyEvaluationFinding) SetTriggerId(v string)`

SetTriggerId sets TriggerId field to given value.


### GetGate

`func (o *PolicyEvaluationFinding) GetGate() string`

GetGate returns the Gate field if non-nil, zero value otherwise.

### GetGateOk

`func (o *PolicyEvaluationFinding) GetGateOk() (*string, bool)`

GetGateOk returns a tuple with the Gate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGate

`func (o *PolicyEvaluationFinding) SetGate(v string)`

SetGate sets Gate field to given value.


### GetTrigger

`func (o *PolicyEvaluationFinding) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *PolicyEvaluationFinding) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *PolicyEvaluationFinding) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.


### GetMessage

`func (o *PolicyEvaluationFinding) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PolicyEvaluationFinding) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PolicyEvaluationFinding) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetAction

`func (o *PolicyEvaluationFinding) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PolicyEvaluationFinding) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PolicyEvaluationFinding) SetAction(v string)`

SetAction sets Action field to given value.


### GetPolicyId

`func (o *PolicyEvaluationFinding) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *PolicyEvaluationFinding) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *PolicyEvaluationFinding) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.


### GetRecommendation

`func (o *PolicyEvaluationFinding) GetRecommendation() string`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *PolicyEvaluationFinding) GetRecommendationOk() (*string, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *PolicyEvaluationFinding) SetRecommendation(v string)`

SetRecommendation sets Recommendation field to given value.


### GetRuleId

`func (o *PolicyEvaluationFinding) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *PolicyEvaluationFinding) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *PolicyEvaluationFinding) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.


### GetAllowlisted

`func (o *PolicyEvaluationFinding) GetAllowlisted() bool`

GetAllowlisted returns the Allowlisted field if non-nil, zero value otherwise.

### GetAllowlistedOk

`func (o *PolicyEvaluationFinding) GetAllowlistedOk() (*bool, bool)`

GetAllowlistedOk returns a tuple with the Allowlisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowlisted

`func (o *PolicyEvaluationFinding) SetAllowlisted(v bool)`

SetAllowlisted sets Allowlisted field to given value.


### GetAllowlistMatch

`func (o *PolicyEvaluationFinding) GetAllowlistMatch() PolicyFindingAllowlistMatch`

GetAllowlistMatch returns the AllowlistMatch field if non-nil, zero value otherwise.

### GetAllowlistMatchOk

`func (o *PolicyEvaluationFinding) GetAllowlistMatchOk() (*PolicyFindingAllowlistMatch, bool)`

GetAllowlistMatchOk returns a tuple with the AllowlistMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowlistMatch

`func (o *PolicyEvaluationFinding) SetAllowlistMatch(v PolicyFindingAllowlistMatch)`

SetAllowlistMatch sets AllowlistMatch field to given value.

### HasAllowlistMatch

`func (o *PolicyEvaluationFinding) HasAllowlistMatch() bool`

HasAllowlistMatch returns a boolean if a field has been set.

### SetAllowlistMatchNil

`func (o *PolicyEvaluationFinding) SetAllowlistMatchNil(b bool)`

 SetAllowlistMatchNil sets the value for AllowlistMatch to be an explicit nil

### UnsetAllowlistMatch
`func (o *PolicyEvaluationFinding) UnsetAllowlistMatch()`

UnsetAllowlistMatch ensures that no value is present for AllowlistMatch, not even an explicit nil
### GetInheritedFromBase

`func (o *PolicyEvaluationFinding) GetInheritedFromBase() bool`

GetInheritedFromBase returns the InheritedFromBase field if non-nil, zero value otherwise.

### GetInheritedFromBaseOk

`func (o *PolicyEvaluationFinding) GetInheritedFromBaseOk() (*bool, bool)`

GetInheritedFromBaseOk returns a tuple with the InheritedFromBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritedFromBase

`func (o *PolicyEvaluationFinding) SetInheritedFromBase(v bool)`

SetInheritedFromBase sets InheritedFromBase field to given value.

### HasInheritedFromBase

`func (o *PolicyEvaluationFinding) HasInheritedFromBase() bool`

HasInheritedFromBase returns a boolean if a field has been set.

### SetInheritedFromBaseNil

`func (o *PolicyEvaluationFinding) SetInheritedFromBaseNil(b bool)`

 SetInheritedFromBaseNil sets the value for InheritedFromBase to be an explicit nil

### UnsetInheritedFromBase
`func (o *PolicyEvaluationFinding) UnsetInheritedFromBase()`

UnsetInheritedFromBase ensures that no value is present for InheritedFromBase, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


