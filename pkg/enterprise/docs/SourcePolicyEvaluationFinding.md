# SourcePolicyEvaluationFinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TriggerId** | Pointer to **string** | ID of this policy trigger finding (can be used to allowlist this finding) | [optional] 
**Gate** | Pointer to **string** | Name of the gate that generated this finding | [optional] 
**Trigger** | Pointer to **string** | Name of the trigger that generated this finding | [optional] 
**Message** | Pointer to **string** | Description of the finding | [optional] 
**Action** | Pointer to **string** | The action associated with this finding | [optional] 
**PolicyId** | Pointer to **string** | ID of the policy that this gate and trigger are a part of | [optional] 
**Recommendation** | Pointer to **string** | User provided details for resolving this finding | [optional] 
**RuleId** | Pointer to **string** | ID of the policy rule that that generated this finding | [optional] 
**Allowlisted** | Pointer to **bool** | Indicates if this finding was allowlisted or not | [optional] 
**AllowlistMatch** | Pointer to [**AnyOfNullTypePolicyEvaluationFindingAllowlistMatch**](anyOf&lt;NullType,PolicyEvaluationFindingAllowlistMatch&gt;.md) |  | [optional] 

## Methods

### NewSourcePolicyEvaluationFinding

`func NewSourcePolicyEvaluationFinding() *SourcePolicyEvaluationFinding`

NewSourcePolicyEvaluationFinding instantiates a new SourcePolicyEvaluationFinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourcePolicyEvaluationFindingWithDefaults

`func NewSourcePolicyEvaluationFindingWithDefaults() *SourcePolicyEvaluationFinding`

NewSourcePolicyEvaluationFindingWithDefaults instantiates a new SourcePolicyEvaluationFinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTriggerId

`func (o *SourcePolicyEvaluationFinding) GetTriggerId() string`

GetTriggerId returns the TriggerId field if non-nil, zero value otherwise.

### GetTriggerIdOk

`func (o *SourcePolicyEvaluationFinding) GetTriggerIdOk() (*string, bool)`

GetTriggerIdOk returns a tuple with the TriggerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerId

`func (o *SourcePolicyEvaluationFinding) SetTriggerId(v string)`

SetTriggerId sets TriggerId field to given value.

### HasTriggerId

`func (o *SourcePolicyEvaluationFinding) HasTriggerId() bool`

HasTriggerId returns a boolean if a field has been set.

### GetGate

`func (o *SourcePolicyEvaluationFinding) GetGate() string`

GetGate returns the Gate field if non-nil, zero value otherwise.

### GetGateOk

`func (o *SourcePolicyEvaluationFinding) GetGateOk() (*string, bool)`

GetGateOk returns a tuple with the Gate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGate

`func (o *SourcePolicyEvaluationFinding) SetGate(v string)`

SetGate sets Gate field to given value.

### HasGate

`func (o *SourcePolicyEvaluationFinding) HasGate() bool`

HasGate returns a boolean if a field has been set.

### GetTrigger

`func (o *SourcePolicyEvaluationFinding) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *SourcePolicyEvaluationFinding) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *SourcePolicyEvaluationFinding) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *SourcePolicyEvaluationFinding) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetMessage

`func (o *SourcePolicyEvaluationFinding) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SourcePolicyEvaluationFinding) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SourcePolicyEvaluationFinding) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SourcePolicyEvaluationFinding) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetAction

`func (o *SourcePolicyEvaluationFinding) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SourcePolicyEvaluationFinding) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SourcePolicyEvaluationFinding) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *SourcePolicyEvaluationFinding) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetPolicyId

`func (o *SourcePolicyEvaluationFinding) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *SourcePolicyEvaluationFinding) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *SourcePolicyEvaluationFinding) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *SourcePolicyEvaluationFinding) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetRecommendation

`func (o *SourcePolicyEvaluationFinding) GetRecommendation() string`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *SourcePolicyEvaluationFinding) GetRecommendationOk() (*string, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *SourcePolicyEvaluationFinding) SetRecommendation(v string)`

SetRecommendation sets Recommendation field to given value.

### HasRecommendation

`func (o *SourcePolicyEvaluationFinding) HasRecommendation() bool`

HasRecommendation returns a boolean if a field has been set.

### GetRuleId

`func (o *SourcePolicyEvaluationFinding) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *SourcePolicyEvaluationFinding) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *SourcePolicyEvaluationFinding) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *SourcePolicyEvaluationFinding) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetAllowlisted

`func (o *SourcePolicyEvaluationFinding) GetAllowlisted() bool`

GetAllowlisted returns the Allowlisted field if non-nil, zero value otherwise.

### GetAllowlistedOk

`func (o *SourcePolicyEvaluationFinding) GetAllowlistedOk() (*bool, bool)`

GetAllowlistedOk returns a tuple with the Allowlisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowlisted

`func (o *SourcePolicyEvaluationFinding) SetAllowlisted(v bool)`

SetAllowlisted sets Allowlisted field to given value.

### HasAllowlisted

`func (o *SourcePolicyEvaluationFinding) HasAllowlisted() bool`

HasAllowlisted returns a boolean if a field has been set.

### GetAllowlistMatch

`func (o *SourcePolicyEvaluationFinding) GetAllowlistMatch() AnyOfNullTypePolicyEvaluationFindingAllowlistMatch`

GetAllowlistMatch returns the AllowlistMatch field if non-nil, zero value otherwise.

### GetAllowlistMatchOk

`func (o *SourcePolicyEvaluationFinding) GetAllowlistMatchOk() (*AnyOfNullTypePolicyEvaluationFindingAllowlistMatch, bool)`

GetAllowlistMatchOk returns a tuple with the AllowlistMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowlistMatch

`func (o *SourcePolicyEvaluationFinding) SetAllowlistMatch(v AnyOfNullTypePolicyEvaluationFindingAllowlistMatch)`

SetAllowlistMatch sets AllowlistMatch field to given value.

### HasAllowlistMatch

`func (o *SourcePolicyEvaluationFinding) HasAllowlistMatch() bool`

HasAllowlistMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


