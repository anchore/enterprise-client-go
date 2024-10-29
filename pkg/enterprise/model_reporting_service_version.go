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
)

// checks if the ReportingServiceVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportingServiceVersion{}

// ReportingServiceVersion Version information for a service
type ReportingServiceVersion struct {
	Service *NotificationServiceVersionService `json:"service,omitempty"`
	Api *NotificationServiceVersionApi `json:"api,omitempty"`
	Db *NotificationServiceVersionDb `json:"db,omitempty"`
	Engine *NotificationServiceVersionEngine `json:"engine,omitempty"`
}

// NewReportingServiceVersion instantiates a new ReportingServiceVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportingServiceVersion() *ReportingServiceVersion {
	this := ReportingServiceVersion{}
	return &this
}

// NewReportingServiceVersionWithDefaults instantiates a new ReportingServiceVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportingServiceVersionWithDefaults() *ReportingServiceVersion {
	this := ReportingServiceVersion{}
	return &this
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *ReportingServiceVersion) GetService() NotificationServiceVersionService {
	if o == nil || IsNil(o.Service) {
		var ret NotificationServiceVersionService
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingServiceVersion) GetServiceOk() (*NotificationServiceVersionService, bool) {
	if o == nil || IsNil(o.Service) {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *ReportingServiceVersion) HasService() bool {
	if o != nil && !IsNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given NotificationServiceVersionService and assigns it to the Service field.
func (o *ReportingServiceVersion) SetService(v NotificationServiceVersionService) {
	o.Service = &v
}

// GetApi returns the Api field value if set, zero value otherwise.
func (o *ReportingServiceVersion) GetApi() NotificationServiceVersionApi {
	if o == nil || IsNil(o.Api) {
		var ret NotificationServiceVersionApi
		return ret
	}
	return *o.Api
}

// GetApiOk returns a tuple with the Api field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingServiceVersion) GetApiOk() (*NotificationServiceVersionApi, bool) {
	if o == nil || IsNil(o.Api) {
		return nil, false
	}
	return o.Api, true
}

// HasApi returns a boolean if a field has been set.
func (o *ReportingServiceVersion) HasApi() bool {
	if o != nil && !IsNil(o.Api) {
		return true
	}

	return false
}

// SetApi gets a reference to the given NotificationServiceVersionApi and assigns it to the Api field.
func (o *ReportingServiceVersion) SetApi(v NotificationServiceVersionApi) {
	o.Api = &v
}

// GetDb returns the Db field value if set, zero value otherwise.
func (o *ReportingServiceVersion) GetDb() NotificationServiceVersionDb {
	if o == nil || IsNil(o.Db) {
		var ret NotificationServiceVersionDb
		return ret
	}
	return *o.Db
}

// GetDbOk returns a tuple with the Db field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingServiceVersion) GetDbOk() (*NotificationServiceVersionDb, bool) {
	if o == nil || IsNil(o.Db) {
		return nil, false
	}
	return o.Db, true
}

// HasDb returns a boolean if a field has been set.
func (o *ReportingServiceVersion) HasDb() bool {
	if o != nil && !IsNil(o.Db) {
		return true
	}

	return false
}

// SetDb gets a reference to the given NotificationServiceVersionDb and assigns it to the Db field.
func (o *ReportingServiceVersion) SetDb(v NotificationServiceVersionDb) {
	o.Db = &v
}

// GetEngine returns the Engine field value if set, zero value otherwise.
func (o *ReportingServiceVersion) GetEngine() NotificationServiceVersionEngine {
	if o == nil || IsNil(o.Engine) {
		var ret NotificationServiceVersionEngine
		return ret
	}
	return *o.Engine
}

// GetEngineOk returns a tuple with the Engine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingServiceVersion) GetEngineOk() (*NotificationServiceVersionEngine, bool) {
	if o == nil || IsNil(o.Engine) {
		return nil, false
	}
	return o.Engine, true
}

// HasEngine returns a boolean if a field has been set.
func (o *ReportingServiceVersion) HasEngine() bool {
	if o != nil && !IsNil(o.Engine) {
		return true
	}

	return false
}

// SetEngine gets a reference to the given NotificationServiceVersionEngine and assigns it to the Engine field.
func (o *ReportingServiceVersion) SetEngine(v NotificationServiceVersionEngine) {
	o.Engine = &v
}

func (o ReportingServiceVersion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportingServiceVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	if !IsNil(o.Api) {
		toSerialize["api"] = o.Api
	}
	if !IsNil(o.Db) {
		toSerialize["db"] = o.Db
	}
	if !IsNil(o.Engine) {
		toSerialize["engine"] = o.Engine
	}
	return toSerialize, nil
}

type NullableReportingServiceVersion struct {
	value *ReportingServiceVersion
	isSet bool
}

func (v NullableReportingServiceVersion) Get() *ReportingServiceVersion {
	return v.value
}

func (v *NullableReportingServiceVersion) Set(val *ReportingServiceVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableReportingServiceVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableReportingServiceVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportingServiceVersion(val *ReportingServiceVersion) *NullableReportingServiceVersion {
	return &NullableReportingServiceVersion{value: val, isSet: true}
}

func (v NullableReportingServiceVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportingServiceVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


