/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.5.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// AnalysisUpdateNotificationAllOf The Notification Object definition for Analysis Update Notifications
type AnalysisUpdateNotificationAllOf struct {
	Data *AnalysisUpdateNotificationData `json:"data,omitempty"`
}

// NewAnalysisUpdateNotificationAllOf instantiates a new AnalysisUpdateNotificationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalysisUpdateNotificationAllOf() *AnalysisUpdateNotificationAllOf {
	this := AnalysisUpdateNotificationAllOf{}
	return &this
}

// NewAnalysisUpdateNotificationAllOfWithDefaults instantiates a new AnalysisUpdateNotificationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalysisUpdateNotificationAllOfWithDefaults() *AnalysisUpdateNotificationAllOf {
	this := AnalysisUpdateNotificationAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AnalysisUpdateNotificationAllOf) GetData() AnalysisUpdateNotificationData {
	if o == nil || o.Data == nil {
		var ret AnalysisUpdateNotificationData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisUpdateNotificationAllOf) GetDataOk() (*AnalysisUpdateNotificationData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AnalysisUpdateNotificationAllOf) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given AnalysisUpdateNotificationData and assigns it to the Data field.
func (o *AnalysisUpdateNotificationAllOf) SetData(v AnalysisUpdateNotificationData) {
	o.Data = &v
}

func (o AnalysisUpdateNotificationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableAnalysisUpdateNotificationAllOf struct {
	value *AnalysisUpdateNotificationAllOf
	isSet bool
}

func (v NullableAnalysisUpdateNotificationAllOf) Get() *AnalysisUpdateNotificationAllOf {
	return v.value
}

func (v *NullableAnalysisUpdateNotificationAllOf) Set(val *AnalysisUpdateNotificationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalysisUpdateNotificationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalysisUpdateNotificationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalysisUpdateNotificationAllOf(val *AnalysisUpdateNotificationAllOf) *NullableAnalysisUpdateNotificationAllOf {
	return &NullableAnalysisUpdateNotificationAllOf{value: val, isSet: true}
}

func (v NullableAnalysisUpdateNotificationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalysisUpdateNotificationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


