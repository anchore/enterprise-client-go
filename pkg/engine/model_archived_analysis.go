/*
Anchore Engine API Server

This is the Anchore Engine API. Provides the primary external API for users of the service.

API version: 0.7.0
Contact: nurmi@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package engine

import (
	"encoding/json"
	"time"
)

// ArchivedAnalysis struct for ArchivedAnalysis
type ArchivedAnalysis struct {
	// The image digest (digest of the manifest describing the image, per docker spec)
	ImageDigest *string `json:"imageDigest,omitempty"`
	// The digest of a parent manifest (for manifest-list images)
	ParentDigest *string `json:"parentDigest,omitempty"`
	// User provided annotations as key-value pairs
	Annotations *interface{} `json:"annotations,omitempty"`
	// The archival status
	Status *string `json:"status,omitempty"`
	// List of tags associated with the image digest
	ImageDetail *[]ArchiveTagEntry `json:"image_detail,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
	AnalyzedAt *time.Time `json:"analyzed_at,omitempty"`
	// The size, in bytes, of the analysis archive file
	ArchiveSizeBytes *int32 `json:"archive_size_bytes,omitempty"`
}

// NewArchivedAnalysis instantiates a new ArchivedAnalysis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArchivedAnalysis() *ArchivedAnalysis {
	this := ArchivedAnalysis{}
	return &this
}

// NewArchivedAnalysisWithDefaults instantiates a new ArchivedAnalysis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArchivedAnalysisWithDefaults() *ArchivedAnalysis {
	this := ArchivedAnalysis{}
	return &this
}

// GetImageDigest returns the ImageDigest field value if set, zero value otherwise.
func (o *ArchivedAnalysis) GetImageDigest() string {
	if o == nil || o.ImageDigest == nil {
		var ret string
		return ret
	}
	return *o.ImageDigest
}

// GetImageDigestOk returns a tuple with the ImageDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchivedAnalysis) GetImageDigestOk() (*string, bool) {
	if o == nil || o.ImageDigest == nil {
		return nil, false
	}
	return o.ImageDigest, true
}

// HasImageDigest returns a boolean if a field has been set.
func (o *ArchivedAnalysis) HasImageDigest() bool {
	if o != nil && o.ImageDigest != nil {
		return true
	}

	return false
}

// SetImageDigest gets a reference to the given string and assigns it to the ImageDigest field.
func (o *ArchivedAnalysis) SetImageDigest(v string) {
	o.ImageDigest = &v
}

// GetParentDigest returns the ParentDigest field value if set, zero value otherwise.
func (o *ArchivedAnalysis) GetParentDigest() string {
	if o == nil || o.ParentDigest == nil {
		var ret string
		return ret
	}
	return *o.ParentDigest
}

// GetParentDigestOk returns a tuple with the ParentDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchivedAnalysis) GetParentDigestOk() (*string, bool) {
	if o == nil || o.ParentDigest == nil {
		return nil, false
	}
	return o.ParentDigest, true
}

// HasParentDigest returns a boolean if a field has been set.
func (o *ArchivedAnalysis) HasParentDigest() bool {
	if o != nil && o.ParentDigest != nil {
		return true
	}

	return false
}

// SetParentDigest gets a reference to the given string and assigns it to the ParentDigest field.
func (o *ArchivedAnalysis) SetParentDigest(v string) {
	o.ParentDigest = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *ArchivedAnalysis) GetAnnotations() interface{} {
	if o == nil || o.Annotations == nil {
		var ret interface{}
		return ret
	}
	return *o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchivedAnalysis) GetAnnotationsOk() (*interface{}, bool) {
	if o == nil || o.Annotations == nil {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *ArchivedAnalysis) HasAnnotations() bool {
	if o != nil && o.Annotations != nil {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given interface{} and assigns it to the Annotations field.
func (o *ArchivedAnalysis) SetAnnotations(v interface{}) {
	o.Annotations = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ArchivedAnalysis) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchivedAnalysis) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ArchivedAnalysis) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ArchivedAnalysis) SetStatus(v string) {
	o.Status = &v
}

// GetImageDetail returns the ImageDetail field value if set, zero value otherwise.
func (o *ArchivedAnalysis) GetImageDetail() []ArchiveTagEntry {
	if o == nil || o.ImageDetail == nil {
		var ret []ArchiveTagEntry
		return ret
	}
	return *o.ImageDetail
}

// GetImageDetailOk returns a tuple with the ImageDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchivedAnalysis) GetImageDetailOk() (*[]ArchiveTagEntry, bool) {
	if o == nil || o.ImageDetail == nil {
		return nil, false
	}
	return o.ImageDetail, true
}

// HasImageDetail returns a boolean if a field has been set.
func (o *ArchivedAnalysis) HasImageDetail() bool {
	if o != nil && o.ImageDetail != nil {
		return true
	}

	return false
}

// SetImageDetail gets a reference to the given []ArchiveTagEntry and assigns it to the ImageDetail field.
func (o *ArchivedAnalysis) SetImageDetail(v []ArchiveTagEntry) {
	o.ImageDetail = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ArchivedAnalysis) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchivedAnalysis) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ArchivedAnalysis) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ArchivedAnalysis) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *ArchivedAnalysis) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchivedAnalysis) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *ArchivedAnalysis) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *ArchivedAnalysis) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetAnalyzedAt returns the AnalyzedAt field value if set, zero value otherwise.
func (o *ArchivedAnalysis) GetAnalyzedAt() time.Time {
	if o == nil || o.AnalyzedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.AnalyzedAt
}

// GetAnalyzedAtOk returns a tuple with the AnalyzedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchivedAnalysis) GetAnalyzedAtOk() (*time.Time, bool) {
	if o == nil || o.AnalyzedAt == nil {
		return nil, false
	}
	return o.AnalyzedAt, true
}

// HasAnalyzedAt returns a boolean if a field has been set.
func (o *ArchivedAnalysis) HasAnalyzedAt() bool {
	if o != nil && o.AnalyzedAt != nil {
		return true
	}

	return false
}

// SetAnalyzedAt gets a reference to the given time.Time and assigns it to the AnalyzedAt field.
func (o *ArchivedAnalysis) SetAnalyzedAt(v time.Time) {
	o.AnalyzedAt = &v
}

// GetArchiveSizeBytes returns the ArchiveSizeBytes field value if set, zero value otherwise.
func (o *ArchivedAnalysis) GetArchiveSizeBytes() int32 {
	if o == nil || o.ArchiveSizeBytes == nil {
		var ret int32
		return ret
	}
	return *o.ArchiveSizeBytes
}

// GetArchiveSizeBytesOk returns a tuple with the ArchiveSizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchivedAnalysis) GetArchiveSizeBytesOk() (*int32, bool) {
	if o == nil || o.ArchiveSizeBytes == nil {
		return nil, false
	}
	return o.ArchiveSizeBytes, true
}

// HasArchiveSizeBytes returns a boolean if a field has been set.
func (o *ArchivedAnalysis) HasArchiveSizeBytes() bool {
	if o != nil && o.ArchiveSizeBytes != nil {
		return true
	}

	return false
}

// SetArchiveSizeBytes gets a reference to the given int32 and assigns it to the ArchiveSizeBytes field.
func (o *ArchivedAnalysis) SetArchiveSizeBytes(v int32) {
	o.ArchiveSizeBytes = &v
}

func (o ArchivedAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ImageDigest != nil {
		toSerialize["imageDigest"] = o.ImageDigest
	}
	if o.ParentDigest != nil {
		toSerialize["parentDigest"] = o.ParentDigest
	}
	if o.Annotations != nil {
		toSerialize["annotations"] = o.Annotations
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ImageDetail != nil {
		toSerialize["image_detail"] = o.ImageDetail
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if o.AnalyzedAt != nil {
		toSerialize["analyzed_at"] = o.AnalyzedAt
	}
	if o.ArchiveSizeBytes != nil {
		toSerialize["archive_size_bytes"] = o.ArchiveSizeBytes
	}
	return json.Marshal(toSerialize)
}

type NullableArchivedAnalysis struct {
	value *ArchivedAnalysis
	isSet bool
}

func (v NullableArchivedAnalysis) Get() *ArchivedAnalysis {
	return v.value
}

func (v *NullableArchivedAnalysis) Set(val *ArchivedAnalysis) {
	v.value = val
	v.isSet = true
}

func (v NullableArchivedAnalysis) IsSet() bool {
	return v.isSet
}

func (v *NullableArchivedAnalysis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArchivedAnalysis(val *ArchivedAnalysis) *NullableArchivedAnalysis {
	return &NullableArchivedAnalysis{value: val, isSet: true}
}

func (v NullableArchivedAnalysis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArchivedAnalysis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


