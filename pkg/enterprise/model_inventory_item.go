/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.6.1
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// InventoryItem Inventory result for a specific Namespace
type InventoryItem struct {
	InventoryType *string `json:"inventory_type,omitempty"`
	Context *string `json:"context,omitempty"`
	ImageTag *string `json:"image_tag,omitempty"`
	ImageDigest *string `json:"image_digest,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
	LastSeen *time.Time `json:"last_seen,omitempty"`
}

// NewInventoryItem instantiates a new InventoryItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryItem() *InventoryItem {
	this := InventoryItem{}
	return &this
}

// NewInventoryItemWithDefaults instantiates a new InventoryItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryItemWithDefaults() *InventoryItem {
	this := InventoryItem{}
	return &this
}

// GetInventoryType returns the InventoryType field value if set, zero value otherwise.
func (o *InventoryItem) GetInventoryType() string {
	if o == nil || o.InventoryType == nil {
		var ret string
		return ret
	}
	return *o.InventoryType
}

// GetInventoryTypeOk returns a tuple with the InventoryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetInventoryTypeOk() (*string, bool) {
	if o == nil || o.InventoryType == nil {
		return nil, false
	}
	return o.InventoryType, true
}

// HasInventoryType returns a boolean if a field has been set.
func (o *InventoryItem) HasInventoryType() bool {
	if o != nil && o.InventoryType != nil {
		return true
	}

	return false
}

// SetInventoryType gets a reference to the given string and assigns it to the InventoryType field.
func (o *InventoryItem) SetInventoryType(v string) {
	o.InventoryType = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *InventoryItem) GetContext() string {
	if o == nil || o.Context == nil {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetContextOk() (*string, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *InventoryItem) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *InventoryItem) SetContext(v string) {
	o.Context = &v
}

// GetImageTag returns the ImageTag field value if set, zero value otherwise.
func (o *InventoryItem) GetImageTag() string {
	if o == nil || o.ImageTag == nil {
		var ret string
		return ret
	}
	return *o.ImageTag
}

// GetImageTagOk returns a tuple with the ImageTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetImageTagOk() (*string, bool) {
	if o == nil || o.ImageTag == nil {
		return nil, false
	}
	return o.ImageTag, true
}

// HasImageTag returns a boolean if a field has been set.
func (o *InventoryItem) HasImageTag() bool {
	if o != nil && o.ImageTag != nil {
		return true
	}

	return false
}

// SetImageTag gets a reference to the given string and assigns it to the ImageTag field.
func (o *InventoryItem) SetImageTag(v string) {
	o.ImageTag = &v
}

// GetImageDigest returns the ImageDigest field value if set, zero value otherwise.
func (o *InventoryItem) GetImageDigest() string {
	if o == nil || o.ImageDigest == nil {
		var ret string
		return ret
	}
	return *o.ImageDigest
}

// GetImageDigestOk returns a tuple with the ImageDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetImageDigestOk() (*string, bool) {
	if o == nil || o.ImageDigest == nil {
		return nil, false
	}
	return o.ImageDigest, true
}

// HasImageDigest returns a boolean if a field has been set.
func (o *InventoryItem) HasImageDigest() bool {
	if o != nil && o.ImageDigest != nil {
		return true
	}

	return false
}

// SetImageDigest gets a reference to the given string and assigns it to the ImageDigest field.
func (o *InventoryItem) SetImageDigest(v string) {
	o.ImageDigest = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *InventoryItem) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *InventoryItem) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *InventoryItem) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *InventoryItem) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *InventoryItem) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *InventoryItem) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetLastSeen returns the LastSeen field value if set, zero value otherwise.
func (o *InventoryItem) GetLastSeen() time.Time {
	if o == nil || o.LastSeen == nil {
		var ret time.Time
		return ret
	}
	return *o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetLastSeenOk() (*time.Time, bool) {
	if o == nil || o.LastSeen == nil {
		return nil, false
	}
	return o.LastSeen, true
}

// HasLastSeen returns a boolean if a field has been set.
func (o *InventoryItem) HasLastSeen() bool {
	if o != nil && o.LastSeen != nil {
		return true
	}

	return false
}

// SetLastSeen gets a reference to the given time.Time and assigns it to the LastSeen field.
func (o *InventoryItem) SetLastSeen(v time.Time) {
	o.LastSeen = &v
}

func (o InventoryItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InventoryType != nil {
		toSerialize["inventory_type"] = o.InventoryType
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if o.ImageTag != nil {
		toSerialize["image_tag"] = o.ImageTag
	}
	if o.ImageDigest != nil {
		toSerialize["image_digest"] = o.ImageDigest
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if o.LastSeen != nil {
		toSerialize["last_seen"] = o.LastSeen
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryItem struct {
	value *InventoryItem
	isSet bool
}

func (v NullableInventoryItem) Get() *InventoryItem {
	return v.value
}

func (v *NullableInventoryItem) Set(val *InventoryItem) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryItem) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryItem(val *InventoryItem) *NullableInventoryItem {
	return &NullableInventoryItem{value: val, isSet: true}
}

func (v NullableInventoryItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


