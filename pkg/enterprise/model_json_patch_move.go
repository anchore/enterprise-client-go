/*
Anchore Enterprise API Server

This is the Anchore Enterprise API. It provides additional external API routes and functionality for enterprise users.

API version: 0.8.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// JsonPatchMove The 'move' operation per RFC6902
type JsonPatchMove struct {
	// Operation ID, referenced for ordering in the
	Id *string `json:"id,omitempty"`
	Op string `json:"op"`
	// A JSONPointer per RFC6901
	Path string `json:"path"`
	// A JSONPointer per RFC6901
	From string `json:"from"`
}

// NewJsonPatchMove instantiates a new JsonPatchMove object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonPatchMove(op string, path string, from string) *JsonPatchMove {
	this := JsonPatchMove{}
	this.Op = op
	this.Path = path
	this.From = from
	return &this
}

// NewJsonPatchMoveWithDefaults instantiates a new JsonPatchMove object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonPatchMoveWithDefaults() *JsonPatchMove {
	this := JsonPatchMove{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *JsonPatchMove) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonPatchMove) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *JsonPatchMove) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *JsonPatchMove) SetId(v string) {
	o.Id = &v
}

// GetOp returns the Op field value
func (o *JsonPatchMove) GetOp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *JsonPatchMove) GetOpOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *JsonPatchMove) SetOp(v string) {
	o.Op = v
}

// GetPath returns the Path field value
func (o *JsonPatchMove) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *JsonPatchMove) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *JsonPatchMove) SetPath(v string) {
	o.Path = v
}

// GetFrom returns the From field value
func (o *JsonPatchMove) GetFrom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *JsonPatchMove) GetFromOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *JsonPatchMove) SetFrom(v string) {
	o.From = v
}

func (o JsonPatchMove) MarshalJSON() ([]byte, error) {
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
		toSerialize["from"] = o.From
	}
	return json.Marshal(toSerialize)
}

type NullableJsonPatchMove struct {
	value *JsonPatchMove
	isSet bool
}

func (v NullableJsonPatchMove) Get() *JsonPatchMove {
	return v.value
}

func (v *NullableJsonPatchMove) Set(val *JsonPatchMove) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonPatchMove) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonPatchMove) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonPatchMove(val *JsonPatchMove) *NullableJsonPatchMove {
	return &NullableJsonPatchMove{value: val, isSet: true}
}

func (v NullableJsonPatchMove) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonPatchMove) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


