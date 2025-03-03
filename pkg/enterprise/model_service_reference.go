/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.10.1
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the ServiceReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceReference{}

// ServiceReference A reference to a service in the system
type ServiceReference struct {
	// The unique id of the host on which the service is executing
	HostId *string `json:"host_id,omitempty"`
	// Registered service name
	ServiceName *string `json:"service_name,omitempty"`
}

// NewServiceReference instantiates a new ServiceReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceReference() *ServiceReference {
	this := ServiceReference{}
	return &this
}

// NewServiceReferenceWithDefaults instantiates a new ServiceReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceReferenceWithDefaults() *ServiceReference {
	this := ServiceReference{}
	return &this
}

// GetHostId returns the HostId field value if set, zero value otherwise.
func (o *ServiceReference) GetHostId() string {
	if o == nil || IsNil(o.HostId) {
		var ret string
		return ret
	}
	return *o.HostId
}

// GetHostIdOk returns a tuple with the HostId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceReference) GetHostIdOk() (*string, bool) {
	if o == nil || IsNil(o.HostId) {
		return nil, false
	}
	return o.HostId, true
}

// HasHostId returns a boolean if a field has been set.
func (o *ServiceReference) HasHostId() bool {
	if o != nil && !IsNil(o.HostId) {
		return true
	}

	return false
}

// SetHostId gets a reference to the given string and assigns it to the HostId field.
func (o *ServiceReference) SetHostId(v string) {
	o.HostId = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *ServiceReference) GetServiceName() string {
	if o == nil || IsNil(o.ServiceName) {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceReference) GetServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceName) {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *ServiceReference) HasServiceName() bool {
	if o != nil && !IsNil(o.ServiceName) {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *ServiceReference) SetServiceName(v string) {
	o.ServiceName = &v
}

func (o ServiceReference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HostId) {
		toSerialize["host_id"] = o.HostId
	}
	if !IsNil(o.ServiceName) {
		toSerialize["service_name"] = o.ServiceName
	}
	return toSerialize, nil
}

type NullableServiceReference struct {
	value *ServiceReference
	isSet bool
}

func (v NullableServiceReference) Get() *ServiceReference {
	return v.value
}

func (v *NullableServiceReference) Set(val *ServiceReference) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceReference) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceReference(val *ServiceReference) *NullableServiceReference {
	return &NullableServiceReference{value: val, isSet: true}
}

func (v NullableServiceReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


