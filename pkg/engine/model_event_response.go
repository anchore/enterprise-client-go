/*
Anchore Engine API Server

This is the Anchore Engine API. Provides the primary external API for users of the service.

API version: 0.1.20
Contact: nurmi@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package engine

import (
	"encoding/json"
	"time"
)

// EventResponse A record of occurance of an asynchronous event triggered either by system or by user activity
type EventResponse struct {
	GeneratedUuid *string `json:"generated_uuid,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Event *EventResponseEvent `json:"event,omitempty"`
}

// NewEventResponse instantiates a new EventResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventResponse() *EventResponse {
	this := EventResponse{}
	return &this
}

// NewEventResponseWithDefaults instantiates a new EventResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventResponseWithDefaults() *EventResponse {
	this := EventResponse{}
	return &this
}

// GetGeneratedUuid returns the GeneratedUuid field value if set, zero value otherwise.
func (o *EventResponse) GetGeneratedUuid() string {
	if o == nil || o.GeneratedUuid == nil {
		var ret string
		return ret
	}
	return *o.GeneratedUuid
}

// GetGeneratedUuidOk returns a tuple with the GeneratedUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponse) GetGeneratedUuidOk() (*string, bool) {
	if o == nil || o.GeneratedUuid == nil {
		return nil, false
	}
	return o.GeneratedUuid, true
}

// HasGeneratedUuid returns a boolean if a field has been set.
func (o *EventResponse) HasGeneratedUuid() bool {
	if o != nil && o.GeneratedUuid != nil {
		return true
	}

	return false
}

// SetGeneratedUuid gets a reference to the given string and assigns it to the GeneratedUuid field.
func (o *EventResponse) SetGeneratedUuid(v string) {
	o.GeneratedUuid = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EventResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EventResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *EventResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *EventResponse) GetEvent() EventResponseEvent {
	if o == nil || o.Event == nil {
		var ret EventResponseEvent
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponse) GetEventOk() (*EventResponseEvent, bool) {
	if o == nil || o.Event == nil {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *EventResponse) HasEvent() bool {
	if o != nil && o.Event != nil {
		return true
	}

	return false
}

// SetEvent gets a reference to the given EventResponseEvent and assigns it to the Event field.
func (o *EventResponse) SetEvent(v EventResponseEvent) {
	o.Event = &v
}

func (o EventResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GeneratedUuid != nil {
		toSerialize["generated_uuid"] = o.GeneratedUuid
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Event != nil {
		toSerialize["event"] = o.Event
	}
	return json.Marshal(toSerialize)
}

type NullableEventResponse struct {
	value *EventResponse
	isSet bool
}

func (v NullableEventResponse) Get() *EventResponse {
	return v.value
}

func (v *NullableEventResponse) Set(val *EventResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEventResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEventResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventResponse(val *EventResponse) *NullableEventResponse {
	return &NullableEventResponse{value: val, isSet: true}
}

func (v NullableEventResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


