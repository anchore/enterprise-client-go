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

// checks if the ECSInventoryContainer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ECSInventoryContainer{}

// ECSInventoryContainer struct for ECSInventoryContainer
type ECSInventoryContainer struct {
	Arn string `json:"arn"`
	TaskArn *string `json:"task_arn,omitempty"`
	ImageTag string `json:"image_tag"`
	ImageDigest *string `json:"image_digest,omitempty"`
}

type _ECSInventoryContainer ECSInventoryContainer

// NewECSInventoryContainer instantiates a new ECSInventoryContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewECSInventoryContainer(arn string, imageTag string) *ECSInventoryContainer {
	this := ECSInventoryContainer{}
	this.Arn = arn
	this.ImageTag = imageTag
	return &this
}

// NewECSInventoryContainerWithDefaults instantiates a new ECSInventoryContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewECSInventoryContainerWithDefaults() *ECSInventoryContainer {
	this := ECSInventoryContainer{}
	return &this
}

// GetArn returns the Arn field value
func (o *ECSInventoryContainer) GetArn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Arn
}

// GetArnOk returns a tuple with the Arn field value
// and a boolean to check if the value has been set.
func (o *ECSInventoryContainer) GetArnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Arn, true
}

// SetArn sets field value
func (o *ECSInventoryContainer) SetArn(v string) {
	o.Arn = v
}

// GetTaskArn returns the TaskArn field value if set, zero value otherwise.
func (o *ECSInventoryContainer) GetTaskArn() string {
	if o == nil || IsNil(o.TaskArn) {
		var ret string
		return ret
	}
	return *o.TaskArn
}

// GetTaskArnOk returns a tuple with the TaskArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSInventoryContainer) GetTaskArnOk() (*string, bool) {
	if o == nil || IsNil(o.TaskArn) {
		return nil, false
	}
	return o.TaskArn, true
}

// HasTaskArn returns a boolean if a field has been set.
func (o *ECSInventoryContainer) HasTaskArn() bool {
	if o != nil && !IsNil(o.TaskArn) {
		return true
	}

	return false
}

// SetTaskArn gets a reference to the given string and assigns it to the TaskArn field.
func (o *ECSInventoryContainer) SetTaskArn(v string) {
	o.TaskArn = &v
}

// GetImageTag returns the ImageTag field value
func (o *ECSInventoryContainer) GetImageTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageTag
}

// GetImageTagOk returns a tuple with the ImageTag field value
// and a boolean to check if the value has been set.
func (o *ECSInventoryContainer) GetImageTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageTag, true
}

// SetImageTag sets field value
func (o *ECSInventoryContainer) SetImageTag(v string) {
	o.ImageTag = v
}

// GetImageDigest returns the ImageDigest field value if set, zero value otherwise.
func (o *ECSInventoryContainer) GetImageDigest() string {
	if o == nil || IsNil(o.ImageDigest) {
		var ret string
		return ret
	}
	return *o.ImageDigest
}

// GetImageDigestOk returns a tuple with the ImageDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSInventoryContainer) GetImageDigestOk() (*string, bool) {
	if o == nil || IsNil(o.ImageDigest) {
		return nil, false
	}
	return o.ImageDigest, true
}

// HasImageDigest returns a boolean if a field has been set.
func (o *ECSInventoryContainer) HasImageDigest() bool {
	if o != nil && !IsNil(o.ImageDigest) {
		return true
	}

	return false
}

// SetImageDigest gets a reference to the given string and assigns it to the ImageDigest field.
func (o *ECSInventoryContainer) SetImageDigest(v string) {
	o.ImageDigest = &v
}

func (o ECSInventoryContainer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ECSInventoryContainer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["arn"] = o.Arn
	if !IsNil(o.TaskArn) {
		toSerialize["task_arn"] = o.TaskArn
	}
	toSerialize["image_tag"] = o.ImageTag
	if !IsNil(o.ImageDigest) {
		toSerialize["image_digest"] = o.ImageDigest
	}
	return toSerialize, nil
}

func (o *ECSInventoryContainer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"arn",
		"image_tag",
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

	varECSInventoryContainer := _ECSInventoryContainer{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varECSInventoryContainer)

	if err != nil {
		return err
	}

	*o = ECSInventoryContainer(varECSInventoryContainer)

	return err
}

type NullableECSInventoryContainer struct {
	value *ECSInventoryContainer
	isSet bool
}

func (v NullableECSInventoryContainer) Get() *ECSInventoryContainer {
	return v.value
}

func (v *NullableECSInventoryContainer) Set(val *ECSInventoryContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableECSInventoryContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableECSInventoryContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECSInventoryContainer(val *ECSInventoryContainer) *NullableECSInventoryContainer {
	return &NullableECSInventoryContainer{value: val, isSet: true}
}

func (v NullableECSInventoryContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECSInventoryContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


