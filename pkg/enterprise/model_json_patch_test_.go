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

// JsonPatchTest The 'test' operation per RFC6902
type JsonPatchTest struct {
	// Operation ID, referenced for ordering in the
	Id *string `json:"id,omitempty"`
	Op string `json:"op"`
	// A JSONPointer per RFC6901
	Path string `json:"path"`
	// Expected value for test
	Value interface{} `json:"value"`
}

// NewJsonPatchTest instantiates a new JsonPatchTest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonPatchTest(op string, path string, value interface{}) *JsonPatchTest {
	this := JsonPatchTest{}
	this.Op = op
	this.Path = path
	this.Value = value
	return &this
}

// NewJsonPatchTestWithDefaults instantiates a new JsonPatchTest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonPatchTestWithDefaults() *JsonPatchTest {
	this := JsonPatchTest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *JsonPatchTest) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonPatchTest) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *JsonPatchTest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *JsonPatchTest) SetId(v string) {
	o.Id = &v
}

// GetOp returns the Op field value
func (o *JsonPatchTest) GetOp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *JsonPatchTest) GetOpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *JsonPatchTest) SetOp(v string) {
	o.Op = v
}

// GetPath returns the Path field value
func (o *JsonPatchTest) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *JsonPatchTest) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *JsonPatchTest) SetPath(v string) {
	o.Path = v
}

// GetValue returns the Value field value
func (o *JsonPatchTest) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *JsonPatchTest) GetValueOk() (interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value, true
}

// SetValue sets field value
func (o *JsonPatchTest) SetValue(v interface{}) {
	o.Value = v
}

func (o JsonPatchTest) MarshalJSON() ([]byte, error) {
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

type NullableJsonPatchTest struct {
	value *JsonPatchTest
	isSet bool
}

func (v NullableJsonPatchTest) Get() *JsonPatchTest {
	return v.value
}

func (v *NullableJsonPatchTest) Set(val *JsonPatchTest) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonPatchTest) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonPatchTest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonPatchTest(val *JsonPatchTest) *NullableJsonPatchTest {
	return &NullableJsonPatchTest{value: val, isSet: true}
}

func (v NullableJsonPatchTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonPatchTest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


