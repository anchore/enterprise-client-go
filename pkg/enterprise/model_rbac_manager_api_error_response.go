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

// RbacManagerApiErrorResponse Generic HTTP API error response
type RbacManagerApiErrorResponse struct {
	Code *int32 `json:"code,omitempty"`
	ErrorType *string `json:"error_type,omitempty"`
	Message *string `json:"message,omitempty"`
	// Details structure for additional information about the error if available. Content and structure will be error specific.
	Detail interface{} `json:"detail,omitempty"`
}

// NewRbacManagerApiErrorResponse instantiates a new RbacManagerApiErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRbacManagerApiErrorResponse() *RbacManagerApiErrorResponse {
	this := RbacManagerApiErrorResponse{}
	return &this
}

// NewRbacManagerApiErrorResponseWithDefaults instantiates a new RbacManagerApiErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRbacManagerApiErrorResponseWithDefaults() *RbacManagerApiErrorResponse {
	this := RbacManagerApiErrorResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *RbacManagerApiErrorResponse) GetCode() int32 {
	if o == nil || o.Code == nil {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerApiErrorResponse) GetCodeOk() (*int32, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *RbacManagerApiErrorResponse) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *RbacManagerApiErrorResponse) SetCode(v int32) {
	o.Code = &v
}

// GetErrorType returns the ErrorType field value if set, zero value otherwise.
func (o *RbacManagerApiErrorResponse) GetErrorType() string {
	if o == nil || o.ErrorType == nil {
		var ret string
		return ret
	}
	return *o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerApiErrorResponse) GetErrorTypeOk() (*string, bool) {
	if o == nil || o.ErrorType == nil {
		return nil, false
	}
	return o.ErrorType, true
}

// HasErrorType returns a boolean if a field has been set.
func (o *RbacManagerApiErrorResponse) HasErrorType() bool {
	if o != nil && o.ErrorType != nil {
		return true
	}

	return false
}

// SetErrorType gets a reference to the given string and assigns it to the ErrorType field.
func (o *RbacManagerApiErrorResponse) SetErrorType(v string) {
	o.ErrorType = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RbacManagerApiErrorResponse) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerApiErrorResponse) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *RbacManagerApiErrorResponse) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RbacManagerApiErrorResponse) SetMessage(v string) {
	o.Message = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *RbacManagerApiErrorResponse) GetDetail() interface{} {
	if o == nil || o.Detail == nil {
		var ret interface{}
		return ret
	}
	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerApiErrorResponse) GetDetailOk() (interface{}, bool) {
	if o == nil || o.Detail == nil {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *RbacManagerApiErrorResponse) HasDetail() bool {
	if o != nil && o.Detail != nil {
		return true
	}

	return false
}

// SetDetail gets a reference to the given interface{} and assigns it to the Detail field.
func (o *RbacManagerApiErrorResponse) SetDetail(v interface{}) {
	o.Detail = v
}

func (o RbacManagerApiErrorResponse) MarshalJSON() ([]byte, error) {
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

type NullableRbacManagerApiErrorResponse struct {
	value *RbacManagerApiErrorResponse
	isSet bool
}

func (v NullableRbacManagerApiErrorResponse) Get() *RbacManagerApiErrorResponse {
	return v.value
}

func (v *NullableRbacManagerApiErrorResponse) Set(val *RbacManagerApiErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRbacManagerApiErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRbacManagerApiErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRbacManagerApiErrorResponse(val *RbacManagerApiErrorResponse) *NullableRbacManagerApiErrorResponse {
	return &NullableRbacManagerApiErrorResponse{value: val, isSet: true}
}

func (v NullableRbacManagerApiErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRbacManagerApiErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


