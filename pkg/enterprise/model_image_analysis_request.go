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
)

// checks if the ImageAnalysisRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageAnalysisRequest{}

// ImageAnalysisRequest A request to add an image to be watched and analyzed by the engine.
type ImageAnalysisRequest struct {
	// Optional. The type of image this is adding, defaults to \"docker\".
	ImageType *string `json:"image_type,omitempty"`
	// Annotations to be associated with the added image in key/value form
	Annotations interface{} `json:"annotations,omitempty"`
	Source *ImageSource `json:"source,omitempty"`
}

// NewImageAnalysisRequest instantiates a new ImageAnalysisRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageAnalysisRequest() *ImageAnalysisRequest {
	this := ImageAnalysisRequest{}
	return &this
}

// NewImageAnalysisRequestWithDefaults instantiates a new ImageAnalysisRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageAnalysisRequestWithDefaults() *ImageAnalysisRequest {
	this := ImageAnalysisRequest{}
	return &this
}

// GetImageType returns the ImageType field value if set, zero value otherwise.
func (o *ImageAnalysisRequest) GetImageType() string {
	if o == nil || IsNil(o.ImageType) {
		var ret string
		return ret
	}
	return *o.ImageType
}

// GetImageTypeOk returns a tuple with the ImageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageAnalysisRequest) GetImageTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ImageType) {
		return nil, false
	}
	return o.ImageType, true
}

// HasImageType returns a boolean if a field has been set.
func (o *ImageAnalysisRequest) HasImageType() bool {
	if o != nil && !IsNil(o.ImageType) {
		return true
	}

	return false
}

// SetImageType gets a reference to the given string and assigns it to the ImageType field.
func (o *ImageAnalysisRequest) SetImageType(v string) {
	o.ImageType = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *ImageAnalysisRequest) GetAnnotations() interface{} {
	if o == nil || IsNil(o.Annotations) {
		var ret interface{}
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageAnalysisRequest) GetAnnotationsOk() (interface{}, bool) {
	if o == nil || IsNil(o.Annotations) {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *ImageAnalysisRequest) HasAnnotations() bool {
	if o != nil && !IsNil(o.Annotations) {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given interface{} and assigns it to the Annotations field.
func (o *ImageAnalysisRequest) SetAnnotations(v interface{}) {
	o.Annotations = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ImageAnalysisRequest) GetSource() ImageSource {
	if o == nil || IsNil(o.Source) {
		var ret ImageSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageAnalysisRequest) GetSourceOk() (*ImageSource, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ImageAnalysisRequest) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given ImageSource and assigns it to the Source field.
func (o *ImageAnalysisRequest) SetSource(v ImageSource) {
	o.Source = &v
}

func (o ImageAnalysisRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageAnalysisRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ImageType) {
		toSerialize["image_type"] = o.ImageType
	}
	if !IsNil(o.Annotations) {
		toSerialize["annotations"] = o.Annotations
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	return toSerialize, nil
}

type NullableImageAnalysisRequest struct {
	value *ImageAnalysisRequest
	isSet bool
}

func (v NullableImageAnalysisRequest) Get() *ImageAnalysisRequest {
	return v.value
}

func (v *NullableImageAnalysisRequest) Set(val *ImageAnalysisRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableImageAnalysisRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableImageAnalysisRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageAnalysisRequest(val *ImageAnalysisRequest) *NullableImageAnalysisRequest {
	return &NullableImageAnalysisRequest{value: val, isSet: true}
}

func (v NullableImageAnalysisRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageAnalysisRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


