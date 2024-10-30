/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.8.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"fmt"
)


// ImportPackageCpesInner struct for ImportPackageCpesInner
type ImportPackageCpesInner struct {
	ImportPackageCPE *ImportPackageCPE
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ImportPackageCpesInner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ImportPackageCPE
	err = json.Unmarshal(data, &dst.ImportPackageCPE);
	if err == nil {
		jsonImportPackageCPE, _ := json.Marshal(dst.ImportPackageCPE)
		if string(jsonImportPackageCPE) == "{}" { // empty struct
			dst.ImportPackageCPE = nil
		} else {
			return nil // data stored in dst.ImportPackageCPE, return on the first match
		}
	} else {
		dst.ImportPackageCPE = nil
	}

	// try to unmarshal JSON data into String
	err = json.Unmarshal(data, &dst.String);
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ImportPackageCpesInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ImportPackageCpesInner) MarshalJSON() ([]byte, error) {
	if src.ImportPackageCPE != nil {
		return json.Marshal(&src.ImportPackageCPE)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableImportPackageCpesInner struct {
	value *ImportPackageCpesInner
	isSet bool
}

func (v NullableImportPackageCpesInner) Get() *ImportPackageCpesInner {
	return v.value
}

func (v *NullableImportPackageCpesInner) Set(val *ImportPackageCpesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableImportPackageCpesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableImportPackageCpesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportPackageCpesInner(val *ImportPackageCpesInner) *NullableImportPackageCpesInner {
	return &NullableImportPackageCpesInner{value: val, isSet: true}
}

func (v NullableImportPackageCpesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportPackageCpesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


