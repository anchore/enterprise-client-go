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

// SourceManifest struct for SourceManifest
type SourceManifest struct {
	Uuid *string `json:"uuid,omitempty"`
	AccountName *string `json:"account_name,omitempty"`
	VcsType *string `json:"vcs_type,omitempty"`
	Host *string `json:"host,omitempty"`
	RepositoryName *string `json:"repository_name,omitempty"`
	Revision *string `json:"revision,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
	AnalysisStatus *string `json:"analysis_status,omitempty"`
	SourceStatus *string `json:"source_status,omitempty"`
	// Array of metadata available
	MetadataRecords []SourceManifestMetadataRecord `json:"metadata_records,omitempty"`
}

// NewSourceManifest instantiates a new SourceManifest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceManifest() *SourceManifest {
	this := SourceManifest{}
	return &this
}

// NewSourceManifestWithDefaults instantiates a new SourceManifest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceManifestWithDefaults() *SourceManifest {
	this := SourceManifest{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *SourceManifest) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceManifest) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *SourceManifest) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *SourceManifest) SetUuid(v string) {
	o.Uuid = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *SourceManifest) GetAccountName() string {
	if o == nil || o.AccountName == nil {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceManifest) GetAccountNameOk() (*string, bool) {
	if o == nil || o.AccountName == nil {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *SourceManifest) HasAccountName() bool {
	if o != nil && o.AccountName != nil {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *SourceManifest) SetAccountName(v string) {
	o.AccountName = &v
}

// GetVcsType returns the VcsType field value if set, zero value otherwise.
func (o *SourceManifest) GetVcsType() string {
	if o == nil || o.VcsType == nil {
		var ret string
		return ret
	}
	return *o.VcsType
}

// GetVcsTypeOk returns a tuple with the VcsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceManifest) GetVcsTypeOk() (*string, bool) {
	if o == nil || o.VcsType == nil {
		return nil, false
	}
	return o.VcsType, true
}

// HasVcsType returns a boolean if a field has been set.
func (o *SourceManifest) HasVcsType() bool {
	if o != nil && o.VcsType != nil {
		return true
	}

	return false
}

// SetVcsType gets a reference to the given string and assigns it to the VcsType field.
func (o *SourceManifest) SetVcsType(v string) {
	o.VcsType = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *SourceManifest) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceManifest) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *SourceManifest) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *SourceManifest) SetHost(v string) {
	o.Host = &v
}

// GetRepositoryName returns the RepositoryName field value if set, zero value otherwise.
func (o *SourceManifest) GetRepositoryName() string {
	if o == nil || o.RepositoryName == nil {
		var ret string
		return ret
	}
	return *o.RepositoryName
}

// GetRepositoryNameOk returns a tuple with the RepositoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceManifest) GetRepositoryNameOk() (*string, bool) {
	if o == nil || o.RepositoryName == nil {
		return nil, false
	}
	return o.RepositoryName, true
}

// HasRepositoryName returns a boolean if a field has been set.
func (o *SourceManifest) HasRepositoryName() bool {
	if o != nil && o.RepositoryName != nil {
		return true
	}

	return false
}

// SetRepositoryName gets a reference to the given string and assigns it to the RepositoryName field.
func (o *SourceManifest) SetRepositoryName(v string) {
	o.RepositoryName = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *SourceManifest) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceManifest) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *SourceManifest) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *SourceManifest) SetRevision(v string) {
	o.Revision = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SourceManifest) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceManifest) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SourceManifest) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SourceManifest) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *SourceManifest) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceManifest) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *SourceManifest) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *SourceManifest) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetAnalysisStatus returns the AnalysisStatus field value if set, zero value otherwise.
func (o *SourceManifest) GetAnalysisStatus() string {
	if o == nil || o.AnalysisStatus == nil {
		var ret string
		return ret
	}
	return *o.AnalysisStatus
}

// GetAnalysisStatusOk returns a tuple with the AnalysisStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceManifest) GetAnalysisStatusOk() (*string, bool) {
	if o == nil || o.AnalysisStatus == nil {
		return nil, false
	}
	return o.AnalysisStatus, true
}

// HasAnalysisStatus returns a boolean if a field has been set.
func (o *SourceManifest) HasAnalysisStatus() bool {
	if o != nil && o.AnalysisStatus != nil {
		return true
	}

	return false
}

// SetAnalysisStatus gets a reference to the given string and assigns it to the AnalysisStatus field.
func (o *SourceManifest) SetAnalysisStatus(v string) {
	o.AnalysisStatus = &v
}

// GetSourceStatus returns the SourceStatus field value if set, zero value otherwise.
func (o *SourceManifest) GetSourceStatus() string {
	if o == nil || o.SourceStatus == nil {
		var ret string
		return ret
	}
	return *o.SourceStatus
}

// GetSourceStatusOk returns a tuple with the SourceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceManifest) GetSourceStatusOk() (*string, bool) {
	if o == nil || o.SourceStatus == nil {
		return nil, false
	}
	return o.SourceStatus, true
}

// HasSourceStatus returns a boolean if a field has been set.
func (o *SourceManifest) HasSourceStatus() bool {
	if o != nil && o.SourceStatus != nil {
		return true
	}

	return false
}

// SetSourceStatus gets a reference to the given string and assigns it to the SourceStatus field.
func (o *SourceManifest) SetSourceStatus(v string) {
	o.SourceStatus = &v
}

// GetMetadataRecords returns the MetadataRecords field value if set, zero value otherwise.
func (o *SourceManifest) GetMetadataRecords() []SourceManifestMetadataRecord {
	if o == nil || o.MetadataRecords == nil {
		var ret []SourceManifestMetadataRecord
		return ret
	}
	return o.MetadataRecords
}

// GetMetadataRecordsOk returns a tuple with the MetadataRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceManifest) GetMetadataRecordsOk() ([]SourceManifestMetadataRecord, bool) {
	if o == nil || o.MetadataRecords == nil {
		return nil, false
	}
	return o.MetadataRecords, true
}

// HasMetadataRecords returns a boolean if a field has been set.
func (o *SourceManifest) HasMetadataRecords() bool {
	if o != nil && o.MetadataRecords != nil {
		return true
	}

	return false
}

// SetMetadataRecords gets a reference to the given []SourceManifestMetadataRecord and assigns it to the MetadataRecords field.
func (o *SourceManifest) SetMetadataRecords(v []SourceManifestMetadataRecord) {
	o.MetadataRecords = v
}

func (o SourceManifest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.AccountName != nil {
		toSerialize["account_name"] = o.AccountName
	}
	if o.VcsType != nil {
		toSerialize["vcs_type"] = o.VcsType
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.RepositoryName != nil {
		toSerialize["repository_name"] = o.RepositoryName
	}
	if o.Revision != nil {
		toSerialize["revision"] = o.Revision
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if o.AnalysisStatus != nil {
		toSerialize["analysis_status"] = o.AnalysisStatus
	}
	if o.SourceStatus != nil {
		toSerialize["source_status"] = o.SourceStatus
	}
	if o.MetadataRecords != nil {
		toSerialize["metadata_records"] = o.MetadataRecords
	}
	return json.Marshal(toSerialize)
}

type NullableSourceManifest struct {
	value *SourceManifest
	isSet bool
}

func (v NullableSourceManifest) Get() *SourceManifest {
	return v.value
}

func (v *NullableSourceManifest) Set(val *SourceManifest) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceManifest) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceManifest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceManifest(val *SourceManifest) *NullableSourceManifest {
	return &NullableSourceManifest{value: val, isSet: true}
}

func (v NullableSourceManifest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceManifest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


