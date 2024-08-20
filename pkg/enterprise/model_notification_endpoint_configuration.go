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

// NotificationEndpointConfiguration Base schema that endpoint configurations extend
type NotificationEndpointConfiguration struct {
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
}

// NewNotificationEndpointConfiguration instantiates a new NotificationEndpointConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationEndpointConfiguration() *NotificationEndpointConfiguration {
	this := NotificationEndpointConfiguration{}
	return &this
}

// NewNotificationEndpointConfigurationWithDefaults instantiates a new NotificationEndpointConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationEndpointConfigurationWithDefaults() *NotificationEndpointConfiguration {
	this := NotificationEndpointConfiguration{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *NotificationEndpointConfiguration) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationEndpointConfiguration) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *NotificationEndpointConfiguration) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *NotificationEndpointConfiguration) SetUuid(v string) {
	o.Uuid = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NotificationEndpointConfiguration) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationEndpointConfiguration) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NotificationEndpointConfiguration) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NotificationEndpointConfiguration) SetDescription(v string) {
	o.Description = &v
}

// GetVerifyTls returns the VerifyTls field value if set, zero value otherwise.
func (o *NotificationEndpointConfiguration) GetVerifyTls() bool {
	if o == nil || o.VerifyTls == nil {
		var ret bool
		return ret
	}
	return *o.VerifyTls
}

// GetVerifyTlsOk returns a tuple with the VerifyTls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationEndpointConfiguration) GetVerifyTlsOk() (*bool, bool) {
	if o == nil || o.VerifyTls == nil {
		return nil, false
	}
	return o.VerifyTls, true
}

// HasVerifyTls returns a boolean if a field has been set.
func (o *NotificationEndpointConfiguration) HasVerifyTls() bool {
	if o != nil && o.VerifyTls != nil {
		return true
	}

	return false
}

// SetVerifyTls gets a reference to the given bool and assigns it to the VerifyTls field.
func (o *NotificationEndpointConfiguration) SetVerifyTls(v bool) {
	o.VerifyTls = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *NotificationEndpointConfiguration) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationEndpointConfiguration) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *NotificationEndpointConfiguration) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *NotificationEndpointConfiguration) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *NotificationEndpointConfiguration) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationEndpointConfiguration) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *NotificationEndpointConfiguration) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *NotificationEndpointConfiguration) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o NotificationEndpointConfiguration) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableNotificationEndpointConfiguration struct {
	value *NotificationEndpointConfiguration
	isSet bool
}

func (v NullableNotificationEndpointConfiguration) Get() *NotificationEndpointConfiguration {
	return v.value
}

func (v *NullableNotificationEndpointConfiguration) Set(val *NotificationEndpointConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationEndpointConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationEndpointConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationEndpointConfiguration(val *NotificationEndpointConfiguration) *NullableNotificationEndpointConfiguration {
	return &NullableNotificationEndpointConfiguration{value: val, isSet: true}
}

func (v NullableNotificationEndpointConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationEndpointConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


