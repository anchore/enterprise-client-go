/*
Anchore Enterprise API Server

This is the Anchore Enterprise API. It provides additional external API routes and functionality for enterprise users.

API version: 0.8.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// ImageArtifact Model for an image artifact
type ImageArtifact struct {
	// The account id
	AccountId *string `json:"account_id,omitempty"`
	// The digest of the image
	ImageDigest *string `json:"image_digest,omitempty"`
	// The distro of the image
	Distro *string `json:"distro,omitempty"`
	// The distro version of the image
	DistroVersion *string `json:"distro_version,omitempty"`
	// the analysis status of image
	AnalysisStatus *string `json:"analysis_status,omitempty"`
	// The status of the image
	ImageStatus *string `json:"image_status,omitempty"`
	// RFC 3339 formatted UTC timestamp when the image was analyzed
	AnalyzedAt *time.Time `json:"analyzed_at,omitempty"`
	// RFC 3339 formatted UTC timestamp when the image was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// RFC 3339 formatted UTC timestamp when the image was last updated
	LastUpdated *time.Time `json:"last_updated,omitempty"`
}

// NewImageArtifact instantiates a new ImageArtifact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageArtifact() *ImageArtifact {
	this := ImageArtifact{}
	return &this
}

// NewImageArtifactWithDefaults instantiates a new ImageArtifact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageArtifactWithDefaults() *ImageArtifact {
	this := ImageArtifact{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *ImageArtifact) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageArtifact) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *ImageArtifact) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *ImageArtifact) SetAccountId(v string) {
	o.AccountId = &v
}

// GetImageDigest returns the ImageDigest field value if set, zero value otherwise.
func (o *ImageArtifact) GetImageDigest() string {
	if o == nil || o.ImageDigest == nil {
		var ret string
		return ret
	}
	return *o.ImageDigest
}

// GetImageDigestOk returns a tuple with the ImageDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageArtifact) GetImageDigestOk() (*string, bool) {
	if o == nil || o.ImageDigest == nil {
		return nil, false
	}
	return o.ImageDigest, true
}

// HasImageDigest returns a boolean if a field has been set.
func (o *ImageArtifact) HasImageDigest() bool {
	if o != nil && o.ImageDigest != nil {
		return true
	}

	return false
}

// SetImageDigest gets a reference to the given string and assigns it to the ImageDigest field.
func (o *ImageArtifact) SetImageDigest(v string) {
	o.ImageDigest = &v
}

// GetDistro returns the Distro field value if set, zero value otherwise.
func (o *ImageArtifact) GetDistro() string {
	if o == nil || o.Distro == nil {
		var ret string
		return ret
	}
	return *o.Distro
}

// GetDistroOk returns a tuple with the Distro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageArtifact) GetDistroOk() (*string, bool) {
	if o == nil || o.Distro == nil {
		return nil, false
	}
	return o.Distro, true
}

// HasDistro returns a boolean if a field has been set.
func (o *ImageArtifact) HasDistro() bool {
	if o != nil && o.Distro != nil {
		return true
	}

	return false
}

// SetDistro gets a reference to the given string and assigns it to the Distro field.
func (o *ImageArtifact) SetDistro(v string) {
	o.Distro = &v
}

// GetDistroVersion returns the DistroVersion field value if set, zero value otherwise.
func (o *ImageArtifact) GetDistroVersion() string {
	if o == nil || o.DistroVersion == nil {
		var ret string
		return ret
	}
	return *o.DistroVersion
}

// GetDistroVersionOk returns a tuple with the DistroVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageArtifact) GetDistroVersionOk() (*string, bool) {
	if o == nil || o.DistroVersion == nil {
		return nil, false
	}
	return o.DistroVersion, true
}

// HasDistroVersion returns a boolean if a field has been set.
func (o *ImageArtifact) HasDistroVersion() bool {
	if o != nil && o.DistroVersion != nil {
		return true
	}

	return false
}

// SetDistroVersion gets a reference to the given string and assigns it to the DistroVersion field.
func (o *ImageArtifact) SetDistroVersion(v string) {
	o.DistroVersion = &v
}

// GetAnalysisStatus returns the AnalysisStatus field value if set, zero value otherwise.
func (o *ImageArtifact) GetAnalysisStatus() string {
	if o == nil || o.AnalysisStatus == nil {
		var ret string
		return ret
	}
	return *o.AnalysisStatus
}

// GetAnalysisStatusOk returns a tuple with the AnalysisStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageArtifact) GetAnalysisStatusOk() (*string, bool) {
	if o == nil || o.AnalysisStatus == nil {
		return nil, false
	}
	return o.AnalysisStatus, true
}

// HasAnalysisStatus returns a boolean if a field has been set.
func (o *ImageArtifact) HasAnalysisStatus() bool {
	if o != nil && o.AnalysisStatus != nil {
		return true
	}

	return false
}

// SetAnalysisStatus gets a reference to the given string and assigns it to the AnalysisStatus field.
func (o *ImageArtifact) SetAnalysisStatus(v string) {
	o.AnalysisStatus = &v
}

// GetImageStatus returns the ImageStatus field value if set, zero value otherwise.
func (o *ImageArtifact) GetImageStatus() string {
	if o == nil || o.ImageStatus == nil {
		var ret string
		return ret
	}
	return *o.ImageStatus
}

// GetImageStatusOk returns a tuple with the ImageStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageArtifact) GetImageStatusOk() (*string, bool) {
	if o == nil || o.ImageStatus == nil {
		return nil, false
	}
	return o.ImageStatus, true
}

// HasImageStatus returns a boolean if a field has been set.
func (o *ImageArtifact) HasImageStatus() bool {
	if o != nil && o.ImageStatus != nil {
		return true
	}

	return false
}

// SetImageStatus gets a reference to the given string and assigns it to the ImageStatus field.
func (o *ImageArtifact) SetImageStatus(v string) {
	o.ImageStatus = &v
}

// GetAnalyzedAt returns the AnalyzedAt field value if set, zero value otherwise.
func (o *ImageArtifact) GetAnalyzedAt() time.Time {
	if o == nil || o.AnalyzedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.AnalyzedAt
}

// GetAnalyzedAtOk returns a tuple with the AnalyzedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageArtifact) GetAnalyzedAtOk() (*time.Time, bool) {
	if o == nil || o.AnalyzedAt == nil {
		return nil, false
	}
	return o.AnalyzedAt, true
}

// HasAnalyzedAt returns a boolean if a field has been set.
func (o *ImageArtifact) HasAnalyzedAt() bool {
	if o != nil && o.AnalyzedAt != nil {
		return true
	}

	return false
}

// SetAnalyzedAt gets a reference to the given time.Time and assigns it to the AnalyzedAt field.
func (o *ImageArtifact) SetAnalyzedAt(v time.Time) {
	o.AnalyzedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ImageArtifact) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageArtifact) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ImageArtifact) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ImageArtifact) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *ImageArtifact) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageArtifact) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *ImageArtifact) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *ImageArtifact) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o ImageArtifact) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.ImageDigest != nil {
		toSerialize["image_digest"] = o.ImageDigest
	}
	if o.Distro != nil {
		toSerialize["distro"] = o.Distro
	}
	if o.DistroVersion != nil {
		toSerialize["distro_version"] = o.DistroVersion
	}
	if o.AnalysisStatus != nil {
		toSerialize["analysis_status"] = o.AnalysisStatus
	}
	if o.ImageStatus != nil {
		toSerialize["image_status"] = o.ImageStatus
	}
	if o.AnalyzedAt != nil {
		toSerialize["analyzed_at"] = o.AnalyzedAt
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	return json.Marshal(toSerialize)
}

type NullableImageArtifact struct {
	value *ImageArtifact
	isSet bool
}

func (v NullableImageArtifact) Get() *ImageArtifact {
	return v.value
}

func (v *NullableImageArtifact) Set(val *ImageArtifact) {
	v.value = val
	v.isSet = true
}

func (v NullableImageArtifact) IsSet() bool {
	return v.isSet
}

func (v *NullableImageArtifact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageArtifact(val *ImageArtifact) *NullableImageArtifact {
	return &NullableImageArtifact{value: val, isSet: true}
}

func (v NullableImageArtifact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageArtifact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


