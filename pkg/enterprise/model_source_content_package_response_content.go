/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 1.0.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// SourceContentPackageResponseContent Package content listings from a source
type SourceContentPackageResponseContent struct {
	Package *string `json:"package,omitempty"`
	Version *string `json:"version,omitempty"`
	Size *string `json:"size,omitempty"`
	Type *string `json:"type,omitempty"`
	Origin *string `json:"origin,omitempty"`
	Licenses *[]string `json:"licenses,omitempty"`
	Location *string `json:"location,omitempty"`
	// A list of Common Platform Enumerations that may uniquely identify the package
	Cpes *[]string `json:"cpes,omitempty"`
	// The type of the metadata entry
	MetadataType *string `json:"metadata_type,omitempty"`
	// Package type specific metadata
	Metadata *interface{} `json:"metadata,omitempty"`
}

// NewSourceContentPackageResponseContent instantiates a new SourceContentPackageResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceContentPackageResponseContent() *SourceContentPackageResponseContent {
	this := SourceContentPackageResponseContent{}
	return &this
}

// NewSourceContentPackageResponseContentWithDefaults instantiates a new SourceContentPackageResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceContentPackageResponseContentWithDefaults() *SourceContentPackageResponseContent {
	this := SourceContentPackageResponseContent{}
	return &this
}

// GetPackage returns the Package field value if set, zero value otherwise.
func (o *SourceContentPackageResponseContent) GetPackage() string {
	if o == nil || o.Package == nil {
		var ret string
		return ret
	}
	return *o.Package
}

// GetPackageOk returns a tuple with the Package field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceContentPackageResponseContent) GetPackageOk() (*string, bool) {
	if o == nil || o.Package == nil {
		return nil, false
	}
	return o.Package, true
}

// HasPackage returns a boolean if a field has been set.
func (o *SourceContentPackageResponseContent) HasPackage() bool {
	if o != nil && o.Package != nil {
		return true
	}

	return false
}

// SetPackage gets a reference to the given string and assigns it to the Package field.
func (o *SourceContentPackageResponseContent) SetPackage(v string) {
	o.Package = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SourceContentPackageResponseContent) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceContentPackageResponseContent) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SourceContentPackageResponseContent) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *SourceContentPackageResponseContent) SetVersion(v string) {
	o.Version = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *SourceContentPackageResponseContent) GetSize() string {
	if o == nil || o.Size == nil {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceContentPackageResponseContent) GetSizeOk() (*string, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *SourceContentPackageResponseContent) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *SourceContentPackageResponseContent) SetSize(v string) {
	o.Size = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SourceContentPackageResponseContent) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceContentPackageResponseContent) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SourceContentPackageResponseContent) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SourceContentPackageResponseContent) SetType(v string) {
	o.Type = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *SourceContentPackageResponseContent) GetOrigin() string {
	if o == nil || o.Origin == nil {
		var ret string
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceContentPackageResponseContent) GetOriginOk() (*string, bool) {
	if o == nil || o.Origin == nil {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *SourceContentPackageResponseContent) HasOrigin() bool {
	if o != nil && o.Origin != nil {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given string and assigns it to the Origin field.
func (o *SourceContentPackageResponseContent) SetOrigin(v string) {
	o.Origin = &v
}

// GetLicenses returns the Licenses field value if set, zero value otherwise.
func (o *SourceContentPackageResponseContent) GetLicenses() []string {
	if o == nil || o.Licenses == nil {
		var ret []string
		return ret
	}
	return *o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceContentPackageResponseContent) GetLicensesOk() (*[]string, bool) {
	if o == nil || o.Licenses == nil {
		return nil, false
	}
	return o.Licenses, true
}

// HasLicenses returns a boolean if a field has been set.
func (o *SourceContentPackageResponseContent) HasLicenses() bool {
	if o != nil && o.Licenses != nil {
		return true
	}

	return false
}

// SetLicenses gets a reference to the given []string and assigns it to the Licenses field.
func (o *SourceContentPackageResponseContent) SetLicenses(v []string) {
	o.Licenses = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *SourceContentPackageResponseContent) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceContentPackageResponseContent) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *SourceContentPackageResponseContent) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *SourceContentPackageResponseContent) SetLocation(v string) {
	o.Location = &v
}

// GetCpes returns the Cpes field value if set, zero value otherwise.
func (o *SourceContentPackageResponseContent) GetCpes() []string {
	if o == nil || o.Cpes == nil {
		var ret []string
		return ret
	}
	return *o.Cpes
}

// GetCpesOk returns a tuple with the Cpes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceContentPackageResponseContent) GetCpesOk() (*[]string, bool) {
	if o == nil || o.Cpes == nil {
		return nil, false
	}
	return o.Cpes, true
}

// HasCpes returns a boolean if a field has been set.
func (o *SourceContentPackageResponseContent) HasCpes() bool {
	if o != nil && o.Cpes != nil {
		return true
	}

	return false
}

// SetCpes gets a reference to the given []string and assigns it to the Cpes field.
func (o *SourceContentPackageResponseContent) SetCpes(v []string) {
	o.Cpes = &v
}

// GetMetadataType returns the MetadataType field value if set, zero value otherwise.
func (o *SourceContentPackageResponseContent) GetMetadataType() string {
	if o == nil || o.MetadataType == nil {
		var ret string
		return ret
	}
	return *o.MetadataType
}

// GetMetadataTypeOk returns a tuple with the MetadataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceContentPackageResponseContent) GetMetadataTypeOk() (*string, bool) {
	if o == nil || o.MetadataType == nil {
		return nil, false
	}
	return o.MetadataType, true
}

// HasMetadataType returns a boolean if a field has been set.
func (o *SourceContentPackageResponseContent) HasMetadataType() bool {
	if o != nil && o.MetadataType != nil {
		return true
	}

	return false
}

// SetMetadataType gets a reference to the given string and assigns it to the MetadataType field.
func (o *SourceContentPackageResponseContent) SetMetadataType(v string) {
	o.MetadataType = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SourceContentPackageResponseContent) GetMetadata() interface{} {
	if o == nil || o.Metadata == nil {
		var ret interface{}
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceContentPackageResponseContent) GetMetadataOk() (*interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SourceContentPackageResponseContent) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *SourceContentPackageResponseContent) SetMetadata(v interface{}) {
	o.Metadata = &v
}

func (o SourceContentPackageResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Package != nil {
		toSerialize["package"] = o.Package
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Origin != nil {
		toSerialize["origin"] = o.Origin
	}
	if o.Licenses != nil {
		toSerialize["licenses"] = o.Licenses
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.Cpes != nil {
		toSerialize["cpes"] = o.Cpes
	}
	if o.MetadataType != nil {
		toSerialize["metadata_type"] = o.MetadataType
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableSourceContentPackageResponseContent struct {
	value *SourceContentPackageResponseContent
	isSet bool
}

func (v NullableSourceContentPackageResponseContent) Get() *SourceContentPackageResponseContent {
	return v.value
}

func (v *NullableSourceContentPackageResponseContent) Set(val *SourceContentPackageResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceContentPackageResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceContentPackageResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceContentPackageResponseContent(val *SourceContentPackageResponseContent) *NullableSourceContentPackageResponseContent {
	return &NullableSourceContentPackageResponseContent{value: val, isSet: true}
}

func (v NullableSourceContentPackageResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceContentPackageResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


