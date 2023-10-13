/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.0.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// ActionPlan describes a remediation action plan object
type ActionPlan struct {
	Type *string `json:"type,omitempty"`
	ImageTag *string `json:"image_tag,omitempty"`
	ImageDigest *string `json:"image_digest,omitempty"`
	PolicyId *string `json:"policy_id,omitempty"`
	Resolutions []ActionPlanResolution `json:"resolutions,omitempty"`
	Endpoint *string `json:"endpoint,omitempty"`
	ConfigurationId *string `json:"configuration_id,omitempty"`
	Subject *string `json:"subject,omitempty"`
	Message *string `json:"message,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
}

// NewActionPlan instantiates a new ActionPlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionPlan() *ActionPlan {
	this := ActionPlan{}
	return &this
}

// NewActionPlanWithDefaults instantiates a new ActionPlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionPlanWithDefaults() *ActionPlan {
	this := ActionPlan{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ActionPlan) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionPlan) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ActionPlan) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ActionPlan) SetType(v string) {
	o.Type = &v
}

// GetImageTag returns the ImageTag field value if set, zero value otherwise.
func (o *ActionPlan) GetImageTag() string {
	if o == nil || o.ImageTag == nil {
		var ret string
		return ret
	}
	return *o.ImageTag
}

// GetImageTagOk returns a tuple with the ImageTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionPlan) GetImageTagOk() (*string, bool) {
	if o == nil || o.ImageTag == nil {
		return nil, false
	}
	return o.ImageTag, true
}

// HasImageTag returns a boolean if a field has been set.
func (o *ActionPlan) HasImageTag() bool {
	if o != nil && o.ImageTag != nil {
		return true
	}

	return false
}

// SetImageTag gets a reference to the given string and assigns it to the ImageTag field.
func (o *ActionPlan) SetImageTag(v string) {
	o.ImageTag = &v
}

// GetImageDigest returns the ImageDigest field value if set, zero value otherwise.
func (o *ActionPlan) GetImageDigest() string {
	if o == nil || o.ImageDigest == nil {
		var ret string
		return ret
	}
	return *o.ImageDigest
}

// GetImageDigestOk returns a tuple with the ImageDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionPlan) GetImageDigestOk() (*string, bool) {
	if o == nil || o.ImageDigest == nil {
		return nil, false
	}
	return o.ImageDigest, true
}

// HasImageDigest returns a boolean if a field has been set.
func (o *ActionPlan) HasImageDigest() bool {
	if o != nil && o.ImageDigest != nil {
		return true
	}

	return false
}

// SetImageDigest gets a reference to the given string and assigns it to the ImageDigest field.
func (o *ActionPlan) SetImageDigest(v string) {
	o.ImageDigest = &v
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *ActionPlan) GetPolicyId() string {
	if o == nil || o.PolicyId == nil {
		var ret string
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionPlan) GetPolicyIdOk() (*string, bool) {
	if o == nil || o.PolicyId == nil {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *ActionPlan) HasPolicyId() bool {
	if o != nil && o.PolicyId != nil {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given string and assigns it to the PolicyId field.
func (o *ActionPlan) SetPolicyId(v string) {
	o.PolicyId = &v
}

// GetResolutions returns the Resolutions field value if set, zero value otherwise.
func (o *ActionPlan) GetResolutions() []ActionPlanResolution {
	if o == nil || o.Resolutions == nil {
		var ret []ActionPlanResolution
		return ret
	}
	return o.Resolutions
}

// GetResolutionsOk returns a tuple with the Resolutions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionPlan) GetResolutionsOk() ([]ActionPlanResolution, bool) {
	if o == nil || o.Resolutions == nil {
		return nil, false
	}
	return o.Resolutions, true
}

// HasResolutions returns a boolean if a field has been set.
func (o *ActionPlan) HasResolutions() bool {
	if o != nil && o.Resolutions != nil {
		return true
	}

	return false
}

// SetResolutions gets a reference to the given []ActionPlanResolution and assigns it to the Resolutions field.
func (o *ActionPlan) SetResolutions(v []ActionPlanResolution) {
	o.Resolutions = v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *ActionPlan) GetEndpoint() string {
	if o == nil || o.Endpoint == nil {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionPlan) GetEndpointOk() (*string, bool) {
	if o == nil || o.Endpoint == nil {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *ActionPlan) HasEndpoint() bool {
	if o != nil && o.Endpoint != nil {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *ActionPlan) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetConfigurationId returns the ConfigurationId field value if set, zero value otherwise.
func (o *ActionPlan) GetConfigurationId() string {
	if o == nil || o.ConfigurationId == nil {
		var ret string
		return ret
	}
	return *o.ConfigurationId
}

// GetConfigurationIdOk returns a tuple with the ConfigurationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionPlan) GetConfigurationIdOk() (*string, bool) {
	if o == nil || o.ConfigurationId == nil {
		return nil, false
	}
	return o.ConfigurationId, true
}

// HasConfigurationId returns a boolean if a field has been set.
func (o *ActionPlan) HasConfigurationId() bool {
	if o != nil && o.ConfigurationId != nil {
		return true
	}

	return false
}

// SetConfigurationId gets a reference to the given string and assigns it to the ConfigurationId field.
func (o *ActionPlan) SetConfigurationId(v string) {
	o.ConfigurationId = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *ActionPlan) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionPlan) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *ActionPlan) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *ActionPlan) SetSubject(v string) {
	o.Subject = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ActionPlan) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionPlan) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ActionPlan) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ActionPlan) SetMessage(v string) {
	o.Message = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *ActionPlan) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionPlan) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *ActionPlan) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *ActionPlan) SetUuid(v string) {
	o.Uuid = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ActionPlan) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionPlan) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ActionPlan) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ActionPlan) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *ActionPlan) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionPlan) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *ActionPlan) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *ActionPlan) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o ActionPlan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.ImageTag != nil {
		toSerialize["image_tag"] = o.ImageTag
	}
	if o.ImageDigest != nil {
		toSerialize["image_digest"] = o.ImageDigest
	}
	if o.PolicyId != nil {
		toSerialize["policy_id"] = o.PolicyId
	}
	if o.Resolutions != nil {
		toSerialize["resolutions"] = o.Resolutions
	}
	if o.Endpoint != nil {
		toSerialize["endpoint"] = o.Endpoint
	}
	if o.ConfigurationId != nil {
		toSerialize["configuration_id"] = o.ConfigurationId
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	return json.Marshal(toSerialize)
}

type NullableActionPlan struct {
	value *ActionPlan
	isSet bool
}

func (v NullableActionPlan) Get() *ActionPlan {
	return v.value
}

func (v *NullableActionPlan) Set(val *ActionPlan) {
	v.value = val
	v.isSet = true
}

func (v NullableActionPlan) IsSet() bool {
	return v.isSet
}

func (v *NullableActionPlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionPlan(val *ActionPlan) *NullableActionPlan {
	return &NullableActionPlan{value: val, isSet: true}
}

func (v NullableActionPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionPlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


