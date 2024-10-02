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
)

// IntegrationRegister Properties of an integration instance, e.g., an agent, to register
type IntegrationRegister struct {
	// identifier for the integration during registration until it has received its integration_id
	RegistrationId string `json:"registration_id"`
	// Unique identifier for the integration instance (among its replicas) during registration
	RegistrationInstanceId string `json:"registration_instance_id"`
	Type IntegrationType `json:"type"`
	// Name of the integration instance
	Name *string `json:"name,omitempty"`
	// Short description of the integration instance
	Description *string `json:"description,omitempty"`
	// Version of the integration instance
	Version *string `json:"version,omitempty"`
	// Timestamp when integration instance was started
	StartedAt *time.Time `json:"started_at,omitempty"`
	// Running time of integration instance in seconds
	Uptime *float32 `json:"uptime,omitempty"`
	// User that the integration instance authenticates as by default
	Username string `json:"username"`
	// List of accounts that the integration instance is statically configured to handle
	ExplicitlyAccountBound []string `json:"explicitly_account_bound,omitempty"`
	// List of namespaces that the integration instance is statically configured to handle
	Namespaces []string `json:"namespaces,omitempty"`
	// Configuration of the integration instance
	Configuration interface{} `json:"configuration,omitempty"`
	// Name of cluster where the integration instance runs
	ClusterName *string `json:"cluster_name,omitempty"`
	// Namespace in which the integration instance runs
	Namespace *string `json:"namespace,omitempty"`
	// Interval (in seconds) between health reports
	HealthReportInterval int32 `json:"health_report_interval"`
}

// NewIntegrationRegister instantiates a new IntegrationRegister object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationRegister(registrationId string, registrationInstanceId string, type_ IntegrationType, username string, healthReportInterval int32) *IntegrationRegister {
	this := IntegrationRegister{}
	this.RegistrationId = registrationId
	this.RegistrationInstanceId = registrationInstanceId
	this.Type = type_
	this.Username = username
	this.HealthReportInterval = healthReportInterval
	return &this
}

// NewIntegrationRegisterWithDefaults instantiates a new IntegrationRegister object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationRegisterWithDefaults() *IntegrationRegister {
	this := IntegrationRegister{}
	return &this
}

// GetRegistrationId returns the RegistrationId field value
func (o *IntegrationRegister) GetRegistrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegistrationId
}

// GetRegistrationIdOk returns a tuple with the RegistrationId field value
// and a boolean to check if the value has been set.
func (o *IntegrationRegister) GetRegistrationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegistrationId, true
}

// SetRegistrationId sets field value
func (o *IntegrationRegister) SetRegistrationId(v string) {
	o.RegistrationId = v
}

// GetRegistrationInstanceId returns the RegistrationInstanceId field value
func (o *IntegrationRegister) GetRegistrationInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegistrationInstanceId
}

// GetRegistrationInstanceIdOk returns a tuple with the RegistrationInstanceId field value
// and a boolean to check if the value has been set.
func (o *IntegrationRegister) GetRegistrationInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegistrationInstanceId, true
}

// SetRegistrationInstanceId sets field value
func (o *IntegrationRegister) SetRegistrationInstanceId(v string) {
	o.RegistrationInstanceId = v
}

// GetType returns the Type field value
func (o *IntegrationRegister) GetType() IntegrationType {
	if o == nil {
		var ret IntegrationType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IntegrationRegister) GetTypeOk() (*IntegrationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *IntegrationRegister) SetType(v IntegrationType) {
	o.Type = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IntegrationRegister) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationRegister) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IntegrationRegister) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IntegrationRegister) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IntegrationRegister) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationRegister) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IntegrationRegister) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IntegrationRegister) SetDescription(v string) {
	o.Description = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *IntegrationRegister) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationRegister) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *IntegrationRegister) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *IntegrationRegister) SetVersion(v string) {
	o.Version = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *IntegrationRegister) GetStartedAt() time.Time {
	if o == nil || o.StartedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationRegister) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || o.StartedAt == nil {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *IntegrationRegister) HasStartedAt() bool {
	if o != nil && o.StartedAt != nil {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *IntegrationRegister) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetUptime returns the Uptime field value if set, zero value otherwise.
func (o *IntegrationRegister) GetUptime() float32 {
	if o == nil || o.Uptime == nil {
		var ret float32
		return ret
	}
	return *o.Uptime
}

// GetUptimeOk returns a tuple with the Uptime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationRegister) GetUptimeOk() (*float32, bool) {
	if o == nil || o.Uptime == nil {
		return nil, false
	}
	return o.Uptime, true
}

// HasUptime returns a boolean if a field has been set.
func (o *IntegrationRegister) HasUptime() bool {
	if o != nil && o.Uptime != nil {
		return true
	}

	return false
}

// SetUptime gets a reference to the given float32 and assigns it to the Uptime field.
func (o *IntegrationRegister) SetUptime(v float32) {
	o.Uptime = &v
}

