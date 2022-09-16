/*
Anchore Engine API Server

This is the Anchore Engine API. Provides the primary external API for users of the service.

API version: 0.2.0
Contact: nurmi@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package engine

import (
	"encoding/json"
)

// PolicyEvalNotificationPayloadAllOf struct for PolicyEvalNotificationPayloadAllOf
type PolicyEvalNotificationPayloadAllOf struct {
	// The Current Policy Evaluation result
	CurrEval *interface{} `json:"curr_eval,omitempty"`
	// The Previous Policy Evaluation result
	LastEval *interface{} `json:"last_eval,omitempty"`
	// List of Corresponding Image Annotations
	Annotations interface{} `json:"annotations,omitempty"`
}

// NewPolicyEvalNotificationPayloadAllOf instantiates a new PolicyEvalNotificationPayloadAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyEvalNotificationPayloadAllOf() *PolicyEvalNotificationPayloadAllOf {
	this := PolicyEvalNotificationPayloadAllOf{}
	return &this
}

// NewPolicyEvalNotificationPayloadAllOfWithDefaults instantiates a new PolicyEvalNotificationPayloadAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyEvalNotificationPayloadAllOfWithDefaults() *PolicyEvalNotificationPayloadAllOf {
	this := PolicyEvalNotificationPayloadAllOf{}
	return &this
}

// GetCurrEval returns the CurrEval field value if set, zero value otherwise.
func (o *PolicyEvalNotificationPayloadAllOf) GetCurrEval() interface{} {
	if o == nil || o.CurrEval == nil {
		var ret interface{}
		return ret
	}
	return *o.CurrEval
}

// GetCurrEvalOk returns a tuple with the CurrEval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvalNotificationPayloadAllOf) GetCurrEvalOk() (*interface{}, bool) {
	if o == nil || o.CurrEval == nil {
		return nil, false
	}
	return o.CurrEval, true
}

// HasCurrEval returns a boolean if a field has been set.
func (o *PolicyEvalNotificationPayloadAllOf) HasCurrEval() bool {
	if o != nil && o.CurrEval != nil {
		return true
	}

	return false
}

// SetCurrEval gets a reference to the given interface{} and assigns it to the CurrEval field.
func (o *PolicyEvalNotificationPayloadAllOf) SetCurrEval(v interface{}) {
	o.CurrEval = &v
}

// GetLastEval returns the LastEval field value if set, zero value otherwise.
func (o *PolicyEvalNotificationPayloadAllOf) GetLastEval() interface{} {
	if o == nil || o.LastEval == nil {
		var ret interface{}
		return ret
	}
	return *o.LastEval
}

// GetLastEvalOk returns a tuple with the LastEval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvalNotificationPayloadAllOf) GetLastEvalOk() (*interface{}, bool) {
	if o == nil || o.LastEval == nil {
		return nil, false
	}
	return o.LastEval, true
}

// HasLastEval returns a boolean if a field has been set.
func (o *PolicyEvalNotificationPayloadAllOf) HasLastEval() bool {
	if o != nil && o.LastEval != nil {
		return true
	}

	return false
}

// SetLastEval gets a reference to the given interface{} and assigns it to the LastEval field.
func (o *PolicyEvalNotificationPayloadAllOf) SetLastEval(v interface{}) {
	o.LastEval = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyEvalNotificationPayloadAllOf) GetAnnotations() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyEvalNotificationPayloadAllOf) GetAnnotationsOk() (*interface{}, bool) {
	if o == nil || o.Annotations == nil {
		return nil, false
	}
	return &o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *PolicyEvalNotificationPayloadAllOf) HasAnnotations() bool {
	if o != nil && o.Annotations != nil {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given interface{} and assigns it to the Annotations field.
func (o *PolicyEvalNotificationPayloadAllOf) SetAnnotations(v interface{}) {
	o.Annotations = v
}

func (o PolicyEvalNotificationPayloadAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullablePolicyEvalNotificationPayloadAllOf struct {
	value *PolicyEvalNotificationPayloadAllOf
	isSet bool
}

func (v NullablePolicyEvalNotificationPayloadAllOf) Get() *PolicyEvalNotificationPayloadAllOf {
	return v.value
}

func (v *NullablePolicyEvalNotificationPayloadAllOf) Set(val *PolicyEvalNotificationPayloadAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyEvalNotificationPayloadAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyEvalNotificationPayloadAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyEvalNotificationPayloadAllOf(val *PolicyEvalNotificationPayloadAllOf) *NullablePolicyEvalNotificationPayloadAllOf {
	return &NullablePolicyEvalNotificationPayloadAllOf{value: val, isSet: true}
}

func (v NullablePolicyEvalNotificationPayloadAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyEvalNotificationPayloadAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


