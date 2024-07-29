/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.6.2
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// GroupSyncResult struct for GroupSyncResult
type GroupSyncResult struct {
	// The name of the group
	Group *string `json:"group,omitempty"`
	Status *string `json:"status,omitempty"`
	// The number of images updated by the this group sync, across all accounts. This is typically only non-zero for vulnerability feeds which update images' vulnerability results during the sync.
	UpdatedImageCount *int32 `json:"updated_image_count,omitempty"`
	// The number of feed data records synced down as either updates or new records
	UpdatedRecordCount *int32 `json:"updated_record_count,omitempty"`
	// The duration of the group sync in seconds
	TotalTimeSeconds *float32 `json:"total_time_seconds,omitempty"`
}

// NewGroupSyncResult instantiates a new GroupSyncResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupSyncResult() *GroupSyncResult {
	this := GroupSyncResult{}
	return &this
}

// NewGroupSyncResultWithDefaults instantiates a new GroupSyncResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupSyncResultWithDefaults() *GroupSyncResult {
	this := GroupSyncResult{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *GroupSyncResult) GetGroup() string {
	if o == nil || o.Group == nil {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSyncResult) GetGroupOk() (*string, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *GroupSyncResult) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *GroupSyncResult) SetGroup(v string) {
	o.Group = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GroupSyncResult) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSyncResult) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GroupSyncResult) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GroupSyncResult) SetStatus(v string) {
	o.Status = &v
}

// GetUpdatedImageCount returns the UpdatedImageCount field value if set, zero value otherwise.
func (o *GroupSyncResult) GetUpdatedImageCount() int32 {
	if o == nil || o.UpdatedImageCount == nil {
		var ret int32
		return ret
	}
	return *o.UpdatedImageCount
}

// GetUpdatedImageCountOk returns a tuple with the UpdatedImageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSyncResult) GetUpdatedImageCountOk() (*int32, bool) {
	if o == nil || o.UpdatedImageCount == nil {
		return nil, false
	}
	return o.UpdatedImageCount, true
}

// HasUpdatedImageCount returns a boolean if a field has been set.
func (o *GroupSyncResult) HasUpdatedImageCount() bool {
	if o != nil && o.UpdatedImageCount != nil {
		return true
	}

	return false
}

// SetUpdatedImageCount gets a reference to the given int32 and assigns it to the UpdatedImageCount field.
func (o *GroupSyncResult) SetUpdatedImageCount(v int32) {
	o.UpdatedImageCount = &v
}

// GetUpdatedRecordCount returns the UpdatedRecordCount field value if set, zero value otherwise.
func (o *GroupSyncResult) GetUpdatedRecordCount() int32 {
	if o == nil || o.UpdatedRecordCount == nil {
		var ret int32
		return ret
	}
	return *o.UpdatedRecordCount
}

// GetUpdatedRecordCountOk returns a tuple with the UpdatedRecordCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSyncResult) GetUpdatedRecordCountOk() (*int32, bool) {
	if o == nil || o.UpdatedRecordCount == nil {
		return nil, false
	}
	return o.UpdatedRecordCount, true
}

// HasUpdatedRecordCount returns a boolean if a field has been set.
func (o *GroupSyncResult) HasUpdatedRecordCount() bool {
	if o != nil && o.UpdatedRecordCount != nil {
		return true
	}

	return false
}

// SetUpdatedRecordCount gets a reference to the given int32 and assigns it to the UpdatedRecordCount field.
func (o *GroupSyncResult) SetUpdatedRecordCount(v int32) {
	o.UpdatedRecordCount = &v
}

// GetTotalTimeSeconds returns the TotalTimeSeconds field value if set, zero value otherwise.
func (o *GroupSyncResult) GetTotalTimeSeconds() float32 {
	if o == nil || o.TotalTimeSeconds == nil {
		var ret float32
		return ret
	}
	return *o.TotalTimeSeconds
}

// GetTotalTimeSecondsOk returns a tuple with the TotalTimeSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSyncResult) GetTotalTimeSecondsOk() (*float32, bool) {
	if o == nil || o.TotalTimeSeconds == nil {
		return nil, false
	}
	return o.TotalTimeSeconds, true
}

// HasTotalTimeSeconds returns a boolean if a field has been set.
func (o *GroupSyncResult) HasTotalTimeSeconds() bool {
	if o != nil && o.TotalTimeSeconds != nil {
		return true
	}

	return false
}

// SetTotalTimeSeconds gets a reference to the given float32 and assigns it to the TotalTimeSeconds field.
func (o *GroupSyncResult) SetTotalTimeSeconds(v float32) {
	o.TotalTimeSeconds = &v
}

func (o GroupSyncResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.UpdatedImageCount != nil {
		toSerialize["updated_image_count"] = o.UpdatedImageCount
	}
	if o.UpdatedRecordCount != nil {
		toSerialize["updated_record_count"] = o.UpdatedRecordCount
	}
	if o.TotalTimeSeconds != nil {
		toSerialize["total_time_seconds"] = o.TotalTimeSeconds
	}
	return json.Marshal(toSerialize)
}

type NullableGroupSyncResult struct {
	value *GroupSyncResult
	isSet bool
}

func (v NullableGroupSyncResult) Get() *GroupSyncResult {
	return v.value
}

func (v *NullableGroupSyncResult) Set(val *GroupSyncResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupSyncResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupSyncResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupSyncResult(val *GroupSyncResult) *NullableGroupSyncResult {
	return &NullableGroupSyncResult{value: val, isSet: true}
}

func (v NullableGroupSyncResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupSyncResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


