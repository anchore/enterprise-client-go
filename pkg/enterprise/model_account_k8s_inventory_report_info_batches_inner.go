/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.12.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the AccountK8sInventoryReportInfoBatchesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountK8sInventoryReportInfoBatchesInner{}

// AccountK8sInventoryReportInfoBatchesInner struct for AccountK8sInventoryReportInfoBatchesInner
type AccountK8sInventoryReportInfoBatchesInner struct {
	// Index of the batch
	BatchIndex int32 `json:"batch_index"`
	// Timestamp when the batch was sent
	SendTimestamp time.Time `json:"send_timestamp"`
	// Error message if the sending was unsuccessful
	Error *string `json:"error,omitempty"`
}

type _AccountK8sInventoryReportInfoBatchesInner AccountK8sInventoryReportInfoBatchesInner

// NewAccountK8sInventoryReportInfoBatchesInner instantiates a new AccountK8sInventoryReportInfoBatchesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountK8sInventoryReportInfoBatchesInner(batchIndex int32, sendTimestamp time.Time) *AccountK8sInventoryReportInfoBatchesInner {
	this := AccountK8sInventoryReportInfoBatchesInner{}
	this.BatchIndex = batchIndex
	this.SendTimestamp = sendTimestamp
	return &this
}

// NewAccountK8sInventoryReportInfoBatchesInnerWithDefaults instantiates a new AccountK8sInventoryReportInfoBatchesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountK8sInventoryReportInfoBatchesInnerWithDefaults() *AccountK8sInventoryReportInfoBatchesInner {
	this := AccountK8sInventoryReportInfoBatchesInner{}
	return &this
}

// GetBatchIndex returns the BatchIndex field value
func (o *AccountK8sInventoryReportInfoBatchesInner) GetBatchIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BatchIndex
}

// GetBatchIndexOk returns a tuple with the BatchIndex field value
// and a boolean to check if the value has been set.
func (o *AccountK8sInventoryReportInfoBatchesInner) GetBatchIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BatchIndex, true
}

// SetBatchIndex sets field value
func (o *AccountK8sInventoryReportInfoBatchesInner) SetBatchIndex(v int32) {
	o.BatchIndex = v
}

// GetSendTimestamp returns the SendTimestamp field value
func (o *AccountK8sInventoryReportInfoBatchesInner) GetSendTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.SendTimestamp
}

// GetSendTimestampOk returns a tuple with the SendTimestamp field value
// and a boolean to check if the value has been set.
func (o *AccountK8sInventoryReportInfoBatchesInner) GetSendTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SendTimestamp, true
}

// SetSendTimestamp sets field value
func (o *AccountK8sInventoryReportInfoBatchesInner) SetSendTimestamp(v time.Time) {
	o.SendTimestamp = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *AccountK8sInventoryReportInfoBatchesInner) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountK8sInventoryReportInfoBatchesInner) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *AccountK8sInventoryReportInfoBatchesInner) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *AccountK8sInventoryReportInfoBatchesInner) SetError(v string) {
	o.Error = &v
}

func (o AccountK8sInventoryReportInfoBatchesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountK8sInventoryReportInfoBatchesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["batch_index"] = o.BatchIndex
	toSerialize["send_timestamp"] = o.SendTimestamp
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

func (o *AccountK8sInventoryReportInfoBatchesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"batch_index",
		"send_timestamp",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAccountK8sInventoryReportInfoBatchesInner := _AccountK8sInventoryReportInfoBatchesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccountK8sInventoryReportInfoBatchesInner)

	if err != nil {
		return err
	}

	*o = AccountK8sInventoryReportInfoBatchesInner(varAccountK8sInventoryReportInfoBatchesInner)

	return err
}

type NullableAccountK8sInventoryReportInfoBatchesInner struct {
	value *AccountK8sInventoryReportInfoBatchesInner
	isSet bool
}

func (v NullableAccountK8sInventoryReportInfoBatchesInner) Get() *AccountK8sInventoryReportInfoBatchesInner {
	return v.value
}

func (v *NullableAccountK8sInventoryReportInfoBatchesInner) Set(val *AccountK8sInventoryReportInfoBatchesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountK8sInventoryReportInfoBatchesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountK8sInventoryReportInfoBatchesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountK8sInventoryReportInfoBatchesInner(val *AccountK8sInventoryReportInfoBatchesInner) *NullableAccountK8sInventoryReportInfoBatchesInner {
	return &NullableAccountK8sInventoryReportInfoBatchesInner{value: val, isSet: true}
}

func (v NullableAccountK8sInventoryReportInfoBatchesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountK8sInventoryReportInfoBatchesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


