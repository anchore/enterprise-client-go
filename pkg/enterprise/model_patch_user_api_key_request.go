/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.11.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the PatchUserApiKeyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchUserApiKeyRequest{}

// PatchUserApiKeyRequest struct for PatchUserApiKeyRequest
type PatchUserApiKeyRequest struct {
	Description *string
	Name *string
	Status *string
}

// NewPatchUserApiKeyRequest instantiates a new PatchUserApiKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchUserApiKeyRequest() *PatchUserApiKeyRequest {
	this := PatchUserApiKeyRequest{}
	return &this
}

// NewPatchUserApiKeyRequestWithDefaults instantiates a new PatchUserApiKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchUserApiKeyRequestWithDefaults() *PatchUserApiKeyRequest {
	this := PatchUserApiKeyRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchUserApiKeyRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchUserApiKeyRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchUserApiKeyRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchUserApiKeyRequest) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchUserApiKeyRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchUserApiKeyRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchUserApiKeyRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchUserApiKeyRequest) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PatchUserApiKeyRequest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchUserApiKeyRequest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PatchUserApiKeyRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PatchUserApiKeyRequest) SetStatus(v string) {
	o.Status = &v
}

func (o PatchUserApiKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchUserApiKeyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullablePatchUserApiKeyRequest struct {
	value *PatchUserApiKeyRequest
	isSet bool
}

func (v NullablePatchUserApiKeyRequest) Get() *PatchUserApiKeyRequest {
	return v.value
}

func (v *NullablePatchUserApiKeyRequest) Set(val *PatchUserApiKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchUserApiKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchUserApiKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchUserApiKeyRequest(val *PatchUserApiKeyRequest) *NullablePatchUserApiKeyRequest {
	return &NullablePatchUserApiKeyRequest{value: val, isSet: true}
}

func (v NullablePatchUserApiKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchUserApiKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


