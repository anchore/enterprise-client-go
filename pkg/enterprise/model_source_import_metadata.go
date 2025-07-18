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
	"bytes"
	"fmt"
)

// checks if the SourceImportMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceImportMetadata{}

// SourceImportMetadata struct for SourceImportMetadata
type SourceImportMetadata struct {
	CiWorkflowName NullableString `json:"ci_workflow_name,omitempty"`
	CiWorkflowExecutionTime NullableTime `json:"ci_workflow_execution_time,omitempty"`
	Host string `json:"host"`
	RepositoryName string `json:"repository_name"`
	BranchName NullableString `json:"branch_name,omitempty"`
	Revision string `json:"revision"`
	ChangeAuthor NullableString `json:"change_author,omitempty"`
	Contents SourceImportMetadataContents `json:"contents"`
}

type _SourceImportMetadata SourceImportMetadata

// NewSourceImportMetadata instantiates a new SourceImportMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceImportMetadata(host string, repositoryName string, revision string, contents SourceImportMetadataContents) *SourceImportMetadata {
	this := SourceImportMetadata{}
	this.Host = host
	this.RepositoryName = repositoryName
	this.Revision = revision
	this.Contents = contents
	return &this
}

// NewSourceImportMetadataWithDefaults instantiates a new SourceImportMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceImportMetadataWithDefaults() *SourceImportMetadata {
	this := SourceImportMetadata{}
	return &this
}

// GetCiWorkflowName returns the CiWorkflowName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceImportMetadata) GetCiWorkflowName() string {
	if o == nil || IsNil(o.CiWorkflowName.Get()) {
		var ret string
		return ret
	}
	return *o.CiWorkflowName.Get()
}

// GetCiWorkflowNameOk returns a tuple with the CiWorkflowName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceImportMetadata) GetCiWorkflowNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CiWorkflowName.Get(), o.CiWorkflowName.IsSet()
}

// HasCiWorkflowName returns a boolean if a field has been set.
func (o *SourceImportMetadata) HasCiWorkflowName() bool {
	if o != nil && o.CiWorkflowName.IsSet() {
		return true
	}

	return false
}

// SetCiWorkflowName gets a reference to the given NullableString and assigns it to the CiWorkflowName field.
func (o *SourceImportMetadata) SetCiWorkflowName(v string) {
	o.CiWorkflowName.Set(&v)
}
// SetCiWorkflowNameNil sets the value for CiWorkflowName to be an explicit nil
func (o *SourceImportMetadata) SetCiWorkflowNameNil() {
	o.CiWorkflowName.Set(nil)
}

// UnsetCiWorkflowName ensures that no value is present for CiWorkflowName, not even an explicit nil
func (o *SourceImportMetadata) UnsetCiWorkflowName() {
	o.CiWorkflowName.Unset()
}

// GetCiWorkflowExecutionTime returns the CiWorkflowExecutionTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceImportMetadata) GetCiWorkflowExecutionTime() time.Time {
	if o == nil || IsNil(o.CiWorkflowExecutionTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CiWorkflowExecutionTime.Get()
}

// GetCiWorkflowExecutionTimeOk returns a tuple with the CiWorkflowExecutionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceImportMetadata) GetCiWorkflowExecutionTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CiWorkflowExecutionTime.Get(), o.CiWorkflowExecutionTime.IsSet()
}

// HasCiWorkflowExecutionTime returns a boolean if a field has been set.
func (o *SourceImportMetadata) HasCiWorkflowExecutionTime() bool {
	if o != nil && o.CiWorkflowExecutionTime.IsSet() {
		return true
	}

	return false
}

// SetCiWorkflowExecutionTime gets a reference to the given NullableTime and assigns it to the CiWorkflowExecutionTime field.
func (o *SourceImportMetadata) SetCiWorkflowExecutionTime(v time.Time) {
	o.CiWorkflowExecutionTime.Set(&v)
}
// SetCiWorkflowExecutionTimeNil sets the value for CiWorkflowExecutionTime to be an explicit nil
func (o *SourceImportMetadata) SetCiWorkflowExecutionTimeNil() {
	o.CiWorkflowExecutionTime.Set(nil)
}

// UnsetCiWorkflowExecutionTime ensures that no value is present for CiWorkflowExecutionTime, not even an explicit nil
func (o *SourceImportMetadata) UnsetCiWorkflowExecutionTime() {
	o.CiWorkflowExecutionTime.Unset()
}

// GetHost returns the Host field value
func (o *SourceImportMetadata) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *SourceImportMetadata) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *SourceImportMetadata) SetHost(v string) {
	o.Host = v
}

// GetRepositoryName returns the RepositoryName field value
func (o *SourceImportMetadata) GetRepositoryName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RepositoryName
}

// GetRepositoryNameOk returns a tuple with the RepositoryName field value
// and a boolean to check if the value has been set.
func (o *SourceImportMetadata) GetRepositoryNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepositoryName, true
}

