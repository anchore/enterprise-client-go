/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.7.2
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the ECSContainers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ECSContainers{}

// ECSContainers Containers defined in ECS
type ECSContainers struct {
	Containers []ECSContainer `json:"containers,omitempty"`
}

// NewECSContainers instantiates a new ECSContainers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewECSContainers() *ECSContainers {
	this := ECSContainers{}
	return &this
}

// NewECSContainersWithDefaults instantiates a new ECSContainers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewECSContainersWithDefaults() *ECSContainers {
	this := ECSContainers{}
	return &this
}

// GetContainers returns the Containers field value if set, zero value otherwise.
func (o *ECSContainers) GetContainers() []ECSContainer {
	if o == nil || IsNil(o.Containers) {
		var ret []ECSContainer
		return ret
	}
	return o.Containers
}

// GetContainersOk returns a tuple with the Containers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSContainers) GetContainersOk() ([]ECSContainer, bool) {
	if o == nil || IsNil(o.Containers) {
		return nil, false
	}
	return o.Containers, true
}

// HasContainers returns a boolean if a field has been set.
func (o *ECSContainers) HasContainers() bool {
	if o != nil && !IsNil(o.Containers) {
		return true
	}

	return false
}

// SetContainers gets a reference to the given []ECSContainer and assigns it to the Containers field.
func (o *ECSContainers) SetContainers(v []ECSContainer) {
	o.Containers = v
}

func (o ECSContainers) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ECSContainers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Containers) {
		toSerialize["containers"] = o.Containers
	}
	return toSerialize, nil
}

type NullableECSContainers struct {
	value *ECSContainers
	isSet bool
}

func (v NullableECSContainers) Get() *ECSContainers {
	return v.value
}

func (v *NullableECSContainers) Set(val *ECSContainers) {
	v.value = val
	v.isSet = true
}

func (v NullableECSContainers) IsSet() bool {
	return v.isSet
}

func (v *NullableECSContainers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECSContainers(val *ECSContainers) *NullableECSContainers {
	return &NullableECSContainers{value: val, isSet: true}
}

func (v NullableECSContainers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECSContainers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


