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

// TriggerSpec Definition of a trigger and its parameters
type TriggerSpec struct {
	// Name of the trigger as it would appear in a policy document
	Name *string `json:"name,omitempty"`
	// Trigger description for what it tests and when it will fire during evaluation
	Description *string `json:"description,omitempty"`
	// State of the trigger
	State *string `json:"state,omitempty"`
	// The name of another trigger that supersedes this on functionally if this is deprecated
	SupersededBy NullableString `json:"superseded_by,omitempty"`
	// The list of parameters that are valid for this trigger
	Parameters []TriggerParamSpec `json:"parameters,omitempty"`
}

// NewTriggerSpec instantiates a new TriggerSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerSpec() *TriggerSpec {
	this := TriggerSpec{}
	return &this
}

// NewTriggerSpecWithDefaults instantiates a new TriggerSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerSpecWithDefaults() *TriggerSpec {
	this := TriggerSpec{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TriggerSpec) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerSpec) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TriggerSpec) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TriggerSpec) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TriggerSpec) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerSpec) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TriggerSpec) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TriggerSpec) SetDescription(v string) {
	o.Description = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *TriggerSpec) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerSpec) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *TriggerSpec) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *TriggerSpec) SetState(v string) {
	o.State = &v
}

// GetSupersededBy returns the SupersededBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TriggerSpec) GetSupersededBy() string {
	if o == nil || o.SupersededBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.SupersededBy.Get()
}

// GetSupersededByOk returns a tuple with the SupersededBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TriggerSpec) GetSupersededByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SupersededBy.Get(), o.SupersededBy.IsSet()
}

// HasSupersededBy returns a boolean if a field has been set.
func (o *TriggerSpec) HasSupersededBy() bool {
	if o != nil && o.SupersededBy.IsSet() {
		return true
	}

	return false
}

// SetSupersededBy gets a reference to the given NullableString and assigns it to the SupersededBy field.
func (o *TriggerSpec) SetSupersededBy(v string) {
	o.SupersededBy.Set(&v)
}
// SetSupersededByNil sets the value for SupersededBy to be an explicit nil
func (o *TriggerSpec) SetSupersededByNil() {
	o.SupersededBy.Set(nil)
}

// UnsetSupersededBy ensures that no value is present for SupersededBy, not even an explicit nil
func (o *TriggerSpec) UnsetSupersededBy() {
	o.SupersededBy.Unset()
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *TriggerSpec) GetParameters() []TriggerParamSpec {
	if o == nil || o.Parameters == nil {
		var ret []TriggerParamSpec
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerSpec) GetParametersOk() ([]TriggerParamSpec, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *TriggerSpec) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []TriggerParamSpec and assigns it to the Parameters field.
func (o *TriggerSpec) SetParameters(v []TriggerParamSpec) {
	o.Parameters = v
}

func (o TriggerSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.SupersededBy.IsSet() {
		toSerialize["superseded_by"] = o.SupersededBy.Get()
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	return json.Marshal(toSerialize)
}

type NullableTriggerSpec struct {
	value *TriggerSpec
	isSet bool
}

func (v NullableTriggerSpec) Get() *TriggerSpec {
	return v.value
}

func (v *NullableTriggerSpec) Set(val *TriggerSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerSpec(val *TriggerSpec) *NullableTriggerSpec {
	return &NullableTriggerSpec{value: val, isSet: true}
}

func (v NullableTriggerSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


