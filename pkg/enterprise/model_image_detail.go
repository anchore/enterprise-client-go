/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 0.1.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// ImageDetail A metadata detail record for a specific image. Multiple detail records may map a single catalog image.
type ImageDetail struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
	// Full docker-pullable tag string referencing the image
	Fulltag *string `json:"fulltag,omitempty"`
	// Full docker-pullable digest string including the registry url and repository necessary get the image
	Fulldigest *string `json:"fulldigest,omitempty"`
	AccountName *string `json:"account_name,omitempty"`
	ImageId *string `json:"image_id,omitempty"`
	Registry *string `json:"registry,omitempty"`
	Repo *string `json:"repo,omitempty"`
	Dockerfile NullableString `json:"dockerfile,omitempty"`
	// The parent Anchore Image record to which this detail maps
	ImageDigest *string `json:"image_digest,omitempty"`
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
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageDetail) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ImageDetail) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
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
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageDetail) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *ImageDetail) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *ImageDetail) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetFulltag returns the Fulltag field value if set, zero value otherwise.
func (o *ImageDetail) GetFulltag() string {
	if o == nil || o.Fulltag == nil {
		var ret string
		return ret
	}
	return *o.Fulltag
}

// GetFulltagOk returns a tuple with the Fulltag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageDetail) GetFulltagOk() (*string, bool) {
	if o == nil || o.Fulltag == nil {
		return nil, false
	}
	return o.Fulltag, true
}

// HasFulltag returns a boolean if a field has been set.
func (o *ImageDetail) HasFulltag() bool {
	if o != nil && o.Fulltag != nil {
		return true
	}

	return false
}

// SetFulltag gets a reference to the given string and assigns it to the Fulltag field.
func (o *ImageDetail) SetFulltag(v string) {
	o.Fulltag = &v
}

// GetFulldigest returns the Fulldigest field value if set, zero value otherwise.
func (o *ImageDetail) GetFulldigest() string {
	if o == nil || o.Fulldigest == nil {
		var ret string
		return ret
	}
	return *o.Fulldigest
}

// GetFulldigestOk returns a tuple with the Fulldigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageDetail) GetFulldigestOk() (*string, bool) {
	if o == nil || o.Fulldigest == nil {
		return nil, false
	}
	return o.Fulldigest, true
}

// HasFulldigest returns a boolean if a field has been set.
func (o *ImageDetail) HasFulldigest() bool {
	if o != nil && o.Fulldigest != nil {
		return true
	}

	return false
}

// SetFulldigest gets a reference to the given string and assigns it to the Fulldigest field.
func (o *ImageDetail) SetFulldigest(v string) {
	o.Fulldigest = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *ImageDetail) GetAccountName() string {
	if o == nil || o.AccountName == nil {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageDetail) GetAccountNameOk() (*string, bool) {
	if o == nil || o.AccountName == nil {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *ImageDetail) HasAccountName() bool {
	if o != nil && o.AccountName != nil {
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
	if o == nil || o.ImageId == nil {
		var ret string
		return ret
	}
	return *o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageDetail) GetImageIdOk() (*string, bool) {
	if o == nil || o.ImageId == nil {
		return nil, false
	}
	return o.ImageId, true
}

// HasImageId returns a boolean if a field has been set.
func (o *ImageDetail) HasImageId() bool {
	if o != nil && o.ImageId != nil {
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
	if o == nil || o.Registry == nil {
		var ret string
		return ret
	}
	return *o.Registry
}

// GetRegistryOk returns a tuple with the Registry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageDetail) GetRegistryOk() (*string, bool) {
	if o == nil || o.Registry == nil {
		return nil, false
	}
	return o.Registry, true
}

// HasRegistry returns a boolean if a field has been set.
func (o *ImageDetail) HasRegistry() bool {
	if o != nil && o.Registry != nil {
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
	if o == nil || o.Repo == nil {
		var ret string
		return ret
	}
	return *o.Repo
}

// GetRepoOk returns a tuple with the Repo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageDetail) GetRepoOk() (*string, bool) {
	if o == nil || o.Repo == nil {
		return nil, false
	}
	return o.Repo, true
}

// HasRepo returns a boolean if a field has been set.
func (o *ImageDetail) HasRepo() bool {
	if o != nil && o.Repo != nil {
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
	if o == nil || o.Dockerfile.Get() == nil {
		var ret string
		return ret
	}
	return *o.Dockerfile.Get()
}

// GetDockerfileOk returns a tuple with the Dockerfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageDetail) GetDockerfileOk() (*string, bool) {
	if o == nil  {
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
	if o == nil || o.ImageDigest == nil {
		var ret string
		return ret
	}
	return *o.ImageDigest
}

// GetImageDigestOk returns a tuple with the ImageDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageDetail) GetImageDigestOk() (*string, bool) {
	if o == nil || o.ImageDigest == nil {
		return nil, false
	}
	return o.ImageDigest, true
}

// HasImageDigest returns a boolean if a field has been set.
func (o *ImageDetail) HasImageDigest() bool {
	if o != nil && o.ImageDigest != nil {
		return true
	}

	return false
}

// SetImageDigest gets a reference to the given string and assigns it to the ImageDigest field.
func (o *ImageDetail) SetImageDigest(v string) {
	o.ImageDigest = &v
}

func (o ImageDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if o.Fulltag != nil {
		toSerialize["fulltag"] = o.Fulltag
	}
	if o.Fulldigest != nil {
		toSerialize["fulldigest"] = o.Fulldigest
	}
	if o.AccountName != nil {
		toSerialize["account_name"] = o.AccountName
	}
	if o.ImageId != nil {
		toSerialize["image_id"] = o.ImageId
	}
	if o.Registry != nil {
		toSerialize["registry"] = o.Registry
	}
	if o.Repo != nil {
		toSerialize["repo"] = o.Repo
	}
	if o.Dockerfile.IsSet() {
		toSerialize["dockerfile"] = o.Dockerfile.Get()
	}
	if o.ImageDigest != nil {
		toSerialize["image_digest"] = o.ImageDigest
	}
	return json.Marshal(toSerialize)
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


