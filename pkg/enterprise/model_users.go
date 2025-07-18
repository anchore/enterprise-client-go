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

// checks if the Users type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Users{}

// Users a list of User objects
type Users struct {
	Items []User `json:"items,omitempty"`
	TotalRows *int32 `json:"total_rows,omitempty"`
}

// NewUsers instantiates a new Users object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsers() *Users {
	this := Users{}
	return &this
}

// NewUsersWithDefaults instantiates a new Users object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersWithDefaults() *Users {
	this := Users{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *Users) GetItems() []User {
	if o == nil || IsNil(o.Items) {
		var ret []User
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Users) GetItemsOk() ([]User, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *Users) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []User and assigns it to the Items field.
func (o *Users) SetItems(v []User) {
	o.Items = v
}

// GetTotalRows returns the TotalRows field value if set, zero value otherwise.
func (o *Users) GetTotalRows() int32 {
	if o == nil || IsNil(o.TotalRows) {
		var ret int32
		return ret
	}
	return *o.TotalRows
}

// GetTotalRowsOk returns a tuple with the TotalRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Users) GetTotalRowsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalRows) {
		return nil, false
	}
	return o.TotalRows, true
}

// HasTotalRows returns a boolean if a field has been set.
func (o *Users) HasTotalRows() bool {
	if o != nil && !IsNil(o.TotalRows) {
		return true
	}

	return false
}

// SetTotalRows gets a reference to the given int32 and assigns it to the TotalRows field.
func (o *Users) SetTotalRows(v int32) {
	o.TotalRows = &v
}

func (o Users) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Users) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.TotalRows) {
		toSerialize["total_rows"] = o.TotalRows
	}
	return toSerialize, nil
}

type NullableUsers struct {
	value *Users
	isSet bool
}

func (v NullableUsers) Get() *Users {
	return v.value
}

func (v *NullableUsers) Set(val *Users) {
	v.value = val
	v.isSet = true
}

func (v NullableUsers) IsSet() bool {
	return v.isSet
}

func (v *NullableUsers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsers(val *Users) *NullableUsers {
	return &NullableUsers{value: val, isSet: true}
}

func (v NullableUsers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


