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

// AllowlistItem Identifies a specific gate and trigger match from a policy against an image and indicates it should be ignored in final policy decisions
type AllowlistItem struct {
	Id *string `json:"id,omitempty"`
	Gate string `json:"gate"`
	TriggerId string `json:"trigger_id"`
	ExpiresOn *time.Time `json:"expires_on,omitempty"`
	// Description of the Allowlist item, human readable
	Description *string `json:"description,omitempty"`
}

// NewAllowlistItem instantiates a new AllowlistItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllowlistItem(gate string, triggerId string) *AllowlistItem {
	this := AllowlistItem{}
	this.Gate = gate
	this.TriggerId = triggerId
	return &this
}

// NewAllowlistItemWithDefaults instantiates a new AllowlistItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllowlistItemWithDefaults() *AllowlistItem {
	this := AllowlistItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AllowlistItem) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowlistItem) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AllowlistItem) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AllowlistItem) SetId(v string) {
	o.Id = &v
}

// GetGate returns the Gate field value
func (o *AllowlistItem) GetGate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gate
}

// GetGateOk returns a tuple with the Gate field value
// and a boolean to check if the value has been set.
func (o *AllowlistItem) GetGateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Gate, true
}

// SetGate sets field value
func (o *AllowlistItem) SetGate(v string) {
	o.Gate = v
}

// GetTriggerId returns the TriggerId field value
func (o *AllowlistItem) GetTriggerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TriggerId
}

// GetTriggerIdOk returns a tuple with the TriggerId field value
// and a boolean to check if the value has been set.
func (o *AllowlistItem) GetTriggerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TriggerId, true
}

// SetTriggerId sets field value
func (o *AllowlistItem) SetTriggerId(v string) {
	o.TriggerId = v
}

// GetExpiresOn returns the ExpiresOn field value if set, zero value otherwise.
func (o *AllowlistItem) GetExpiresOn() time.Time {
	if o == nil || o.ExpiresOn == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresOn
}

// GetExpiresOnOk returns a tuple with the ExpiresOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowlistItem) GetExpiresOnOk() (*time.Time, bool) {
	if o == nil || o.ExpiresOn == nil {
		return nil, false
	}
	return o.ExpiresOn, true
}

// HasExpiresOn returns a boolean if a field has been set.
func (o *AllowlistItem) HasExpiresOn() bool {
	if o != nil && o.ExpiresOn != nil {
		return true
	}

	return false
}

// SetExpiresOn gets a reference to the given time.Time and assigns it to the ExpiresOn field.
func (o *AllowlistItem) SetExpiresOn(v time.Time) {
	o.ExpiresOn = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AllowlistItem) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowlistItem) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AllowlistItem) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AllowlistItem) SetDescription(v string) {
	o.Description = &v
}

func (o AllowlistItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["gate"] = o.Gate
	}
	if true {
		toSerialize["trigger_id"] = o.TriggerId
	}
	if o.ExpiresOn != nil {
		toSerialize["expires_on"] = o.ExpiresOn
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableAllowlistItem struct {
	value *AllowlistItem
	isSet bool
}

func (v NullableAllowlistItem) Get() *AllowlistItem {
	return v.value
}

func (v *NullableAllowlistItem) Set(val *AllowlistItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAllowlistItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAllowlistItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllowlistItem(val *AllowlistItem) *NullableAllowlistItem {
	return &NullableAllowlistItem{value: val, isSet: true}
}

func (v NullableAllowlistItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllowlistItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


