/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.5.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// NotificationGitHubEndpointConfigurationBase Configuration for GitHub endpoint (Base model)
type NotificationGitHubEndpointConfigurationBase struct {
	// The instance identifier for the configuration
	Uuid *string `json:"uuid,omitempty"`
	// User friendly name or description for the configuration
	Description *string `json:"description,omitempty"`
	// Verify the cert if using tls for connecting externally. Defaults to true if not specified
	VerifyTls *bool `json:"verify_tls,omitempty"`
	// Timestamp for last modification to the record
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Timestamp for last modification to the record
	LastUpdated *time.Time `json:"last_updated,omitempty"`
	// Github API endpoint, defaults to https://api.github.com if not specified
	Url *string `json:"url,omitempty"`
	// GitHub username for creating issues
	Username *string `json:"username,omitempty"`
	// Personal access token for the GitHub account
	AccessToken *string `json:"access_token,omitempty"`
	// Owner of the repository to create issues against
	Owner *string `json:"owner,omitempty"`
	// Name of the repository to create issues against
	Repository *string `json:"repository,omitempty"`
	// Number of the milestone to associate with the issue
	Milestone *int32 `json:"milestone,omitempty"`
	// List of labels to associate with the issue
	Labels []string `json:"labels,omitempty"`
	// List of user logins to assign to the issue.
	Assignees []string `json:"assignees,omitempty"`
}

// NewNotificationGitHubEndpointConfigurationBase instantiates a new NotificationGitHubEndpointConfigurationBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationGitHubEndpointConfigurationBase() *NotificationGitHubEndpointConfigurationBase {
	this := NotificationGitHubEndpointConfigurationBase{}
	return &this
}

// NewNotificationGitHubEndpointConfigurationBaseWithDefaults instantiates a new NotificationGitHubEndpointConfigurationBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationGitHubEndpointConfigurationBaseWithDefaults() *NotificationGitHubEndpointConfigurationBase {
	this := NotificationGitHubEndpointConfigurationBase{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *NotificationGitHubEndpointConfigurationBase) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationBase) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationBase) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *NotificationGitHubEndpointConfigurationBase) SetUuid(v string) {
	o.Uuid = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NotificationGitHubEndpointConfigurationBase) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationBase) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationBase) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NotificationGitHubEndpointConfigurationBase) SetDescription(v string) {
	o.Description = &v
}

// GetVerifyTls returns the VerifyTls field value if set, zero value otherwise.
func (o *NotificationGitHubEndpointConfigurationBase) GetVerifyTls() bool {
	if o == nil || o.VerifyTls == nil {
		var ret bool
		return ret
	}
	return *o.VerifyTls
}

// GetVerifyTlsOk returns a tuple with the VerifyTls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationBase) GetVerifyTlsOk() (*bool, bool) {
	if o == nil || o.VerifyTls == nil {
		return nil, false
	}
	return o.VerifyTls, true
}

// HasVerifyTls returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationBase) HasVerifyTls() bool {
	if o != nil && o.VerifyTls != nil {
		return true
	}

	return false
}

