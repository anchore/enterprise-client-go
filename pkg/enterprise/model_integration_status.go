/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.10.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// checks if the IntegrationStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationStatus{}

// IntegrationStatus Status of integration instance as perceived by Enterprise
type IntegrationStatus struct {
	State *IntegrationLifecycleState `json:"state,omitempty"`
	Reason *string `json:"reason,omitempty"`
	Details interface{} `json:"details,omitempty"`
	// Timestamp when status was updated
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewIntegrationStatus instantiates a new IntegrationStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationStatus() *IntegrationStatus {
	this := IntegrationStatus{}
	return &this
}

// NewIntegrationStatusWithDefaults instantiates a new IntegrationStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationStatusWithDefaults() *IntegrationStatus {
	this := IntegrationStatus{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *IntegrationStatus) GetState() IntegrationLifecycleState {
	if o == nil || IsNil(o.State) {
		var ret IntegrationLifecycleState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationStatus) GetStateOk() (*IntegrationLifecycleState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *IntegrationStatus) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given IntegrationLifecycleState and assigns it to the State field.
func (o *IntegrationStatus) SetState(v IntegrationLifecycleState) {
	o.State = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *IntegrationStatus) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationStatus) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *IntegrationStatus) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *IntegrationStatus) SetReason(v string) {
	o.Reason = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *IntegrationStatus) GetDetails() interface{} {
	if o == nil || IsNil(o.Details) {
		var ret interface{}
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationStatus) GetDetailsOk() (interface{}, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *IntegrationStatus) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given interface{} and assigns it to the Details field.
func (o *IntegrationStatus) SetDetails(v interface{}) {
	o.Details = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *IntegrationStatus) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationStatus) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *IntegrationStatus) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *IntegrationStatus) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o IntegrationStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationStatus) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableIntegrationStatus struct {
	value *IntegrationStatus
	isSet bool
}

func (v NullableIntegrationStatus) Get() *IntegrationStatus {
	return v.value
}

func (v *NullableIntegrationStatus) Set(val *IntegrationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationStatus(val *IntegrationStatus) *NullableIntegrationStatus {
	return &NullableIntegrationStatus{value: val, isSet: true}
}

func (v NullableIntegrationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


