/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.2.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// ECSInventoryTask struct for ECSInventoryTask
type ECSInventoryTask struct {
	Arn string `json:"arn"`
	ServiceArn *string `json:"service_arn,omitempty"`
	TaskDefinitionArn *string `json:"task_definition_arn,omitempty"`
	Tags *map[string]string `json:"tags,omitempty"`
}

// NewECSInventoryTask instantiates a new ECSInventoryTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewECSInventoryTask(arn string) *ECSInventoryTask {
	this := ECSInventoryTask{}
	this.Arn = arn
	return &this
}

// NewECSInventoryTaskWithDefaults instantiates a new ECSInventoryTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewECSInventoryTaskWithDefaults() *ECSInventoryTask {
	this := ECSInventoryTask{}
	return &this
}

// GetArn returns the Arn field value
func (o *ECSInventoryTask) GetArn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Arn
}

// GetArnOk returns a tuple with the Arn field value
// and a boolean to check if the value has been set.
func (o *ECSInventoryTask) GetArnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Arn, true
}

// SetArn sets field value
func (o *ECSInventoryTask) SetArn(v string) {
	o.Arn = v
}

// GetServiceArn returns the ServiceArn field value if set, zero value otherwise.
func (o *ECSInventoryTask) GetServiceArn() string {
	if o == nil || o.ServiceArn == nil {
		var ret string
		return ret
	}
	return *o.ServiceArn
}

// GetServiceArnOk returns a tuple with the ServiceArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSInventoryTask) GetServiceArnOk() (*string, bool) {
	if o == nil || o.ServiceArn == nil {
		return nil, false
	}
	return o.ServiceArn, true
}

// HasServiceArn returns a boolean if a field has been set.
func (o *ECSInventoryTask) HasServiceArn() bool {
	if o != nil && o.ServiceArn != nil {
		return true
	}

	return false
}

// SetServiceArn gets a reference to the given string and assigns it to the ServiceArn field.
func (o *ECSInventoryTask) SetServiceArn(v string) {
	o.ServiceArn = &v
}

// GetTaskDefinitionArn returns the TaskDefinitionArn field value if set, zero value otherwise.
func (o *ECSInventoryTask) GetTaskDefinitionArn() string {
	if o == nil || o.TaskDefinitionArn == nil {
		var ret string
		return ret
	}
	return *o.TaskDefinitionArn
}

// GetTaskDefinitionArnOk returns a tuple with the TaskDefinitionArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSInventoryTask) GetTaskDefinitionArnOk() (*string, bool) {
	if o == nil || o.TaskDefinitionArn == nil {
		return nil, false
	}
	return o.TaskDefinitionArn, true
}

// HasTaskDefinitionArn returns a boolean if a field has been set.
func (o *ECSInventoryTask) HasTaskDefinitionArn() bool {
	if o != nil && o.TaskDefinitionArn != nil {
		return true
	}

	return false
}

// SetTaskDefinitionArn gets a reference to the given string and assigns it to the TaskDefinitionArn field.
func (o *ECSInventoryTask) SetTaskDefinitionArn(v string) {
	o.TaskDefinitionArn = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ECSInventoryTask) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSInventoryTask) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ECSInventoryTask) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *ECSInventoryTask) SetTags(v map[string]string) {
	o.Tags = &v
}

func (o ECSInventoryTask) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["arn"] = o.Arn
	}
	if o.ServiceArn != nil {
		toSerialize["service_arn"] = o.ServiceArn
	}
	if o.TaskDefinitionArn != nil {
		toSerialize["task_definition_arn"] = o.TaskDefinitionArn
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableECSInventoryTask struct {
	value *ECSInventoryTask
	isSet bool
}

func (v NullableECSInventoryTask) Get() *ECSInventoryTask {
	return v.value
}

func (v *NullableECSInventoryTask) Set(val *ECSInventoryTask) {
	v.value = val
	v.isSet = true
}

func (v NullableECSInventoryTask) IsSet() bool {
	return v.isSet
}

func (v *NullableECSInventoryTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECSInventoryTask(val *ECSInventoryTask) *NullableECSInventoryTask {
	return &NullableECSInventoryTask{value: val, isSet: true}
}

func (v NullableECSInventoryTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECSInventoryTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


