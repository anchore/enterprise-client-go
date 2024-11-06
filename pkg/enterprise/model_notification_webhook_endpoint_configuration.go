/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.9.0
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

// checks if the NotificationWebhookEndpointConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationWebhookEndpointConfiguration{}

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
	Url string `json:"url" validate:"regexp=https?:\\/\\/.*"`
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	// Verify SSL certificates for HTTPS requests, disabled by default
	VerifySsl *bool `json:"verify_ssl,omitempty"`
}

type _NotificationWebhookEndpointConfiguration NotificationWebhookEndpointConfiguration

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
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationWebhookEndpointConfiguration) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *NotificationWebhookEndpointConfiguration) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
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
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationWebhookEndpointConfiguration) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NotificationWebhookEndpointConfiguration) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
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
	if o == nil || IsNil(o.VerifyTls) {
		var ret bool
		return ret
	}
	return *o.VerifyTls
}

// GetVerifyTlsOk returns a tuple with the VerifyTls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationWebhookEndpointConfiguration) GetVerifyTlsOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifyTls) {
		return nil, false
	}
	return o.VerifyTls, true
}

// HasVerifyTls returns a boolean if a field has been set.
func (o *NotificationWebhookEndpointConfiguration) HasVerifyTls() bool {
	if o != nil && !IsNil(o.VerifyTls) {
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
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationWebhookEndpointConfiguration) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *NotificationWebhookEndpointConfiguration) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
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
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationWebhookEndpointConfiguration) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *NotificationWebhookEndpointConfiguration) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
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
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationWebhookEndpointConfiguration) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *NotificationWebhookEndpointConfiguration) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
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
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationWebhookEndpointConfiguration) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *NotificationWebhookEndpointConfiguration) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
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
	if o == nil || IsNil(o.VerifySsl) {
		var ret bool
		return ret
	}
	return *o.VerifySsl
}

// GetVerifySslOk returns a tuple with the VerifySsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationWebhookEndpointConfiguration) GetVerifySslOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifySsl) {
		return nil, false
	}
	return o.VerifySsl, true
}

// HasVerifySsl returns a boolean if a field has been set.
func (o *NotificationWebhookEndpointConfiguration) HasVerifySsl() bool {
	if o != nil && !IsNil(o.VerifySsl) {
		return true
	}

	return false
}

// SetVerifySsl gets a reference to the given bool and assigns it to the VerifySsl field.
func (o *NotificationWebhookEndpointConfiguration) SetVerifySsl(v bool) {
	o.VerifySsl = &v
}

func (o NotificationWebhookEndpointConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationWebhookEndpointConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.VerifyTls) {
		toSerialize["verify_tls"] = o.VerifyTls
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["last_updated"] = o.LastUpdated
	}
	toSerialize["url"] = o.Url
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.VerifySsl) {
		toSerialize["verify_ssl"] = o.VerifySsl
	}
	return toSerialize, nil
}

func (o *NotificationWebhookEndpointConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
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

	varNotificationWebhookEndpointConfiguration := _NotificationWebhookEndpointConfiguration{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNotificationWebhookEndpointConfiguration)

	if err != nil {
		return err
	}

	*o = NotificationWebhookEndpointConfiguration(varNotificationWebhookEndpointConfiguration)

	return err
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


