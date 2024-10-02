/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.7.2
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"fmt"
)

// IntegrationHealthState Health state of the integration instance * HEALTHY: All good. No errors. * UNHEALTHY: Not well. Probably errors. 
type IntegrationHealthState string

// List of IntegrationHealthState
const (
	HEALTHY IntegrationHealthState = "HEALTHY"
	UNHEALTHY IntegrationHealthState = "UNHEALTHY"
)

// All allowed values of IntegrationHealthState enum
var AllowedIntegrationHealthStateEnumValues = []IntegrationHealthState{
	"HEALTHY",
	"UNHEALTHY",
}

func (v *IntegrationHealthState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IntegrationHealthState(value)
	for _, existing := range AllowedIntegrationHealthStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IntegrationHealthState", value)
}

// NewIntegrationHealthStateFromValue returns a pointer to a valid IntegrationHealthState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIntegrationHealthStateFromValue(v string) (*IntegrationHealthState, error) {
	ev := IntegrationHealthState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IntegrationHealthState: valid values are %v", v, AllowedIntegrationHealthStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IntegrationHealthState) IsValid() bool {
	for _, existing := range AllowedIntegrationHealthStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IntegrationHealthState value
func (v IntegrationHealthState) Ptr() *IntegrationHealthState {
	return &v
}

type NullableIntegrationHealthState struct {
	value *IntegrationHealthState
	isSet bool
}

func (v NullableIntegrationHealthState) Get() *IntegrationHealthState {
	return v.value
}

func (v *NullableIntegrationHealthState) Set(val *IntegrationHealthState) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationHealthState) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationHealthState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationHealthState(val *IntegrationHealthState) *NullableIntegrationHealthState {
	return &NullableIntegrationHealthState{value: val, isSet: true}
}

func (v NullableIntegrationHealthState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationHealthState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

