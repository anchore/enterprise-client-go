/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.10.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the ECSTasks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ECSTasks{}

// ECSTasks Tasks defined in ECS
type ECSTasks struct {
	Tasks []ECSTask `json:"tasks,omitempty"`
}

// NewECSTasks instantiates a new ECSTasks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewECSTasks() *ECSTasks {
	this := ECSTasks{}
	return &this
}

// NewECSTasksWithDefaults instantiates a new ECSTasks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewECSTasksWithDefaults() *ECSTasks {
	this := ECSTasks{}
	return &this
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *ECSTasks) GetTasks() []ECSTask {
	if o == nil || IsNil(o.Tasks) {
		var ret []ECSTask
		return ret
	}
	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSTasks) GetTasksOk() ([]ECSTask, bool) {
	if o == nil || IsNil(o.Tasks) {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *ECSTasks) HasTasks() bool {
	if o != nil && !IsNil(o.Tasks) {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []ECSTask and assigns it to the Tasks field.
func (o *ECSTasks) SetTasks(v []ECSTask) {
	o.Tasks = v
}

func (o ECSTasks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ECSTasks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tasks) {
		toSerialize["tasks"] = o.Tasks
	}
	return toSerialize, nil
}

type NullableECSTasks struct {
	value *ECSTasks
	isSet bool
}

func (v NullableECSTasks) Get() *ECSTasks {
	return v.value
}

func (v *NullableECSTasks) Set(val *ECSTasks) {
	v.value = val
	v.isSet = true
}

func (v NullableECSTasks) IsSet() bool {
	return v.isSet
}

func (v *NullableECSTasks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECSTasks(val *ECSTasks) *NullableECSTasks {
	return &NullableECSTasks{value: val, isSet: true}
}

func (v NullableECSTasks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECSTasks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


