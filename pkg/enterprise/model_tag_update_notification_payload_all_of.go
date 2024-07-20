/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.6.1
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// TagUpdateNotificationPayloadAllOf struct for TagUpdateNotificationPayloadAllOf
type TagUpdateNotificationPayloadAllOf struct {
	// A list containing the current image digest
	CurrEval []interface{} `json:"curr_eval,omitempty"`
	// A list containing the previous image digests
	LastEval []interface{} `json:"last_eval,omitempty"`
	// List of Corresponding Image Annotations
	Annotations interface{} `json:"annotations,omitempty"`
}

// NewTagUpdateNotificationPayloadAllOf instantiates a new TagUpdateNotificationPayloadAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagUpdateNotificationPayloadAllOf() *TagUpdateNotificationPayloadAllOf {
	this := TagUpdateNotificationPayloadAllOf{}
	return &this
}

// NewTagUpdateNotificationPayloadAllOfWithDefaults instantiates a new TagUpdateNotificationPayloadAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagUpdateNotificationPayloadAllOfWithDefaults() *TagUpdateNotificationPayloadAllOf {
	this := TagUpdateNotificationPayloadAllOf{}
	return &this
}

// GetCurrEval returns the CurrEval field value if set, zero value otherwise.
func (o *TagUpdateNotificationPayloadAllOf) GetCurrEval() []interface{} {
	if o == nil || o.CurrEval == nil {
		var ret []interface{}
		return ret
	}
	return o.CurrEval
}

// GetCurrEvalOk returns a tuple with the CurrEval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagUpdateNotificationPayloadAllOf) GetCurrEvalOk() ([]interface{}, bool) {
	if o == nil || o.CurrEval == nil {
		return nil, false
	}
	return o.CurrEval, true
}

// HasCurrEval returns a boolean if a field has been set.
func (o *TagUpdateNotificationPayloadAllOf) HasCurrEval() bool {
	if o != nil && o.CurrEval != nil {
		return true
	}

	return false
}

// SetCurrEval gets a reference to the given []interface{} and assigns it to the CurrEval field.
func (o *TagUpdateNotificationPayloadAllOf) SetCurrEval(v []interface{}) {
	o.CurrEval = v
}

// GetLastEval returns the LastEval field value if set, zero value otherwise.
func (o *TagUpdateNotificationPayloadAllOf) GetLastEval() []interface{} {
	if o == nil || o.LastEval == nil {
		var ret []interface{}
		return ret
	}
	return o.LastEval
}

// GetLastEvalOk returns a tuple with the LastEval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagUpdateNotificationPayloadAllOf) GetLastEvalOk() ([]interface{}, bool) {
	if o == nil || o.LastEval == nil {
		return nil, false
	}
	return o.LastEval, true
}

// HasLastEval returns a boolean if a field has been set.
func (o *TagUpdateNotificationPayloadAllOf) HasLastEval() bool {
	if o != nil && o.LastEval != nil {
		return true
	}

	return false
}

// SetLastEval gets a reference to the given []interface{} and assigns it to the LastEval field.
func (o *TagUpdateNotificationPayloadAllOf) SetLastEval(v []interface{}) {
	o.LastEval = v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagUpdateNotificationPayloadAllOf) GetAnnotations() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagUpdateNotificationPayloadAllOf) GetAnnotationsOk() (interface{}, bool) {
	if o == nil || o.Annotations == nil {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *TagUpdateNotificationPayloadAllOf) HasAnnotations() bool {
	if o != nil && o.Annotations != nil {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given interface{} and assigns it to the Annotations field.
func (o *TagUpdateNotificationPayloadAllOf) SetAnnotations(v interface{}) {
	o.Annotations = v
}

func (o TagUpdateNotificationPayloadAllOf) MarshalJSON() ([]byte, error) {
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

type NullableTagUpdateNotificationPayloadAllOf struct {
	value *TagUpdateNotificationPayloadAllOf
	isSet bool
}

func (v NullableTagUpdateNotificationPayloadAllOf) Get() *TagUpdateNotificationPayloadAllOf {
	return v.value
}

func (v *NullableTagUpdateNotificationPayloadAllOf) Set(val *TagUpdateNotificationPayloadAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTagUpdateNotificationPayloadAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTagUpdateNotificationPayloadAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagUpdateNotificationPayloadAllOf(val *TagUpdateNotificationPayloadAllOf) *NullableTagUpdateNotificationPayloadAllOf {
	return &NullableTagUpdateNotificationPayloadAllOf{value: val, isSet: true}
}

func (v NullableTagUpdateNotificationPayloadAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagUpdateNotificationPayloadAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


