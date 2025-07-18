/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.12.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the KubernetesInventory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesInventory{}

// KubernetesInventory struct for KubernetesInventory
type KubernetesInventory struct {
	ClusterName string `json:"cluster_name"`
	Timestamp time.Time `json:"timestamp"`
	Namespaces []KubernetesInventoryNamespace `json:"namespaces,omitempty"`
	Nodes []KubernetesInventoryNode `json:"nodes,omitempty"`
	Pods []KubernetesInventoryPod `json:"pods,omitempty"`
	Containers []KubernetesInventoryContainer `json:"containers,omitempty"`
}

type _KubernetesInventory KubernetesInventory

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
	if o == nil || IsNil(o.Namespaces) {
		var ret []KubernetesInventoryNamespace
		return ret
	}
	return o.Namespaces
}

// GetNamespacesOk returns a tuple with the Namespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInventory) GetNamespacesOk() ([]KubernetesInventoryNamespace, bool) {
	if o == nil || IsNil(o.Namespaces) {
		return nil, false
	}
	return o.Namespaces, true
}

// HasNamespaces returns a boolean if a field has been set.
func (o *KubernetesInventory) HasNamespaces() bool {
	if o != nil && !IsNil(o.Namespaces) {
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
	if o == nil || IsNil(o.Nodes) {
		var ret []KubernetesInventoryNode
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInventory) GetNodesOk() ([]KubernetesInventoryNode, bool) {
	if o == nil || IsNil(o.Nodes) {
		return nil, false
	}
	return o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *KubernetesInventory) HasNodes() bool {
	if o != nil && !IsNil(o.Nodes) {
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
	if o == nil || IsNil(o.Pods) {
		var ret []KubernetesInventoryPod
		return ret
	}
	return o.Pods
}

// GetPodsOk returns a tuple with the Pods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInventory) GetPodsOk() ([]KubernetesInventoryPod, bool) {
	if o == nil || IsNil(o.Pods) {
		return nil, false
	}
	return o.Pods, true
}

// HasPods returns a boolean if a field has been set.
func (o *KubernetesInventory) HasPods() bool {
	if o != nil && !IsNil(o.Pods) {
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
	if o == nil || IsNil(o.Containers) {
		var ret []KubernetesInventoryContainer
		return ret
	}
	return o.Containers
}

// GetContainersOk returns a tuple with the Containers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInventory) GetContainersOk() ([]KubernetesInventoryContainer, bool) {
	if o == nil || IsNil(o.Containers) {
		return nil, false
	}
	return o.Containers, true
}

// HasContainers returns a boolean if a field has been set.
func (o *KubernetesInventory) HasContainers() bool {
	if o != nil && !IsNil(o.Containers) {
		return true
	}

	return false
}

// SetContainers gets a reference to the given []KubernetesInventoryContainer and assigns it to the Containers field.
func (o *KubernetesInventory) SetContainers(v []KubernetesInventoryContainer) {
	o.Containers = v
}

func (o KubernetesInventory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesInventory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cluster_name"] = o.ClusterName
	toSerialize["timestamp"] = o.Timestamp
	if !IsNil(o.Namespaces) {
		toSerialize["namespaces"] = o.Namespaces
	}
	if !IsNil(o.Nodes) {
		toSerialize["nodes"] = o.Nodes
	}
	if !IsNil(o.Pods) {
		toSerialize["pods"] = o.Pods
	}
	if !IsNil(o.Containers) {
		toSerialize["containers"] = o.Containers
	}
	return toSerialize, nil
}

func (o *KubernetesInventory) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cluster_name",
		"timestamp",
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

	varKubernetesInventory := _KubernetesInventory{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKubernetesInventory)

	if err != nil {
		return err
	}

	*o = KubernetesInventory(varKubernetesInventory)

	return err
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


