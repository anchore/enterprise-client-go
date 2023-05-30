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
)

// NotificationBase base object for Notifications (every notification has this basic structure)
type NotificationBase struct {
	QueueId *string `json:"queueId,omitempty"`
	UserId *string `json:"userId,omitempty"`
	DataId *string `json:"dataId,omitempty"`
	CreatedAt *int32 `json:"created_at,omitempty"`
	LastUpdated *int32 `json:"last_updated,omitempty"`
	RecordStateKey *string `json:"record_state_key,omitempty"`
	RecordStateVal NullableString `json:"record_state_val,omitempty"`
	Tries *int32 `json:"tries,omitempty"`
	MaxTries *int32 `json:"max_tries,omitempty"`
}

// NewNotificationBase instantiates a new NotificationBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationBase() *NotificationBase {
	this := NotificationBase{}
	var recordStateKey string = "active"
	this.RecordStateKey = &recordStateKey
	return &this
}

// NewNotificationBaseWithDefaults instantiates a new NotificationBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationBaseWithDefaults() *NotificationBase {
	this := NotificationBase{}
	var recordStateKey string = "active"
	this.RecordStateKey = &recordStateKey
	return &this
}

// GetQueueId returns the QueueId field value if set, zero value otherwise.
func (o *NotificationBase) GetQueueId() string {
	if o == nil || o.QueueId == nil {
		var ret string
		return ret
	}
	return *o.QueueId
}

// GetQueueIdOk returns a tuple with the QueueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationBase) GetQueueIdOk() (*string, bool) {
	if o == nil || o.QueueId == nil {
		return nil, false
	}
	return o.QueueId, true
}

// HasQueueId returns a boolean if a field has been set.
func (o *NotificationBase) HasQueueId() bool {
	if o != nil && o.QueueId != nil {
		return true
	}

	return false
}

// SetQueueId gets a reference to the given string and assigns it to the QueueId field.
func (o *NotificationBase) SetQueueId(v string) {
	o.QueueId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *NotificationBase) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationBase) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *NotificationBase) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *NotificationBase) SetUserId(v string) {
	o.UserId = &v
}

// GetDataId returns the DataId field value if set, zero value otherwise.
func (o *NotificationBase) GetDataId() string {
	if o == nil || o.DataId == nil {
		var ret string
		return ret
	}
	return *o.DataId
}

// GetDataIdOk returns a tuple with the DataId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationBase) GetDataIdOk() (*string, bool) {
	if o == nil || o.DataId == nil {
		return nil, false
	}
	return o.DataId, true
}

// HasDataId returns a boolean if a field has been set.
func (o *NotificationBase) HasDataId() bool {
	if o != nil && o.DataId != nil {
		return true
	}

	return false
}

// SetDataId gets a reference to the given string and assigns it to the DataId field.
func (o *NotificationBase) SetDataId(v string) {
	o.DataId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *NotificationBase) GetCreatedAt() int32 {
	if o == nil || o.CreatedAt == nil {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationBase) GetCreatedAtOk() (*int32, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *NotificationBase) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *NotificationBase) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *NotificationBase) GetLastUpdated() int32 {
	if o == nil || o.LastUpdated == nil {
		var ret int32
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationBase) GetLastUpdatedOk() (*int32, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *NotificationBase) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given int32 and assigns it to the LastUpdated field.
func (o *NotificationBase) SetLastUpdated(v int32) {
	o.LastUpdated = &v
}

// GetRecordStateKey returns the RecordStateKey field value if set, zero value otherwise.
func (o *NotificationBase) GetRecordStateKey() string {
	if o == nil || o.RecordStateKey == nil {
		var ret string
		return ret
	}
	return *o.RecordStateKey
}

// GetRecordStateKeyOk returns a tuple with the RecordStateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationBase) GetRecordStateKeyOk() (*string, bool) {
	if o == nil || o.RecordStateKey == nil {
		return nil, false
	}
	return o.RecordStateKey, true
}

// HasRecordStateKey returns a boolean if a field has been set.
func (o *NotificationBase) HasRecordStateKey() bool {
	if o != nil && o.RecordStateKey != nil {
		return true
	}

	return false
}

