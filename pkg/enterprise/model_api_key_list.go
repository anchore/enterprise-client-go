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

// ApiKeyList struct for ApiKeyList
type ApiKeyList struct {
	Items []UserApiKey `json:"items,omitempty"`
}

// NewApiKeyList instantiates a new ApiKeyList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiKeyList() *ApiKeyList {
	this := ApiKeyList{}
	return &this
}

// NewApiKeyListWithDefaults instantiates a new ApiKeyList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiKeyListWithDefaults() *ApiKeyList {
	this := ApiKeyList{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ApiKeyList) GetItems() []UserApiKey {
	if o == nil || o.Items == nil {
		var ret []UserApiKey
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyList) GetItemsOk() ([]UserApiKey, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ApiKeyList) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []UserApiKey and assigns it to the Items field.
func (o *ApiKeyList) SetItems(v []UserApiKey) {
	o.Items = v
}

func (o ApiKeyList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableApiKeyList struct {
	value *ApiKeyList
	isSet bool
}

func (v NullableApiKeyList) Get() *ApiKeyList {
	return v.value
}

func (v *NullableApiKeyList) Set(val *ApiKeyList) {
	v.value = val
	v.isSet = true
}

func (v NullableApiKeyList) IsSet() bool {
	return v.isSet
}

func (v *NullableApiKeyList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiKeyList(val *ApiKeyList) *NullableApiKeyList {
	return &NullableApiKeyList{value: val, isSet: true}
}

func (v NullableApiKeyList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiKeyList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


