/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.8.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// Integration Information about an integration instance, e.g., an agent
type Integration struct {
	// Globally unique identifier for the integration instance
	Uuid string `json:"uuid"`
	Type IntegrationType `json:"type"`
	// Name of the integration instance
	Name *string `json:"name,omitempty"`
	// Version of the integration instance
	Version *string `json:"version,omitempty"`
	ReportedStatus *IntegrationReportedStatus `json:"reported_status,omitempty"`
	IntegrationStatus *IntegrationStatus `json:"integration_status,omitempty"`
	// Timestamp of last received health report from integration instance
	LastSeen *time.Time `json:"last_seen,omitempty"`
	// Running time of integration instance in seconds
	Uptime *float32 `json:"uptime,omitempty"`
	// Account that the integration instance reports to by default
	AccountName string `json:"account_name"`
	// User that the integration instance authenticates as by default
	Username string `json:"username"`
	// List of accounts that the integration instance is statically configured to handle
	ExplicitlyAccountBound []string `json:"explicitly_account_bound,omitempty"`
	// List of accounts that the integration instance handles
	Accounts []string `json:"accounts,omitempty"`
	// Name of cluster where the integration instance runs
	ClusterName *string `json:"cluster_name,omitempty"`
	// Namespace in which the integration instance runs
	Namespace *string `json:"namespace,omitempty"`
	// Interval (in seconds) between health reports
	HealthReportInterval *int32 `json:"health_report_interval,omitempty"`
	// identifier for the integration during registration until it has received its integration_id
	RegistrationId *string `json:"registration_id,omitempty"`
	// Unique identifier for the integration instance (among its replicas) during registration
	RegistrationInstanceId *string `json:"registration_instance_id,omitempty"`
	// Short description of the integration instance
	Description *string `json:"description,omitempty"`
	// Timestamp when integration instance was started
	StartedAt *time.Time `json:"started_at,omitempty"`
	// List of namespaces handled by the integration instance
	Namespaces []string `json:"namespaces,omitempty"`
	// Configuration of the integration instance
	Configuration interface{} `json:"configuration,omitempty"`
}

// NewIntegration instantiates a new Integration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegration(uuid string, type_ IntegrationType, accountName string, username string) *Integration {
	this := Integration{}
	this.Uuid = uuid
	this.Type = type_
	this.AccountName = accountName
	this.Username = username
	return &this
}

// NewIntegrationWithDefaults instantiates a new Integration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationWithDefaults() *Integration {
	this := Integration{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *Integration) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *Integration) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *Integration) SetUuid(v string) {
	o.Uuid = v
}

// GetType returns the Type field value
func (o *Integration) GetType() IntegrationType {
	if o == nil {
		var ret IntegrationType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Integration) GetTypeOk() (*IntegrationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Integration) SetType(v IntegrationType) {
	o.Type = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Integration) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Integration) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Integration) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Integration) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Integration) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Integration) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Integration) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Integration) SetVersion(v string) {
	o.Version = &v
}

// GetReportedStatus returns the ReportedStatus field value if set, zero value otherwise.
func (o *Integration) GetReportedStatus() IntegrationReportedStatus {
	if o == nil || o.ReportedStatus == nil {
		var ret IntegrationReportedStatus
		return ret
	}
	return *o.ReportedStatus
}

// GetReportedStatusOk returns a tuple with the ReportedStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Integration) GetReportedStatusOk() (*IntegrationReportedStatus, bool) {
	if o == nil || o.ReportedStatus == nil {
		return nil, false
	}
	return o.ReportedStatus, true
}

// HasReportedStatus returns a boolean if a field has been set.
func (o *Integration) HasReportedStatus() bool {
	if o != nil && o.ReportedStatus != nil {
		return true
	}

	return false
}

// SetReportedStatus gets a reference to the given IntegrationReportedStatus and assigns it to the ReportedStatus field.
func (o *Integration) SetReportedStatus(v IntegrationReportedStatus) {
	o.ReportedStatus = &v
}

// GetIntegrationStatus returns the IntegrationStatus field value if set, zero value otherwise.
func (o *Integration) GetIntegrationStatus() IntegrationStatus {
	if o == nil || o.IntegrationStatus == nil {
		var ret IntegrationStatus
		return ret
	}
	return *o.IntegrationStatus
}

// GetIntegrationStatusOk returns a tuple with the IntegrationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Integration) GetIntegrationStatusOk() (*IntegrationStatus, bool) {
	if o == nil || o.IntegrationStatus == nil {
		return nil, false
	}
	return o.IntegrationStatus, true
}

// HasIntegrationStatus returns a boolean if a field has been set.
func (o *Integration) HasIntegrationStatus() bool {
	if o != nil && o.IntegrationStatus != nil {
		return true
	}

	return false
}

// SetIntegrationStatus gets a reference to the given IntegrationStatus and assigns it to the IntegrationStatus field.
func (o *Integration) SetIntegrationStatus(v IntegrationStatus) {
	o.IntegrationStatus = &v
}

