/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.10.2
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the NotificationSlackTestResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationSlackTestResult{}

// NotificationSlackTestResult struct for NotificationSlackTestResult
type NotificationSlackTestResult struct {
	Status *string `json:"status,omitempty"`
	Response *string `json:"response,omitempty"`
}

// NewNotificationSlackTestResult instantiates a new NotificationSlackTestResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationSlackTestResult() *NotificationSlackTestResult {
	this := NotificationSlackTestResult{}
	return &this
}

// NewNotificationSlackTestResultWithDefaults instantiates a new NotificationSlackTestResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationSlackTestResultWithDefaults() *NotificationSlackTestResult {
	this := NotificationSlackTestResult{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NotificationSlackTestResult) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSlackTestResult) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NotificationSlackTestResult) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *NotificationSlackTestResult) SetStatus(v string) {
	o.Status = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *NotificationSlackTestResult) GetResponse() string {
	if o == nil || IsNil(o.Response) {
		var ret string
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSlackTestResult) GetResponseOk() (*string, bool) {
	if o == nil || IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *NotificationSlackTestResult) HasResponse() bool {
	if o != nil && !IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given string and assigns it to the Response field.
func (o *NotificationSlackTestResult) SetResponse(v string) {
	o.Response = &v
}

func (o NotificationSlackTestResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationSlackTestResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Response) {
		toSerialize["response"] = o.Response
	}
	return toSerialize, nil
}

type NullableNotificationSlackTestResult struct {
	value *NotificationSlackTestResult
	isSet bool
}

func (v NullableNotificationSlackTestResult) Get() *NotificationSlackTestResult {
	return v.value
}

func (v *NullableNotificationSlackTestResult) Set(val *NotificationSlackTestResult) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationSlackTestResult) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationSlackTestResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationSlackTestResult(val *NotificationSlackTestResult) *NullableNotificationSlackTestResult {
	return &NullableNotificationSlackTestResult{value: val, isSet: true}
}

func (v NullableNotificationSlackTestResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationSlackTestResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


