/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.9.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the TagUpdateNotificationPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagUpdateNotificationPayload{}

// TagUpdateNotificationPayload struct for TagUpdateNotificationPayload
type TagUpdateNotificationPayload struct {
	AccountName *string `json:"account_name,omitempty"`
	SubscriptionKey *string `json:"subscription_key,omitempty"`
	SubscriptionType *string `json:"subscription_type,omitempty"`
	NotificationId *string `json:"notification_id,omitempty"`
	// A list containing the current image digest
	CurrEval []interface{} `json:"curr_eval,omitempty"`
	// A list containing the previous image digests
	LastEval []interface{} `json:"last_eval,omitempty"`
	// List of Corresponding Image Annotations
	Annotations interface{} `json:"annotations,omitempty"`
}

// NewTagUpdateNotificationPayload instantiates a new TagUpdateNotificationPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagUpdateNotificationPayload() *TagUpdateNotificationPayload {
	this := TagUpdateNotificationPayload{}
	return &this
}

// NewTagUpdateNotificationPayloadWithDefaults instantiates a new TagUpdateNotificationPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagUpdateNotificationPayloadWithDefaults() *TagUpdateNotificationPayload {
	this := TagUpdateNotificationPayload{}
	return &this
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *TagUpdateNotificationPayload) GetAccountName() string {
	if o == nil || IsNil(o.AccountName) {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagUpdateNotificationPayload) GetAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.AccountName) {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *TagUpdateNotificationPayload) HasAccountName() bool {
	if o != nil && !IsNil(o.AccountName) {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *TagUpdateNotificationPayload) SetAccountName(v string) {
	o.AccountName = &v
}

// GetSubscriptionKey returns the SubscriptionKey field value if set, zero value otherwise.
func (o *TagUpdateNotificationPayload) GetSubscriptionKey() string {
	if o == nil || IsNil(o.SubscriptionKey) {
		var ret string
		return ret
	}
	return *o.SubscriptionKey
}

// GetSubscriptionKeyOk returns a tuple with the SubscriptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagUpdateNotificationPayload) GetSubscriptionKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionKey) {
		return nil, false
	}
	return o.SubscriptionKey, true
}

// HasSubscriptionKey returns a boolean if a field has been set.
func (o *TagUpdateNotificationPayload) HasSubscriptionKey() bool {
	if o != nil && !IsNil(o.SubscriptionKey) {
		return true
	}

	return false
}

// SetSubscriptionKey gets a reference to the given string and assigns it to the SubscriptionKey field.
func (o *TagUpdateNotificationPayload) SetSubscriptionKey(v string) {
	o.SubscriptionKey = &v
}

// GetSubscriptionType returns the SubscriptionType field value if set, zero value otherwise.
func (o *TagUpdateNotificationPayload) GetSubscriptionType() string {
	if o == nil || IsNil(o.SubscriptionType) {
		var ret string
		return ret
	}
	return *o.SubscriptionType
}

// GetSubscriptionTypeOk returns a tuple with the SubscriptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagUpdateNotificationPayload) GetSubscriptionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionType) {
		return nil, false
	}
	return o.SubscriptionType, true
}

// HasSubscriptionType returns a boolean if a field has been set.
func (o *TagUpdateNotificationPayload) HasSubscriptionType() bool {
	if o != nil && !IsNil(o.SubscriptionType) {
		return true
	}

	return false
}

// SetSubscriptionType gets a reference to the given string and assigns it to the SubscriptionType field.
func (o *TagUpdateNotificationPayload) SetSubscriptionType(v string) {
	o.SubscriptionType = &v
}

// GetNotificationId returns the NotificationId field value if set, zero value otherwise.
func (o *TagUpdateNotificationPayload) GetNotificationId() string {
	if o == nil || IsNil(o.NotificationId) {
		var ret string
		return ret
	}
	return *o.NotificationId
}

