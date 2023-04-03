/*
Anchore Enterprise API Server

This is the Anchore Enterprise API. It provides additional external API routes and functionality for enterprise users.

API version: 0.7.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// RuntimeComplianceCheck The result of a runtime compliance check
type RuntimeComplianceCheck struct {
	// The type of runtime compliance check
	CheckType string `json:"check_type"`
	// The result of the runtime compliance check
	Result *string `json:"result,omitempty"`
	// The pod the check was run against
	Pod *string `json:"pod,omitempty"`
	// The namespace of the pod the check was run against
	Namespace *string `json:"namespace,omitempty"`
	// The tag of image in the pod the check was run against
	ImageTag *string `json:"image_tag,omitempty"`
	// The digest of the pod the check was run against
	ImageDigest *string `json:"image_digest,omitempty"`
	// RFC 3339 formatted UTC timestamp when the runtime check started
	StartTime *time.Time `json:"start_time,omitempty"`
	// RFC 3339 formatted UTC timestamp when the runtime check ended
	EndTime *time.Time `json:"end_time,omitempty"`
	// Ids of the files generated by the runtime compliance check
	ComplianceFileIds *map[string]string `json:"compliance_file_ids,omitempty"`
}

// NewRuntimeComplianceCheck instantiates a new RuntimeComplianceCheck object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuntimeComplianceCheck(checkType string) *RuntimeComplianceCheck {
	this := RuntimeComplianceCheck{}
	this.CheckType = checkType
	return &this
}

// NewRuntimeComplianceCheckWithDefaults instantiates a new RuntimeComplianceCheck object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuntimeComplianceCheckWithDefaults() *RuntimeComplianceCheck {
	this := RuntimeComplianceCheck{}
	return &this
}

// GetCheckType returns the CheckType field value
func (o *RuntimeComplianceCheck) GetCheckType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CheckType
}

// GetCheckTypeOk returns a tuple with the CheckType field value
// and a boolean to check if the value has been set.
func (o *RuntimeComplianceCheck) GetCheckTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CheckType, true
}

// SetCheckType sets field value
func (o *RuntimeComplianceCheck) SetCheckType(v string) {
	o.CheckType = v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *RuntimeComplianceCheck) GetResult() string {
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuntimeComplianceCheck) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *RuntimeComplianceCheck) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *RuntimeComplianceCheck) SetResult(v string) {
	o.Result = &v
}

// GetPod returns the Pod field value if set, zero value otherwise.
func (o *RuntimeComplianceCheck) GetPod() string {
	if o == nil || o.Pod == nil {
		var ret string
		return ret
	}
	return *o.Pod
}

// GetPodOk returns a tuple with the Pod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuntimeComplianceCheck) GetPodOk() (*string, bool) {
	if o == nil || o.Pod == nil {
		return nil, false
	}
	return o.Pod, true
}

// HasPod returns a boolean if a field has been set.
func (o *RuntimeComplianceCheck) HasPod() bool {
	if o != nil && o.Pod != nil {
		return true
	}

	return false
}

// SetPod gets a reference to the given string and assigns it to the Pod field.
func (o *RuntimeComplianceCheck) SetPod(v string) {
	o.Pod = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *RuntimeComplianceCheck) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuntimeComplianceCheck) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *RuntimeComplianceCheck) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *RuntimeComplianceCheck) SetNamespace(v string) {
	o.Namespace = &v
}

// GetImageTag returns the ImageTag field value if set, zero value otherwise.
func (o *RuntimeComplianceCheck) GetImageTag() string {
	if o == nil || o.ImageTag == nil {
		var ret string
		return ret
	}
	return *o.ImageTag
}

// GetImageTagOk returns a tuple with the ImageTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuntimeComplianceCheck) GetImageTagOk() (*string, bool) {
	if o == nil || o.ImageTag == nil {
		return nil, false
	}
	return o.ImageTag, true
}

// HasImageTag returns a boolean if a field has been set.
func (o *RuntimeComplianceCheck) HasImageTag() bool {
	if o != nil && o.ImageTag != nil {
		return true
	}

	return false
}

// SetImageTag gets a reference to the given string and assigns it to the ImageTag field.
func (o *RuntimeComplianceCheck) SetImageTag(v string) {
	o.ImageTag = &v
}

// GetImageDigest returns the ImageDigest field value if set, zero value otherwise.
func (o *RuntimeComplianceCheck) GetImageDigest() string {
	if o == nil || o.ImageDigest == nil {
		var ret string
		return ret
	}
	return *o.ImageDigest
}

// GetImageDigestOk returns a tuple with the ImageDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuntimeComplianceCheck) GetImageDigestOk() (*string, bool) {
	if o == nil || o.ImageDigest == nil {
		return nil, false
	}
	return o.ImageDigest, true
}

// HasImageDigest returns a boolean if a field has been set.
func (o *RuntimeComplianceCheck) HasImageDigest() bool {
	if o != nil && o.ImageDigest != nil {
		return true
	}

	return false
}

// SetImageDigest gets a reference to the given string and assigns it to the ImageDigest field.
func (o *RuntimeComplianceCheck) SetImageDigest(v string) {
	o.ImageDigest = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *RuntimeComplianceCheck) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuntimeComplianceCheck) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *RuntimeComplianceCheck) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *RuntimeComplianceCheck) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *RuntimeComplianceCheck) GetEndTime() time.Time {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuntimeComplianceCheck) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *RuntimeComplianceCheck) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *RuntimeComplianceCheck) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetComplianceFileIds returns the ComplianceFileIds field value if set, zero value otherwise.
func (o *RuntimeComplianceCheck) GetComplianceFileIds() map[string]string {
	if o == nil || o.ComplianceFileIds == nil {
		var ret map[string]string
		return ret
	}
	return *o.ComplianceFileIds
}

// GetComplianceFileIdsOk returns a tuple with the ComplianceFileIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuntimeComplianceCheck) GetComplianceFileIdsOk() (*map[string]string, bool) {
	if o == nil || o.ComplianceFileIds == nil {
		return nil, false
	}
	return o.ComplianceFileIds, true
}

// HasComplianceFileIds returns a boolean if a field has been set.
func (o *RuntimeComplianceCheck) HasComplianceFileIds() bool {
	if o != nil && o.ComplianceFileIds != nil {
		return true
	}

	return false
}

// SetComplianceFileIds gets a reference to the given map[string]string and assigns it to the ComplianceFileIds field.
func (o *RuntimeComplianceCheck) SetComplianceFileIds(v map[string]string) {
	o.ComplianceFileIds = &v
}

func (o RuntimeComplianceCheck) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["check_type"] = o.CheckType
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.Pod != nil {
		toSerialize["pod"] = o.Pod
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.ImageTag != nil {
		toSerialize["image_tag"] = o.ImageTag
	}
	if o.ImageDigest != nil {
		toSerialize["image_digest"] = o.ImageDigest
	}
	if o.StartTime != nil {
		toSerialize["start_time"] = o.StartTime
	}
	if o.EndTime != nil {
		toSerialize["end_time"] = o.EndTime
	}
	if o.ComplianceFileIds != nil {
		toSerialize["compliance_file_ids"] = o.ComplianceFileIds
	}
	return json.Marshal(toSerialize)
}

type NullableRuntimeComplianceCheck struct {
	value *RuntimeComplianceCheck
	isSet bool
}

func (v NullableRuntimeComplianceCheck) Get() *RuntimeComplianceCheck {
	return v.value
}

func (v *NullableRuntimeComplianceCheck) Set(val *RuntimeComplianceCheck) {
	v.value = val
	v.isSet = true
}

func (v NullableRuntimeComplianceCheck) IsSet() bool {
	return v.isSet
}

func (v *NullableRuntimeComplianceCheck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuntimeComplianceCheck(val *RuntimeComplianceCheck) *NullableRuntimeComplianceCheck {
	return &NullableRuntimeComplianceCheck{value: val, isSet: true}
}

func (v NullableRuntimeComplianceCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuntimeComplianceCheck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


