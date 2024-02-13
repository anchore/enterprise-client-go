/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.2.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"fmt"
)

// ImportPackageCpesInner - struct for ImportPackageCpesInner
type ImportPackageCpesInner struct {
	ImportPackageCPE *ImportPackageCPE
	String *string
}

// ImportPackageCPEAsImportPackageCpesInner is a convenience function that returns ImportPackageCPE wrapped in ImportPackageCpesInner
func ImportPackageCPEAsImportPackageCpesInner(v *ImportPackageCPE) ImportPackageCpesInner {
	return ImportPackageCpesInner{
		ImportPackageCPE: v,
	}
}

// stringAsImportPackageCpesInner is a convenience function that returns string wrapped in ImportPackageCpesInner
func StringAsImportPackageCpesInner(v *string) ImportPackageCpesInner {
	return ImportPackageCpesInner{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ImportPackageCpesInner) UnmarshalJSON(data []byte) error {
	var err error
        match := 0
        // try to unmarshal data into ImportPackageCPE
        err = json.Unmarshal(data, &dst.ImportPackageCPE)
        if err == nil {
                jsonImportPackageCPE, _ := json.Marshal(dst.ImportPackageCPE)
                if string(jsonImportPackageCPE) == "{}" { // empty struct
                        dst.ImportPackageCPE = nil
                } else {
                        match++
                }
        } else {
                dst.ImportPackageCPE = nil
        }

        // try to unmarshal data into String
        err = json.Unmarshal(data, &dst.String)
        if err == nil {
                jsonstring, _ := json.Marshal(dst.String)
                if string(jsonstring) == "{}" { // empty struct
                        dst.String = nil
                } else {
                        match++
                }
        } else {
                dst.String = nil
        }

        if match > 1 { // more than 1 match
                // reset to nil
                dst.ImportPackageCPE = nil
                dst.String = nil

                return fmt.Errorf("Data matches more than one schema in oneOf(ImportPackageCpesInner)")
        } else if match == 1 {
                return nil // exactly one match
        } else { // no match
                return fmt.Errorf("Data failed to match schemas in oneOf(ImportPackageCpesInner)")
        }
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ImportPackageCpesInner) MarshalJSON() ([]byte, error) {
	if src.ImportPackageCPE != nil {
		return json.Marshal(&src.ImportPackageCPE)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ImportPackageCpesInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ImportPackageCPE != nil {
		return obj.ImportPackageCPE
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
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

