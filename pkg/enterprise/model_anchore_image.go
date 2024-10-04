/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.7.2
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// checks if the AnchoreImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnchoreImage{}

// AnchoreImage A unique image in the engine. May have multiple tags or references. Unique to an image content across registries or repositories.
type AnchoreImage struct {
	// A metadata content record for a specific image, containing different content type entries
	ImageContent interface{} `json:"image_content,omitempty"`
	// Details specific to an image reference and type such as tag and image source
	ImageDetail []ImageDetail `json:"image_detail,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	ImageDigest *string `json:"image_digest,omitempty"`
	ParentDigest *string `json:"parent_digest,omitempty"`
	AccountName *string `json:"account_name,omitempty"`
	Annotations interface{} `json:"annotations,omitempty"`
	// State of the image
	ImageStatus *string `json:"image_status,omitempty"`
	// A state value for the current status of the analysis progress of the image
	AnalysisStatus *string `json:"analysis_status,omitempty"`
	// The version of the record, used for internal schema updates and data migrations.
	RecordVersion *string `json:"record_version,omitempty"`
	AnalysisStatusDetail []AnalysisStatusDetail `json:"analysis_status_detail,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AnchoreImage AnchoreImage

// NewAnchoreImage instantiates a new AnchoreImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnchoreImage() *AnchoreImage {
	this := AnchoreImage{}
	return &this
}

// NewAnchoreImageWithDefaults instantiates a new AnchoreImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnchoreImageWithDefaults() *AnchoreImage {
	this := AnchoreImage{}
	return &this
}

// GetImageContent returns the ImageContent field value if set, zero value otherwise.
func (o *AnchoreImage) GetImageContent() interface{} {
	if o == nil || IsNil(o.ImageContent) {
		var ret interface{}
		return ret
	}
	return o.ImageContent
}

// GetImageContentOk returns a tuple with the ImageContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnchoreImage) GetImageContentOk() (interface{}, bool) {
	if o == nil || IsNil(o.ImageContent) {
		return nil, false
	}
	return o.ImageContent, true
}

// HasImageContent returns a boolean if a field has been set.
func (o *AnchoreImage) HasImageContent() bool {
	if o != nil && !IsNil(o.ImageContent) {
		return true
	}

	return false
}

// SetImageContent gets a reference to the given interface{} and assigns it to the ImageContent field.
func (o *AnchoreImage) SetImageContent(v interface{}) {
	o.ImageContent = v
}

// GetImageDetail returns the ImageDetail field value if set, zero value otherwise.
func (o *AnchoreImage) GetImageDetail() []ImageDetail {
	if o == nil || IsNil(o.ImageDetail) {
		var ret []ImageDetail
		return ret
	}
	return o.ImageDetail
}

// GetImageDetailOk returns a tuple with the ImageDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnchoreImage) GetImageDetailOk() ([]ImageDetail, bool) {
	if o == nil || IsNil(o.ImageDetail) {
		return nil, false
	}
	return o.ImageDetail, true
}

// HasImageDetail returns a boolean if a field has been set.
func (o *AnchoreImage) HasImageDetail() bool {
	if o != nil && !IsNil(o.ImageDetail) {
		return true
	}

	return false
}

// SetImageDetail gets a reference to the given []ImageDetail and assigns it to the ImageDetail field.
func (o *AnchoreImage) SetImageDetail(v []ImageDetail) {
	o.ImageDetail = v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *AnchoreImage) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnchoreImage) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *AnchoreImage) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *AnchoreImage) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AnchoreImage) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnchoreImage) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AnchoreImage) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AnchoreImage) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetImageDigest returns the ImageDigest field value if set, zero value otherwise.
func (o *AnchoreImage) GetImageDigest() string {
	if o == nil || IsNil(o.ImageDigest) {
		var ret string
		return ret
	}
	return *o.ImageDigest
}

// GetImageDigestOk returns a tuple with the ImageDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnchoreImage) GetImageDigestOk() (*string, bool) {
	if o == nil || IsNil(o.ImageDigest) {
		return nil, false
	}
	return o.ImageDigest, true
}