// SetRecordStateKey gets a reference to the given string and assigns it to the RecordStateKey field.
func (o *NotificationBase) SetRecordStateKey(v string) {
	o.RecordStateKey = &v
}

// GetRecordStateVal returns the RecordStateVal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationBase) GetRecordStateVal() string {
	if o == nil || o.RecordStateVal.Get() == nil {
		var ret string
		return ret
	}
	return *o.RecordStateVal.Get()
}

// GetRecordStateValOk returns a tuple with the RecordStateVal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationBase) GetRecordStateValOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RecordStateVal.Get(), o.RecordStateVal.IsSet()
}

// HasRecordStateVal returns a boolean if a field has been set.
func (o *NotificationBase) HasRecordStateVal() bool {
	if o != nil && o.RecordStateVal.IsSet() {
		return true
	}

	return false
}

// SetRecordStateVal gets a reference to the given NullableString and assigns it to the RecordStateVal field.
func (o *NotificationBase) SetRecordStateVal(v string) {
	o.RecordStateVal.Set(&v)
}
// SetRecordStateValNil sets the value for RecordStateVal to be an explicit nil
func (o *NotificationBase) SetRecordStateValNil() {
	o.RecordStateVal.Set(nil)
}

// UnsetRecordStateVal ensures that no value is present for RecordStateVal, not even an explicit nil
func (o *NotificationBase) UnsetRecordStateVal() {
	o.RecordStateVal.Unset()
}

// GetTries returns the Tries field value if set, zero value otherwise.
func (o *NotificationBase) GetTries() int32 {
	if o == nil || o.Tries == nil {
		var ret int32
		return ret
	}
	return *o.Tries
}

// GetTriesOk returns a tuple with the Tries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationBase) GetTriesOk() (*int32, bool) {
	if o == nil || o.Tries == nil {
		return nil, false
	}
	return o.Tries, true
}

// HasTries returns a boolean if a field has been set.
func (o *NotificationBase) HasTries() bool {
	if o != nil && o.Tries != nil {
		return true
	}

	return false
}

// SetTries gets a reference to the given int32 and assigns it to the Tries field.
func (o *NotificationBase) SetTries(v int32) {
	o.Tries = &v
}

// GetMaxTries returns the MaxTries field value if set, zero value otherwise.
func (o *NotificationBase) GetMaxTries() int32 {
	if o == nil || o.MaxTries == nil {
		var ret int32
		return ret
	}
	return *o.MaxTries
}

// GetMaxTriesOk returns a tuple with the MaxTries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationBase) GetMaxTriesOk() (*int32, bool) {
	if o == nil || o.MaxTries == nil {
		return nil, false
	}
	return o.MaxTries, true
}

// HasMaxTries returns a boolean if a field has been set.
func (o *NotificationBase) HasMaxTries() bool {
	if o != nil && o.MaxTries != nil {
		return true
	}

	return false
}

// SetMaxTries gets a reference to the given int32 and assigns it to the MaxTries field.
func (o *NotificationBase) SetMaxTries(v int32) {
	o.MaxTries = &v
}

func (o NotificationBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.QueueId != nil {
		toSerialize["queueId"] = o.QueueId
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	if o.DataId != nil {
		toSerialize["dataId"] = o.DataId
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if o.RecordStateKey != nil {
		toSerialize["record_state_key"] = o.RecordStateKey
	}
	if o.RecordStateVal.IsSet() {
		toSerialize["record_state_val"] = o.RecordStateVal.Get()
	}
	if o.Tries != nil {
		toSerialize["tries"] = o.Tries
	}
	if o.MaxTries != nil {
		toSerialize["max_tries"] = o.MaxTries
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationBase struct {
	value *NotificationBase
	isSet bool
}

func (v NullableNotificationBase) Get() *NotificationBase {
	return v.value
}

func (v *NullableNotificationBase) Set(val *NotificationBase) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationBase) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationBase(val *NotificationBase) *NullableNotificationBase {
	return &NullableNotificationBase{value: val, isSet: true}
}

func (v NullableNotificationBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


