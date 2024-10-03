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

// checks if the DeploymentHistory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploymentHistory{}

// DeploymentHistory struct for DeploymentHistory
type DeploymentHistory struct {
	DeploymentId *string `json:"deployment_id,omitempty"`
	UpgradeId *string `json:"upgrade_id,omitempty"`
	ToSystemVersion *string `json:"to_system_version,omitempty"`
	FromSystemVersion *string `json:"from_system_version,omitempty"`
	ToDatabaseVersion *string `json:"to_database_version,omitempty"`
	FromDatabaseVersion *string `json:"from_database_version,omitempty"`
	Outcome *string `json:"outcome,omitempty"`
	DbUpgradeDuration *float32 `json:"db_upgrade_duration,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// NewDeploymentHistory instantiates a new DeploymentHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentHistory() *DeploymentHistory {
	this := DeploymentHistory{}
	return &this
}

// NewDeploymentHistoryWithDefaults instantiates a new DeploymentHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentHistoryWithDefaults() *DeploymentHistory {
	this := DeploymentHistory{}
	return &this
}

// GetDeploymentId returns the DeploymentId field value if set, zero value otherwise.
func (o *DeploymentHistory) GetDeploymentId() string {
	if o == nil || IsNil(o.DeploymentId) {
		var ret string
		return ret
	}
	return *o.DeploymentId
}

// GetDeploymentIdOk returns a tuple with the DeploymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistory) GetDeploymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeploymentId) {
		return nil, false
	}
	return o.DeploymentId, true
}

// HasDeploymentId returns a boolean if a field has been set.
func (o *DeploymentHistory) HasDeploymentId() bool {
	if o != nil && !IsNil(o.DeploymentId) {
		return true
	}

	return false
}

// SetDeploymentId gets a reference to the given string and assigns it to the DeploymentId field.
func (o *DeploymentHistory) SetDeploymentId(v string) {
	o.DeploymentId = &v
}

// GetUpgradeId returns the UpgradeId field value if set, zero value otherwise.
func (o *DeploymentHistory) GetUpgradeId() string {
	if o == nil || IsNil(o.UpgradeId) {
		var ret string
		return ret
	}
	return *o.UpgradeId
}

// GetUpgradeIdOk returns a tuple with the UpgradeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistory) GetUpgradeIdOk() (*string, bool) {
	if o == nil || IsNil(o.UpgradeId) {
		return nil, false
	}
	return o.UpgradeId, true
}

// HasUpgradeId returns a boolean if a field has been set.
func (o *DeploymentHistory) HasUpgradeId() bool {
	if o != nil && !IsNil(o.UpgradeId) {
		return true
	}

	return false
}

// SetUpgradeId gets a reference to the given string and assigns it to the UpgradeId field.
func (o *DeploymentHistory) SetUpgradeId(v string) {
	o.UpgradeId = &v
}

// GetToSystemVersion returns the ToSystemVersion field value if set, zero value otherwise.
func (o *DeploymentHistory) GetToSystemVersion() string {
	if o == nil || IsNil(o.ToSystemVersion) {
		var ret string
		return ret
	}
	return *o.ToSystemVersion
}

// GetToSystemVersionOk returns a tuple with the ToSystemVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistory) GetToSystemVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ToSystemVersion) {
		return nil, false
	}
	return o.ToSystemVersion, true
}

// HasToSystemVersion returns a boolean if a field has been set.
func (o *DeploymentHistory) HasToSystemVersion() bool {
	if o != nil && !IsNil(o.ToSystemVersion) {
		return true
	}

	return false
}

// SetToSystemVersion gets a reference to the given string and assigns it to the ToSystemVersion field.
func (o *DeploymentHistory) SetToSystemVersion(v string) {
	o.ToSystemVersion = &v
}

// GetFromSystemVersion returns the FromSystemVersion field value if set, zero value otherwise.
func (o *DeploymentHistory) GetFromSystemVersion() string {
	if o == nil || IsNil(o.FromSystemVersion) {
		var ret string
		return ret
	}
	return *o.FromSystemVersion
}

// GetFromSystemVersionOk returns a tuple with the FromSystemVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistory) GetFromSystemVersionOk() (*string, bool) {
	if o == nil || IsNil(o.FromSystemVersion) {
		return nil, false
	}
	return o.FromSystemVersion, true
}

// HasFromSystemVersion returns a boolean if a field has been set.
func (o *DeploymentHistory) HasFromSystemVersion() bool {
	if o != nil && !IsNil(o.FromSystemVersion) {
		return true
	}

	return false
}

// SetFromSystemVersion gets a reference to the given string and assigns it to the FromSystemVersion field.
func (o *DeploymentHistory) SetFromSystemVersion(v string) {
	o.FromSystemVersion = &v
}

// GetToDatabaseVersion returns the ToDatabaseVersion field value if set, zero value otherwise.
func (o *DeploymentHistory) GetToDatabaseVersion() string {
	if o == nil || IsNil(o.ToDatabaseVersion) {
		var ret string
		return ret
	}
	return *o.ToDatabaseVersion
}

// GetToDatabaseVersionOk returns a tuple with the ToDatabaseVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistory) GetToDatabaseVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ToDatabaseVersion) {
		return nil, false
	}
	return o.ToDatabaseVersion, true
}

// HasToDatabaseVersion returns a boolean if a field has been set.
func (o *DeploymentHistory) HasToDatabaseVersion() bool {
	if o != nil && !IsNil(o.ToDatabaseVersion) {
		return true
	}

	return false
}

// SetToDatabaseVersion gets a reference to the given string and assigns it to the ToDatabaseVersion field.
func (o *DeploymentHistory) SetToDatabaseVersion(v string) {
	o.ToDatabaseVersion = &v
}

// GetFromDatabaseVersion returns the FromDatabaseVersion field value if set, zero value otherwise.
func (o *DeploymentHistory) GetFromDatabaseVersion() string {
	if o == nil || IsNil(o.FromDatabaseVersion) {
		var ret string
		return ret
	}
	return *o.FromDatabaseVersion
}

// GetFromDatabaseVersionOk returns a tuple with the FromDatabaseVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistory) GetFromDatabaseVersionOk() (*string, bool) {
	if o == nil || IsNil(o.FromDatabaseVersion) {
		return nil, false
	}
	return o.FromDatabaseVersion, true
}

// HasFromDatabaseVersion returns a boolean if a field has been set.
func (o *DeploymentHistory) HasFromDatabaseVersion() bool {
	if o != nil && !IsNil(o.FromDatabaseVersion) {
		return true
	}

	return false
}

// SetFromDatabaseVersion gets a reference to the given string and assigns it to the FromDatabaseVersion field.
func (o *DeploymentHistory) SetFromDatabaseVersion(v string) {
	o.FromDatabaseVersion = &v
}

// GetOutcome returns the Outcome field value if set, zero value otherwise.
func (o *DeploymentHistory) GetOutcome() string {
	if o == nil || IsNil(o.Outcome) {
		var ret string
		return ret
	}
	return *o.Outcome
}

// GetOutcomeOk returns a tuple with the Outcome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistory) GetOutcomeOk() (*string, bool) {
	if o == nil || IsNil(o.Outcome) {
		return nil, false
	}
	return o.Outcome, true
}

// HasOutcome returns a boolean if a field has been set.
func (o *DeploymentHistory) HasOutcome() bool {
	if o != nil && !IsNil(o.Outcome) {
		return true
	}

	return false
}

// SetOutcome gets a reference to the given string and assigns it to the Outcome field.
func (o *DeploymentHistory) SetOutcome(v string) {
	o.Outcome = &v
}

// GetDbUpgradeDuration returns the DbUpgradeDuration field value if set, zero value otherwise.
func (o *DeploymentHistory) GetDbUpgradeDuration() float32 {
	if o == nil || IsNil(o.DbUpgradeDuration) {
		var ret float32
		return ret
	}
	return *o.DbUpgradeDuration
}

// GetDbUpgradeDurationOk returns a tuple with the DbUpgradeDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistory) GetDbUpgradeDurationOk() (*float32, bool) {
	if o == nil || IsNil(o.DbUpgradeDuration) {
		return nil, false
	}
	return o.DbUpgradeDuration, true
}

// HasDbUpgradeDuration returns a boolean if a field has been set.
func (o *DeploymentHistory) HasDbUpgradeDuration() bool {
	if o != nil && !IsNil(o.DbUpgradeDuration) {
		return true
	}

	return false
}

// SetDbUpgradeDuration gets a reference to the given float32 and assigns it to the DbUpgradeDuration field.
func (o *DeploymentHistory) SetDbUpgradeDuration(v float32) {
	o.DbUpgradeDuration = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DeploymentHistory) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistory) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DeploymentHistory) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DeploymentHistory) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o DeploymentHistory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploymentHistory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeploymentId) {
		toSerialize["deployment_id"] = o.DeploymentId
	}
	if !IsNil(o.UpgradeId) {
		toSerialize["upgrade_id"] = o.UpgradeId
	}
	if !IsNil(o.ToSystemVersion) {
		toSerialize["to_system_version"] = o.ToSystemVersion
	}
	if !IsNil(o.FromSystemVersion) {
		toSerialize["from_system_version"] = o.FromSystemVersion
	}
	if !IsNil(o.ToDatabaseVersion) {
		toSerialize["to_database_version"] = o.ToDatabaseVersion
	}
	if !IsNil(o.FromDatabaseVersion) {
		toSerialize["from_database_version"] = o.FromDatabaseVersion
	}
	if !IsNil(o.Outcome) {
		toSerialize["outcome"] = o.Outcome
	}
	if !IsNil(o.DbUpgradeDuration) {
		toSerialize["db_upgrade_duration"] = o.DbUpgradeDuration
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	return toSerialize, nil
}

type NullableDeploymentHistory struct {
	value *DeploymentHistory
	isSet bool
}

func (v NullableDeploymentHistory) Get() *DeploymentHistory {
	return v.value
}

func (v *NullableDeploymentHistory) Set(val *DeploymentHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentHistory(val *DeploymentHistory) *NullableDeploymentHistory {
	return &NullableDeploymentHistory{value: val, isSet: true}
}

func (v NullableDeploymentHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


