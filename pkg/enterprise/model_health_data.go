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
	"bytes"
	"fmt"
)

// checks if the HealthData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HealthData{}

// HealthData Integration instance specific health data
type HealthData struct {
	Type IntegrationType `json:"type"`
	// Version of the health data schema
	Version int32 `json:"version"`
	// List of errors
	Errors []string `json:"errors,omitempty"`
	// Information about sent inventory report for accounts
	AccountK8sInventoryReports *map[string]AccountK8sInventoryReportInfo `json:"account_k8s_inventory_reports,omitempty"`
}

type _HealthData HealthData

// NewHealthData instantiates a new HealthData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthData(type_ IntegrationType, version int32) *HealthData {
	this := HealthData{}
	this.Type = type_
	this.Version = version
	return &this
}

// NewHealthDataWithDefaults instantiates a new HealthData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthDataWithDefaults() *HealthData {
	this := HealthData{}
	return &this
}

// GetType returns the Type field value
func (o *HealthData) GetType() IntegrationType {
	if o == nil {
		var ret IntegrationType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HealthData) GetTypeOk() (*IntegrationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *HealthData) SetType(v IntegrationType) {
	o.Type = v
}

// GetVersion returns the Version field value
func (o *HealthData) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *HealthData) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *HealthData) SetVersion(v int32) {
	o.Version = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *HealthData) GetErrors() []string {
	if o == nil || IsNil(o.Errors) {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthData) GetErrorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *HealthData) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *HealthData) SetErrors(v []string) {
	o.Errors = v
}

// GetAccountK8sInventoryReports returns the AccountK8sInventoryReports field value if set, zero value otherwise.
func (o *HealthData) GetAccountK8sInventoryReports() map[string]AccountK8sInventoryReportInfo {
	if o == nil || IsNil(o.AccountK8sInventoryReports) {
		var ret map[string]AccountK8sInventoryReportInfo
		return ret
	}
	return *o.AccountK8sInventoryReports
}

// GetAccountK8sInventoryReportsOk returns a tuple with the AccountK8sInventoryReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthData) GetAccountK8sInventoryReportsOk() (*map[string]AccountK8sInventoryReportInfo, bool) {
	if o == nil || IsNil(o.AccountK8sInventoryReports) {
		return nil, false
	}
	return o.AccountK8sInventoryReports, true
}

// HasAccountK8sInventoryReports returns a boolean if a field has been set.
func (o *HealthData) HasAccountK8sInventoryReports() bool {
	if o != nil && !IsNil(o.AccountK8sInventoryReports) {
		return true
	}

	return false
}

// SetAccountK8sInventoryReports gets a reference to the given map[string]AccountK8sInventoryReportInfo and assigns it to the AccountK8sInventoryReports field.
func (o *HealthData) SetAccountK8sInventoryReports(v map[string]AccountK8sInventoryReportInfo) {
	o.AccountK8sInventoryReports = &v
}

func (o HealthData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HealthData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["version"] = o.Version
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.AccountK8sInventoryReports) {
		toSerialize["account_k8s_inventory_reports"] = o.AccountK8sInventoryReports
	}
	return toSerialize, nil
}

func (o *HealthData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"version",
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

	varHealthData := _HealthData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHealthData)

	if err != nil {
		return err
	}

	*o = HealthData(varHealthData)

	return err
}

type NullableHealthData struct {
	value *HealthData
	isSet bool
}

func (v NullableHealthData) Get() *HealthData {
	return v.value
}

func (v *NullableHealthData) Set(val *HealthData) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthData) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthData(val *HealthData) *NullableHealthData {
	return &NullableHealthData{value: val, isSet: true}
}

func (v NullableHealthData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


