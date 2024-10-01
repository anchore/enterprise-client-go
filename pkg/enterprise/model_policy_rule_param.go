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

// PolicyRuleParam struct for PolicyRuleParam
type PolicyRuleParam struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

// NewPolicyRuleParam instantiates a new PolicyRuleParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyRuleParam(name string, value string) *PolicyRuleParam {
	this := PolicyRuleParam{}
	this.Name = name
	this.Value = value
	return &this
}

// NewPolicyRuleParamWithDefaults instantiates a new PolicyRuleParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyRuleParamWithDefaults() *PolicyRuleParam {
	this := PolicyRuleParam{}
	return &this
}

// GetName returns the Name field value
func (o *PolicyRuleParam) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PolicyRuleParam) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PolicyRuleParam) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *PolicyRuleParam) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *PolicyRuleParam) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *PolicyRuleParam) SetValue(v string) {
	o.Value = v
}

func (o PolicyRuleParam) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullablePolicyRuleParam struct {
	value *PolicyRuleParam
	isSet bool
}

func (v NullablePolicyRuleParam) Get() *PolicyRuleParam {
	return v.value
}

func (v *NullablePolicyRuleParam) Set(val *PolicyRuleParam) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyRuleParam) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyRuleParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyRuleParam(val *PolicyRuleParam) *NullablePolicyRuleParam {
	return &NullablePolicyRuleParam{value: val, isSet: true}
}

func (v NullablePolicyRuleParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyRuleParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


