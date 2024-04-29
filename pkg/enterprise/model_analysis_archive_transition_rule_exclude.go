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

// AnalysisArchiveTransitionRuleExclude Which Images to exclude from auto-archiving logic
type AnalysisArchiveTransitionRuleExclude struct {
	Selector *ImageSelector `json:"selector,omitempty"`
	// How long the image selected will be excluded from the archive transition
	ExpirationDays *int32 `json:"expiration_days,omitempty"`
	// Exclude image from archive if last seen in inventory within defined number of days
	LastSeenInDays *int32 `json:"last_seen_in_days,omitempty"`
}

// NewAnalysisArchiveTransitionRuleExclude instantiates a new AnalysisArchiveTransitionRuleExclude object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalysisArchiveTransitionRuleExclude() *AnalysisArchiveTransitionRuleExclude {
	this := AnalysisArchiveTransitionRuleExclude{}
	return &this
}

// NewAnalysisArchiveTransitionRuleExcludeWithDefaults instantiates a new AnalysisArchiveTransitionRuleExclude object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalysisArchiveTransitionRuleExcludeWithDefaults() *AnalysisArchiveTransitionRuleExclude {
	this := AnalysisArchiveTransitionRuleExclude{}
	return &this
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *AnalysisArchiveTransitionRuleExclude) GetSelector() ImageSelector {
	if o == nil || o.Selector == nil {
		var ret ImageSelector
		return ret
	}
	return *o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisArchiveTransitionRuleExclude) GetSelectorOk() (*ImageSelector, bool) {
	if o == nil || o.Selector == nil {
		return nil, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *AnalysisArchiveTransitionRuleExclude) HasSelector() bool {
	if o != nil && o.Selector != nil {
		return true
	}

	return false
}

// SetSelector gets a reference to the given ImageSelector and assigns it to the Selector field.
func (o *AnalysisArchiveTransitionRuleExclude) SetSelector(v ImageSelector) {
	o.Selector = &v
}

// GetExpirationDays returns the ExpirationDays field value if set, zero value otherwise.
func (o *AnalysisArchiveTransitionRuleExclude) GetExpirationDays() int32 {
	if o == nil || o.ExpirationDays == nil {
		var ret int32
		return ret
	}
	return *o.ExpirationDays
}

// GetExpirationDaysOk returns a tuple with the ExpirationDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisArchiveTransitionRuleExclude) GetExpirationDaysOk() (*int32, bool) {
	if o == nil || o.ExpirationDays == nil {
		return nil, false
	}
	return o.ExpirationDays, true
}

// HasExpirationDays returns a boolean if a field has been set.
func (o *AnalysisArchiveTransitionRuleExclude) HasExpirationDays() bool {
	if o != nil && o.ExpirationDays != nil {
		return true
	}

	return false
}

// SetExpirationDays gets a reference to the given int32 and assigns it to the ExpirationDays field.
func (o *AnalysisArchiveTransitionRuleExclude) SetExpirationDays(v int32) {
	o.ExpirationDays = &v
}

// GetLastSeenInDays returns the LastSeenInDays field value if set, zero value otherwise.
func (o *AnalysisArchiveTransitionRuleExclude) GetLastSeenInDays() int32 {
	if o == nil || o.LastSeenInDays == nil {
		var ret int32
		return ret
	}
	return *o.LastSeenInDays
}

// GetLastSeenInDaysOk returns a tuple with the LastSeenInDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisArchiveTransitionRuleExclude) GetLastSeenInDaysOk() (*int32, bool) {
	if o == nil || o.LastSeenInDays == nil {
		return nil, false
	}
	return o.LastSeenInDays, true
}

// HasLastSeenInDays returns a boolean if a field has been set.
func (o *AnalysisArchiveTransitionRuleExclude) HasLastSeenInDays() bool {
	if o != nil && o.LastSeenInDays != nil {
		return true
	}

	return false
}

// SetLastSeenInDays gets a reference to the given int32 and assigns it to the LastSeenInDays field.
func (o *AnalysisArchiveTransitionRuleExclude) SetLastSeenInDays(v int32) {
	o.LastSeenInDays = &v
}

func (o AnalysisArchiveTransitionRuleExclude) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Selector != nil {
		toSerialize["selector"] = o.Selector
	}
	if o.ExpirationDays != nil {
		toSerialize["expiration_days"] = o.ExpirationDays
	}
	if o.LastSeenInDays != nil {
		toSerialize["last_seen_in_days"] = o.LastSeenInDays
	}
	return json.Marshal(toSerialize)
}

type NullableAnalysisArchiveTransitionRuleExclude struct {
	value *AnalysisArchiveTransitionRuleExclude
	isSet bool
}

func (v NullableAnalysisArchiveTransitionRuleExclude) Get() *AnalysisArchiveTransitionRuleExclude {
	return v.value
}

func (v *NullableAnalysisArchiveTransitionRuleExclude) Set(val *AnalysisArchiveTransitionRuleExclude) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalysisArchiveTransitionRuleExclude) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalysisArchiveTransitionRuleExclude) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalysisArchiveTransitionRuleExclude(val *AnalysisArchiveTransitionRuleExclude) *NullableAnalysisArchiveTransitionRuleExclude {
	return &NullableAnalysisArchiveTransitionRuleExclude{value: val, isSet: true}
}

func (v NullableAnalysisArchiveTransitionRuleExclude) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalysisArchiveTransitionRuleExclude) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


