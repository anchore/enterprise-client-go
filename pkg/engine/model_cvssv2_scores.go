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

// CVSSV2Scores struct for CVSSV2Scores
type CVSSV2Scores struct {
	BaseScore NullableFloat32 `json:"base_score,omitempty"`
	ExploitabilityScore NullableFloat32 `json:"exploitability_score,omitempty"`
	ImpactScore NullableFloat32 `json:"impact_score,omitempty"`
}

// NewCVSSV2Scores instantiates a new CVSSV2Scores object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCVSSV2Scores() *CVSSV2Scores {
	this := CVSSV2Scores{}
	return &this
}

// NewCVSSV2ScoresWithDefaults instantiates a new CVSSV2Scores object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCVSSV2ScoresWithDefaults() *CVSSV2Scores {
	this := CVSSV2Scores{}
	return &this
}

// GetBaseScore returns the BaseScore field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CVSSV2Scores) GetBaseScore() float32 {
	if o == nil || o.BaseScore.Get() == nil {
		var ret float32
		return ret
	}
	return *o.BaseScore.Get()
}

// GetBaseScoreOk returns a tuple with the BaseScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CVSSV2Scores) GetBaseScoreOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BaseScore.Get(), o.BaseScore.IsSet()
}

// HasBaseScore returns a boolean if a field has been set.
func (o *CVSSV2Scores) HasBaseScore() bool {
	if o != nil && o.BaseScore.IsSet() {
		return true
	}

	return false
}

// SetBaseScore gets a reference to the given NullableFloat32 and assigns it to the BaseScore field.
func (o *CVSSV2Scores) SetBaseScore(v float32) {
	o.BaseScore.Set(&v)
}
// SetBaseScoreNil sets the value for BaseScore to be an explicit nil
func (o *CVSSV2Scores) SetBaseScoreNil() {
	o.BaseScore.Set(nil)
}

// UnsetBaseScore ensures that no value is present for BaseScore, not even an explicit nil
func (o *CVSSV2Scores) UnsetBaseScore() {
	o.BaseScore.Unset()
}

// GetExploitabilityScore returns the ExploitabilityScore field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CVSSV2Scores) GetExploitabilityScore() float32 {
	if o == nil || o.ExploitabilityScore.Get() == nil {
		var ret float32
		return ret
	}
	return *o.ExploitabilityScore.Get()
}

// GetExploitabilityScoreOk returns a tuple with the ExploitabilityScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CVSSV2Scores) GetExploitabilityScoreOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExploitabilityScore.Get(), o.ExploitabilityScore.IsSet()
}

// HasExploitabilityScore returns a boolean if a field has been set.
func (o *CVSSV2Scores) HasExploitabilityScore() bool {
	if o != nil && o.ExploitabilityScore.IsSet() {
		return true
	}

	return false
}

// SetExploitabilityScore gets a reference to the given NullableFloat32 and assigns it to the ExploitabilityScore field.
func (o *CVSSV2Scores) SetExploitabilityScore(v float32) {
	o.ExploitabilityScore.Set(&v)
}
// SetExploitabilityScoreNil sets the value for ExploitabilityScore to be an explicit nil
func (o *CVSSV2Scores) SetExploitabilityScoreNil() {
	o.ExploitabilityScore.Set(nil)
}

// UnsetExploitabilityScore ensures that no value is present for ExploitabilityScore, not even an explicit nil
func (o *CVSSV2Scores) UnsetExploitabilityScore() {
	o.ExploitabilityScore.Unset()
}

// GetImpactScore returns the ImpactScore field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CVSSV2Scores) GetImpactScore() float32 {
	if o == nil || o.ImpactScore.Get() == nil {
		var ret float32
		return ret
	}
	return *o.ImpactScore.Get()
}

// GetImpactScoreOk returns a tuple with the ImpactScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CVSSV2Scores) GetImpactScoreOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ImpactScore.Get(), o.ImpactScore.IsSet()
}

// HasImpactScore returns a boolean if a field has been set.
func (o *CVSSV2Scores) HasImpactScore() bool {
	if o != nil && o.ImpactScore.IsSet() {
		return true
	}

	return false
}

// SetImpactScore gets a reference to the given NullableFloat32 and assigns it to the ImpactScore field.
func (o *CVSSV2Scores) SetImpactScore(v float32) {
	o.ImpactScore.Set(&v)
}
// SetImpactScoreNil sets the value for ImpactScore to be an explicit nil
func (o *CVSSV2Scores) SetImpactScoreNil() {
	o.ImpactScore.Set(nil)
}

// UnsetImpactScore ensures that no value is present for ImpactScore, not even an explicit nil
func (o *CVSSV2Scores) UnsetImpactScore() {
	o.ImpactScore.Unset()
}

func (o CVSSV2Scores) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BaseScore.IsSet() {
		toSerialize["base_score"] = o.BaseScore.Get()
	}
	if o.ExploitabilityScore.IsSet() {
		toSerialize["exploitability_score"] = o.ExploitabilityScore.Get()
	}
	if o.ImpactScore.IsSet() {
		toSerialize["impact_score"] = o.ImpactScore.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCVSSV2Scores struct {
	value *CVSSV2Scores
	isSet bool
}

func (v NullableCVSSV2Scores) Get() *CVSSV2Scores {
	return v.value
}

func (v *NullableCVSSV2Scores) Set(val *CVSSV2Scores) {
	v.value = val
	v.isSet = true
}

func (v NullableCVSSV2Scores) IsSet() bool {
	return v.isSet
}

func (v *NullableCVSSV2Scores) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCVSSV2Scores(val *CVSSV2Scores) *NullableCVSSV2Scores {
	return &NullableCVSSV2Scores{value: val, isSet: true}
}

func (v NullableCVSSV2Scores) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCVSSV2Scores) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


