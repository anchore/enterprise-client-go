/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.10.2
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"fmt"
)

// checks if the NativeSBOMPackage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NativeSBOMPackage{}

// NativeSBOMPackage struct for NativeSBOMPackage
type NativeSBOMPackage struct {
	Id *string `json:"id,omitempty"`
	Name string `json:"name"`
	Version string `json:"version"`
	Type string `json:"type"`
	FoundBy *string `json:"foundBy,omitempty"`
	Locations []NativeSBOMPackageLocation `json:"locations"`
	Licenses []ImportPackageLicensesInner `json:"licenses"`
	Language string `json:"language"`
	Cpes []NativeSBOMPackageCpesInner `json:"cpes"`
	Purl *string `json:"purl,omitempty"`
	MetadataType NullableString `json:"metadataType,omitempty"`
	Metadata interface{} `json:"metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NativeSBOMPackage NativeSBOMPackage

// NewNativeSBOMPackage instantiates a new NativeSBOMPackage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNativeSBOMPackage(name string, version string, type_ string, locations []NativeSBOMPackageLocation, licenses []ImportPackageLicensesInner, language string, cpes []NativeSBOMPackageCpesInner) *NativeSBOMPackage {
	this := NativeSBOMPackage{}
	this.Name = name
	this.Version = version
	this.Type = type_
	this.Locations = locations
	this.Licenses = licenses
	this.Language = language
	this.Cpes = cpes
	return &this
}

// NewNativeSBOMPackageWithDefaults instantiates a new NativeSBOMPackage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNativeSBOMPackageWithDefaults() *NativeSBOMPackage {
	this := NativeSBOMPackage{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NativeSBOMPackage) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NativeSBOMPackage) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NativeSBOMPackage) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NativeSBOMPackage) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *NativeSBOMPackage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NativeSBOMPackage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NativeSBOMPackage) SetName(v string) {
	o.Name = v
}

// GetVersion returns the Version field value
func (o *NativeSBOMPackage) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *NativeSBOMPackage) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *NativeSBOMPackage) SetVersion(v string) {
	o.Version = v
}

// GetType returns the Type field value
func (o *NativeSBOMPackage) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NativeSBOMPackage) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NativeSBOMPackage) SetType(v string) {
	o.Type = v
}

// GetFoundBy returns the FoundBy field value if set, zero value otherwise.
func (o *NativeSBOMPackage) GetFoundBy() string {
	if o == nil || IsNil(o.FoundBy) {
		var ret string
		return ret
	}
	return *o.FoundBy
}

// GetFoundByOk returns a tuple with the FoundBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NativeSBOMPackage) GetFoundByOk() (*string, bool) {
	if o == nil || IsNil(o.FoundBy) {
		return nil, false
	}
	return o.FoundBy, true
}

// HasFoundBy returns a boolean if a field has been set.
func (o *NativeSBOMPackage) HasFoundBy() bool {
	if o != nil && !IsNil(o.FoundBy) {
		return true
	}

	return false
}

// SetFoundBy gets a reference to the given string and assigns it to the FoundBy field.
func (o *NativeSBOMPackage) SetFoundBy(v string) {
	o.FoundBy = &v
}

// GetLocations returns the Locations field value
func (o *NativeSBOMPackage) GetLocations() []NativeSBOMPackageLocation {
	if o == nil {
		var ret []NativeSBOMPackageLocation
		return ret
	}

	return o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value
// and a boolean to check if the value has been set.
func (o *NativeSBOMPackage) GetLocationsOk() ([]NativeSBOMPackageLocation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Locations, true
}

// SetLocations sets field value
func (o *NativeSBOMPackage) SetLocations(v []NativeSBOMPackageLocation) {
	o.Locations = v
}

// GetLicenses returns the Licenses field value
func (o *NativeSBOMPackage) GetLicenses() []ImportPackageLicensesInner {
	if o == nil {
		var ret []ImportPackageLicensesInner
		return ret
	}

	return o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value
// and a boolean to check if the value has been set.
func (o *NativeSBOMPackage) GetLicensesOk() ([]ImportPackageLicensesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Licenses, true
}

// SetLicenses sets field value
func (o *NativeSBOMPackage) SetLicenses(v []ImportPackageLicensesInner) {
	o.Licenses = v
}

// GetLanguage returns the Language field value
func (o *NativeSBOMPackage) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *NativeSBOMPackage) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *NativeSBOMPackage) SetLanguage(v string) {
	o.Language = v
}

// GetCpes returns the Cpes field value
func (o *NativeSBOMPackage) GetCpes() []NativeSBOMPackageCpesInner {
	if o == nil {
		var ret []NativeSBOMPackageCpesInner
		return ret
	}

	return o.Cpes
}

// GetCpesOk returns a tuple with the Cpes field value
// and a boolean to check if the value has been set.
func (o *NativeSBOMPackage) GetCpesOk() ([]NativeSBOMPackageCpesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cpes, true
}

// SetCpes sets field value
func (o *NativeSBOMPackage) SetCpes(v []NativeSBOMPackageCpesInner) {
	o.Cpes = v
}

// GetPurl returns the Purl field value if set, zero value otherwise.
func (o *NativeSBOMPackage) GetPurl() string {
	if o == nil || IsNil(o.Purl) {
		var ret string
		return ret
	}
	return *o.Purl
}

// GetPurlOk returns a tuple with the Purl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NativeSBOMPackage) GetPurlOk() (*string, bool) {
	if o == nil || IsNil(o.Purl) {
		return nil, false
	}
	return o.Purl, true
}

// HasPurl returns a boolean if a field has been set.
func (o *NativeSBOMPackage) HasPurl() bool {
	if o != nil && !IsNil(o.Purl) {
		return true
	}

	return false
}

// SetPurl gets a reference to the given string and assigns it to the Purl field.
func (o *NativeSBOMPackage) SetPurl(v string) {
	o.Purl = &v
}

// GetMetadataType returns the MetadataType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NativeSBOMPackage) GetMetadataType() string {
	if o == nil || IsNil(o.MetadataType.Get()) {
		var ret string
		return ret
	}
	return *o.MetadataType.Get()
}

// GetMetadataTypeOk returns a tuple with the MetadataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NativeSBOMPackage) GetMetadataTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetadataType.Get(), o.MetadataType.IsSet()
}

// HasMetadataType returns a boolean if a field has been set.
func (o *NativeSBOMPackage) HasMetadataType() bool {
	if o != nil && o.MetadataType.IsSet() {
		return true
	}

	return false
}

// SetMetadataType gets a reference to the given NullableString and assigns it to the MetadataType field.
func (o *NativeSBOMPackage) SetMetadataType(v string) {
	o.MetadataType.Set(&v)
}
// SetMetadataTypeNil sets the value for MetadataType to be an explicit nil
func (o *NativeSBOMPackage) SetMetadataTypeNil() {
	o.MetadataType.Set(nil)
}

// UnsetMetadataType ensures that no value is present for MetadataType, not even an explicit nil
func (o *NativeSBOMPackage) UnsetMetadataType() {
	o.MetadataType.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NativeSBOMPackage) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NativeSBOMPackage) GetMetadataOk() (interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *NativeSBOMPackage) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *NativeSBOMPackage) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o NativeSBOMPackage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NativeSBOMPackage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	toSerialize["version"] = o.Version
	toSerialize["type"] = o.Type
	if !IsNil(o.FoundBy) {
		toSerialize["foundBy"] = o.FoundBy
	}
	toSerialize["locations"] = o.Locations
	toSerialize["licenses"] = o.Licenses
	toSerialize["language"] = o.Language
	toSerialize["cpes"] = o.Cpes
	if !IsNil(o.Purl) {
		toSerialize["purl"] = o.Purl
	}
	if o.MetadataType.IsSet() {
		toSerialize["metadataType"] = o.MetadataType.Get()
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NativeSBOMPackage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"version",
		"type",
		"locations",
		"licenses",
		"language",
		"cpes",
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

	varNativeSBOMPackage := _NativeSBOMPackage{}

	err = json.Unmarshal(data, &varNativeSBOMPackage)

	if err != nil {
		return err
	}

	*o = NativeSBOMPackage(varNativeSBOMPackage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "version")
		delete(additionalProperties, "type")
		delete(additionalProperties, "foundBy")
		delete(additionalProperties, "locations")
		delete(additionalProperties, "licenses")
		delete(additionalProperties, "language")
		delete(additionalProperties, "cpes")
		delete(additionalProperties, "purl")
		delete(additionalProperties, "metadataType")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNativeSBOMPackage struct {
	value *NativeSBOMPackage
	isSet bool
}

func (v NullableNativeSBOMPackage) Get() *NativeSBOMPackage {
	return v.value
}

func (v *NullableNativeSBOMPackage) Set(val *NativeSBOMPackage) {
	v.value = val
	v.isSet = true
}

func (v NullableNativeSBOMPackage) IsSet() bool {
	return v.isSet
}

func (v *NullableNativeSBOMPackage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNativeSBOMPackage(val *NativeSBOMPackage) *NullableNativeSBOMPackage {
	return &NullableNativeSBOMPackage{value: val, isSet: true}
}

func (v NullableNativeSBOMPackage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNativeSBOMPackage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


