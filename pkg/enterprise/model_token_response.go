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
)

// TokenResponse struct for TokenResponse
type TokenResponse struct {
	// The token content
	AccessToken string `json:"access_token"`
	// The lifetime of the access token in seconds
	ExpiresIn *float32 `json:"expires_in,omitempty"`
	// The type of the token returned
	TokenType *string `json:"token_type,omitempty"`
	// The refresh token content
	RefreshToken *string `json:"refresh_token,omitempty"`
}

// NewTokenResponse instantiates a new TokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenResponse(accessToken string) *TokenResponse {
	this := TokenResponse{}
	this.AccessToken = accessToken
	return &this
}

// NewTokenResponseWithDefaults instantiates a new TokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenResponseWithDefaults() *TokenResponse {
	this := TokenResponse{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *TokenResponse) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *TokenResponse) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *TokenResponse) GetExpiresIn() float32 {
	if o == nil || o.ExpiresIn == nil {
		var ret float32
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetExpiresInOk() (*float32, bool) {
	if o == nil || o.ExpiresIn == nil {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *TokenResponse) HasExpiresIn() bool {
	if o != nil && o.ExpiresIn != nil {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given float32 and assigns it to the ExpiresIn field.
func (o *TokenResponse) SetExpiresIn(v float32) {
	o.ExpiresIn = &v
}

// GetTokenType returns the TokenType field value if set, zero value otherwise.
func (o *TokenResponse) GetTokenType() string {
	if o == nil || o.TokenType == nil {
		var ret string
		return ret
	}
	return *o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetTokenTypeOk() (*string, bool) {
	if o == nil || o.TokenType == nil {
		return nil, false
	}
	return o.TokenType, true
}

// HasTokenType returns a boolean if a field has been set.
func (o *TokenResponse) HasTokenType() bool {
	if o != nil && o.TokenType != nil {
		return true
	}

	return false
}

// SetTokenType gets a reference to the given string and assigns it to the TokenType field.
func (o *TokenResponse) SetTokenType(v string) {
	o.TokenType = &v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *TokenResponse) GetRefreshToken() string {
	if o == nil || o.RefreshToken == nil {
		var ret string
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetRefreshTokenOk() (*string, bool) {
	if o == nil || o.RefreshToken == nil {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *TokenResponse) HasRefreshToken() bool {
	if o != nil && o.RefreshToken != nil {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given string and assigns it to the RefreshToken field.
func (o *TokenResponse) SetRefreshToken(v string) {
	o.RefreshToken = &v
}

func (o TokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	if o.ExpiresIn != nil {
		toSerialize["expires_in"] = o.ExpiresIn
	}
	if o.TokenType != nil {
		toSerialize["token_type"] = o.TokenType
	}
	if o.RefreshToken != nil {
		toSerialize["refresh_token"] = o.RefreshToken
	}
	return json.Marshal(toSerialize)
}

type NullableTokenResponse struct {
	value *TokenResponse
	isSet bool
}

func (v NullableTokenResponse) Get() *TokenResponse {
	return v.value
}

func (v *NullableTokenResponse) Set(val *TokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenResponse(val *TokenResponse) *NullableTokenResponse {
	return &NullableTokenResponse{value: val, isSet: true}
}

func (v NullableTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


