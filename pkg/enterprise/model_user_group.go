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
	"time"
	"bytes"
	"fmt"
)

// checks if the UserGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserGroup{}

// UserGroup struct for UserGroup
type UserGroup struct {
	// The name of the user group
	Name string "json:\"name\" validate:\"regexp=^[a-zA-Z0-9][ a-zA-Z0-9@.!#$+=^_`~;-]{1,126}[a-zA-Z0-9_]$\""
	// The description of the user group
	Description *string `json:"description,omitempty"`
	// The unique identifier for the user group
	GroupUuid string `json:"group_uuid"`
	// The timestamp of when the user group was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The timestamp of the last update to this user group
	LastUpdated *time.Time `json:"last_updated,omitempty"`
	AccountRoles *UserGroupRoles `json:"account_roles,omitempty"`
}

type _UserGroup UserGroup

// NewUserGroup instantiates a new UserGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGroup(name string, groupUuid string) *UserGroup {
	this := UserGroup{}
	this.Name = name
	this.GroupUuid = groupUuid
	return &this
}

// NewUserGroupWithDefaults instantiates a new UserGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGroupWithDefaults() *UserGroup {
	this := UserGroup{}
	return &this
}

// GetName returns the Name field value
func (o *UserGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserGroup) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UserGroup) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UserGroup) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UserGroup) SetDescription(v string) {
	o.Description = &v
}

// GetGroupUuid returns the GroupUuid field value
func (o *UserGroup) GetGroupUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupUuid
}

// GetGroupUuidOk returns a tuple with the GroupUuid field value
// and a boolean to check if the value has been set.
func (o *UserGroup) GetGroupUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupUuid, true
}

// SetGroupUuid sets field value
func (o *UserGroup) SetGroupUuid(v string) {
	o.GroupUuid = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *UserGroup) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroup) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *UserGroup) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *UserGroup) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *UserGroup) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroup) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *UserGroup) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *UserGroup) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetAccountRoles returns the AccountRoles field value if set, zero value otherwise.
func (o *UserGroup) GetAccountRoles() UserGroupRoles {
	if o == nil || IsNil(o.AccountRoles) {
		var ret UserGroupRoles
		return ret
	}
	return *o.AccountRoles
}

// GetAccountRolesOk returns a tuple with the AccountRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroup) GetAccountRolesOk() (*UserGroupRoles, bool) {
	if o == nil || IsNil(o.AccountRoles) {
		return nil, false
	}
	return o.AccountRoles, true
}

// HasAccountRoles returns a boolean if a field has been set.
func (o *UserGroup) HasAccountRoles() bool {
	if o != nil && !IsNil(o.AccountRoles) {
		return true
	}

	return false
}

// SetAccountRoles gets a reference to the given UserGroupRoles and assigns it to the AccountRoles field.
func (o *UserGroup) SetAccountRoles(v UserGroupRoles) {
	o.AccountRoles = &v
}

func (o UserGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["group_uuid"] = o.GroupUuid
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if !IsNil(o.AccountRoles) {
		toSerialize["account_roles"] = o.AccountRoles
	}
	return toSerialize, nil
}

func (o *UserGroup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"group_uuid",
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

	varUserGroup := _UserGroup{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserGroup)

	if err != nil {
		return err
	}

	*o = UserGroup(varUserGroup)

	return err
}

type NullableUserGroup struct {
	value *UserGroup
	isSet bool
}

func (v NullableUserGroup) Get() *UserGroup {
	return v.value
}

func (v *NullableUserGroup) Set(val *UserGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGroup(val *UserGroup) *NullableUserGroup {
	return &NullableUserGroup{value: val, isSet: true}
}

func (v NullableUserGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


