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

// SourceManifestMetadataRecord Metadata associated with a source upload
type SourceManifestMetadataRecord struct {
	Uuid *string `json:"uuid,omitempty"`
	CiWorkflowName NullableString `json:"ci_workflow_name,omitempty"`
	CiWorkflowExecutionTime NullableTime `json:"ci_workflow_execution_time,omitempty"`
	BranchName NullableString `json:"branch_name,omitempty"`
	ChangeAuthor NullableString `json:"change_author,omitempty"`
}

// NewSourceManifestMetadataRecord instantiates a new SourceManifestMetadataRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceManifestMetadataRecord() *SourceManifestMetadataRecord {
	this := SourceManifestMetadataRecord{}
	return &this
}

// NewSourceManifestMetadataRecordWithDefaults instantiates a new SourceManifestMetadataRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceManifestMetadataRecordWithDefaults() *SourceManifestMetadataRecord {
	this := SourceManifestMetadataRecord{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *SourceManifestMetadataRecord) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceManifestMetadataRecord) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *SourceManifestMetadataRecord) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *SourceManifestMetadataRecord) SetUuid(v string) {
	o.Uuid = &v
}

// GetCiWorkflowName returns the CiWorkflowName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceManifestMetadataRecord) GetCiWorkflowName() string {
	if o == nil || o.CiWorkflowName.Get() == nil {
		var ret string
		return ret
	}
	return *o.CiWorkflowName.Get()
}

// GetCiWorkflowNameOk returns a tuple with the CiWorkflowName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceManifestMetadataRecord) GetCiWorkflowNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CiWorkflowName.Get(), o.CiWorkflowName.IsSet()
}

// HasCiWorkflowName returns a boolean if a field has been set.
func (o *SourceManifestMetadataRecord) HasCiWorkflowName() bool {
	if o != nil && o.CiWorkflowName.IsSet() {
		return true
	}

	return false
}

// SetCiWorkflowName gets a reference to the given NullableString and assigns it to the CiWorkflowName field.
func (o *SourceManifestMetadataRecord) SetCiWorkflowName(v string) {
	o.CiWorkflowName.Set(&v)
}
// SetCiWorkflowNameNil sets the value for CiWorkflowName to be an explicit nil
func (o *SourceManifestMetadataRecord) SetCiWorkflowNameNil() {
	o.CiWorkflowName.Set(nil)
}

// UnsetCiWorkflowName ensures that no value is present for CiWorkflowName, not even an explicit nil
func (o *SourceManifestMetadataRecord) UnsetCiWorkflowName() {
	o.CiWorkflowName.Unset()
}

// GetCiWorkflowExecutionTime returns the CiWorkflowExecutionTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceManifestMetadataRecord) GetCiWorkflowExecutionTime() time.Time {
	if o == nil || o.CiWorkflowExecutionTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CiWorkflowExecutionTime.Get()
}

// GetCiWorkflowExecutionTimeOk returns a tuple with the CiWorkflowExecutionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceManifestMetadataRecord) GetCiWorkflowExecutionTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CiWorkflowExecutionTime.Get(), o.CiWorkflowExecutionTime.IsSet()
}

// HasCiWorkflowExecutionTime returns a boolean if a field has been set.
func (o *SourceManifestMetadataRecord) HasCiWorkflowExecutionTime() bool {
	if o != nil && o.CiWorkflowExecutionTime.IsSet() {
		return true
	}

	return false
}

// SetCiWorkflowExecutionTime gets a reference to the given NullableTime and assigns it to the CiWorkflowExecutionTime field.
func (o *SourceManifestMetadataRecord) SetCiWorkflowExecutionTime(v time.Time) {
	o.CiWorkflowExecutionTime.Set(&v)
}
// SetCiWorkflowExecutionTimeNil sets the value for CiWorkflowExecutionTime to be an explicit nil
func (o *SourceManifestMetadataRecord) SetCiWorkflowExecutionTimeNil() {
	o.CiWorkflowExecutionTime.Set(nil)
}

