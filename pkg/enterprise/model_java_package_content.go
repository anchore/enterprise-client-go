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

// JavaPackageContent struct for JavaPackageContent
type JavaPackageContent struct {
	Package *string `json:"package,omitempty"`
	ImplementationVersion *string `json:"implementation-version,omitempty"`
	SpecificationVersion *string `json:"specification-version,omitempty"`
	MavenVersion *string `json:"maven-version,omitempty"`
	Location *string `json:"location,omitempty"`
	Type *string `json:"type,omitempty"`
	Origin *string `json:"origin,omitempty"`
	Licenses []string `json:"licenses,omitempty"`
	// A list of Common Platform Enumerations that may uniquely identify the package
	Cpes []string `json:"cpes,omitempty"`
	Purl *string `json:"purl,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewJavaPackageContent instantiates a new JavaPackageContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJavaPackageContent() *JavaPackageContent {
	this := JavaPackageContent{}
	return &this
}

// NewJavaPackageContentWithDefaults instantiates a new JavaPackageContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJavaPackageContentWithDefaults() *JavaPackageContent {
	this := JavaPackageContent{}
	return &this
}

// GetPackage returns the Package field value if set, zero value otherwise.
func (o *JavaPackageContent) GetPackage() string {
	if o == nil || o.Package == nil {
		var ret string
		return ret
	}
	return *o.Package
}

// GetPackageOk returns a tuple with the Package field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JavaPackageContent) GetPackageOk() (*string, bool) {
	if o == nil || o.Package == nil {
		return nil, false
	}
	return o.Package, true
}

// HasPackage returns a boolean if a field has been set.
func (o *JavaPackageContent) HasPackage() bool {
	if o != nil && o.Package != nil {
		return true
	}

	return false
}

// SetPackage gets a reference to the given string and assigns it to the Package field.
func (o *JavaPackageContent) SetPackage(v string) {
	o.Package = &v
}

// GetImplementationVersion returns the ImplementationVersion field value if set, zero value otherwise.
func (o *JavaPackageContent) GetImplementationVersion() string {
	if o == nil || o.ImplementationVersion == nil {
		var ret string
		return ret
	}
	return *o.ImplementationVersion
}

// GetImplementationVersionOk returns a tuple with the ImplementationVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JavaPackageContent) GetImplementationVersionOk() (*string, bool) {
	if o == nil || o.ImplementationVersion == nil {
		return nil, false
	}
	return o.ImplementationVersion, true
}

// HasImplementationVersion returns a boolean if a field has been set.
func (o *JavaPackageContent) HasImplementationVersion() bool {
	if o != nil && o.ImplementationVersion != nil {
		return true
	}

	return false
}

// SetImplementationVersion gets a reference to the given string and assigns it to the ImplementationVersion field.
func (o *JavaPackageContent) SetImplementationVersion(v string) {
	o.ImplementationVersion = &v
}

// GetSpecificationVersion returns the SpecificationVersion field value if set, zero value otherwise.
func (o *JavaPackageContent) GetSpecificationVersion() string {
	if o == nil || o.SpecificationVersion == nil {
		var ret string
		return ret
	}
	return *o.SpecificationVersion
}

// GetSpecificationVersionOk returns a tuple with the SpecificationVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JavaPackageContent) GetSpecificationVersionOk() (*string, bool) {
	if o == nil || o.SpecificationVersion == nil {
		return nil, false
	}
	return o.SpecificationVersion, true
}

// HasSpecificationVersion returns a boolean if a field has been set.
func (o *JavaPackageContent) HasSpecificationVersion() bool {
	if o != nil && o.SpecificationVersion != nil {
		return true
	}

	return false
}

// SetSpecificationVersion gets a reference to the given string and assigns it to the SpecificationVersion field.
func (o *JavaPackageContent) SetSpecificationVersion(v string) {
	o.SpecificationVersion = &v
}

// GetMavenVersion returns the MavenVersion field value if set, zero value otherwise.
func (o *JavaPackageContent) GetMavenVersion() string {
	if o == nil || o.MavenVersion == nil {
		var ret string
		return ret
	}
	return *o.MavenVersion
}

// GetMavenVersionOk returns a tuple with the MavenVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JavaPackageContent) GetMavenVersionOk() (*string, bool) {
	if o == nil || o.MavenVersion == nil {
		return nil, false
	}
	return o.MavenVersion, true
}

// HasMavenVersion returns a boolean if a field has been set.
func (o *JavaPackageContent) HasMavenVersion() bool {
	if o != nil && o.MavenVersion != nil {
		return true
	}

	return false
}

// SetMavenVersion gets a reference to the given string and assigns it to the MavenVersion field.
func (o *JavaPackageContent) SetMavenVersion(v string) {
	o.MavenVersion = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *JavaPackageContent) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JavaPackageContent) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *JavaPackageContent) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *JavaPackageContent) SetLocation(v string) {
	o.Location = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *JavaPackageContent) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JavaPackageContent) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *JavaPackageContent) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *JavaPackageContent) SetType(v string) {
	o.Type = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *JavaPackageContent) GetOrigin() string {
	if o == nil || o.Origin == nil {
		var ret string
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JavaPackageContent) GetOriginOk() (*string, bool) {
	if o == nil || o.Origin == nil {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *JavaPackageContent) HasOrigin() bool {
	if o != nil && o.Origin != nil {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given string and assigns it to the Origin field.
func (o *JavaPackageContent) SetOrigin(v string) {
	o.Origin = &v
}

// GetLicenses returns the Licenses field value if set, zero value otherwise.
func (o *JavaPackageContent) GetLicenses() []string {
	if o == nil || o.Licenses == nil {
		var ret []string
		return ret
	}
	return o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JavaPackageContent) GetLicensesOk() ([]string, bool) {
	if o == nil || o.Licenses == nil {
		return nil, false
	}
	return o.Licenses, true
}

// HasLicenses returns a boolean if a field has been set.
func (o *JavaPackageContent) HasLicenses() bool {
	if o != nil && o.Licenses != nil {
		return true
	}

	return false
}

// SetLicenses gets a reference to the given []string and assigns it to the Licenses field.
func (o *JavaPackageContent) SetLicenses(v []string) {
	o.Licenses = v
}

// GetCpes returns the Cpes field value if set, zero value otherwise.
func (o *JavaPackageContent) GetCpes() []string {
	if o == nil || o.Cpes == nil {
		var ret []string
		return ret
	}
	return o.Cpes
}

// GetCpesOk returns a tuple with the Cpes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JavaPackageContent) GetCpesOk() ([]string, bool) {
	if o == nil || o.Cpes == nil {
		return nil, false
	}
	return o.Cpes, true
}

// HasCpes returns a boolean if a field has been set.
func (o *JavaPackageContent) HasCpes() bool {
	if o != nil && o.Cpes != nil {
		return true
	}

	return false
}

// SetCpes gets a reference to the given []string and assigns it to the Cpes field.
func (o *JavaPackageContent) SetCpes(v []string) {
	o.Cpes = v
}

// GetPurl returns the Purl field value if set, zero value otherwise.
func (o *JavaPackageContent) GetPurl() string {
	if o == nil || o.Purl == nil {
		var ret string
		return ret
	}
	return *o.Purl
}

// GetPurlOk returns a tuple with the Purl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JavaPackageContent) GetPurlOk() (*string, bool) {
	if o == nil || o.Purl == nil {
		return nil, false
	}
	return o.Purl, true
}

// HasPurl returns a boolean if a field has been set.
func (o *JavaPackageContent) HasPurl() bool {
	if o != nil && o.Purl != nil {
		return true
	}

	return false
}

// SetPurl gets a reference to the given string and assigns it to the Purl field.
func (o *JavaPackageContent) SetPurl(v string) {
	o.Purl = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *JavaPackageContent) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JavaPackageContent) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *JavaPackageContent) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *JavaPackageContent) SetVersion(v string) {
	o.Version = &v
}

func (o JavaPackageContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Package != nil {
		toSerialize["package"] = o.Package
	}
	if o.ImplementationVersion != nil {
		toSerialize["implementation-version"] = o.ImplementationVersion
	}
	if o.SpecificationVersion != nil {
		toSerialize["specification-version"] = o.SpecificationVersion
	}
	if o.MavenVersion != nil {
		toSerialize["maven-version"] = o.MavenVersion
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
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
	if o.Cpes != nil {
		toSerialize["cpes"] = o.Cpes
	}
	if o.Purl != nil {
		toSerialize["purl"] = o.Purl
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableJavaPackageContent struct {
	value *JavaPackageContent
	isSet bool
}

func (v NullableJavaPackageContent) Get() *JavaPackageContent {
	return v.value
}

func (v *NullableJavaPackageContent) Set(val *JavaPackageContent) {
	v.value = val
	v.isSet = true
}

func (v NullableJavaPackageContent) IsSet() bool {
	return v.isSet
}

func (v *NullableJavaPackageContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJavaPackageContent(val *JavaPackageContent) *NullableJavaPackageContent {
	return &NullableJavaPackageContent{value: val, isSet: true}
}

func (v NullableJavaPackageContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJavaPackageContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


