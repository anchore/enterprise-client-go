/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.12.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the VulnUpdateNotificationPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VulnUpdateNotificationPayload{}

// VulnUpdateNotificationPayload struct for VulnUpdateNotificationPayload
type VulnUpdateNotificationPayload struct {
	AccountName *string `json:"account_name,omitempty"`
	SubscriptionKey *string `json:"subscription_key,omitempty"`
	SubscriptionType *string `json:"subscription_type,omitempty"`
	NotificationId *string `json:"notification_id,omitempty"`
	DiffVulnerabilityResult *VulnDiffResult `json:"diff_vulnerability_result,omitempty"`
	ImageDigest *string `json:"image_digest,omitempty"`
	// List of Corresponding Image Annotations
	Annotations interface{} `json:"annotations,omitempty"`
}

// NewVulnUpdateNotificationPayload instantiates a new VulnUpdateNotificationPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVulnUpdateNotificationPayload() *VulnUpdateNotificationPayload {
	this := VulnUpdateNotificationPayload{}
	return &this
}

// NewVulnUpdateNotificationPayloadWithDefaults instantiates a new VulnUpdateNotificationPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVulnUpdateNotificationPayloadWithDefaults() *VulnUpdateNotificationPayload {
	this := VulnUpdateNotificationPayload{}
	return &this
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *VulnUpdateNotificationPayload) GetAccountName() string {
	if o == nil || IsNil(o.AccountName) {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnUpdateNotificationPayload) GetAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.AccountName) {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *VulnUpdateNotificationPayload) HasAccountName() bool {
	if o != nil && !IsNil(o.AccountName) {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *VulnUpdateNotificationPayload) SetAccountName(v string) {
	o.AccountName = &v
}

// GetSubscriptionKey returns the SubscriptionKey field value if set, zero value otherwise.
func (o *VulnUpdateNotificationPayload) GetSubscriptionKey() string {
	if o == nil || IsNil(o.SubscriptionKey) {
		var ret string
		return ret
	}
	return *o.SubscriptionKey
}

// GetSubscriptionKeyOk returns a tuple with the SubscriptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnUpdateNotificationPayload) GetSubscriptionKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionKey) {
		return nil, false
	}
	return o.SubscriptionKey, true
}

// HasSubscriptionKey returns a boolean if a field has been set.
func (o *VulnUpdateNotificationPayload) HasSubscriptionKey() bool {
	if o != nil && !IsNil(o.SubscriptionKey) {
		return true
	}

	return false
}

// SetSubscriptionKey gets a reference to the given string and assigns it to the SubscriptionKey field.
func (o *VulnUpdateNotificationPayload) SetSubscriptionKey(v string) {
	o.SubscriptionKey = &v
}

// GetSubscriptionType returns the SubscriptionType field value if set, zero value otherwise.
func (o *VulnUpdateNotificationPayload) GetSubscriptionType() string {
	if o == nil || IsNil(o.SubscriptionType) {
		var ret string
		return ret
	}
	return *o.SubscriptionType
}

// GetSubscriptionTypeOk returns a tuple with the SubscriptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnUpdateNotificationPayload) GetSubscriptionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionType) {
		return nil, false
	}
	return o.SubscriptionType, true
}

// HasSubscriptionType returns a boolean if a field has been set.
func (o *VulnUpdateNotificationPayload) HasSubscriptionType() bool {
	if o != nil && !IsNil(o.SubscriptionType) {
		return true
	}

	return false
}

// SetSubscriptionType gets a reference to the given string and assigns it to the SubscriptionType field.
func (o *VulnUpdateNotificationPayload) SetSubscriptionType(v string) {
	o.SubscriptionType = &v
}

// GetNotificationId returns the NotificationId field value if set, zero value otherwise.
func (o *VulnUpdateNotificationPayload) GetNotificationId() string {
	if o == nil || IsNil(o.NotificationId) {
		var ret string
		return ret
	}
	return *o.NotificationId
}

// GetNotificationIdOk returns a tuple with the NotificationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnUpdateNotificationPayload) GetNotificationIdOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationId) {
		return nil, false
	}
	return o.NotificationId, true
}

// HasNotificationId returns a boolean if a field has been set.
func (o *VulnUpdateNotificationPayload) HasNotificationId() bool {
	if o != nil && !IsNil(o.NotificationId) {
		return true
	}

	return false
}

// SetNotificationId gets a reference to the given string and assigns it to the NotificationId field.
func (o *VulnUpdateNotificationPayload) SetNotificationId(v string) {
	o.NotificationId = &v
}

