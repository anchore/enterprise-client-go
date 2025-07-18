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
)

// checks if the ReportingStatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportingStatusResponse{}

// ReportingStatusResponse System status response
type ReportingStatusResponse struct {
	Busy *bool `json:"busy,omitempty"`
	Up *bool `json:"up,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewReportingStatusResponse instantiates a new ReportingStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportingStatusResponse() *ReportingStatusResponse {
	this := ReportingStatusResponse{}
	return &this
}

// NewReportingStatusResponseWithDefaults instantiates a new ReportingStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportingStatusResponseWithDefaults() *ReportingStatusResponse {
	this := ReportingStatusResponse{}
	return &this
}

// GetBusy returns the Busy field value if set, zero value otherwise.
func (o *ReportingStatusResponse) GetBusy() bool {
	if o == nil || IsNil(o.Busy) {
		var ret bool
		return ret
	}
	return *o.Busy
}

// GetBusyOk returns a tuple with the Busy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingStatusResponse) GetBusyOk() (*bool, bool) {
	if o == nil || IsNil(o.Busy) {
		return nil, false
	}
	return o.Busy, true
}

// HasBusy returns a boolean if a field has been set.
func (o *ReportingStatusResponse) HasBusy() bool {
	if o != nil && !IsNil(o.Busy) {
		return true
	}

	return false
}

// SetBusy gets a reference to the given bool and assigns it to the Busy field.
func (o *ReportingStatusResponse) SetBusy(v bool) {
	o.Busy = &v
}

// GetUp returns the Up field value if set, zero value otherwise.
func (o *ReportingStatusResponse) GetUp() bool {
	if o == nil || IsNil(o.Up) {
		var ret bool
		return ret
	}
	return *o.Up
}

// GetUpOk returns a tuple with the Up field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingStatusResponse) GetUpOk() (*bool, bool) {
	if o == nil || IsNil(o.Up) {
		return nil, false
	}
	return o.Up, true
}

// HasUp returns a boolean if a field has been set.
func (o *ReportingStatusResponse) HasUp() bool {
	if o != nil && !IsNil(o.Up) {
		return true
	}

	return false
}

// SetUp gets a reference to the given bool and assigns it to the Up field.
func (o *ReportingStatusResponse) SetUp(v bool) {
	o.Up = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ReportingStatusResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingStatusResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ReportingStatusResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ReportingStatusResponse) SetMessage(v string) {
	o.Message = &v
}

func (o ReportingStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportingStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Busy) {
		toSerialize["busy"] = o.Busy
	}
	if !IsNil(o.Up) {
		toSerialize["up"] = o.Up
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableReportingStatusResponse struct {
	value *ReportingStatusResponse
	isSet bool
}

func (v NullableReportingStatusResponse) Get() *ReportingStatusResponse {
	return v.value
}

func (v *NullableReportingStatusResponse) Set(val *ReportingStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReportingStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReportingStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportingStatusResponse(val *ReportingStatusResponse) *NullableReportingStatusResponse {
	return &NullableReportingStatusResponse{value: val, isSet: true}
}

func (v NullableReportingStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportingStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


