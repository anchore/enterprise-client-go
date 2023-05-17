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
	"time"
)

// ApplicationVersionSbom A combined sbom for the artifacts associated with an application version
type ApplicationVersionSbom struct {
	Application *Application `json:"application,omitempty"`
	ApplicationVersion *ApplicationVersion `json:"application_version,omitempty"`
	// RFC 3339 formatted UTC timestamp when the application version sbom was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	SourceSboms *[]interface{} `json:"source_sboms,omitempty"`
	ImageSboms *[]interface{} `json:"image_sboms,omitempty"`
}

// NewApplicationVersionSbom instantiates a new ApplicationVersionSbom object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationVersionSbom() *ApplicationVersionSbom {
	this := ApplicationVersionSbom{}
	return &this
}

// NewApplicationVersionSbomWithDefaults instantiates a new ApplicationVersionSbom object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationVersionSbomWithDefaults() *ApplicationVersionSbom {
	this := ApplicationVersionSbom{}
	return &this
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *ApplicationVersionSbom) GetApplication() Application {
	if o == nil || o.Application == nil {
		var ret Application
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationVersionSbom) GetApplicationOk() (*Application, bool) {
	if o == nil || o.Application == nil {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *ApplicationVersionSbom) HasApplication() bool {
	if o != nil && o.Application != nil {
		return true
	}

	return false
}

// SetApplication gets a reference to the given Application and assigns it to the Application field.
func (o *ApplicationVersionSbom) SetApplication(v Application) {
	o.Application = &v
}

// GetApplicationVersion returns the ApplicationVersion field value if set, zero value otherwise.
func (o *ApplicationVersionSbom) GetApplicationVersion() ApplicationVersion {
	if o == nil || o.ApplicationVersion == nil {
		var ret ApplicationVersion
		return ret
	}
	return *o.ApplicationVersion
}

// GetApplicationVersionOk returns a tuple with the ApplicationVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationVersionSbom) GetApplicationVersionOk() (*ApplicationVersion, bool) {
	if o == nil || o.ApplicationVersion == nil {
		return nil, false
	}
	return o.ApplicationVersion, true
}

// HasApplicationVersion returns a boolean if a field has been set.
func (o *ApplicationVersionSbom) HasApplicationVersion() bool {
	if o != nil && o.ApplicationVersion != nil {
		return true
	}

	return false
}

// SetApplicationVersion gets a reference to the given ApplicationVersion and assigns it to the ApplicationVersion field.
func (o *ApplicationVersionSbom) SetApplicationVersion(v ApplicationVersion) {
	o.ApplicationVersion = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ApplicationVersionSbom) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationVersionSbom) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ApplicationVersionSbom) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ApplicationVersionSbom) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetSourceSboms returns the SourceSboms field value if set, zero value otherwise.
func (o *ApplicationVersionSbom) GetSourceSboms() []interface{} {
	if o == nil || o.SourceSboms == nil {
		var ret []interface{}
		return ret
	}
	return *o.SourceSboms
}

// GetSourceSbomsOk returns a tuple with the SourceSboms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationVersionSbom) GetSourceSbomsOk() (*[]interface{}, bool) {
	if o == nil || o.SourceSboms == nil {
		return nil, false
	}
	return o.SourceSboms, true
}

// HasSourceSboms returns a boolean if a field has been set.
func (o *ApplicationVersionSbom) HasSourceSboms() bool {
	if o != nil && o.SourceSboms != nil {
		return true
	}

	return false
}

// SetSourceSboms gets a reference to the given []interface{} and assigns it to the SourceSboms field.
func (o *ApplicationVersionSbom) SetSourceSboms(v []interface{}) {
	o.SourceSboms = &v
}

// GetImageSboms returns the ImageSboms field value if set, zero value otherwise.
func (o *ApplicationVersionSbom) GetImageSboms() []interface{} {
	if o == nil || o.ImageSboms == nil {
		var ret []interface{}
		return ret
	}
	return *o.ImageSboms
}

// GetImageSbomsOk returns a tuple with the ImageSboms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationVersionSbom) GetImageSbomsOk() (*[]interface{}, bool) {
	if o == nil || o.ImageSboms == nil {
		return nil, false
	}
	return o.ImageSboms, true
}

// HasImageSboms returns a boolean if a field has been set.
func (o *ApplicationVersionSbom) HasImageSboms() bool {
	if o != nil && o.ImageSboms != nil {
		return true
	}

	return false
}

// SetImageSboms gets a reference to the given []interface{} and assigns it to the ImageSboms field.
func (o *ApplicationVersionSbom) SetImageSboms(v []interface{}) {
	o.ImageSboms = &v
}

func (o ApplicationVersionSbom) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Application != nil {
		toSerialize["application"] = o.Application
	}
	if o.ApplicationVersion != nil {
		toSerialize["application_version"] = o.ApplicationVersion
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.SourceSboms != nil {
		toSerialize["source_sboms"] = o.SourceSboms
	}
	if o.ImageSboms != nil {
		toSerialize["image_sboms"] = o.ImageSboms
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationVersionSbom struct {
	value *ApplicationVersionSbom
	isSet bool
}

func (v NullableApplicationVersionSbom) Get() *ApplicationVersionSbom {
	return v.value
}

func (v *NullableApplicationVersionSbom) Set(val *ApplicationVersionSbom) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationVersionSbom) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationVersionSbom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationVersionSbom(val *ApplicationVersionSbom) *NullableApplicationVersionSbom {
	return &NullableApplicationVersionSbom{value: val, isSet: true}
}

func (v NullableApplicationVersionSbom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationVersionSbom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


