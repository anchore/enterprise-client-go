/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.6.2
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// KubernetesInventory struct for KubernetesInventory
type KubernetesInventory struct {
	ClusterName string `json:"cluster_name"`
	Timestamp time.Time `json:"timestamp"`
	Namespaces []KubernetesInventoryNamespace `json:"namespaces,omitempty"`
	Nodes []KubernetesInventoryNode `json:"nodes,omitempty"`
	Pods []KubernetesInventoryPod `json:"pods,omitempty"`
	Containers []KubernetesInventoryContainer `json:"containers,omitempty"`
}

// NewKubernetesInventory instantiates a new KubernetesInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesInventory(clusterName string, timestamp time.Time) *KubernetesInventory {
	this := KubernetesInventory{}
	this.ClusterName = clusterName
	this.Timestamp = timestamp
	return &this
}

// NewKubernetesInventoryWithDefaults instantiates a new KubernetesInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesInventoryWithDefaults() *KubernetesInventory {
	this := KubernetesInventory{}
	return &this
}

// GetClusterName returns the ClusterName field value
func (o *KubernetesInventory) GetClusterName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value
// and a boolean to check if the value has been set.
func (o *KubernetesInventory) GetClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterName, true
}

// SetClusterName sets field value
func (o *KubernetesInventory) SetClusterName(v string) {
	o.ClusterName = v
}

// GetTimestamp returns the Timestamp field value
func (o *KubernetesInventory) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *KubernetesInventory) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *KubernetesInventory) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetNamespaces returns the Namespaces field value if set, zero value otherwise.
func (o *KubernetesInventory) GetNamespaces() []KubernetesInventoryNamespace {
	if o == nil || o.Namespaces == nil {
		var ret []KubernetesInventoryNamespace
		return ret
	}
	return o.Namespaces
}

// GetNamespacesOk returns a tuple with the Namespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInventory) GetNamespacesOk() ([]KubernetesInventoryNamespace, bool) {
	if o == nil || o.Namespaces == nil {
		return nil, false
	}
	return o.Namespaces, true
}

// HasNamespaces returns a boolean if a field has been set.
func (o *KubernetesInventory) HasNamespaces() bool {
	if o != nil && o.Namespaces != nil {
		return true
	}

	return false
}

// SetNamespaces gets a reference to the given []KubernetesInventoryNamespace and assigns it to the Namespaces field.
func (o *KubernetesInventory) SetNamespaces(v []KubernetesInventoryNamespace) {
	o.Namespaces = v
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *KubernetesInventory) GetNodes() []KubernetesInventoryNode {
	if o == nil || o.Nodes == nil {
		var ret []KubernetesInventoryNode
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInventory) GetNodesOk() ([]KubernetesInventoryNode, bool) {
	if o == nil || o.Nodes == nil {
		return nil, false
	}
	return o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *KubernetesInventory) HasNodes() bool {
	if o != nil && o.Nodes != nil {
		return true
	}

	return false
}

// SetNodes gets a reference to the given []KubernetesInventoryNode and assigns it to the Nodes field.
func (o *KubernetesInventory) SetNodes(v []KubernetesInventoryNode) {
	o.Nodes = v
}

// GetPods returns the Pods field value if set, zero value otherwise.
func (o *KubernetesInventory) GetPods() []KubernetesInventoryPod {
	if o == nil || o.Pods == nil {
		var ret []KubernetesInventoryPod
		return ret
	}
	return o.Pods
}

// GetPodsOk returns a tuple with the Pods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInventory) GetPodsOk() ([]KubernetesInventoryPod, bool) {
	if o == nil || o.Pods == nil {
		return nil, false
	}
	return o.Pods, true
}

// HasPods returns a boolean if a field has been set.
func (o *KubernetesInventory) HasPods() bool {
	if o != nil && o.Pods != nil {
		return true
	}

	return false
}

// SetPods gets a reference to the given []KubernetesInventoryPod and assigns it to the Pods field.
func (o *KubernetesInventory) SetPods(v []KubernetesInventoryPod) {
	o.Pods = v
}

// GetContainers returns the Containers field value if set, zero value otherwise.
func (o *KubernetesInventory) GetContainers() []KubernetesInventoryContainer {
	if o == nil || o.Containers == nil {
		var ret []KubernetesInventoryContainer
		return ret
	}
	return o.Containers
}

// GetContainersOk returns a tuple with the Containers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInventory) GetContainersOk() ([]KubernetesInventoryContainer, bool) {
	if o == nil || o.Containers == nil {
		return nil, false
	}
	return o.Containers, true
}

// HasContainers returns a boolean if a field has been set.
func (o *KubernetesInventory) HasContainers() bool {
	if o != nil && o.Containers != nil {
		return true
	}

	return false
}

// SetContainers gets a reference to the given []KubernetesInventoryContainer and assigns it to the Containers field.
func (o *KubernetesInventory) SetContainers(v []KubernetesInventoryContainer) {
	o.Containers = v
}

func (o KubernetesInventory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cluster_name"] = o.ClusterName
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Namespaces != nil {
		toSerialize["namespaces"] = o.Namespaces
	}
	if o.Nodes != nil {
		toSerialize["nodes"] = o.Nodes
	}
	if o.Pods != nil {
		toSerialize["pods"] = o.Pods
	}
	if o.Containers != nil {
		toSerialize["containers"] = o.Containers
	}
	return json.Marshal(toSerialize)
}

type NullableKubernetesInventory struct {
	value *KubernetesInventory
	isSet bool
}

func (v NullableKubernetesInventory) Get() *KubernetesInventory {
	return v.value
}

func (v *NullableKubernetesInventory) Set(val *KubernetesInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesInventory(val *KubernetesInventory) *NullableKubernetesInventory {
	return &NullableKubernetesInventory{value: val, isSet: true}
}

func (v NullableKubernetesInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


