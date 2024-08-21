/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.4.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// NativeSBOM struct for NativeSBOM
type NativeSBOM struct {
	Artifacts []NativeSBOMPackage `json:"artifacts"`
	Source NativeSBOMSource `json:"source"`
	Distro NativeSBOMDistribution `json:"distro"`
	Descriptor *NativeSBOMDescriptor `json:"descriptor,omitempty"`
	Schema *NativeSBOMSchema `json:"schema,omitempty"`
	ArtifactRelationships []NativeSBOMPackageRelationship `json:"artifactRelationships,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NativeSBOM NativeSBOM

// NewNativeSBOM instantiates a new NativeSBOM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNativeSBOM(artifacts []NativeSBOMPackage, source NativeSBOMSource, distro NativeSBOMDistribution) *NativeSBOM {
	this := NativeSBOM{}
	this.Artifacts = artifacts
	this.Source = source
	this.Distro = distro
	return &this
}

// NewNativeSBOMWithDefaults instantiates a new NativeSBOM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNativeSBOMWithDefaults() *NativeSBOM {
	this := NativeSBOM{}
	return &this
}

// GetArtifacts returns the Artifacts field value
func (o *NativeSBOM) GetArtifacts() []NativeSBOMPackage {
	if o == nil {
		var ret []NativeSBOMPackage
		return ret
	}

	return o.Artifacts
}

// GetArtifactsOk returns a tuple with the Artifacts field value
// and a boolean to check if the value has been set.
func (o *NativeSBOM) GetArtifactsOk() ([]NativeSBOMPackage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Artifacts, true
}

// SetArtifacts sets field value
func (o *NativeSBOM) SetArtifacts(v []NativeSBOMPackage) {
	o.Artifacts = v
}

// GetSource returns the Source field value
func (o *NativeSBOM) GetSource() NativeSBOMSource {
	if o == nil {
		var ret NativeSBOMSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *NativeSBOM) GetSourceOk() (*NativeSBOMSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *NativeSBOM) SetSource(v NativeSBOMSource) {
	o.Source = v
}

// GetDistro returns the Distro field value
func (o *NativeSBOM) GetDistro() NativeSBOMDistribution {
	if o == nil {
		var ret NativeSBOMDistribution
		return ret
	}

	return o.Distro
}

// GetDistroOk returns a tuple with the Distro field value
// and a boolean to check if the value has been set.
func (o *NativeSBOM) GetDistroOk() (*NativeSBOMDistribution, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Distro, true
}

// SetDistro sets field value
func (o *NativeSBOM) SetDistro(v NativeSBOMDistribution) {
	o.Distro = v
}

// GetDescriptor returns the Descriptor field value if set, zero value otherwise.
func (o *NativeSBOM) GetDescriptor() NativeSBOMDescriptor {
	if o == nil || o.Descriptor == nil {
		var ret NativeSBOMDescriptor
		return ret
	}
	return *o.Descriptor
}

// GetDescriptorOk returns a tuple with the Descriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NativeSBOM) GetDescriptorOk() (*NativeSBOMDescriptor, bool) {
	if o == nil || o.Descriptor == nil {
		return nil, false
	}
	return o.Descriptor, true
}

// HasDescriptor returns a boolean if a field has been set.
func (o *NativeSBOM) HasDescriptor() bool {
	if o != nil && o.Descriptor != nil {
		return true
	}

	return false
}

// SetDescriptor gets a reference to the given NativeSBOMDescriptor and assigns it to the Descriptor field.
func (o *NativeSBOM) SetDescriptor(v NativeSBOMDescriptor) {
	o.Descriptor = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *NativeSBOM) GetSchema() NativeSBOMSchema {
	if o == nil || o.Schema == nil {
		var ret NativeSBOMSchema
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NativeSBOM) GetSchemaOk() (*NativeSBOMSchema, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *NativeSBOM) HasSchema() bool {
	if o != nil && o.Schema != nil {
		return true
	}

	return false
}

// SetSchema gets a reference to the given NativeSBOMSchema and assigns it to the Schema field.
func (o *NativeSBOM) SetSchema(v NativeSBOMSchema) {
	o.Schema = &v
}

// GetArtifactRelationships returns the ArtifactRelationships field value if set, zero value otherwise.
func (o *NativeSBOM) GetArtifactRelationships() []NativeSBOMPackageRelationship {
	if o == nil || o.ArtifactRelationships == nil {
		var ret []NativeSBOMPackageRelationship
		return ret
	}
	return o.ArtifactRelationships
}

// GetArtifactRelationshipsOk returns a tuple with the ArtifactRelationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NativeSBOM) GetArtifactRelationshipsOk() ([]NativeSBOMPackageRelationship, bool) {
	if o == nil || o.ArtifactRelationships == nil {
		return nil, false
	}
	return o.ArtifactRelationships, true
}

// HasArtifactRelationships returns a boolean if a field has been set.
func (o *NativeSBOM) HasArtifactRelationships() bool {
	if o != nil && o.ArtifactRelationships != nil {
		return true
	}

	return false
}

// SetArtifactRelationships gets a reference to the given []NativeSBOMPackageRelationship and assigns it to the ArtifactRelationships field.
func (o *NativeSBOM) SetArtifactRelationships(v []NativeSBOMPackageRelationship) {
	o.ArtifactRelationships = v
}

func (o NativeSBOM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["artifacts"] = o.Artifacts
	}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["distro"] = o.Distro
	}
	if o.Descriptor != nil {
		toSerialize["descriptor"] = o.Descriptor
	}
	if o.Schema != nil {
		toSerialize["schema"] = o.Schema
	}
	if o.ArtifactRelationships != nil {
		toSerialize["artifactRelationships"] = o.ArtifactRelationships
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NativeSBOM) UnmarshalJSON(bytes []byte) (err error) {
	varNativeSBOM := _NativeSBOM{}

	if err = json.Unmarshal(bytes, &varNativeSBOM); err == nil {
		*o = NativeSBOM(varNativeSBOM)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "artifacts")
		delete(additionalProperties, "source")
		delete(additionalProperties, "distro")
		delete(additionalProperties, "descriptor")
		delete(additionalProperties, "schema")
		delete(additionalProperties, "artifactRelationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNativeSBOM struct {
	value *NativeSBOM
	isSet bool
}

func (v NullableNativeSBOM) Get() *NativeSBOM {
	return v.value
}

func (v *NullableNativeSBOM) Set(val *NativeSBOM) {
	v.value = val
	v.isSet = true
}

func (v NullableNativeSBOM) IsSet() bool {
	return v.isSet
}

func (v *NullableNativeSBOM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNativeSBOM(val *NativeSBOM) *NullableNativeSBOM {
	return &NullableNativeSBOM{value: val, isSet: true}
}

func (v NullableNativeSBOM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNativeSBOM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


