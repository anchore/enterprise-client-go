/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.6.1
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// PolicyEvaluationRemediation struct for PolicyEvaluationRemediation
type PolicyEvaluationRemediation struct {
	// Anchore generated options for resolving a finding
	Suggestions []PolicyEvaluationRemediationSuggestion `json:"suggestions"`
	// List of trigger IDs that these remediation suggestions apply to
	TriggerIds []string `json:"trigger_ids"`
}

// NewPolicyEvaluationRemediation instantiates a new PolicyEvaluationRemediation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyEvaluationRemediation(suggestions []PolicyEvaluationRemediationSuggestion, triggerIds []string) *PolicyEvaluationRemediation {
	this := PolicyEvaluationRemediation{}
	this.Suggestions = suggestions
	this.TriggerIds = triggerIds
	return &this
}

// NewPolicyEvaluationRemediationWithDefaults instantiates a new PolicyEvaluationRemediation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyEvaluationRemediationWithDefaults() *PolicyEvaluationRemediation {
	this := PolicyEvaluationRemediation{}
	return &this
}

// GetSuggestions returns the Suggestions field value
func (o *PolicyEvaluationRemediation) GetSuggestions() []PolicyEvaluationRemediationSuggestion {
	if o == nil {
		var ret []PolicyEvaluationRemediationSuggestion
		return ret
	}

	return o.Suggestions
}

// GetSuggestionsOk returns a tuple with the Suggestions field value
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationRemediation) GetSuggestionsOk() ([]PolicyEvaluationRemediationSuggestion, bool) {
	if o == nil {
		return nil, false
	}
	return o.Suggestions, true
}

// SetSuggestions sets field value
func (o *PolicyEvaluationRemediation) SetSuggestions(v []PolicyEvaluationRemediationSuggestion) {
	o.Suggestions = v
}

// GetTriggerIds returns the TriggerIds field value
func (o *PolicyEvaluationRemediation) GetTriggerIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TriggerIds
}

// GetTriggerIdsOk returns a tuple with the TriggerIds field value
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationRemediation) GetTriggerIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TriggerIds, true
}

// SetTriggerIds sets field value
func (o *PolicyEvaluationRemediation) SetTriggerIds(v []string) {
	o.TriggerIds = v
}

func (o PolicyEvaluationRemediation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["suggestions"] = o.Suggestions
	}
	if true {
		toSerialize["trigger_ids"] = o.TriggerIds
	}
	return json.Marshal(toSerialize)
}

type NullablePolicyEvaluationRemediation struct {
	value *PolicyEvaluationRemediation
	isSet bool
}

func (v NullablePolicyEvaluationRemediation) Get() *PolicyEvaluationRemediation {
	return v.value
}

func (v *NullablePolicyEvaluationRemediation) Set(val *PolicyEvaluationRemediation) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyEvaluationRemediation) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyEvaluationRemediation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyEvaluationRemediation(val *PolicyEvaluationRemediation) *NullablePolicyEvaluationRemediation {
	return &NullablePolicyEvaluationRemediation{value: val, isSet: true}
}

func (v NullablePolicyEvaluationRemediation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyEvaluationRemediation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


