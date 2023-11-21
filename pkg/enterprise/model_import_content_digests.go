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
)

// ImportContentDigests Digest of content to use in the final import
type ImportContentDigests struct {
	// Digest to use for the packages content
	Packages string `json:"packages"`
	// Digest for reference content for image config
	ImageConfig string `json:"image_config"`
	// Digest to reference content for the image manifest
	Manifest string `json:"manifest"`
	// Digest for reference content for parent manifest
	ParentManifest *string `json:"parent_manifest,omitempty"`
	// Digest for reference content for dockerfile
	Dockerfile *string `json:"dockerfile,omitempty"`
	// Digest for reference content for secret search results
	SecretSearches *string `json:"secret_searches,omitempty"`
	// Digest for reference content for content search results
	ContentSearches *string `json:"content_searches,omitempty"`
	// Digest for reference content for file retrieve content
	FileContents *string `json:"file_contents,omitempty"`
}

// NewImportContentDigests instantiates a new ImportContentDigests object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportContentDigests(packages string, imageConfig string, manifest string) *ImportContentDigests {
	this := ImportContentDigests{}
	this.Packages = packages
	this.ImageConfig = imageConfig
	this.Manifest = manifest
	return &this
}

// NewImportContentDigestsWithDefaults instantiates a new ImportContentDigests object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportContentDigestsWithDefaults() *ImportContentDigests {
	this := ImportContentDigests{}
	return &this
}

// GetPackages returns the Packages field value
func (o *ImportContentDigests) GetPackages() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Packages
}

// GetPackagesOk returns a tuple with the Packages field value
// and a boolean to check if the value has been set.
func (o *ImportContentDigests) GetPackagesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Packages, true
}

// SetPackages sets field value
func (o *ImportContentDigests) SetPackages(v string) {
	o.Packages = v
}

// GetImageConfig returns the ImageConfig field value
func (o *ImportContentDigests) GetImageConfig() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageConfig
}

// GetImageConfigOk returns a tuple with the ImageConfig field value
// and a boolean to check if the value has been set.
func (o *ImportContentDigests) GetImageConfigOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageConfig, true
}

// SetImageConfig sets field value
func (o *ImportContentDigests) SetImageConfig(v string) {
	o.ImageConfig = v
}

// GetManifest returns the Manifest field value
func (o *ImportContentDigests) GetManifest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Manifest
}

// GetManifestOk returns a tuple with the Manifest field value
// and a boolean to check if the value has been set.
func (o *ImportContentDigests) GetManifestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Manifest, true
}

// SetManifest sets field value
func (o *ImportContentDigests) SetManifest(v string) {
	o.Manifest = v
}

// GetParentManifest returns the ParentManifest field value if set, zero value otherwise.
func (o *ImportContentDigests) GetParentManifest() string {
	if o == nil || o.ParentManifest == nil {
		var ret string
		return ret
	}
	return *o.ParentManifest
}

// GetParentManifestOk returns a tuple with the ParentManifest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportContentDigests) GetParentManifestOk() (*string, bool) {
	if o == nil || o.ParentManifest == nil {
		return nil, false
	}
	return o.ParentManifest, true
}

// HasParentManifest returns a boolean if a field has been set.
func (o *ImportContentDigests) HasParentManifest() bool {
	if o != nil && o.ParentManifest != nil {
		return true
	}

	return false
}

// SetParentManifest gets a reference to the given string and assigns it to the ParentManifest field.
func (o *ImportContentDigests) SetParentManifest(v string) {
	o.ParentManifest = &v
}

// GetDockerfile returns the Dockerfile field value if set, zero value otherwise.
func (o *ImportContentDigests) GetDockerfile() string {
	if o == nil || o.Dockerfile == nil {
		var ret string
		return ret
	}
	return *o.Dockerfile
}

// GetDockerfileOk returns a tuple with the Dockerfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportContentDigests) GetDockerfileOk() (*string, bool) {
	if o == nil || o.Dockerfile == nil {
		return nil, false
	}
	return o.Dockerfile, true
}

// HasDockerfile returns a boolean if a field has been set.
func (o *ImportContentDigests) HasDockerfile() bool {
	if o != nil && o.Dockerfile != nil {
		return true
	}

	return false
}

