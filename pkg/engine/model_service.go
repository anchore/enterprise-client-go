/*
Anchore Engine API Server

This is the Anchore Engine API. Provides the primary external API for users of the service.

API version: 0.1.21
Contact: nurmi@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package engine

import (
	"encoding/json"
)

// Service A service status record
type Service struct {
	// The unique id of the host on which the service is executing
	Hostid *string `json:"hostid,omitempty"`
	// Registered service name
	Servicename *string `json:"servicename,omitempty"`
	// The url to reach the service, including port as needed
	BaseUrl *string `json:"base_url,omitempty"`
	// A state indicating the condition of the service. Normal operation is 'registered'
	StatusMessage *string `json:"status_message,omitempty"`
	ServiceDetail *StatusResponse `json:"service_detail,omitempty"`
	Status *bool `json:"status,omitempty"`
	// The version of the service as reported by the service implementation on registration
	Version *string `json:"version,omitempty"`
}

// NewService instantiates a new Service object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewService() *Service {
	this := Service{}
	return &this
}

// NewServiceWithDefaults instantiates a new Service object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceWithDefaults() *Service {
	this := Service{}
	return &this
}

// GetHostid returns the Hostid field value if set, zero value otherwise.
func (o *Service) GetHostid() string {
	if o == nil || o.Hostid == nil {
		var ret string
		return ret
	}
	return *o.Hostid
}

// GetHostidOk returns a tuple with the Hostid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetHostidOk() (*string, bool) {
	if o == nil || o.Hostid == nil {
		return nil, false
	}
	return o.Hostid, true
}

// HasHostid returns a boolean if a field has been set.
func (o *Service) HasHostid() bool {
	if o != nil && o.Hostid != nil {
		return true
	}

	return false
}

// SetHostid gets a reference to the given string and assigns it to the Hostid field.
func (o *Service) SetHostid(v string) {
	o.Hostid = &v
}

// GetServicename returns the Servicename field value if set, zero value otherwise.
func (o *Service) GetServicename() string {
	if o == nil || o.Servicename == nil {
		var ret string
		return ret
	}
	return *o.Servicename
}

// GetServicenameOk returns a tuple with the Servicename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetServicenameOk() (*string, bool) {
	if o == nil || o.Servicename == nil {
		return nil, false
	}
	return o.Servicename, true
}

// HasServicename returns a boolean if a field has been set.
func (o *Service) HasServicename() bool {
	if o != nil && o.Servicename != nil {
		return true
	}

	return false
}

// SetServicename gets a reference to the given string and assigns it to the Servicename field.
func (o *Service) SetServicename(v string) {
	o.Servicename = &v
}

// GetBaseUrl returns the BaseUrl field value if set, zero value otherwise.
func (o *Service) GetBaseUrl() string {
	if o == nil || o.BaseUrl == nil {
		var ret string
		return ret
	}
	return *o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetBaseUrlOk() (*string, bool) {
	if o == nil || o.BaseUrl == nil {
		return nil, false
	}
	return o.BaseUrl, true
}

// HasBaseUrl returns a boolean if a field has been set.
func (o *Service) HasBaseUrl() bool {
	if o != nil && o.BaseUrl != nil {
		return true
	}

	return false
}

// SetBaseUrl gets a reference to the given string and assigns it to the BaseUrl field.
func (o *Service) SetBaseUrl(v string) {
	o.BaseUrl = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *Service) GetStatusMessage() string {
	if o == nil || o.StatusMessage == nil {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetStatusMessageOk() (*string, bool) {
	if o == nil || o.StatusMessage == nil {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *Service) HasStatusMessage() bool {
	if o != nil && o.StatusMessage != nil {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *Service) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetServiceDetail returns the ServiceDetail field value if set, zero value otherwise.
func (o *Service) GetServiceDetail() StatusResponse {
	if o == nil || o.ServiceDetail == nil {
		var ret StatusResponse
		return ret
	}
	return *o.ServiceDetail
}

// GetServiceDetailOk returns a tuple with the ServiceDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetServiceDetailOk() (*StatusResponse, bool) {
	if o == nil || o.ServiceDetail == nil {
		return nil, false
	}
	return o.ServiceDetail, true
}

// HasServiceDetail returns a boolean if a field has been set.
func (o *Service) HasServiceDetail() bool {
	if o != nil && o.ServiceDetail != nil {
		return true
	}

	return false
}

// SetServiceDetail gets a reference to the given StatusResponse and assigns it to the ServiceDetail field.
func (o *Service) SetServiceDetail(v StatusResponse) {
	o.ServiceDetail = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Service) GetStatus() bool {
	if o == nil || o.Status == nil {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetStatusOk() (*bool, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Service) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *Service) SetStatus(v bool) {
	o.Status = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Service) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Service) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Service) SetVersion(v string) {
	o.Version = &v
}

func (o Service) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hostid != nil {
		toSerialize["hostid"] = o.Hostid
	}
	if o.Servicename != nil {
		toSerialize["servicename"] = o.Servicename
	}
	if o.BaseUrl != nil {
		toSerialize["base_url"] = o.BaseUrl
	}
	if o.StatusMessage != nil {
		toSerialize["status_message"] = o.StatusMessage
	}
	if o.ServiceDetail != nil {
		toSerialize["service_detail"] = o.ServiceDetail
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableService struct {
	value *Service
	isSet bool
}

func (v NullableService) Get() *Service {
	return v.value
}

func (v *NullableService) Set(val *Service) {
	v.value = val
	v.isSet = true
}

func (v NullableService) IsSet() bool {
	return v.isSet
}

func (v *NullableService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableService(val *Service) *NullableService {
	return &NullableService{value: val, isSet: true}
}

func (v NullableService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


