/*
Anchore Engine API Server

This is the Anchore Engine API. Provides the primary external API for users of the service.

API version: 0.7.0
Contact: nurmi@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package engine

import (
	"encoding/json"
)

// VulnerablePackageReference A record of a software item which is vulnerable or carries a fix for a vulnerability
type VulnerablePackageReference struct {
	// Package name
	Name *string `json:"name,omitempty"`
	// A version for the package. If null, then references all versions
	Version NullableString `json:"version,omitempty"`
	// Package type (e.g. package, rpm, deb, apk, jar, npm, gem, ...)
	Type *string `json:"type,omitempty"`
	// Severity of vulnerability affecting package
	Severity *string `json:"severity,omitempty"`
	// Vulnerability namespace of affected package
	Namespace *string `json:"namespace,omitempty"`
}

// NewVulnerablePackageReference instantiates a new VulnerablePackageReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVulnerablePackageReference() *VulnerablePackageReference {
	this := VulnerablePackageReference{}
	return &this
}

// NewVulnerablePackageReferenceWithDefaults instantiates a new VulnerablePackageReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVulnerablePackageReferenceWithDefaults() *VulnerablePackageReference {
	this := VulnerablePackageReference{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VulnerablePackageReference) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnerablePackageReference) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VulnerablePackageReference) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VulnerablePackageReference) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VulnerablePackageReference) GetVersion() string {
	if o == nil || o.Version.Get() == nil {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VulnerablePackageReference) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *VulnerablePackageReference) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *VulnerablePackageReference) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *VulnerablePackageReference) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *VulnerablePackageReference) UnsetVersion() {
	o.Version.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VulnerablePackageReference) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnerablePackageReference) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VulnerablePackageReference) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VulnerablePackageReference) SetType(v string) {
	o.Type = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *VulnerablePackageReference) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnerablePackageReference) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *VulnerablePackageReference) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *VulnerablePackageReference) SetSeverity(v string) {
	o.Severity = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *VulnerablePackageReference) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnerablePackageReference) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *VulnerablePackageReference) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *VulnerablePackageReference) SetNamespace(v string) {
	o.Namespace = &v
}

func (o VulnerablePackageReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	return json.Marshal(toSerialize)
}

type NullableVulnerablePackageReference struct {
	value *VulnerablePackageReference
	isSet bool
}

func (v NullableVulnerablePackageReference) Get() *VulnerablePackageReference {
	return v.value
}

func (v *NullableVulnerablePackageReference) Set(val *VulnerablePackageReference) {
	v.value = val
	v.isSet = true
}

func (v NullableVulnerablePackageReference) IsSet() bool {
	return v.isSet
}

func (v *NullableVulnerablePackageReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVulnerablePackageReference(val *VulnerablePackageReference) *NullableVulnerablePackageReference {
	return &NullableVulnerablePackageReference{value: val, isSet: true}
}

func (v NullableVulnerablePackageReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVulnerablePackageReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


