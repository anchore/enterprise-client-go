/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.5.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// SourcePolicyEvaluationFinding struct for SourcePolicyEvaluationFinding
type SourcePolicyEvaluationFinding struct {
	// ID of this policy trigger finding (can be used to allowlist this finding)
	TriggerId *string `json:"trigger_id,omitempty"`
	// Name of the gate that generated this finding
	Gate *string `json:"gate,omitempty"`
	// Name of the trigger that generated this finding
	Trigger *string `json:"trigger,omitempty"`
	// Description of the finding
	Message *string `json:"message,omitempty"`
	// The action associated with this finding
	Action *string `json:"action,omitempty"`
	// ID of the policy that this gate and trigger are a part of
	PolicyId *string `json:"policy_id,omitempty"`
	// User provided details for resolving this finding
	Recommendation *string `json:"recommendation,omitempty"`
	// ID of the policy rule that that generated this finding
	RuleId *string `json:"rule_id,omitempty"`
	// Indicates if this finding was allowlisted or not
	Allowlisted *bool `json:"allowlisted,omitempty"`
	AllowlistMatch *PolicyEvaluationFindingAllowlistMatch `json:"allowlist_match,omitempty"`
}

// NewSourcePolicyEvaluationFinding instantiates a new SourcePolicyEvaluationFinding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourcePolicyEvaluationFinding() *SourcePolicyEvaluationFinding {
	this := SourcePolicyEvaluationFinding{}
	return &this
}

// NewSourcePolicyEvaluationFindingWithDefaults instantiates a new SourcePolicyEvaluationFinding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourcePolicyEvaluationFindingWithDefaults() *SourcePolicyEvaluationFinding {
	this := SourcePolicyEvaluationFinding{}
	return &this
}

// GetTriggerId returns the TriggerId field value if set, zero value otherwise.
func (o *SourcePolicyEvaluationFinding) GetTriggerId() string {
	if o == nil || o.TriggerId == nil {
		var ret string
		return ret
	}
	return *o.TriggerId
}

// GetTriggerIdOk returns a tuple with the TriggerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcePolicyEvaluationFinding) GetTriggerIdOk() (*string, bool) {
	if o == nil || o.TriggerId == nil {
		return nil, false
	}
	return o.TriggerId, true
}

// HasTriggerId returns a boolean if a field has been set.
func (o *SourcePolicyEvaluationFinding) HasTriggerId() bool {
	if o != nil && o.TriggerId != nil {
		return true
	}

	return false
}

// SetTriggerId gets a reference to the given string and assigns it to the TriggerId field.
func (o *SourcePolicyEvaluationFinding) SetTriggerId(v string) {
	o.TriggerId = &v
}

// GetGate returns the Gate field value if set, zero value otherwise.
func (o *SourcePolicyEvaluationFinding) GetGate() string {
	if o == nil || o.Gate == nil {
		var ret string
		return ret
	}
	return *o.Gate
}

// GetGateOk returns a tuple with the Gate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcePolicyEvaluationFinding) GetGateOk() (*string, bool) {
	if o == nil || o.Gate == nil {
		return nil, false
	}
	return o.Gate, true
}

// HasGate returns a boolean if a field has been set.
func (o *SourcePolicyEvaluationFinding) HasGate() bool {
	if o != nil && o.Gate != nil {
		return true
	}

	return false
}

// SetGate gets a reference to the given string and assigns it to the Gate field.
func (o *SourcePolicyEvaluationFinding) SetGate(v string) {
	o.Gate = &v
}

// GetTrigger returns the Trigger field value if set, zero value otherwise.
func (o *SourcePolicyEvaluationFinding) GetTrigger() string {
	if o == nil || o.Trigger == nil {
		var ret string
		return ret
	}
	return *o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcePolicyEvaluationFinding) GetTriggerOk() (*string, bool) {
	if o == nil || o.Trigger == nil {
		return nil, false
	}
	return o.Trigger, true
}

// HasTrigger returns a boolean if a field has been set.
func (o *SourcePolicyEvaluationFinding) HasTrigger() bool {
	if o != nil && o.Trigger != nil {
		return true
	}

	return false
}