// GetLastSeen returns the LastSeen field value if set, zero value otherwise.
func (o *Integration) GetLastSeen() time.Time {
	if o == nil || o.LastSeen == nil {
		var ret time.Time
		return ret
	}
	return *o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Integration) GetLastSeenOk() (*time.Time, bool) {
	if o == nil || o.LastSeen == nil {
		return nil, false
	}
	return o.LastSeen, true
}

// HasLastSeen returns a boolean if a field has been set.
func (o *Integration) HasLastSeen() bool {
	if o != nil && o.LastSeen != nil {
		return true
	}

	return false
}

// SetLastSeen gets a reference to the given time.Time and assigns it to the LastSeen field.
func (o *Integration) SetLastSeen(v time.Time) {
	o.LastSeen = &v
}

// GetUptime returns the Uptime field value if set, zero value otherwise.
func (o *Integration) GetUptime() float32 {
	if o == nil || o.Uptime == nil {
		var ret float32
		return ret
	}
	return *o.Uptime
}

// GetUptimeOk returns a tuple with the Uptime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Integration) GetUptimeOk() (*float32, bool) {
	if o == nil || o.Uptime == nil {
		return nil, false
	}
	return o.Uptime, true
}

// HasUptime returns a boolean if a field has been set.
func (o *Integration) HasUptime() bool {
	if o != nil && o.Uptime != nil {
		return true
	}

	return false
}

// SetUptime gets a reference to the given float32 and assigns it to the Uptime field.
func (o *Integration) SetUptime(v float32) {
	o.Uptime = &v
}

// GetAccountName returns the AccountName field value
func (o *Integration) GetAccountName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value
// and a boolean to check if the value has been set.
func (o *Integration) GetAccountNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountName, true
}

// SetAccountName sets field value
func (o *Integration) SetAccountName(v string) {
	o.AccountName = v
}

// GetUsername returns the Username field value
func (o *Integration) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *Integration) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *Integration) SetUsername(v string) {
	o.Username = v
}

// GetExplicitlyAccountBound returns the ExplicitlyAccountBound field value if set, zero value otherwise.
func (o *Integration) GetExplicitlyAccountBound() []string {
	if o == nil || o.ExplicitlyAccountBound == nil {
		var ret []string
		return ret
	}
	return o.ExplicitlyAccountBound
}

// GetExplicitlyAccountBoundOk returns a tuple with the ExplicitlyAccountBound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Integration) GetExplicitlyAccountBoundOk() ([]string, bool) {
	if o == nil || o.ExplicitlyAccountBound == nil {
		return nil, false
	}
	return o.ExplicitlyAccountBound, true
}

// HasExplicitlyAccountBound returns a boolean if a field has been set.
func (o *Integration) HasExplicitlyAccountBound() bool {
	if o != nil && o.ExplicitlyAccountBound != nil {
		return true
	}

	return false
}

// SetExplicitlyAccountBound gets a reference to the given []string and assigns it to the ExplicitlyAccountBound field.
func (o *Integration) SetExplicitlyAccountBound(v []string) {
	o.ExplicitlyAccountBound = v
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *Integration) GetAccounts() []string {
	if o == nil || o.Accounts == nil {
		var ret []string
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Integration) GetAccountsOk() ([]string, bool) {
	if o == nil || o.Accounts == nil {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *Integration) HasAccounts() bool {
	if o != nil && o.Accounts != nil {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []string and assigns it to the Accounts field.
func (o *Integration) SetAccounts(v []string) {
	o.Accounts = v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *Integration) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Integration) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *Integration) HasClusterName() bool {
	if o != nil && o.ClusterName != nil {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *Integration) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *Integration) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Integration) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *Integration) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *Integration) SetNamespace(v string) {
	o.Namespace = &v
}

// GetHealthReportInterval returns the HealthReportInterval field value if set, zero value otherwise.
func (o *Integration) GetHealthReportInterval() int32 {
	if o == nil || o.HealthReportInterval == nil {
		var ret int32
		return ret
	}
	return *o.HealthReportInterval
}

// GetHealthReportIntervalOk returns a tuple with the HealthReportInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Integration) GetHealthReportIntervalOk() (*int32, bool) {
	if o == nil || o.HealthReportInterval == nil {
		return nil, false
	}
	return o.HealthReportInterval, true
}

// HasHealthReportInterval returns a boolean if a field has been set.
func (o *Integration) HasHealthReportInterval() bool {
	if o != nil && o.HealthReportInterval != nil {
		return true
	}

	return false
}

// SetHealthReportInterval gets a reference to the given int32 and assigns it to the HealthReportInterval field.
func (o *Integration) SetHealthReportInterval(v int32) {
	o.HealthReportInterval = &v
}

// GetRegistrationId returns the RegistrationId field value if set, zero value otherwise.
func (o *Integration) GetRegistrationId() string {
	if o == nil || o.RegistrationId == nil {
		var ret string
		return ret
	}
	return *o.RegistrationId
}

