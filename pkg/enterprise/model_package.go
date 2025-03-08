/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.10.1
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the Package type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Package{}

// Package A normalized and simplified package that can represent any package type
type Package struct {
	Name *string `json:"name,omitempty"`
	Version *string `json:"version,omitempty"`
	Release *string `json:"release,omitempty"`
	Sourcepkg *string `json:"sourcepkg,omitempty"`
	Location *string `json:"location,omitempty"`
	Origin *string `json:"origin,omitempty"`
	Size *int32 `json:"size,omitempty"`
	Licenses []string `json:"licenses,omitempty"`
	// The type of the metadata entry
	MetadataType *string `json:"metadata_type,omitempty"`
	// Package type specific metadata
	Metadata interface{} `json:"metadata,omitempty"`
	// Spec version for java packages
	SpecificationVersion *string `json:"specification_version,omitempty"`
	// Implementation version for java packages
	ImplementationVersion *string `json:"implementation_version,omitempty"`
	// Maven version for java packages
	MavenVersion *string `json:"maven_version,omitempty"`
	// List of CPE strings for this package
	Cpes []string `json:"cpes,omitempty"`
	Purl *string `json:"purl,omitempty"`
}

// NewPackage instantiates a new Package object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackage() *Package {
	this := Package{}
	return &this
}

// NewPackageWithDefaults instantiates a new Package object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageWithDefaults() *Package {
	this := Package{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Package) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Package) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Package) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Package) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Package) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Package) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Package) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Package) SetVersion(v string) {
	o.Version = &v
}

// GetRelease returns the Release field value if set, zero value otherwise.
func (o *Package) GetRelease() string {
	if o == nil || IsNil(o.Release) {
		var ret string
		return ret
	}
	return *o.Release
}

// GetReleaseOk returns a tuple with the Release field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Package) GetReleaseOk() (*string, bool) {
	if o == nil || IsNil(o.Release) {
		return nil, false
	}
	return o.Release, true
}

// HasRelease returns a boolean if a field has been set.
func (o *Package) HasRelease() bool {
	if o != nil && !IsNil(o.Release) {
		return true
	}

	return false
}

// SetRelease gets a reference to the given string and assigns it to the Release field.
func (o *Package) SetRelease(v string) {
	o.Release = &v
}

// GetSourcepkg returns the Sourcepkg field value if set, zero value otherwise.
func (o *Package) GetSourcepkg() string {
	if o == nil || IsNil(o.Sourcepkg) {
		var ret string
		return ret
	}
	return *o.Sourcepkg
}

// GetSourcepkgOk returns a tuple with the Sourcepkg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Package) GetSourcepkgOk() (*string, bool) {
	if o == nil || IsNil(o.Sourcepkg) {
		return nil, false
	}
	return o.Sourcepkg, true
}

// HasSourcepkg returns a boolean if a field has been set.
func (o *Package) HasSourcepkg() bool {
	if o != nil && !IsNil(o.Sourcepkg) {
		return true
	}

	return false
}

// SetSourcepkg gets a reference to the given string and assigns it to the Sourcepkg field.
func (o *Package) SetSourcepkg(v string) {
	o.Sourcepkg = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *Package) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Package) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *Package) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *Package) SetLocation(v string) {
	o.Location = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *Package) GetOrigin() string {
	if o == nil || IsNil(o.Origin) {
		var ret string
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Package) GetOriginOk() (*string, bool) {
	if o == nil || IsNil(o.Origin) {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *Package) HasOrigin() bool {
	if o != nil && !IsNil(o.Origin) {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given string and assigns it to the Origin field.
func (o *Package) SetOrigin(v string) {
	o.Origin = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *Package) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Package) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *Package) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *Package) SetSize(v int32) {
	o.Size = &v
}

// GetLicenses returns the Licenses field value if set, zero value otherwise.
func (o *Package) GetLicenses() []string {
	if o == nil || IsNil(o.Licenses) {
		var ret []string
		return ret
	}
	return o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Package) GetLicensesOk() ([]string, bool) {
	if o == nil || IsNil(o.Licenses) {
		return nil, false
	}
	return o.Licenses, true
}

// HasLicenses returns a boolean if a field has been set.
func (o *Package) HasLicenses() bool {
	if o != nil && !IsNil(o.Licenses) {
		return true
	}

	return false
}

// SetLicenses gets a reference to the given []string and assigns it to the Licenses field.
func (o *Package) SetLicenses(v []string) {
	o.Licenses = v
}

// GetMetadataType returns the MetadataType field value if set, zero value otherwise.
func (o *Package) GetMetadataType() string {
	if o == nil || IsNil(o.MetadataType) {
		var ret string
		return ret
	}
	return *o.MetadataType
}

// GetMetadataTypeOk returns a tuple with the MetadataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Package) GetMetadataTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MetadataType) {
		return nil, false
	}
	return o.MetadataType, true
}

// HasMetadataType returns a boolean if a field has been set.
func (o *Package) HasMetadataType() bool {
	if o != nil && !IsNil(o.MetadataType) {
		return true
	}

	return false
}

// SetMetadataType gets a reference to the given string and assigns it to the MetadataType field.
func (o *Package) SetMetadataType(v string) {
	o.MetadataType = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Package) GetMetadata() interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Package) GetMetadataOk() (interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Package) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *Package) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetSpecificationVersion returns the SpecificationVersion field value if set, zero value otherwise.
