/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.12.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PolicyEvaluationRemediation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyEvaluationRemediation{}

// PolicyEvaluationRemediation struct for PolicyEvaluationRemediation
type PolicyEvaluationRemediation struct {
	// Anchore generated options for resolving a finding
	Suggestions []PolicyEvaluationRemediationSuggestion `json:"suggestions"`
	// List of trigger IDs that these remediation suggestions apply to
	TriggerIds []string `json:"trigger_ids"`
}

type _PolicyEvaluationRemediation PolicyEvaluationRemediation

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
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyEvaluationRemediation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["suggestions"] = o.Suggestions
	toSerialize["trigger_ids"] = o.TriggerIds
	return toSerialize, nil
}

func (o *PolicyEvaluationRemediation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"suggestions",
		"trigger_ids",
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

	varPolicyEvaluationRemediation := _PolicyEvaluationRemediation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPolicyEvaluationRemediation)

	if err != nil {
		return err
	}

	*o = PolicyEvaluationRemediation(varPolicyEvaluationRemediation)

	return err
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


