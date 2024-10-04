/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.7.2
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PolicyEvaluationFinding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyEvaluationFinding{}

// PolicyEvaluationFinding struct for PolicyEvaluationFinding
type PolicyEvaluationFinding struct {
	// ID of this policy trigger finding (can be used to allowlist this finding)
	TriggerId string `json:"trigger_id"`
	// Name of the gate that generated this finding
	Gate string `json:"gate"`
	// Name of the trigger that generated this finding
	Trigger string `json:"trigger"`
	// Description of the finding
	Message string `json:"message"`
	// The action associated with this finding
	Action string `json:"action"`
	// ID of the policy that this gate and trigger are a part of
	PolicyId string `json:"policy_id"`
	// User provided details for resolving this finding
	Recommendation string `json:"recommendation"`
	// ID of the policy rule that that generated this finding
	RuleId string `json:"rule_id"`
	// Indicates if this finding was allowlisted or not
	Allowlisted bool `json:"allowlisted"`
	AllowlistMatch NullablePolicyFindingAllowlistMatch `json:"allowlist_match,omitempty"`
	// Indicates if this finding was found in the base image
	InheritedFromBase NullableBool `json:"inherited_from_base,omitempty"`
}

type _PolicyEvaluationFinding PolicyEvaluationFinding

// NewPolicyEvaluationFinding instantiates a new PolicyEvaluationFinding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyEvaluationFinding(triggerId string, gate string, trigger string, message string, action string, policyId string, recommendation string, ruleId string, allowlisted bool) *PolicyEvaluationFinding {
	this := PolicyEvaluationFinding{}
	this.TriggerId = triggerId
	this.Gate = gate
	this.Trigger = trigger
	this.Message = message
	this.Action = action
	this.PolicyId = policyId
	this.Recommendation = recommendation
	this.RuleId = ruleId
	this.Allowlisted = allowlisted
	return &this
}

// NewPolicyEvaluationFindingWithDefaults instantiates a new PolicyEvaluationFinding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyEvaluationFindingWithDefaults() *PolicyEvaluationFinding {
	this := PolicyEvaluationFinding{}
	return &this
}

// GetTriggerId returns the TriggerId field value
func (o *PolicyEvaluationFinding) GetTriggerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TriggerId
}

// GetTriggerIdOk returns a tuple with the TriggerId field value
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationFinding) GetTriggerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggerId, true
}

// SetTriggerId sets field value
func (o *PolicyEvaluationFinding) SetTriggerId(v string) {
	o.TriggerId = v
}

// GetGate returns the Gate field value
func (o *PolicyEvaluationFinding) GetGate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gate
}

// GetGateOk returns a tuple with the Gate field value
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationFinding) GetGateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gate, true
}

// SetGate sets field value
func (o *PolicyEvaluationFinding) SetGate(v string) {
	o.Gate = v
}

// GetTrigger returns the Trigger field value
func (o *PolicyEvaluationFinding) GetTrigger() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationFinding) GetTriggerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Trigger, true
}

// SetTrigger sets field value
func (o *PolicyEvaluationFinding) SetTrigger(v string) {
	o.Trigger = v
}

// GetMessage returns the Message field value
func (o *PolicyEvaluationFinding) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationFinding) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *PolicyEvaluationFinding) SetMessage(v string) {
	o.Message = v
}

// GetAction returns the Action field value
func (o *PolicyEvaluationFinding) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationFinding) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *PolicyEvaluationFinding) SetAction(v string) {
	o.Action = v
}

// GetPolicyId returns the PolicyId field value
func (o *PolicyEvaluationFinding) GetPolicyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationFinding) GetPolicyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyId, true
}

// SetPolicyId sets field value
func (o *PolicyEvaluationFinding) SetPolicyId(v string) {
	o.PolicyId = v
}

// GetRecommendation returns the Recommendation field value
func (o *PolicyEvaluationFinding) GetRecommendation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Recommendation
}

// GetRecommendationOk returns a tuple with the Recommendation field value
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationFinding) GetRecommendationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recommendation, true
}

// SetRecommendation sets field value
func (o *PolicyEvaluationFinding) SetRecommendation(v string) {
	o.Recommendation = v
}

// GetRuleId returns the RuleId field value
func (o *PolicyEvaluationFinding) GetRuleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationFinding) GetRuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleId, true
}

// SetRuleId sets field value
func (o *PolicyEvaluationFinding) SetRuleId(v string) {
	o.RuleId = v
}

// GetAllowlisted returns the Allowlisted field value
func (o *PolicyEvaluationFinding) GetAllowlisted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Allowlisted
}

// GetAllowlistedOk returns a tuple with the Allowlisted field value
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationFinding) GetAllowlistedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Allowlisted, true
}

// SetAllowlisted sets field value
func (o *PolicyEvaluationFinding) SetAllowlisted(v bool) {
	o.Allowlisted = v
}

