/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.8.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// checks if the ImageDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageDetail{}

// ImageDetail A metadata detail record for a specific image. Multiple detail records may map to a single catalog image. For example, an image having multiple tags associated with a single image digest will generate multiple image detail records.
type ImageDetail struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
	// Full docker-pullable tag string referencing the image
	FullTag *string `json:"full_tag,omitempty"`
	// Full docker-pullable digest string including the registry url and repository necessary get the image
	FullDigest *string `json:"full_digest,omitempty"`
	AccountName *string `json:"account_name,omitempty"`
	ImageId *string `json:"image_id,omitempty"`
	Registry *string `json:"registry,omitempty"`
	Repo *string `json:"repo,omitempty"`
	Dockerfile NullableString `json:"dockerfile,omitempty"`
	// The parent Anchore Image record to which this detail maps
	ImageDigest *string `json:"image_digest,omitempty"`
	Tag *string `json:"tag,omitempty"`
	TagDetectedAt *time.Time `json:"tag_detected_at,omitempty"`
}

// NewImageDetail instantiates a new ImageDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageDetail() *ImageDetail {
	this := ImageDetail{}
	return &this
}

// NewImageDetailWithDefaults instantiates a new ImageDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageDetailWithDefaults() *ImageDetail {
	this := ImageDetail{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ImageDetail) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageDetail) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ImageDetail) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ImageDetail) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *ImageDetail) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageDetail) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *ImageDetail) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *ImageDetail) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetFullTag returns the FullTag field value if set, zero value otherwise.
func (o *ImageDetail) GetFullTag() string {
	if o == nil || IsNil(o.FullTag) {
		var ret string
		return ret
	}
	return *o.FullTag
}

// GetFullTagOk returns a tuple with the FullTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageDetail) GetFullTagOk() (*string, bool) {
	if o == nil || IsNil(o.FullTag) {
		return nil, false
	}
	return o.FullTag, true
}

// HasFullTag returns a boolean if a field has been set.
func (o *ImageDetail) HasFullTag() bool {
	if o != nil && !IsNil(o.FullTag) {
		return true
	}

	return false
}

// SetFullTag gets a reference to the given string and assigns it to the FullTag field.
func (o *ImageDetail) SetFullTag(v string) {
	o.FullTag = &v
}

// GetFullDigest returns the FullDigest field value if set, zero value otherwise.
func (o *ImageDetail) GetFullDigest() string {
	if o == nil || IsNil(o.FullDigest) {
		var ret string
		return ret
	}
	return *o.FullDigest
}

// GetFullDigestOk returns a tuple with the FullDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageDetail) GetFullDigestOk() (*string, bool) {
	if o == nil || IsNil(o.FullDigest) {
		return nil, false
	}
	return o.FullDigest, true
}

// HasFullDigest returns a boolean if a field has been set.
func (o *ImageDetail) HasFullDigest() bool {
	if o != nil && !IsNil(o.FullDigest) {
		return true
	}

	return false
}

// SetFullDigest gets a reference to the given string and assigns it to the FullDigest field.
func (o *ImageDetail) SetFullDigest(v string) {
	o.FullDigest = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *ImageDetail) GetAccountName() string {
	if o == nil || IsNil(o.AccountName) {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageDetail) GetAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.AccountName) {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *ImageDetail) HasAccountName() bool {
	if o != nil && !IsNil(o.AccountName) {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *ImageDetail) SetAccountName(v string) {
	o.AccountName = &v
}

// GetImageId returns the ImageId field value if set, zero value otherwise.
func (o *ImageDetail) GetImageId() string {
	if o == nil || IsNil(o.ImageId) {
		var ret string
		return ret
	}
	return *o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageDetail) GetImageIdOk() (*string, bool) {
	if o == nil || IsNil(o.ImageId) {
		return nil, false
	}
	return o.ImageId, true
}

// HasImageId returns a boolean if a field has been set.
func (o *ImageDetail) HasImageId() bool {
	if o != nil && !IsNil(o.ImageId) {
		return true
	}

	return false
}

// SetImageId gets a reference to the given string and assigns it to the ImageId field.
func (o *ImageDetail) SetImageId(v string) {
	o.ImageId = &v
}

// GetRegistry returns the Registry field value if set, zero value otherwise.
func (o *ImageDetail) GetRegistry() string {
	if o == nil || IsNil(o.Registry) {
		var ret string
		return ret
	}
	return *o.Registry
}

// GetRegistryOk returns a tuple with the Registry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageDetail) GetRegistryOk() (*string, bool) {
	if o == nil || IsNil(o.Registry) {
		return nil, false
	}
	return o.Registry, true
}

// HasRegistry returns a boolean if a field has been set.
func (o *ImageDetail) HasRegistry() bool {
	if o != nil && !IsNil(o.Registry) {
		return true
	}

	return false
}

// SetRegistry gets a reference to the given string and assigns it to the Registry field.
func (o *ImageDetail) SetRegistry(v string) {
	o.Registry = &v
}

// GetRepo returns the Repo field value if set, zero value otherwise.
func (o *ImageDetail) GetRepo() string {
	if o == nil || IsNil(o.Repo) {
		var ret string
		return ret
	}
	return *o.Repo
}

