/*
Anchore Engine API Server

This is the Anchore Engine API. Provides the primary external API for users of the service.

API version: 0.6.0
Contact: nurmi@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package engine

import (
	"encoding/json"
)

// VulnUpdateNotification struct for VulnUpdateNotification
type VulnUpdateNotification struct {
	QueueId *string `json:"queue_id,omitempty"`
	AccountName *string `json:"account_name,omitempty"`
	DataId *string `json:"data_id,omitempty"`
	CreatedAt *int32 `json:"created_at,omitempty"`
	LastUpdated *int32 `json:"last_updated,omitempty"`
	RecordStateKey *string `json:"record_state_key,omitempty"`
	RecordStateVal NullableString `json:"record_state_val,omitempty"`
	Tries *int32 `json:"tries,omitempty"`
	MaxTries *int32 `json:"max_tries,omitempty"`
	Data *VulnUpdateNotificationData `json:"data,omitempty"`
}

// NewVulnUpdateNotification instantiates a new VulnUpdateNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVulnUpdateNotification() *VulnUpdateNotification {
	this := VulnUpdateNotification{}
	var recordStateKey string = "active"
	this.RecordStateKey = &recordStateKey
	return &this
}

// NewVulnUpdateNotificationWithDefaults instantiates a new VulnUpdateNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVulnUpdateNotificationWithDefaults() *VulnUpdateNotification {
	this := VulnUpdateNotification{}
	var recordStateKey string = "active"
	this.RecordStateKey = &recordStateKey
	return &this
}

// GetQueueId returns the QueueId field value if set, zero value otherwise.
func (o *VulnUpdateNotification) GetQueueId() string {
	if o == nil || o.QueueId == nil {
		var ret string
		return ret
	}
	return *o.QueueId
}

// GetQueueIdOk returns a tuple with the QueueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnUpdateNotification) GetQueueIdOk() (*string, bool) {
	if o == nil || o.QueueId == nil {
		return nil, false
	}
	return o.QueueId, true
}

// HasQueueId returns a boolean if a field has been set.
func (o *VulnUpdateNotification) HasQueueId() bool {
	if o != nil && o.QueueId != nil {
		return true
	}

	return false
}

// SetQueueId gets a reference to the given string and assigns it to the QueueId field.
func (o *VulnUpdateNotification) SetQueueId(v string) {
	o.QueueId = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *VulnUpdateNotification) GetAccountName() string {
	if o == nil || o.AccountName == nil {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnUpdateNotification) GetAccountNameOk() (*string, bool) {
	if o == nil || o.AccountName == nil {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *VulnUpdateNotification) HasAccountName() bool {
	if o != nil && o.AccountName != nil {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *VulnUpdateNotification) SetAccountName(v string) {
	o.AccountName = &v
}

// GetDataId returns the DataId field value if set, zero value otherwise.
func (o *VulnUpdateNotification) GetDataId() string {
	if o == nil || o.DataId == nil {
		var ret string
		return ret
	}
	return *o.DataId
}

// GetDataIdOk returns a tuple with the DataId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnUpdateNotification) GetDataIdOk() (*string, bool) {
	if o == nil || o.DataId == nil {
		return nil, false
	}
	return o.DataId, true
}

// HasDataId returns a boolean if a field has been set.
func (o *VulnUpdateNotification) HasDataId() bool {
	if o != nil && o.DataId != nil {
		return true
	}

	return false
}

// SetDataId gets a reference to the given string and assigns it to the DataId field.
func (o *VulnUpdateNotification) SetDataId(v string) {
	o.DataId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *VulnUpdateNotification) GetCreatedAt() int32 {
	if o == nil || o.CreatedAt == nil {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnUpdateNotification) GetCreatedAtOk() (*int32, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *VulnUpdateNotification) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *VulnUpdateNotification) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *VulnUpdateNotification) GetLastUpdated() int32 {
	if o == nil || o.LastUpdated == nil {
		var ret int32
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnUpdateNotification) GetLastUpdatedOk() (*int32, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *VulnUpdateNotification) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given int32 and assigns it to the LastUpdated field.
func (o *VulnUpdateNotification) SetLastUpdated(v int32) {
	o.LastUpdated = &v
}

// GetRecordStateKey returns the RecordStateKey field value if set, zero value otherwise.
func (o *VulnUpdateNotification) GetRecordStateKey() string {
	if o == nil || o.RecordStateKey == nil {
		var ret string
		return ret
	}
	return *o.RecordStateKey
}

// GetRecordStateKeyOk returns a tuple with the RecordStateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnUpdateNotification) GetRecordStateKeyOk() (*string, bool) {
	if o == nil || o.RecordStateKey == nil {
		return nil, false
	}
	return o.RecordStateKey, true
}

// HasRecordStateKey returns a boolean if a field has been set.
func (o *VulnUpdateNotification) HasRecordStateKey() bool {
	if o != nil && o.RecordStateKey != nil {
		return true
	}

	return false
}

// SetRecordStateKey gets a reference to the given string and assigns it to the RecordStateKey field.
func (o *VulnUpdateNotification) SetRecordStateKey(v string) {
	o.RecordStateKey = &v
}

// GetRecordStateVal returns the RecordStateVal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VulnUpdateNotification) GetRecordStateVal() string {
	if o == nil || o.RecordStateVal.Get() == nil {
		var ret string
		return ret
	}
	return *o.RecordStateVal.Get()
}

// GetRecordStateValOk returns a tuple with the RecordStateVal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VulnUpdateNotification) GetRecordStateValOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RecordStateVal.Get(), o.RecordStateVal.IsSet()
}

// HasRecordStateVal returns a boolean if a field has been set.
func (o *VulnUpdateNotification) HasRecordStateVal() bool {
	if o != nil && o.RecordStateVal.IsSet() {
		return true
	}

	return false
}

// SetRecordStateVal gets a reference to the given NullableString and assigns it to the RecordStateVal field.
func (o *VulnUpdateNotification) SetRecordStateVal(v string) {
	o.RecordStateVal.Set(&v)
}
// SetRecordStateValNil sets the value for RecordStateVal to be an explicit nil
func (o *VulnUpdateNotification) SetRecordStateValNil() {
	o.RecordStateVal.Set(nil)
}

// UnsetRecordStateVal ensures that no value is present for RecordStateVal, not even an explicit nil
func (o *VulnUpdateNotification) UnsetRecordStateVal() {
	o.RecordStateVal.Unset()
}

// GetTries returns the Tries field value if set, zero value otherwise.
func (o *VulnUpdateNotification) GetTries() int32 {
	if o == nil || o.Tries == nil {
		var ret int32
		return ret
	}
	return *o.Tries
}

// GetTriesOk returns a tuple with the Tries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnUpdateNotification) GetTriesOk() (*int32, bool) {
	if o == nil || o.Tries == nil {
		return nil, false
	}
	return o.Tries, true
}

// HasTries returns a boolean if a field has been set.
func (o *VulnUpdateNotification) HasTries() bool {
	if o != nil && o.Tries != nil {
		return true
	}

	return false
}

// SetTries gets a reference to the given int32 and assigns it to the Tries field.
func (o *VulnUpdateNotification) SetTries(v int32) {
	o.Tries = &v
}

// GetMaxTries returns the MaxTries field value if set, zero value otherwise.
func (o *VulnUpdateNotification) GetMaxTries() int32 {
	if o == nil || o.MaxTries == nil {
		var ret int32
		return ret
	}
	return *o.MaxTries
}

// GetMaxTriesOk returns a tuple with the MaxTries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnUpdateNotification) GetMaxTriesOk() (*int32, bool) {
	if o == nil || o.MaxTries == nil {
		return nil, false
	}
	return o.MaxTries, true
}

// HasMaxTries returns a boolean if a field has been set.
func (o *VulnUpdateNotification) HasMaxTries() bool {
	if o != nil && o.MaxTries != nil {
		return true
	}

	return false
}

// SetMaxTries gets a reference to the given int32 and assigns it to the MaxTries field.
func (o *VulnUpdateNotification) SetMaxTries(v int32) {
	o.MaxTries = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *VulnUpdateNotification) GetData() VulnUpdateNotificationData {
	if o == nil || o.Data == nil {
		var ret VulnUpdateNotificationData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnUpdateNotification) GetDataOk() (*VulnUpdateNotificationData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *VulnUpdateNotification) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given VulnUpdateNotificationData and assigns it to the Data field.
func (o *VulnUpdateNotification) SetData(v VulnUpdateNotificationData) {
	o.Data = &v
}

func (o VulnUpdateNotification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.QueueId != nil {
		toSerialize["queue_id"] = o.QueueId
	}
	if o.AccountName != nil {
		toSerialize["account_name"] = o.AccountName
	}
	if o.DataId != nil {
		toSerialize["data_id"] = o.DataId
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
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableVulnUpdateNotification struct {
	value *VulnUpdateNotification
	isSet bool
}

func (v NullableVulnUpdateNotification) Get() *VulnUpdateNotification {
	return v.value
}

func (v *NullableVulnUpdateNotification) Set(val *VulnUpdateNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableVulnUpdateNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableVulnUpdateNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVulnUpdateNotification(val *VulnUpdateNotification) *NullableVulnUpdateNotification {
	return &NullableVulnUpdateNotification{value: val, isSet: true}
}

func (v NullableVulnUpdateNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVulnUpdateNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