// GetUsername returns the Username field value
func (o *IntegrationRegister) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *IntegrationRegister) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *IntegrationRegister) SetUsername(v string) {
	o.Username = v
}

// GetExplicitlyAccountBound returns the ExplicitlyAccountBound field value if set, zero value otherwise.
func (o *IntegrationRegister) GetExplicitlyAccountBound() []string {
	if o == nil || o.ExplicitlyAccountBound == nil {
		var ret []string
		return ret
	}
	return o.ExplicitlyAccountBound
}

// GetExplicitlyAccountBoundOk returns a tuple with the ExplicitlyAccountBound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationRegister) GetExplicitlyAccountBoundOk() ([]string, bool) {
	if o == nil || o.ExplicitlyAccountBound == nil {
		return nil, false
	}
	return o.ExplicitlyAccountBound, true
}

// HasExplicitlyAccountBound returns a boolean if a field has been set.
func (o *IntegrationRegister) HasExplicitlyAccountBound() bool {
	if o != nil && o.ExplicitlyAccountBound != nil {
		return true
	}

	return false
}

// SetExplicitlyAccountBound gets a reference to the given []string and assigns it to the ExplicitlyAccountBound field.
func (o *IntegrationRegister) SetExplicitlyAccountBound(v []string) {
	o.ExplicitlyAccountBound = v
}

// GetNamespaces returns the Namespaces field value if set, zero value otherwise.
func (o *IntegrationRegister) GetNamespaces() []string {
	if o == nil || o.Namespaces == nil {
		var ret []string
		return ret
	}
	return o.Namespaces
}

// GetNamespacesOk returns a tuple with the Namespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationRegister) GetNamespacesOk() ([]string, bool) {
	if o == nil || o.Namespaces == nil {
		return nil, false
	}
	return o.Namespaces, true
}

// HasNamespaces returns a boolean if a field has been set.
func (o *IntegrationRegister) HasNamespaces() bool {
	if o != nil && o.Namespaces != nil {
		return true
	}

	return false
}

// SetNamespaces gets a reference to the given []string and assigns it to the Namespaces field.
func (o *IntegrationRegister) SetNamespaces(v []string) {
	o.Namespaces = v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *IntegrationRegister) GetConfiguration() interface{} {
	if o == nil || o.Configuration == nil {
		var ret interface{}
		return ret
	}
	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationRegister) GetConfigurationOk() (interface{}, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *IntegrationRegister) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given interface{} and assigns it to the Configuration field.
func (o *IntegrationRegister) SetConfiguration(v interface{}) {
	o.Configuration = v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *IntegrationRegister) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationRegister) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *IntegrationRegister) HasClusterName() bool {
	if o != nil && o.ClusterName != nil {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *IntegrationRegister) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *IntegrationRegister) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationRegister) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *IntegrationRegister) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *IntegrationRegister) SetNamespace(v string) {
	o.Namespace = &v
}

// GetHealthReportInterval returns the HealthReportInterval field value
func (o *IntegrationRegister) GetHealthReportInterval() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.HealthReportInterval
}

// GetHealthReportIntervalOk returns a tuple with the HealthReportInterval field value
// and a boolean to check if the value has been set.
func (o *IntegrationRegister) GetHealthReportIntervalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HealthReportInterval, true
}

// SetHealthReportInterval sets field value
func (o *IntegrationRegister) SetHealthReportInterval(v int32) {
	o.HealthReportInterval = v
}

func (o IntegrationRegister) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["registration_id"] = o.RegistrationId
	}
	if true {
		toSerialize["registration_instance_id"] = o.RegistrationInstanceId
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.StartedAt != nil {
		toSerialize["started_at"] = o.StartedAt
	}
	if o.Uptime != nil {
		toSerialize["uptime"] = o.Uptime
	}
	if true {
		toSerialize["username"] = o.Username
	}
	if o.ExplicitlyAccountBound != nil {
		toSerialize["explicitly_account_bound"] = o.ExplicitlyAccountBound
	}
	if o.Namespaces != nil {
		toSerialize["namespaces"] = o.Namespaces
	}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	if o.ClusterName != nil {
		toSerialize["cluster_name"] = o.ClusterName
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if true {
		toSerialize["health_report_interval"] = o.HealthReportInterval
	}
	return json.Marshal(toSerialize)
}

type NullableIntegrationRegister struct {
	value *IntegrationRegister
	isSet bool
}

func (v NullableIntegrationRegister) Get() *IntegrationRegister {
	return v.value
}

func (v *NullableIntegrationRegister) Set(val *IntegrationRegister) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationRegister) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationRegister) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationRegister(val *IntegrationRegister) *NullableIntegrationRegister {
	return &NullableIntegrationRegister{value: val, isSet: true}
}

func (v NullableIntegrationRegister) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationRegister) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


