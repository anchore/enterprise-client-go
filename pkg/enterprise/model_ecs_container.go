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
	"bytes"
	"fmt"
)

// checks if the ECSContainer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ECSContainer{}

// ECSContainer struct for ECSContainer
type ECSContainer struct {
	Arn string `json:"arn"`
	TaskArn string `json:"task_arn"`
	AccountName string `json:"account_name"`
	Context string `json:"context"`
	ImageTag string `json:"image_tag"`
	ImageDigest string `json:"image_digest"`
}

type _ECSContainer ECSContainer

// NewECSContainer instantiates a new ECSContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewECSContainer(arn string, taskArn string, accountName string, context string, imageTag string, imageDigest string) *ECSContainer {
	this := ECSContainer{}
	this.Arn = arn
	this.TaskArn = taskArn
	this.AccountName = accountName
	this.Context = context
	this.ImageTag = imageTag
	this.ImageDigest = imageDigest
	return &this
}

// NewECSContainerWithDefaults instantiates a new ECSContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewECSContainerWithDefaults() *ECSContainer {
	this := ECSContainer{}
	return &this
}

// GetArn returns the Arn field value
func (o *ECSContainer) GetArn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Arn
}

// GetArnOk returns a tuple with the Arn field value
// and a boolean to check if the value has been set.
func (o *ECSContainer) GetArnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Arn, true
}

// SetArn sets field value
func (o *ECSContainer) SetArn(v string) {
	o.Arn = v
}

// GetTaskArn returns the TaskArn field value
func (o *ECSContainer) GetTaskArn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TaskArn
}

// GetTaskArnOk returns a tuple with the TaskArn field value
// and a boolean to check if the value has been set.
func (o *ECSContainer) GetTaskArnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskArn, true
}

// SetTaskArn sets field value
func (o *ECSContainer) SetTaskArn(v string) {
	o.TaskArn = v
}

// GetAccountName returns the AccountName field value
func (o *ECSContainer) GetAccountName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value
// and a boolean to check if the value has been set.
func (o *ECSContainer) GetAccountNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountName, true
}

// SetAccountName sets field value
func (o *ECSContainer) SetAccountName(v string) {
	o.AccountName = v
}

// GetContext returns the Context field value
func (o *ECSContainer) GetContext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *ECSContainer) GetContextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *ECSContainer) SetContext(v string) {
	o.Context = v
}

// GetImageTag returns the ImageTag field value
func (o *ECSContainer) GetImageTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageTag
}

// GetImageTagOk returns a tuple with the ImageTag field value
// and a boolean to check if the value has been set.
func (o *ECSContainer) GetImageTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageTag, true
}

// SetImageTag sets field value
func (o *ECSContainer) SetImageTag(v string) {
	o.ImageTag = v
}

// GetImageDigest returns the ImageDigest field value
func (o *ECSContainer) GetImageDigest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageDigest
}

// GetImageDigestOk returns a tuple with the ImageDigest field value
// and a boolean to check if the value has been set.
func (o *ECSContainer) GetImageDigestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageDigest, true
}

// SetImageDigest sets field value
func (o *ECSContainer) SetImageDigest(v string) {
	o.ImageDigest = v
}

func (o ECSContainer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ECSContainer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["arn"] = o.Arn
	toSerialize["task_arn"] = o.TaskArn
	toSerialize["account_name"] = o.AccountName
	toSerialize["context"] = o.Context
	toSerialize["image_tag"] = o.ImageTag
	toSerialize["image_digest"] = o.ImageDigest
	return toSerialize, nil
}

func (o *ECSContainer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"arn",
		"task_arn",
		"account_name",
		"context",
		"image_tag",
		"image_digest",
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

	varECSContainer := _ECSContainer{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varECSContainer)

	if err != nil {
		return err
	}

	*o = ECSContainer(varECSContainer)

	return err
}

type NullableECSContainer struct {
	value *ECSContainer
	isSet bool
}

func (v NullableECSContainer) Get() *ECSContainer {
	return v.value
}

func (v *NullableECSContainer) Set(val *ECSContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableECSContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableECSContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECSContainer(val *ECSContainer) *NullableECSContainer {
	return &NullableECSContainer{value: val, isSet: true}
}

func (v NullableECSContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECSContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


