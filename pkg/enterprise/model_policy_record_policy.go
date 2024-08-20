/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.6.2
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"fmt"
)

// PolicyRecordPolicy struct for PolicyRecordPolicy
type PolicyRecordPolicy struct {
	NullType *NullType
	Policy *Policy
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PolicyRecordPolicy) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into NullType
	err = json.Unmarshal(data, &dst.NullType);
	if err == nil {
		jsonNullType, _ := json.Marshal(dst.NullType)
		if string(jsonNullType) == "{}" { // empty struct
			dst.NullType = nil
		} else {
			return nil // data stored in dst.NullType, return on the first match
		}
	} else {
		dst.NullType = nil
	}

	// try to unmarshal JSON data into Policy
	err = json.Unmarshal(data, &dst.Policy);
	if err == nil {
		jsonPolicy, _ := json.Marshal(dst.Policy)
		if string(jsonPolicy) == "{}" { // empty struct
			dst.Policy = nil
		} else {
			return nil // data stored in dst.Policy, return on the first match
		}
	} else {
		dst.Policy = nil
	}

	return fmt.Errorf("Data failed to match schemas in anyOf(PolicyRecordPolicy)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PolicyRecordPolicy) MarshalJSON() ([]byte, error) {
	if src.NullType != nil {
		return json.Marshal(&src.NullType)
	}

	if src.Policy != nil {
		return json.Marshal(&src.Policy)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePolicyRecordPolicy struct {
	value *PolicyRecordPolicy
	isSet bool
}

func (v NullablePolicyRecordPolicy) Get() *PolicyRecordPolicy {
	return v.value
}

func (v *NullablePolicyRecordPolicy) Set(val *PolicyRecordPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyRecordPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyRecordPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyRecordPolicy(val *PolicyRecordPolicy) *NullablePolicyRecordPolicy {
	return &NullablePolicyRecordPolicy{value: val, isSet: true}
}

func (v NullablePolicyRecordPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyRecordPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