// SetTrigger gets a reference to the given string and assigns it to the Trigger field.
func (o *SourcePolicyEvaluationFinding) SetTrigger(v string) {
	o.Trigger = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SourcePolicyEvaluationFinding) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcePolicyEvaluationFinding) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SourcePolicyEvaluationFinding) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SourcePolicyEvaluationFinding) SetMessage(v string) {
	o.Message = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *SourcePolicyEvaluationFinding) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcePolicyEvaluationFinding) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *SourcePolicyEvaluationFinding) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *SourcePolicyEvaluationFinding) SetAction(v string) {
	o.Action = &v
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *SourcePolicyEvaluationFinding) GetPolicyId() string {
	if o == nil || o.PolicyId == nil {
		var ret string
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcePolicyEvaluationFinding) GetPolicyIdOk() (*string, bool) {
	if o == nil || o.PolicyId == nil {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *SourcePolicyEvaluationFinding) HasPolicyId() bool {
	if o != nil && o.PolicyId != nil {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given string and assigns it to the PolicyId field.
func (o *SourcePolicyEvaluationFinding) SetPolicyId(v string) {
	o.PolicyId = &v
}

// GetRecommendation returns the Recommendation field value if set, zero value otherwise.
func (o *SourcePolicyEvaluationFinding) GetRecommendation() string {
	if o == nil || o.Recommendation == nil {
		var ret string
		return ret
	}
	return *o.Recommendation
}

// GetRecommendationOk returns a tuple with the Recommendation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcePolicyEvaluationFinding) GetRecommendationOk() (*string, bool) {
	if o == nil || o.Recommendation == nil {
		return nil, false
	}
	return o.Recommendation, true
}

// HasRecommendation returns a boolean if a field has been set.
func (o *SourcePolicyEvaluationFinding) HasRecommendation() bool {
	if o != nil && o.Recommendation != nil {
		return true
	}

	return false
}

// SetRecommendation gets a reference to the given string and assigns it to the Recommendation field.
func (o *SourcePolicyEvaluationFinding) SetRecommendation(v string) {
	o.Recommendation = &v
}

// GetRuleId returns the RuleId field value if set, zero value otherwise.
func (o *SourcePolicyEvaluationFinding) GetRuleId() string {
	if o == nil || o.RuleId == nil {
		var ret string
		return ret
	}
	return *o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcePolicyEvaluationFinding) GetRuleIdOk() (*string, bool) {
	if o == nil || o.RuleId == nil {
		return nil, false
	}
	return o.RuleId, true
}

// HasRuleId returns a boolean if a field has been set.
func (o *SourcePolicyEvaluationFinding) HasRuleId() bool {
	if o != nil && o.RuleId != nil {
		return true
	}

	return false
}

// SetRuleId gets a reference to the given string and assigns it to the RuleId field.
func (o *SourcePolicyEvaluationFinding) SetRuleId(v string) {
	o.RuleId = &v
}

// GetAllowlisted returns the Allowlisted field value if set, zero value otherwise.
func (o *SourcePolicyEvaluationFinding) GetAllowlisted() bool {
	if o == nil || o.Allowlisted == nil {
		var ret bool
		return ret
	}
	return *o.Allowlisted
}

// GetAllowlistedOk returns a tuple with the Allowlisted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcePolicyEvaluationFinding) GetAllowlistedOk() (*bool, bool) {
	if o == nil || o.Allowlisted == nil {
		return nil, false
	}
	return o.Allowlisted, true
}

// HasAllowlisted returns a boolean if a field has been set.
func (o *SourcePolicyEvaluationFinding) HasAllowlisted() bool {
	if o != nil && o.Allowlisted != nil {
		return true
	}

	return false
}

// SetAllowlisted gets a reference to the given bool and assigns it to the Allowlisted field.
func (o *SourcePolicyEvaluationFinding) SetAllowlisted(v bool) {
	o.Allowlisted = &v
}

// GetAllowlistMatch returns the AllowlistMatch field value if set, zero value otherwise.
func (o *SourcePolicyEvaluationFinding) GetAllowlistMatch() PolicyEvaluationFindingAllowlistMatch {
	if o == nil || o.AllowlistMatch == nil {
		var ret PolicyEvaluationFindingAllowlistMatch
		return ret
	}
	return *o.AllowlistMatch
}

// GetAllowlistMatchOk returns a tuple with the AllowlistMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcePolicyEvaluationFinding) GetAllowlistMatchOk() (*PolicyEvaluationFindingAllowlistMatch, bool) {
	if o == nil || o.AllowlistMatch == nil {
		return nil, false
	}
	return o.AllowlistMatch, true
}

// HasAllowlistMatch returns a boolean if a field has been set.
func (o *SourcePolicyEvaluationFinding) HasAllowlistMatch() bool {
	if o != nil && o.AllowlistMatch != nil {
		return true
	}

	return false
}

// SetAllowlistMatch gets a reference to the given PolicyEvaluationFindingAllowlistMatch and assigns it to the AllowlistMatch field.
func (o *SourcePolicyEvaluationFinding) SetAllowlistMatch(v PolicyEvaluationFindingAllowlistMatch) {
	o.AllowlistMatch = &v
}

func (o SourcePolicyEvaluationFinding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TriggerId != nil {
		toSerialize["trigger_id"] = o.TriggerId
	}
	if o.Gate != nil {
		toSerialize["gate"] = o.Gate
	}
	if o.Trigger != nil {
		toSerialize["trigger"] = o.Trigger
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.PolicyId != nil {
		toSerialize["policy_id"] = o.PolicyId
	}
	if o.Recommendation != nil {
		toSerialize["recommendation"] = o.Recommendation
	}
	if o.RuleId != nil {
		toSerialize["rule_id"] = o.RuleId
	}
	if o.Allowlisted != nil {
		toSerialize["allowlisted"] = o.Allowlisted
	}
	if o.AllowlistMatch != nil {
		toSerialize["allowlist_match"] = o.AllowlistMatch
	}
	return json.Marshal(toSerialize)
}

type NullableSourcePolicyEvaluationFinding struct {
	value *SourcePolicyEvaluationFinding
	isSet bool
}

func (v NullableSourcePolicyEvaluationFinding) Get() *SourcePolicyEvaluationFinding {
	return v.value
}

func (v *NullableSourcePolicyEvaluationFinding) Set(val *SourcePolicyEvaluationFinding) {
	v.value = val
	v.isSet = true
}

func (v NullableSourcePolicyEvaluationFinding) IsSet() bool {
	return v.isSet
}

func (v *NullableSourcePolicyEvaluationFinding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourcePolicyEvaluationFinding(val *SourcePolicyEvaluationFinding) *NullableSourcePolicyEvaluationFinding {
	return &NullableSourcePolicyEvaluationFinding{value: val, isSet: true}
}

func (v NullableSourcePolicyEvaluationFinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourcePolicyEvaluationFinding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