// SetVerifyTls gets a reference to the given bool and assigns it to the VerifyTls field.
func (o *NotificationGitHubEndpointConfigurationBase) SetVerifyTls(v bool) {
	o.VerifyTls = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *NotificationGitHubEndpointConfigurationBase) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationBase) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationBase) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *NotificationGitHubEndpointConfigurationBase) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *NotificationGitHubEndpointConfigurationBase) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationBase) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationBase) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *NotificationGitHubEndpointConfigurationBase) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *NotificationGitHubEndpointConfigurationBase) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationBase) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationBase) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *NotificationGitHubEndpointConfigurationBase) SetUrl(v string) {
	o.Url = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *NotificationGitHubEndpointConfigurationBase) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationBase) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationBase) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *NotificationGitHubEndpointConfigurationBase) SetUsername(v string) {
	o.Username = &v
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *NotificationGitHubEndpointConfigurationBase) GetAccessToken() string {
	if o == nil || o.AccessToken == nil {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationBase) GetAccessTokenOk() (*string, bool) {
	if o == nil || o.AccessToken == nil {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationBase) HasAccessToken() bool {
	if o != nil && o.AccessToken != nil {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *NotificationGitHubEndpointConfigurationBase) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *NotificationGitHubEndpointConfigurationBase) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationBase) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationBase) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *NotificationGitHubEndpointConfigurationBase) SetOwner(v string) {
	o.Owner = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *NotificationGitHubEndpointConfigurationBase) GetRepository() string {
	if o == nil || o.Repository == nil {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationBase) GetRepositoryOk() (*string, bool) {
	if o == nil || o.Repository == nil {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationBase) HasRepository() bool {
	if o != nil && o.Repository != nil {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *NotificationGitHubEndpointConfigurationBase) SetRepository(v string) {
	o.Repository = &v
}

// GetMilestone returns the Milestone field value if set, zero value otherwise.
func (o *NotificationGitHubEndpointConfigurationBase) GetMilestone() int32 {
	if o == nil || o.Milestone == nil {
		var ret int32
		return ret
	}
	return *o.Milestone
}

// GetMilestoneOk returns a tuple with the Milestone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationBase) GetMilestoneOk() (*int32, bool) {
	if o == nil || o.Milestone == nil {
		return nil, false
	}
	return o.Milestone, true
}

// HasMilestone returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationBase) HasMilestone() bool {
	if o != nil && o.Milestone != nil {
		return true
	}

	return false
}

// SetMilestone gets a reference to the given int32 and assigns it to the Milestone field.
func (o *NotificationGitHubEndpointConfigurationBase) SetMilestone(v int32) {
	o.Milestone = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *NotificationGitHubEndpointConfigurationBase) GetLabels() []string {
	if o == nil || o.Labels == nil {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationBase) GetLabelsOk() ([]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationBase) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *NotificationGitHubEndpointConfigurationBase) SetLabels(v []string) {
	o.Labels = v
}

// GetAssignees returns the Assignees field value if set, zero value otherwise.
func (o *NotificationGitHubEndpointConfigurationBase) GetAssignees() []string {
	if o == nil || o.Assignees == nil {
		var ret []string
		return ret
	}
	return o.Assignees
}

// GetAssigneesOk returns a tuple with the Assignees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationBase) GetAssigneesOk() ([]string, bool) {
	if o == nil || o.Assignees == nil {
		return nil, false
	}
	return o.Assignees, true
}

// HasAssignees returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationBase) HasAssignees() bool {
	if o != nil && o.Assignees != nil {
		return true
	}

	return false
}

// SetAssignees gets a reference to the given []string and assigns it to the Assignees field.
func (o *NotificationGitHubEndpointConfigurationBase) SetAssignees(v []string) {
	o.Assignees = v
}

func (o NotificationGitHubEndpointConfigurationBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.VerifyTls != nil {
		toSerialize["verify_tls"] = o.VerifyTls
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.AccessToken != nil {
		toSerialize["access_token"] = o.AccessToken
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.Repository != nil {
		toSerialize["repository"] = o.Repository
	}
	if o.Milestone != nil {
		toSerialize["milestone"] = o.Milestone
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.Assignees != nil {
		toSerialize["assignees"] = o.Assignees
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationGitHubEndpointConfigurationBase struct {
	value *NotificationGitHubEndpointConfigurationBase
	isSet bool
}

func (v NullableNotificationGitHubEndpointConfigurationBase) Get() *NotificationGitHubEndpointConfigurationBase {
	return v.value
}

func (v *NullableNotificationGitHubEndpointConfigurationBase) Set(val *NotificationGitHubEndpointConfigurationBase) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationGitHubEndpointConfigurationBase) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationGitHubEndpointConfigurationBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationGitHubEndpointConfigurationBase(val *NotificationGitHubEndpointConfigurationBase) *NullableNotificationGitHubEndpointConfigurationBase {
	return &NullableNotificationGitHubEndpointConfigurationBase{value: val, isSet: true}
}

func (v NullableNotificationGitHubEndpointConfigurationBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationGitHubEndpointConfigurationBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


