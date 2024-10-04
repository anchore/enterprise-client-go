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
	"time"
	"bytes"
	"fmt"
)

// checks if the NotificationSMTPEndpointConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationSMTPEndpointConfiguration{}

// NotificationSMTPEndpointConfiguration Configuration for email via smtp endpoint
type NotificationSMTPEndpointConfiguration struct {
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
	Host string `json:"host"`
	Port int32 `json:"port"`
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	// Encrypt the SMTP connection with TLS. Defaults to true
	UseTls *bool `json:"use_tls,omitempty"`
	// The from address to use for emails send by this configuration
	From string `json:"from"`
	// The address to which the emails are sent
	To string `json:"to"`
}

type _NotificationSMTPEndpointConfiguration NotificationSMTPEndpointConfiguration

// NewNotificationSMTPEndpointConfiguration instantiates a new NotificationSMTPEndpointConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationSMTPEndpointConfiguration(host string, port int32, from string, to string) *NotificationSMTPEndpointConfiguration {
	this := NotificationSMTPEndpointConfiguration{}
	this.Host = host
	this.Port = port
	this.From = from
	this.To = to
	return &this
}

// NewNotificationSMTPEndpointConfigurationWithDefaults instantiates a new NotificationSMTPEndpointConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationSMTPEndpointConfigurationWithDefaults() *NotificationSMTPEndpointConfiguration {
	this := NotificationSMTPEndpointConfiguration{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *NotificationSMTPEndpointConfiguration) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSMTPEndpointConfiguration) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *NotificationSMTPEndpointConfiguration) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *NotificationSMTPEndpointConfiguration) SetUuid(v string) {
	o.Uuid = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NotificationSMTPEndpointConfiguration) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSMTPEndpointConfiguration) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NotificationSMTPEndpointConfiguration) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NotificationSMTPEndpointConfiguration) SetDescription(v string) {
	o.Description = &v
}

// GetVerifyTls returns the VerifyTls field value if set, zero value otherwise.
func (o *NotificationSMTPEndpointConfiguration) GetVerifyTls() bool {
	if o == nil || IsNil(o.VerifyTls) {
		var ret bool
		return ret
	}
	return *o.VerifyTls
}

// GetVerifyTlsOk returns a tuple with the VerifyTls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSMTPEndpointConfiguration) GetVerifyTlsOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifyTls) {
		return nil, false
	}
	return o.VerifyTls, true
}

// HasVerifyTls returns a boolean if a field has been set.
func (o *NotificationSMTPEndpointConfiguration) HasVerifyTls() bool {
	if o != nil && !IsNil(o.VerifyTls) {
		return true
	}

	return false
}

// SetVerifyTls gets a reference to the given bool and assigns it to the VerifyTls field.
func (o *NotificationSMTPEndpointConfiguration) SetVerifyTls(v bool) {
	o.VerifyTls = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *NotificationSMTPEndpointConfiguration) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSMTPEndpointConfiguration) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *NotificationSMTPEndpointConfiguration) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *NotificationSMTPEndpointConfiguration) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *NotificationSMTPEndpointConfiguration) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSMTPEndpointConfiguration) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *NotificationSMTPEndpointConfiguration) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *NotificationSMTPEndpointConfiguration) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetHost returns the Host field value
func (o *NotificationSMTPEndpointConfiguration) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *NotificationSMTPEndpointConfiguration) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *NotificationSMTPEndpointConfiguration) SetHost(v string) {
	o.Host = v
}

// GetPort returns the Port field value
func (o *NotificationSMTPEndpointConfiguration) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *NotificationSMTPEndpointConfiguration) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *NotificationSMTPEndpointConfiguration) SetPort(v int32) {
	o.Port = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *NotificationSMTPEndpointConfiguration) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSMTPEndpointConfiguration) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *NotificationSMTPEndpointConfiguration) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *NotificationSMTPEndpointConfiguration) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *NotificationSMTPEndpointConfiguration) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSMTPEndpointConfiguration) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *NotificationSMTPEndpointConfiguration) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *NotificationSMTPEndpointConfiguration) SetPassword(v string) {
	o.Password = &v
}

// GetUseTls returns the UseTls field value if set, zero value otherwise.
func (o *NotificationSMTPEndpointConfiguration) GetUseTls() bool {
	if o == nil || IsNil(o.UseTls) {
		var ret bool
		return ret
	}
	return *o.UseTls
}

// GetUseTlsOk returns a tuple with the UseTls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSMTPEndpointConfiguration) GetUseTlsOk() (*bool, bool) {
	if o == nil || IsNil(o.UseTls) {
		return nil, false
	}
	return o.UseTls, true
}

// HasUseTls returns a boolean if a field has been set.
func (o *NotificationSMTPEndpointConfiguration) HasUseTls() bool {
	if o != nil && !IsNil(o.UseTls) {
		return true
	}

	return false
}

// SetUseTls gets a reference to the given bool and assigns it to the UseTls field.
func (o *NotificationSMTPEndpointConfiguration) SetUseTls(v bool) {
	o.UseTls = &v
}

// GetFrom returns the From field value
func (o *NotificationSMTPEndpointConfiguration) GetFrom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *NotificationSMTPEndpointConfiguration) GetFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *NotificationSMTPEndpointConfiguration) SetFrom(v string) {
	o.From = v
}

// GetTo returns the To field value
func (o *NotificationSMTPEndpointConfiguration) GetTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *NotificationSMTPEndpointConfiguration) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *NotificationSMTPEndpointConfiguration) SetTo(v string) {
	o.To = v
}

func (o NotificationSMTPEndpointConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationSMTPEndpointConfiguration) ToMap() (map[string]interface{}, error) {
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
	toSerialize["host"] = o.Host
	toSerialize["port"] = o.Port
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.UseTls) {
		toSerialize["use_tls"] = o.UseTls
	}
	toSerialize["from"] = o.From
	toSerialize["to"] = o.To
	return toSerialize, nil
}

func (o *NotificationSMTPEndpointConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"host",
		"port",
		"from",
		"to",
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

	varNotificationSMTPEndpointConfiguration := _NotificationSMTPEndpointConfiguration{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNotificationSMTPEndpointConfiguration)

	if err != nil {
		return err
	}

	*o = NotificationSMTPEndpointConfiguration(varNotificationSMTPEndpointConfiguration)

	return err
}

type NullableNotificationSMTPEndpointConfiguration struct {
	value *NotificationSMTPEndpointConfiguration
	isSet bool
}

func (v NullableNotificationSMTPEndpointConfiguration) Get() *NotificationSMTPEndpointConfiguration {
	return v.value
}

func (v *NullableNotificationSMTPEndpointConfiguration) Set(val *NotificationSMTPEndpointConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationSMTPEndpointConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationSMTPEndpointConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationSMTPEndpointConfiguration(val *NotificationSMTPEndpointConfiguration) *NullableNotificationSMTPEndpointConfiguration {
	return &NullableNotificationSMTPEndpointConfiguration{value: val, isSet: true}
}

func (v NullableNotificationSMTPEndpointConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationSMTPEndpointConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


