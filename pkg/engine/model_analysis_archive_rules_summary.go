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
	"time"
)

// AnalysisArchiveRulesSummary Summary of the transition rule set
type AnalysisArchiveRulesSummary struct {
	// The number of rules for this account
	Count *int32 `json:"count,omitempty"`
	// The newest last_updated timestamp from the set of rules
	LastUpdated *time.Time `json:"last_updated,omitempty"`
}

// NewAnalysisArchiveRulesSummary instantiates a new AnalysisArchiveRulesSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalysisArchiveRulesSummary() *AnalysisArchiveRulesSummary {
	this := AnalysisArchiveRulesSummary{}
	return &this
}

// NewAnalysisArchiveRulesSummaryWithDefaults instantiates a new AnalysisArchiveRulesSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalysisArchiveRulesSummaryWithDefaults() *AnalysisArchiveRulesSummary {
	this := AnalysisArchiveRulesSummary{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *AnalysisArchiveRulesSummary) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisArchiveRulesSummary) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *AnalysisArchiveRulesSummary) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *AnalysisArchiveRulesSummary) SetCount(v int32) {
	o.Count = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *AnalysisArchiveRulesSummary) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisArchiveRulesSummary) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *AnalysisArchiveRulesSummary) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *AnalysisArchiveRulesSummary) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o AnalysisArchiveRulesSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	return json.Marshal(toSerialize)
}

type NullableAnalysisArchiveRulesSummary struct {
	value *AnalysisArchiveRulesSummary
	isSet bool
}

func (v NullableAnalysisArchiveRulesSummary) Get() *AnalysisArchiveRulesSummary {
	return v.value
}

func (v *NullableAnalysisArchiveRulesSummary) Set(val *AnalysisArchiveRulesSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalysisArchiveRulesSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalysisArchiveRulesSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalysisArchiveRulesSummary(val *AnalysisArchiveRulesSummary) *NullableAnalysisArchiveRulesSummary {
	return &NullableAnalysisArchiveRulesSummary{value: val, isSet: true}
}

func (v NullableAnalysisArchiveRulesSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalysisArchiveRulesSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


