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

// FeedDataRecord A data record
type FeedDataRecord struct {
	// The unique identifier for the data record
	Id *string `json:"id,omitempty"`
	// The name of the data record
	Name *string `json:"name,omitempty"`
	// The data set type of the record
	Dataset *string `json:"dataset,omitempty"`
	// The format of the data record
	DataFormat *string `json:"data_format,omitempty"`
	// The checksum of the data record
	Checksum *string `json:"checksum,omitempty"`
	// The build timestamp of the data record
	BuiltAt *time.Time `json:"built_at,omitempty"`
	// The creation timestamp of the data record
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The last update timestamp of the data record
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewFeedDataRecord instantiates a new FeedDataRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeedDataRecord() *FeedDataRecord {
	this := FeedDataRecord{}
	return &this
}

// NewFeedDataRecordWithDefaults instantiates a new FeedDataRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeedDataRecordWithDefaults() *FeedDataRecord {
	this := FeedDataRecord{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FeedDataRecord) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedDataRecord) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FeedDataRecord) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FeedDataRecord) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FeedDataRecord) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedDataRecord) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FeedDataRecord) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FeedDataRecord) SetName(v string) {
	o.Name = &v
}

// GetDataset returns the Dataset field value if set, zero value otherwise.
func (o *FeedDataRecord) GetDataset() string {
	if o == nil || o.Dataset == nil {
		var ret string
		return ret
	}
	return *o.Dataset
}

// GetDatasetOk returns a tuple with the Dataset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedDataRecord) GetDatasetOk() (*string, bool) {
	if o == nil || o.Dataset == nil {
		return nil, false
	}
	return o.Dataset, true
}

// HasDataset returns a boolean if a field has been set.
func (o *FeedDataRecord) HasDataset() bool {
	if o != nil && o.Dataset != nil {
		return true
	}

	return false
}

// SetDataset gets a reference to the given string and assigns it to the Dataset field.
func (o *FeedDataRecord) SetDataset(v string) {
	o.Dataset = &v
}

// GetDataFormat returns the DataFormat field value if set, zero value otherwise.
func (o *FeedDataRecord) GetDataFormat() string {
	if o == nil || o.DataFormat == nil {
		var ret string
		return ret
	}
	return *o.DataFormat
}

// GetDataFormatOk returns a tuple with the DataFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedDataRecord) GetDataFormatOk() (*string, bool) {
	if o == nil || o.DataFormat == nil {
		return nil, false
	}
	return o.DataFormat, true
}

// HasDataFormat returns a boolean if a field has been set.
func (o *FeedDataRecord) HasDataFormat() bool {
	if o != nil && o.DataFormat != nil {
		return true
	}

	return false
}

// SetDataFormat gets a reference to the given string and assigns it to the DataFormat field.
func (o *FeedDataRecord) SetDataFormat(v string) {
	o.DataFormat = &v
}

// GetChecksum returns the Checksum field value if set, zero value otherwise.
func (o *FeedDataRecord) GetChecksum() string {
	if o == nil || o.Checksum == nil {
		var ret string
		return ret
	}
	return *o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedDataRecord) GetChecksumOk() (*string, bool) {
	if o == nil || o.Checksum == nil {
		return nil, false
	}
	return o.Checksum, true
}

// HasChecksum returns a boolean if a field has been set.
func (o *FeedDataRecord) HasChecksum() bool {
	if o != nil && o.Checksum != nil {
		return true
	}

	return false
}

// SetChecksum gets a reference to the given string and assigns it to the Checksum field.
func (o *FeedDataRecord) SetChecksum(v string) {
	o.Checksum = &v
}

// GetBuiltAt returns the BuiltAt field value if set, zero value otherwise.
func (o *FeedDataRecord) GetBuiltAt() time.Time {
	if o == nil || o.BuiltAt == nil {
		var ret time.Time
		return ret
	}
	return *o.BuiltAt
}

// GetBuiltAtOk returns a tuple with the BuiltAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedDataRecord) GetBuiltAtOk() (*time.Time, bool) {
	if o == nil || o.BuiltAt == nil {
		return nil, false
	}
	return o.BuiltAt, true
}

// HasBuiltAt returns a boolean if a field has been set.
func (o *FeedDataRecord) HasBuiltAt() bool {
	if o != nil && o.BuiltAt != nil {
		return true
	}

	return false
}

// SetBuiltAt gets a reference to the given time.Time and assigns it to the BuiltAt field.
func (o *FeedDataRecord) SetBuiltAt(v time.Time) {
	o.BuiltAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FeedDataRecord) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedDataRecord) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FeedDataRecord) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *FeedDataRecord) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *FeedDataRecord) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedDataRecord) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *FeedDataRecord) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *FeedDataRecord) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o FeedDataRecord) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Dataset != nil {
		toSerialize["dataset"] = o.Dataset
	}
	if o.DataFormat != nil {
		toSerialize["data_format"] = o.DataFormat
	}
	if o.Checksum != nil {
		toSerialize["checksum"] = o.Checksum
	}
	if o.BuiltAt != nil {
		toSerialize["built_at"] = o.BuiltAt
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableFeedDataRecord struct {
	value *FeedDataRecord
	isSet bool
}

func (v NullableFeedDataRecord) Get() *FeedDataRecord {
	return v.value
}

func (v *NullableFeedDataRecord) Set(val *FeedDataRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedDataRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedDataRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedDataRecord(val *FeedDataRecord) *NullableFeedDataRecord {
	return &NullableFeedDataRecord{value: val, isSet: true}
}

func (v NullableFeedDataRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedDataRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


