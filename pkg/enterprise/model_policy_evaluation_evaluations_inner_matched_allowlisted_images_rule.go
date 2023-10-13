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

// PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule struct for PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule
type PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule struct {
	ImageSelectionRule *ImageSelectionRule
	NullType *NullType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ImageSelectionRule
	err = json.Unmarshal(data, &dst.ImageSelectionRule);
	if err == nil {
		jsonImageSelectionRule, _ := json.Marshal(dst.ImageSelectionRule)
		if string(jsonImageSelectionRule) == "{}" { // empty struct
			dst.ImageSelectionRule = nil
		} else {
			return nil // data stored in dst.ImageSelectionRule, return on the first match
		}
	} else {
		dst.ImageSelectionRule = nil
	}

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

	return fmt.Errorf("Data failed to match schemas in anyOf(PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule) MarshalJSON() ([]byte, error) {
	if src.ImageSelectionRule != nil {
		return json.Marshal(&src.ImageSelectionRule)
	}

	if src.NullType != nil {
		return json.Marshal(&src.NullType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule struct {
	value *PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule
	isSet bool
}

func (v NullablePolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule) Get() *PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule {
	return v.value
}

func (v *NullablePolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule) Set(val *PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule(val *PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule) *NullablePolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule {
	return &NullablePolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule{value: val, isSet: true}
}

func (v NullablePolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


