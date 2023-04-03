/*
Anchore Engine API Server

This is the Anchore Engine API. Provides the primary external API for users of the service.

API version: 0.6.0
Contact: nurmi@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package engine

import (
	"encoding/json"
)

// ImportPackage struct for ImportPackage
type ImportPackage struct {
	Id *string `json:"id,omitempty"`
	Name string `json:"name"`
	Version string `json:"version"`
	Type string `json:"type"`
	FoundBy *string `json:"foundBy,omitempty"`
	Locations []ImportPackageLocation `json:"locations"`
	Licenses []string `json:"licenses"`
	Language string `json:"language"`
	Cpes []string `json:"cpes"`
	Purl *string `json:"purl,omitempty"`
	MetadataType string `json:"metadataType"`
	Metadata *interface{} `json:"metadata,omitempty"`
}

// NewImportPackage instantiates a new ImportPackage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportPackage(name string, version string, type_ string, locations []ImportPackageLocation, licenses []string, language string, cpes []string, metadataType string) *ImportPackage {
	this := ImportPackage{}
	this.Name = name
	this.Version = version
	this.Type = type_
	this.Locations = locations
	this.Licenses = licenses
	this.Language = language
	this.Cpes = cpes
	this.MetadataType = metadataType
	return &this
}

// NewImportPackageWithDefaults instantiates a new ImportPackage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportPackageWithDefaults() *ImportPackage {
	this := ImportPackage{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ImportPackage) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportPackage) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ImportPackage) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ImportPackage) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *ImportPackage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ImportPackage) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ImportPackage) SetName(v string) {
	o.Name = v
}

// GetVersion returns the Version field value
func (o *ImportPackage) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ImportPackage) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ImportPackage) SetVersion(v string) {
	o.Version = v
}

// GetType returns the Type field value
func (o *ImportPackage) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ImportPackage) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ImportPackage) SetType(v string) {
	o.Type = v
}

// GetFoundBy returns the FoundBy field value if set, zero value otherwise.
func (o *ImportPackage) GetFoundBy() string {
	if o == nil || o.FoundBy == nil {
		var ret string
		return ret
	}
	return *o.FoundBy
}

// GetFoundByOk returns a tuple with the FoundBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportPackage) GetFoundByOk() (*string, bool) {
	if o == nil || o.FoundBy == nil {
		return nil, false
	}
	return o.FoundBy, true
}

// HasFoundBy returns a boolean if a field has been set.
func (o *ImportPackage) HasFoundBy() bool {
	if o != nil && o.FoundBy != nil {
		return true
	}

	return false
}

// SetFoundBy gets a reference to the given string and assigns it to the FoundBy field.
func (o *ImportPackage) SetFoundBy(v string) {
	o.FoundBy = &v
}

// GetLocations returns the Locations field value
func (o *ImportPackage) GetLocations() []ImportPackageLocation {
	if o == nil {
		var ret []ImportPackageLocation
		return ret
	}

	return o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value
// and a boolean to check if the value has been set.
func (o *ImportPackage) GetLocationsOk() (*[]ImportPackageLocation, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Locations, true
}

// SetLocations sets field value
func (o *ImportPackage) SetLocations(v []ImportPackageLocation) {
	o.Locations = v
}

// GetLicenses returns the Licenses field value
func (o *ImportPackage) GetLicenses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value
// and a boolean to check if the value has been set.
func (o *ImportPackage) GetLicensesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Licenses, true
}

// SetLicenses sets field value
func (o *ImportPackage) SetLicenses(v []string) {
	o.Licenses = v
}

// GetLanguage returns the Language field value
func (o *ImportPackage) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *ImportPackage) GetLanguageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *ImportPackage) SetLanguage(v string) {
	o.Language = v
}

// GetCpes returns the Cpes field value
func (o *ImportPackage) GetCpes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Cpes
}

// GetCpesOk returns a tuple with the Cpes field value
// and a boolean to check if the value has been set.
func (o *ImportPackage) GetCpesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cpes, true
}

// SetCpes sets field value
func (o *ImportPackage) SetCpes(v []string) {
	o.Cpes = v
}

// GetPurl returns the Purl field value if set, zero value otherwise.
func (o *ImportPackage) GetPurl() string {
	if o == nil || o.Purl == nil {
		var ret string
		return ret
	}
	return *o.Purl
}

// GetPurlOk returns a tuple with the Purl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportPackage) GetPurlOk() (*string, bool) {
	if o == nil || o.Purl == nil {
		return nil, false
	}
	return o.Purl, true
}

// HasPurl returns a boolean if a field has been set.
func (o *ImportPackage) HasPurl() bool {
	if o != nil && o.Purl != nil {
		return true
	}

	return false
}

// SetPurl gets a reference to the given string and assigns it to the Purl field.
func (o *ImportPackage) SetPurl(v string) {
	o.Purl = &v
}

// GetMetadataType returns the MetadataType field value
func (o *ImportPackage) GetMetadataType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetadataType
}

// GetMetadataTypeOk returns a tuple with the MetadataType field value
// and a boolean to check if the value has been set.
func (o *ImportPackage) GetMetadataTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MetadataType, true
}

// SetMetadataType sets field value
func (o *ImportPackage) SetMetadataType(v string) {
	o.MetadataType = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ImportPackage) GetMetadata() interface{} {
	if o == nil || o.Metadata == nil {
		var ret interface{}
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportPackage) GetMetadataOk() (*interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ImportPackage) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *ImportPackage) SetMetadata(v interface{}) {
	o.Metadata = &v
}

func (o ImportPackage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.FoundBy != nil {
		toSerialize["foundBy"] = o.FoundBy
	}
	if true {
		toSerialize["locations"] = o.Locations
	}
	if true {
		toSerialize["licenses"] = o.Licenses
	}
	if true {
		toSerialize["language"] = o.Language
	}
	if true {
		toSerialize["cpes"] = o.Cpes
	}
	if o.Purl != nil {
		toSerialize["purl"] = o.Purl
	}
	if true {
		toSerialize["metadataType"] = o.MetadataType
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableImportPackage struct {
	value *ImportPackage
	isSet bool
}

func (v NullableImportPackage) Get() *ImportPackage {
	return v.value
}

func (v *NullableImportPackage) Set(val *ImportPackage) {
	v.value = val
	v.isSet = true
}

func (v NullableImportPackage) IsSet() bool {
	return v.isSet
}

func (v *NullableImportPackage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportPackage(val *ImportPackage) *NullableImportPackage {
	return &NullableImportPackage{value: val, isSet: true}
}

func (v NullableImportPackage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportPackage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


