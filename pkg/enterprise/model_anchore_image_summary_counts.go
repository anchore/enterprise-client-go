/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.11.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the AnchoreImageSummaryCounts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnchoreImageSummaryCounts{}

// AnchoreImageSummaryCounts Analysis stats for a repo
type AnchoreImageSummaryCounts struct {
	Tags *int32 `json:"tags,omitempty"`
	Images *int32 `json:"images,omitempty"`
	Analyzed *int32 `json:"analyzed,omitempty"`
	Analyzing *int32 `json:"analyzing,omitempty"`
	Pending *int32 `json:"pending,omitempty"`
	Failed *int32 `json:"failed,omitempty"`
}

// NewAnchoreImageSummaryCounts instantiates a new AnchoreImageSummaryCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnchoreImageSummaryCounts() *AnchoreImageSummaryCounts {
	this := AnchoreImageSummaryCounts{}
	return &this
}

// NewAnchoreImageSummaryCountsWithDefaults instantiates a new AnchoreImageSummaryCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnchoreImageSummaryCountsWithDefaults() *AnchoreImageSummaryCounts {
	this := AnchoreImageSummaryCounts{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AnchoreImageSummaryCounts) GetTags() int32 {
	if o == nil || IsNil(o.Tags) {
		var ret int32
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnchoreImageSummaryCounts) GetTagsOk() (*int32, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AnchoreImageSummaryCounts) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given int32 and assigns it to the Tags field.
func (o *AnchoreImageSummaryCounts) SetTags(v int32) {
	o.Tags = &v
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *AnchoreImageSummaryCounts) GetImages() int32 {
	if o == nil || IsNil(o.Images) {
		var ret int32
		return ret
	}
	return *o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnchoreImageSummaryCounts) GetImagesOk() (*int32, bool) {
	if o == nil || IsNil(o.Images) {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *AnchoreImageSummaryCounts) HasImages() bool {
	if o != nil && !IsNil(o.Images) {
		return true
	}

	return false
}

// SetImages gets a reference to the given int32 and assigns it to the Images field.
func (o *AnchoreImageSummaryCounts) SetImages(v int32) {
	o.Images = &v
}

// GetAnalyzed returns the Analyzed field value if set, zero value otherwise.
func (o *AnchoreImageSummaryCounts) GetAnalyzed() int32 {
	if o == nil || IsNil(o.Analyzed) {
		var ret int32
		return ret
	}
	return *o.Analyzed
}

// GetAnalyzedOk returns a tuple with the Analyzed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnchoreImageSummaryCounts) GetAnalyzedOk() (*int32, bool) {
	if o == nil || IsNil(o.Analyzed) {
		return nil, false
	}
	return o.Analyzed, true
}

// HasAnalyzed returns a boolean if a field has been set.
func (o *AnchoreImageSummaryCounts) HasAnalyzed() bool {
	if o != nil && !IsNil(o.Analyzed) {
		return true
	}

	return false
}

// SetAnalyzed gets a reference to the given int32 and assigns it to the Analyzed field.
func (o *AnchoreImageSummaryCounts) SetAnalyzed(v int32) {
	o.Analyzed = &v
}

// GetAnalyzing returns the Analyzing field value if set, zero value otherwise.
func (o *AnchoreImageSummaryCounts) GetAnalyzing() int32 {
	if o == nil || IsNil(o.Analyzing) {
		var ret int32
		return ret
	}
	return *o.Analyzing
}

// GetAnalyzingOk returns a tuple with the Analyzing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnchoreImageSummaryCounts) GetAnalyzingOk() (*int32, bool) {
	if o == nil || IsNil(o.Analyzing) {
		return nil, false
	}
	return o.Analyzing, true
}

// HasAnalyzing returns a boolean if a field has been set.
func (o *AnchoreImageSummaryCounts) HasAnalyzing() bool {
	if o != nil && !IsNil(o.Analyzing) {
		return true
	}

	return false
}

// SetAnalyzing gets a reference to the given int32 and assigns it to the Analyzing field.
func (o *AnchoreImageSummaryCounts) SetAnalyzing(v int32) {
	o.Analyzing = &v
}

// GetPending returns the Pending field value if set, zero value otherwise.
func (o *AnchoreImageSummaryCounts) GetPending() int32 {
	if o == nil || IsNil(o.Pending) {
		var ret int32
		return ret
	}
	return *o.Pending
}

// GetPendingOk returns a tuple with the Pending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnchoreImageSummaryCounts) GetPendingOk() (*int32, bool) {
	if o == nil || IsNil(o.Pending) {
		return nil, false
	}
	return o.Pending, true
}

// HasPending returns a boolean if a field has been set.
func (o *AnchoreImageSummaryCounts) HasPending() bool {
	if o != nil && !IsNil(o.Pending) {
		return true
	}

	return false
}

// SetPending gets a reference to the given int32 and assigns it to the Pending field.
func (o *AnchoreImageSummaryCounts) SetPending(v int32) {
	o.Pending = &v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *AnchoreImageSummaryCounts) GetFailed() int32 {
	if o == nil || IsNil(o.Failed) {
		var ret int32
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnchoreImageSummaryCounts) GetFailedOk() (*int32, bool) {
	if o == nil || IsNil(o.Failed) {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *AnchoreImageSummaryCounts) HasFailed() bool {
	if o != nil && !IsNil(o.Failed) {
		return true
	}

	return false
}

// SetFailed gets a reference to the given int32 and assigns it to the Failed field.
func (o *AnchoreImageSummaryCounts) SetFailed(v int32) {
	o.Failed = &v
}

func (o AnchoreImageSummaryCounts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnchoreImageSummaryCounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Images) {
		toSerialize["images"] = o.Images
	}
	if !IsNil(o.Analyzed) {
		toSerialize["analyzed"] = o.Analyzed
	}
	if !IsNil(o.Analyzing) {
		toSerialize["analyzing"] = o.Analyzing
	}
	if !IsNil(o.Pending) {
		toSerialize["pending"] = o.Pending
	}
	if !IsNil(o.Failed) {
		toSerialize["failed"] = o.Failed
	}
	return toSerialize, nil
}

type NullableAnchoreImageSummaryCounts struct {
	value *AnchoreImageSummaryCounts
	isSet bool
}

func (v NullableAnchoreImageSummaryCounts) Get() *AnchoreImageSummaryCounts {
	return v.value
}

func (v *NullableAnchoreImageSummaryCounts) Set(val *AnchoreImageSummaryCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableAnchoreImageSummaryCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableAnchoreImageSummaryCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnchoreImageSummaryCounts(val *AnchoreImageSummaryCounts) *NullableAnchoreImageSummaryCounts {
	return &NullableAnchoreImageSummaryCounts{value: val, isSet: true}
}

func (v NullableAnchoreImageSummaryCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnchoreImageSummaryCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