// GetAllowlistMatch returns the AllowlistMatch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyEvaluationFinding) GetAllowlistMatch() PolicyFindingAllowlistMatch {
	if o == nil || IsNil(o.AllowlistMatch.Get()) {
		var ret PolicyFindingAllowlistMatch
		return ret
	}
	return *o.AllowlistMatch.Get()
}

// GetAllowlistMatchOk returns a tuple with the AllowlistMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyEvaluationFinding) GetAllowlistMatchOk() (*PolicyFindingAllowlistMatch, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowlistMatch.Get(), o.AllowlistMatch.IsSet()
}

// HasAllowlistMatch returns a boolean if a field has been set.
func (o *PolicyEvaluationFinding) HasAllowlistMatch() bool {
	if o != nil && o.AllowlistMatch.IsSet() {
		return true
	}

	return false
}

// SetAllowlistMatch gets a reference to the given NullablePolicyFindingAllowlistMatch and assigns it to the AllowlistMatch field.
func (o *PolicyEvaluationFinding) SetAllowlistMatch(v PolicyFindingAllowlistMatch) {
	o.AllowlistMatch.Set(&v)
}
// SetAllowlistMatchNil sets the value for AllowlistMatch to be an explicit nil
func (o *PolicyEvaluationFinding) SetAllowlistMatchNil() {
	o.AllowlistMatch.Set(nil)
}

// UnsetAllowlistMatch ensures that no value is present for AllowlistMatch, not even an explicit nil
func (o *PolicyEvaluationFinding) UnsetAllowlistMatch() {
	o.AllowlistMatch.Unset()
}

// GetInheritedFromBase returns the InheritedFromBase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyEvaluationFinding) GetInheritedFromBase() bool {
	if o == nil || IsNil(o.InheritedFromBase.Get()) {
		var ret bool
		return ret
	}
	return *o.InheritedFromBase.Get()
}

// GetInheritedFromBaseOk returns a tuple with the InheritedFromBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyEvaluationFinding) GetInheritedFromBaseOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.InheritedFromBase.Get(), o.InheritedFromBase.IsSet()
}

// HasInheritedFromBase returns a boolean if a field has been set.
func (o *PolicyEvaluationFinding) HasInheritedFromBase() bool {
	if o != nil && o.InheritedFromBase.IsSet() {
		return true
	}

	return false
}

// SetInheritedFromBase gets a reference to the given NullableBool and assigns it to the InheritedFromBase field.
func (o *PolicyEvaluationFinding) SetInheritedFromBase(v bool) {
	o.InheritedFromBase.Set(&v)
}
// SetInheritedFromBaseNil sets the value for InheritedFromBase to be an explicit nil
func (o *PolicyEvaluationFinding) SetInheritedFromBaseNil() {
	o.InheritedFromBase.Set(nil)
}

// UnsetInheritedFromBase ensures that no value is present for InheritedFromBase, not even an explicit nil
func (o *PolicyEvaluationFinding) UnsetInheritedFromBase() {
	o.InheritedFromBase.Unset()
}

func (o PolicyEvaluationFinding) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyEvaluationFinding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["trigger_id"] = o.TriggerId
	toSerialize["gate"] = o.Gate
	toSerialize["trigger"] = o.Trigger
	toSerialize["message"] = o.Message
	toSerialize["action"] = o.Action
	toSerialize["policy_id"] = o.PolicyId
	toSerialize["recommendation"] = o.Recommendation
	toSerialize["rule_id"] = o.RuleId
	toSerialize["allowlisted"] = o.Allowlisted
	if o.AllowlistMatch.IsSet() {
		toSerialize["allowlist_match"] = o.AllowlistMatch.Get()
	}
	if o.InheritedFromBase.IsSet() {
		toSerialize["inherited_from_base"] = o.InheritedFromBase.Get()
	}
	return toSerialize, nil
}

func (o *PolicyEvaluationFinding) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"trigger_id",
		"gate",
		"trigger",
		"message",
		"action",
		"policy_id",
		"recommendation",
		"rule_id",
		"allowlisted",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPolicyEvaluationFinding := _PolicyEvaluationFinding{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPolicyEvaluationFinding)

	if err != nil {
		return err
	}

	*o = PolicyEvaluationFinding(varPolicyEvaluationFinding)

	return err
}

type NullablePolicyEvaluationFinding struct {
	value *PolicyEvaluationFinding
	isSet bool
}

func (v NullablePolicyEvaluationFinding) Get() *PolicyEvaluationFinding {
	return v.value
}

func (v *NullablePolicyEvaluationFinding) Set(val *PolicyEvaluationFinding) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyEvaluationFinding) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyEvaluationFinding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyEvaluationFinding(val *PolicyEvaluationFinding) *NullablePolicyEvaluationFinding {
	return &NullablePolicyEvaluationFinding{value: val, isSet: true}
}

func (v NullablePolicyEvaluationFinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyEvaluationFinding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