// UnsetCiWorkflowExecutionTime ensures that no value is present for CiWorkflowExecutionTime, not even an explicit nil
func (o *SourceManifestMetadataRecord) UnsetCiWorkflowExecutionTime() {
	o.CiWorkflowExecutionTime.Unset()
}

// GetBranchName returns the BranchName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceManifestMetadataRecord) GetBranchName() string {
	if o == nil || o.BranchName.Get() == nil {
		var ret string
		return ret
	}
	return *o.BranchName.Get()
}

// GetBranchNameOk returns a tuple with the BranchName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceManifestMetadataRecord) GetBranchNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BranchName.Get(), o.BranchName.IsSet()
}

// HasBranchName returns a boolean if a field has been set.
func (o *SourceManifestMetadataRecord) HasBranchName() bool {
	if o != nil && o.BranchName.IsSet() {
		return true
	}

	return false
}

// SetBranchName gets a reference to the given NullableString and assigns it to the BranchName field.
func (o *SourceManifestMetadataRecord) SetBranchName(v string) {
	o.BranchName.Set(&v)
}
// SetBranchNameNil sets the value for BranchName to be an explicit nil
func (o *SourceManifestMetadataRecord) SetBranchNameNil() {
	o.BranchName.Set(nil)
}

// UnsetBranchName ensures that no value is present for BranchName, not even an explicit nil
func (o *SourceManifestMetadataRecord) UnsetBranchName() {
	o.BranchName.Unset()
}

// GetChangeAuthor returns the ChangeAuthor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceManifestMetadataRecord) GetChangeAuthor() string {
	if o == nil || o.ChangeAuthor.Get() == nil {
		var ret string
		return ret
	}
	return *o.ChangeAuthor.Get()
}

// GetChangeAuthorOk returns a tuple with the ChangeAuthor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceManifestMetadataRecord) GetChangeAuthorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChangeAuthor.Get(), o.ChangeAuthor.IsSet()
}

// HasChangeAuthor returns a boolean if a field has been set.
func (o *SourceManifestMetadataRecord) HasChangeAuthor() bool {
	if o != nil && o.ChangeAuthor.IsSet() {
		return true
	}

	return false
}

// SetChangeAuthor gets a reference to the given NullableString and assigns it to the ChangeAuthor field.
func (o *SourceManifestMetadataRecord) SetChangeAuthor(v string) {
	o.ChangeAuthor.Set(&v)
}
// SetChangeAuthorNil sets the value for ChangeAuthor to be an explicit nil
func (o *SourceManifestMetadataRecord) SetChangeAuthorNil() {
	o.ChangeAuthor.Set(nil)
}

// UnsetChangeAuthor ensures that no value is present for ChangeAuthor, not even an explicit nil
func (o *SourceManifestMetadataRecord) UnsetChangeAuthor() {
	o.ChangeAuthor.Unset()
}

func (o SourceManifestMetadataRecord) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.CiWorkflowName.IsSet() {
		toSerialize["ci_workflow_name"] = o.CiWorkflowName.Get()
	}
	if o.CiWorkflowExecutionTime.IsSet() {
		toSerialize["ci_workflow_execution_time"] = o.CiWorkflowExecutionTime.Get()
	}
	if o.BranchName.IsSet() {
		toSerialize["branch_name"] = o.BranchName.Get()
	}
	if o.ChangeAuthor.IsSet() {
		toSerialize["change_author"] = o.ChangeAuthor.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSourceManifestMetadataRecord struct {
	value *SourceManifestMetadataRecord
	isSet bool
}

func (v NullableSourceManifestMetadataRecord) Get() *SourceManifestMetadataRecord {
	return v.value
}

func (v *NullableSourceManifestMetadataRecord) Set(val *SourceManifestMetadataRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceManifestMetadataRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceManifestMetadataRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceManifestMetadataRecord(val *SourceManifestMetadataRecord) *NullableSourceManifestMetadataRecord {
	return &NullableSourceManifestMetadataRecord{value: val, isSet: true}
}

func (v NullableSourceManifestMetadataRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceManifestMetadataRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


