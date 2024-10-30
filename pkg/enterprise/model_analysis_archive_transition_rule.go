/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.8.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the AnalysisArchiveTransitionRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnalysisArchiveTransitionRule{}

// AnalysisArchiveTransitionRule A rule for auto-archiving image analysis by time and/or tag-history
type AnalysisArchiveTransitionRule struct {
	Selector *ImageSelector `json:"selector,omitempty"`
	// Unique identifier for archive rule
	RuleId *string `json:"rule_id,omitempty"`
	// Number of images mapped to the tag that are newer
	TagVersionsNewer *int32 `json:"tag_versions_newer,omitempty"`
	// Matches if the analysis is strictly older than this number of days
	AnalysisAgeDays *int32 `json:"analysis_age_days,omitempty"`
	// The type of transition to make. If \"archive\", then archive an image from the working set and remove it from the working set. If \"delete\", then match against archived images and delete from the archive if match.
	Transition string `json:"transition"`
	// True if the rule applies to all accounts in the system. This is only available to admin users to update/modify, but all users with permission to list rules can see them
	SystemGlobal *bool `json:"system_global,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
	Exclude *AnalysisArchiveTransitionRuleExclude `json:"exclude,omitempty"`
	// This is the maximum number of image analyses an account can have. Can only be set on system_global rules
	MaxImagesPerAccount *int32 `json:"max_images_per_account,omitempty"`
}

type _AnalysisArchiveTransitionRule AnalysisArchiveTransitionRule

// NewAnalysisArchiveTransitionRule instantiates a new AnalysisArchiveTransitionRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalysisArchiveTransitionRule(transition string) *AnalysisArchiveTransitionRule {
	this := AnalysisArchiveTransitionRule{}
	this.Transition = transition
	return &this
}

// NewAnalysisArchiveTransitionRuleWithDefaults instantiates a new AnalysisArchiveTransitionRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalysisArchiveTransitionRuleWithDefaults() *AnalysisArchiveTransitionRule {
	this := AnalysisArchiveTransitionRule{}
	return &this
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *AnalysisArchiveTransitionRule) GetSelector() ImageSelector {
	if o == nil || IsNil(o.Selector) {
		var ret ImageSelector
		return ret
	}
	return *o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisArchiveTransitionRule) GetSelectorOk() (*ImageSelector, bool) {
	if o == nil || IsNil(o.Selector) {
		return nil, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *AnalysisArchiveTransitionRule) HasSelector() bool {
	if o != nil && !IsNil(o.Selector) {
		return true
	}

	return false
}

// SetSelector gets a reference to the given ImageSelector and assigns it to the Selector field.
func (o *AnalysisArchiveTransitionRule) SetSelector(v ImageSelector) {
	o.Selector = &v
}

// GetRuleId returns the RuleId field value if set, zero value otherwise.
func (o *AnalysisArchiveTransitionRule) GetRuleId() string {
	if o == nil || IsNil(o.RuleId) {
		var ret string
		return ret
	}
	return *o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisArchiveTransitionRule) GetRuleIdOk() (*string, bool) {
	if o == nil || IsNil(o.RuleId) {
		return nil, false
	}
	return o.RuleId, true
}

// HasRuleId returns a boolean if a field has been set.
func (o *AnalysisArchiveTransitionRule) HasRuleId() bool {
	if o != nil && !IsNil(o.RuleId) {
		return true
	}

	return false
}

// SetRuleId gets a reference to the given string and assigns it to the RuleId field.
func (o *AnalysisArchiveTransitionRule) SetRuleId(v string) {
	o.RuleId = &v
}

// GetTagVersionsNewer returns the TagVersionsNewer field value if set, zero value otherwise.
func (o *AnalysisArchiveTransitionRule) GetTagVersionsNewer() int32 {
	if o == nil || IsNil(o.TagVersionsNewer) {
		var ret int32
		return ret
	}
	return *o.TagVersionsNewer
}

// GetTagVersionsNewerOk returns a tuple with the TagVersionsNewer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisArchiveTransitionRule) GetTagVersionsNewerOk() (*int32, bool) {
	if o == nil || IsNil(o.TagVersionsNewer) {
		return nil, false
	}
	return o.TagVersionsNewer, true
}

// HasTagVersionsNewer returns a boolean if a field has been set.
func (o *AnalysisArchiveTransitionRule) HasTagVersionsNewer() bool {
	if o != nil && !IsNil(o.TagVersionsNewer) {
		return true
	}

	return false
}

// SetTagVersionsNewer gets a reference to the given int32 and assigns it to the TagVersionsNewer field.
func (o *AnalysisArchiveTransitionRule) SetTagVersionsNewer(v int32) {
	o.TagVersionsNewer = &v
}

// GetAnalysisAgeDays returns the AnalysisAgeDays field value if set, zero value otherwise.
func (o *AnalysisArchiveTransitionRule) GetAnalysisAgeDays() int32 {
	if o == nil || IsNil(o.AnalysisAgeDays) {
		var ret int32
		return ret
	}
	return *o.AnalysisAgeDays
}

// GetAnalysisAgeDaysOk returns a tuple with the AnalysisAgeDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisArchiveTransitionRule) GetAnalysisAgeDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.AnalysisAgeDays) {
		return nil, false
	}
	return o.AnalysisAgeDays, true
}

// HasAnalysisAgeDays returns a boolean if a field has been set.
func (o *AnalysisArchiveTransitionRule) HasAnalysisAgeDays() bool {
	if o != nil && !IsNil(o.AnalysisAgeDays) {
		return true
	}

	return false
}

// SetAnalysisAgeDays gets a reference to the given int32 and assigns it to the AnalysisAgeDays field.
func (o *AnalysisArchiveTransitionRule) SetAnalysisAgeDays(v int32) {
	o.AnalysisAgeDays = &v
}

// GetTransition returns the Transition field value
func (o *AnalysisArchiveTransitionRule) GetTransition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Transition
}

// GetTransitionOk returns a tuple with the Transition field value
// and a boolean to check if the value has been set.
func (o *AnalysisArchiveTransitionRule) GetTransitionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transition, true
}

// SetTransition sets field value
func (o *AnalysisArchiveTransitionRule) SetTransition(v string) {
	o.Transition = v
}

// GetSystemGlobal returns the SystemGlobal field value if set, zero value otherwise.
func (o *AnalysisArchiveTransitionRule) GetSystemGlobal() bool {
	if o == nil || IsNil(o.SystemGlobal) {
		var ret bool
		return ret
	}
	return *o.SystemGlobal
}

// GetSystemGlobalOk returns a tuple with the SystemGlobal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisArchiveTransitionRule) GetSystemGlobalOk() (*bool, bool) {
	if o == nil || IsNil(o.SystemGlobal) {
		return nil, false
	}
	return o.SystemGlobal, true
}

// HasSystemGlobal returns a boolean if a field has been set.
func (o *AnalysisArchiveTransitionRule) HasSystemGlobal() bool {
	if o != nil && !IsNil(o.SystemGlobal) {
		return true
	}

	return false
}

// SetSystemGlobal gets a reference to the given bool and assigns it to the SystemGlobal field.
func (o *AnalysisArchiveTransitionRule) SetSystemGlobal(v bool) {
	o.SystemGlobal = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AnalysisArchiveTransitionRule) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisArchiveTransitionRule) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AnalysisArchiveTransitionRule) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AnalysisArchiveTransitionRule) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *AnalysisArchiveTransitionRule) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisArchiveTransitionRule) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *AnalysisArchiveTransitionRule) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *AnalysisArchiveTransitionRule) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetExclude returns the Exclude field value if set, zero value otherwise.
func (o *AnalysisArchiveTransitionRule) GetExclude() AnalysisArchiveTransitionRuleExclude {
	if o == nil || IsNil(o.Exclude) {
		var ret AnalysisArchiveTransitionRuleExclude
		return ret
	}
	return *o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisArchiveTransitionRule) GetExcludeOk() (*AnalysisArchiveTransitionRuleExclude, bool) {
	if o == nil || IsNil(o.Exclude) {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *AnalysisArchiveTransitionRule) HasExclude() bool {
	if o != nil && !IsNil(o.Exclude) {
		return true
	}

	return false
}

// SetExclude gets a reference to the given AnalysisArchiveTransitionRuleExclude and assigns it to the Exclude field.
func (o *AnalysisArchiveTransitionRule) SetExclude(v AnalysisArchiveTransitionRuleExclude) {
	o.Exclude = &v
}

// GetMaxImagesPerAccount returns the MaxImagesPerAccount field value if set, zero value otherwise.
func (o *AnalysisArchiveTransitionRule) GetMaxImagesPerAccount() int32 {
	if o == nil || IsNil(o.MaxImagesPerAccount) {
		var ret int32
		return ret
	}
	return *o.MaxImagesPerAccount
}

// GetMaxImagesPerAccountOk returns a tuple with the MaxImagesPerAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisArchiveTransitionRule) GetMaxImagesPerAccountOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxImagesPerAccount) {
		return nil, false
	}
	return o.MaxImagesPerAccount, true
}

// HasMaxImagesPerAccount returns a boolean if a field has been set.
func (o *AnalysisArchiveTransitionRule) HasMaxImagesPerAccount() bool {
	if o != nil && !IsNil(o.MaxImagesPerAccount) {
		return true
	}

	return false
}

// SetMaxImagesPerAccount gets a reference to the given int32 and assigns it to the MaxImagesPerAccount field.
func (o *AnalysisArchiveTransitionRule) SetMaxImagesPerAccount(v int32) {
	o.MaxImagesPerAccount = &v
}

func (o AnalysisArchiveTransitionRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnalysisArchiveTransitionRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Selector) {
		toSerialize["selector"] = o.Selector
	}
	if !IsNil(o.RuleId) {
		toSerialize["rule_id"] = o.RuleId
	}
	if !IsNil(o.TagVersionsNewer) {
		toSerialize["tag_versions_newer"] = o.TagVersionsNewer
	}
	if !IsNil(o.AnalysisAgeDays) {
		toSerialize["analysis_age_days"] = o.AnalysisAgeDays
	}
	toSerialize["transition"] = o.Transition
	if !IsNil(o.SystemGlobal) {
		toSerialize["system_global"] = o.SystemGlobal
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if !IsNil(o.Exclude) {
		toSerialize["exclude"] = o.Exclude
	}
	if !IsNil(o.MaxImagesPerAccount) {
		toSerialize["max_images_per_account"] = o.MaxImagesPerAccount
	}
	return toSerialize, nil
}

func (o *AnalysisArchiveTransitionRule) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"transition",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAnalysisArchiveTransitionRule := _AnalysisArchiveTransitionRule{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAnalysisArchiveTransitionRule)

	if err != nil {
		return err
	}

	*o = AnalysisArchiveTransitionRule(varAnalysisArchiveTransitionRule)

	return err
}

type NullableAnalysisArchiveTransitionRule struct {
	value *AnalysisArchiveTransitionRule
	isSet bool
}

func (v NullableAnalysisArchiveTransitionRule) Get() *AnalysisArchiveTransitionRule {
	return v.value
}

func (v *NullableAnalysisArchiveTransitionRule) Set(val *AnalysisArchiveTransitionRule) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalysisArchiveTransitionRule) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalysisArchiveTransitionRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalysisArchiveTransitionRule(val *AnalysisArchiveTransitionRule) *NullableAnalysisArchiveTransitionRule {
	return &NullableAnalysisArchiveTransitionRule{value: val, isSet: true}
}

func (v NullableAnalysisArchiveTransitionRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalysisArchiveTransitionRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


