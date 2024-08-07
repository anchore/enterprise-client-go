/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.6.1
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// NotificationTeamsTestResult struct for NotificationTeamsTestResult
type NotificationTeamsTestResult struct {
	Status *string `json:"status,omitempty"`
	Response *string `json:"response,omitempty"`
}

// NewNotificationTeamsTestResult instantiates a new NotificationTeamsTestResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationTeamsTestResult() *NotificationTeamsTestResult {
	this := NotificationTeamsTestResult{}
	return &this
}

// NewNotificationTeamsTestResultWithDefaults instantiates a new NotificationTeamsTestResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationTeamsTestResultWithDefaults() *NotificationTeamsTestResult {
	this := NotificationTeamsTestResult{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NotificationTeamsTestResult) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationTeamsTestResult) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NotificationTeamsTestResult) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *NotificationTeamsTestResult) SetStatus(v string) {
	o.Status = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *NotificationTeamsTestResult) GetResponse() string {
	if o == nil || o.Response == nil {
		var ret string
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationTeamsTestResult) GetResponseOk() (*string, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *NotificationTeamsTestResult) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given string and assigns it to the Response field.
func (o *NotificationTeamsTestResult) SetResponse(v string) {
	o.Response = &v
}

func (o NotificationTeamsTestResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationTeamsTestResult struct {
	value *NotificationTeamsTestResult
	isSet bool
}

func (v NullableNotificationTeamsTestResult) Get() *NotificationTeamsTestResult {
	return v.value
}

func (v *NullableNotificationTeamsTestResult) Set(val *NotificationTeamsTestResult) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationTeamsTestResult) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationTeamsTestResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationTeamsTestResult(val *NotificationTeamsTestResult) *NullableNotificationTeamsTestResult {
	return &NullableNotificationTeamsTestResult{value: val, isSet: true}
}

func (v NullableNotificationTeamsTestResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationTeamsTestResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