// GetRepoOk returns a tuple with the Repo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageDetail) GetRepoOk() (*string, bool) {
	if o == nil || IsNil(o.Repo) {
		return nil, false
	}
	return o.Repo, true
}

// HasRepo returns a boolean if a field has been set.
func (o *ImageDetail) HasRepo() bool {
	if o != nil && !IsNil(o.Repo) {
		return true
	}

	return false
}

// SetRepo gets a reference to the given string and assigns it to the Repo field.
func (o *ImageDetail) SetRepo(v string) {
	o.Repo = &v
}

// GetDockerfile returns the Dockerfile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageDetail) GetDockerfile() string {
	if o == nil || IsNil(o.Dockerfile.Get()) {
		var ret string
		return ret
	}
	return *o.Dockerfile.Get()
}

// GetDockerfileOk returns a tuple with the Dockerfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageDetail) GetDockerfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dockerfile.Get(), o.Dockerfile.IsSet()
}

// HasDockerfile returns a boolean if a field has been set.
func (o *ImageDetail) HasDockerfile() bool {
	if o != nil && o.Dockerfile.IsSet() {
		return true
	}

	return false
}

// SetDockerfile gets a reference to the given NullableString and assigns it to the Dockerfile field.
func (o *ImageDetail) SetDockerfile(v string) {
	o.Dockerfile.Set(&v)
}
// SetDockerfileNil sets the value for Dockerfile to be an explicit nil
func (o *ImageDetail) SetDockerfileNil() {
	o.Dockerfile.Set(nil)
}

// UnsetDockerfile ensures that no value is present for Dockerfile, not even an explicit nil
func (o *ImageDetail) UnsetDockerfile() {
	o.Dockerfile.Unset()
}

// GetImageDigest returns the ImageDigest field value if set, zero value otherwise.
func (o *ImageDetail) GetImageDigest() string {
	if o == nil || IsNil(o.ImageDigest) {
		var ret string
		return ret
	}
	return *o.ImageDigest
}

// GetImageDigestOk returns a tuple with the ImageDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageDetail) GetImageDigestOk() (*string, bool) {
	if o == nil || IsNil(o.ImageDigest) {
		return nil, false
	}
	return o.ImageDigest, true
}

// HasImageDigest returns a boolean if a field has been set.
func (o *ImageDetail) HasImageDigest() bool {
	if o != nil && !IsNil(o.ImageDigest) {
		return true
	}

	return false
}

// SetImageDigest gets a reference to the given string and assigns it to the ImageDigest field.
func (o *ImageDetail) SetImageDigest(v string) {
	o.ImageDigest = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *ImageDetail) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageDetail) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *ImageDetail) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *ImageDetail) SetTag(v string) {
	o.Tag = &v
}

// GetTagDetectedAt returns the TagDetectedAt field value if set, zero value otherwise.
func (o *ImageDetail) GetTagDetectedAt() time.Time {
	if o == nil || IsNil(o.TagDetectedAt) {
		var ret time.Time
		return ret
	}
	return *o.TagDetectedAt
}

// GetTagDetectedAtOk returns a tuple with the TagDetectedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageDetail) GetTagDetectedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TagDetectedAt) {
		return nil, false
	}
	return o.TagDetectedAt, true
}

// HasTagDetectedAt returns a boolean if a field has been set.
func (o *ImageDetail) HasTagDetectedAt() bool {
	if o != nil && !IsNil(o.TagDetectedAt) {
		return true
	}

	return false
}

// SetTagDetectedAt gets a reference to the given time.Time and assigns it to the TagDetectedAt field.
func (o *ImageDetail) SetTagDetectedAt(v time.Time) {
	o.TagDetectedAt = &v
}

func (o ImageDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if !IsNil(o.FullTag) {
		toSerialize["full_tag"] = o.FullTag
	}
	if !IsNil(o.FullDigest) {
		toSerialize["full_digest"] = o.FullDigest
	}
	if !IsNil(o.AccountName) {
		toSerialize["account_name"] = o.AccountName
	}
	if !IsNil(o.ImageId) {
		toSerialize["image_id"] = o.ImageId
	}
	if !IsNil(o.Registry) {
		toSerialize["registry"] = o.Registry
	}
	if !IsNil(o.Repo) {
		toSerialize["repo"] = o.Repo
	}
	if o.Dockerfile.IsSet() {
		toSerialize["dockerfile"] = o.Dockerfile.Get()
	}
	if !IsNil(o.ImageDigest) {
		toSerialize["image_digest"] = o.ImageDigest
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.TagDetectedAt) {
		toSerialize["tag_detected_at"] = o.TagDetectedAt
	}
	return toSerialize, nil
}

type NullableImageDetail struct {
	value *ImageDetail
	isSet bool
}

func (v NullableImageDetail) Get() *ImageDetail {
	return v.value
}

func (v *NullableImageDetail) Set(val *ImageDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableImageDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableImageDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageDetail(val *ImageDetail) *NullableImageDetail {
	return &NullableImageDetail{value: val, isSet: true}
}

func (v NullableImageDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


