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
)

// checks if the ImportDistribution type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportDistribution{}

// ImportDistribution struct for ImportDistribution
type ImportDistribution struct {
	Name NullableString `json:"name,omitempty"`
	Id NullableString `json:"id,omitempty"`
	Version NullableString `json:"version,omitempty"`
	VersionID NullableString `json:"versionID,omitempty"`
	IdLike interface{} `json:"idLike,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ImportDistribution ImportDistribution

// NewImportDistribution instantiates a new ImportDistribution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportDistribution() *ImportDistribution {
	this := ImportDistribution{}
	return &this
}

// NewImportDistributionWithDefaults instantiates a new ImportDistribution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportDistributionWithDefaults() *ImportDistribution {
	this := ImportDistribution{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportDistribution) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportDistribution) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ImportDistribution) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ImportDistribution) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ImportDistribution) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ImportDistribution) UnsetName() {
	o.Name.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportDistribution) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportDistribution) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ImportDistribution) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *ImportDistribution) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ImportDistribution) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ImportDistribution) UnsetId() {
	o.Id.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportDistribution) GetVersion() string {
	if o == nil || IsNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportDistribution) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *ImportDistribution) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *ImportDistribution) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *ImportDistribution) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *ImportDistribution) UnsetVersion() {
	o.Version.Unset()
}

// GetVersionID returns the VersionID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportDistribution) GetVersionID() string {
	if o == nil || IsNil(o.VersionID.Get()) {
		var ret string
		return ret
	}
	return *o.VersionID.Get()
}

// GetVersionIDOk returns a tuple with the VersionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportDistribution) GetVersionIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VersionID.Get(), o.VersionID.IsSet()
}

// HasVersionID returns a boolean if a field has been set.
func (o *ImportDistribution) HasVersionID() bool {
	if o != nil && o.VersionID.IsSet() {
		return true
	}

	return false
}

// SetVersionID gets a reference to the given NullableString and assigns it to the VersionID field.
func (o *ImportDistribution) SetVersionID(v string) {
	o.VersionID.Set(&v)
}
// SetVersionIDNil sets the value for VersionID to be an explicit nil
func (o *ImportDistribution) SetVersionIDNil() {
	o.VersionID.Set(nil)
}

// UnsetVersionID ensures that no value is present for VersionID, not even an explicit nil
func (o *ImportDistribution) UnsetVersionID() {
	o.VersionID.Unset()
}

// GetIdLike returns the IdLike field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportDistribution) GetIdLike() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.IdLike
}

// GetIdLikeOk returns a tuple with the IdLike field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportDistribution) GetIdLikeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.IdLike) {
		return nil, false
	}
	return &o.IdLike, true
}

// HasIdLike returns a boolean if a field has been set.
func (o *ImportDistribution) HasIdLike() bool {
	if o != nil && !IsNil(o.IdLike) {
		return true
	}

	return false
}

// SetIdLike gets a reference to the given interface{} and assigns it to the IdLike field.
func (o *ImportDistribution) SetIdLike(v interface{}) {
	o.IdLike = v
}

func (o ImportDistribution) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportDistribution) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if o.VersionID.IsSet() {
		toSerialize["versionID"] = o.VersionID.Get()
	}
	if o.IdLike != nil {
		toSerialize["idLike"] = o.IdLike
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ImportDistribution) UnmarshalJSON(data []byte) (err error) {
	varImportDistribution := _ImportDistribution{}

	err = json.Unmarshal(data, &varImportDistribution)

	if err != nil {
		return err
	}

	*o = ImportDistribution(varImportDistribution)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "id")
		delete(additionalProperties, "version")
		delete(additionalProperties, "versionID")
		delete(additionalProperties, "idLike")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImportDistribution struct {
	value *ImportDistribution
	isSet bool
}

func (v NullableImportDistribution) Get() *ImportDistribution {
	return v.value
}

func (v *NullableImportDistribution) Set(val *ImportDistribution) {
	v.value = val
	v.isSet = true
}

func (v NullableImportDistribution) IsSet() bool {
	return v.isSet
}

func (v *NullableImportDistribution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportDistribution(val *ImportDistribution) *NullableImportDistribution {
	return &NullableImportDistribution{value: val, isSet: true}
}

func (v NullableImportDistribution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportDistribution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


