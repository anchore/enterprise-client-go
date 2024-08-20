/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.6.2
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// EventResponseEvent struct for EventResponseEvent
type EventResponseEvent struct {
	Source *EventResponseEventSource `json:"source,omitempty"`
	Resource *EventResponseEventResource `json:"resource,omitempty"`
	Type *string `json:"type,omitempty"`
	Category *string `json:"category,omitempty"`
	Level *string `json:"level,omitempty"`
	Message *string `json:"message,omitempty"`
	Details interface{} `json:"details,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// NewEventResponseEvent instantiates a new EventResponseEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventResponseEvent() *EventResponseEvent {
	this := EventResponseEvent{}
	return &this
}

// NewEventResponseEventWithDefaults instantiates a new EventResponseEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventResponseEventWithDefaults() *EventResponseEvent {
	this := EventResponseEvent{}
	return &this
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *EventResponseEvent) GetSource() EventResponseEventSource {
	if o == nil || o.Source == nil {
		var ret EventResponseEventSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponseEvent) GetSourceOk() (*EventResponseEventSource, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *EventResponseEvent) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given EventResponseEventSource and assigns it to the Source field.
func (o *EventResponseEvent) SetSource(v EventResponseEventSource) {
	o.Source = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *EventResponseEvent) GetResource() EventResponseEventResource {
	if o == nil || o.Resource == nil {
		var ret EventResponseEventResource
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponseEvent) GetResourceOk() (*EventResponseEventResource, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *EventResponseEvent) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given EventResponseEventResource and assigns it to the Resource field.
func (o *EventResponseEvent) SetResource(v EventResponseEventResource) {
	o.Resource = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EventResponseEvent) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponseEvent) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EventResponseEvent) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EventResponseEvent) SetType(v string) {
	o.Type = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *EventResponseEvent) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponseEvent) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *EventResponseEvent) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *EventResponseEvent) SetCategory(v string) {
	o.Category = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *EventResponseEvent) GetLevel() string {
	if o == nil || o.Level == nil {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponseEvent) GetLevelOk() (*string, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *EventResponseEvent) HasLevel() bool {
	if o != nil && o.Level != nil {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *EventResponseEvent) SetLevel(v string) {
	o.Level = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *EventResponseEvent) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponseEvent) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *EventResponseEvent) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *EventResponseEvent) SetMessage(v string) {
	o.Message = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *EventResponseEvent) GetDetails() interface{} {
	if o == nil || o.Details == nil {
		var ret interface{}
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponseEvent) GetDetailsOk() (interface{}, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *EventResponseEvent) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given interface{} and assigns it to the Details field.
func (o *EventResponseEvent) SetDetails(v interface{}) {
	o.Details = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *EventResponseEvent) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponseEvent) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *EventResponseEvent) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *EventResponseEvent) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

func (o EventResponseEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Resource != nil {
		toSerialize["resource"] = o.Resource
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.Level != nil {
		toSerialize["level"] = o.Level
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	return json.Marshal(toSerialize)
}

type NullableEventResponseEvent struct {
	value *EventResponseEvent
	isSet bool
}

func (v NullableEventResponseEvent) Get() *EventResponseEvent {
	return v.value
}

func (v *NullableEventResponseEvent) Set(val *EventResponseEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableEventResponseEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableEventResponseEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventResponseEvent(val *EventResponseEvent) *NullableEventResponseEvent {
	return &NullableEventResponseEvent{value: val, isSet: true}
}

func (v NullableEventResponseEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventResponseEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


