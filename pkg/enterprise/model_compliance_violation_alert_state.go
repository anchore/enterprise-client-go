/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.2.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// ComplianceViolationAlertState State change for an existing ComplianceViolationAlert
type ComplianceViolationAlertState struct {
	// The new state of the compliance violation alert
	State string `json:"state"`
}

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
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
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


