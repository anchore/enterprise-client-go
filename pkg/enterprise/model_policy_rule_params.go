/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

<<<<<<< HEAD
API version: 2.0.0
=======
API version: 0.1.0
>>>>>>> 48fc108 (feat: updated the enterprise ref)
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// PolicyRuleParams struct for PolicyRuleParams
type PolicyRuleParams struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

// NewPolicyRuleParams instantiates a new PolicyRuleParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyRuleParams(name string, value string) *PolicyRuleParams {
	this := PolicyRuleParams{}
	this.Name = name
	this.Value = value
	return &this
}

// NewPolicyRuleParamsWithDefaults instantiates a new PolicyRuleParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyRuleParamsWithDefaults() *PolicyRuleParams {
	this := PolicyRuleParams{}
	return &this
}

// GetName returns the Name field value
func (o *PolicyRuleParams) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PolicyRuleParams) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PolicyRuleParams) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *PolicyRuleParams) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *PolicyRuleParams) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *PolicyRuleParams) SetValue(v string) {
	o.Value = v
}

func (o PolicyRuleParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullablePolicyRuleParams struct {
	value *PolicyRuleParams
	isSet bool
}

func (v NullablePolicyRuleParams) Get() *PolicyRuleParams {
	return v.value
}

func (v *NullablePolicyRuleParams) Set(val *PolicyRuleParams) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyRuleParams) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyRuleParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyRuleParams(val *PolicyRuleParams) *NullablePolicyRuleParams {
	return &NullablePolicyRuleParams{value: val, isSet: true}
}

func (v NullablePolicyRuleParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyRuleParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


