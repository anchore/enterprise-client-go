/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

<<<<<<< HEAD
API version: 2.0.0
=======
API version: 0.1.0
>>>>>>> 48fc108 (feat: updated the enterprise ref)
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// AlertSummary A summary of the stateful indicator of a specific event in the system
type AlertSummary struct {
	// Identifier for the alert
	Uuid *string `json:"uuid,omitempty"`
	// Type of the alert
	Type *string `json:"type,omitempty"`
	// Current state of the alert
	State *string `json:"state,omitempty"`
	ResourceLabels *[]ResourceLabel `json:"resource_labels,omitempty"`
	// Account that closed the alert
	ClosedBy *string `json:"closed_by,omitempty"`
	// Reason for closing the alert
	ClosedReason *string `json:"closed_reason,omitempty"`
	// RFC 3339 formatted UTC timestamp when the alert was generated
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// RFC 3339 formatted UTC timestamp when the alert was last modified
	LastUpdated *time.Time `json:"last_updated,omitempty"`
}

// NewAlertSummary instantiates a new AlertSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertSummary() *AlertSummary {
	this := AlertSummary{}
	return &this
}

// NewAlertSummaryWithDefaults instantiates a new AlertSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertSummaryWithDefaults() *AlertSummary {
	this := AlertSummary{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *AlertSummary) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertSummary) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *AlertSummary) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *AlertSummary) SetUuid(v string) {
	o.Uuid = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AlertSummary) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertSummary) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AlertSummary) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AlertSummary) SetType(v string) {
	o.Type = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AlertSummary) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertSummary) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AlertSummary) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *AlertSummary) SetState(v string) {
	o.State = &v
}

// GetResourceLabels returns the ResourceLabels field value if set, zero value otherwise.
func (o *AlertSummary) GetResourceLabels() []ResourceLabel {
	if o == nil || o.ResourceLabels == nil {
		var ret []ResourceLabel
		return ret
	}
	return *o.ResourceLabels
}

// GetResourceLabelsOk returns a tuple with the ResourceLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertSummary) GetResourceLabelsOk() (*[]ResourceLabel, bool) {
	if o == nil || o.ResourceLabels == nil {
		return nil, false
	}
	return o.ResourceLabels, true
}

// HasResourceLabels returns a boolean if a field has been set.
func (o *AlertSummary) HasResourceLabels() bool {
	if o != nil && o.ResourceLabels != nil {
		return true
	}

	return false
}

// SetResourceLabels gets a reference to the given []ResourceLabel and assigns it to the ResourceLabels field.
func (o *AlertSummary) SetResourceLabels(v []ResourceLabel) {
	o.ResourceLabels = &v
}

// GetClosedBy returns the ClosedBy field value if set, zero value otherwise.
func (o *AlertSummary) GetClosedBy() string {
	if o == nil || o.ClosedBy == nil {
		var ret string
		return ret
	}
	return *o.ClosedBy
}

// GetClosedByOk returns a tuple with the ClosedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertSummary) GetClosedByOk() (*string, bool) {
	if o == nil || o.ClosedBy == nil {
		return nil, false
	}
	return o.ClosedBy, true
}

// HasClosedBy returns a boolean if a field has been set.
func (o *AlertSummary) HasClosedBy() bool {
	if o != nil && o.ClosedBy != nil {
		return true
	}

	return false
}

// SetClosedBy gets a reference to the given string and assigns it to the ClosedBy field.
func (o *AlertSummary) SetClosedBy(v string) {
	o.ClosedBy = &v
}

// GetClosedReason returns the ClosedReason field value if set, zero value otherwise.
func (o *AlertSummary) GetClosedReason() string {
	if o == nil || o.ClosedReason == nil {
		var ret string
		return ret
	}
	return *o.ClosedReason
}

// GetClosedReasonOk returns a tuple with the ClosedReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertSummary) GetClosedReasonOk() (*string, bool) {
	if o == nil || o.ClosedReason == nil {
		return nil, false
	}
	return o.ClosedReason, true
}

// HasClosedReason returns a boolean if a field has been set.
func (o *AlertSummary) HasClosedReason() bool {
	if o != nil && o.ClosedReason != nil {
		return true
	}

	return false
}

// SetClosedReason gets a reference to the given string and assigns it to the ClosedReason field.
func (o *AlertSummary) SetClosedReason(v string) {
	o.ClosedReason = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AlertSummary) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertSummary) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AlertSummary) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AlertSummary) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *AlertSummary) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertSummary) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *AlertSummary) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *AlertSummary) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o AlertSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.ResourceLabels != nil {
		toSerialize["resource_labels"] = o.ResourceLabels
	}
	if o.ClosedBy != nil {
		toSerialize["closed_by"] = o.ClosedBy
	}
	if o.ClosedReason != nil {
		toSerialize["closed_reason"] = o.ClosedReason
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	return json.Marshal(toSerialize)
}

type NullableAlertSummary struct {
	value *AlertSummary
	isSet bool
}

func (v NullableAlertSummary) Get() *AlertSummary {
	return v.value
}

func (v *NullableAlertSummary) Set(val *AlertSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertSummary(val *AlertSummary) *NullableAlertSummary {
	return &NullableAlertSummary{value: val, isSet: true}
}

func (v NullableAlertSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


