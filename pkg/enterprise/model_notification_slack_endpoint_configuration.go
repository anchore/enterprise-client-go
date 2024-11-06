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

// checks if the NotificationSlackEndpointConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationSlackEndpointConfiguration{}

// NotificationSlackEndpointConfiguration Configuration for slack endpoint
type NotificationSlackEndpointConfiguration struct {
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
}

type _NotificationSlackEndpointConfiguration NotificationSlackEndpointConfiguration

// NewNotificationSlackEndpointConfiguration instantiates a new NotificationSlackEndpointConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationSlackEndpointConfiguration(url string) *NotificationSlackEndpointConfiguration {
	this := NotificationSlackEndpointConfiguration{}
	this.Url = url
	return &this
}

// NewNotificationSlackEndpointConfigurationWithDefaults instantiates a new NotificationSlackEndpointConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationSlackEndpointConfigurationWithDefaults() *NotificationSlackEndpointConfiguration {
	this := NotificationSlackEndpointConfiguration{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *NotificationSlackEndpointConfiguration) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSlackEndpointConfiguration) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *NotificationSlackEndpointConfiguration) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *NotificationSlackEndpointConfiguration) SetUuid(v string) {
	o.Uuid = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NotificationSlackEndpointConfiguration) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSlackEndpointConfiguration) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NotificationSlackEndpointConfiguration) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NotificationSlackEndpointConfiguration) SetDescription(v string) {
	o.Description = &v
}

// GetVerifyTls returns the VerifyTls field value if set, zero value otherwise.
func (o *NotificationSlackEndpointConfiguration) GetVerifyTls() bool {
	if o == nil || IsNil(o.VerifyTls) {
		var ret bool
		return ret
	}
	return *o.VerifyTls
}

// GetVerifyTlsOk returns a tuple with the VerifyTls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSlackEndpointConfiguration) GetVerifyTlsOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifyTls) {
		return nil, false
	}
	return o.VerifyTls, true
}

// HasVerifyTls returns a boolean if a field has been set.
func (o *NotificationSlackEndpointConfiguration) HasVerifyTls() bool {
	if o != nil && !IsNil(o.VerifyTls) {
		return true
	}

	return false
}

// SetVerifyTls gets a reference to the given bool and assigns it to the VerifyTls field.
func (o *NotificationSlackEndpointConfiguration) SetVerifyTls(v bool) {
	o.VerifyTls = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *NotificationSlackEndpointConfiguration) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSlackEndpointConfiguration) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *NotificationSlackEndpointConfiguration) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *NotificationSlackEndpointConfiguration) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *NotificationSlackEndpointConfiguration) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSlackEndpointConfiguration) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *NotificationSlackEndpointConfiguration) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *NotificationSlackEndpointConfiguration) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetUrl returns the Url field value
func (o *NotificationSlackEndpointConfiguration) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NotificationSlackEndpointConfiguration) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NotificationSlackEndpointConfiguration) SetUrl(v string) {
	o.Url = v
}

func (o NotificationSlackEndpointConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationSlackEndpointConfiguration) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

func (o *NotificationSlackEndpointConfiguration) UnmarshalJSON(data []byte) (err error) {
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

	varNotificationSlackEndpointConfiguration := _NotificationSlackEndpointConfiguration{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNotificationSlackEndpointConfiguration)

	if err != nil {
		return err
	}

	*o = NotificationSlackEndpointConfiguration(varNotificationSlackEndpointConfiguration)

	return err
}

type NullableNotificationSlackEndpointConfiguration struct {
	value *NotificationSlackEndpointConfiguration
	isSet bool
}

func (v NullableNotificationSlackEndpointConfiguration) Get() *NotificationSlackEndpointConfiguration {
	return v.value
}

func (v *NullableNotificationSlackEndpointConfiguration) Set(val *NotificationSlackEndpointConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationSlackEndpointConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationSlackEndpointConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationSlackEndpointConfiguration(val *NotificationSlackEndpointConfiguration) *NullableNotificationSlackEndpointConfiguration {
	return &NullableNotificationSlackEndpointConfiguration{value: val, isSet: true}
}

func (v NullableNotificationSlackEndpointConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationSlackEndpointConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


