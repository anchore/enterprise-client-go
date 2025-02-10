/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.10.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UserCreationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserCreationRequest{}

// UserCreationRequest A payload for creating a new user, includes the username and password in a single request
type UserCreationRequest struct {
	// The username for authentication. If the user_type is 'native', this name must not contain a colon character as per RFC 2617 (HTTP Basic and Digest Authentication). If the user_type is 'saml', then colons are allowed in the name since HTTP Basic auth is not used for that user type.
	Username string "json:\"username\" validate:\"regexp=^[a-zA-Z0-9][ a-zA-Z0-9@.!#$+=^_`~;:-]{1,126}[a-zA-Z0-9_]$\""
	// The initial password for the user, must be at least 6 characters, up to 128. This must be null when the user_type is not 'native'.
	Password *string `json:"password,omitempty" validate:"regexp=.{6,128}$"`
	// The user's type. A Native user authenticates using user/password log on. All other users will authenticate with an IDP.
	UserType *string `json:"user_type,omitempty"`
	// If the user is authenticating via an IDP, this is the name of the IDP. A 'native' user should have this set to null.
	IdpName *string `json:"idp_name,omitempty"`
}

type _UserCreationRequest UserCreationRequest

// NewUserCreationRequest instantiates a new UserCreationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserCreationRequest(username string) *UserCreationRequest {
	this := UserCreationRequest{}
	this.Username = username
	return &this
}

// NewUserCreationRequestWithDefaults instantiates a new UserCreationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserCreationRequestWithDefaults() *UserCreationRequest {
	this := UserCreationRequest{}
	return &this
}

// GetUsername returns the Username field value
func (o *UserCreationRequest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *UserCreationRequest) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *UserCreationRequest) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UserCreationRequest) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreationRequest) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UserCreationRequest) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UserCreationRequest) SetPassword(v string) {
	o.Password = &v
}

// GetUserType returns the UserType field value if set, zero value otherwise.
func (o *UserCreationRequest) GetUserType() string {
	if o == nil || IsNil(o.UserType) {
		var ret string
		return ret
	}
	return *o.UserType
}

// GetUserTypeOk returns a tuple with the UserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreationRequest) GetUserTypeOk() (*string, bool) {
	if o == nil || IsNil(o.UserType) {
		return nil, false
	}
	return o.UserType, true
}

// HasUserType returns a boolean if a field has been set.
func (o *UserCreationRequest) HasUserType() bool {
	if o != nil && !IsNil(o.UserType) {
		return true
	}

	return false
}

// SetUserType gets a reference to the given string and assigns it to the UserType field.
func (o *UserCreationRequest) SetUserType(v string) {
	o.UserType = &v
}

// GetIdpName returns the IdpName field value if set, zero value otherwise.
func (o *UserCreationRequest) GetIdpName() string {
	if o == nil || IsNil(o.IdpName) {
		var ret string
		return ret
	}
	return *o.IdpName
}

// GetIdpNameOk returns a tuple with the IdpName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreationRequest) GetIdpNameOk() (*string, bool) {
	if o == nil || IsNil(o.IdpName) {
		return nil, false
	}
	return o.IdpName, true
}

// HasIdpName returns a boolean if a field has been set.
func (o *UserCreationRequest) HasIdpName() bool {
	if o != nil && !IsNil(o.IdpName) {
		return true
	}

	return false
}

// SetIdpName gets a reference to the given string and assigns it to the IdpName field.
func (o *UserCreationRequest) SetIdpName(v string) {
	o.IdpName = &v
}

func (o UserCreationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserCreationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["username"] = o.Username
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.UserType) {
		toSerialize["user_type"] = o.UserType
	}
	if !IsNil(o.IdpName) {
		toSerialize["idp_name"] = o.IdpName
	}
	return toSerialize, nil
}

func (o *UserCreationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"username",
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

	varUserCreationRequest := _UserCreationRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserCreationRequest)

	if err != nil {
		return err
	}

	*o = UserCreationRequest(varUserCreationRequest)

	return err
}

type NullableUserCreationRequest struct {
	value *UserCreationRequest
	isSet bool
}

func (v NullableUserCreationRequest) Get() *UserCreationRequest {
	return v.value
}

func (v *NullableUserCreationRequest) Set(val *UserCreationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCreationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCreationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCreationRequest(val *UserCreationRequest) *NullableUserCreationRequest {
	return &NullableUserCreationRequest{value: val, isSet: true}
}

func (v NullableUserCreationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCreationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


