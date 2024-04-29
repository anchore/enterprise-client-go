/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.5.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// NotificationSMTPEndpointConfigurationAllOf struct for NotificationSMTPEndpointConfigurationAllOf
type NotificationSMTPEndpointConfigurationAllOf struct {
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

// NewNotificationSMTPEndpointConfigurationAllOf instantiates a new NotificationSMTPEndpointConfigurationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationSMTPEndpointConfigurationAllOf(host string, port int32, from string, to string) *NotificationSMTPEndpointConfigurationAllOf {
	this := NotificationSMTPEndpointConfigurationAllOf{}
	this.Host = host
	this.Port = port
	this.From = from
	this.To = to
	return &this
}

// NewNotificationSMTPEndpointConfigurationAllOfWithDefaults instantiates a new NotificationSMTPEndpointConfigurationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationSMTPEndpointConfigurationAllOfWithDefaults() *NotificationSMTPEndpointConfigurationAllOf {
	this := NotificationSMTPEndpointConfigurationAllOf{}
	return &this
}

// GetHost returns the Host field value
func (o *NotificationSMTPEndpointConfigurationAllOf) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *NotificationSMTPEndpointConfigurationAllOf) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *NotificationSMTPEndpointConfigurationAllOf) SetHost(v string) {
	o.Host = v
}

// GetPort returns the Port field value
func (o *NotificationSMTPEndpointConfigurationAllOf) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *NotificationSMTPEndpointConfigurationAllOf) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *NotificationSMTPEndpointConfigurationAllOf) SetPort(v int32) {
	o.Port = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *NotificationSMTPEndpointConfigurationAllOf) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSMTPEndpointConfigurationAllOf) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *NotificationSMTPEndpointConfigurationAllOf) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *NotificationSMTPEndpointConfigurationAllOf) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *NotificationSMTPEndpointConfigurationAllOf) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSMTPEndpointConfigurationAllOf) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *NotificationSMTPEndpointConfigurationAllOf) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *NotificationSMTPEndpointConfigurationAllOf) SetPassword(v string) {
	o.Password = &v
}

// GetUseTls returns the UseTls field value if set, zero value otherwise.
func (o *NotificationSMTPEndpointConfigurationAllOf) GetUseTls() bool {
	if o == nil || o.UseTls == nil {
		var ret bool
		return ret
	}
	return *o.UseTls
}

// GetUseTlsOk returns a tuple with the UseTls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSMTPEndpointConfigurationAllOf) GetUseTlsOk() (*bool, bool) {
	if o == nil || o.UseTls == nil {
		return nil, false
	}
	return o.UseTls, true
}

// HasUseTls returns a boolean if a field has been set.
func (o *NotificationSMTPEndpointConfigurationAllOf) HasUseTls() bool {
	if o != nil && o.UseTls != nil {
		return true
	}

	return false
}

// SetUseTls gets a reference to the given bool and assigns it to the UseTls field.
func (o *NotificationSMTPEndpointConfigurationAllOf) SetUseTls(v bool) {
	o.UseTls = &v
}

// GetFrom returns the From field value
func (o *NotificationSMTPEndpointConfigurationAllOf) GetFrom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *NotificationSMTPEndpointConfigurationAllOf) GetFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *NotificationSMTPEndpointConfigurationAllOf) SetFrom(v string) {
	o.From = v
}

// GetTo returns the To field value
func (o *NotificationSMTPEndpointConfigurationAllOf) GetTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *NotificationSMTPEndpointConfigurationAllOf) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *NotificationSMTPEndpointConfigurationAllOf) SetTo(v string) {
	o.To = v
}

func (o NotificationSMTPEndpointConfigurationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["host"] = o.Host
	}
	if true {
		toSerialize["port"] = o.Port
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.UseTls != nil {
		toSerialize["use_tls"] = o.UseTls
	}
	if true {
		toSerialize["from"] = o.From
	}
	if true {
		toSerialize["to"] = o.To
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationSMTPEndpointConfigurationAllOf struct {
	value *NotificationSMTPEndpointConfigurationAllOf
	isSet bool
}

func (v NullableNotificationSMTPEndpointConfigurationAllOf) Get() *NotificationSMTPEndpointConfigurationAllOf {
	return v.value
}

func (v *NullableNotificationSMTPEndpointConfigurationAllOf) Set(val *NotificationSMTPEndpointConfigurationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationSMTPEndpointConfigurationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationSMTPEndpointConfigurationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationSMTPEndpointConfigurationAllOf(val *NotificationSMTPEndpointConfigurationAllOf) *NullableNotificationSMTPEndpointConfigurationAllOf {
	return &NullableNotificationSMTPEndpointConfigurationAllOf{value: val, isSet: true}
}

func (v NullableNotificationSMTPEndpointConfigurationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationSMTPEndpointConfigurationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


