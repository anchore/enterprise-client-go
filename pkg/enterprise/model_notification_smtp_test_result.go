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

// NotificationSMTPTestResult struct for NotificationSMTPTestResult
type NotificationSMTPTestResult struct {
	Status *string `json:"status,omitempty"`
	Response *string `json:"response,omitempty"`
}

// NewNotificationSMTPTestResult instantiates a new NotificationSMTPTestResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationSMTPTestResult() *NotificationSMTPTestResult {
	this := NotificationSMTPTestResult{}
	return &this
}

// NewNotificationSMTPTestResultWithDefaults instantiates a new NotificationSMTPTestResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationSMTPTestResultWithDefaults() *NotificationSMTPTestResult {
	this := NotificationSMTPTestResult{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NotificationSMTPTestResult) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSMTPTestResult) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NotificationSMTPTestResult) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *NotificationSMTPTestResult) SetStatus(v string) {
	o.Status = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *NotificationSMTPTestResult) GetResponse() string {
	if o == nil || o.Response == nil {
		var ret string
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSMTPTestResult) GetResponseOk() (*string, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *NotificationSMTPTestResult) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given string and assigns it to the Response field.
func (o *NotificationSMTPTestResult) SetResponse(v string) {
	o.Response = &v
}

func (o NotificationSMTPTestResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationSMTPTestResult struct {
	value *NotificationSMTPTestResult
	isSet bool
}

func (v NullableNotificationSMTPTestResult) Get() *NotificationSMTPTestResult {
	return v.value
}

func (v *NullableNotificationSMTPTestResult) Set(val *NotificationSMTPTestResult) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationSMTPTestResult) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationSMTPTestResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationSMTPTestResult(val *NotificationSMTPTestResult) *NullableNotificationSMTPTestResult {
	return &NullableNotificationSMTPTestResult{value: val, isSet: true}
}

func (v NullableNotificationSMTPTestResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationSMTPTestResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


