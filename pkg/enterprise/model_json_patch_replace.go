/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

<<<<<<< HEAD
API version: 2.0.0
=======
API version: 0.1.0
>>>>>>> 48fc108 (feat: updated the enterprise ref)
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// JsonPatchReplace The 'replace' operation per RFC6902
type JsonPatchReplace struct {
	// Operation ID, referenced for ordering in the
	Id *string `json:"id,omitempty"`
	Op string `json:"op"`
	// A JSONPointer per RFC6901
	Path string `json:"path"`
	// A valid json value, can be any valid json type
	Value interface{} `json:"value"`
}

// NewJsonPatchReplace instantiates a new JsonPatchReplace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonPatchReplace(op string, path string, value interface{}) *JsonPatchReplace {
	this := JsonPatchReplace{}
	this.Op = op
	this.Path = path
	this.Value = value
	return &this
}

// NewJsonPatchReplaceWithDefaults instantiates a new JsonPatchReplace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonPatchReplaceWithDefaults() *JsonPatchReplace {
	this := JsonPatchReplace{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *JsonPatchReplace) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonPatchReplace) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *JsonPatchReplace) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *JsonPatchReplace) SetId(v string) {
	o.Id = &v
}

// GetOp returns the Op field value
func (o *JsonPatchReplace) GetOp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *JsonPatchReplace) GetOpOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *JsonPatchReplace) SetOp(v string) {
	o.Op = v
}

// GetPath returns the Path field value
func (o *JsonPatchReplace) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *JsonPatchReplace) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *JsonPatchReplace) SetPath(v string) {
	o.Path = v
}

// GetValue returns the Value field value
func (o *JsonPatchReplace) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *JsonPatchReplace) GetValueOk() (*interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *JsonPatchReplace) SetValue(v interface{}) {
	o.Value = v
}

func (o JsonPatchReplace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["op"] = o.Op
	}
	if true {
		toSerialize["path"] = o.Path
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableJsonPatchReplace struct {
	value *JsonPatchReplace
	isSet bool
}

func (v NullableJsonPatchReplace) Get() *JsonPatchReplace {
	return v.value
}

func (v *NullableJsonPatchReplace) Set(val *JsonPatchReplace) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonPatchReplace) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonPatchReplace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonPatchReplace(val *JsonPatchReplace) *NullableJsonPatchReplace {
	return &NullableJsonPatchReplace{value: val, isSet: true}
}

func (v NullableJsonPatchReplace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonPatchReplace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


