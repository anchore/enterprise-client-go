/*
Anchore Enterprise API Server

This is the Anchore Enterprise API. It provides additional external API routes and functionality for enterprise users.

API version: 0.7.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// RelationshipSbomDiff The diff of two sboms with context applied in each difference. The \"added\" and \"removed\" directions depend on the relationship to which this diff applies. A relationship defines a source, a target, and a type. For example, a relationship of type \"contains\" with a source of an image and a target of a source revision will indicate that the diff is between the source repo sbom and the image sbom. Added packages are present in the image but not in the source, removed are present in the source revision but not in the image, etc. 
type RelationshipSbomDiff struct {
	// Packages added based on the type of relationship. A \"contains\" relationship means packages present in the source artifact (image) not present in the target (source repo) of the relationship.
	SourceOnly *[]Package `json:"source_only,omitempty"`
	// Packages removed based on the type of relationship. A \"contains\" relationship means packages not present in the source artifact (image) present in the target (source repo) of the relationship.
	TargetOnly *[]Package `json:"target_only,omitempty"`
	SourceModified *[]ModifiedPackage `json:"source_modified,omitempty"`
	BothUnmodified *[]Package `json:"both_unmodified,omitempty"`
}

// NewRelationshipSbomDiff instantiates a new RelationshipSbomDiff object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipSbomDiff() *RelationshipSbomDiff {
	this := RelationshipSbomDiff{}
	return &this
}

// NewRelationshipSbomDiffWithDefaults instantiates a new RelationshipSbomDiff object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipSbomDiffWithDefaults() *RelationshipSbomDiff {
	this := RelationshipSbomDiff{}
	return &this
}

// GetSourceOnly returns the SourceOnly field value if set, zero value otherwise.
func (o *RelationshipSbomDiff) GetSourceOnly() []Package {
	if o == nil || o.SourceOnly == nil {
		var ret []Package
		return ret
	}
	return *o.SourceOnly
}

// GetSourceOnlyOk returns a tuple with the SourceOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipSbomDiff) GetSourceOnlyOk() (*[]Package, bool) {
	if o == nil || o.SourceOnly == nil {
		return nil, false
	}
	return o.SourceOnly, true
}

// HasSourceOnly returns a boolean if a field has been set.
func (o *RelationshipSbomDiff) HasSourceOnly() bool {
	if o != nil && o.SourceOnly != nil {
		return true
	}

	return false
}

// SetSourceOnly gets a reference to the given []Package and assigns it to the SourceOnly field.
func (o *RelationshipSbomDiff) SetSourceOnly(v []Package) {
	o.SourceOnly = &v
}

// GetTargetOnly returns the TargetOnly field value if set, zero value otherwise.
func (o *RelationshipSbomDiff) GetTargetOnly() []Package {
	if o == nil || o.TargetOnly == nil {
		var ret []Package
		return ret
	}
	return *o.TargetOnly
}

// GetTargetOnlyOk returns a tuple with the TargetOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipSbomDiff) GetTargetOnlyOk() (*[]Package, bool) {
	if o == nil || o.TargetOnly == nil {
		return nil, false
	}
	return o.TargetOnly, true
}

// HasTargetOnly returns a boolean if a field has been set.
func (o *RelationshipSbomDiff) HasTargetOnly() bool {
	if o != nil && o.TargetOnly != nil {
		return true
	}

	return false
}

// SetTargetOnly gets a reference to the given []Package and assigns it to the TargetOnly field.
func (o *RelationshipSbomDiff) SetTargetOnly(v []Package) {
	o.TargetOnly = &v
}

// GetSourceModified returns the SourceModified field value if set, zero value otherwise.
func (o *RelationshipSbomDiff) GetSourceModified() []ModifiedPackage {
	if o == nil || o.SourceModified == nil {
		var ret []ModifiedPackage
		return ret
	}
	return *o.SourceModified
}

// GetSourceModifiedOk returns a tuple with the SourceModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipSbomDiff) GetSourceModifiedOk() (*[]ModifiedPackage, bool) {
	if o == nil || o.SourceModified == nil {
		return nil, false
	}
	return o.SourceModified, true
}

// HasSourceModified returns a boolean if a field has been set.
func (o *RelationshipSbomDiff) HasSourceModified() bool {
	if o != nil && o.SourceModified != nil {
		return true
	}

	return false
}

// SetSourceModified gets a reference to the given []ModifiedPackage and assigns it to the SourceModified field.
func (o *RelationshipSbomDiff) SetSourceModified(v []ModifiedPackage) {
	o.SourceModified = &v
}

// GetBothUnmodified returns the BothUnmodified field value if set, zero value otherwise.
func (o *RelationshipSbomDiff) GetBothUnmodified() []Package {
	if o == nil || o.BothUnmodified == nil {
		var ret []Package
		return ret
	}
	return *o.BothUnmodified
}

// GetBothUnmodifiedOk returns a tuple with the BothUnmodified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipSbomDiff) GetBothUnmodifiedOk() (*[]Package, bool) {
	if o == nil || o.BothUnmodified == nil {
		return nil, false
	}
	return o.BothUnmodified, true
}

// HasBothUnmodified returns a boolean if a field has been set.
func (o *RelationshipSbomDiff) HasBothUnmodified() bool {
	if o != nil && o.BothUnmodified != nil {
		return true
	}

	return false
}

// SetBothUnmodified gets a reference to the given []Package and assigns it to the BothUnmodified field.
func (o *RelationshipSbomDiff) SetBothUnmodified(v []Package) {
	o.BothUnmodified = &v
}

func (o RelationshipSbomDiff) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SourceOnly != nil {
		toSerialize["source_only"] = o.SourceOnly
	}
	if o.TargetOnly != nil {
		toSerialize["target_only"] = o.TargetOnly
	}
	if o.SourceModified != nil {
		toSerialize["source_modified"] = o.SourceModified
	}
	if o.BothUnmodified != nil {
		toSerialize["both_unmodified"] = o.BothUnmodified
	}
	return json.Marshal(toSerialize)
}

type NullableRelationshipSbomDiff struct {
	value *RelationshipSbomDiff
	isSet bool
}

func (v NullableRelationshipSbomDiff) Get() *RelationshipSbomDiff {
	return v.value
}

func (v *NullableRelationshipSbomDiff) Set(val *RelationshipSbomDiff) {
	v.value = val
	v.isSet = true
}

func (v NullableRelationshipSbomDiff) IsSet() bool {
	return v.isSet
}

func (v *NullableRelationshipSbomDiff) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelationshipSbomDiff(val *RelationshipSbomDiff) *NullableRelationshipSbomDiff {
	return &NullableRelationshipSbomDiff{value: val, isSet: true}
}

func (v NullableRelationshipSbomDiff) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelationshipSbomDiff) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