// GetDiffVulnerabilityResult returns the DiffVulnerabilityResult field value if set, zero value otherwise.
func (o *VulnUpdateNotificationPayload) GetDiffVulnerabilityResult() VulnDiffResult {
	if o == nil || IsNil(o.DiffVulnerabilityResult) {
		var ret VulnDiffResult
		return ret
	}
	return *o.DiffVulnerabilityResult
}

// GetDiffVulnerabilityResultOk returns a tuple with the DiffVulnerabilityResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnUpdateNotificationPayload) GetDiffVulnerabilityResultOk() (*VulnDiffResult, bool) {
	if o == nil || IsNil(o.DiffVulnerabilityResult) {
		return nil, false
	}
	return o.DiffVulnerabilityResult, true
}

// HasDiffVulnerabilityResult returns a boolean if a field has been set.
func (o *VulnUpdateNotificationPayload) HasDiffVulnerabilityResult() bool {
	if o != nil && !IsNil(o.DiffVulnerabilityResult) {
		return true
	}

	return false
}

// SetDiffVulnerabilityResult gets a reference to the given VulnDiffResult and assigns it to the DiffVulnerabilityResult field.
func (o *VulnUpdateNotificationPayload) SetDiffVulnerabilityResult(v VulnDiffResult) {
	o.DiffVulnerabilityResult = &v
}

// GetImageDigest returns the ImageDigest field value if set, zero value otherwise.
func (o *VulnUpdateNotificationPayload) GetImageDigest() string {
	if o == nil || IsNil(o.ImageDigest) {
		var ret string
		return ret
	}
	return *o.ImageDigest
}

// GetImageDigestOk returns a tuple with the ImageDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnUpdateNotificationPayload) GetImageDigestOk() (*string, bool) {
	if o == nil || IsNil(o.ImageDigest) {
		return nil, false
	}
	return o.ImageDigest, true
}

// HasImageDigest returns a boolean if a field has been set.
func (o *VulnUpdateNotificationPayload) HasImageDigest() bool {
	if o != nil && !IsNil(o.ImageDigest) {
		return true
	}

	return false
}

// SetImageDigest gets a reference to the given string and assigns it to the ImageDigest field.
func (o *VulnUpdateNotificationPayload) SetImageDigest(v string) {
	o.ImageDigest = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VulnUpdateNotificationPayload) GetAnnotations() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VulnUpdateNotificationPayload) GetAnnotationsOk() (interface{}, bool) {
	if o == nil || IsNil(o.Annotations) {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *VulnUpdateNotificationPayload) HasAnnotations() bool {
	if o != nil && !IsNil(o.Annotations) {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given interface{} and assigns it to the Annotations field.
func (o *VulnUpdateNotificationPayload) SetAnnotations(v interface{}) {
	o.Annotations = v
}

func (o VulnUpdateNotificationPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VulnUpdateNotificationPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountName) {
		toSerialize["account_name"] = o.AccountName
	}
	if !IsNil(o.SubscriptionKey) {
		toSerialize["subscription_key"] = o.SubscriptionKey
	}
	if !IsNil(o.SubscriptionType) {
		toSerialize["subscription_type"] = o.SubscriptionType
	}
	if !IsNil(o.NotificationId) {
		toSerialize["notification_id"] = o.NotificationId
	}
	if !IsNil(o.DiffVulnerabilityResult) {
		toSerialize["diff_vulnerability_result"] = o.DiffVulnerabilityResult
	}
	if !IsNil(o.ImageDigest) {
		toSerialize["image_digest"] = o.ImageDigest
	}
	if o.Annotations != nil {
		toSerialize["annotations"] = o.Annotations
	}
	return toSerialize, nil
}

type NullableVulnUpdateNotificationPayload struct {
	value *VulnUpdateNotificationPayload
	isSet bool
}

func (v NullableVulnUpdateNotificationPayload) Get() *VulnUpdateNotificationPayload {
	return v.value
}

func (v *NullableVulnUpdateNotificationPayload) Set(val *VulnUpdateNotificationPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableVulnUpdateNotificationPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableVulnUpdateNotificationPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVulnUpdateNotificationPayload(val *VulnUpdateNotificationPayload) *NullableVulnUpdateNotificationPayload {
	return &NullableVulnUpdateNotificationPayload{value: val, isSet: true}
}

func (v NullableVulnUpdateNotificationPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVulnUpdateNotificationPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


