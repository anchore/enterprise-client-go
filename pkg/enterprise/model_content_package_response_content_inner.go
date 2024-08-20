/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.6.2
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"fmt"
)

// ContentPackageResponseContentInner - struct for ContentPackageResponseContentInner
type ContentPackageResponseContentInner struct {
	JavaPackageContent *JavaPackageContent
	PackageContent *PackageContent
}

// JavaPackageContentAsContentPackageResponseContentInner is a convenience function that returns JavaPackageContent wrapped in ContentPackageResponseContentInner
func JavaPackageContentAsContentPackageResponseContentInner(v *JavaPackageContent) ContentPackageResponseContentInner {
	return ContentPackageResponseContentInner{
		JavaPackageContent: v,
	}
}

// PackageContentAsContentPackageResponseContentInner is a convenience function that returns PackageContent wrapped in ContentPackageResponseContentInner
func PackageContentAsContentPackageResponseContentInner(v *PackageContent) ContentPackageResponseContentInner {
	return ContentPackageResponseContentInner{
		PackageContent: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ContentPackageResponseContentInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into JavaPackageContent
	err = newStrictDecoder(data).Decode(&dst.JavaPackageContent)
	if err == nil {
		jsonJavaPackageContent, _ := json.Marshal(dst.JavaPackageContent)
		if string(jsonJavaPackageContent) == "{}" { // empty struct
			dst.JavaPackageContent = nil
		} else {
			match++
		}
	} else {
		dst.JavaPackageContent = nil
	}

	// try to unmarshal data into PackageContent
	err = newStrictDecoder(data).Decode(&dst.PackageContent)
	if err == nil {
		jsonPackageContent, _ := json.Marshal(dst.PackageContent)
		if string(jsonPackageContent) == "{}" { // empty struct
			dst.PackageContent = nil
		} else {
			match++
		}
	} else {
		dst.PackageContent = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.JavaPackageContent = nil
		dst.PackageContent = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(ContentPackageResponseContentInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(ContentPackageResponseContentInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ContentPackageResponseContentInner) MarshalJSON() ([]byte, error) {
	if src.JavaPackageContent != nil {
		return json.Marshal(&src.JavaPackageContent)
	}

	if src.PackageContent != nil {
		return json.Marshal(&src.PackageContent)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ContentPackageResponseContentInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.JavaPackageContent != nil {
		return obj.JavaPackageContent
	}

	if obj.PackageContent != nil {
		return obj.PackageContent
	}

	// all schemas are nil
	return nil
}

type NullableContentPackageResponseContentInner struct {
	value *ContentPackageResponseContentInner
	isSet bool
}

func (v NullableContentPackageResponseContentInner) Get() *ContentPackageResponseContentInner {
	return v.value
}

func (v *NullableContentPackageResponseContentInner) Set(val *ContentPackageResponseContentInner) {
	v.value = val
	v.isSet = true
}

func (v NullableContentPackageResponseContentInner) IsSet() bool {
	return v.isSet
}

func (v *NullableContentPackageResponseContentInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentPackageResponseContentInner(val *ContentPackageResponseContentInner) *NullableContentPackageResponseContentInner {
	return &NullableContentPackageResponseContentInner{value: val, isSet: true}
}

func (v NullableContentPackageResponseContentInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentPackageResponseContentInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


