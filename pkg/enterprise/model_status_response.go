/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.1.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// StatusResponse System status response
type StatusResponse struct {
	Available *bool `json:"available,omitempty"`
	Busy *bool `json:"busy,omitempty"`
	Up *bool `json:"up,omitempty"`
	Message *string `json:"message,omitempty"`
	Version *string `json:"version,omitempty"`
	DbVersion *string `json:"db_version,omitempty"`
	Detail interface{} `json:"detail,omitempty"`
}

// NewStatusResponse instantiates a new StatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusResponse() *StatusResponse {
	this := StatusResponse{}
	return &this
}

// NewStatusResponseWithDefaults instantiates a new StatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusResponseWithDefaults() *StatusResponse {
	this := StatusResponse{}
	return &this
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *StatusResponse) GetAvailable() bool {
	if o == nil || o.Available == nil {
		var ret bool
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusResponse) GetAvailableOk() (*bool, bool) {
	if o == nil || o.Available == nil {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *StatusResponse) HasAvailable() bool {
	if o != nil && o.Available != nil {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given bool and assigns it to the Available field.
func (o *StatusResponse) SetAvailable(v bool) {
	o.Available = &v
}

// GetBusy returns the Busy field value if set, zero value otherwise.
func (o *StatusResponse) GetBusy() bool {
	if o == nil || o.Busy == nil {
		var ret bool
		return ret
	}
	return *o.Busy
}

// GetBusyOk returns a tuple with the Busy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusResponse) GetBusyOk() (*bool, bool) {
	if o == nil || o.Busy == nil {
		return nil, false
	}
	return o.Busy, true
}

// HasBusy returns a boolean if a field has been set.
func (o *StatusResponse) HasBusy() bool {
	if o != nil && o.Busy != nil {
		return true
	}

	return false
}

// SetBusy gets a reference to the given bool and assigns it to the Busy field.
func (o *StatusResponse) SetBusy(v bool) {
	o.Busy = &v
}

// GetUp returns the Up field value if set, zero value otherwise.
func (o *StatusResponse) GetUp() bool {
	if o == nil || o.Up == nil {
		var ret bool
		return ret
	}
	return *o.Up
}

// GetUpOk returns a tuple with the Up field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusResponse) GetUpOk() (*bool, bool) {
	if o == nil || o.Up == nil {
		return nil, false
	}
	return o.Up, true
}

// HasUp returns a boolean if a field has been set.
func (o *StatusResponse) HasUp() bool {
	if o != nil && o.Up != nil {
		return true
	}

	return false
}

// SetUp gets a reference to the given bool and assigns it to the Up field.
func (o *StatusResponse) SetUp(v bool) {
	o.Up = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *StatusResponse) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusResponse) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *StatusResponse) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *StatusResponse) SetMessage(v string) {
	o.Message = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *StatusResponse) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusResponse) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *StatusResponse) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *StatusResponse) SetVersion(v string) {
	o.Version = &v
}

// GetDbVersion returns the DbVersion field value if set, zero value otherwise.
func (o *StatusResponse) GetDbVersion() string {
	if o == nil || o.DbVersion == nil {
		var ret string
		return ret
	}
	return *o.DbVersion
}

// GetDbVersionOk returns a tuple with the DbVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusResponse) GetDbVersionOk() (*string, bool) {
	if o == nil || o.DbVersion == nil {
		return nil, false
	}
	return o.DbVersion, true
}

// HasDbVersion returns a boolean if a field has been set.
func (o *StatusResponse) HasDbVersion() bool {
	if o != nil && o.DbVersion != nil {
		return true
	}

	return false
}

// SetDbVersion gets a reference to the given string and assigns it to the DbVersion field.
func (o *StatusResponse) SetDbVersion(v string) {
	o.DbVersion = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *StatusResponse) GetDetail() interface{} {
	if o == nil || o.Detail == nil {
		var ret interface{}
		return ret
	}
	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusResponse) GetDetailOk() (interface{}, bool) {
	if o == nil || o.Detail == nil {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *StatusResponse) HasDetail() bool {
	if o != nil && o.Detail != nil {
		return true
	}

	return false
}

// SetDetail gets a reference to the given interface{} and assigns it to the Detail field.
func (o *StatusResponse) SetDetail(v interface{}) {
	o.Detail = v
}

func (o StatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Available != nil {
		toSerialize["available"] = o.Available
	}
	if o.Busy != nil {
		toSerialize["busy"] = o.Busy
	}
	if o.Up != nil {
		toSerialize["up"] = o.Up
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.DbVersion != nil {
		toSerialize["db_version"] = o.DbVersion
	}
	if o.Detail != nil {
		toSerialize["detail"] = o.Detail
	}
	return json.Marshal(toSerialize)
}

type NullableStatusResponse struct {
	value *StatusResponse
	isSet bool
}

func (v NullableStatusResponse) Get() *StatusResponse {
	return v.value
}

func (v *NullableStatusResponse) Set(val *StatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusResponse(val *StatusResponse) *NullableStatusResponse {
	return &NullableStatusResponse{value: val, isSet: true}
}

func (v NullableStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


