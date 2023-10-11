/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.0.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// checks if the ECSInventory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ECSInventory{}

// ECSInventory struct for ECSInventory
type ECSInventory struct {
	ClusterArn string `json:"cluster_arn"`
	Timestamp time.Time `json:"timestamp"`
	Tasks []ECSInventoryTasksInner `json:"tasks,omitempty"`
	Containers []ECSInventoryContainersInner `json:"containers,omitempty"`
	Services []ECSInventoryServicesInner `json:"services,omitempty"`
}

// NewECSInventory instantiates a new ECSInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewECSInventory(clusterArn string, timestamp time.Time) *ECSInventory {
	this := ECSInventory{}
	this.ClusterArn = clusterArn
	this.Timestamp = timestamp
	return &this
}

// NewECSInventoryWithDefaults instantiates a new ECSInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewECSInventoryWithDefaults() *ECSInventory {
	this := ECSInventory{}
	return &this
}

// GetClusterArn returns the ClusterArn field value
func (o *ECSInventory) GetClusterArn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterArn
}

// GetClusterArnOk returns a tuple with the ClusterArn field value
// and a boolean to check if the value has been set.
func (o *ECSInventory) GetClusterArnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterArn, true
}

// SetClusterArn sets field value
func (o *ECSInventory) SetClusterArn(v string) {
	o.ClusterArn = v
}

// GetTimestamp returns the Timestamp field value
func (o *ECSInventory) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *ECSInventory) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *ECSInventory) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetTasks returns the Tasks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ECSInventory) GetTasks() []ECSInventoryTasksInner {
	if o == nil {
		var ret []ECSInventoryTasksInner
		return ret
	}
	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ECSInventory) GetTasksOk() ([]ECSInventoryTasksInner, bool) {
	if o == nil || IsNil(o.Tasks) {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *ECSInventory) HasTasks() bool {
	if o != nil && IsNil(o.Tasks) {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []ECSInventoryTasksInner and assigns it to the Tasks field.
func (o *ECSInventory) SetTasks(v []ECSInventoryTasksInner) {
	o.Tasks = v
}

// GetContainers returns the Containers field value if set, zero value otherwise.
func (o *ECSInventory) GetContainers() []ECSInventoryContainersInner {
	if o == nil || IsNil(o.Containers) {
		var ret []ECSInventoryContainersInner
		return ret
	}
	return o.Containers
}

// GetContainersOk returns a tuple with the Containers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSInventory) GetContainersOk() ([]ECSInventoryContainersInner, bool) {
	if o == nil || IsNil(o.Containers) {
		return nil, false
	}
	return o.Containers, true
}

// HasContainers returns a boolean if a field has been set.
func (o *ECSInventory) HasContainers() bool {
	if o != nil && !IsNil(o.Containers) {
		return true
	}

	return false
}

// SetContainers gets a reference to the given []ECSInventoryContainersInner and assigns it to the Containers field.
func (o *ECSInventory) SetContainers(v []ECSInventoryContainersInner) {
	o.Containers = v
}

// GetServices returns the Services field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ECSInventory) GetServices() []ECSInventoryServicesInner {
	if o == nil {
		var ret []ECSInventoryServicesInner
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ECSInventory) GetServicesOk() ([]ECSInventoryServicesInner, bool) {
	if o == nil || IsNil(o.Services) {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *ECSInventory) HasServices() bool {
	if o != nil && IsNil(o.Services) {
		return true
	}

	return false
}

// SetServices gets a reference to the given []ECSInventoryServicesInner and assigns it to the Services field.
func (o *ECSInventory) SetServices(v []ECSInventoryServicesInner) {
	o.Services = v
}

func (o ECSInventory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ECSInventory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cluster_arn"] = o.ClusterArn
	toSerialize["timestamp"] = o.Timestamp
	if o.Tasks != nil {
		toSerialize["tasks"] = o.Tasks
	}
	if !IsNil(o.Containers) {
		toSerialize["containers"] = o.Containers
	}
	if o.Services != nil {
		toSerialize["services"] = o.Services
	}
	return toSerialize, nil
}

type NullableECSInventory struct {
	value *ECSInventory
	isSet bool
}

func (v NullableECSInventory) Get() *ECSInventory {
	return v.value
}

func (v *NullableECSInventory) Set(val *ECSInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableECSInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableECSInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECSInventory(val *ECSInventory) *NullableECSInventory {
	return &NullableECSInventory{value: val, isSet: true}
}

func (v NullableECSInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECSInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


