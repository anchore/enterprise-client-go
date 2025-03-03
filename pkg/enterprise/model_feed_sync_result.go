/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.10.1
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the FeedSyncResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeedSyncResult{}

// FeedSyncResult The result of a sync of a single feed
type FeedSyncResult struct {
	// The name of the feed synced
	Feed *string `json:"feed,omitempty"`
	// The result of the sync operations, either co
	Status *string `json:"status,omitempty"`
	// The duration, in seconds, of the sync of the feed, the sum of all the group syncs
	TotalTimeSeconds *float32 `json:"total_time_seconds,omitempty"`
	// Array of group sync results
	Groups []GroupSyncResult `json:"groups,omitempty"`
}

// NewFeedSyncResult instantiates a new FeedSyncResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeedSyncResult() *FeedSyncResult {
	this := FeedSyncResult{}
	return &this
}

// NewFeedSyncResultWithDefaults instantiates a new FeedSyncResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeedSyncResultWithDefaults() *FeedSyncResult {
	this := FeedSyncResult{}
	return &this
}

// GetFeed returns the Feed field value if set, zero value otherwise.
func (o *FeedSyncResult) GetFeed() string {
	if o == nil || IsNil(o.Feed) {
		var ret string
		return ret
	}
	return *o.Feed
}

// GetFeedOk returns a tuple with the Feed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedSyncResult) GetFeedOk() (*string, bool) {
	if o == nil || IsNil(o.Feed) {
		return nil, false
	}
	return o.Feed, true
}

// HasFeed returns a boolean if a field has been set.
func (o *FeedSyncResult) HasFeed() bool {
	if o != nil && !IsNil(o.Feed) {
		return true
	}

	return false
}

// SetFeed gets a reference to the given string and assigns it to the Feed field.
func (o *FeedSyncResult) SetFeed(v string) {
	o.Feed = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FeedSyncResult) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedSyncResult) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FeedSyncResult) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *FeedSyncResult) SetStatus(v string) {
	o.Status = &v
}

// GetTotalTimeSeconds returns the TotalTimeSeconds field value if set, zero value otherwise.
func (o *FeedSyncResult) GetTotalTimeSeconds() float32 {
	if o == nil || IsNil(o.TotalTimeSeconds) {
		var ret float32
		return ret
	}
	return *o.TotalTimeSeconds
}

// GetTotalTimeSecondsOk returns a tuple with the TotalTimeSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedSyncResult) GetTotalTimeSecondsOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalTimeSeconds) {
		return nil, false
	}
	return o.TotalTimeSeconds, true
}

// HasTotalTimeSeconds returns a boolean if a field has been set.
func (o *FeedSyncResult) HasTotalTimeSeconds() bool {
	if o != nil && !IsNil(o.TotalTimeSeconds) {
		return true
	}

	return false
}

// SetTotalTimeSeconds gets a reference to the given float32 and assigns it to the TotalTimeSeconds field.
func (o *FeedSyncResult) SetTotalTimeSeconds(v float32) {
	o.TotalTimeSeconds = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *FeedSyncResult) GetGroups() []GroupSyncResult {
	if o == nil || IsNil(o.Groups) {
		var ret []GroupSyncResult
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedSyncResult) GetGroupsOk() ([]GroupSyncResult, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *FeedSyncResult) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []GroupSyncResult and assigns it to the Groups field.
func (o *FeedSyncResult) SetGroups(v []GroupSyncResult) {
	o.Groups = v
}

func (o FeedSyncResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeedSyncResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Feed) {
		toSerialize["feed"] = o.Feed
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TotalTimeSeconds) {
		toSerialize["total_time_seconds"] = o.TotalTimeSeconds
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	return toSerialize, nil
}

type NullableFeedSyncResult struct {
	value *FeedSyncResult
	isSet bool
}

func (v NullableFeedSyncResult) Get() *FeedSyncResult {
	return v.value
}

func (v *NullableFeedSyncResult) Set(val *FeedSyncResult) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedSyncResult) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedSyncResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedSyncResult(val *FeedSyncResult) *NullableFeedSyncResult {
	return &NullableFeedSyncResult{value: val, isSet: true}
}

func (v NullableFeedSyncResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedSyncResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


