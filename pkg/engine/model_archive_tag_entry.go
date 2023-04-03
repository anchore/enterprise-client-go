/*
Anchore Engine API Server

This is the Anchore Engine API. Provides the primary external API for users of the service.

API version: 0.6.0
Contact: nurmi@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package engine

import (
	"encoding/json"
	"time"
)

// ArchiveTagEntry A docker-pullable tag value as well as deconstructed components
type ArchiveTagEntry struct {
	// The pullable string for the tag. E.g. \"docker.io/library/node:latest\"
	Pullstring *string `json:"pullstring,omitempty"`
	// The registry hostname:port section of the pull string
	Registry *string `json:"registry,omitempty"`
	// The repository section of the pull string
	Repository *string `json:"repository,omitempty"`
	// The tag-only section of the pull string
	Tag *string `json:"tag,omitempty"`
	// The timestamp at which the Anchore Engine detected this tag was mapped to the image digest. Does not necessarily indicate when the tag was actually pushed to the registry.
	DetectedAt *time.Time `json:"detected_at,omitempty"`
	// The timestamp at which Anchore Engine archived this image digest.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The timestamp that the last change was made to this record.
	LastUpdated *time.Time `json:"last_updated,omitempty"`
}

// NewArchiveTagEntry instantiates a new ArchiveTagEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArchiveTagEntry() *ArchiveTagEntry {
	this := ArchiveTagEntry{}
	return &this
}

// NewArchiveTagEntryWithDefaults instantiates a new ArchiveTagEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArchiveTagEntryWithDefaults() *ArchiveTagEntry {
	this := ArchiveTagEntry{}
	return &this
}

// GetPullstring returns the Pullstring field value if set, zero value otherwise.
func (o *ArchiveTagEntry) GetPullstring() string {
	if o == nil || o.Pullstring == nil {
		var ret string
		return ret
	}
	return *o.Pullstring
}

// GetPullstringOk returns a tuple with the Pullstring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveTagEntry) GetPullstringOk() (*string, bool) {
	if o == nil || o.Pullstring == nil {
		return nil, false
	}
	return o.Pullstring, true
}

// HasPullstring returns a boolean if a field has been set.
func (o *ArchiveTagEntry) HasPullstring() bool {
	if o != nil && o.Pullstring != nil {
		return true
	}

	return false
}

// SetPullstring gets a reference to the given string and assigns it to the Pullstring field.
func (o *ArchiveTagEntry) SetPullstring(v string) {
	o.Pullstring = &v
}

// GetRegistry returns the Registry field value if set, zero value otherwise.
func (o *ArchiveTagEntry) GetRegistry() string {
	if o == nil || o.Registry == nil {
		var ret string
		return ret
	}
	return *o.Registry
}

// GetRegistryOk returns a tuple with the Registry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveTagEntry) GetRegistryOk() (*string, bool) {
	if o == nil || o.Registry == nil {
		return nil, false
	}
	return o.Registry, true
}

// HasRegistry returns a boolean if a field has been set.
func (o *ArchiveTagEntry) HasRegistry() bool {
	if o != nil && o.Registry != nil {
		return true
	}

	return false
}

// SetRegistry gets a reference to the given string and assigns it to the Registry field.
func (o *ArchiveTagEntry) SetRegistry(v string) {
	o.Registry = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *ArchiveTagEntry) GetRepository() string {
	if o == nil || o.Repository == nil {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveTagEntry) GetRepositoryOk() (*string, bool) {
	if o == nil || o.Repository == nil {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *ArchiveTagEntry) HasRepository() bool {
	if o != nil && o.Repository != nil {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *ArchiveTagEntry) SetRepository(v string) {
	o.Repository = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *ArchiveTagEntry) GetTag() string {
	if o == nil || o.Tag == nil {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveTagEntry) GetTagOk() (*string, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *ArchiveTagEntry) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *ArchiveTagEntry) SetTag(v string) {
	o.Tag = &v
}

// GetDetectedAt returns the DetectedAt field value if set, zero value otherwise.
func (o *ArchiveTagEntry) GetDetectedAt() time.Time {
	if o == nil || o.DetectedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.DetectedAt
}

// GetDetectedAtOk returns a tuple with the DetectedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveTagEntry) GetDetectedAtOk() (*time.Time, bool) {
	if o == nil || o.DetectedAt == nil {
		return nil, false
	}
	return o.DetectedAt, true
}

// HasDetectedAt returns a boolean if a field has been set.
func (o *ArchiveTagEntry) HasDetectedAt() bool {
	if o != nil && o.DetectedAt != nil {
		return true
	}

	return false
}

// SetDetectedAt gets a reference to the given time.Time and assigns it to the DetectedAt field.
func (o *ArchiveTagEntry) SetDetectedAt(v time.Time) {
	o.DetectedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ArchiveTagEntry) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveTagEntry) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ArchiveTagEntry) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ArchiveTagEntry) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *ArchiveTagEntry) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveTagEntry) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *ArchiveTagEntry) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *ArchiveTagEntry) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o ArchiveTagEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Pullstring != nil {
		toSerialize["pullstring"] = o.Pullstring
	}
	if o.Registry != nil {
		toSerialize["registry"] = o.Registry
	}
	if o.Repository != nil {
		toSerialize["repository"] = o.Repository
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	if o.DetectedAt != nil {
		toSerialize["detected_at"] = o.DetectedAt
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	return json.Marshal(toSerialize)
}

type NullableArchiveTagEntry struct {
	value *ArchiveTagEntry
	isSet bool
}

func (v NullableArchiveTagEntry) Get() *ArchiveTagEntry {
	return v.value
}

func (v *NullableArchiveTagEntry) Set(val *ArchiveTagEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableArchiveTagEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableArchiveTagEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArchiveTagEntry(val *ArchiveTagEntry) *NullableArchiveTagEntry {
	return &NullableArchiveTagEntry{value: val, isSet: true}
}

func (v NullableArchiveTagEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArchiveTagEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


