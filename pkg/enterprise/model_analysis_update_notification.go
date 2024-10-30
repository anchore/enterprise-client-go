/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.8.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the AnalysisUpdateNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnalysisUpdateNotification{}

// AnalysisUpdateNotification struct for AnalysisUpdateNotification
type AnalysisUpdateNotification struct {
	QueueId *string `json:"queue_id,omitempty"`
	AccountName *string `json:"account_name,omitempty"`
	DataId *string `json:"data_id,omitempty"`
	CreatedAt *int32 `json:"created_at,omitempty"`
	LastUpdated *int32 `json:"last_updated,omitempty"`
	RecordStateKey *string `json:"record_state_key,omitempty"`
	RecordStateVal NullableString `json:"record_state_val,omitempty"`
	Tries *int32 `json:"tries,omitempty"`
	MaxTries *int32 `json:"max_tries,omitempty"`
	Data *AnalysisUpdateNotificationData `json:"data,omitempty"`
}

// NewAnalysisUpdateNotification instantiates a new AnalysisUpdateNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalysisUpdateNotification() *AnalysisUpdateNotification {
	this := AnalysisUpdateNotification{}
	var recordStateKey string = "active"
	this.RecordStateKey = &recordStateKey
	return &this
}

// NewAnalysisUpdateNotificationWithDefaults instantiates a new AnalysisUpdateNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalysisUpdateNotificationWithDefaults() *AnalysisUpdateNotification {
	this := AnalysisUpdateNotification{}
	var recordStateKey string = "active"
	this.RecordStateKey = &recordStateKey
	return &this
}

// GetQueueId returns the QueueId field value if set, zero value otherwise.
func (o *AnalysisUpdateNotification) GetQueueId() string {
	if o == nil || IsNil(o.QueueId) {
		var ret string
		return ret
	}
	return *o.QueueId
}

// GetQueueIdOk returns a tuple with the QueueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisUpdateNotification) GetQueueIdOk() (*string, bool) {
	if o == nil || IsNil(o.QueueId) {
		return nil, false
	}
	return o.QueueId, true
}

// HasQueueId returns a boolean if a field has been set.
func (o *AnalysisUpdateNotification) HasQueueId() bool {
	if o != nil && !IsNil(o.QueueId) {
		return true
	}

	return false
}

// SetQueueId gets a reference to the given string and assigns it to the QueueId field.
func (o *AnalysisUpdateNotification) SetQueueId(v string) {
	o.QueueId = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *AnalysisUpdateNotification) GetAccountName() string {
	if o == nil || IsNil(o.AccountName) {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisUpdateNotification) GetAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.AccountName) {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *AnalysisUpdateNotification) HasAccountName() bool {
	if o != nil && !IsNil(o.AccountName) {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *AnalysisUpdateNotification) SetAccountName(v string) {
	o.AccountName = &v
}

// GetDataId returns the DataId field value if set, zero value otherwise.
func (o *AnalysisUpdateNotification) GetDataId() string {
	if o == nil || IsNil(o.DataId) {
		var ret string
		return ret
	}
	return *o.DataId
}

// GetDataIdOk returns a tuple with the DataId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisUpdateNotification) GetDataIdOk() (*string, bool) {
	if o == nil || IsNil(o.DataId) {
		return nil, false
	}
	return o.DataId, true
}

// HasDataId returns a boolean if a field has been set.
func (o *AnalysisUpdateNotification) HasDataId() bool {
	if o != nil && !IsNil(o.DataId) {
		return true
	}

	return false
}

// SetDataId gets a reference to the given string and assigns it to the DataId field.
func (o *AnalysisUpdateNotification) SetDataId(v string) {
	o.DataId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AnalysisUpdateNotification) GetCreatedAt() int32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisUpdateNotification) GetCreatedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AnalysisUpdateNotification) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *AnalysisUpdateNotification) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *AnalysisUpdateNotification) GetLastUpdated() int32 {
	if o == nil || IsNil(o.LastUpdated) {
		var ret int32
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisUpdateNotification) GetLastUpdatedOk() (*int32, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *AnalysisUpdateNotification) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given int32 and assigns it to the LastUpdated field.
func (o *AnalysisUpdateNotification) SetLastUpdated(v int32) {
	o.LastUpdated = &v
}

// GetRecordStateKey returns the RecordStateKey field value if set, zero value otherwise.
func (o *AnalysisUpdateNotification) GetRecordStateKey() string {
	if o == nil || IsNil(o.RecordStateKey) {
		var ret string
		return ret
	}
	return *o.RecordStateKey
}

// GetRecordStateKeyOk returns a tuple with the RecordStateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisUpdateNotification) GetRecordStateKeyOk() (*string, bool) {
	if o == nil || IsNil(o.RecordStateKey) {
		return nil, false
	}
	return o.RecordStateKey, true
}

// HasRecordStateKey returns a boolean if a field has been set.
func (o *AnalysisUpdateNotification) HasRecordStateKey() bool {
	if o != nil && !IsNil(o.RecordStateKey) {
		return true
	}

	return false
}

