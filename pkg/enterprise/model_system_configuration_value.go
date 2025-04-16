/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.11.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// SystemConfigurationValue - struct for SystemConfigurationValue
type SystemConfigurationValue struct {
	ArrayOfSystemConfigurationValueOneOfInner *[]SystemConfigurationValueOneOfInner
	Bool *bool
	Float32 *float32
	String *string
}

// []SystemConfigurationValueOneOfInnerAsSystemConfigurationValue is a convenience function that returns []SystemConfigurationValueOneOfInner wrapped in SystemConfigurationValue
func ArrayOfSystemConfigurationValueOneOfInnerAsSystemConfigurationValue(v *[]SystemConfigurationValueOneOfInner) SystemConfigurationValue {
	return SystemConfigurationValue{
		ArrayOfSystemConfigurationValueOneOfInner: v,
	}
}

// boolAsSystemConfigurationValue is a convenience function that returns bool wrapped in SystemConfigurationValue
func BoolAsSystemConfigurationValue(v *bool) SystemConfigurationValue {
	return SystemConfigurationValue{
		Bool: v,
	}
}

// float32AsSystemConfigurationValue is a convenience function that returns float32 wrapped in SystemConfigurationValue
func Float32AsSystemConfigurationValue(v *float32) SystemConfigurationValue {
	return SystemConfigurationValue{
		Float32: v,
	}
}

// stringAsSystemConfigurationValue is a convenience function that returns string wrapped in SystemConfigurationValue
func StringAsSystemConfigurationValue(v *string) SystemConfigurationValue {
	return SystemConfigurationValue{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SystemConfigurationValue) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into ArrayOfSystemConfigurationValueOneOfInner
	err = newStrictDecoder(data).Decode(&dst.ArrayOfSystemConfigurationValueOneOfInner)
	if err == nil {
		jsonArrayOfSystemConfigurationValueOneOfInner, _ := json.Marshal(dst.ArrayOfSystemConfigurationValueOneOfInner)
		if string(jsonArrayOfSystemConfigurationValueOneOfInner) == "{}" { // empty struct
			dst.ArrayOfSystemConfigurationValueOneOfInner = nil
		} else {
			if err = validator.Validate(dst.ArrayOfSystemConfigurationValueOneOfInner); err != nil {
				dst.ArrayOfSystemConfigurationValueOneOfInner = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfSystemConfigurationValueOneOfInner = nil
	}

	// try to unmarshal data into Bool
	err = newStrictDecoder(data).Decode(&dst.Bool)
	if err == nil {
		jsonBool, _ := json.Marshal(dst.Bool)
		if string(jsonBool) == "{}" { // empty struct
			dst.Bool = nil
		} else {
			if err = validator.Validate(dst.Bool); err != nil {
				dst.Bool = nil
			} else {
				match++
			}
		}
	} else {
		dst.Bool = nil
	}

	// try to unmarshal data into Float32
	err = newStrictDecoder(data).Decode(&dst.Float32)
	if err == nil {
		jsonFloat32, _ := json.Marshal(dst.Float32)
		if string(jsonFloat32) == "{}" { // empty struct
			dst.Float32 = nil
		} else {
			if err = validator.Validate(dst.Float32); err != nil {
				dst.Float32 = nil
			} else {
				match++
			}
		}
	} else {
		dst.Float32 = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			if err = validator.Validate(dst.String); err != nil {
				dst.String = nil
			} else {
				match++
			}
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ArrayOfSystemConfigurationValueOneOfInner = nil
		dst.Bool = nil
		dst.Float32 = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SystemConfigurationValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SystemConfigurationValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SystemConfigurationValue) MarshalJSON() ([]byte, error) {
	if src.ArrayOfSystemConfigurationValueOneOfInner != nil {
		return json.Marshal(&src.ArrayOfSystemConfigurationValueOneOfInner)
	}

	if src.Bool != nil {
		return json.Marshal(&src.Bool)
	}

	if src.Float32 != nil {
		return json.Marshal(&src.Float32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SystemConfigurationValue) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfSystemConfigurationValueOneOfInner != nil {
		return obj.ArrayOfSystemConfigurationValueOneOfInner
	}

	if obj.Bool != nil {
		return obj.Bool
	}

	if obj.Float32 != nil {
		return obj.Float32
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableSystemConfigurationValue struct {
	value *SystemConfigurationValue
	isSet bool
}

func (v NullableSystemConfigurationValue) Get() *SystemConfigurationValue {
	return v.value
}

func (v *NullableSystemConfigurationValue) Set(val *SystemConfigurationValue) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemConfigurationValue) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemConfigurationValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemConfigurationValue(val *SystemConfigurationValue) *NullableSystemConfigurationValue {
	return &NullableSystemConfigurationValue{value: val, isSet: true}
}

func (v NullableSystemConfigurationValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemConfigurationValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


