/*
Anchore Engine API Server

This is the Anchore Engine API. Provides the primary external API for users of the service.

API version: 0.6.0
Contact: nurmi@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package engine

import (
	"encoding/json"
)

// AnalysisUpdateNotificationPayload struct for AnalysisUpdateNotificationPayload
type AnalysisUpdateNotificationPayload struct {
	AccountName *string `json:"account_name,omitempty"`
	SubscriptionKey *string `json:"subscription_key,omitempty"`
	SubscriptionType *string `json:"subscription_type,omitempty"`
	NotificationId *string `json:"notification_id,omitempty"`
	CurrEval *AnalysisUpdateEval `json:"curr_eval,omitempty"`
	LastEval *AnalysisUpdateEval `json:"last_eval,omitempty"`
	// List of Corresponding Image Annotations
	Annotations interface{} `json:"annotations,omitempty"`
}

// NewAnalysisUpdateNotificationPayload instantiates a new AnalysisUpdateNotificationPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalysisUpdateNotificationPayload() *AnalysisUpdateNotificationPayload {
	this := AnalysisUpdateNotificationPayload{}
	return &this
}

// NewAnalysisUpdateNotificationPayloadWithDefaults instantiates a new AnalysisUpdateNotificationPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalysisUpdateNotificationPayloadWithDefaults() *AnalysisUpdateNotificationPayload {
	this := AnalysisUpdateNotificationPayload{}
	return &this
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *AnalysisUpdateNotificationPayload) GetAccountName() string {
	if o == nil || o.AccountName == nil {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisUpdateNotificationPayload) GetAccountNameOk() (*string, bool) {
	if o == nil || o.AccountName == nil {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *AnalysisUpdateNotificationPayload) HasAccountName() bool {
	if o != nil && o.AccountName != nil {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *AnalysisUpdateNotificationPayload) SetAccountName(v string) {
	o.AccountName = &v
}

// GetSubscriptionKey returns the SubscriptionKey field value if set, zero value otherwise.
func (o *AnalysisUpdateNotificationPayload) GetSubscriptionKey() string {
	if o == nil || o.SubscriptionKey == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionKey
}

// GetSubscriptionKeyOk returns a tuple with the SubscriptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisUpdateNotificationPayload) GetSubscriptionKeyOk() (*string, bool) {
	if o == nil || o.SubscriptionKey == nil {
		return nil, false
	}
	return o.SubscriptionKey, true
}

// HasSubscriptionKey returns a boolean if a field has been set.
func (o *AnalysisUpdateNotificationPayload) HasSubscriptionKey() bool {
	if o != nil && o.SubscriptionKey != nil {
		return true
	}

	return false
}

// SetSubscriptionKey gets a reference to the given string and assigns it to the SubscriptionKey field.
func (o *AnalysisUpdateNotificationPayload) SetSubscriptionKey(v string) {
	o.SubscriptionKey = &v
}

// GetSubscriptionType returns the SubscriptionType field value if set, zero value otherwise.
func (o *AnalysisUpdateNotificationPayload) GetSubscriptionType() string {
	if o == nil || o.SubscriptionType == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionType
}

// GetSubscriptionTypeOk returns a tuple with the SubscriptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisUpdateNotificationPayload) GetSubscriptionTypeOk() (*string, bool) {
	if o == nil || o.SubscriptionType == nil {
		return nil, false
	}
	return o.SubscriptionType, true
}

// HasSubscriptionType returns a boolean if a field has been set.
func (o *AnalysisUpdateNotificationPayload) HasSubscriptionType() bool {
	if o != nil && o.SubscriptionType != nil {
		return true
	}

	return false
}

// SetSubscriptionType gets a reference to the given string and assigns it to the SubscriptionType field.
func (o *AnalysisUpdateNotificationPayload) SetSubscriptionType(v string) {
	o.SubscriptionType = &v
}

// GetNotificationId returns the NotificationId field value if set, zero value otherwise.
func (o *AnalysisUpdateNotificationPayload) GetNotificationId() string {
	if o == nil || o.NotificationId == nil {
		var ret string
		return ret
	}
	return *o.NotificationId
}

// GetNotificationIdOk returns a tuple with the NotificationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisUpdateNotificationPayload) GetNotificationIdOk() (*string, bool) {
	if o == nil || o.NotificationId == nil {
		return nil, false
	}
	return o.NotificationId, true
}

// HasNotificationId returns a boolean if a field has been set.
func (o *AnalysisUpdateNotificationPayload) HasNotificationId() bool {
	if o != nil && o.NotificationId != nil {
		return true
	}

	return false
}

// SetNotificationId gets a reference to the given string and assigns it to the NotificationId field.
func (o *AnalysisUpdateNotificationPayload) SetNotificationId(v string) {
	o.NotificationId = &v
}

// GetCurrEval returns the CurrEval field value if set, zero value otherwise.
func (o *AnalysisUpdateNotificationPayload) GetCurrEval() AnalysisUpdateEval {
	if o == nil || o.CurrEval == nil {
		var ret AnalysisUpdateEval
		return ret
	}
	return *o.CurrEval
}

// GetCurrEvalOk returns a tuple with the CurrEval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisUpdateNotificationPayload) GetCurrEvalOk() (*AnalysisUpdateEval, bool) {
	if o == nil || o.CurrEval == nil {
		return nil, false
	}
	return o.CurrEval, true
}

// HasCurrEval returns a boolean if a field has been set.
func (o *AnalysisUpdateNotificationPayload) HasCurrEval() bool {
	if o != nil && o.CurrEval != nil {
		return true
	}

	return false
}

// SetCurrEval gets a reference to the given AnalysisUpdateEval and assigns it to the CurrEval field.
func (o *AnalysisUpdateNotificationPayload) SetCurrEval(v AnalysisUpdateEval) {
	o.CurrEval = &v
}

// GetLastEval returns the LastEval field value if set, zero value otherwise.
func (o *AnalysisUpdateNotificationPayload) GetLastEval() AnalysisUpdateEval {
	if o == nil || o.LastEval == nil {
		var ret AnalysisUpdateEval
		return ret
	}
	return *o.LastEval
}

// GetLastEvalOk returns a tuple with the LastEval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisUpdateNotificationPayload) GetLastEvalOk() (*AnalysisUpdateEval, bool) {
	if o == nil || o.LastEval == nil {
		return nil, false
	}
	return o.LastEval, true
}

// HasLastEval returns a boolean if a field has been set.
func (o *AnalysisUpdateNotificationPayload) HasLastEval() bool {
	if o != nil && o.LastEval != nil {
		return true
	}

	return false
}

// SetLastEval gets a reference to the given AnalysisUpdateEval and assigns it to the LastEval field.
func (o *AnalysisUpdateNotificationPayload) SetLastEval(v AnalysisUpdateEval) {
	o.LastEval = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnalysisUpdateNotificationPayload) GetAnnotations() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnalysisUpdateNotificationPayload) GetAnnotationsOk() (*interface{}, bool) {
	if o == nil || o.Annotations == nil {
		return nil, false
	}
	return &o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *AnalysisUpdateNotificationPayload) HasAnnotations() bool {
	if o != nil && o.Annotations != nil {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given interface{} and assigns it to the Annotations field.
func (o *AnalysisUpdateNotificationPayload) SetAnnotations(v interface{}) {
	o.Annotations = v
}

func (o AnalysisUpdateNotificationPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountName != nil {
		toSerialize["account_name"] = o.AccountName
	}
	if o.SubscriptionKey != nil {
		toSerialize["subscription_key"] = o.SubscriptionKey
	}
	if o.SubscriptionType != nil {
		toSerialize["subscription_type"] = o.SubscriptionType
	}
	if o.NotificationId != nil {
		toSerialize["notification_id"] = o.NotificationId
	}
	if o.CurrEval != nil {
		toSerialize["curr_eval"] = o.CurrEval
	}
	if o.LastEval != nil {
		toSerialize["last_eval"] = o.LastEval
	}
	if o.Annotations != nil {
		toSerialize["annotations"] = o.Annotations
	}
	return json.Marshal(toSerialize)
}

type NullableAnalysisUpdateNotificationPayload struct {
	value *AnalysisUpdateNotificationPayload
	isSet bool
}

func (v NullableAnalysisUpdateNotificationPayload) Get() *AnalysisUpdateNotificationPayload {
	return v.value
}

func (v *NullableAnalysisUpdateNotificationPayload) Set(val *AnalysisUpdateNotificationPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalysisUpdateNotificationPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalysisUpdateNotificationPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalysisUpdateNotificationPayload(val *AnalysisUpdateNotificationPayload) *NullableAnalysisUpdateNotificationPayload {
	return &NullableAnalysisUpdateNotificationPayload{value: val, isSet: true}
}

func (v NullableAnalysisUpdateNotificationPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalysisUpdateNotificationPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


