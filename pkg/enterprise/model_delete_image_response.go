/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.10.1
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the DeleteImageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteImageResponse{}

// DeleteImageResponse Image deletion response containing status and details
type DeleteImageResponse struct {
	ImageDigest string `json:"image_digest"`
	// Current status of the image deletion
	Status string `json:"status"`
	Detail *string `json:"detail,omitempty"`
}

type _DeleteImageResponse DeleteImageResponse

// NewDeleteImageResponse instantiates a new DeleteImageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteImageResponse(imageDigest string, status string) *DeleteImageResponse {
	this := DeleteImageResponse{}
	this.ImageDigest = imageDigest
	this.Status = status
	return &this
}

// NewDeleteImageResponseWithDefaults instantiates a new DeleteImageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteImageResponseWithDefaults() *DeleteImageResponse {
	this := DeleteImageResponse{}
	return &this
}

// GetImageDigest returns the ImageDigest field value
func (o *DeleteImageResponse) GetImageDigest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageDigest
}

// GetImageDigestOk returns a tuple with the ImageDigest field value
// and a boolean to check if the value has been set.
func (o *DeleteImageResponse) GetImageDigestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageDigest, true
}

// SetImageDigest sets field value
func (o *DeleteImageResponse) SetImageDigest(v string) {
	o.ImageDigest = v
}

// GetStatus returns the Status field value
func (o *DeleteImageResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DeleteImageResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DeleteImageResponse) SetStatus(v string) {
	o.Status = v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *DeleteImageResponse) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteImageResponse) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *DeleteImageResponse) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *DeleteImageResponse) SetDetail(v string) {
	o.Detail = &v
}

func (o DeleteImageResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteImageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["image_digest"] = o.ImageDigest
	toSerialize["status"] = o.Status
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	return toSerialize, nil
}

func (o *DeleteImageResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"image_digest",
		"status",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDeleteImageResponse := _DeleteImageResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeleteImageResponse)

	if err != nil {
		return err
	}

	*o = DeleteImageResponse(varDeleteImageResponse)

	return err
}

type NullableDeleteImageResponse struct {
	value *DeleteImageResponse
	isSet bool
}

func (v NullableDeleteImageResponse) Get() *DeleteImageResponse {
	return v.value
}

func (v *NullableDeleteImageResponse) Set(val *DeleteImageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteImageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteImageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteImageResponse(val *DeleteImageResponse) *NullableDeleteImageResponse {
	return &NullableDeleteImageResponse{value: val, isSet: true}
}

func (v NullableDeleteImageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteImageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


