/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.3.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// NotificationGitHubTestResult struct for NotificationGitHubTestResult
type NotificationGitHubTestResult struct {
	Status *string `json:"status,omitempty"`
	Response *string `json:"response,omitempty"`
}

// NewNotificationGitHubTestResult instantiates a new NotificationGitHubTestResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationGitHubTestResult() *NotificationGitHubTestResult {
	this := NotificationGitHubTestResult{}
	return &this
}

// NewNotificationGitHubTestResultWithDefaults instantiates a new NotificationGitHubTestResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationGitHubTestResultWithDefaults() *NotificationGitHubTestResult {
	this := NotificationGitHubTestResult{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NotificationGitHubTestResult) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubTestResult) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NotificationGitHubTestResult) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *NotificationGitHubTestResult) SetStatus(v string) {
	o.Status = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *NotificationGitHubTestResult) GetResponse() string {
	if o == nil || o.Response == nil {
		var ret string
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubTestResult) GetResponseOk() (*string, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *NotificationGitHubTestResult) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given string and assigns it to the Response field.
func (o *NotificationGitHubTestResult) SetResponse(v string) {
	o.Response = &v
}

func (o NotificationGitHubTestResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationGitHubTestResult struct {
	value *NotificationGitHubTestResult
	isSet bool
}

func (v NullableNotificationGitHubTestResult) Get() *NotificationGitHubTestResult {
	return v.value
}

func (v *NullableNotificationGitHubTestResult) Set(val *NotificationGitHubTestResult) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationGitHubTestResult) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationGitHubTestResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationGitHubTestResult(val *NotificationGitHubTestResult) *NullableNotificationGitHubTestResult {
	return &NullableNotificationGitHubTestResult{value: val, isSet: true}
}

func (v NullableNotificationGitHubTestResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationGitHubTestResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


