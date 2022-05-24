/*
Anchore Engine API Server

This is the Anchore Engine API. Provides the primary external API for users of the service.

API version: 0.1.20
Contact: nurmi@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package engine

import (
	"encoding/json"
)

// Whitelist A collection of whitelist items to match a policy evaluation against.
type Whitelist struct {
	Id string `json:"id"`
	Name *string `json:"name,omitempty"`
	Version string `json:"version"`
	Comment *string `json:"comment,omitempty"`
	Items *[]WhitelistItem `json:"items,omitempty"`
}

// NewWhitelist instantiates a new Whitelist object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhitelist(id string, version string) *Whitelist {
	this := Whitelist{}
	this.Id = id
	this.Version = version
	return &this
}

// NewWhitelistWithDefaults instantiates a new Whitelist object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhitelistWithDefaults() *Whitelist {
	this := Whitelist{}
	return &this
}

// GetId returns the Id field value
func (o *Whitelist) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Whitelist) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Whitelist) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Whitelist) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Whitelist) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Whitelist) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Whitelist) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value
func (o *Whitelist) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *Whitelist) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *Whitelist) SetVersion(v string) {
	o.Version = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *Whitelist) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Whitelist) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *Whitelist) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *Whitelist) SetComment(v string) {
	o.Comment = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *Whitelist) GetItems() []WhitelistItem {
	if o == nil || o.Items == nil {
		var ret []WhitelistItem
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Whitelist) GetItemsOk() (*[]WhitelistItem, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *Whitelist) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []WhitelistItem and assigns it to the Items field.
func (o *Whitelist) SetItems(v []WhitelistItem) {
	o.Items = &v
}

func (o Whitelist) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableWhitelist struct {
	value *Whitelist
	isSet bool
}

func (v NullableWhitelist) Get() *Whitelist {
	return v.value
}

func (v *NullableWhitelist) Set(val *Whitelist) {
	v.value = val
	v.isSet = true
}

func (v NullableWhitelist) IsSet() bool {
	return v.isSet
}

func (v *NullableWhitelist) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhitelist(val *Whitelist) *NullableWhitelist {
	return &NullableWhitelist{value: val, isSet: true}
}

func (v NullableWhitelist) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhitelist) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


