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

// Subscription Subscription entry
type Subscription struct {
	// The key value that the subscription references. E.g. a tag value or a repo name.
	SubscriptionKey *string `json:"subscription_key,omitempty"`
	// The type of the subscription
	SubscriptionType *string `json:"subscription_type,omitempty"`
	// The value of the subscription target
	SubscriptionValue NullableString `json:"subscription_value,omitempty"`
	// The account_name of the subscribed user
	AccountName *string `json:"account_name,omitempty"`
	// Is the subscription currently active
	Active *bool `json:"active,omitempty"`
	// the unique id for this subscription record
	SubscriptionId *string `json:"subscription_id,omitempty"`
}

// NewSubscription instantiates a new Subscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscription() *Subscription {
	this := Subscription{}
	return &this
}

// NewSubscriptionWithDefaults instantiates a new Subscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionWithDefaults() *Subscription {
	this := Subscription{}
	return &this
}

// GetSubscriptionKey returns the SubscriptionKey field value if set, zero value otherwise.
func (o *Subscription) GetSubscriptionKey() string {
	if o == nil || o.SubscriptionKey == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionKey
}

// GetSubscriptionKeyOk returns a tuple with the SubscriptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetSubscriptionKeyOk() (*string, bool) {
	if o == nil || o.SubscriptionKey == nil {
		return nil, false
	}
	return o.SubscriptionKey, true
}

// HasSubscriptionKey returns a boolean if a field has been set.
func (o *Subscription) HasSubscriptionKey() bool {
	if o != nil && o.SubscriptionKey != nil {
		return true
	}

	return false
}

// SetSubscriptionKey gets a reference to the given string and assigns it to the SubscriptionKey field.
func (o *Subscription) SetSubscriptionKey(v string) {
	o.SubscriptionKey = &v
}

// GetSubscriptionType returns the SubscriptionType field value if set, zero value otherwise.
func (o *Subscription) GetSubscriptionType() string {
	if o == nil || o.SubscriptionType == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionType
}

// GetSubscriptionTypeOk returns a tuple with the SubscriptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetSubscriptionTypeOk() (*string, bool) {
	if o == nil || o.SubscriptionType == nil {
		return nil, false
	}
	return o.SubscriptionType, true
}

// HasSubscriptionType returns a boolean if a field has been set.
func (o *Subscription) HasSubscriptionType() bool {
	if o != nil && o.SubscriptionType != nil {
		return true
	}

	return false
}

// SetSubscriptionType gets a reference to the given string and assigns it to the SubscriptionType field.
func (o *Subscription) SetSubscriptionType(v string) {
	o.SubscriptionType = &v
}

// GetSubscriptionValue returns the SubscriptionValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Subscription) GetSubscriptionValue() string {
	if o == nil || o.SubscriptionValue.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionValue.Get()
}

// GetSubscriptionValueOk returns a tuple with the SubscriptionValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Subscription) GetSubscriptionValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionValue.Get(), o.SubscriptionValue.IsSet()
}

// HasSubscriptionValue returns a boolean if a field has been set.
func (o *Subscription) HasSubscriptionValue() bool {
	if o != nil && o.SubscriptionValue.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionValue gets a reference to the given NullableString and assigns it to the SubscriptionValue field.
func (o *Subscription) SetSubscriptionValue(v string) {
	o.SubscriptionValue.Set(&v)
}
// SetSubscriptionValueNil sets the value for SubscriptionValue to be an explicit nil
func (o *Subscription) SetSubscriptionValueNil() {
	o.SubscriptionValue.Set(nil)
}

// UnsetSubscriptionValue ensures that no value is present for SubscriptionValue, not even an explicit nil
func (o *Subscription) UnsetSubscriptionValue() {
	o.SubscriptionValue.Unset()
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *Subscription) GetAccountName() string {
	if o == nil || o.AccountName == nil {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetAccountNameOk() (*string, bool) {
	if o == nil || o.AccountName == nil {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *Subscription) HasAccountName() bool {
	if o != nil && o.AccountName != nil {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *Subscription) SetAccountName(v string) {
	o.AccountName = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *Subscription) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *Subscription) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *Subscription) SetActive(v bool) {
	o.Active = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *Subscription) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || o.SubscriptionId == nil {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *Subscription) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId != nil {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *Subscription) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

func (o Subscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SubscriptionKey != nil {
		toSerialize["subscription_key"] = o.SubscriptionKey
	}
	if o.SubscriptionType != nil {
		toSerialize["subscription_type"] = o.SubscriptionType
	}
	if o.SubscriptionValue.IsSet() {
		toSerialize["subscription_value"] = o.SubscriptionValue.Get()
	}
	if o.AccountName != nil {
		toSerialize["account_name"] = o.AccountName
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.SubscriptionId != nil {
		toSerialize["subscription_id"] = o.SubscriptionId
	}
	return json.Marshal(toSerialize)
}

type NullableSubscription struct {
	value *Subscription
	isSet bool
}

func (v NullableSubscription) Get() *Subscription {
	return v.value
}

func (v *NullableSubscription) Set(val *Subscription) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscription(val *Subscription) *NullableSubscription {
	return &NullableSubscription{value: val, isSet: true}
}

func (v NullableSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


