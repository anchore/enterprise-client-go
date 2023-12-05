/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.0.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"fmt"
)

// PolicyEvaluationFindingAllowlistMatch struct for PolicyEvaluationFindingAllowlistMatch
type PolicyEvaluationFindingAllowlistMatch struct {
	NullType *NullType
	PolicyFindingAllowlistMatch *PolicyFindingAllowlistMatch
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PolicyEvaluationFindingAllowlistMatch) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal JSON data into PolicyFindingAllowlistMatch
	err = json.Unmarshal(data, &dst.PolicyFindingAllowlistMatch);
	if err == nil {
		jsonPolicyFindingAllowlistMatch, _ := json.Marshal(dst.PolicyFindingAllowlistMatch)
		if string(jsonPolicyFindingAllowlistMatch) == "{}" { // empty struct
			dst.PolicyFindingAllowlistMatch = nil
		} else {
			return nil // data stored in dst.PolicyFindingAllowlistMatch, return on the first match
		}
	} else {
		dst.PolicyFindingAllowlistMatch = nil
	}

	return fmt.Errorf("Data failed to match schemas in anyOf(PolicyEvaluationFindingAllowlistMatch)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PolicyEvaluationFindingAllowlistMatch) MarshalJSON() ([]byte, error) {
	if src.NullType != nil {
		return json.Marshal(&src.NullType)
	}

	if src.PolicyFindingAllowlistMatch != nil {
		return json.Marshal(&src.PolicyFindingAllowlistMatch)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePolicyEvaluationFindingAllowlistMatch struct {
	value *PolicyEvaluationFindingAllowlistMatch
	isSet bool
}

func (v NullablePolicyEvaluationFindingAllowlistMatch) Get() *PolicyEvaluationFindingAllowlistMatch {
	return v.value
}

func (v *NullablePolicyEvaluationFindingAllowlistMatch) Set(val *PolicyEvaluationFindingAllowlistMatch) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyEvaluationFindingAllowlistMatch) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyEvaluationFindingAllowlistMatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyEvaluationFindingAllowlistMatch(val *PolicyEvaluationFindingAllowlistMatch) *NullablePolicyEvaluationFindingAllowlistMatch {
	return &NullablePolicyEvaluationFindingAllowlistMatch{value: val, isSet: true}
}

func (v NullablePolicyEvaluationFindingAllowlistMatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyEvaluationFindingAllowlistMatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


