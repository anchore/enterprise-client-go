/*
Anchore Engine API Server

This is the Anchore Engine API. Provides the primary external API for users of the service.

API version: 0.1.21
Contact: nurmi@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package engine

import (
	"encoding/json"
)

// SubscriptionRequest A subscription entry to add to the system
type SubscriptionRequest struct {
	SubscriptionKey *string `json:"subscription_key,omitempty"`
	SubscriptionValue NullableString `json:"subscription_value,omitempty"`
	SubscriptionType *string `json:"subscription_type,omitempty"`
}

// NewSubscriptionRequest instantiates a new SubscriptionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionRequest() *SubscriptionRequest {
	this := SubscriptionRequest{}
	return &this
}

// NewSubscriptionRequestWithDefaults instantiates a new SubscriptionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionRequestWithDefaults() *SubscriptionRequest {
	this := SubscriptionRequest{}
	return &this
}

// GetSubscriptionKey returns the SubscriptionKey field value if set, zero value otherwise.
func (o *SubscriptionRequest) GetSubscriptionKey() string {
	if o == nil || o.SubscriptionKey == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionKey
}

// GetSubscriptionKeyOk returns a tuple with the SubscriptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionRequest) GetSubscriptionKeyOk() (*string, bool) {
	if o == nil || o.SubscriptionKey == nil {
		return nil, false
	}
	return o.SubscriptionKey, true
}

// HasSubscriptionKey returns a boolean if a field has been set.
func (o *SubscriptionRequest) HasSubscriptionKey() bool {
	if o != nil && o.SubscriptionKey != nil {
		return true
	}

	return false
}

// SetSubscriptionKey gets a reference to the given string and assigns it to the SubscriptionKey field.
func (o *SubscriptionRequest) SetSubscriptionKey(v string) {
	o.SubscriptionKey = &v
}

// GetSubscriptionValue returns the SubscriptionValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionRequest) GetSubscriptionValue() string {
	if o == nil || o.SubscriptionValue.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionValue.Get()
}

// GetSubscriptionValueOk returns a tuple with the SubscriptionValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionRequest) GetSubscriptionValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubscriptionValue.Get(), o.SubscriptionValue.IsSet()
}

// HasSubscriptionValue returns a boolean if a field has been set.
func (o *SubscriptionRequest) HasSubscriptionValue() bool {
	if o != nil && o.SubscriptionValue.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionValue gets a reference to the given NullableString and assigns it to the SubscriptionValue field.
func (o *SubscriptionRequest) SetSubscriptionValue(v string) {
	o.SubscriptionValue.Set(&v)
}
// SetSubscriptionValueNil sets the value for SubscriptionValue to be an explicit nil
func (o *SubscriptionRequest) SetSubscriptionValueNil() {
	o.SubscriptionValue.Set(nil)
}

// UnsetSubscriptionValue ensures that no value is present for SubscriptionValue, not even an explicit nil
func (o *SubscriptionRequest) UnsetSubscriptionValue() {
	o.SubscriptionValue.Unset()
}

// GetSubscriptionType returns the SubscriptionType field value if set, zero value otherwise.
func (o *SubscriptionRequest) GetSubscriptionType() string {
	if o == nil || o.SubscriptionType == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionType
}

// GetSubscriptionTypeOk returns a tuple with the SubscriptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionRequest) GetSubscriptionTypeOk() (*string, bool) {
	if o == nil || o.SubscriptionType == nil {
		return nil, false
	}
	return o.SubscriptionType, true
}

// HasSubscriptionType returns a boolean if a field has been set.
func (o *SubscriptionRequest) HasSubscriptionType() bool {
	if o != nil && o.SubscriptionType != nil {
		return true
	}

	return false
}

// SetSubscriptionType gets a reference to the given string and assigns it to the SubscriptionType field.
func (o *SubscriptionRequest) SetSubscriptionType(v string) {
	o.SubscriptionType = &v
}

func (o SubscriptionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SubscriptionKey != nil {
		toSerialize["subscription_key"] = o.SubscriptionKey
	}
	if o.SubscriptionValue.IsSet() {
		toSerialize["subscription_value"] = o.SubscriptionValue.Get()
	}
	if o.SubscriptionType != nil {
		toSerialize["subscription_type"] = o.SubscriptionType
	}
	return json.Marshal(toSerialize)
}

type NullableSubscriptionRequest struct {
	value *SubscriptionRequest
	isSet bool
}

func (v NullableSubscriptionRequest) Get() *SubscriptionRequest {
	return v.value
}

func (v *NullableSubscriptionRequest) Set(val *SubscriptionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionRequest(val *SubscriptionRequest) *NullableSubscriptionRequest {
	return &NullableSubscriptionRequest{value: val, isSet: true}
}

func (v NullableSubscriptionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


