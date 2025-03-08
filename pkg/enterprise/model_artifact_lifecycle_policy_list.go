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
)

// checks if the ArtifactLifecyclePolicyList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArtifactLifecyclePolicyList{}

// ArtifactLifecyclePolicyList struct for ArtifactLifecyclePolicyList
type ArtifactLifecyclePolicyList struct {
	Items []ArtifactLifecyclePolicyResponse `json:"items,omitempty"`
}

// NewArtifactLifecyclePolicyList instantiates a new ArtifactLifecyclePolicyList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtifactLifecyclePolicyList() *ArtifactLifecyclePolicyList {
	this := ArtifactLifecyclePolicyList{}
	return &this
}

// NewArtifactLifecyclePolicyListWithDefaults instantiates a new ArtifactLifecyclePolicyList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtifactLifecyclePolicyListWithDefaults() *ArtifactLifecyclePolicyList {
	this := ArtifactLifecyclePolicyList{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ArtifactLifecyclePolicyList) GetItems() []ArtifactLifecyclePolicyResponse {
	if o == nil || IsNil(o.Items) {
		var ret []ArtifactLifecyclePolicyResponse
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactLifecyclePolicyList) GetItemsOk() ([]ArtifactLifecyclePolicyResponse, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ArtifactLifecyclePolicyList) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ArtifactLifecyclePolicyResponse and assigns it to the Items field.
func (o *ArtifactLifecyclePolicyList) SetItems(v []ArtifactLifecyclePolicyResponse) {
	o.Items = v
}

func (o ArtifactLifecyclePolicyList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArtifactLifecyclePolicyList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableArtifactLifecyclePolicyList struct {
	value *ArtifactLifecyclePolicyList
	isSet bool
}

func (v NullableArtifactLifecyclePolicyList) Get() *ArtifactLifecyclePolicyList {
	return v.value
}

func (v *NullableArtifactLifecyclePolicyList) Set(val *ArtifactLifecyclePolicyList) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifactLifecyclePolicyList) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifactLifecyclePolicyList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifactLifecyclePolicyList(val *ArtifactLifecyclePolicyList) *NullableArtifactLifecyclePolicyList {
	return &NullableArtifactLifecyclePolicyList{value: val, isSet: true}
}

func (v NullableArtifactLifecyclePolicyList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtifactLifecyclePolicyList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


