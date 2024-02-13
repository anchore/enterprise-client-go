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

// AnchoreImageTagSummary A unique image in the engine.
type AnchoreImageTagSummary struct {
	ImageDigest *string `json:"image_digest,omitempty"`
	ParentDigest *string `json:"parent_digest,omitempty"`
	ImageId *string `json:"image_id,omitempty"`
	AnalysisStatus *string `json:"analysis_status,omitempty"`
	FullTag *string `json:"full_tag,omitempty"`
	CreatedAt *int32 `json:"created_at,omitempty"`
	AnalyzedAt *int32 `json:"analyzed_at,omitempty"`
	TagDetectedAt *int32 `json:"tag_detected_at,omitempty"`
	ImageStatus *string `json:"image_status,omitempty"`
}

// NewAnchoreImageTagSummary instantiates a new AnchoreImageTagSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnchoreImageTagSummary() *AnchoreImageTagSummary {
	this := AnchoreImageTagSummary{}
	return &this
}

// NewAnchoreImageTagSummaryWithDefaults instantiates a new AnchoreImageTagSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnchoreImageTagSummaryWithDefaults() *AnchoreImageTagSummary {
	this := AnchoreImageTagSummary{}
	return &this
}

// GetImageDigest returns the ImageDigest field value if set, zero value otherwise.
func (o *AnchoreImageTagSummary) GetImageDigest() string {
	if o == nil || o.ImageDigest == nil {
		var ret string
		return ret
	}
	return *o.ImageDigest
}

// GetImageDigestOk returns a tuple with the ImageDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnchoreImageTagSummary) GetImageDigestOk() (*string, bool) {
	if o == nil || o.ImageDigest == nil {
		return nil, false
	}
	return o.ImageDigest, true
}

// HasImageDigest returns a boolean if a field has been set.
func (o *AnchoreImageTagSummary) HasImageDigest() bool {
	if o != nil && o.ImageDigest != nil {
		return true
	}

	return false
}

// SetImageDigest gets a reference to the given string and assigns it to the ImageDigest field.
func (o *AnchoreImageTagSummary) SetImageDigest(v string) {
	o.ImageDigest = &v
}

// GetParentDigest returns the ParentDigest field value if set, zero value otherwise.
func (o *AnchoreImageTagSummary) GetParentDigest() string {
	if o == nil || o.ParentDigest == nil {
		var ret string
		return ret
	}
	return *o.ParentDigest
}

// GetParentDigestOk returns a tuple with the ParentDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnchoreImageTagSummary) GetParentDigestOk() (*string, bool) {
	if o == nil || o.ParentDigest == nil {
		return nil, false
	}
	return o.ParentDigest, true
}

// HasParentDigest returns a boolean if a field has been set.
func (o *AnchoreImageTagSummary) HasParentDigest() bool {
	if o != nil && o.ParentDigest != nil {
		return true
	}

	return false
}

// SetParentDigest gets a reference to the given string and assigns it to the ParentDigest field.
func (o *AnchoreImageTagSummary) SetParentDigest(v string) {
	o.ParentDigest = &v
}

// GetImageId returns the ImageId field value if set, zero value otherwise.
func (o *AnchoreImageTagSummary) GetImageId() string {
	if o == nil || o.ImageId == nil {
		var ret string
		return ret
	}
	return *o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnchoreImageTagSummary) GetImageIdOk() (*string, bool) {
	if o == nil || o.ImageId == nil {
		return nil, false
	}
	return o.ImageId, true
}

// HasImageId returns a boolean if a field has been set.
func (o *AnchoreImageTagSummary) HasImageId() bool {
	if o != nil && o.ImageId != nil {
		return true
	}

	return false
}

// SetImageId gets a reference to the given string and assigns it to the ImageId field.
func (o *AnchoreImageTagSummary) SetImageId(v string) {
	o.ImageId = &v
}

// GetAnalysisStatus returns the AnalysisStatus field value if set, zero value otherwise.
func (o *AnchoreImageTagSummary) GetAnalysisStatus() string {
	if o == nil || o.AnalysisStatus == nil {
		var ret string
		return ret
	}
	return *o.AnalysisStatus
}

// GetAnalysisStatusOk returns a tuple with the AnalysisStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnchoreImageTagSummary) GetAnalysisStatusOk() (*string, bool) {
	if o == nil || o.AnalysisStatus == nil {
		return nil, false
	}
	return o.AnalysisStatus, true
}

// HasAnalysisStatus returns a boolean if a field has been set.
func (o *AnchoreImageTagSummary) HasAnalysisStatus() bool {
	if o != nil && o.AnalysisStatus != nil {
		return true
	}

	return false
}

// SetAnalysisStatus gets a reference to the given string and assigns it to the AnalysisStatus field.
func (o *AnchoreImageTagSummary) SetAnalysisStatus(v string) {
	o.AnalysisStatus = &v
}

// GetFullTag returns the FullTag field value if set, zero value otherwise.
func (o *AnchoreImageTagSummary) GetFullTag() string {
	if o == nil || o.FullTag == nil {
		var ret string
		return ret
	}
	return *o.FullTag
}

// GetFullTagOk returns a tuple with the FullTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnchoreImageTagSummary) GetFullTagOk() (*string, bool) {
	if o == nil || o.FullTag == nil {
		return nil, false
	}
	return o.FullTag, true
}

