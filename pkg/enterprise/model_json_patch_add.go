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
	"bytes"
	"fmt"
)

// checks if the JsonPatchAdd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JsonPatchAdd{}

// JsonPatchAdd The 'add' operation per RFC6902
type JsonPatchAdd struct {
	// Operation ID, referenced for ordering in the
	Id *string `json:"id,omitempty"`
	Op string `json:"op"`
	// A JSONPointer per RFC6901
	Path string `json:"path" validate:"regexp=^(\\/[^\\/~]*(~[01][^\\/~]*)*)*$"`
	// A valid json value, can be any valid json type
	Value interface{} `json:"value"`
}

type _JsonPatchAdd JsonPatchAdd

// NewJsonPatchAdd instantiates a new JsonPatchAdd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonPatchAdd(op string, path string, value interface{}) *JsonPatchAdd {
	this := JsonPatchAdd{}
	this.Op = op
	this.Path = path
	this.Value = value
	return &this
}

// NewJsonPatchAddWithDefaults instantiates a new JsonPatchAdd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonPatchAddWithDefaults() *JsonPatchAdd {
	this := JsonPatchAdd{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *JsonPatchAdd) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonPatchAdd) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *JsonPatchAdd) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *JsonPatchAdd) SetId(v string) {
	o.Id = &v
}

// GetOp returns the Op field value
func (o *JsonPatchAdd) GetOp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *JsonPatchAdd) GetOpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *JsonPatchAdd) SetOp(v string) {
	o.Op = v
}

// GetPath returns the Path field value
func (o *JsonPatchAdd) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *JsonPatchAdd) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *JsonPatchAdd) SetPath(v string) {
	o.Path = v
}

// GetValue returns the Value field value
func (o *JsonPatchAdd) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *JsonPatchAdd) GetValueOk() (interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value, true
}

// SetValue sets field value
func (o *JsonPatchAdd) SetValue(v interface{}) {
	o.Value = v
}

func (o JsonPatchAdd) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JsonPatchAdd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["op"] = o.Op
	toSerialize["path"] = o.Path
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *JsonPatchAdd) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"op",
		"path",
		"value",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varJsonPatchAdd := _JsonPatchAdd{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varJsonPatchAdd)

	if err != nil {
		return err
	}

	*o = JsonPatchAdd(varJsonPatchAdd)

	return err
}

type NullableJsonPatchAdd struct {
	value *JsonPatchAdd
	isSet bool
}

func (v NullableJsonPatchAdd) Get() *JsonPatchAdd {
	return v.value
}

func (v *NullableJsonPatchAdd) Set(val *JsonPatchAdd) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonPatchAdd) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonPatchAdd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonPatchAdd(val *JsonPatchAdd) *NullableJsonPatchAdd {
	return &NullableJsonPatchAdd{value: val, isSet: true}
}

func (v NullableJsonPatchAdd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonPatchAdd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


