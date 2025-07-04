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
	"time"
	"bytes"
	"fmt"
)

// checks if the Correction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Correction{}

// Correction Defines a correction object for false positive management
type Correction struct {
	// Identifier for the correction
	Uuid *string `json:"uuid,omitempty"`
	// Type of correction
	Type string `json:"type"`
	Description *string `json:"description,omitempty"`
	Match CorrectionMatch `json:"match"`
	Replace []CorrectionFieldMatch `json:"replace"`
	// RFC 3339 formatted UTC timestamp when the correction was generated
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// RFC 3339 formatted UTC timestamp when the correction was last modified
	LastUpdated *time.Time `json:"last_updated,omitempty"`
}

type _Correction Correction

// NewCorrection instantiates a new Correction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCorrection(type_ string, match CorrectionMatch, replace []CorrectionFieldMatch) *Correction {
	this := Correction{}
	this.Type = type_
	this.Match = match
	this.Replace = replace
	return &this
}

// NewCorrectionWithDefaults instantiates a new Correction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCorrectionWithDefaults() *Correction {
	this := Correction{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *Correction) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Correction) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *Correction) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *Correction) SetUuid(v string) {
	o.Uuid = &v
}

// GetType returns the Type field value
func (o *Correction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Correction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Correction) SetType(v string) {
	o.Type = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Correction) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Correction) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Correction) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Correction) SetDescription(v string) {
	o.Description = &v
}

// GetMatch returns the Match field value
func (o *Correction) GetMatch() CorrectionMatch {
	if o == nil {
		var ret CorrectionMatch
		return ret
	}

	return o.Match
}

// GetMatchOk returns a tuple with the Match field value
// and a boolean to check if the value has been set.
func (o *Correction) GetMatchOk() (*CorrectionMatch, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Match, true
}

// SetMatch sets field value
func (o *Correction) SetMatch(v CorrectionMatch) {
	o.Match = v
}

// GetReplace returns the Replace field value
func (o *Correction) GetReplace() []CorrectionFieldMatch {
	if o == nil {
		var ret []CorrectionFieldMatch
		return ret
	}

	return o.Replace
}

// GetReplaceOk returns a tuple with the Replace field value
// and a boolean to check if the value has been set.
func (o *Correction) GetReplaceOk() ([]CorrectionFieldMatch, bool) {
	if o == nil {
		return nil, false
	}
	return o.Replace, true
}

// SetReplace sets field value
func (o *Correction) SetReplace(v []CorrectionFieldMatch) {
	o.Replace = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Correction) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Correction) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Correction) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Correction) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *Correction) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Correction) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *Correction) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *Correction) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o Correction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Correction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["match"] = o.Match
	toSerialize["replace"] = o.Replace
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["last_updated"] = o.LastUpdated
	}
	return toSerialize, nil
}

func (o *Correction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"match",
		"replace",
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

	varCorrection := _Correction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCorrection)

	if err != nil {
		return err
	}

	*o = Correction(varCorrection)

	return err
}

type NullableCorrection struct {
	value *Correction
	isSet bool
}

func (v NullableCorrection) Get() *Correction {
	return v.value
}

func (v *NullableCorrection) Set(val *Correction) {
	v.value = val
	v.isSet = true
}

func (v NullableCorrection) IsSet() bool {
	return v.isSet
}

func (v *NullableCorrection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCorrection(val *Correction) *NullableCorrection {
	return &NullableCorrection{value: val, isSet: true}
}

func (v NullableCorrection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCorrection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


