/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.10.2
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// checks if the FeedMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeedMetadata{}

// FeedMetadata Metadata on the feeds based on findings from querying the endpoints.
type FeedMetadata struct {
	// name of the feed
	Name *string `json:"name,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The last time the policy-engine service pinged the feed service to see if there was a new grypedb available.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The build timestamp of the feed
	Built *time.Time `json:"built,omitempty"`
	Groups []FeedGroupMetadata `json:"groups,omitempty"`
	// The last time that policy-engine service downloaded a new grypedb.
	LastFullSync *time.Time `json:"last_full_sync,omitempty"`
	// If feed is enabled
	Enabled *bool `json:"enabled,omitempty"`
	// The name of the dataset that provides this feed
	DatasetName *string `json:"dataset_name,omitempty"`
	// The checksum of the dataset that provides this feed
	DatasetChecksum *string `json:"dataset_checksum,omitempty"`
}

// NewFeedMetadata instantiates a new FeedMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeedMetadata() *FeedMetadata {
	this := FeedMetadata{}
	return &this
}

// NewFeedMetadataWithDefaults instantiates a new FeedMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeedMetadataWithDefaults() *FeedMetadata {
	this := FeedMetadata{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FeedMetadata) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedMetadata) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FeedMetadata) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FeedMetadata) SetName(v string) {
	o.Name = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FeedMetadata) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedMetadata) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FeedMetadata) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *FeedMetadata) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *FeedMetadata) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedMetadata) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *FeedMetadata) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *FeedMetadata) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetBuilt returns the Built field value if set, zero value otherwise.
func (o *FeedMetadata) GetBuilt() time.Time {
	if o == nil || IsNil(o.Built) {
		var ret time.Time
		return ret
	}
	return *o.Built
}

// GetBuiltOk returns a tuple with the Built field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedMetadata) GetBuiltOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Built) {
		return nil, false
	}
	return o.Built, true
}

// HasBuilt returns a boolean if a field has been set.
func (o *FeedMetadata) HasBuilt() bool {
	if o != nil && !IsNil(o.Built) {
		return true
	}

	return false
}

// SetBuilt gets a reference to the given time.Time and assigns it to the Built field.
func (o *FeedMetadata) SetBuilt(v time.Time) {
	o.Built = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *FeedMetadata) GetGroups() []FeedGroupMetadata {
	if o == nil || IsNil(o.Groups) {
		var ret []FeedGroupMetadata
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedMetadata) GetGroupsOk() ([]FeedGroupMetadata, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *FeedMetadata) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []FeedGroupMetadata and assigns it to the Groups field.
func (o *FeedMetadata) SetGroups(v []FeedGroupMetadata) {
	o.Groups = v
}

// GetLastFullSync returns the LastFullSync field value if set, zero value otherwise.
func (o *FeedMetadata) GetLastFullSync() time.Time {
	if o == nil || IsNil(o.LastFullSync) {
		var ret time.Time
		return ret
	}
	return *o.LastFullSync
}

// GetLastFullSyncOk returns a tuple with the LastFullSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedMetadata) GetLastFullSyncOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastFullSync) {
		return nil, false
	}
	return o.LastFullSync, true
}

// HasLastFullSync returns a boolean if a field has been set.
func (o *FeedMetadata) HasLastFullSync() bool {
	if o != nil && !IsNil(o.LastFullSync) {
		return true
	}

	return false
}

// SetLastFullSync gets a reference to the given time.Time and assigns it to the LastFullSync field.
func (o *FeedMetadata) SetLastFullSync(v time.Time) {
	o.LastFullSync = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *FeedMetadata) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedMetadata) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *FeedMetadata) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *FeedMetadata) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDatasetName returns the DatasetName field value if set, zero value otherwise.
func (o *FeedMetadata) GetDatasetName() string {
	if o == nil || IsNil(o.DatasetName) {
		var ret string
		return ret
	}
	return *o.DatasetName
}

// GetDatasetNameOk returns a tuple with the DatasetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedMetadata) GetDatasetNameOk() (*string, bool) {
	if o == nil || IsNil(o.DatasetName) {
		return nil, false
	}
	return o.DatasetName, true
}

// HasDatasetName returns a boolean if a field has been set.
func (o *FeedMetadata) HasDatasetName() bool {
	if o != nil && !IsNil(o.DatasetName) {
		return true
	}

	return false
}

// SetDatasetName gets a reference to the given string and assigns it to the DatasetName field.
func (o *FeedMetadata) SetDatasetName(v string) {
	o.DatasetName = &v
}

// GetDatasetChecksum returns the DatasetChecksum field value if set, zero value otherwise.
func (o *FeedMetadata) GetDatasetChecksum() string {
	if o == nil || IsNil(o.DatasetChecksum) {
		var ret string
		return ret
	}
	return *o.DatasetChecksum
}

// GetDatasetChecksumOk returns a tuple with the DatasetChecksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedMetadata) GetDatasetChecksumOk() (*string, bool) {
	if o == nil || IsNil(o.DatasetChecksum) {
		return nil, false
	}
	return o.DatasetChecksum, true
}

// HasDatasetChecksum returns a boolean if a field has been set.
func (o *FeedMetadata) HasDatasetChecksum() bool {
	if o != nil && !IsNil(o.DatasetChecksum) {
		return true
	}

	return false
}

// SetDatasetChecksum gets a reference to the given string and assigns it to the DatasetChecksum field.
func (o *FeedMetadata) SetDatasetChecksum(v string) {
	o.DatasetChecksum = &v
}

func (o FeedMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeedMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Built) {
		toSerialize["built"] = o.Built
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.LastFullSync) {
		toSerialize["last_full_sync"] = o.LastFullSync
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.DatasetName) {
		toSerialize["dataset_name"] = o.DatasetName
	}
	if !IsNil(o.DatasetChecksum) {
		toSerialize["dataset_checksum"] = o.DatasetChecksum
	}
	return toSerialize, nil
}

type NullableFeedMetadata struct {
	value *FeedMetadata
	isSet bool
}

func (v NullableFeedMetadata) Get() *FeedMetadata {
	return v.value
}

func (v *NullableFeedMetadata) Set(val *FeedMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedMetadata(val *FeedMetadata) *NullableFeedMetadata {
	return &NullableFeedMetadata{value: val, isSet: true}
}

func (v NullableFeedMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


