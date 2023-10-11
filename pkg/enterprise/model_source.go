/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.0.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// checks if the Source type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Source{}

// Source struct for Source
type Source struct {
	// A system-assigned identifier unique for each source analysis
	Uuid *string `json:"uuid,omitempty"`
	// The anchore account name that owns this resource
	AccountName *string `json:"account_name,omitempty"`
	// Host name for the repository location (e.g. github.com)
	Host *string `json:"host,omitempty"`
	// The name of the repository on the host (e.g. 'anchore/anchore-engine')
	RepositoryName *string `json:"repository_name,omitempty"`
	// The commit ID for a git repository
	Revision *string `json:"revision,omitempty"`
	// The analysis state of the source
	AnalysisStatus *string `json:"analysis_status,omitempty"`
	// The state of the source
	SourceStatus *string `json:"source_status,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
}

// NewSource instantiates a new Source object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSource() *Source {
	this := Source{}
	return &this
}

// NewSourceWithDefaults instantiates a new Source object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceWithDefaults() *Source {
	this := Source{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *Source) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *Source) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *Source) SetUuid(v string) {
	o.Uuid = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *Source) GetAccountName() string {
	if o == nil || IsNil(o.AccountName) {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.AccountName) {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *Source) HasAccountName() bool {
	if o != nil && !IsNil(o.AccountName) {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *Source) SetAccountName(v string) {
	o.AccountName = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *Source) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *Source) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *Source) SetHost(v string) {
	o.Host = &v
}

// GetRepositoryName returns the RepositoryName field value if set, zero value otherwise.
func (o *Source) GetRepositoryName() string {
	if o == nil || IsNil(o.RepositoryName) {
		var ret string
		return ret
	}
	return *o.RepositoryName
}

// GetRepositoryNameOk returns a tuple with the RepositoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetRepositoryNameOk() (*string, bool) {
	if o == nil || IsNil(o.RepositoryName) {
		return nil, false
	}
	return o.RepositoryName, true
}

// HasRepositoryName returns a boolean if a field has been set.
func (o *Source) HasRepositoryName() bool {
	if o != nil && !IsNil(o.RepositoryName) {
		return true
	}

	return false
}

// SetRepositoryName gets a reference to the given string and assigns it to the RepositoryName field.
func (o *Source) SetRepositoryName(v string) {
	o.RepositoryName = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *Source) GetRevision() string {
	if o == nil || IsNil(o.Revision) {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetRevisionOk() (*string, bool) {
	if o == nil || IsNil(o.Revision) {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *Source) HasRevision() bool {
	if o != nil && !IsNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *Source) SetRevision(v string) {
	o.Revision = &v
}

// GetAnalysisStatus returns the AnalysisStatus field value if set, zero value otherwise.
func (o *Source) GetAnalysisStatus() string {
	if o == nil || IsNil(o.AnalysisStatus) {
		var ret string
		return ret
	}
	return *o.AnalysisStatus
}

// GetAnalysisStatusOk returns a tuple with the AnalysisStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetAnalysisStatusOk() (*string, bool) {
	if o == nil || IsNil(o.AnalysisStatus) {
		return nil, false
	}
	return o.AnalysisStatus, true
}

// HasAnalysisStatus returns a boolean if a field has been set.
func (o *Source) HasAnalysisStatus() bool {
	if o != nil && !IsNil(o.AnalysisStatus) {
		return true
	}

	return false
}

// SetAnalysisStatus gets a reference to the given string and assigns it to the AnalysisStatus field.
func (o *Source) SetAnalysisStatus(v string) {
	o.AnalysisStatus = &v
}

// GetSourceStatus returns the SourceStatus field value if set, zero value otherwise.
func (o *Source) GetSourceStatus() string {
	if o == nil || IsNil(o.SourceStatus) {
		var ret string
		return ret
	}
	return *o.SourceStatus
}

// GetSourceStatusOk returns a tuple with the SourceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetSourceStatusOk() (*string, bool) {
	if o == nil || IsNil(o.SourceStatus) {
		return nil, false
	}
	return o.SourceStatus, true
}

// HasSourceStatus returns a boolean if a field has been set.
func (o *Source) HasSourceStatus() bool {
	if o != nil && !IsNil(o.SourceStatus) {
		return true
	}

	return false
}

// SetSourceStatus gets a reference to the given string and assigns it to the SourceStatus field.
func (o *Source) SetSourceStatus(v string) {
	o.SourceStatus = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Source) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Source) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Source) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *Source) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *Source) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *Source) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o Source) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Source) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.AccountName) {
		toSerialize["account_name"] = o.AccountName
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.RepositoryName) {
		toSerialize["repository_name"] = o.RepositoryName
	}
	if !IsNil(o.Revision) {
		toSerialize["revision"] = o.Revision
	}
	if !IsNil(o.AnalysisStatus) {
		toSerialize["analysis_status"] = o.AnalysisStatus
	}
	if !IsNil(o.SourceStatus) {
		toSerialize["source_status"] = o.SourceStatus
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["last_updated"] = o.LastUpdated
	}
	return toSerialize, nil
}

type NullableSource struct {
	value *Source
	isSet bool
}

func (v NullableSource) Get() *Source {
	return v.value
}

func (v *NullableSource) Set(val *Source) {
	v.value = val
	v.isSet = true
}

func (v NullableSource) IsSet() bool {
	return v.isSet
}

func (v *NullableSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSource(val *Source) *NullableSource {
	return &NullableSource{value: val, isSet: true}
}

func (v NullableSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


