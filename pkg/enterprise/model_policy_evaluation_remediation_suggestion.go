/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.4.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// PolicyEvaluationRemediationSuggestion struct for PolicyEvaluationRemediationSuggestion
type PolicyEvaluationRemediationSuggestion struct {
	// The suggestion for resolving a finding
	Message string `json:"message"`
	// Indicates whether this suggestion is recommended
	Preferred bool `json:"preferred"`
}

// NewPolicyEvaluationRemediationSuggestion instantiates a new PolicyEvaluationRemediationSuggestion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyEvaluationRemediationSuggestion(message string, preferred bool) *PolicyEvaluationRemediationSuggestion {
	this := PolicyEvaluationRemediationSuggestion{}
	this.Message = message
	this.Preferred = preferred
	return &this
}

// NewPolicyEvaluationRemediationSuggestionWithDefaults instantiates a new PolicyEvaluationRemediationSuggestion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyEvaluationRemediationSuggestionWithDefaults() *PolicyEvaluationRemediationSuggestion {
	this := PolicyEvaluationRemediationSuggestion{}
	return &this
}

// GetMessage returns the Message field value
func (o *PolicyEvaluationRemediationSuggestion) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationRemediationSuggestion) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *PolicyEvaluationRemediationSuggestion) SetMessage(v string) {
	o.Message = v
}

// GetPreferred returns the Preferred field value
func (o *PolicyEvaluationRemediationSuggestion) GetPreferred() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Preferred
}

// GetPreferredOk returns a tuple with the Preferred field value
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationRemediationSuggestion) GetPreferredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Preferred, true
}

// SetPreferred sets field value
func (o *PolicyEvaluationRemediationSuggestion) SetPreferred(v bool) {
	o.Preferred = v
}

func (o PolicyEvaluationRemediationSuggestion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["preferred"] = o.Preferred
	}
	return json.Marshal(toSerialize)
}

type NullablePolicyEvaluationRemediationSuggestion struct {
	value *PolicyEvaluationRemediationSuggestion
	isSet bool
}

func (v NullablePolicyEvaluationRemediationSuggestion) Get() *PolicyEvaluationRemediationSuggestion {
	return v.value
}

func (v *NullablePolicyEvaluationRemediationSuggestion) Set(val *PolicyEvaluationRemediationSuggestion) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyEvaluationRemediationSuggestion) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyEvaluationRemediationSuggestion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyEvaluationRemediationSuggestion(val *PolicyEvaluationRemediationSuggestion) *NullablePolicyEvaluationRemediationSuggestion {
	return &NullablePolicyEvaluationRemediationSuggestion{value: val, isSet: true}
}

func (v NullablePolicyEvaluationRemediationSuggestion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyEvaluationRemediationSuggestion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


