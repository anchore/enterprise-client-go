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

// checks if the GenericNotificationPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenericNotificationPayload{}

// GenericNotificationPayload Parent class for Notification Payloads
type GenericNotificationPayload struct {
	AccountName *string `json:"account_name,omitempty"`
	SubscriptionKey *string `json:"subscription_key,omitempty"`
	SubscriptionType *string `json:"subscription_type,omitempty"`
	NotificationId *string `json:"notification_id,omitempty"`
}

// NewGenericNotificationPayload instantiates a new GenericNotificationPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericNotificationPayload() *GenericNotificationPayload {
	this := GenericNotificationPayload{}
	return &this
}

// NewGenericNotificationPayloadWithDefaults instantiates a new GenericNotificationPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericNotificationPayloadWithDefaults() *GenericNotificationPayload {
	this := GenericNotificationPayload{}
	return &this
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *GenericNotificationPayload) GetAccountName() string {
	if o == nil || IsNil(o.AccountName) {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericNotificationPayload) GetAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.AccountName) {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *GenericNotificationPayload) HasAccountName() bool {
	if o != nil && !IsNil(o.AccountName) {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *GenericNotificationPayload) SetAccountName(v string) {
	o.AccountName = &v
}

// GetSubscriptionKey returns the SubscriptionKey field value if set, zero value otherwise.
func (o *GenericNotificationPayload) GetSubscriptionKey() string {
	if o == nil || IsNil(o.SubscriptionKey) {
		var ret string
		return ret
	}
	return *o.SubscriptionKey
}

// GetSubscriptionKeyOk returns a tuple with the SubscriptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericNotificationPayload) GetSubscriptionKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionKey) {
		return nil, false
	}
	return o.SubscriptionKey, true
}

// HasSubscriptionKey returns a boolean if a field has been set.
func (o *GenericNotificationPayload) HasSubscriptionKey() bool {
	if o != nil && !IsNil(o.SubscriptionKey) {
		return true
	}

	return false
}

// SetSubscriptionKey gets a reference to the given string and assigns it to the SubscriptionKey field.
func (o *GenericNotificationPayload) SetSubscriptionKey(v string) {
	o.SubscriptionKey = &v
}

// GetSubscriptionType returns the SubscriptionType field value if set, zero value otherwise.
func (o *GenericNotificationPayload) GetSubscriptionType() string {
	if o == nil || IsNil(o.SubscriptionType) {
		var ret string
		return ret
	}
	return *o.SubscriptionType
}

// GetSubscriptionTypeOk returns a tuple with the SubscriptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericNotificationPayload) GetSubscriptionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionType) {
		return nil, false
	}
	return o.SubscriptionType, true
}

// HasSubscriptionType returns a boolean if a field has been set.
func (o *GenericNotificationPayload) HasSubscriptionType() bool {
	if o != nil && !IsNil(o.SubscriptionType) {
		return true
	}

	return false
}

// SetSubscriptionType gets a reference to the given string and assigns it to the SubscriptionType field.
func (o *GenericNotificationPayload) SetSubscriptionType(v string) {
	o.SubscriptionType = &v
}

// GetNotificationId returns the NotificationId field value if set, zero value otherwise.
func (o *GenericNotificationPayload) GetNotificationId() string {
	if o == nil || IsNil(o.NotificationId) {
		var ret string
		return ret
	}
	return *o.NotificationId
}

// GetNotificationIdOk returns a tuple with the NotificationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericNotificationPayload) GetNotificationIdOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationId) {
		return nil, false
	}
	return o.NotificationId, true
}

// HasNotificationId returns a boolean if a field has been set.
func (o *GenericNotificationPayload) HasNotificationId() bool {
	if o != nil && !IsNil(o.NotificationId) {
		return true
	}

	return false
}

// SetNotificationId gets a reference to the given string and assigns it to the NotificationId field.
func (o *GenericNotificationPayload) SetNotificationId(v string) {
	o.NotificationId = &v
}

func (o GenericNotificationPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenericNotificationPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountName) {
		toSerialize["account_name"] = o.AccountName
	}
	if !IsNil(o.SubscriptionKey) {
		toSerialize["subscription_key"] = o.SubscriptionKey
	}
	if !IsNil(o.SubscriptionType) {
		toSerialize["subscription_type"] = o.SubscriptionType
	}
	if !IsNil(o.NotificationId) {
		toSerialize["notification_id"] = o.NotificationId
	}
	return toSerialize, nil
}

type NullableGenericNotificationPayload struct {
	value *GenericNotificationPayload
	isSet bool
}

func (v NullableGenericNotificationPayload) Get() *GenericNotificationPayload {
	return v.value
}

func (v *NullableGenericNotificationPayload) Set(val *GenericNotificationPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericNotificationPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericNotificationPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericNotificationPayload(val *GenericNotificationPayload) *NullableGenericNotificationPayload {
	return &NullableGenericNotificationPayload{value: val, isSet: true}
}

func (v NullableGenericNotificationPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericNotificationPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


