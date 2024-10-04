/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.7.2
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the FeedDataRecords type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeedDataRecords{}

// FeedDataRecords A list of data records
type FeedDataRecords struct {
	Records []FeedDataRecord `json:"records,omitempty"`
}

// NewFeedDataRecords instantiates a new FeedDataRecords object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeedDataRecords() *FeedDataRecords {
	this := FeedDataRecords{}
	return &this
}

// NewFeedDataRecordsWithDefaults instantiates a new FeedDataRecords object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeedDataRecordsWithDefaults() *FeedDataRecords {
	this := FeedDataRecords{}
	return &this
}

// GetRecords returns the Records field value if set, zero value otherwise.
func (o *FeedDataRecords) GetRecords() []FeedDataRecord {
	if o == nil || IsNil(o.Records) {
		var ret []FeedDataRecord
		return ret
	}
	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedDataRecords) GetRecordsOk() ([]FeedDataRecord, bool) {
	if o == nil || IsNil(o.Records) {
		return nil, false
	}
	return o.Records, true
}

// HasRecords returns a boolean if a field has been set.
func (o *FeedDataRecords) HasRecords() bool {
	if o != nil && !IsNil(o.Records) {
		return true
	}

	return false
}

// SetRecords gets a reference to the given []FeedDataRecord and assigns it to the Records field.
func (o *FeedDataRecords) SetRecords(v []FeedDataRecord) {
	o.Records = v
}

func (o FeedDataRecords) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeedDataRecords) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Records) {
		toSerialize["records"] = o.Records
	}
	return toSerialize, nil
}

type NullableFeedDataRecords struct {
	value *FeedDataRecords
	isSet bool
}

func (v NullableFeedDataRecords) Get() *FeedDataRecords {
	return v.value
}

func (v *NullableFeedDataRecords) Set(val *FeedDataRecords) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedDataRecords) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedDataRecords) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedDataRecords(val *FeedDataRecords) *NullableFeedDataRecords {
	return &NullableFeedDataRecords{value: val, isSet: true}
}

func (v NullableFeedDataRecords) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedDataRecords) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