// HasImageDigest returns a boolean if a field has been set.
func (o *AnchoreImage) HasImageDigest() bool {
	if o != nil && !IsNil(o.ImageDigest) {
		return true
	}

	return false
}

// SetImageDigest gets a reference to the given string and assigns it to the ImageDigest field.
func (o *AnchoreImage) SetImageDigest(v string) {
	o.ImageDigest = &v
}

// GetParentDigest returns the ParentDigest field value if set, zero value otherwise.
func (o *AnchoreImage) GetParentDigest() string {
	if o == nil || IsNil(o.ParentDigest) {
		var ret string
		return ret
	}
	return *o.ParentDigest
}

// GetParentDigestOk returns a tuple with the ParentDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnchoreImage) GetParentDigestOk() (*string, bool) {
	if o == nil || IsNil(o.ParentDigest) {
		return nil, false
	}
	return o.ParentDigest, true
}

// HasParentDigest returns a boolean if a field has been set.
func (o *AnchoreImage) HasParentDigest() bool {
	if o != nil && !IsNil(o.ParentDigest) {
		return true
	}

	return false
}

// SetParentDigest gets a reference to the given string and assigns it to the ParentDigest field.
func (o *AnchoreImage) SetParentDigest(v string) {
	o.ParentDigest = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *AnchoreImage) GetAccountName() string {
	if o == nil || IsNil(o.AccountName) {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnchoreImage) GetAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.AccountName) {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *AnchoreImage) HasAccountName() bool {
	if o != nil && !IsNil(o.AccountName) {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *AnchoreImage) SetAccountName(v string) {
	o.AccountName = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *AnchoreImage) GetAnnotations() interface{} {
	if o == nil || IsNil(o.Annotations) {
		var ret interface{}
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnchoreImage) GetAnnotationsOk() (interface{}, bool) {
	if o == nil || IsNil(o.Annotations) {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *AnchoreImage) HasAnnotations() bool {
	if o != nil && !IsNil(o.Annotations) {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given interface{} and assigns it to the Annotations field.
func (o *AnchoreImage) SetAnnotations(v interface{}) {
	o.Annotations = v
}

// GetImageStatus returns the ImageStatus field value if set, zero value otherwise.
func (o *AnchoreImage) GetImageStatus() string {
	if o == nil || IsNil(o.ImageStatus) {
		var ret string
		return ret
	}
	return *o.ImageStatus
}

// GetImageStatusOk returns a tuple with the ImageStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnchoreImage) GetImageStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ImageStatus) {
		return nil, false
	}
	return o.ImageStatus, true
}

// HasImageStatus returns a boolean if a field has been set.
func (o *AnchoreImage) HasImageStatus() bool {
	if o != nil && !IsNil(o.ImageStatus) {
		return true
	}

	return false
}

// SetImageStatus gets a reference to the given string and assigns it to the ImageStatus field.
func (o *AnchoreImage) SetImageStatus(v string) {
	o.ImageStatus = &v
}

// GetAnalysisStatus returns the AnalysisStatus field value if set, zero value otherwise.
func (o *AnchoreImage) GetAnalysisStatus() string {
	if o == nil || IsNil(o.AnalysisStatus) {
		var ret string
		return ret
	}
	return *o.AnalysisStatus
}

// GetAnalysisStatusOk returns a tuple with the AnalysisStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnchoreImage) GetAnalysisStatusOk() (*string, bool) {
	if o == nil || IsNil(o.AnalysisStatus) {
		return nil, false
	}
	return o.AnalysisStatus, true
}

// HasAnalysisStatus returns a boolean if a field has been set.
func (o *AnchoreImage) HasAnalysisStatus() bool {
	if o != nil && !IsNil(o.AnalysisStatus) {
		return true
	}

	return false
}

// SetAnalysisStatus gets a reference to the given string and assigns it to the AnalysisStatus field.
func (o *AnchoreImage) SetAnalysisStatus(v string) {
	o.AnalysisStatus = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *AnchoreImage) GetRecordVersion() string {
	if o == nil || IsNil(o.RecordVersion) {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnchoreImage) GetRecordVersionOk() (*string, bool) {
	if o == nil || IsNil(o.RecordVersion) {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *AnchoreImage) HasRecordVersion() bool {
	if o != nil && !IsNil(o.RecordVersion) {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *AnchoreImage) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetAnalysisStatusDetail returns the AnalysisStatusDetail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnchoreImage) GetAnalysisStatusDetail() []AnalysisStatusDetail {
	if o == nil {
		var ret []AnalysisStatusDetail
		return ret
	}
	return o.AnalysisStatusDetail
}

// GetAnalysisStatusDetailOk returns a tuple with the AnalysisStatusDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnchoreImage) GetAnalysisStatusDetailOk() ([]AnalysisStatusDetail, bool) {
	if o == nil || IsNil(o.AnalysisStatusDetail) {
		return nil, false
	}
	return o.AnalysisStatusDetail, true
}

// HasAnalysisStatusDetail returns a boolean if a field has been set.
func (o *AnchoreImage) HasAnalysisStatusDetail() bool {
	if o != nil && !IsNil(o.AnalysisStatusDetail) {
		return true
	}

	return false
}

// SetAnalysisStatusDetail gets a reference to the given []AnalysisStatusDetail and assigns it to the AnalysisStatusDetail field.
func (o *AnchoreImage) SetAnalysisStatusDetail(v []AnalysisStatusDetail) {
	o.AnalysisStatusDetail = v
}

func (o AnchoreImage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnchoreImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ImageContent) {
		toSerialize["image_content"] = o.ImageContent
	}
	if !IsNil(o.ImageDetail) {
		toSerialize["image_detail"] = o.ImageDetail
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.ImageDigest) {
		toSerialize["image_digest"] = o.ImageDigest
	}
	if !IsNil(o.ParentDigest) {
		toSerialize["parent_digest"] = o.ParentDigest
	}
	if !IsNil(o.AccountName) {
		toSerialize["account_name"] = o.AccountName
	}
	if !IsNil(o.Annotations) {
		toSerialize["annotations"] = o.Annotations
	}
	if !IsNil(o.ImageStatus) {
		toSerialize["image_status"] = o.ImageStatus
	}
	if !IsNil(o.AnalysisStatus) {
		toSerialize["analysis_status"] = o.AnalysisStatus
	}
	if !IsNil(o.RecordVersion) {
		toSerialize["record_version"] = o.RecordVersion
	}
	if o.AnalysisStatusDetail != nil {
		toSerialize["analysis_status_detail"] = o.AnalysisStatusDetail
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AnchoreImage) UnmarshalJSON(data []byte) (err error) {
	varAnchoreImage := _AnchoreImage{}

	err = json.Unmarshal(data, &varAnchoreImage)

	if err != nil {
		return err
	}

	*o = AnchoreImage(varAnchoreImage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "image_content")
		delete(additionalProperties, "image_detail")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "image_digest")
		delete(additionalProperties, "parent_digest")
		delete(additionalProperties, "account_name")
		delete(additionalProperties, "annotations")
		delete(additionalProperties, "image_status")
		delete(additionalProperties, "analysis_status")
		delete(additionalProperties, "record_version")
		delete(additionalProperties, "analysis_status_detail")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAnchoreImage struct {
	value *AnchoreImage
	isSet bool
}

func (v NullableAnchoreImage) Get() *AnchoreImage {
	return v.value
}

func (v *NullableAnchoreImage) Set(val *AnchoreImage) {
	v.value = val
	v.isSet = true
}

func (v NullableAnchoreImage) IsSet() bool {
	return v.isSet
}

func (v *NullableAnchoreImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnchoreImage(val *AnchoreImage) *NullableAnchoreImage {
	return &NullableAnchoreImage{value: val, isSet: true}
}

func (v NullableAnchoreImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnchoreImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


