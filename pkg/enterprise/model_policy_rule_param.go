/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.10.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PolicyRuleParam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyRuleParam{}

// PolicyRuleParam struct for PolicyRuleParam
type PolicyRuleParam struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

type _PolicyRuleParam PolicyRuleParam

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
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyRuleParam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *PolicyRuleParam) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"value",
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

	varPolicyRuleParam := _PolicyRuleParam{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPolicyRuleParam)

	if err != nil {
		return err
	}

	*o = PolicyRuleParam(varPolicyRuleParam)

	return err
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


