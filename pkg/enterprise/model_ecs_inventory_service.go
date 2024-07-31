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

// ECSInventoryService struct for ECSInventoryService
type ECSInventoryService struct {
	Arn string `json:"arn"`
	Tags *map[string]string `json:"tags,omitempty"`
}

// NewECSInventoryService instantiates a new ECSInventoryService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewECSInventoryService(arn string) *ECSInventoryService {
	this := ECSInventoryService{}
	this.Arn = arn
	return &this
}

// NewECSInventoryServiceWithDefaults instantiates a new ECSInventoryService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewECSInventoryServiceWithDefaults() *ECSInventoryService {
	this := ECSInventoryService{}
	return &this
}

// GetArn returns the Arn field value
func (o *ECSInventoryService) GetArn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Arn
}

// GetArnOk returns a tuple with the Arn field value
// and a boolean to check if the value has been set.
func (o *ECSInventoryService) GetArnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Arn, true
}

// SetArn sets field value
func (o *ECSInventoryService) SetArn(v string) {
	o.Arn = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ECSInventoryService) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSInventoryService) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ECSInventoryService) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *ECSInventoryService) SetTags(v map[string]string) {
	o.Tags = &v
}

func (o ECSInventoryService) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["arn"] = o.Arn
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableECSInventoryService struct {
	value *ECSInventoryService
	isSet bool
}

func (v NullableECSInventoryService) Get() *ECSInventoryService {
	return v.value
}

func (v *NullableECSInventoryService) Set(val *ECSInventoryService) {
	v.value = val
	v.isSet = true
}

func (v NullableECSInventoryService) IsSet() bool {
	return v.isSet
}

func (v *NullableECSInventoryService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECSInventoryService(val *ECSInventoryService) *NullableECSInventoryService {
	return &NullableECSInventoryService{value: val, isSet: true}
}

func (v NullableECSInventoryService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECSInventoryService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


