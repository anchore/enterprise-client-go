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

// ComplianceResource A resource that ties compliance related artifacts - image digest, tag and policy
type ComplianceResource struct {
	ImageDigest *string `json:"image_digest,omitempty"`
	PolicyId *string `json:"policy_id,omitempty"`
	ImageTag *string `json:"image_tag,omitempty"`
	Registry *string `json:"registry,omitempty"`
	Repository *string `json:"repository,omitempty"`
	EvaluationId *string `json:"evaluation_id,omitempty"`
	EvaluatedAt *time.Time `json:"evaluated_at,omitempty"`
}

// NewComplianceResource instantiates a new ComplianceResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComplianceResource() *ComplianceResource {
	this := ComplianceResource{}
	return &this
}

// NewComplianceResourceWithDefaults instantiates a new ComplianceResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComplianceResourceWithDefaults() *ComplianceResource {
	this := ComplianceResource{}
	return &this
}

// GetImageDigest returns the ImageDigest field value if set, zero value otherwise.
func (o *ComplianceResource) GetImageDigest() string {
	if o == nil || o.ImageDigest == nil {
		var ret string
		return ret
	}
	return *o.ImageDigest
}

// GetImageDigestOk returns a tuple with the ImageDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplianceResource) GetImageDigestOk() (*string, bool) {
	if o == nil || o.ImageDigest == nil {
		return nil, false
	}
	return o.ImageDigest, true
}

// HasImageDigest returns a boolean if a field has been set.
func (o *ComplianceResource) HasImageDigest() bool {
	if o != nil && o.ImageDigest != nil {
		return true
	}

	return false
}

// SetImageDigest gets a reference to the given string and assigns it to the ImageDigest field.
func (o *ComplianceResource) SetImageDigest(v string) {
	o.ImageDigest = &v
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *ComplianceResource) GetPolicyId() string {
	if o == nil || o.PolicyId == nil {
		var ret string
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplianceResource) GetPolicyIdOk() (*string, bool) {
	if o == nil || o.PolicyId == nil {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *ComplianceResource) HasPolicyId() bool {
	if o != nil && o.PolicyId != nil {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given string and assigns it to the PolicyId field.
func (o *ComplianceResource) SetPolicyId(v string) {
	o.PolicyId = &v
}

// GetImageTag returns the ImageTag field value if set, zero value otherwise.
func (o *ComplianceResource) GetImageTag() string {
	if o == nil || o.ImageTag == nil {
		var ret string
		return ret
	}
	return *o.ImageTag
}

// GetImageTagOk returns a tuple with the ImageTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplianceResource) GetImageTagOk() (*string, bool) {
	if o == nil || o.ImageTag == nil {
		return nil, false
	}
	return o.ImageTag, true
}

// HasImageTag returns a boolean if a field has been set.
func (o *ComplianceResource) HasImageTag() bool {
	if o != nil && o.ImageTag != nil {
		return true
	}

	return false
}

// SetImageTag gets a reference to the given string and assigns it to the ImageTag field.
func (o *ComplianceResource) SetImageTag(v string) {
	o.ImageTag = &v
}

// GetRegistry returns the Registry field value if set, zero value otherwise.
func (o *ComplianceResource) GetRegistry() string {
	if o == nil || o.Registry == nil {
		var ret string
		return ret
	}
	return *o.Registry
}

// GetRegistryOk returns a tuple with the Registry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplianceResource) GetRegistryOk() (*string, bool) {
	if o == nil || o.Registry == nil {
		return nil, false
	}
	return o.Registry, true
}

// HasRegistry returns a boolean if a field has been set.
func (o *ComplianceResource) HasRegistry() bool {
	if o != nil && o.Registry != nil {
		return true
	}

	return false
}

// SetRegistry gets a reference to the given string and assigns it to the Registry field.
func (o *ComplianceResource) SetRegistry(v string) {
	o.Registry = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *ComplianceResource) GetRepository() string {
	if o == nil || o.Repository == nil {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplianceResource) GetRepositoryOk() (*string, bool) {
	if o == nil || o.Repository == nil {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *ComplianceResource) HasRepository() bool {
	if o != nil && o.Repository != nil {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *ComplianceResource) SetRepository(v string) {
	o.Repository = &v
}

// GetEvaluationId returns the EvaluationId field value if set, zero value otherwise.
func (o *ComplianceResource) GetEvaluationId() string {
	if o == nil || o.EvaluationId == nil {
		var ret string
		return ret
	}
	return *o.EvaluationId
}

// GetEvaluationIdOk returns a tuple with the EvaluationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplianceResource) GetEvaluationIdOk() (*string, bool) {
	if o == nil || o.EvaluationId == nil {
		return nil, false
	}
	return o.EvaluationId, true
}

// HasEvaluationId returns a boolean if a field has been set.
func (o *ComplianceResource) HasEvaluationId() bool {
	if o != nil && o.EvaluationId != nil {
		return true
	}

	return false
}

// SetEvaluationId gets a reference to the given string and assigns it to the EvaluationId field.
func (o *ComplianceResource) SetEvaluationId(v string) {
	o.EvaluationId = &v
}

// GetEvaluatedAt returns the EvaluatedAt field value if set, zero value otherwise.
func (o *ComplianceResource) GetEvaluatedAt() time.Time {
	if o == nil || o.EvaluatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.EvaluatedAt
}

// GetEvaluatedAtOk returns a tuple with the EvaluatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplianceResource) GetEvaluatedAtOk() (*time.Time, bool) {
	if o == nil || o.EvaluatedAt == nil {
		return nil, false
	}
	return o.EvaluatedAt, true
}

// HasEvaluatedAt returns a boolean if a field has been set.
func (o *ComplianceResource) HasEvaluatedAt() bool {
	if o != nil && o.EvaluatedAt != nil {
		return true
	}

	return false
}

// SetEvaluatedAt gets a reference to the given time.Time and assigns it to the EvaluatedAt field.
func (o *ComplianceResource) SetEvaluatedAt(v time.Time) {
	o.EvaluatedAt = &v
}

func (o ComplianceResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ImageDigest != nil {
		toSerialize["image_digest"] = o.ImageDigest
	}
	if o.PolicyId != nil {
		toSerialize["policy_id"] = o.PolicyId
	}
	if o.ImageTag != nil {
		toSerialize["image_tag"] = o.ImageTag
	}
	if o.Registry != nil {
		toSerialize["registry"] = o.Registry
	}
	if o.Repository != nil {
		toSerialize["repository"] = o.Repository
	}
	if o.EvaluationId != nil {
		toSerialize["evaluation_id"] = o.EvaluationId
	}
	if o.EvaluatedAt != nil {
		toSerialize["evaluated_at"] = o.EvaluatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableComplianceResource struct {
	value *ComplianceResource
	isSet bool
}

func (v NullableComplianceResource) Get() *ComplianceResource {
	return v.value
}

func (v *NullableComplianceResource) Set(val *ComplianceResource) {
	v.value = val
	v.isSet = true
}

func (v NullableComplianceResource) IsSet() bool {
	return v.isSet
}

func (v *NullableComplianceResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComplianceResource(val *ComplianceResource) *NullableComplianceResource {
	return &NullableComplianceResource{value: val, isSet: true}
}

func (v NullableComplianceResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComplianceResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