// SetRepositoryName sets field value
func (o *SourceImportMetadata) SetRepositoryName(v string) {
	o.RepositoryName = v
}

// GetBranchName returns the BranchName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceImportMetadata) GetBranchName() string {
	if o == nil || IsNil(o.BranchName.Get()) {
		var ret string
		return ret
	}
	return *o.BranchName.Get()
}

// GetBranchNameOk returns a tuple with the BranchName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceImportMetadata) GetBranchNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BranchName.Get(), o.BranchName.IsSet()
}

// HasBranchName returns a boolean if a field has been set.
func (o *SourceImportMetadata) HasBranchName() bool {
	if o != nil && o.BranchName.IsSet() {
		return true
	}

	return false
}

// SetBranchName gets a reference to the given NullableString and assigns it to the BranchName field.
func (o *SourceImportMetadata) SetBranchName(v string) {
	o.BranchName.Set(&v)
}
// SetBranchNameNil sets the value for BranchName to be an explicit nil
func (o *SourceImportMetadata) SetBranchNameNil() {
	o.BranchName.Set(nil)
}

// UnsetBranchName ensures that no value is present for BranchName, not even an explicit nil
func (o *SourceImportMetadata) UnsetBranchName() {
	o.BranchName.Unset()
}

// GetRevision returns the Revision field value
func (o *SourceImportMetadata) GetRevision() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *SourceImportMetadata) GetRevisionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value
func (o *SourceImportMetadata) SetRevision(v string) {
	o.Revision = v
}

// GetChangeAuthor returns the ChangeAuthor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceImportMetadata) GetChangeAuthor() string {
	if o == nil || IsNil(o.ChangeAuthor.Get()) {
		var ret string
		return ret
	}
	return *o.ChangeAuthor.Get()
}

// GetChangeAuthorOk returns a tuple with the ChangeAuthor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceImportMetadata) GetChangeAuthorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChangeAuthor.Get(), o.ChangeAuthor.IsSet()
}

// HasChangeAuthor returns a boolean if a field has been set.
func (o *SourceImportMetadata) HasChangeAuthor() bool {
	if o != nil && o.ChangeAuthor.IsSet() {
		return true
	}

	return false
}

// SetChangeAuthor gets a reference to the given NullableString and assigns it to the ChangeAuthor field.
func (o *SourceImportMetadata) SetChangeAuthor(v string) {
	o.ChangeAuthor.Set(&v)
}
// SetChangeAuthorNil sets the value for ChangeAuthor to be an explicit nil
func (o *SourceImportMetadata) SetChangeAuthorNil() {
	o.ChangeAuthor.Set(nil)
}

// UnsetChangeAuthor ensures that no value is present for ChangeAuthor, not even an explicit nil
func (o *SourceImportMetadata) UnsetChangeAuthor() {
	o.ChangeAuthor.Unset()
}

// GetContents returns the Contents field value
func (o *SourceImportMetadata) GetContents() SourceImportMetadataContents {
	if o == nil {
		var ret SourceImportMetadataContents
		return ret
	}

	return o.Contents
}

// GetContentsOk returns a tuple with the Contents field value
// and a boolean to check if the value has been set.
func (o *SourceImportMetadata) GetContentsOk() (*SourceImportMetadataContents, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contents, true
}

// SetContents sets field value
func (o *SourceImportMetadata) SetContents(v SourceImportMetadataContents) {
	o.Contents = v
}

func (o SourceImportMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceImportMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CiWorkflowName.IsSet() {
		toSerialize["ci_workflow_name"] = o.CiWorkflowName.Get()
	}
	if o.CiWorkflowExecutionTime.IsSet() {
		toSerialize["ci_workflow_execution_time"] = o.CiWorkflowExecutionTime.Get()
	}
	toSerialize["host"] = o.Host
	toSerialize["repository_name"] = o.RepositoryName
	if o.BranchName.IsSet() {
		toSerialize["branch_name"] = o.BranchName.Get()
	}
	toSerialize["revision"] = o.Revision
	if o.ChangeAuthor.IsSet() {
		toSerialize["change_author"] = o.ChangeAuthor.Get()
	}
	toSerialize["contents"] = o.Contents
	return toSerialize, nil
}

func (o *SourceImportMetadata) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"host",
		"repository_name",
		"revision",
		"contents",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSourceImportMetadata := _SourceImportMetadata{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSourceImportMetadata)

	if err != nil {
		return err
	}

	*o = SourceImportMetadata(varSourceImportMetadata)

	return err
}

type NullableSourceImportMetadata struct {
	value *SourceImportMetadata
	isSet bool
}

func (v NullableSourceImportMetadata) Get() *SourceImportMetadata {
	return v.value
}

func (v *NullableSourceImportMetadata) Set(val *SourceImportMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceImportMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceImportMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceImportMetadata(val *SourceImportMetadata) *NullableSourceImportMetadata {
	return &NullableSourceImportMetadata{value: val, isSet: true}
}

func (v NullableSourceImportMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceImportMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


