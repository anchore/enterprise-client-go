/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.12.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the IntegrationReportedStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationReportedStatus{}

// IntegrationReportedStatus Status of the integration as perceived by the integration instance itself
type IntegrationReportedStatus struct {
	State *IntegrationHealthState `json:"state,omitempty"`
	Reason *string `json:"reason,omitempty"`
	Details interface{} `json:"details,omitempty"`
}

// NewIntegrationReportedStatus instantiates a new IntegrationReportedStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationReportedStatus() *IntegrationReportedStatus {
	this := IntegrationReportedStatus{}
	return &this
}

// NewIntegrationReportedStatusWithDefaults instantiates a new IntegrationReportedStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationReportedStatusWithDefaults() *IntegrationReportedStatus {
	this := IntegrationReportedStatus{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *IntegrationReportedStatus) GetState() IntegrationHealthState {
	if o == nil || IsNil(o.State) {
		var ret IntegrationHealthState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationReportedStatus) GetStateOk() (*IntegrationHealthState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *IntegrationReportedStatus) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given IntegrationHealthState and assigns it to the State field.
func (o *IntegrationReportedStatus) SetState(v IntegrationHealthState) {
	o.State = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *IntegrationReportedStatus) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationReportedStatus) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *IntegrationReportedStatus) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *IntegrationReportedStatus) SetReason(v string) {
	o.Reason = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *IntegrationReportedStatus) GetDetails() interface{} {
	if o == nil || IsNil(o.Details) {
		var ret interface{}
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationReportedStatus) GetDetailsOk() (interface{}, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *IntegrationReportedStatus) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given interface{} and assigns it to the Details field.
func (o *IntegrationReportedStatus) SetDetails(v interface{}) {
	o.Details = v
}

func (o IntegrationReportedStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationReportedStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableIntegrationReportedStatus struct {
	value *IntegrationReportedStatus
	isSet bool
}

func (v NullableIntegrationReportedStatus) Get() *IntegrationReportedStatus {
	return v.value
}

func (v *NullableIntegrationReportedStatus) Set(val *IntegrationReportedStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationReportedStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationReportedStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationReportedStatus(val *IntegrationReportedStatus) *NullableIntegrationReportedStatus {
	return &NullableIntegrationReportedStatus{value: val, isSet: true}
}

func (v NullableIntegrationReportedStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationReportedStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


