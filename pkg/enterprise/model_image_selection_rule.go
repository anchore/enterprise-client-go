/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

<<<<<<< HEAD
API version: 2.0.0
=======
API version: 0.1.0
>>>>>>> 48fc108 (feat: updated the enterprise ref)
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// ImageSelectionRule struct for ImageSelectionRule
type ImageSelectionRule struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Registry string `json:"registry"`
	Repository string `json:"repository"`
	Image ImageRef `json:"image"`
	// Description of the Allowlist or Denylist image match, human readable
	Description *string `json:"description,omitempty"`
}

// NewImageSelectionRule instantiates a new ImageSelectionRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageSelectionRule(id string, name string, registry string, repository string, image ImageRef) *ImageSelectionRule {
	this := ImageSelectionRule{}
	this.Id = id
	this.Name = name
	this.Registry = registry
	this.Repository = repository
	this.Image = image
	return &this
}

// NewImageSelectionRuleWithDefaults instantiates a new ImageSelectionRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageSelectionRuleWithDefaults() *ImageSelectionRule {
	this := ImageSelectionRule{}
	return &this
}

// GetId returns the Id field value
func (o *ImageSelectionRule) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ImageSelectionRule) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ImageSelectionRule) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ImageSelectionRule) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ImageSelectionRule) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ImageSelectionRule) SetName(v string) {
	o.Name = v
}

// GetRegistry returns the Registry field value
func (o *ImageSelectionRule) GetRegistry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Registry
}

// GetRegistryOk returns a tuple with the Registry field value
// and a boolean to check if the value has been set.
func (o *ImageSelectionRule) GetRegistryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Registry, true
}

// SetRegistry sets field value
func (o *ImageSelectionRule) SetRegistry(v string) {
	o.Registry = v
}

// GetRepository returns the Repository field value
func (o *ImageSelectionRule) GetRepository() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *ImageSelectionRule) GetRepositoryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *ImageSelectionRule) SetRepository(v string) {
	o.Repository = v
}

// GetImage returns the Image field value
func (o *ImageSelectionRule) GetImage() ImageRef {
	if o == nil {
		var ret ImageRef
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *ImageSelectionRule) GetImageOk() (*ImageRef, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *ImageSelectionRule) SetImage(v ImageRef) {
	o.Image = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ImageSelectionRule) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageSelectionRule) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ImageSelectionRule) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ImageSelectionRule) SetDescription(v string) {
	o.Description = &v
}

func (o ImageSelectionRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["registry"] = o.Registry
	}
	if true {
		toSerialize["repository"] = o.Repository
	}
	if true {
		toSerialize["image"] = o.Image
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableImageSelectionRule struct {
	value *ImageSelectionRule
	isSet bool
}

func (v NullableImageSelectionRule) Get() *ImageSelectionRule {
	return v.value
}

func (v *NullableImageSelectionRule) Set(val *ImageSelectionRule) {
	v.value = val
	v.isSet = true
}

func (v NullableImageSelectionRule) IsSet() bool {
	return v.isSet
}

func (v *NullableImageSelectionRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageSelectionRule(val *ImageSelectionRule) *NullableImageSelectionRule {
	return &NullableImageSelectionRule{value: val, isSet: true}
}

func (v NullableImageSelectionRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageSelectionRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


