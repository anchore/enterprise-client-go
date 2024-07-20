/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.6.1
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// EventResponseEventResource struct for EventResponseEventResource
type EventResponseEventResource struct {
	AccountName *string `json:"account_name,omitempty"`
	Id *string `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewEventResponseEventResource instantiates a new EventResponseEventResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventResponseEventResource() *EventResponseEventResource {
	this := EventResponseEventResource{}
	return &this
}

// NewEventResponseEventResourceWithDefaults instantiates a new EventResponseEventResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventResponseEventResourceWithDefaults() *EventResponseEventResource {
	this := EventResponseEventResource{}
	return &this
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *EventResponseEventResource) GetAccountName() string {
	if o == nil || o.AccountName == nil {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponseEventResource) GetAccountNameOk() (*string, bool) {
	if o == nil || o.AccountName == nil {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *EventResponseEventResource) HasAccountName() bool {
	if o != nil && o.AccountName != nil {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *EventResponseEventResource) SetAccountName(v string) {
	o.AccountName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EventResponseEventResource) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponseEventResource) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EventResponseEventResource) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EventResponseEventResource) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EventResponseEventResource) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponseEventResource) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EventResponseEventResource) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EventResponseEventResource) SetType(v string) {
	o.Type = &v
}

func (o EventResponseEventResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountName != nil {
		toSerialize["account_name"] = o.AccountName
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableEventResponseEventResource struct {
	value *EventResponseEventResource
	isSet bool
}

func (v NullableEventResponseEventResource) Get() *EventResponseEventResource {
	return v.value
}

func (v *NullableEventResponseEventResource) Set(val *EventResponseEventResource) {
	v.value = val
	v.isSet = true
}

func (v NullableEventResponseEventResource) IsSet() bool {
	return v.isSet
}

func (v *NullableEventResponseEventResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventResponseEventResource(val *EventResponseEventResource) *NullableEventResponseEventResource {
	return &NullableEventResponseEventResource{value: val, isSet: true}
}

func (v NullableEventResponseEventResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventResponseEventResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


