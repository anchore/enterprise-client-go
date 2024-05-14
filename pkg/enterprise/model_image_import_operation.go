/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.5.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// ImageImportOperation An import record, creating a unique identifier for referencing the operation as well as its state
type ImageImportOperation struct {
	Uuid *string `json:"uuid,omitempty"`
	Status *string `json:"status,omitempty"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// NewImageImportOperation instantiates a new ImageImportOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageImportOperation() *ImageImportOperation {
	this := ImageImportOperation{}
	return &this
}

// NewImageImportOperationWithDefaults instantiates a new ImageImportOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageImportOperationWithDefaults() *ImageImportOperation {
	this := ImageImportOperation{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *ImageImportOperation) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageImportOperation) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *ImageImportOperation) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *ImageImportOperation) SetUuid(v string) {
	o.Uuid = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ImageImportOperation) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageImportOperation) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ImageImportOperation) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ImageImportOperation) SetStatus(v string) {
	o.Status = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *ImageImportOperation) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageImportOperation) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *ImageImportOperation) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *ImageImportOperation) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ImageImportOperation) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageImportOperation) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ImageImportOperation) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ImageImportOperation) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o ImageImportOperation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableImageImportOperation struct {
	value *ImageImportOperation
	isSet bool
}

func (v NullableImageImportOperation) Get() *ImageImportOperation {
	return v.value
}

func (v *NullableImageImportOperation) Set(val *ImageImportOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableImageImportOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableImageImportOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageImportOperation(val *ImageImportOperation) *NullableImageImportOperation {
	return &NullableImageImportOperation{value: val, isSet: true}
}

func (v NullableImageImportOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageImportOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


