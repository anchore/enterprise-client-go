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
	"fmt"
)

// PatchUserApiKeyRequest struct for PatchUserApiKeyRequest
type PatchUserApiKeyRequest struct {
	inter interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PatchUserApiKeyRequest) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into interface{}
	err = json.Unmarshal(data, &dst.inter);
	if err == nil {
		jsoninterface, _ := json.Marshal(dst.inter)
		if string(jsoninterface) == "{}" { // empty struct
			dst.inter = nil
		} else {
			return nil // data stored in dst.interface{}, return on the first match
		}
	} else {
		dst.inter = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(PatchUserApiKeyRequest)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PatchUserApiKeyRequest) MarshalJSON() ([]byte, error) {
	if src.inter != nil {
		return json.Marshal(&src.inter)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePatchUserApiKeyRequest struct {
	value *PatchUserApiKeyRequest
	isSet bool
}

func (v NullablePatchUserApiKeyRequest) Get() *PatchUserApiKeyRequest {
	return v.value
}

func (v *NullablePatchUserApiKeyRequest) Set(val *PatchUserApiKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchUserApiKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchUserApiKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchUserApiKeyRequest(val *PatchUserApiKeyRequest) *NullablePatchUserApiKeyRequest {
	return &NullablePatchUserApiKeyRequest{value: val, isSet: true}
}

func (v NullablePatchUserApiKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchUserApiKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


