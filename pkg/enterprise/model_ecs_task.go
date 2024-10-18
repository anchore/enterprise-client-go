/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.8.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ECSTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ECSTask{}

// ECSTask struct for ECSTask
type ECSTask struct {
	Arn string `json:"arn"`
	ClusterArn string `json:"cluster_arn"`
	ServiceArn string `json:"service_arn"`
	TaskDefinitionArn string `json:"task_definition_arn"`
	Tags map[string]string `json:"tags"`
	AccountName string `json:"account_name"`
}

type _ECSTask ECSTask

// NewECSTask instantiates a new ECSTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewECSTask(arn string, clusterArn string, serviceArn string, taskDefinitionArn string, tags map[string]string, accountName string) *ECSTask {
	this := ECSTask{}
	this.Arn = arn
	this.ClusterArn = clusterArn
	this.ServiceArn = serviceArn
	this.TaskDefinitionArn = taskDefinitionArn
	this.Tags = tags
	this.AccountName = accountName
	return &this
}

// NewECSTaskWithDefaults instantiates a new ECSTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewECSTaskWithDefaults() *ECSTask {
	this := ECSTask{}
	return &this
}

// GetArn returns the Arn field value
func (o *ECSTask) GetArn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Arn
}

// GetArnOk returns a tuple with the Arn field value
// and a boolean to check if the value has been set.
func (o *ECSTask) GetArnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Arn, true
}

// SetArn sets field value
func (o *ECSTask) SetArn(v string) {
	o.Arn = v
}

// GetClusterArn returns the ClusterArn field value
func (o *ECSTask) GetClusterArn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterArn
}

// GetClusterArnOk returns a tuple with the ClusterArn field value
// and a boolean to check if the value has been set.
func (o *ECSTask) GetClusterArnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterArn, true
}

// SetClusterArn sets field value
func (o *ECSTask) SetClusterArn(v string) {
	o.ClusterArn = v
}

// GetServiceArn returns the ServiceArn field value
func (o *ECSTask) GetServiceArn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceArn
}

// GetServiceArnOk returns a tuple with the ServiceArn field value
// and a boolean to check if the value has been set.
func (o *ECSTask) GetServiceArnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceArn, true
}

// SetServiceArn sets field value
func (o *ECSTask) SetServiceArn(v string) {
	o.ServiceArn = v
}

// GetTaskDefinitionArn returns the TaskDefinitionArn field value
func (o *ECSTask) GetTaskDefinitionArn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TaskDefinitionArn
}

// GetTaskDefinitionArnOk returns a tuple with the TaskDefinitionArn field value
// and a boolean to check if the value has been set.
func (o *ECSTask) GetTaskDefinitionArnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskDefinitionArn, true
}

// SetTaskDefinitionArn sets field value
func (o *ECSTask) SetTaskDefinitionArn(v string) {
	o.TaskDefinitionArn = v
}

// GetTags returns the Tags field value
func (o *ECSTask) GetTags() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *ECSTask) GetTagsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value
func (o *ECSTask) SetTags(v map[string]string) {
	o.Tags = v
}

// GetAccountName returns the AccountName field value
func (o *ECSTask) GetAccountName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value
// and a boolean to check if the value has been set.
func (o *ECSTask) GetAccountNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountName, true
}

// SetAccountName sets field value
func (o *ECSTask) SetAccountName(v string) {
	o.AccountName = v
}

func (o ECSTask) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ECSTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["arn"] = o.Arn
	toSerialize["cluster_arn"] = o.ClusterArn
	toSerialize["service_arn"] = o.ServiceArn
	toSerialize["task_definition_arn"] = o.TaskDefinitionArn
	toSerialize["tags"] = o.Tags
	toSerialize["account_name"] = o.AccountName
	return toSerialize, nil
}

func (o *ECSTask) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"arn",
		"cluster_arn",
		"service_arn",
		"task_definition_arn",
		"tags",
		"account_name",
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

	varECSTask := _ECSTask{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varECSTask)

	if err != nil {
		return err
	}

	*o = ECSTask(varECSTask)

	return err
}

type NullableECSTask struct {
	value *ECSTask
	isSet bool
}

func (v NullableECSTask) Get() *ECSTask {
	return v.value
}

func (v *NullableECSTask) Set(val *ECSTask) {
	v.value = val
	v.isSet = true
}

func (v NullableECSTask) IsSet() bool {
	return v.isSet
}

func (v *NullableECSTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECSTask(val *ECSTask) *NullableECSTask {
	return &NullableECSTask{value: val, isSet: true}
}

func (v NullableECSTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECSTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


