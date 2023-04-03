/*
Anchore Enterprise API Server

This is the Anchore Enterprise API. It provides additional external API routes and functionality for enterprise users.

API version: 0.7.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// ApplicationVersion A representation of an SLDC application
type ApplicationVersion struct {
	// The id of the application version
	ApplicationVersionId *string `json:"application_version_id,omitempty"`
	// The id of the application
	ApplicationId *string `json:"application_id,omitempty"`
	// The name of the application version. The name must be unique per application
	VersionName string `json:"version_name"`
	// RFC 3339 formatted UTC timestamp when the application was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// RFC 3339 formatted UTC timestamp when the application was last updated
	LastUpdated *time.Time `json:"last_updated,omitempty"`
}

// NewApplicationVersion instantiates a new ApplicationVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationVersion(versionName string) *ApplicationVersion {
	this := ApplicationVersion{}
	this.VersionName = versionName
	return &this
}

// NewApplicationVersionWithDefaults instantiates a new ApplicationVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationVersionWithDefaults() *ApplicationVersion {
	this := ApplicationVersion{}
	return &this
}

// GetApplicationVersionId returns the ApplicationVersionId field value if set, zero value otherwise.
func (o *ApplicationVersion) GetApplicationVersionId() string {
	if o == nil || o.ApplicationVersionId == nil {
		var ret string
		return ret
	}
	return *o.ApplicationVersionId
}

// GetApplicationVersionIdOk returns a tuple with the ApplicationVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationVersion) GetApplicationVersionIdOk() (*string, bool) {
	if o == nil || o.ApplicationVersionId == nil {
		return nil, false
	}
	return o.ApplicationVersionId, true
}

// HasApplicationVersionId returns a boolean if a field has been set.
func (o *ApplicationVersion) HasApplicationVersionId() bool {
	if o != nil && o.ApplicationVersionId != nil {
		return true
	}

	return false
}

// SetApplicationVersionId gets a reference to the given string and assigns it to the ApplicationVersionId field.
func (o *ApplicationVersion) SetApplicationVersionId(v string) {
	o.ApplicationVersionId = &v
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *ApplicationVersion) GetApplicationId() string {
	if o == nil || o.ApplicationId == nil {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationVersion) GetApplicationIdOk() (*string, bool) {
	if o == nil || o.ApplicationId == nil {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *ApplicationVersion) HasApplicationId() bool {
	if o != nil && o.ApplicationId != nil {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *ApplicationVersion) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetVersionName returns the VersionName field value
func (o *ApplicationVersion) GetVersionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionName
}

// GetVersionNameOk returns a tuple with the VersionName field value
// and a boolean to check if the value has been set.
func (o *ApplicationVersion) GetVersionNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VersionName, true
}

// SetVersionName sets field value
func (o *ApplicationVersion) SetVersionName(v string) {
	o.VersionName = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ApplicationVersion) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationVersion) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ApplicationVersion) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ApplicationVersion) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *ApplicationVersion) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationVersion) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *ApplicationVersion) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *ApplicationVersion) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o ApplicationVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApplicationVersionId != nil {
		toSerialize["application_version_id"] = o.ApplicationVersionId
	}
	if o.ApplicationId != nil {
		toSerialize["application_id"] = o.ApplicationId
	}
	if true {
		toSerialize["version_name"] = o.VersionName
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationVersion struct {
	value *ApplicationVersion
	isSet bool
}

func (v NullableApplicationVersion) Get() *ApplicationVersion {
	return v.value
}

func (v *NullableApplicationVersion) Set(val *ApplicationVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationVersion(val *ApplicationVersion) *NullableApplicationVersion {
	return &NullableApplicationVersion{value: val, isSet: true}
}

func (v NullableApplicationVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