// SetDockerfile gets a reference to the given string and assigns it to the Dockerfile field.
func (o *ImportContentDigests) SetDockerfile(v string) {
	o.Dockerfile = &v
}

// GetSecretSearches returns the SecretSearches field value if set, zero value otherwise.
func (o *ImportContentDigests) GetSecretSearches() string {
	if o == nil || o.SecretSearches == nil {
		var ret string
		return ret
	}
	return *o.SecretSearches
}

// GetSecretSearchesOk returns a tuple with the SecretSearches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportContentDigests) GetSecretSearchesOk() (*string, bool) {
	if o == nil || o.SecretSearches == nil {
		return nil, false
	}
	return o.SecretSearches, true
}

// HasSecretSearches returns a boolean if a field has been set.
func (o *ImportContentDigests) HasSecretSearches() bool {
	if o != nil && o.SecretSearches != nil {
		return true
	}

	return false
}

// SetSecretSearches gets a reference to the given string and assigns it to the SecretSearches field.
func (o *ImportContentDigests) SetSecretSearches(v string) {
	o.SecretSearches = &v
}

// GetContentSearches returns the ContentSearches field value if set, zero value otherwise.
func (o *ImportContentDigests) GetContentSearches() string {
	if o == nil || o.ContentSearches == nil {
		var ret string
		return ret
	}
	return *o.ContentSearches
}

// GetContentSearchesOk returns a tuple with the ContentSearches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportContentDigests) GetContentSearchesOk() (*string, bool) {
	if o == nil || o.ContentSearches == nil {
		return nil, false
	}
	return o.ContentSearches, true
}

// HasContentSearches returns a boolean if a field has been set.
func (o *ImportContentDigests) HasContentSearches() bool {
	if o != nil && o.ContentSearches != nil {
		return true
	}

	return false
}

// SetContentSearches gets a reference to the given string and assigns it to the ContentSearches field.
func (o *ImportContentDigests) SetContentSearches(v string) {
	o.ContentSearches = &v
}

// GetFileContents returns the FileContents field value if set, zero value otherwise.
func (o *ImportContentDigests) GetFileContents() string {
	if o == nil || o.FileContents == nil {
		var ret string
		return ret
	}
	return *o.FileContents
}

// GetFileContentsOk returns a tuple with the FileContents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportContentDigests) GetFileContentsOk() (*string, bool) {
	if o == nil || o.FileContents == nil {
		return nil, false
	}
	return o.FileContents, true
}

// HasFileContents returns a boolean if a field has been set.
func (o *ImportContentDigests) HasFileContents() bool {
	if o != nil && o.FileContents != nil {
		return true
	}

	return false
}

// SetFileContents gets a reference to the given string and assigns it to the FileContents field.
func (o *ImportContentDigests) SetFileContents(v string) {
	o.FileContents = &v
}

func (o ImportContentDigests) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["packages"] = o.Packages
	}
	if true {
		toSerialize["image_config"] = o.ImageConfig
	}
	if true {
		toSerialize["manifest"] = o.Manifest
	}
	if o.ParentManifest != nil {
		toSerialize["parent_manifest"] = o.ParentManifest
	}
	if o.Dockerfile != nil {
		toSerialize["dockerfile"] = o.Dockerfile
	}
	if o.SecretSearches != nil {
		toSerialize["secret_searches"] = o.SecretSearches
	}
	if o.ContentSearches != nil {
		toSerialize["content_searches"] = o.ContentSearches
	}
	if o.FileContents != nil {
		toSerialize["file_contents"] = o.FileContents
	}
	return json.Marshal(toSerialize)
}

type NullableImportContentDigests struct {
	value *ImportContentDigests
	isSet bool
}

func (v NullableImportContentDigests) Get() *ImportContentDigests {
	return v.value
}

func (v *NullableImportContentDigests) Set(val *ImportContentDigests) {
	v.value = val
	v.isSet = true
}

func (v NullableImportContentDigests) IsSet() bool {
	return v.isSet
}

func (v *NullableImportContentDigests) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportContentDigests(val *ImportContentDigests) *NullableImportContentDigests {
	return &NullableImportContentDigests{value: val, isSet: true}
}

func (v NullableImportContentDigests) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportContentDigests) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


