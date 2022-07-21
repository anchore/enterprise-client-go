/*
Anchore Engine API Server

This is the Anchore Engine API. Provides the primary external API for users of the service.

API version: 0.1.20
Contact: nurmi@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package engine

import (
	"encoding/json"
)

// TriggerParamSpec struct for TriggerParamSpec
type TriggerParamSpec struct {
	// Parameter name as it appears in policy document
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	// An example value for the parameter (encoded as a string if the parameter is an object or list type)
	Example NullableString `json:"example,omitempty"`
	// State of the trigger parameter
	State *string `json:"state,omitempty"`
	// The name of another trigger that supercedes this on functionally if this is deprecated
	SupercededBy NullableString `json:"superceded_by,omitempty"`
	// Is this a required parameter or optional
	Required *bool `json:"required,omitempty"`
	// If present, a definition for validation of input. Typically a jsonschema object that can be used to validate an input against.
	Validator *interface{} `json:"validator,omitempty"`
}

// NewTriggerParamSpec instantiates a new TriggerParamSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerParamSpec() *TriggerParamSpec {
	this := TriggerParamSpec{}
	return &this
}

// NewTriggerParamSpecWithDefaults instantiates a new TriggerParamSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerParamSpecWithDefaults() *TriggerParamSpec {
	this := TriggerParamSpec{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TriggerParamSpec) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerParamSpec) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TriggerParamSpec) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TriggerParamSpec) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TriggerParamSpec) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerParamSpec) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TriggerParamSpec) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TriggerParamSpec) SetDescription(v string) {
	o.Description = &v
}

// GetExample returns the Example field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TriggerParamSpec) GetExample() string {
	if o == nil || o.Example.Get() == nil {
		var ret string
		return ret
	}
	return *o.Example.Get()
}

// GetExampleOk returns a tuple with the Example field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TriggerParamSpec) GetExampleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Example.Get(), o.Example.IsSet()
}

// HasExample returns a boolean if a field has been set.
func (o *TriggerParamSpec) HasExample() bool {
	if o != nil && o.Example.IsSet() {
		return true
	}

	return false
}

// SetExample gets a reference to the given NullableString and assigns it to the Example field.
func (o *TriggerParamSpec) SetExample(v string) {
	o.Example.Set(&v)
}
// SetExampleNil sets the value for Example to be an explicit nil
func (o *TriggerParamSpec) SetExampleNil() {
	o.Example.Set(nil)
}

// UnsetExample ensures that no value is present for Example, not even an explicit nil
func (o *TriggerParamSpec) UnsetExample() {
	o.Example.Unset()
}

// GetState returns the State field value if set, zero value otherwise.
func (o *TriggerParamSpec) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerParamSpec) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *TriggerParamSpec) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *TriggerParamSpec) SetState(v string) {
	o.State = &v
}

// GetSupercededBy returns the SupercededBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TriggerParamSpec) GetSupercededBy() string {
	if o == nil || o.SupercededBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.SupercededBy.Get()
}

// GetSupercededByOk returns a tuple with the SupercededBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TriggerParamSpec) GetSupercededByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SupercededBy.Get(), o.SupercededBy.IsSet()
}

// HasSupercededBy returns a boolean if a field has been set.
func (o *TriggerParamSpec) HasSupercededBy() bool {
	if o != nil && o.SupercededBy.IsSet() {
		return true
	}

	return false
}

// SetSupercededBy gets a reference to the given NullableString and assigns it to the SupercededBy field.
func (o *TriggerParamSpec) SetSupercededBy(v string) {
	o.SupercededBy.Set(&v)
}
// SetSupercededByNil sets the value for SupercededBy to be an explicit nil
func (o *TriggerParamSpec) SetSupercededByNil() {
	o.SupercededBy.Set(nil)
}

// UnsetSupercededBy ensures that no value is present for SupercededBy, not even an explicit nil
func (o *TriggerParamSpec) UnsetSupercededBy() {
	o.SupercededBy.Unset()
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *TriggerParamSpec) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerParamSpec) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *TriggerParamSpec) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *TriggerParamSpec) SetRequired(v bool) {
	o.Required = &v
}

// GetValidator returns the Validator field value if set, zero value otherwise.
func (o *TriggerParamSpec) GetValidator() interface{} {
	if o == nil || o.Validator == nil {
		var ret interface{}
		return ret
	}
	return *o.Validator
}

// GetValidatorOk returns a tuple with the Validator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerParamSpec) GetValidatorOk() (*interface{}, bool) {
	if o == nil || o.Validator == nil {
		return nil, false
	}
	return o.Validator, true
}

// HasValidator returns a boolean if a field has been set.
func (o *TriggerParamSpec) HasValidator() bool {
	if o != nil && o.Validator != nil {
		return true
	}

	return false
}

// SetValidator gets a reference to the given interface{} and assigns it to the Validator field.
func (o *TriggerParamSpec) SetValidator(v interface{}) {
	o.Validator = &v
}

func (o TriggerParamSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Example.IsSet() {
		toSerialize["example"] = o.Example.Get()
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.SupercededBy.IsSet() {
		toSerialize["superceded_by"] = o.SupercededBy.Get()
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	if o.Validator != nil {
		toSerialize["validator"] = o.Validator
	}
	return json.Marshal(toSerialize)
}

type NullableTriggerParamSpec struct {
	value *TriggerParamSpec
	isSet bool
}

func (v NullableTriggerParamSpec) Get() *TriggerParamSpec {
	return v.value
}

func (v *NullableTriggerParamSpec) Set(val *TriggerParamSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerParamSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerParamSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerParamSpec(val *TriggerParamSpec) *NullableTriggerParamSpec {
	return &NullableTriggerParamSpec{value: val, isSet: true}
}

func (v NullableTriggerParamSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerParamSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

