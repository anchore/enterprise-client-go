/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.4.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// ECSInventory struct for ECSInventory
type ECSInventory struct {
	ClusterArn string `json:"cluster_arn"`
	Timestamp time.Time `json:"timestamp"`
	Tasks []ECSInventoryTask `json:"tasks,omitempty"`
	Containers []ECSInventoryContainer `json:"containers,omitempty"`
	Services []ECSInventoryService `json:"services,omitempty"`
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
func (o *ECSInventory) GetTasks() []ECSInventoryTask {
	if o == nil {
		var ret []ECSInventoryTask
		return ret
	}
	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ECSInventory) GetTasksOk() ([]ECSInventoryTask, bool) {
	if o == nil || o.Tasks == nil {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *ECSInventory) HasTasks() bool {
	if o != nil && o.Tasks != nil {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []ECSInventoryTask and assigns it to the Tasks field.
func (o *ECSInventory) SetTasks(v []ECSInventoryTask) {
	o.Tasks = v
}

// GetContainers returns the Containers field value if set, zero value otherwise.
func (o *ECSInventory) GetContainers() []ECSInventoryContainer {
	if o == nil || o.Containers == nil {
		var ret []ECSInventoryContainer
		return ret
	}
	return o.Containers
}

// GetContainersOk returns a tuple with the Containers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSInventory) GetContainersOk() ([]ECSInventoryContainer, bool) {
	if o == nil || o.Containers == nil {
		return nil, false
	}
	return o.Containers, true
}

// HasContainers returns a boolean if a field has been set.
func (o *ECSInventory) HasContainers() bool {
	if o != nil && o.Containers != nil {
		return true
	}

	return false
}

// SetContainers gets a reference to the given []ECSInventoryContainer and assigns it to the Containers field.
func (o *ECSInventory) SetContainers(v []ECSInventoryContainer) {
	o.Containers = v
}

// GetServices returns the Services field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ECSInventory) GetServices() []ECSInventoryService {
	if o == nil {
		var ret []ECSInventoryService
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ECSInventory) GetServicesOk() ([]ECSInventoryService, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *ECSInventory) HasServices() bool {
	if o != nil && o.Services != nil {
		return true
	}

	return false
}

// SetServices gets a reference to the given []ECSInventoryService and assigns it to the Services field.
func (o *ECSInventory) SetServices(v []ECSInventoryService) {
	o.Services = v
}

func (o ECSInventory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cluster_arn"] = o.ClusterArn
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Tasks != nil {
		toSerialize["tasks"] = o.Tasks
	}
	if o.Containers != nil {
		toSerialize["containers"] = o.Containers
	}
	if o.Services != nil {
		toSerialize["services"] = o.Services
	}
	return json.Marshal(toSerialize)
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


