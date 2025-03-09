/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.10.2
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the JsonPatchRemove type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JsonPatchRemove{}

// JsonPatchRemove The 'remove' operation per RFC6902
type JsonPatchRemove struct {
	// Operation ID, referenced for ordering in the
	Id *string `json:"id,omitempty"`
	Op string `json:"op"`
	// A JSONPointer per RFC6901
	Path string `json:"path" validate:"regexp=^(\\/[^\\/~]*(~[01][^\\/~]*)*)*$"`
}

type _JsonPatchRemove JsonPatchRemove

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
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonPatchRemove) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *JsonPatchRemove) HasId() bool {
	if o != nil && !IsNil(o.Id) {
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
	if o == nil {
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
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *JsonPatchRemove) SetPath(v string) {
	o.Path = v
}

func (o JsonPatchRemove) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JsonPatchRemove) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["op"] = o.Op
	toSerialize["path"] = o.Path
	return toSerialize, nil
}

func (o *JsonPatchRemove) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"op",
		"path",
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

	varJsonPatchRemove := _JsonPatchRemove{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varJsonPatchRemove)

	if err != nil {
		return err
	}

	*o = JsonPatchRemove(varJsonPatchRemove)

	return err
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


