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
)

// checks if the EventsList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventsList{}

// EventsList Response envelope for paginated listing of events
type EventsList struct {
	// List of events
	Results []EventResponse `json:"results,omitempty"`
	// Boolean flag, True indicates there are more events and False otherwise
	NextPage *bool `json:"next_page,omitempty"`
	// Number of events in this page
	ItemCount *int32 `json:"item_count,omitempty"`
	// Page number of this result set
	Page *int32 `json:"page,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EventsList EventsList

// NewEventsList instantiates a new EventsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventsList() *EventsList {
	this := EventsList{}
	return &this
}

// NewEventsListWithDefaults instantiates a new EventsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventsListWithDefaults() *EventsList {
	this := EventsList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *EventsList) GetResults() []EventResponse {
	if o == nil || IsNil(o.Results) {
		var ret []EventResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsList) GetResultsOk() ([]EventResponse, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *EventsList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []EventResponse and assigns it to the Results field.
func (o *EventsList) SetResults(v []EventResponse) {
	o.Results = v
}

// GetNextPage returns the NextPage field value if set, zero value otherwise.
func (o *EventsList) GetNextPage() bool {
	if o == nil || IsNil(o.NextPage) {
		var ret bool
		return ret
	}
	return *o.NextPage
}

// GetNextPageOk returns a tuple with the NextPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsList) GetNextPageOk() (*bool, bool) {
	if o == nil || IsNil(o.NextPage) {
		return nil, false
	}
	return o.NextPage, true
}

// HasNextPage returns a boolean if a field has been set.
func (o *EventsList) HasNextPage() bool {
	if o != nil && !IsNil(o.NextPage) {
		return true
	}

	return false
}

// SetNextPage gets a reference to the given bool and assigns it to the NextPage field.
func (o *EventsList) SetNextPage(v bool) {
	o.NextPage = &v
}

// GetItemCount returns the ItemCount field value if set, zero value otherwise.
func (o *EventsList) GetItemCount() int32 {
	if o == nil || IsNil(o.ItemCount) {
		var ret int32
		return ret
	}
	return *o.ItemCount
}

// GetItemCountOk returns a tuple with the ItemCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsList) GetItemCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ItemCount) {
		return nil, false
	}
	return o.ItemCount, true
}

// HasItemCount returns a boolean if a field has been set.
func (o *EventsList) HasItemCount() bool {
	if o != nil && !IsNil(o.ItemCount) {
		return true
	}

	return false
}

// SetItemCount gets a reference to the given int32 and assigns it to the ItemCount field.
func (o *EventsList) SetItemCount(v int32) {
	o.ItemCount = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *EventsList) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsList) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *EventsList) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *EventsList) SetPage(v int32) {
	o.Page = &v
}

func (o EventsList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventsList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.NextPage) {
		toSerialize["next_page"] = o.NextPage
	}
	if !IsNil(o.ItemCount) {
		toSerialize["item_count"] = o.ItemCount
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EventsList) UnmarshalJSON(data []byte) (err error) {
	varEventsList := _EventsList{}

	err = json.Unmarshal(data, &varEventsList)

	if err != nil {
		return err
	}

	*o = EventsList(varEventsList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		delete(additionalProperties, "next_page")
		delete(additionalProperties, "item_count")
		delete(additionalProperties, "page")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEventsList struct {
	value *EventsList
	isSet bool
}

func (v NullableEventsList) Get() *EventsList {
	return v.value
}

func (v *NullableEventsList) Set(val *EventsList) {
	v.value = val
	v.isSet = true
}

func (v NullableEventsList) IsSet() bool {
	return v.isSet
}

func (v *NullableEventsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventsList(val *EventsList) *NullableEventsList {
	return &NullableEventsList{value: val, isSet: true}
}

func (v NullableEventsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


