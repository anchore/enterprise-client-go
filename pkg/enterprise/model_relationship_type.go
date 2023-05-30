/*
Anchore Enterprise API Server

This is the Anchore Enterprise API. It provides additional external API routes and functionality for enterprise users.

API version: 0.8.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"fmt"
)

// RelationshipType The type of relationship between to artifacts
type RelationshipType string

// List of RelationshipType
const (
	CONTAINS RelationshipType = "contains"
	CONTAINED_BY RelationshipType = "contained_by"
)

var allowedRelationshipTypeEnumValues = []RelationshipType{
	"contains",
	"contained_by",
}

func (v *RelationshipType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RelationshipType(value)
	for _, existing := range allowedRelationshipTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RelationshipType", value)
}

// NewRelationshipTypeFromValue returns a pointer to a valid RelationshipType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRelationshipTypeFromValue(v string) (*RelationshipType, error) {
	ev := RelationshipType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RelationshipType: valid values are %v", v, allowedRelationshipTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RelationshipType) IsValid() bool {
	for _, existing := range allowedRelationshipTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RelationshipType value
func (v RelationshipType) Ptr() *RelationshipType {
	return &v
}

type NullableRelationshipType struct {
	value *RelationshipType
	isSet bool
}

func (v NullableRelationshipType) Get() *RelationshipType {
	return v.value
}

func (v *NullableRelationshipType) Set(val *RelationshipType) {
	v.value = val
	v.isSet = true
}

func (v NullableRelationshipType) IsSet() bool {
	return v.isSet
}

func (v *NullableRelationshipType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelationshipType(val *RelationshipType) *NullableRelationshipType {
	return &NullableRelationshipType{value: val, isSet: true}
}

func (v NullableRelationshipType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelationshipType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

