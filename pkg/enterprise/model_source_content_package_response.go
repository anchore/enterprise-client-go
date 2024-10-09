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
)

// SourceContentPackageResponse Package content listings from analysis sbom
type SourceContentPackageResponse struct {
	SourceId *string `json:"source_id,omitempty"`
	ContentType *string `json:"content_type,omitempty"`
	Content []SourceContentPackageResponseContent `json:"content,omitempty"`
}

// NewSourceContentPackageResponse instantiates a new SourceContentPackageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceContentPackageResponse() *SourceContentPackageResponse {
	this := SourceContentPackageResponse{}
	return &this
}

// NewSourceContentPackageResponseWithDefaults instantiates a new SourceContentPackageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceContentPackageResponseWithDefaults() *SourceContentPackageResponse {
	this := SourceContentPackageResponse{}
	return &this
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *SourceContentPackageResponse) GetSourceId() string {
	if o == nil || o.SourceId == nil {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceContentPackageResponse) GetSourceIdOk() (*string, bool) {
	if o == nil || o.SourceId == nil {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *SourceContentPackageResponse) HasSourceId() bool {
	if o != nil && o.SourceId != nil {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *SourceContentPackageResponse) SetSourceId(v string) {
	o.SourceId = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *SourceContentPackageResponse) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceContentPackageResponse) GetContentTypeOk() (*string, bool) {
	if o == nil || o.ContentType == nil {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *SourceContentPackageResponse) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *SourceContentPackageResponse) SetContentType(v string) {
	o.ContentType = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *SourceContentPackageResponse) GetContent() []SourceContentPackageResponseContent {
	if o == nil || o.Content == nil {
		var ret []SourceContentPackageResponseContent
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceContentPackageResponse) GetContentOk() ([]SourceContentPackageResponseContent, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *SourceContentPackageResponse) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given []SourceContentPackageResponseContent and assigns it to the Content field.
func (o *SourceContentPackageResponse) SetContent(v []SourceContentPackageResponseContent) {
	o.Content = v
}

func (o SourceContentPackageResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SourceId != nil {
		toSerialize["source_id"] = o.SourceId
	}
	if o.ContentType != nil {
		toSerialize["content_type"] = o.ContentType
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	return json.Marshal(toSerialize)
}

type NullableSourceContentPackageResponse struct {
	value *SourceContentPackageResponse
	isSet bool
}

func (v NullableSourceContentPackageResponse) Get() *SourceContentPackageResponse {
	return v.value
}

func (v *NullableSourceContentPackageResponse) Set(val *SourceContentPackageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceContentPackageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceContentPackageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceContentPackageResponse(val *SourceContentPackageResponse) *NullableSourceContentPackageResponse {
	return &NullableSourceContentPackageResponse{value: val, isSet: true}
}

func (v NullableSourceContentPackageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceContentPackageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


