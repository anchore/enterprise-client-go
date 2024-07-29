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

// AnalysisUpdateNotificationPayloadAllOf struct for AnalysisUpdateNotificationPayloadAllOf
type AnalysisUpdateNotificationPayloadAllOf struct {
	CurrEval *AnalysisUpdateEval `json:"curr_eval,omitempty"`
	LastEval *AnalysisUpdateEval `json:"last_eval,omitempty"`
	// List of Corresponding Image Annotations
	Annotations interface{} `json:"annotations,omitempty"`
}

// NewAnalysisUpdateNotificationPayloadAllOf instantiates a new AnalysisUpdateNotificationPayloadAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalysisUpdateNotificationPayloadAllOf() *AnalysisUpdateNotificationPayloadAllOf {
	this := AnalysisUpdateNotificationPayloadAllOf{}
	return &this
}

// NewAnalysisUpdateNotificationPayloadAllOfWithDefaults instantiates a new AnalysisUpdateNotificationPayloadAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalysisUpdateNotificationPayloadAllOfWithDefaults() *AnalysisUpdateNotificationPayloadAllOf {
	this := AnalysisUpdateNotificationPayloadAllOf{}
	return &this
}

// GetCurrEval returns the CurrEval field value if set, zero value otherwise.
func (o *AnalysisUpdateNotificationPayloadAllOf) GetCurrEval() AnalysisUpdateEval {
	if o == nil || o.CurrEval == nil {
		var ret AnalysisUpdateEval
		return ret
	}
	return *o.CurrEval
}

// GetCurrEvalOk returns a tuple with the CurrEval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisUpdateNotificationPayloadAllOf) GetCurrEvalOk() (*AnalysisUpdateEval, bool) {
	if o == nil || o.CurrEval == nil {
		return nil, false
	}
	return o.CurrEval, true
}

// HasCurrEval returns a boolean if a field has been set.
func (o *AnalysisUpdateNotificationPayloadAllOf) HasCurrEval() bool {
	if o != nil && o.CurrEval != nil {
		return true
	}

	return false
}

// SetCurrEval gets a reference to the given AnalysisUpdateEval and assigns it to the CurrEval field.
func (o *AnalysisUpdateNotificationPayloadAllOf) SetCurrEval(v AnalysisUpdateEval) {
	o.CurrEval = &v
}

// GetLastEval returns the LastEval field value if set, zero value otherwise.
func (o *AnalysisUpdateNotificationPayloadAllOf) GetLastEval() AnalysisUpdateEval {
	if o == nil || o.LastEval == nil {
		var ret AnalysisUpdateEval
		return ret
	}
	return *o.LastEval
}

// GetLastEvalOk returns a tuple with the LastEval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisUpdateNotificationPayloadAllOf) GetLastEvalOk() (*AnalysisUpdateEval, bool) {
	if o == nil || o.LastEval == nil {
		return nil, false
	}
	return o.LastEval, true
}

// HasLastEval returns a boolean if a field has been set.
func (o *AnalysisUpdateNotificationPayloadAllOf) HasLastEval() bool {
	if o != nil && o.LastEval != nil {
		return true
	}

	return false
}

// SetLastEval gets a reference to the given AnalysisUpdateEval and assigns it to the LastEval field.
func (o *AnalysisUpdateNotificationPayloadAllOf) SetLastEval(v AnalysisUpdateEval) {
	o.LastEval = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnalysisUpdateNotificationPayloadAllOf) GetAnnotations() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnalysisUpdateNotificationPayloadAllOf) GetAnnotationsOk() (interface{}, bool) {
	if o == nil || o.Annotations == nil {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *AnalysisUpdateNotificationPayloadAllOf) HasAnnotations() bool {
	if o != nil && o.Annotations != nil {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given interface{} and assigns it to the Annotations field.
func (o *AnalysisUpdateNotificationPayloadAllOf) SetAnnotations(v interface{}) {
	o.Annotations = v
}

func (o AnalysisUpdateNotificationPayloadAllOf) MarshalJSON() ([]byte, error) {
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

type NullableAnalysisUpdateNotificationPayloadAllOf struct {
	value *AnalysisUpdateNotificationPayloadAllOf
	isSet bool
}

func (v NullableAnalysisUpdateNotificationPayloadAllOf) Get() *AnalysisUpdateNotificationPayloadAllOf {
	return v.value
}

func (v *NullableAnalysisUpdateNotificationPayloadAllOf) Set(val *AnalysisUpdateNotificationPayloadAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalysisUpdateNotificationPayloadAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalysisUpdateNotificationPayloadAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalysisUpdateNotificationPayloadAllOf(val *AnalysisUpdateNotificationPayloadAllOf) *NullableAnalysisUpdateNotificationPayloadAllOf {
	return &NullableAnalysisUpdateNotificationPayloadAllOf{value: val, isSet: true}
}

func (v NullableAnalysisUpdateNotificationPayloadAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalysisUpdateNotificationPayloadAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