// HasFullTag returns a boolean if a field has been set.
func (o *AnchoreImageTagSummary) HasFullTag() bool {
	if o != nil && o.FullTag != nil {
		return true
	}

	return false
}

// SetFullTag gets a reference to the given string and assigns it to the FullTag field.
func (o *AnchoreImageTagSummary) SetFullTag(v string) {
	o.FullTag = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AnchoreImageTagSummary) GetCreatedAt() int32 {
	if o == nil || o.CreatedAt == nil {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnchoreImageTagSummary) GetCreatedAtOk() (*int32, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AnchoreImageTagSummary) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *AnchoreImageTagSummary) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetAnalyzedAt returns the AnalyzedAt field value if set, zero value otherwise.
func (o *AnchoreImageTagSummary) GetAnalyzedAt() int32 {
	if o == nil || o.AnalyzedAt == nil {
		var ret int32
		return ret
	}
	return *o.AnalyzedAt
}

// GetAnalyzedAtOk returns a tuple with the AnalyzedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnchoreImageTagSummary) GetAnalyzedAtOk() (*int32, bool) {
	if o == nil || o.AnalyzedAt == nil {
		return nil, false
	}
	return o.AnalyzedAt, true
}

// HasAnalyzedAt returns a boolean if a field has been set.
func (o *AnchoreImageTagSummary) HasAnalyzedAt() bool {
	if o != nil && o.AnalyzedAt != nil {
		return true
	}

	return false
}

// SetAnalyzedAt gets a reference to the given int32 and assigns it to the AnalyzedAt field.
func (o *AnchoreImageTagSummary) SetAnalyzedAt(v int32) {
	o.AnalyzedAt = &v
}

// GetTagDetectedAt returns the TagDetectedAt field value if set, zero value otherwise.
func (o *AnchoreImageTagSummary) GetTagDetectedAt() int32 {
	if o == nil || o.TagDetectedAt == nil {
		var ret int32
		return ret
	}
	return *o.TagDetectedAt
}

// GetTagDetectedAtOk returns a tuple with the TagDetectedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnchoreImageTagSummary) GetTagDetectedAtOk() (*int32, bool) {
	if o == nil || o.TagDetectedAt == nil {
		return nil, false
	}
	return o.TagDetectedAt, true
}

// HasTagDetectedAt returns a boolean if a field has been set.
func (o *AnchoreImageTagSummary) HasTagDetectedAt() bool {
	if o != nil && o.TagDetectedAt != nil {
		return true
	}

	return false
}

// SetTagDetectedAt gets a reference to the given int32 and assigns it to the TagDetectedAt field.
func (o *AnchoreImageTagSummary) SetTagDetectedAt(v int32) {
	o.TagDetectedAt = &v
}

// GetImageStatus returns the ImageStatus field value if set, zero value otherwise.
func (o *AnchoreImageTagSummary) GetImageStatus() string {
	if o == nil || o.ImageStatus == nil {
		var ret string
		return ret
	}
	return *o.ImageStatus
}

// GetImageStatusOk returns a tuple with the ImageStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnchoreImageTagSummary) GetImageStatusOk() (*string, bool) {
	if o == nil || o.ImageStatus == nil {
		return nil, false
	}
	return o.ImageStatus, true
}

// HasImageStatus returns a boolean if a field has been set.
func (o *AnchoreImageTagSummary) HasImageStatus() bool {
	if o != nil && o.ImageStatus != nil {
		return true
	}

	return false
}

// SetImageStatus gets a reference to the given string and assigns it to the ImageStatus field.
func (o *AnchoreImageTagSummary) SetImageStatus(v string) {
	o.ImageStatus = &v
}

func (o AnchoreImageTagSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ImageDigest != nil {
		toSerialize["image_digest"] = o.ImageDigest
	}
	if o.ParentDigest != nil {
		toSerialize["parent_digest"] = o.ParentDigest
	}
	if o.ImageId != nil {
		toSerialize["image_id"] = o.ImageId
	}
	if o.AnalysisStatus != nil {
		toSerialize["analysis_status"] = o.AnalysisStatus
	}
	if o.FullTag != nil {
		toSerialize["full_tag"] = o.FullTag
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.AnalyzedAt != nil {
		toSerialize["analyzed_at"] = o.AnalyzedAt
	}
	if o.TagDetectedAt != nil {
		toSerialize["tag_detected_at"] = o.TagDetectedAt
	}
	if o.ImageStatus != nil {
		toSerialize["image_status"] = o.ImageStatus
	}
	return json.Marshal(toSerialize)
}

type NullableAnchoreImageTagSummary struct {
	value *AnchoreImageTagSummary
	isSet bool
}

func (v NullableAnchoreImageTagSummary) Get() *AnchoreImageTagSummary {
	return v.value
}

func (v *NullableAnchoreImageTagSummary) Set(val *AnchoreImageTagSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableAnchoreImageTagSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableAnchoreImageTagSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnchoreImageTagSummary(val *AnchoreImageTagSummary) *NullableAnchoreImageTagSummary {
	return &NullableAnchoreImageTagSummary{value: val, isSet: true}
}

func (v NullableAnchoreImageTagSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnchoreImageTagSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


