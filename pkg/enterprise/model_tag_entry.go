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

// TagEntry A docker-pullable tag value as well as deconstructed components
type TagEntry struct {
	// The pullable string for the tag. E.g. \"docker.io/library/node:latest\"
	Fulltag *string `json:"fulltag,omitempty"`
	// The registry hostname:port section of the pull string
	Registry *string `json:"registry,omitempty"`
	// The repository section of the pull string
	Repo *string `json:"repo,omitempty"`
	// The tag-only section of the pull string
	Tag *string `json:"tag,omitempty"`
	// The timestamp at which the Anchore Engine detected this tag was mapped to the image digest. Does not necessarily indicate when the tag was actually pushed to the registry.
	TagDetectedAt *time.Time `json:"tag_detected_at,omitempty"`
}

// NewTagEntry instantiates a new TagEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagEntry() *TagEntry {
	this := TagEntry{}
	return &this
}

// NewTagEntryWithDefaults instantiates a new TagEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagEntryWithDefaults() *TagEntry {
	this := TagEntry{}
	return &this
}

// GetFulltag returns the Fulltag field value if set, zero value otherwise.
func (o *TagEntry) GetFulltag() string {
	if o == nil || o.Fulltag == nil {
		var ret string
		return ret
	}
	return *o.Fulltag
}

// GetFulltagOk returns a tuple with the Fulltag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagEntry) GetFulltagOk() (*string, bool) {
	if o == nil || o.Fulltag == nil {
		return nil, false
	}
	return o.Fulltag, true
}

// HasFulltag returns a boolean if a field has been set.
func (o *TagEntry) HasFulltag() bool {
	if o != nil && o.Fulltag != nil {
		return true
	}

	return false
}

// SetFulltag gets a reference to the given string and assigns it to the Fulltag field.
func (o *TagEntry) SetFulltag(v string) {
	o.Fulltag = &v
}

// GetRegistry returns the Registry field value if set, zero value otherwise.
func (o *TagEntry) GetRegistry() string {
	if o == nil || o.Registry == nil {
		var ret string
		return ret
	}
	return *o.Registry
}

// GetRegistryOk returns a tuple with the Registry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagEntry) GetRegistryOk() (*string, bool) {
	if o == nil || o.Registry == nil {
		return nil, false
	}
	return o.Registry, true
}

// HasRegistry returns a boolean if a field has been set.
func (o *TagEntry) HasRegistry() bool {
	if o != nil && o.Registry != nil {
		return true
	}

	return false
}

// SetRegistry gets a reference to the given string and assigns it to the Registry field.
func (o *TagEntry) SetRegistry(v string) {
	o.Registry = &v
}

// GetRepo returns the Repo field value if set, zero value otherwise.
func (o *TagEntry) GetRepo() string {
	if o == nil || o.Repo == nil {
		var ret string
		return ret
	}
	return *o.Repo
}

// GetRepoOk returns a tuple with the Repo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagEntry) GetRepoOk() (*string, bool) {
	if o == nil || o.Repo == nil {
		return nil, false
	}
	return o.Repo, true
}

// HasRepo returns a boolean if a field has been set.
func (o *TagEntry) HasRepo() bool {
	if o != nil && o.Repo != nil {
		return true
	}

	return false
}

// SetRepo gets a reference to the given string and assigns it to the Repo field.
func (o *TagEntry) SetRepo(v string) {
	o.Repo = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *TagEntry) GetTag() string {
	if o == nil || o.Tag == nil {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagEntry) GetTagOk() (*string, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *TagEntry) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *TagEntry) SetTag(v string) {
	o.Tag = &v
}

// GetTagDetectedAt returns the TagDetectedAt field value if set, zero value otherwise.
func (o *TagEntry) GetTagDetectedAt() time.Time {
	if o == nil || o.TagDetectedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.TagDetectedAt
}

// GetTagDetectedAtOk returns a tuple with the TagDetectedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagEntry) GetTagDetectedAtOk() (*time.Time, bool) {
	if o == nil || o.TagDetectedAt == nil {
		return nil, false
	}
	return o.TagDetectedAt, true
}

// HasTagDetectedAt returns a boolean if a field has been set.
func (o *TagEntry) HasTagDetectedAt() bool {
	if o != nil && o.TagDetectedAt != nil {
		return true
	}

	return false
}

// SetTagDetectedAt gets a reference to the given time.Time and assigns it to the TagDetectedAt field.
func (o *TagEntry) SetTagDetectedAt(v time.Time) {
	o.TagDetectedAt = &v
}

func (o TagEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Fulltag != nil {
		toSerialize["fulltag"] = o.Fulltag
	}
	if o.Registry != nil {
		toSerialize["registry"] = o.Registry
	}
	if o.Repo != nil {
		toSerialize["repo"] = o.Repo
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	if o.TagDetectedAt != nil {
		toSerialize["tag_detected_at"] = o.TagDetectedAt
	}
	return json.Marshal(toSerialize)
}

type NullableTagEntry struct {
	value *TagEntry
	isSet bool
}

func (v NullableTagEntry) Get() *TagEntry {
	return v.value
}

func (v *NullableTagEntry) Set(val *TagEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableTagEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableTagEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagEntry(val *TagEntry) *NullableTagEntry {
	return &NullableTagEntry{value: val, isSet: true}
}

func (v NullableTagEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


