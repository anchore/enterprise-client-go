/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.10.1
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// checks if the IntegrationAdditionalProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationAdditionalProperties{}

// IntegrationAdditionalProperties struct for IntegrationAdditionalProperties
type IntegrationAdditionalProperties struct {
	// Short description of the integration instance
	Description *string `json:"description,omitempty"`
	// Timestamp when integration instance was started
	StartedAt *time.Time `json:"started_at,omitempty"`
	// List of namespaces handled by the integration instance
	Namespaces []string `json:"namespaces,omitempty"`
	// Configuration of the integration instance
	Configuration interface{} `json:"configuration,omitempty"`
}

// NewIntegrationAdditionalProperties instantiates a new IntegrationAdditionalProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationAdditionalProperties() *IntegrationAdditionalProperties {
	this := IntegrationAdditionalProperties{}
	return &this
}

// NewIntegrationAdditionalPropertiesWithDefaults instantiates a new IntegrationAdditionalProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationAdditionalPropertiesWithDefaults() *IntegrationAdditionalProperties {
	this := IntegrationAdditionalProperties{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IntegrationAdditionalProperties) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationAdditionalProperties) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IntegrationAdditionalProperties) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IntegrationAdditionalProperties) SetDescription(v string) {
	o.Description = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *IntegrationAdditionalProperties) GetStartedAt() time.Time {
	if o == nil || IsNil(o.StartedAt) {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationAdditionalProperties) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartedAt) {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *IntegrationAdditionalProperties) HasStartedAt() bool {
	if o != nil && !IsNil(o.StartedAt) {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *IntegrationAdditionalProperties) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetNamespaces returns the Namespaces field value if set, zero value otherwise.
func (o *IntegrationAdditionalProperties) GetNamespaces() []string {
	if o == nil || IsNil(o.Namespaces) {
		var ret []string
		return ret
	}
	return o.Namespaces
}

// GetNamespacesOk returns a tuple with the Namespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationAdditionalProperties) GetNamespacesOk() ([]string, bool) {
	if o == nil || IsNil(o.Namespaces) {
		return nil, false
	}
	return o.Namespaces, true
}

// HasNamespaces returns a boolean if a field has been set.
func (o *IntegrationAdditionalProperties) HasNamespaces() bool {
	if o != nil && !IsNil(o.Namespaces) {
		return true
	}

	return false
}

// SetNamespaces gets a reference to the given []string and assigns it to the Namespaces field.
func (o *IntegrationAdditionalProperties) SetNamespaces(v []string) {
	o.Namespaces = v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *IntegrationAdditionalProperties) GetConfiguration() interface{} {
	if o == nil || IsNil(o.Configuration) {
		var ret interface{}
		return ret
	}
	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationAdditionalProperties) GetConfigurationOk() (interface{}, bool) {
	if o == nil || IsNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *IntegrationAdditionalProperties) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given interface{} and assigns it to the Configuration field.
func (o *IntegrationAdditionalProperties) SetConfiguration(v interface{}) {
	o.Configuration = v
}

func (o IntegrationAdditionalProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationAdditionalProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.StartedAt) {
		toSerialize["started_at"] = o.StartedAt
	}
	if !IsNil(o.Namespaces) {
		toSerialize["namespaces"] = o.Namespaces
	}
	if !IsNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	return toSerialize, nil
}

type NullableIntegrationAdditionalProperties struct {
	value *IntegrationAdditionalProperties
	isSet bool
}

func (v NullableIntegrationAdditionalProperties) Get() *IntegrationAdditionalProperties {
	return v.value
}

func (v *NullableIntegrationAdditionalProperties) Set(val *IntegrationAdditionalProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationAdditionalProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationAdditionalProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationAdditionalProperties(val *IntegrationAdditionalProperties) *NullableIntegrationAdditionalProperties {
	return &NullableIntegrationAdditionalProperties{value: val, isSet: true}
}

func (v NullableIntegrationAdditionalProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationAdditionalProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


