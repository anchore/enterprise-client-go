/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.7.2
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the Service type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Service{}

// Service A service status record
type Service struct {
	// The unique id of the host on which the service is executing
	HostId *string `json:"host_id,omitempty"`
	// Registered service name
	ServiceName *string `json:"service_name,omitempty"`
	// The url to reach the service, including port as needed
	BaseUrl *string `json:"base_url,omitempty"`
	// A state indicating the condition of the service. Normal operation is 'registered'
	StatusMessage *string `json:"status_message,omitempty"`
	ServiceDetail *StatusResponse `json:"service_detail,omitempty"`
	Status *bool `json:"status,omitempty"`
	// The version of the service as reported by the service implementation on registration
	Version *string `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Service Service

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

// GetHostId returns the HostId field value if set, zero value otherwise.
func (o *Service) GetHostId() string {
	if o == nil || IsNil(o.HostId) {
		var ret string
		return ret
	}
	return *o.HostId
}

// GetHostIdOk returns a tuple with the HostId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetHostIdOk() (*string, bool) {
	if o == nil || IsNil(o.HostId) {
		return nil, false
	}
	return o.HostId, true
}

// HasHostId returns a boolean if a field has been set.
func (o *Service) HasHostId() bool {
	if o != nil && !IsNil(o.HostId) {
		return true
	}

	return false
}

// SetHostId gets a reference to the given string and assigns it to the HostId field.
func (o *Service) SetHostId(v string) {
	o.HostId = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *Service) GetServiceName() string {
	if o == nil || IsNil(o.ServiceName) {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceName) {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *Service) HasServiceName() bool {
	if o != nil && !IsNil(o.ServiceName) {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *Service) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetBaseUrl returns the BaseUrl field value if set, zero value otherwise.
func (o *Service) GetBaseUrl() string {
	if o == nil || IsNil(o.BaseUrl) {
		var ret string
		return ret
	}
	return *o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetBaseUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BaseUrl) {
		return nil, false
	}
	return o.BaseUrl, true
}

// HasBaseUrl returns a boolean if a field has been set.
func (o *Service) HasBaseUrl() bool {
	if o != nil && !IsNil(o.BaseUrl) {
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
	if o == nil || IsNil(o.StatusMessage) {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetStatusMessageOk() (*string, bool) {
	if o == nil || IsNil(o.StatusMessage) {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *Service) HasStatusMessage() bool {
	if o != nil && !IsNil(o.StatusMessage) {
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
	if o == nil || IsNil(o.ServiceDetail) {
		var ret StatusResponse
		return ret
	}
	return *o.ServiceDetail
}

// GetServiceDetailOk returns a tuple with the ServiceDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetServiceDetailOk() (*StatusResponse, bool) {
	if o == nil || IsNil(o.ServiceDetail) {
		return nil, false
	}
	return o.ServiceDetail, true
}

// HasServiceDetail returns a boolean if a field has been set.
func (o *Service) HasServiceDetail() bool {
	if o != nil && !IsNil(o.ServiceDetail) {
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
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Service) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
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
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Service) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Service) SetVersion(v string) {
	o.Version = &v
}

func (o Service) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Service) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HostId) {
		toSerialize["host_id"] = o.HostId
	}
	if !IsNil(o.ServiceName) {
		toSerialize["service_name"] = o.ServiceName
	}
	if !IsNil(o.BaseUrl) {
		toSerialize["base_url"] = o.BaseUrl
	}
	if !IsNil(o.StatusMessage) {
		toSerialize["status_message"] = o.StatusMessage
	}
	if !IsNil(o.ServiceDetail) {
		toSerialize["service_detail"] = o.ServiceDetail
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Service) UnmarshalJSON(data []byte) (err error) {
	varService := _Service{}

	err = json.Unmarshal(data, &varService)

	if err != nil {
		return err
	}

	*o = Service(varService)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "host_id")
		delete(additionalProperties, "service_name")
		delete(additionalProperties, "base_url")
		delete(additionalProperties, "status_message")
		delete(additionalProperties, "service_detail")
		delete(additionalProperties, "status")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
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


