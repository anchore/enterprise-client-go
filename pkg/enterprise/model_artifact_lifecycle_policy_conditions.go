/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.9.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ArtifactLifecyclePolicyConditions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArtifactLifecyclePolicyConditions{}

// ArtifactLifecyclePolicyConditions struct for ArtifactLifecyclePolicyConditions
type ArtifactLifecyclePolicyConditions struct {
	// The version of these policy conditions.
	Version *int32 `json:"version,omitempty"`
	// Select an image even if it exists in the runtime inventory, otherwise if false do not action anything that exists in runtime inventory
	EvenIfExistsInRuntimeInventory bool `json:"even_if_exists_in_runtime_inventory"`
	// An image analysis must be this many days old before it will be considered for processing. An integer value less than or equal to zero will cause this field to be ignored.
	DaysSinceAnalyzed int32 `json:"days_since_analyzed"`
	// Include base images in the policy selection criteria.
	IncludeBaseImages *bool `json:"include_base_images,omitempty"`
	// The type of artifact that will be processed.
	ArtifactType string `json:"artifact_type"`
}

type _ArtifactLifecyclePolicyConditions ArtifactLifecyclePolicyConditions

// NewArtifactLifecyclePolicyConditions instantiates a new ArtifactLifecyclePolicyConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtifactLifecyclePolicyConditions(evenIfExistsInRuntimeInventory bool, daysSinceAnalyzed int32, artifactType string) *ArtifactLifecyclePolicyConditions {
	this := ArtifactLifecyclePolicyConditions{}
	this.EvenIfExistsInRuntimeInventory = evenIfExistsInRuntimeInventory
	this.DaysSinceAnalyzed = daysSinceAnalyzed
	this.ArtifactType = artifactType
	return &this
}

// NewArtifactLifecyclePolicyConditionsWithDefaults instantiates a new ArtifactLifecyclePolicyConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtifactLifecyclePolicyConditionsWithDefaults() *ArtifactLifecyclePolicyConditions {
	this := ArtifactLifecyclePolicyConditions{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ArtifactLifecyclePolicyConditions) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactLifecyclePolicyConditions) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ArtifactLifecyclePolicyConditions) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *ArtifactLifecyclePolicyConditions) SetVersion(v int32) {
	o.Version = &v
}

// GetEvenIfExistsInRuntimeInventory returns the EvenIfExistsInRuntimeInventory field value
func (o *ArtifactLifecyclePolicyConditions) GetEvenIfExistsInRuntimeInventory() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EvenIfExistsInRuntimeInventory
}

// GetEvenIfExistsInRuntimeInventoryOk returns a tuple with the EvenIfExistsInRuntimeInventory field value
// and a boolean to check if the value has been set.
func (o *ArtifactLifecyclePolicyConditions) GetEvenIfExistsInRuntimeInventoryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvenIfExistsInRuntimeInventory, true
}

// SetEvenIfExistsInRuntimeInventory sets field value
func (o *ArtifactLifecyclePolicyConditions) SetEvenIfExistsInRuntimeInventory(v bool) {
	o.EvenIfExistsInRuntimeInventory = v
}

// GetDaysSinceAnalyzed returns the DaysSinceAnalyzed field value
func (o *ArtifactLifecyclePolicyConditions) GetDaysSinceAnalyzed() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DaysSinceAnalyzed
}

// GetDaysSinceAnalyzedOk returns a tuple with the DaysSinceAnalyzed field value
// and a boolean to check if the value has been set.
func (o *ArtifactLifecyclePolicyConditions) GetDaysSinceAnalyzedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DaysSinceAnalyzed, true
}

// SetDaysSinceAnalyzed sets field value
func (o *ArtifactLifecyclePolicyConditions) SetDaysSinceAnalyzed(v int32) {
	o.DaysSinceAnalyzed = v
}

// GetIncludeBaseImages returns the IncludeBaseImages field value if set, zero value otherwise.
func (o *ArtifactLifecyclePolicyConditions) GetIncludeBaseImages() bool {
	if o == nil || IsNil(o.IncludeBaseImages) {
		var ret bool
		return ret
	}
	return *o.IncludeBaseImages
}

// GetIncludeBaseImagesOk returns a tuple with the IncludeBaseImages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactLifecyclePolicyConditions) GetIncludeBaseImagesOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeBaseImages) {
		return nil, false
	}
	return o.IncludeBaseImages, true
}

// HasIncludeBaseImages returns a boolean if a field has been set.
func (o *ArtifactLifecyclePolicyConditions) HasIncludeBaseImages() bool {
	if o != nil && !IsNil(o.IncludeBaseImages) {
		return true
	}

	return false
}

// SetIncludeBaseImages gets a reference to the given bool and assigns it to the IncludeBaseImages field.
func (o *ArtifactLifecyclePolicyConditions) SetIncludeBaseImages(v bool) {
	o.IncludeBaseImages = &v
}

// GetArtifactType returns the ArtifactType field value
func (o *ArtifactLifecyclePolicyConditions) GetArtifactType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ArtifactType
}

// GetArtifactTypeOk returns a tuple with the ArtifactType field value
// and a boolean to check if the value has been set.
func (o *ArtifactLifecyclePolicyConditions) GetArtifactTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArtifactType, true
}

// SetArtifactType sets field value
func (o *ArtifactLifecyclePolicyConditions) SetArtifactType(v string) {
	o.ArtifactType = v
}

func (o ArtifactLifecyclePolicyConditions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArtifactLifecyclePolicyConditions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	toSerialize["even_if_exists_in_runtime_inventory"] = o.EvenIfExistsInRuntimeInventory
	toSerialize["days_since_analyzed"] = o.DaysSinceAnalyzed
	if !IsNil(o.IncludeBaseImages) {
		toSerialize["include_base_images"] = o.IncludeBaseImages
	}
	toSerialize["artifact_type"] = o.ArtifactType
	return toSerialize, nil
}

func (o *ArtifactLifecyclePolicyConditions) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"even_if_exists_in_runtime_inventory",
		"days_since_analyzed",
		"artifact_type",
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

	varArtifactLifecyclePolicyConditions := _ArtifactLifecyclePolicyConditions{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varArtifactLifecyclePolicyConditions)

	if err != nil {
		return err
	}

	*o = ArtifactLifecyclePolicyConditions(varArtifactLifecyclePolicyConditions)

	return err
}

type NullableArtifactLifecyclePolicyConditions struct {
	value *ArtifactLifecyclePolicyConditions
	isSet bool
}

func (v NullableArtifactLifecyclePolicyConditions) Get() *ArtifactLifecyclePolicyConditions {
	return v.value
}

func (v *NullableArtifactLifecyclePolicyConditions) Set(val *ArtifactLifecyclePolicyConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifactLifecyclePolicyConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifactLifecyclePolicyConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifactLifecyclePolicyConditions(val *ArtifactLifecyclePolicyConditions) *NullableArtifactLifecyclePolicyConditions {
	return &NullableArtifactLifecyclePolicyConditions{value: val, isSet: true}
}

func (v NullableArtifactLifecyclePolicyConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtifactLifecyclePolicyConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


