/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.7.2
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// checks if the EventResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventResponse{}

// EventResponse A record of occurrence of an asynchronous event triggered either by system or by user activity
type EventResponse struct {
	EventId *string `json:"event_id,omitempty"`
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

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *EventResponse) GetEventId() string {
	if o == nil || IsNil(o.EventId) {
		var ret string
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponse) GetEventIdOk() (*string, bool) {
	if o == nil || IsNil(o.EventId) {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *EventResponse) HasEventId() bool {
	if o != nil && !IsNil(o.EventId) {
		return true
	}

	return false
}

// SetEventId gets a reference to the given string and assigns it to the EventId field.
func (o *EventResponse) SetEventId(v string) {
	o.EventId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EventResponse) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EventResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
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
	if o == nil || IsNil(o.Event) {
		var ret EventResponseEvent
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponse) GetEventOk() (*EventResponseEvent, bool) {
	if o == nil || IsNil(o.Event) {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *EventResponse) HasEvent() bool {
	if o != nil && !IsNil(o.Event) {
		return true
	}

	return false
}

// SetEvent gets a reference to the given EventResponseEvent and assigns it to the Event field.
func (o *EventResponse) SetEvent(v EventResponseEvent) {
	o.Event = &v
}

func (o EventResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventId) {
		toSerialize["event_id"] = o.EventId
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Event) {
		toSerialize["event"] = o.Event
	}
	return toSerialize, nil
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


