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

// checks if the ComplianceViolationAlertState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComplianceViolationAlertState{}

// ComplianceViolationAlertState State change for an existing ComplianceViolationAlert
type ComplianceViolationAlertState struct {
	// The new state of the compliance violation alert
	State string `json:"state"`
}

type _ComplianceViolationAlertState ComplianceViolationAlertState

// NewComplianceViolationAlertState instantiates a new ComplianceViolationAlertState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComplianceViolationAlertState(state string) *ComplianceViolationAlertState {
	this := ComplianceViolationAlertState{}
	this.State = state
	return &this
}

// NewComplianceViolationAlertStateWithDefaults instantiates a new ComplianceViolationAlertState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComplianceViolationAlertStateWithDefaults() *ComplianceViolationAlertState {
	this := ComplianceViolationAlertState{}
	return &this
}

// GetState returns the State field value
func (o *ComplianceViolationAlertState) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ComplianceViolationAlertState) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ComplianceViolationAlertState) SetState(v string) {
	o.State = v
}

func (o ComplianceViolationAlertState) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComplianceViolationAlertState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["state"] = o.State
	return toSerialize, nil
}

func (o *ComplianceViolationAlertState) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"state",
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

	varComplianceViolationAlertState := _ComplianceViolationAlertState{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varComplianceViolationAlertState)

	if err != nil {
		return err
	}

	*o = ComplianceViolationAlertState(varComplianceViolationAlertState)

	return err
}

type NullableComplianceViolationAlertState struct {
	value *ComplianceViolationAlertState
	isSet bool
}

func (v NullableComplianceViolationAlertState) Get() *ComplianceViolationAlertState {
	return v.value
}

func (v *NullableComplianceViolationAlertState) Set(val *ComplianceViolationAlertState) {
	v.value = val
	v.isSet = true
}

func (v NullableComplianceViolationAlertState) IsSet() bool {
	return v.isSet
}

func (v *NullableComplianceViolationAlertState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComplianceViolationAlertState(val *ComplianceViolationAlertState) *NullableComplianceViolationAlertState {
	return &NullableComplianceViolationAlertState{value: val, isSet: true}
}

func (v NullableComplianceViolationAlertState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComplianceViolationAlertState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


