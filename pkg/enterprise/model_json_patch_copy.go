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

// checks if the JsonPatchCopy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JsonPatchCopy{}

// JsonPatchCopy The 'copy' operation per RFC6902
type JsonPatchCopy struct {
	// Operation ID, referenced for ordering in the
	Id *string `json:"id,omitempty"`
	Op string `json:"op"`
	// A JSONPointer per RFC6901
	Path string `json:"path" validate:"regexp=^(\\/[^\\/~]*(~[01][^\\/~]*)*)*$"`
	// A JSONPointer per RFC6901
	From string `json:"from" validate:"regexp=^(\\/[^\\/~]*(~[01][^\\/~]*)*)*$"`
}

type _JsonPatchCopy JsonPatchCopy

// NewJsonPatchCopy instantiates a new JsonPatchCopy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonPatchCopy(op string, path string, from string) *JsonPatchCopy {
	this := JsonPatchCopy{}
	this.Op = op
	this.Path = path
	this.From = from
	return &this
}

// NewJsonPatchCopyWithDefaults instantiates a new JsonPatchCopy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonPatchCopyWithDefaults() *JsonPatchCopy {
	this := JsonPatchCopy{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *JsonPatchCopy) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonPatchCopy) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *JsonPatchCopy) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *JsonPatchCopy) SetId(v string) {
	o.Id = &v
}

// GetOp returns the Op field value
func (o *JsonPatchCopy) GetOp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *JsonPatchCopy) GetOpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *JsonPatchCopy) SetOp(v string) {
	o.Op = v
}

// GetPath returns the Path field value
func (o *JsonPatchCopy) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *JsonPatchCopy) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *JsonPatchCopy) SetPath(v string) {
	o.Path = v
}

// GetFrom returns the From field value
func (o *JsonPatchCopy) GetFrom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *JsonPatchCopy) GetFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *JsonPatchCopy) SetFrom(v string) {
	o.From = v
}

func (o JsonPatchCopy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JsonPatchCopy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["op"] = o.Op
	toSerialize["path"] = o.Path
	toSerialize["from"] = o.From
	return toSerialize, nil
}

func (o *JsonPatchCopy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"op",
		"path",
		"from",
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

	varJsonPatchCopy := _JsonPatchCopy{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varJsonPatchCopy)

	if err != nil {
		return err
	}

	*o = JsonPatchCopy(varJsonPatchCopy)

	return err
}

type NullableJsonPatchCopy struct {
	value *JsonPatchCopy
	isSet bool
}

func (v NullableJsonPatchCopy) Get() *JsonPatchCopy {
	return v.value
}

func (v *NullableJsonPatchCopy) Set(val *JsonPatchCopy) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonPatchCopy) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonPatchCopy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonPatchCopy(val *JsonPatchCopy) *NullableJsonPatchCopy {
	return &NullableJsonPatchCopy{value: val, isSet: true}
}

func (v NullableJsonPatchCopy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonPatchCopy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


