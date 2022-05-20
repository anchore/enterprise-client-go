/*
Anchore Enterprise API Server

This is the Anchore Enterprise API. It provides additional external API routes and functionality for enterprise users.

API version: 0.2.1
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// Application A representation of an SLDC application
type Application struct {
	// The id of the application
	ApplicationId *string `json:"application_id,omitempty"`
	// The name of the application
	Name *string `json:"name,omitempty"`
	// The description of the application
	Description *string `json:"description,omitempty"`
	// List of application versions
	ApplicationVersions *[]ApplicationVersion `json:"application_versions,omitempty"`
	// RFC 3339 formatted UTC timestamp when the application was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// RFC 3339 formatted UTC timestamp when the application was last updated
	LastUpdated *time.Time `json:"last_updated,omitempty"`
}

// NewApplication instantiates a new Application object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplication() *Application {
	this := Application{}
	return &this
}

// NewApplicationWithDefaults instantiates a new Application object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationWithDefaults() *Application {
	this := Application{}
	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *Application) GetApplicationId() string {
	if o == nil || o.ApplicationId == nil {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetApplicationIdOk() (*string, bool) {
	if o == nil || o.ApplicationId == nil {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *Application) HasApplicationId() bool {
	if o != nil && o.ApplicationId != nil {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *Application) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Application) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Application) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Application) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Application) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Application) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Application) SetDescription(v string) {
	o.Description = &v
}

// GetApplicationVersions returns the ApplicationVersions field value if set, zero value otherwise.
func (o *Application) GetApplicationVersions() []ApplicationVersion {
	if o == nil || o.ApplicationVersions == nil {
		var ret []ApplicationVersion
		return ret
	}
	return *o.ApplicationVersions
}

// GetApplicationVersionsOk returns a tuple with the ApplicationVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetApplicationVersionsOk() (*[]ApplicationVersion, bool) {
	if o == nil || o.ApplicationVersions == nil {
		return nil, false
	}
	return o.ApplicationVersions, true
}

// HasApplicationVersions returns a boolean if a field has been set.
func (o *Application) HasApplicationVersions() bool {
	if o != nil && o.ApplicationVersions != nil {
		return true
	}

	return false
}

// SetApplicationVersions gets a reference to the given []ApplicationVersion and assigns it to the ApplicationVersions field.
func (o *Application) SetApplicationVersions(v []ApplicationVersion) {
	o.ApplicationVersions = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Application) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Application) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Application) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *Application) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *Application) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *Application) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o Application) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApplicationId != nil {
		toSerialize["application_id"] = o.ApplicationId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ApplicationVersions != nil {
		toSerialize["application_versions"] = o.ApplicationVersions
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	return json.Marshal(toSerialize)
}

type NullableApplication struct {
	value *Application
	isSet bool
}

func (v NullableApplication) Get() *Application {
	return v.value
}

func (v *NullableApplication) Set(val *Application) {
	v.value = val
	v.isSet = true
}

func (v NullableApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplication(val *Application) *NullableApplication {
	return &NullableApplication{value: val, isSet: true}
}

func (v NullableApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


