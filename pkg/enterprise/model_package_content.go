/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.6.1
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// PackageContent struct for PackageContent
type PackageContent struct {
	Package *string `json:"package,omitempty"`
	Version *string `json:"version,omitempty"`
	Size *string `json:"size,omitempty"`
	Type *string `json:"type,omitempty"`
	Origin *string `json:"origin,omitempty"`
	Licenses []string `json:"licenses,omitempty"`
	Location *string `json:"location,omitempty"`
	// A list of Common Platform Enumerations that may uniquely identify the package
	Cpes []string `json:"cpes,omitempty"`
	// The type of the metadata entry
	MetadataType *string `json:"metadata_type,omitempty"`
	// Package type specific metadata
	Metadata interface{} `json:"metadata,omitempty"`
	Purl *string `json:"purl,omitempty"`
}

// NewPackageContent instantiates a new PackageContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageContent() *PackageContent {
	this := PackageContent{}
	return &this
}

// NewPackageContentWithDefaults instantiates a new PackageContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageContentWithDefaults() *PackageContent {
	this := PackageContent{}
	return &this
}

// GetPackage returns the Package field value if set, zero value otherwise.
func (o *PackageContent) GetPackage() string {
	if o == nil || o.Package == nil {
		var ret string
		return ret
	}
	return *o.Package
}

// GetPackageOk returns a tuple with the Package field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageContent) GetPackageOk() (*string, bool) {
	if o == nil || o.Package == nil {
		return nil, false
	}
	return o.Package, true
}

// HasPackage returns a boolean if a field has been set.
func (o *PackageContent) HasPackage() bool {
	if o != nil && o.Package != nil {
		return true
	}

	return false
}

// SetPackage gets a reference to the given string and assigns it to the Package field.
func (o *PackageContent) SetPackage(v string) {
	o.Package = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PackageContent) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageContent) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PackageContent) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *PackageContent) SetVersion(v string) {
	o.Version = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *PackageContent) GetSize() string {
	if o == nil || o.Size == nil {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageContent) GetSizeOk() (*string, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *PackageContent) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *PackageContent) SetSize(v string) {
	o.Size = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PackageContent) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageContent) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PackageContent) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PackageContent) SetType(v string) {
	o.Type = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *PackageContent) GetOrigin() string {
	if o == nil || o.Origin == nil {
		var ret string
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageContent) GetOriginOk() (*string, bool) {
	if o == nil || o.Origin == nil {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *PackageContent) HasOrigin() bool {
	if o != nil && o.Origin != nil {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given string and assigns it to the Origin field.
func (o *PackageContent) SetOrigin(v string) {
	o.Origin = &v
}

// GetLicenses returns the Licenses field value if set, zero value otherwise.
func (o *PackageContent) GetLicenses() []string {
	if o == nil || o.Licenses == nil {
		var ret []string
		return ret
	}
	return o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageContent) GetLicensesOk() ([]string, bool) {
	if o == nil || o.Licenses == nil {
		return nil, false
	}
	return o.Licenses, true
}

// HasLicenses returns a boolean if a field has been set.
func (o *PackageContent) HasLicenses() bool {
	if o != nil && o.Licenses != nil {
		return true
	}

	return false
}

// SetLicenses gets a reference to the given []string and assigns it to the Licenses field.
func (o *PackageContent) SetLicenses(v []string) {
	o.Licenses = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *PackageContent) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageContent) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *PackageContent) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *PackageContent) SetLocation(v string) {
	o.Location = &v
}

// GetCpes returns the Cpes field value if set, zero value otherwise.
func (o *PackageContent) GetCpes() []string {
	if o == nil || o.Cpes == nil {
		var ret []string
		return ret
	}
	return o.Cpes
}

// GetCpesOk returns a tuple with the Cpes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageContent) GetCpesOk() ([]string, bool) {
	if o == nil || o.Cpes == nil {
		return nil, false
	}
	return o.Cpes, true
}

// HasCpes returns a boolean if a field has been set.
func (o *PackageContent) HasCpes() bool {
	if o != nil && o.Cpes != nil {
		return true
	}

	return false
}

// SetCpes gets a reference to the given []string and assigns it to the Cpes field.
func (o *PackageContent) SetCpes(v []string) {
	o.Cpes = v
}

// GetMetadataType returns the MetadataType field value if set, zero value otherwise.
func (o *PackageContent) GetMetadataType() string {
	if o == nil || o.MetadataType == nil {
		var ret string
		return ret
	}
	return *o.MetadataType
}

// GetMetadataTypeOk returns a tuple with the MetadataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageContent) GetMetadataTypeOk() (*string, bool) {
	if o == nil || o.MetadataType == nil {
		return nil, false
	}
	return o.MetadataType, true
}

// HasMetadataType returns a boolean if a field has been set.
func (o *PackageContent) HasMetadataType() bool {
	if o != nil && o.MetadataType != nil {
		return true
	}

	return false
}

// SetMetadataType gets a reference to the given string and assigns it to the MetadataType field.
func (o *PackageContent) SetMetadataType(v string) {
	o.MetadataType = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PackageContent) GetMetadata() interface{} {
	if o == nil || o.Metadata == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageContent) GetMetadataOk() (interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PackageContent) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PackageContent) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetPurl returns the Purl field value if set, zero value otherwise.
func (o *PackageContent) GetPurl() string {
	if o == nil || o.Purl == nil {
		var ret string
		return ret
	}
	return *o.Purl
}

// GetPurlOk returns a tuple with the Purl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageContent) GetPurlOk() (*string, bool) {
	if o == nil || o.Purl == nil {
		return nil, false
	}
	return o.Purl, true
}

// HasPurl returns a boolean if a field has been set.
func (o *PackageContent) HasPurl() bool {
	if o != nil && o.Purl != nil {
		return true
	}

	return false
}

// SetPurl gets a reference to the given string and assigns it to the Purl field.
func (o *PackageContent) SetPurl(v string) {
	o.Purl = &v
}

func (o PackageContent) MarshalJSON() ([]byte, error) {
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
	if o.Purl != nil {
		toSerialize["purl"] = o.Purl
	}
	return json.Marshal(toSerialize)
}

type NullablePackageContent struct {
	value *PackageContent
	isSet bool
}

func (v NullablePackageContent) Get() *PackageContent {
	return v.value
}

func (v *NullablePackageContent) Set(val *PackageContent) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageContent) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageContent(val *PackageContent) *NullablePackageContent {
	return &NullablePackageContent{value: val, isSet: true}
}

func (v NullablePackageContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


