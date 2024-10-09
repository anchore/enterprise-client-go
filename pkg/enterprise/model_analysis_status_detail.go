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

// AnalysisStatusDetail The detail of an analysis status change recording which service initiated the state change, when, and which transition
type AnalysisStatusDetail struct {
	FromStatus string `json:"from_status"`
	ToStatus string `json:"to_status"`
	Timestamp string `json:"timestamp"`
	Source ServiceReference `json:"source"`
}

// NewAnalysisStatusDetail instantiates a new AnalysisStatusDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalysisStatusDetail(fromStatus string, toStatus string, timestamp string, source ServiceReference) *AnalysisStatusDetail {
	this := AnalysisStatusDetail{}
	this.FromStatus = fromStatus
	this.ToStatus = toStatus
	this.Timestamp = timestamp
	this.Source = source
	return &this
}

// NewAnalysisStatusDetailWithDefaults instantiates a new AnalysisStatusDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalysisStatusDetailWithDefaults() *AnalysisStatusDetail {
	this := AnalysisStatusDetail{}
	return &this
}

// GetFromStatus returns the FromStatus field value
func (o *AnalysisStatusDetail) GetFromStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromStatus
}

// GetFromStatusOk returns a tuple with the FromStatus field value
// and a boolean to check if the value has been set.
func (o *AnalysisStatusDetail) GetFromStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromStatus, true
}

// SetFromStatus sets field value
func (o *AnalysisStatusDetail) SetFromStatus(v string) {
	o.FromStatus = v
}

// GetToStatus returns the ToStatus field value
func (o *AnalysisStatusDetail) GetToStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToStatus
}

// GetToStatusOk returns a tuple with the ToStatus field value
// and a boolean to check if the value has been set.
func (o *AnalysisStatusDetail) GetToStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToStatus, true
}

// SetToStatus sets field value
func (o *AnalysisStatusDetail) SetToStatus(v string) {
	o.ToStatus = v
}

// GetTimestamp returns the Timestamp field value
func (o *AnalysisStatusDetail) GetTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *AnalysisStatusDetail) GetTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *AnalysisStatusDetail) SetTimestamp(v string) {
	o.Timestamp = v
}

// GetSource returns the Source field value
func (o *AnalysisStatusDetail) GetSource() ServiceReference {
	if o == nil {
		var ret ServiceReference
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *AnalysisStatusDetail) GetSourceOk() (*ServiceReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *AnalysisStatusDetail) SetSource(v ServiceReference) {
	o.Source = v
}

func (o AnalysisStatusDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["from_status"] = o.FromStatus
	}
	if true {
		toSerialize["to_status"] = o.ToStatus
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	if true {
		toSerialize["source"] = o.Source
	}
	return json.Marshal(toSerialize)
}

type NullableAnalysisStatusDetail struct {
	value *AnalysisStatusDetail
	isSet bool
}

func (v NullableAnalysisStatusDetail) Get() *AnalysisStatusDetail {
	return v.value
}

func (v *NullableAnalysisStatusDetail) Set(val *AnalysisStatusDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalysisStatusDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalysisStatusDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalysisStatusDetail(val *AnalysisStatusDetail) *NullableAnalysisStatusDetail {
	return &NullableAnalysisStatusDetail{value: val, isSet: true}
}

func (v NullableAnalysisStatusDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalysisStatusDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