// GetRegistrationIdOk returns a tuple with the RegistrationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Integration) GetRegistrationIdOk() (*string, bool) {
	if o == nil || o.RegistrationId == nil {
		return nil, false
	}
	return o.RegistrationId, true
}

// HasRegistrationId returns a boolean if a field has been set.
func (o *Integration) HasRegistrationId() bool {
	if o != nil && o.RegistrationId != nil {
		return true
	}

	return false
}

// SetRegistrationId gets a reference to the given string and assigns it to the RegistrationId field.
func (o *Integration) SetRegistrationId(v string) {
	o.RegistrationId = &v
}

// GetRegistrationInstanceId returns the RegistrationInstanceId field value if set, zero value otherwise.
func (o *Integration) GetRegistrationInstanceId() string {
	if o == nil || o.RegistrationInstanceId == nil {
		var ret string
		return ret
	}
	return *o.RegistrationInstanceId
}

// GetRegistrationInstanceIdOk returns a tuple with the RegistrationInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Integration) GetRegistrationInstanceIdOk() (*string, bool) {
	if o == nil || o.RegistrationInstanceId == nil {
		return nil, false
	}
	return o.RegistrationInstanceId, true
}

// HasRegistrationInstanceId returns a boolean if a field has been set.
func (o *Integration) HasRegistrationInstanceId() bool {
	if o != nil && o.RegistrationInstanceId != nil {
		return true
	}

	return false
}

// SetRegistrationInstanceId gets a reference to the given string and assigns it to the RegistrationInstanceId field.
func (o *Integration) SetRegistrationInstanceId(v string) {
	o.RegistrationInstanceId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Integration) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Integration) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Integration) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Integration) SetDescription(v string) {
	o.Description = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *Integration) GetStartedAt() time.Time {
	if o == nil || o.StartedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Integration) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || o.StartedAt == nil {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *Integration) HasStartedAt() bool {
	if o != nil && o.StartedAt != nil {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *Integration) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetNamespaces returns the Namespaces field value if set, zero value otherwise.
func (o *Integration) GetNamespaces() []string {
	if o == nil || o.Namespaces == nil {
		var ret []string
		return ret
	}
	return o.Namespaces
}

// GetNamespacesOk returns a tuple with the Namespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Integration) GetNamespacesOk() ([]string, bool) {
	if o == nil || o.Namespaces == nil {
		return nil, false
	}
	return o.Namespaces, true
}

// HasNamespaces returns a boolean if a field has been set.
func (o *Integration) HasNamespaces() bool {
	if o != nil && o.Namespaces != nil {
		return true
	}

	return false
}

// SetNamespaces gets a reference to the given []string and assigns it to the Namespaces field.
func (o *Integration) SetNamespaces(v []string) {
	o.Namespaces = v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *Integration) GetConfiguration() interface{} {
	if o == nil || o.Configuration == nil {
		var ret interface{}
		return ret
	}
	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Integration) GetConfigurationOk() (interface{}, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *Integration) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given interface{} and assigns it to the Configuration field.
func (o *Integration) SetConfiguration(v interface{}) {
	o.Configuration = v
}

func (o Integration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.ReportedStatus != nil {
		toSerialize["reported_status"] = o.ReportedStatus
	}
	if o.IntegrationStatus != nil {
		toSerialize["integration_status"] = o.IntegrationStatus
	}
	if o.LastSeen != nil {
		toSerialize["last_seen"] = o.LastSeen
	}
	if o.Uptime != nil {
		toSerialize["uptime"] = o.Uptime
	}
	if true {
		toSerialize["account_name"] = o.AccountName
	}
	if true {
		toSerialize["username"] = o.Username
	}
	if o.ExplicitlyAccountBound != nil {
		toSerialize["explicitly_account_bound"] = o.ExplicitlyAccountBound
	}
	if o.Accounts != nil {
		toSerialize["accounts"] = o.Accounts
	}
	if o.ClusterName != nil {
		toSerialize["cluster_name"] = o.ClusterName
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.HealthReportInterval != nil {
		toSerialize["health_report_interval"] = o.HealthReportInterval
	}
	if o.RegistrationId != nil {
		toSerialize["registration_id"] = o.RegistrationId
	}
	if o.RegistrationInstanceId != nil {
		toSerialize["registration_instance_id"] = o.RegistrationInstanceId
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.StartedAt != nil {
		toSerialize["started_at"] = o.StartedAt
	}
	if o.Namespaces != nil {
		toSerialize["namespaces"] = o.Namespaces
	}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	return json.Marshal(toSerialize)
}

type NullableIntegration struct {
	value *Integration
	isSet bool
}

func (v NullableIntegration) Get() *Integration {
	return v.value
}

func (v *NullableIntegration) Set(val *Integration) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegration(val *Integration) *NullableIntegration {
	return &NullableIntegration{value: val, isSet: true}
}

func (v NullableIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


