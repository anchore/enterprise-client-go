/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.6.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// RbacManagerServiceVersion Version information for a service
type RbacManagerServiceVersion struct {
	Service *NotificationServiceVersionService `json:"service,omitempty"`
	Api *NotificationServiceVersionApi `json:"api,omitempty"`
	Db *NotificationServiceVersionDb `json:"db,omitempty"`
	Engine *NotificationServiceVersionEngine `json:"engine,omitempty"`
}

// NewRbacManagerServiceVersion instantiates a new RbacManagerServiceVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRbacManagerServiceVersion() *RbacManagerServiceVersion {
	this := RbacManagerServiceVersion{}
	return &this
}

// NewRbacManagerServiceVersionWithDefaults instantiates a new RbacManagerServiceVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRbacManagerServiceVersionWithDefaults() *RbacManagerServiceVersion {
	this := RbacManagerServiceVersion{}
	return &this
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *RbacManagerServiceVersion) GetService() NotificationServiceVersionService {
	if o == nil || o.Service == nil {
		var ret NotificationServiceVersionService
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerServiceVersion) GetServiceOk() (*NotificationServiceVersionService, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *RbacManagerServiceVersion) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given NotificationServiceVersionService and assigns it to the Service field.
func (o *RbacManagerServiceVersion) SetService(v NotificationServiceVersionService) {
	o.Service = &v
}

// GetApi returns the Api field value if set, zero value otherwise.
func (o *RbacManagerServiceVersion) GetApi() NotificationServiceVersionApi {
	if o == nil || o.Api == nil {
		var ret NotificationServiceVersionApi
		return ret
	}
	return *o.Api
}

// GetApiOk returns a tuple with the Api field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerServiceVersion) GetApiOk() (*NotificationServiceVersionApi, bool) {
	if o == nil || o.Api == nil {
		return nil, false
	}
	return o.Api, true
}

// HasApi returns a boolean if a field has been set.
func (o *RbacManagerServiceVersion) HasApi() bool {
	if o != nil && o.Api != nil {
		return true
	}

	return false
}

// SetApi gets a reference to the given NotificationServiceVersionApi and assigns it to the Api field.
func (o *RbacManagerServiceVersion) SetApi(v NotificationServiceVersionApi) {
	o.Api = &v
}

// GetDb returns the Db field value if set, zero value otherwise.
func (o *RbacManagerServiceVersion) GetDb() NotificationServiceVersionDb {
	if o == nil || o.Db == nil {
		var ret NotificationServiceVersionDb
		return ret
	}
	return *o.Db
}

// GetDbOk returns a tuple with the Db field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerServiceVersion) GetDbOk() (*NotificationServiceVersionDb, bool) {
	if o == nil || o.Db == nil {
		return nil, false
	}
	return o.Db, true
}

// HasDb returns a boolean if a field has been set.
func (o *RbacManagerServiceVersion) HasDb() bool {
	if o != nil && o.Db != nil {
		return true
	}

	return false
}

// SetDb gets a reference to the given NotificationServiceVersionDb and assigns it to the Db field.
func (o *RbacManagerServiceVersion) SetDb(v NotificationServiceVersionDb) {
	o.Db = &v
}

// GetEngine returns the Engine field value if set, zero value otherwise.
func (o *RbacManagerServiceVersion) GetEngine() NotificationServiceVersionEngine {
	if o == nil || o.Engine == nil {
		var ret NotificationServiceVersionEngine
		return ret
	}
	return *o.Engine
}

// GetEngineOk returns a tuple with the Engine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerServiceVersion) GetEngineOk() (*NotificationServiceVersionEngine, bool) {
	if o == nil || o.Engine == nil {
		return nil, false
	}
	return o.Engine, true
}

// HasEngine returns a boolean if a field has been set.
func (o *RbacManagerServiceVersion) HasEngine() bool {
	if o != nil && o.Engine != nil {
		return true
	}

	return false
}

// SetEngine gets a reference to the given NotificationServiceVersionEngine and assigns it to the Engine field.
func (o *RbacManagerServiceVersion) SetEngine(v NotificationServiceVersionEngine) {
	o.Engine = &v
}

func (o RbacManagerServiceVersion) MarshalJSON() ([]byte, error) {
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
	if o.Engine != nil {
		toSerialize["engine"] = o.Engine
	}
	return json.Marshal(toSerialize)
}

type NullableRbacManagerServiceVersion struct {
	value *RbacManagerServiceVersion
	isSet bool
}

func (v NullableRbacManagerServiceVersion) Get() *RbacManagerServiceVersion {
	return v.value
}

func (v *NullableRbacManagerServiceVersion) Set(val *RbacManagerServiceVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableRbacManagerServiceVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableRbacManagerServiceVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRbacManagerServiceVersion(val *RbacManagerServiceVersion) *NullableRbacManagerServiceVersion {
	return &NullableRbacManagerServiceVersion{value: val, isSet: true}
}

func (v NullableRbacManagerServiceVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRbacManagerServiceVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


