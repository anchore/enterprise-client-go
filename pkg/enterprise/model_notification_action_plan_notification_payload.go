/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.3.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// NotificationActionPlanNotificationPayload struct for NotificationActionPlanNotificationPayload
type NotificationActionPlanNotificationPayload struct {
	Type *string `json:"type,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
	ImageTag *string `json:"image_tag,omitempty"`
	ImageDigest *string `json:"image_digest,omitempty"`
	BundleId *string `json:"bundle_id,omitempty"`
	Subject *string `json:"subject,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewNotificationActionPlanNotificationPayload instantiates a new NotificationActionPlanNotificationPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationActionPlanNotificationPayload() *NotificationActionPlanNotificationPayload {
	this := NotificationActionPlanNotificationPayload{}
	return &this
}

// NewNotificationActionPlanNotificationPayloadWithDefaults instantiates a new NotificationActionPlanNotificationPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationActionPlanNotificationPayloadWithDefaults() *NotificationActionPlanNotificationPayload {
	this := NotificationActionPlanNotificationPayload{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NotificationActionPlanNotificationPayload) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationActionPlanNotificationPayload) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NotificationActionPlanNotificationPayload) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NotificationActionPlanNotificationPayload) SetType(v string) {
	o.Type = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *NotificationActionPlanNotificationPayload) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationActionPlanNotificationPayload) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *NotificationActionPlanNotificationPayload) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *NotificationActionPlanNotificationPayload) SetUuid(v string) {
	o.Uuid = &v
}

// GetImageTag returns the ImageTag field value if set, zero value otherwise.
func (o *NotificationActionPlanNotificationPayload) GetImageTag() string {
	if o == nil || o.ImageTag == nil {
		var ret string
		return ret
	}
	return *o.ImageTag
}

// GetImageTagOk returns a tuple with the ImageTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationActionPlanNotificationPayload) GetImageTagOk() (*string, bool) {
	if o == nil || o.ImageTag == nil {
		return nil, false
	}
	return o.ImageTag, true
}

// HasImageTag returns a boolean if a field has been set.
func (o *NotificationActionPlanNotificationPayload) HasImageTag() bool {
	if o != nil && o.ImageTag != nil {
		return true
	}

	return false
}

// SetImageTag gets a reference to the given string and assigns it to the ImageTag field.
func (o *NotificationActionPlanNotificationPayload) SetImageTag(v string) {
	o.ImageTag = &v
}

// GetImageDigest returns the ImageDigest field value if set, zero value otherwise.
func (o *NotificationActionPlanNotificationPayload) GetImageDigest() string {
	if o == nil || o.ImageDigest == nil {
		var ret string
		return ret
	}
	return *o.ImageDigest
}

// GetImageDigestOk returns a tuple with the ImageDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationActionPlanNotificationPayload) GetImageDigestOk() (*string, bool) {
	if o == nil || o.ImageDigest == nil {
		return nil, false
	}
	return o.ImageDigest, true
}

// HasImageDigest returns a boolean if a field has been set.
func (o *NotificationActionPlanNotificationPayload) HasImageDigest() bool {
	if o != nil && o.ImageDigest != nil {
		return true
	}

	return false
}

// SetImageDigest gets a reference to the given string and assigns it to the ImageDigest field.
func (o *NotificationActionPlanNotificationPayload) SetImageDigest(v string) {
	o.ImageDigest = &v
}

// GetBundleId returns the BundleId field value if set, zero value otherwise.
func (o *NotificationActionPlanNotificationPayload) GetBundleId() string {
	if o == nil || o.BundleId == nil {
		var ret string
		return ret
	}
	return *o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationActionPlanNotificationPayload) GetBundleIdOk() (*string, bool) {
	if o == nil || o.BundleId == nil {
		return nil, false
	}
	return o.BundleId, true
}

// HasBundleId returns a boolean if a field has been set.
func (o *NotificationActionPlanNotificationPayload) HasBundleId() bool {
	if o != nil && o.BundleId != nil {
		return true
	}

	return false
}

// SetBundleId gets a reference to the given string and assigns it to the BundleId field.
func (o *NotificationActionPlanNotificationPayload) SetBundleId(v string) {
	o.BundleId = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *NotificationActionPlanNotificationPayload) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationActionPlanNotificationPayload) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *NotificationActionPlanNotificationPayload) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *NotificationActionPlanNotificationPayload) SetSubject(v string) {
	o.Subject = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *NotificationActionPlanNotificationPayload) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationActionPlanNotificationPayload) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *NotificationActionPlanNotificationPayload) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *NotificationActionPlanNotificationPayload) SetMessage(v string) {
	o.Message = &v
}

func (o NotificationActionPlanNotificationPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.ImageTag != nil {
		toSerialize["image_tag"] = o.ImageTag
	}
	if o.ImageDigest != nil {
		toSerialize["image_digest"] = o.ImageDigest
	}
	if o.BundleId != nil {
		toSerialize["bundle_id"] = o.BundleId
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationActionPlanNotificationPayload struct {
	value *NotificationActionPlanNotificationPayload
	isSet bool
}

func (v NullableNotificationActionPlanNotificationPayload) Get() *NotificationActionPlanNotificationPayload {
	return v.value
}

func (v *NullableNotificationActionPlanNotificationPayload) Set(val *NotificationActionPlanNotificationPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationActionPlanNotificationPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationActionPlanNotificationPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationActionPlanNotificationPayload(val *NotificationActionPlanNotificationPayload) *NullableNotificationActionPlanNotificationPayload {
	return &NullableNotificationActionPlanNotificationPayload{value: val, isSet: true}
}

func (v NullableNotificationActionPlanNotificationPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationActionPlanNotificationPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


