/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.12.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// checks if the SourceImportOperation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceImportOperation{}

// SourceImportOperation An import record, creating a unique identifier for referencing the operation as well as its state
type SourceImportOperation struct {
	Uuid *string `json:"uuid,omitempty"`
	Status *string `json:"status,omitempty"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
}

// NewSourceImportOperation instantiates a new SourceImportOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceImportOperation() *SourceImportOperation {
	this := SourceImportOperation{}
	return &this
}

// NewSourceImportOperationWithDefaults instantiates a new SourceImportOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceImportOperationWithDefaults() *SourceImportOperation {
	this := SourceImportOperation{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *SourceImportOperation) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceImportOperation) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *SourceImportOperation) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *SourceImportOperation) SetUuid(v string) {
	o.Uuid = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SourceImportOperation) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceImportOperation) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SourceImportOperation) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SourceImportOperation) SetStatus(v string) {
	o.Status = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *SourceImportOperation) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceImportOperation) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *SourceImportOperation) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *SourceImportOperation) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SourceImportOperation) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceImportOperation) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SourceImportOperation) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SourceImportOperation) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *SourceImportOperation) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceImportOperation) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *SourceImportOperation) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *SourceImportOperation) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o SourceImportOperation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceImportOperation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["last_updated"] = o.LastUpdated
	}
	return toSerialize, nil
}

type NullableSourceImportOperation struct {
	value *SourceImportOperation
	isSet bool
}

func (v NullableSourceImportOperation) Get() *SourceImportOperation {
	return v.value
}

func (v *NullableSourceImportOperation) Set(val *SourceImportOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceImportOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceImportOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceImportOperation(val *SourceImportOperation) *NullableSourceImportOperation {
	return &NullableSourceImportOperation{value: val, isSet: true}
}

func (v NullableSourceImportOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceImportOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