// GetNotificationIdOk returns a tuple with the NotificationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagUpdateNotificationPayload) GetNotificationIdOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationId) {
		return nil, false
	}
	return o.NotificationId, true
}

// HasNotificationId returns a boolean if a field has been set.
func (o *TagUpdateNotificationPayload) HasNotificationId() bool {
	if o != nil && !IsNil(o.NotificationId) {
		return true
	}

	return false
}

// SetNotificationId gets a reference to the given string and assigns it to the NotificationId field.
func (o *TagUpdateNotificationPayload) SetNotificationId(v string) {
	o.NotificationId = &v
}

// GetCurrEval returns the CurrEval field value if set, zero value otherwise.
func (o *TagUpdateNotificationPayload) GetCurrEval() []interface{} {
	if o == nil || IsNil(o.CurrEval) {
		var ret []interface{}
		return ret
	}
	return o.CurrEval
}

// GetCurrEvalOk returns a tuple with the CurrEval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagUpdateNotificationPayload) GetCurrEvalOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.CurrEval) {
		return nil, false
	}
	return o.CurrEval, true
}

// HasCurrEval returns a boolean if a field has been set.
func (o *TagUpdateNotificationPayload) HasCurrEval() bool {
	if o != nil && !IsNil(o.CurrEval) {
		return true
	}

	return false
}

// SetCurrEval gets a reference to the given []interface{} and assigns it to the CurrEval field.
func (o *TagUpdateNotificationPayload) SetCurrEval(v []interface{}) {
	o.CurrEval = v
}

// GetLastEval returns the LastEval field value if set, zero value otherwise.
func (o *TagUpdateNotificationPayload) GetLastEval() []interface{} {
	if o == nil || IsNil(o.LastEval) {
		var ret []interface{}
		return ret
	}
	return o.LastEval
}

// GetLastEvalOk returns a tuple with the LastEval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagUpdateNotificationPayload) GetLastEvalOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.LastEval) {
		return nil, false
	}
	return o.LastEval, true
}

// HasLastEval returns a boolean if a field has been set.
func (o *TagUpdateNotificationPayload) HasLastEval() bool {
	if o != nil && !IsNil(o.LastEval) {
		return true
	}

	return false
}

// SetLastEval gets a reference to the given []interface{} and assigns it to the LastEval field.
func (o *TagUpdateNotificationPayload) SetLastEval(v []interface{}) {
	o.LastEval = v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagUpdateNotificationPayload) GetAnnotations() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagUpdateNotificationPayload) GetAnnotationsOk() (interface{}, bool) {
	if o == nil || IsNil(o.Annotations) {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *TagUpdateNotificationPayload) HasAnnotations() bool {
	if o != nil && !IsNil(o.Annotations) {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given interface{} and assigns it to the Annotations field.
func (o *TagUpdateNotificationPayload) SetAnnotations(v interface{}) {
	o.Annotations = v
}

func (o TagUpdateNotificationPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagUpdateNotificationPayload) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.CurrEval) {
		toSerialize["curr_eval"] = o.CurrEval
	}
	if !IsNil(o.LastEval) {
		toSerialize["last_eval"] = o.LastEval
	}
	if o.Annotations != nil {
		toSerialize["annotations"] = o.Annotations
	}
	return toSerialize, nil
}

type NullableTagUpdateNotificationPayload struct {
	value *TagUpdateNotificationPayload
	isSet bool
}

func (v NullableTagUpdateNotificationPayload) Get() *TagUpdateNotificationPayload {
	return v.value
}

func (v *NullableTagUpdateNotificationPayload) Set(val *TagUpdateNotificationPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableTagUpdateNotificationPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableTagUpdateNotificationPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagUpdateNotificationPayload(val *TagUpdateNotificationPayload) *NullableTagUpdateNotificationPayload {
	return &NullableTagUpdateNotificationPayload{value: val, isSet: true}
}

func (v NullableTagUpdateNotificationPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagUpdateNotificationPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


