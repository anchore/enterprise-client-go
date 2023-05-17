/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 1.0.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// ServiceVersion Version information for a service
type ServiceVersion struct {
	Service *NotificationServiceVersionService `json:"service,omitempty"`
	Api *NotificationServiceVersionApi `json:"api,omitempty"`
	Db *NotificationServiceVersionDb `json:"db,omitempty"`
}

// NewServiceVersion instantiates a new ServiceVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceVersion() *ServiceVersion {
	this := ServiceVersion{}
	return &this
}

// NewServiceVersionWithDefaults instantiates a new ServiceVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceVersionWithDefaults() *ServiceVersion {
	this := ServiceVersion{}
	return &this
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *ServiceVersion) GetService() NotificationServiceVersionService {
	if o == nil || o.Service == nil {
		var ret NotificationServiceVersionService
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVersion) GetServiceOk() (*NotificationServiceVersionService, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *ServiceVersion) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given NotificationServiceVersionService and assigns it to the Service field.
func (o *ServiceVersion) SetService(v NotificationServiceVersionService) {
	o.Service = &v
}

// GetApi returns the Api field value if set, zero value otherwise.
func (o *ServiceVersion) GetApi() NotificationServiceVersionApi {
	if o == nil || o.Api == nil {
		var ret NotificationServiceVersionApi
		return ret
	}
	return *o.Api
}

// GetApiOk returns a tuple with the Api field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVersion) GetApiOk() (*NotificationServiceVersionApi, bool) {
	if o == nil || o.Api == nil {
		return nil, false
	}
	return o.Api, true
}

// HasApi returns a boolean if a field has been set.
func (o *ServiceVersion) HasApi() bool {
	if o != nil && o.Api != nil {
		return true
	}

	return false
}

// SetApi gets a reference to the given NotificationServiceVersionApi and assigns it to the Api field.
func (o *ServiceVersion) SetApi(v NotificationServiceVersionApi) {
	o.Api = &v
}

// GetDb returns the Db field value if set, zero value otherwise.
func (o *ServiceVersion) GetDb() NotificationServiceVersionDb {
	if o == nil || o.Db == nil {
		var ret NotificationServiceVersionDb
		return ret
	}
	return *o.Db
}

// GetDbOk returns a tuple with the Db field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVersion) GetDbOk() (*NotificationServiceVersionDb, bool) {
	if o == nil || o.Db == nil {
		return nil, false
	}
	return o.Db, true
}

// HasDb returns a boolean if a field has been set.
func (o *ServiceVersion) HasDb() bool {
	if o != nil && o.Db != nil {
		return true
	}

	return false
}

// SetDb gets a reference to the given NotificationServiceVersionDb and assigns it to the Db field.
func (o *ServiceVersion) SetDb(v NotificationServiceVersionDb) {
	o.Db = &v
}

func (o ServiceVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.Api != nil {
		toSerialize["api"] = o.Api
	}
	if o.Db != nil {
		toSerialize["db"] = o.Db
	}
	return json.Marshal(toSerialize)
}

type NullableServiceVersion struct {
	value *ServiceVersion
	isSet bool
}

func (v NullableServiceVersion) Get() *ServiceVersion {
	return v.value
}

func (v *NullableServiceVersion) Set(val *ServiceVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceVersion(val *ServiceVersion) *NullableServiceVersion {
	return &NullableServiceVersion{value: val, isSet: true}
}

func (v NullableServiceVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


