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
	"time"
)

// NotificationGitHubEndpointConfigurationPost Configuration for GitHub endpoint (Base model)
type NotificationGitHubEndpointConfigurationPost struct {
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
	Username string `json:"username"`
	// Personal access token for the GitHub account
	AccessToken string `json:"access_token"`
	// Owner of the repository to create issues against
	Owner string `json:"owner"`
	// Name of the repository to create issues against
	Repository string `json:"repository"`
	// Number of the milestone to associate with the issue
	Milestone *int32 `json:"milestone,omitempty"`
	// List of labels to associate with the issue
	Labels []string `json:"labels,omitempty"`
	// List of user logins to assign to the issue.
	Assignees []string `json:"assignees,omitempty"`
}

// NewNotificationGitHubEndpointConfigurationPost instantiates a new NotificationGitHubEndpointConfigurationPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationGitHubEndpointConfigurationPost(username string, accessToken string, owner string, repository string) *NotificationGitHubEndpointConfigurationPost {
	this := NotificationGitHubEndpointConfigurationPost{}
	this.Username = username
	this.AccessToken = accessToken
	this.Owner = owner
	this.Repository = repository
	return &this
}

// NewNotificationGitHubEndpointConfigurationPostWithDefaults instantiates a new NotificationGitHubEndpointConfigurationPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationGitHubEndpointConfigurationPostWithDefaults() *NotificationGitHubEndpointConfigurationPost {
	this := NotificationGitHubEndpointConfigurationPost{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *NotificationGitHubEndpointConfigurationPost) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationPost) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationPost) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *NotificationGitHubEndpointConfigurationPost) SetUuid(v string) {
	o.Uuid = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NotificationGitHubEndpointConfigurationPost) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationPost) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationPost) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NotificationGitHubEndpointConfigurationPost) SetDescription(v string) {
	o.Description = &v
}

// GetVerifyTls returns the VerifyTls field value if set, zero value otherwise.
func (o *NotificationGitHubEndpointConfigurationPost) GetVerifyTls() bool {
	if o == nil || o.VerifyTls == nil {
		var ret bool
		return ret
	}
	return *o.VerifyTls
}

// GetVerifyTlsOk returns a tuple with the VerifyTls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationPost) GetVerifyTlsOk() (*bool, bool) {
	if o == nil || o.VerifyTls == nil {
		return nil, false
	}
	return o.VerifyTls, true
}

// HasVerifyTls returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationPost) HasVerifyTls() bool {
	if o != nil && o.VerifyTls != nil {
		return true
	}

	return false
}

// SetVerifyTls gets a reference to the given bool and assigns it to the VerifyTls field.
func (o *NotificationGitHubEndpointConfigurationPost) SetVerifyTls(v bool) {
	o.VerifyTls = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *NotificationGitHubEndpointConfigurationPost) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationPost) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationPost) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *NotificationGitHubEndpointConfigurationPost) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *NotificationGitHubEndpointConfigurationPost) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationPost) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationPost) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *NotificationGitHubEndpointConfigurationPost) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *NotificationGitHubEndpointConfigurationPost) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationPost) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationPost) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *NotificationGitHubEndpointConfigurationPost) SetUrl(v string) {
	o.Url = &v
}

// GetUsername returns the Username field value
func (o *NotificationGitHubEndpointConfigurationPost) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationPost) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *NotificationGitHubEndpointConfigurationPost) SetUsername(v string) {
	o.Username = v
}

// GetAccessToken returns the AccessToken field value
func (o *NotificationGitHubEndpointConfigurationPost) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationPost) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *NotificationGitHubEndpointConfigurationPost) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetOwner returns the Owner field value
func (o *NotificationGitHubEndpointConfigurationPost) GetOwner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationPost) GetOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *NotificationGitHubEndpointConfigurationPost) SetOwner(v string) {
	o.Owner = v
}

// GetRepository returns the Repository field value
func (o *NotificationGitHubEndpointConfigurationPost) GetRepository() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationPost) GetRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *NotificationGitHubEndpointConfigurationPost) SetRepository(v string) {
	o.Repository = v
}

// GetMilestone returns the Milestone field value if set, zero value otherwise.
func (o *NotificationGitHubEndpointConfigurationPost) GetMilestone() int32 {
	if o == nil || o.Milestone == nil {
		var ret int32
		return ret
	}
	return *o.Milestone
}

// GetMilestoneOk returns a tuple with the Milestone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationPost) GetMilestoneOk() (*int32, bool) {
	if o == nil || o.Milestone == nil {
		return nil, false
	}
	return o.Milestone, true
}

// HasMilestone returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationPost) HasMilestone() bool {
	if o != nil && o.Milestone != nil {
		return true
	}

	return false
}

// SetMilestone gets a reference to the given int32 and assigns it to the Milestone field.
func (o *NotificationGitHubEndpointConfigurationPost) SetMilestone(v int32) {
	o.Milestone = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *NotificationGitHubEndpointConfigurationPost) GetLabels() []string {
	if o == nil || o.Labels == nil {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationPost) GetLabelsOk() ([]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationPost) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *NotificationGitHubEndpointConfigurationPost) SetLabels(v []string) {
	o.Labels = v
}

// GetAssignees returns the Assignees field value if set, zero value otherwise.
func (o *NotificationGitHubEndpointConfigurationPost) GetAssignees() []string {
	if o == nil || o.Assignees == nil {
		var ret []string
		return ret
	}
	return o.Assignees
}

// GetAssigneesOk returns a tuple with the Assignees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationPost) GetAssigneesOk() ([]string, bool) {
	if o == nil || o.Assignees == nil {
		return nil, false
	}
	return o.Assignees, true
}

// HasAssignees returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationPost) HasAssignees() bool {
	if o != nil && o.Assignees != nil {
		return true
	}

	return false
}

// SetAssignees gets a reference to the given []string and assigns it to the Assignees field.
func (o *NotificationGitHubEndpointConfigurationPost) SetAssignees(v []string) {
	o.Assignees = v
}

func (o NotificationGitHubEndpointConfigurationPost) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	if true {
		toSerialize["owner"] = o.Owner
	}
	if true {
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

type NullableNotificationGitHubEndpointConfigurationPost struct {
	value *NotificationGitHubEndpointConfigurationPost
	isSet bool
}

func (v NullableNotificationGitHubEndpointConfigurationPost) Get() *NotificationGitHubEndpointConfigurationPost {
	return v.value
}

func (v *NullableNotificationGitHubEndpointConfigurationPost) Set(val *NotificationGitHubEndpointConfigurationPost) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationGitHubEndpointConfigurationPost) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationGitHubEndpointConfigurationPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationGitHubEndpointConfigurationPost(val *NotificationGitHubEndpointConfigurationPost) *NullableNotificationGitHubEndpointConfigurationPost {
	return &NullableNotificationGitHubEndpointConfigurationPost{value: val, isSet: true}
}

func (v NullableNotificationGitHubEndpointConfigurationPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationGitHubEndpointConfigurationPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


