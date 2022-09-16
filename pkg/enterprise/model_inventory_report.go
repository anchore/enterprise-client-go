/*
Anchore Enterprise API Server

This is the Anchore Enterprise API. It provides additional external API routes and functionality for enterprise users.

API version: 0.3.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// InventoryReport Defines the object that Anchore expects to be provided for a given Image Inventory
type InventoryReport struct {
	ClusterName *string `json:"cluster_name,omitempty"`
	InventoryType *string `json:"inventory_type,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
	Results *[]InventoryReportItem `json:"results,omitempty"`
}

// NewInventoryReport instantiates a new InventoryReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryReport() *InventoryReport {
	this := InventoryReport{}
	return &this
}

// NewInventoryReportWithDefaults instantiates a new InventoryReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryReportWithDefaults() *InventoryReport {
	this := InventoryReport{}
	return &this
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *InventoryReport) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryReport) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *InventoryReport) HasClusterName() bool {
	if o != nil && o.ClusterName != nil {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *InventoryReport) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetInventoryType returns the InventoryType field value if set, zero value otherwise.
func (o *InventoryReport) GetInventoryType() string {
	if o == nil || o.InventoryType == nil {
		var ret string
		return ret
	}
	return *o.InventoryType
}

// GetInventoryTypeOk returns a tuple with the InventoryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryReport) GetInventoryTypeOk() (*string, bool) {
	if o == nil || o.InventoryType == nil {
		return nil, false
	}
	return o.InventoryType, true
}

// HasInventoryType returns a boolean if a field has been set.
func (o *InventoryReport) HasInventoryType() bool {
	if o != nil && o.InventoryType != nil {
		return true
	}

	return false
}

// SetInventoryType gets a reference to the given string and assigns it to the InventoryType field.
func (o *InventoryReport) SetInventoryType(v string) {
	o.InventoryType = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *InventoryReport) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryReport) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *InventoryReport) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *InventoryReport) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *InventoryReport) GetResults() []InventoryReportItem {
	if o == nil || o.Results == nil {
		var ret []InventoryReportItem
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryReport) GetResultsOk() (*[]InventoryReportItem, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *InventoryReport) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []InventoryReportItem and assigns it to the Results field.
func (o *InventoryReport) SetResults(v []InventoryReportItem) {
	o.Results = &v
}

func (o InventoryReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClusterName != nil {
		toSerialize["cluster_name"] = o.ClusterName
	}
	if o.InventoryType != nil {
		toSerialize["inventory_type"] = o.InventoryType
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryReport struct {
	value *InventoryReport
	isSet bool
}

func (v NullableInventoryReport) Get() *InventoryReport {
	return v.value
}

func (v *NullableInventoryReport) Set(val *InventoryReport) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryReport) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryReport(val *InventoryReport) *NullableInventoryReport {
	return &NullableInventoryReport{value: val, isSet: true}
}

func (v NullableInventoryReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


