/*
Anchore Enterprise API Server

This is the Anchore Enterprise API. It provides additional external API routes and functionality for enterprise users.

API version: 0.5.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// JsonPatchRemove The 'remove' operation per RFC6902
type JsonPatchRemove struct {
	// Operation ID, referenced for ordering in the
	Id *string `json:"id,omitempty"`
	Op string `json:"op"`
	// A JSONPointer per RFC6901
	Path string `json:"path"`
}

// NewJsonPatchRemove instantiates a new JsonPatchRemove object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonPatchRemove(op string, path string) *JsonPatchRemove {
	this := JsonPatchRemove{}
	this.Op = op
	this.Path = path
	return &this
}

// NewJsonPatchRemoveWithDefaults instantiates a new JsonPatchRemove object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonPatchRemoveWithDefaults() *JsonPatchRemove {
	this := JsonPatchRemove{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *JsonPatchRemove) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonPatchRemove) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *JsonPatchRemove) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *JsonPatchRemove) SetId(v string) {
	o.Id = &v
}

// GetOp returns the Op field value
func (o *JsonPatchRemove) GetOp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *JsonPatchRemove) GetOpOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *JsonPatchRemove) SetOp(v string) {
	o.Op = v
}

// GetPath returns the Path field value
func (o *JsonPatchRemove) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *JsonPatchRemove) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *JsonPatchRemove) SetPath(v string) {
	o.Path = v
}

func (o JsonPatchRemove) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableJsonPatchRemove struct {
	value *JsonPatchRemove
	isSet bool
}

func (v NullableJsonPatchRemove) Get() *JsonPatchRemove {
	return v.value
}

func (v *NullableJsonPatchRemove) Set(val *JsonPatchRemove) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonPatchRemove) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonPatchRemove) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonPatchRemove(val *JsonPatchRemove) *NullableJsonPatchRemove {
	return &NullableJsonPatchRemove{value: val, isSet: true}
}

func (v NullableJsonPatchRemove) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonPatchRemove) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


