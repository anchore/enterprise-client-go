/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.10.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the NotificationWebhookTestResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationWebhookTestResult{}

// NotificationWebhookTestResult struct for NotificationWebhookTestResult
type NotificationWebhookTestResult struct {
	Status *string `json:"status,omitempty"`
	Response *string `json:"response,omitempty"`
}

// NewNotificationWebhookTestResult instantiates a new NotificationWebhookTestResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationWebhookTestResult() *NotificationWebhookTestResult {
	this := NotificationWebhookTestResult{}
	return &this
}

// NewNotificationWebhookTestResultWithDefaults instantiates a new NotificationWebhookTestResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationWebhookTestResultWithDefaults() *NotificationWebhookTestResult {
	this := NotificationWebhookTestResult{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NotificationWebhookTestResult) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationWebhookTestResult) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NotificationWebhookTestResult) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *NotificationWebhookTestResult) SetStatus(v string) {
	o.Status = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *NotificationWebhookTestResult) GetResponse() string {
	if o == nil || IsNil(o.Response) {
		var ret string
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationWebhookTestResult) GetResponseOk() (*string, bool) {
	if o == nil || IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *NotificationWebhookTestResult) HasResponse() bool {
	if o != nil && !IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given string and assigns it to the Response field.
func (o *NotificationWebhookTestResult) SetResponse(v string) {
	o.Response = &v
}

func (o NotificationWebhookTestResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationWebhookTestResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Response) {
		toSerialize["response"] = o.Response
	}
	return toSerialize, nil
}

type NullableNotificationWebhookTestResult struct {
	value *NotificationWebhookTestResult
	isSet bool
}

func (v NullableNotificationWebhookTestResult) Get() *NotificationWebhookTestResult {
	return v.value
}

func (v *NullableNotificationWebhookTestResult) Set(val *NotificationWebhookTestResult) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationWebhookTestResult) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationWebhookTestResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationWebhookTestResult(val *NotificationWebhookTestResult) *NullableNotificationWebhookTestResult {
	return &NullableNotificationWebhookTestResult{value: val, isSet: true}
}

func (v NullableNotificationWebhookTestResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationWebhookTestResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


