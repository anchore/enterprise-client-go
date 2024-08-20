/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.6.2
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// NotificationApiErrorResponse Generic HTTP API error response
type NotificationApiErrorResponse struct {
	Code *int32 `json:"code,omitempty"`
	ErrorType *string `json:"error_type,omitempty"`
	Message *string `json:"message,omitempty"`
	// Details structure for additional information about the error if available. Content and structure will be error specific.
	Detail interface{} `json:"detail,omitempty"`
}

// NewNotificationApiErrorResponse instantiates a new NotificationApiErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationApiErrorResponse() *NotificationApiErrorResponse {
	this := NotificationApiErrorResponse{}
	return &this
}

// NewNotificationApiErrorResponseWithDefaults instantiates a new NotificationApiErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationApiErrorResponseWithDefaults() *NotificationApiErrorResponse {
	this := NotificationApiErrorResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *NotificationApiErrorResponse) GetCode() int32 {
	if o == nil || o.Code == nil {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationApiErrorResponse) GetCodeOk() (*int32, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *NotificationApiErrorResponse) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *NotificationApiErrorResponse) SetCode(v int32) {
	o.Code = &v
}

// GetErrorType returns the ErrorType field value if set, zero value otherwise.
func (o *NotificationApiErrorResponse) GetErrorType() string {
	if o == nil || o.ErrorType == nil {
		var ret string
		return ret
	}
	return *o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationApiErrorResponse) GetErrorTypeOk() (*string, bool) {
	if o == nil || o.ErrorType == nil {
		return nil, false
	}
	return o.ErrorType, true
}

// HasErrorType returns a boolean if a field has been set.
func (o *NotificationApiErrorResponse) HasErrorType() bool {
	if o != nil && o.ErrorType != nil {
		return true
	}

	return false
}

// SetErrorType gets a reference to the given string and assigns it to the ErrorType field.
func (o *NotificationApiErrorResponse) SetErrorType(v string) {
	o.ErrorType = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *NotificationApiErrorResponse) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationApiErrorResponse) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *NotificationApiErrorResponse) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *NotificationApiErrorResponse) SetMessage(v string) {
	o.Message = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *NotificationApiErrorResponse) GetDetail() interface{} {
	if o == nil || o.Detail == nil {
		var ret interface{}
		return ret
	}
	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationApiErrorResponse) GetDetailOk() (interface{}, bool) {
	if o == nil || o.Detail == nil {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *NotificationApiErrorResponse) HasDetail() bool {
	if o != nil && o.Detail != nil {
		return true
	}

	return false
}

// SetDetail gets a reference to the given interface{} and assigns it to the Detail field.
func (o *NotificationApiErrorResponse) SetDetail(v interface{}) {
	o.Detail = v
}

func (o NotificationApiErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.ErrorType != nil {
		toSerialize["error_type"] = o.ErrorType
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Detail != nil {
		toSerialize["detail"] = o.Detail
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationApiErrorResponse struct {
	value *NotificationApiErrorResponse
	isSet bool
}

func (v NullableNotificationApiErrorResponse) Get() *NotificationApiErrorResponse {
	return v.value
}

func (v *NullableNotificationApiErrorResponse) Set(val *NotificationApiErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationApiErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationApiErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationApiErrorResponse(val *NotificationApiErrorResponse) *NullableNotificationApiErrorResponse {
	return &NullableNotificationApiErrorResponse{value: val, isSet: true}
}

func (v NullableNotificationApiErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationApiErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


