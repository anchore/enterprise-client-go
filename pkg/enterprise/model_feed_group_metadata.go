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
	"time"
)

// checks if the FeedGroupMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeedGroupMetadata{}

// FeedGroupMetadata struct for FeedGroupMetadata
type FeedGroupMetadata struct {
	Name *string `json:"name,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	LastSync *time.Time `json:"last_sync,omitempty"`
	RecordCount *int32 `json:"record_count,omitempty"`
	// If group is enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// NewFeedGroupMetadata instantiates a new FeedGroupMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeedGroupMetadata() *FeedGroupMetadata {
	this := FeedGroupMetadata{}
	return &this
}

// NewFeedGroupMetadataWithDefaults instantiates a new FeedGroupMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeedGroupMetadataWithDefaults() *FeedGroupMetadata {
	this := FeedGroupMetadata{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FeedGroupMetadata) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedGroupMetadata) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FeedGroupMetadata) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FeedGroupMetadata) SetName(v string) {
	o.Name = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FeedGroupMetadata) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedGroupMetadata) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FeedGroupMetadata) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *FeedGroupMetadata) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastSync returns the LastSync field value if set, zero value otherwise.
func (o *FeedGroupMetadata) GetLastSync() time.Time {
	if o == nil || IsNil(o.LastSync) {
		var ret time.Time
		return ret
	}
	return *o.LastSync
}

// GetLastSyncOk returns a tuple with the LastSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedGroupMetadata) GetLastSyncOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastSync) {
		return nil, false
	}
	return o.LastSync, true
}

// HasLastSync returns a boolean if a field has been set.
func (o *FeedGroupMetadata) HasLastSync() bool {
	if o != nil && !IsNil(o.LastSync) {
		return true
	}

	return false
}

// SetLastSync gets a reference to the given time.Time and assigns it to the LastSync field.
func (o *FeedGroupMetadata) SetLastSync(v time.Time) {
	o.LastSync = &v
}

// GetRecordCount returns the RecordCount field value if set, zero value otherwise.
func (o *FeedGroupMetadata) GetRecordCount() int32 {
	if o == nil || IsNil(o.RecordCount) {
		var ret int32
		return ret
	}
	return *o.RecordCount
}

// GetRecordCountOk returns a tuple with the RecordCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedGroupMetadata) GetRecordCountOk() (*int32, bool) {
	if o == nil || IsNil(o.RecordCount) {
		return nil, false
	}
	return o.RecordCount, true
}

// HasRecordCount returns a boolean if a field has been set.
func (o *FeedGroupMetadata) HasRecordCount() bool {
	if o != nil && !IsNil(o.RecordCount) {
		return true
	}

	return false
}

// SetRecordCount gets a reference to the given int32 and assigns it to the RecordCount field.
func (o *FeedGroupMetadata) SetRecordCount(v int32) {
	o.RecordCount = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *FeedGroupMetadata) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedGroupMetadata) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *FeedGroupMetadata) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *FeedGroupMetadata) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o FeedGroupMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeedGroupMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.LastSync) {
		toSerialize["last_sync"] = o.LastSync
	}
	if !IsNil(o.RecordCount) {
		toSerialize["record_count"] = o.RecordCount
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableFeedGroupMetadata struct {
	value *FeedGroupMetadata
	isSet bool
}

func (v NullableFeedGroupMetadata) Get() *FeedGroupMetadata {
	return v.value
}

func (v *NullableFeedGroupMetadata) Set(val *FeedGroupMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedGroupMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedGroupMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedGroupMetadata(val *FeedGroupMetadata) *NullableFeedGroupMetadata {
	return &NullableFeedGroupMetadata{value: val, isSet: true}
}

func (v NullableFeedGroupMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedGroupMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