// SetRecordStateKey gets a reference to the given string and assigns it to the RecordStateKey field.
func (o *AnalysisUpdateNotification) SetRecordStateKey(v string) {
	o.RecordStateKey = &v
}

// GetRecordStateVal returns the RecordStateVal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnalysisUpdateNotification) GetRecordStateVal() string {
	if o == nil || IsNil(o.RecordStateVal.Get()) {
		var ret string
		return ret
	}
	return *o.RecordStateVal.Get()
}

// GetRecordStateValOk returns a tuple with the RecordStateVal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnalysisUpdateNotification) GetRecordStateValOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecordStateVal.Get(), o.RecordStateVal.IsSet()
}

// HasRecordStateVal returns a boolean if a field has been set.
func (o *AnalysisUpdateNotification) HasRecordStateVal() bool {
	if o != nil && o.RecordStateVal.IsSet() {
		return true
	}

	return false
}

// SetRecordStateVal gets a reference to the given NullableString and assigns it to the RecordStateVal field.
func (o *AnalysisUpdateNotification) SetRecordStateVal(v string) {
	o.RecordStateVal.Set(&v)
}
// SetRecordStateValNil sets the value for RecordStateVal to be an explicit nil
func (o *AnalysisUpdateNotification) SetRecordStateValNil() {
	o.RecordStateVal.Set(nil)
}

// UnsetRecordStateVal ensures that no value is present for RecordStateVal, not even an explicit nil
func (o *AnalysisUpdateNotification) UnsetRecordStateVal() {
	o.RecordStateVal.Unset()
}

// GetTries returns the Tries field value if set, zero value otherwise.
func (o *AnalysisUpdateNotification) GetTries() int32 {
	if o == nil || IsNil(o.Tries) {
		var ret int32
		return ret
	}
	return *o.Tries
}

// GetTriesOk returns a tuple with the Tries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisUpdateNotification) GetTriesOk() (*int32, bool) {
	if o == nil || IsNil(o.Tries) {
		return nil, false
	}
	return o.Tries, true
}

// HasTries returns a boolean if a field has been set.
func (o *AnalysisUpdateNotification) HasTries() bool {
	if o != nil && !IsNil(o.Tries) {
		return true
	}

	return false
}

// SetTries gets a reference to the given int32 and assigns it to the Tries field.
func (o *AnalysisUpdateNotification) SetTries(v int32) {
	o.Tries = &v
}

// GetMaxTries returns the MaxTries field value if set, zero value otherwise.
func (o *AnalysisUpdateNotification) GetMaxTries() int32 {
	if o == nil || IsNil(o.MaxTries) {
		var ret int32
		return ret
	}
	return *o.MaxTries
}

// GetMaxTriesOk returns a tuple with the MaxTries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisUpdateNotification) GetMaxTriesOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxTries) {
		return nil, false
	}
	return o.MaxTries, true
}

// HasMaxTries returns a boolean if a field has been set.
func (o *AnalysisUpdateNotification) HasMaxTries() bool {
	if o != nil && !IsNil(o.MaxTries) {
		return true
	}

	return false
}

// SetMaxTries gets a reference to the given int32 and assigns it to the MaxTries field.
func (o *AnalysisUpdateNotification) SetMaxTries(v int32) {
	o.MaxTries = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AnalysisUpdateNotification) GetData() AnalysisUpdateNotificationData {
	if o == nil || IsNil(o.Data) {
		var ret AnalysisUpdateNotificationData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisUpdateNotification) GetDataOk() (*AnalysisUpdateNotificationData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AnalysisUpdateNotification) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AnalysisUpdateNotificationData and assigns it to the Data field.
func (o *AnalysisUpdateNotification) SetData(v AnalysisUpdateNotificationData) {
	o.Data = &v
}

func (o AnalysisUpdateNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnalysisUpdateNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.QueueId) {
		toSerialize["queue_id"] = o.QueueId
	}
	if !IsNil(o.AccountName) {
		toSerialize["account_name"] = o.AccountName
	}
	if !IsNil(o.DataId) {
		toSerialize["data_id"] = o.DataId
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if !IsNil(o.RecordStateKey) {
		toSerialize["record_state_key"] = o.RecordStateKey
	}
	if o.RecordStateVal.IsSet() {
		toSerialize["record_state_val"] = o.RecordStateVal.Get()
	}
	if !IsNil(o.Tries) {
		toSerialize["tries"] = o.Tries
	}
	if !IsNil(o.MaxTries) {
		toSerialize["max_tries"] = o.MaxTries
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAnalysisUpdateNotification struct {
	value *AnalysisUpdateNotification
	isSet bool
}

func (v NullableAnalysisUpdateNotification) Get() *AnalysisUpdateNotification {
	return v.value
}

func (v *NullableAnalysisUpdateNotification) Set(val *AnalysisUpdateNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalysisUpdateNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalysisUpdateNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalysisUpdateNotification(val *AnalysisUpdateNotification) *NullableAnalysisUpdateNotification {
	return &NullableAnalysisUpdateNotification{value: val, isSet: true}
}

func (v NullableAnalysisUpdateNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalysisUpdateNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


