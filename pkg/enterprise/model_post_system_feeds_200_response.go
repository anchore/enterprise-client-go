/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.10.1
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the PostSystemFeeds200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostSystemFeeds200Response{}

// PostSystemFeeds200Response struct for PostSystemFeeds200Response
type PostSystemFeeds200Response struct {
	// Message from the operation
	Message *string `json:"message,omitempty"`
}

// NewPostSystemFeeds200Response instantiates a new PostSystemFeeds200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostSystemFeeds200Response() *PostSystemFeeds200Response {
	this := PostSystemFeeds200Response{}
	return &this
}

// NewPostSystemFeeds200ResponseWithDefaults instantiates a new PostSystemFeeds200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostSystemFeeds200ResponseWithDefaults() *PostSystemFeeds200Response {
	this := PostSystemFeeds200Response{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *PostSystemFeeds200Response) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSystemFeeds200Response) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *PostSystemFeeds200Response) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *PostSystemFeeds200Response) SetMessage(v string) {
	o.Message = &v
}

func (o PostSystemFeeds200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostSystemFeeds200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullablePostSystemFeeds200Response struct {
	value *PostSystemFeeds200Response
	isSet bool
}

func (v NullablePostSystemFeeds200Response) Get() *PostSystemFeeds200Response {
	return v.value
}

func (v *NullablePostSystemFeeds200Response) Set(val *PostSystemFeeds200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePostSystemFeeds200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePostSystemFeeds200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostSystemFeeds200Response(val *PostSystemFeeds200Response) *NullablePostSystemFeeds200Response {
	return &NullablePostSystemFeeds200Response{value: val, isSet: true}
}

func (v NullablePostSystemFeeds200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostSystemFeeds200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


