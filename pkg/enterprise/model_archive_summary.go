/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.6.1
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// ArchiveSummary A summarization of the available archives, a place to for long-term storage of audit, analysis, or other data to remove it from the system's working set but keep it available.
type ArchiveSummary struct {
	Images *AnalysisArchiveSummary `json:"images,omitempty"`
	Rules *AnalysisArchiveRulesSummary `json:"rules,omitempty"`
}

// NewArchiveSummary instantiates a new ArchiveSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArchiveSummary() *ArchiveSummary {
	this := ArchiveSummary{}
	return &this
}

// NewArchiveSummaryWithDefaults instantiates a new ArchiveSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArchiveSummaryWithDefaults() *ArchiveSummary {
	this := ArchiveSummary{}
	return &this
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *ArchiveSummary) GetImages() AnalysisArchiveSummary {
	if o == nil || o.Images == nil {
		var ret AnalysisArchiveSummary
		return ret
	}
	return *o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveSummary) GetImagesOk() (*AnalysisArchiveSummary, bool) {
	if o == nil || o.Images == nil {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *ArchiveSummary) HasImages() bool {
	if o != nil && o.Images != nil {
		return true
	}

	return false
}

// SetImages gets a reference to the given AnalysisArchiveSummary and assigns it to the Images field.
func (o *ArchiveSummary) SetImages(v AnalysisArchiveSummary) {
	o.Images = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *ArchiveSummary) GetRules() AnalysisArchiveRulesSummary {
	if o == nil || o.Rules == nil {
		var ret AnalysisArchiveRulesSummary
		return ret
	}
	return *o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveSummary) GetRulesOk() (*AnalysisArchiveRulesSummary, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *ArchiveSummary) HasRules() bool {
	if o != nil && o.Rules != nil {
		return true
	}

	return false
}

// SetRules gets a reference to the given AnalysisArchiveRulesSummary and assigns it to the Rules field.
func (o *ArchiveSummary) SetRules(v AnalysisArchiveRulesSummary) {
	o.Rules = &v
}

func (o ArchiveSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Images != nil {
		toSerialize["images"] = o.Images
	}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableArchiveSummary struct {
	value *ArchiveSummary
	isSet bool
}

func (v NullableArchiveSummary) Get() *ArchiveSummary {
	return v.value
}

func (v *NullableArchiveSummary) Set(val *ArchiveSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableArchiveSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableArchiveSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArchiveSummary(val *ArchiveSummary) *NullableArchiveSummary {
	return &NullableArchiveSummary{value: val, isSet: true}
}

func (v NullableArchiveSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArchiveSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


