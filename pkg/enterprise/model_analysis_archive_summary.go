/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 1.0.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// AnalysisArchiveSummary A summarization of the analysis archive, including size, counts, etc. This archive stores image analysis only, never the actual image content or layers.
type AnalysisArchiveSummary struct {
	// The number of unique images (digests) in the archive
	TotalImageCount *int32 `json:"total_image_count,omitempty"`
	// The number of tag records (registry/repo:tag pull strings) in the archive. This may include repeated tags but will always have a unique tag->digest mapping per record.
	TotalTagCount *int32 `json:"total_tag_count,omitempty"`
	// The total sum of all the bytes stored to the backing storage. Accounts for anchore-applied compression, but not compression by the underlying storage system.
	TotalDataBytes *int32 `json:"total_data_bytes,omitempty"`
	// The timestamp of the most recent archived image
	LastUpdated *time.Time `json:"last_updated,omitempty"`
}

// NewAnalysisArchiveSummary instantiates a new AnalysisArchiveSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalysisArchiveSummary() *AnalysisArchiveSummary {
	this := AnalysisArchiveSummary{}
	return &this
}

// NewAnalysisArchiveSummaryWithDefaults instantiates a new AnalysisArchiveSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalysisArchiveSummaryWithDefaults() *AnalysisArchiveSummary {
	this := AnalysisArchiveSummary{}
	return &this
}

// GetTotalImageCount returns the TotalImageCount field value if set, zero value otherwise.
func (o *AnalysisArchiveSummary) GetTotalImageCount() int32 {
	if o == nil || o.TotalImageCount == nil {
		var ret int32
		return ret
	}
	return *o.TotalImageCount
}

// GetTotalImageCountOk returns a tuple with the TotalImageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisArchiveSummary) GetTotalImageCountOk() (*int32, bool) {
	if o == nil || o.TotalImageCount == nil {
		return nil, false
	}
	return o.TotalImageCount, true
}

// HasTotalImageCount returns a boolean if a field has been set.
func (o *AnalysisArchiveSummary) HasTotalImageCount() bool {
	if o != nil && o.TotalImageCount != nil {
		return true
	}

	return false
}

// SetTotalImageCount gets a reference to the given int32 and assigns it to the TotalImageCount field.
func (o *AnalysisArchiveSummary) SetTotalImageCount(v int32) {
	o.TotalImageCount = &v
}

// GetTotalTagCount returns the TotalTagCount field value if set, zero value otherwise.
func (o *AnalysisArchiveSummary) GetTotalTagCount() int32 {
	if o == nil || o.TotalTagCount == nil {
		var ret int32
		return ret
	}
	return *o.TotalTagCount
}

// GetTotalTagCountOk returns a tuple with the TotalTagCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisArchiveSummary) GetTotalTagCountOk() (*int32, bool) {
	if o == nil || o.TotalTagCount == nil {
		return nil, false
	}
	return o.TotalTagCount, true
}

// HasTotalTagCount returns a boolean if a field has been set.
func (o *AnalysisArchiveSummary) HasTotalTagCount() bool {
	if o != nil && o.TotalTagCount != nil {
		return true
	}

	return false
}

// SetTotalTagCount gets a reference to the given int32 and assigns it to the TotalTagCount field.
func (o *AnalysisArchiveSummary) SetTotalTagCount(v int32) {
	o.TotalTagCount = &v
}

// GetTotalDataBytes returns the TotalDataBytes field value if set, zero value otherwise.
func (o *AnalysisArchiveSummary) GetTotalDataBytes() int32 {
	if o == nil || o.TotalDataBytes == nil {
		var ret int32
		return ret
	}
	return *o.TotalDataBytes
}

// GetTotalDataBytesOk returns a tuple with the TotalDataBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisArchiveSummary) GetTotalDataBytesOk() (*int32, bool) {
	if o == nil || o.TotalDataBytes == nil {
		return nil, false
	}
	return o.TotalDataBytes, true
}

// HasTotalDataBytes returns a boolean if a field has been set.
func (o *AnalysisArchiveSummary) HasTotalDataBytes() bool {
	if o != nil && o.TotalDataBytes != nil {
		return true
	}

	return false
}

// SetTotalDataBytes gets a reference to the given int32 and assigns it to the TotalDataBytes field.
func (o *AnalysisArchiveSummary) SetTotalDataBytes(v int32) {
	o.TotalDataBytes = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *AnalysisArchiveSummary) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisArchiveSummary) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *AnalysisArchiveSummary) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *AnalysisArchiveSummary) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o AnalysisArchiveSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TotalImageCount != nil {
		toSerialize["total_image_count"] = o.TotalImageCount
	}
	if o.TotalTagCount != nil {
		toSerialize["total_tag_count"] = o.TotalTagCount
	}
	if o.TotalDataBytes != nil {
		toSerialize["total_data_bytes"] = o.TotalDataBytes
	}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	return json.Marshal(toSerialize)
}

type NullableAnalysisArchiveSummary struct {
	value *AnalysisArchiveSummary
	isSet bool
}

func (v NullableAnalysisArchiveSummary) Get() *AnalysisArchiveSummary {
	return v.value
}

func (v *NullableAnalysisArchiveSummary) Set(val *AnalysisArchiveSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalysisArchiveSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalysisArchiveSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalysisArchiveSummary(val *AnalysisArchiveSummary) *NullableAnalysisArchiveSummary {
	return &NullableAnalysisArchiveSummary{value: val, isSet: true}
}

func (v NullableAnalysisArchiveSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalysisArchiveSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

