/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.2.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// NotificationTeamsEndpointConfiguration Configuration for MS Teams endpoint
type NotificationTeamsEndpointConfiguration struct {
	// The instance identifier for the configuration
	Uuid *string `json:"uuid,omitempty"`
	// User friendly name or description for the configuration
	Description *string `json:"description,omitempty"`
	// Verify the cert if using tls for connecting externally. Defaults to true if not specified
	VerifyTls *bool `json:"verify_tls,omitempty"`
	// Timestamp for last modification to the record
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Timestamp for last modification to the record
	LastUpdated *time.Time `json:"last_updated,omitempty"`
	// url to POST to, including any query parameters, should begin with 'http://' or 'https://'
	Url string `json:"url"`
}

// NewNotificationTeamsEndpointConfiguration instantiates a new NotificationTeamsEndpointConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationTeamsEndpointConfiguration(url string) *NotificationTeamsEndpointConfiguration {
	this := NotificationTeamsEndpointConfiguration{}
	this.Url = url
	return &this
}

// NewNotificationTeamsEndpointConfigurationWithDefaults instantiates a new NotificationTeamsEndpointConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationTeamsEndpointConfigurationWithDefaults() *NotificationTeamsEndpointConfiguration {
	this := NotificationTeamsEndpointConfiguration{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *NotificationTeamsEndpointConfiguration) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationTeamsEndpointConfiguration) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *NotificationTeamsEndpointConfiguration) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *NotificationTeamsEndpointConfiguration) SetUuid(v string) {
	o.Uuid = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NotificationTeamsEndpointConfiguration) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationTeamsEndpointConfiguration) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NotificationTeamsEndpointConfiguration) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NotificationTeamsEndpointConfiguration) SetDescription(v string) {
	o.Description = &v
}

// GetVerifyTls returns the VerifyTls field value if set, zero value otherwise.
func (o *NotificationTeamsEndpointConfiguration) GetVerifyTls() bool {
	if o == nil || o.VerifyTls == nil {
		var ret bool
		return ret
	}
	return *o.VerifyTls
}

// GetVerifyTlsOk returns a tuple with the VerifyTls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationTeamsEndpointConfiguration) GetVerifyTlsOk() (*bool, bool) {
	if o == nil || o.VerifyTls == nil {
		return nil, false
	}
	return o.VerifyTls, true
}

// HasVerifyTls returns a boolean if a field has been set.
func (o *NotificationTeamsEndpointConfiguration) HasVerifyTls() bool {
	if o != nil && o.VerifyTls != nil {
		return true
	}

	return false
}

// SetVerifyTls gets a reference to the given bool and assigns it to the VerifyTls field.
func (o *NotificationTeamsEndpointConfiguration) SetVerifyTls(v bool) {
	o.VerifyTls = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *NotificationTeamsEndpointConfiguration) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationTeamsEndpointConfiguration) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *NotificationTeamsEndpointConfiguration) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *NotificationTeamsEndpointConfiguration) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *NotificationTeamsEndpointConfiguration) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationTeamsEndpointConfiguration) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *NotificationTeamsEndpointConfiguration) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *NotificationTeamsEndpointConfiguration) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetUrl returns the Url field value
func (o *NotificationTeamsEndpointConfiguration) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NotificationTeamsEndpointConfiguration) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NotificationTeamsEndpointConfiguration) SetUrl(v string) {
	o.Url = v
}

func (o NotificationTeamsEndpointConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.VerifyTls != nil {
		toSerialize["verify_tls"] = o.VerifyTls
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if true {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationTeamsEndpointConfiguration struct {
	value *NotificationTeamsEndpointConfiguration
	isSet bool
}

func (v NullableNotificationTeamsEndpointConfiguration) Get() *NotificationTeamsEndpointConfiguration {
	return v.value
}

func (v *NullableNotificationTeamsEndpointConfiguration) Set(val *NotificationTeamsEndpointConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationTeamsEndpointConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationTeamsEndpointConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationTeamsEndpointConfiguration(val *NotificationTeamsEndpointConfiguration) *NullableNotificationTeamsEndpointConfiguration {
	return &NullableNotificationTeamsEndpointConfiguration{value: val, isSet: true}
}

func (v NullableNotificationTeamsEndpointConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationTeamsEndpointConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


