/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.3.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// NotificationWebhookEndpointConfiguration Configuration for Webhook endpoint
type NotificationWebhookEndpointConfiguration struct {
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
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	// Verify SSL certificates for HTTPS requests, disabled by default
	VerifySsl *bool `json:"verify_ssl,omitempty"`
}

// NewNotificationWebhookEndpointConfiguration instantiates a new NotificationWebhookEndpointConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationWebhookEndpointConfiguration(url string) *NotificationWebhookEndpointConfiguration {
	this := NotificationWebhookEndpointConfiguration{}
	this.Url = url
	return &this
}

// NewNotificationWebhookEndpointConfigurationWithDefaults instantiates a new NotificationWebhookEndpointConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationWebhookEndpointConfigurationWithDefaults() *NotificationWebhookEndpointConfiguration {
	this := NotificationWebhookEndpointConfiguration{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *NotificationWebhookEndpointConfiguration) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationWebhookEndpointConfiguration) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *NotificationWebhookEndpointConfiguration) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *NotificationWebhookEndpointConfiguration) SetUuid(v string) {
	o.Uuid = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NotificationWebhookEndpointConfiguration) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationWebhookEndpointConfiguration) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NotificationWebhookEndpointConfiguration) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NotificationWebhookEndpointConfiguration) SetDescription(v string) {
	o.Description = &v
}

// GetVerifyTls returns the VerifyTls field value if set, zero value otherwise.
func (o *NotificationWebhookEndpointConfiguration) GetVerifyTls() bool {
	if o == nil || o.VerifyTls == nil {
		var ret bool
		return ret
	}
	return *o.VerifyTls
}

// GetVerifyTlsOk returns a tuple with the VerifyTls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationWebhookEndpointConfiguration) GetVerifyTlsOk() (*bool, bool) {
	if o == nil || o.VerifyTls == nil {
		return nil, false
	}
	return o.VerifyTls, true
}

// HasVerifyTls returns a boolean if a field has been set.
func (o *NotificationWebhookEndpointConfiguration) HasVerifyTls() bool {
	if o != nil && o.VerifyTls != nil {
		return true
	}

	return false
}

// SetVerifyTls gets a reference to the given bool and assigns it to the VerifyTls field.
func (o *NotificationWebhookEndpointConfiguration) SetVerifyTls(v bool) {
	o.VerifyTls = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *NotificationWebhookEndpointConfiguration) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationWebhookEndpointConfiguration) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *NotificationWebhookEndpointConfiguration) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *NotificationWebhookEndpointConfiguration) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *NotificationWebhookEndpointConfiguration) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationWebhookEndpointConfiguration) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *NotificationWebhookEndpointConfiguration) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *NotificationWebhookEndpointConfiguration) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetUrl returns the Url field value
func (o *NotificationWebhookEndpointConfiguration) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NotificationWebhookEndpointConfiguration) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NotificationWebhookEndpointConfiguration) SetUrl(v string) {
	o.Url = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *NotificationWebhookEndpointConfiguration) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationWebhookEndpointConfiguration) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *NotificationWebhookEndpointConfiguration) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *NotificationWebhookEndpointConfiguration) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *NotificationWebhookEndpointConfiguration) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationWebhookEndpointConfiguration) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *NotificationWebhookEndpointConfiguration) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *NotificationWebhookEndpointConfiguration) SetPassword(v string) {
	o.Password = &v
}

// GetVerifySsl returns the VerifySsl field value if set, zero value otherwise.
func (o *NotificationWebhookEndpointConfiguration) GetVerifySsl() bool {
	if o == nil || o.VerifySsl == nil {
		var ret bool
		return ret
	}
	return *o.VerifySsl
}

// GetVerifySslOk returns a tuple with the VerifySsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationWebhookEndpointConfiguration) GetVerifySslOk() (*bool, bool) {
	if o == nil || o.VerifySsl == nil {
		return nil, false
	}
	return o.VerifySsl, true
}

// HasVerifySsl returns a boolean if a field has been set.
func (o *NotificationWebhookEndpointConfiguration) HasVerifySsl() bool {
	if o != nil && o.VerifySsl != nil {
		return true
	}

	return false
}

// SetVerifySsl gets a reference to the given bool and assigns it to the VerifySsl field.
func (o *NotificationWebhookEndpointConfiguration) SetVerifySsl(v bool) {
	o.VerifySsl = &v
}

func (o NotificationWebhookEndpointConfiguration) MarshalJSON() ([]byte, error) {
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
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.VerifySsl != nil {
		toSerialize["verify_ssl"] = o.VerifySsl
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationWebhookEndpointConfiguration struct {
	value *NotificationWebhookEndpointConfiguration
	isSet bool
}

func (v NullableNotificationWebhookEndpointConfiguration) Get() *NotificationWebhookEndpointConfiguration {
	return v.value
}

func (v *NullableNotificationWebhookEndpointConfiguration) Set(val *NotificationWebhookEndpointConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationWebhookEndpointConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationWebhookEndpointConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationWebhookEndpointConfiguration(val *NotificationWebhookEndpointConfiguration) *NullableNotificationWebhookEndpointConfiguration {
	return &NullableNotificationWebhookEndpointConfiguration{value: val, isSet: true}
}

func (v NullableNotificationWebhookEndpointConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationWebhookEndpointConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