func (o *Package) GetSpecificationVersion() string {
	if o == nil || IsNil(o.SpecificationVersion) {
		var ret string
		return ret
	}
	return *o.SpecificationVersion
}

// GetSpecificationVersionOk returns a tuple with the SpecificationVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Package) GetSpecificationVersionOk() (*string, bool) {
	if o == nil || IsNil(o.SpecificationVersion) {
		return nil, false
	}
	return o.SpecificationVersion, true
}

// HasSpecificationVersion returns a boolean if a field has been set.
func (o *Package) HasSpecificationVersion() bool {
	if o != nil && !IsNil(o.SpecificationVersion) {
		return true
	}

	return false
}

// SetSpecificationVersion gets a reference to the given string and assigns it to the SpecificationVersion field.
func (o *Package) SetSpecificationVersion(v string) {
	o.SpecificationVersion = &v
}

// GetImplementationVersion returns the ImplementationVersion field value if set, zero value otherwise.
func (o *Package) GetImplementationVersion() string {
	if o == nil || IsNil(o.ImplementationVersion) {
		var ret string
		return ret
	}
	return *o.ImplementationVersion
}

// GetImplementationVersionOk returns a tuple with the ImplementationVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Package) GetImplementationVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ImplementationVersion) {
		return nil, false
	}
	return o.ImplementationVersion, true
}

// HasImplementationVersion returns a boolean if a field has been set.
func (o *Package) HasImplementationVersion() bool {
	if o != nil && !IsNil(o.ImplementationVersion) {
		return true
	}

	return false
}

// SetImplementationVersion gets a reference to the given string and assigns it to the ImplementationVersion field.
func (o *Package) SetImplementationVersion(v string) {
	o.ImplementationVersion = &v
}

// GetMavenVersion returns the MavenVersion field value if set, zero value otherwise.
func (o *Package) GetMavenVersion() string {
	if o == nil || IsNil(o.MavenVersion) {
		var ret string
		return ret
	}
	return *o.MavenVersion
}

// GetMavenVersionOk returns a tuple with the MavenVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Package) GetMavenVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MavenVersion) {
		return nil, false
	}
	return o.MavenVersion, true
}

// HasMavenVersion returns a boolean if a field has been set.
func (o *Package) HasMavenVersion() bool {
	if o != nil && !IsNil(o.MavenVersion) {
		return true
	}

	return false
}

// SetMavenVersion gets a reference to the given string and assigns it to the MavenVersion field.
func (o *Package) SetMavenVersion(v string) {
	o.MavenVersion = &v
}

// GetCpes returns the Cpes field value if set, zero value otherwise.
func (o *Package) GetCpes() []string {
	if o == nil || IsNil(o.Cpes) {
		var ret []string
		return ret
	}
	return o.Cpes
}

// GetCpesOk returns a tuple with the Cpes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Package) GetCpesOk() ([]string, bool) {
	if o == nil || IsNil(o.Cpes) {
		return nil, false
	}
	return o.Cpes, true
}

// HasCpes returns a boolean if a field has been set.
func (o *Package) HasCpes() bool {
	if o != nil && !IsNil(o.Cpes) {
		return true
	}

	return false
}

// SetCpes gets a reference to the given []string and assigns it to the Cpes field.
func (o *Package) SetCpes(v []string) {
	o.Cpes = v
}

// GetPurl returns the Purl field value if set, zero value otherwise.
func (o *Package) GetPurl() string {
	if o == nil || IsNil(o.Purl) {
		var ret string
		return ret
	}
	return *o.Purl
}

// GetPurlOk returns a tuple with the Purl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Package) GetPurlOk() (*string, bool) {
	if o == nil || IsNil(o.Purl) {
		return nil, false
	}
	return o.Purl, true
}

// HasPurl returns a boolean if a field has been set.
func (o *Package) HasPurl() bool {
	if o != nil && !IsNil(o.Purl) {
		return true
	}

	return false
}

// SetPurl gets a reference to the given string and assigns it to the Purl field.
func (o *Package) SetPurl(v string) {
	o.Purl = &v
}

func (o Package) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Package) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Release) {
		toSerialize["release"] = o.Release
	}
	if !IsNil(o.Sourcepkg) {
		toSerialize["sourcepkg"] = o.Sourcepkg
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Origin) {
		toSerialize["origin"] = o.Origin
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Licenses) {
		toSerialize["licenses"] = o.Licenses
	}
	if !IsNil(o.MetadataType) {
		toSerialize["metadata_type"] = o.MetadataType
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.SpecificationVersion) {
		toSerialize["specification_version"] = o.SpecificationVersion
	}
	if !IsNil(o.ImplementationVersion) {
		toSerialize["implementation_version"] = o.ImplementationVersion
	}
	if !IsNil(o.MavenVersion) {
		toSerialize["maven_version"] = o.MavenVersion
	}
	if !IsNil(o.Cpes) {
		toSerialize["cpes"] = o.Cpes
	}
	if !IsNil(o.Purl) {
		toSerialize["purl"] = o.Purl
	}
	return toSerialize, nil
}

type NullablePackage struct {
	value *Package
	isSet bool
}

func (v NullablePackage) Get() *Package {
	return v.value
}

func (v *NullablePackage) Set(val *Package) {
	v.value = val
	v.isSet = true
}

func (v NullablePackage) IsSet() bool {
	return v.isSet
}

func (v *NullablePackage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackage(val *Package) *NullablePackage {
	return &NullablePackage{value: val, isSet: true}
}

func (v NullablePackage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


