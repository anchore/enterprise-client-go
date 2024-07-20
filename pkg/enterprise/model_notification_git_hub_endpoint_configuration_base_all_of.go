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

// NotificationGitHubEndpointConfigurationBaseAllOf struct for NotificationGitHubEndpointConfigurationBaseAllOf
type NotificationGitHubEndpointConfigurationBaseAllOf struct {
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

// NewNotificationGitHubEndpointConfigurationBaseAllOf instantiates a new NotificationGitHubEndpointConfigurationBaseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationGitHubEndpointConfigurationBaseAllOf() *NotificationGitHubEndpointConfigurationBaseAllOf {
	this := NotificationGitHubEndpointConfigurationBaseAllOf{}
	return &this
}

// NewNotificationGitHubEndpointConfigurationBaseAllOfWithDefaults instantiates a new NotificationGitHubEndpointConfigurationBaseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationGitHubEndpointConfigurationBaseAllOfWithDefaults() *NotificationGitHubEndpointConfigurationBaseAllOf {
	this := NotificationGitHubEndpointConfigurationBaseAllOf{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *NotificationGitHubEndpointConfigurationBaseAllOf) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationBaseAllOf) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationBaseAllOf) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *NotificationGitHubEndpointConfigurationBaseAllOf) SetUrl(v string) {
	o.Url = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *NotificationGitHubEndpointConfigurationBaseAllOf) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationBaseAllOf) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationBaseAllOf) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *NotificationGitHubEndpointConfigurationBaseAllOf) SetUsername(v string) {
	o.Username = &v
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *NotificationGitHubEndpointConfigurationBaseAllOf) GetAccessToken() string {
	if o == nil || o.AccessToken == nil {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationBaseAllOf) GetAccessTokenOk() (*string, bool) {
	if o == nil || o.AccessToken == nil {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationBaseAllOf) HasAccessToken() bool {
	if o != nil && o.AccessToken != nil {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *NotificationGitHubEndpointConfigurationBaseAllOf) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *NotificationGitHubEndpointConfigurationBaseAllOf) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationBaseAllOf) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationBaseAllOf) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *NotificationGitHubEndpointConfigurationBaseAllOf) SetOwner(v string) {
	o.Owner = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *NotificationGitHubEndpointConfigurationBaseAllOf) GetRepository() string {
	if o == nil || o.Repository == nil {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationBaseAllOf) GetRepositoryOk() (*string, bool) {
	if o == nil || o.Repository == nil {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationBaseAllOf) HasRepository() bool {
	if o != nil && o.Repository != nil {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *NotificationGitHubEndpointConfigurationBaseAllOf) SetRepository(v string) {
	o.Repository = &v
}

// GetMilestone returns the Milestone field value if set, zero value otherwise.
func (o *NotificationGitHubEndpointConfigurationBaseAllOf) GetMilestone() int32 {
	if o == nil || o.Milestone == nil {
		var ret int32
		return ret
	}
	return *o.Milestone
}

// GetMilestoneOk returns a tuple with the Milestone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationBaseAllOf) GetMilestoneOk() (*int32, bool) {
	if o == nil || o.Milestone == nil {
		return nil, false
	}
	return o.Milestone, true
}

// HasMilestone returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationBaseAllOf) HasMilestone() bool {
	if o != nil && o.Milestone != nil {
		return true
	}

	return false
}

// SetMilestone gets a reference to the given int32 and assigns it to the Milestone field.
func (o *NotificationGitHubEndpointConfigurationBaseAllOf) SetMilestone(v int32) {
	o.Milestone = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *NotificationGitHubEndpointConfigurationBaseAllOf) GetLabels() []string {
	if o == nil || o.Labels == nil {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationBaseAllOf) GetLabelsOk() ([]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationBaseAllOf) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *NotificationGitHubEndpointConfigurationBaseAllOf) SetLabels(v []string) {
	o.Labels = v
}

// GetAssignees returns the Assignees field value if set, zero value otherwise.
func (o *NotificationGitHubEndpointConfigurationBaseAllOf) GetAssignees() []string {
	if o == nil || o.Assignees == nil {
		var ret []string
		return ret
	}
	return o.Assignees
}

// GetAssigneesOk returns a tuple with the Assignees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationBaseAllOf) GetAssigneesOk() ([]string, bool) {
	if o == nil || o.Assignees == nil {
		return nil, false
	}
	return o.Assignees, true
}

// HasAssignees returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationBaseAllOf) HasAssignees() bool {
	if o != nil && o.Assignees != nil {
		return true
	}

	return false
}

// SetAssignees gets a reference to the given []string and assigns it to the Assignees field.
func (o *NotificationGitHubEndpointConfigurationBaseAllOf) SetAssignees(v []string) {
	o.Assignees = v
}

func (o NotificationGitHubEndpointConfigurationBaseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableNotificationGitHubEndpointConfigurationBaseAllOf struct {
	value *NotificationGitHubEndpointConfigurationBaseAllOf
	isSet bool
}

func (v NullableNotificationGitHubEndpointConfigurationBaseAllOf) Get() *NotificationGitHubEndpointConfigurationBaseAllOf {
	return v.value
}

func (v *NullableNotificationGitHubEndpointConfigurationBaseAllOf) Set(val *NotificationGitHubEndpointConfigurationBaseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationGitHubEndpointConfigurationBaseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationGitHubEndpointConfigurationBaseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationGitHubEndpointConfigurationBaseAllOf(val *NotificationGitHubEndpointConfigurationBaseAllOf) *NullableNotificationGitHubEndpointConfigurationBaseAllOf {
	return &NullableNotificationGitHubEndpointConfigurationBaseAllOf{value: val, isSet: true}
}

func (v NullableNotificationGitHubEndpointConfigurationBaseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationGitHubEndpointConfigurationBaseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


