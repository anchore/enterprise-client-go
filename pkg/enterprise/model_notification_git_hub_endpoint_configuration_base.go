/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.10.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// checks if the NotificationGitHubEndpointConfigurationBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationGitHubEndpointConfigurationBase{}

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
	Url *string `json:"url,omitempty" validate:"regexp=https?:\\/\\/.*"`
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
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationBase) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationBase) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
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
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationBase) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationBase) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
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
	if o == nil || IsNil(o.VerifyTls) {
		var ret bool
		return ret
	}
	return *o.VerifyTls
}

// GetVerifyTlsOk returns a tuple with the VerifyTls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationBase) GetVerifyTlsOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifyTls) {
		return nil, false
	}
	return o.VerifyTls, true
}

// HasVerifyTls returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationBase) HasVerifyTls() bool {
	if o != nil && !IsNil(o.VerifyTls) {
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
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationBase) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationBase) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
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
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationBase) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationBase) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
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
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationBase) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationBase) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
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
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationBase) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationBase) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
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
	if o == nil || IsNil(o.AccessToken) {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationBase) GetAccessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.AccessToken) {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationBase) HasAccessToken() bool {
	if o != nil && !IsNil(o.AccessToken) {
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
	if o == nil || IsNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationBase) GetOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationBase) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
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
	if o == nil || IsNil(o.Repository) {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationBase) GetRepositoryOk() (*string, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationBase) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
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
	if o == nil || IsNil(o.Milestone) {
		var ret int32
		return ret
	}
	return *o.Milestone
}

// GetMilestoneOk returns a tuple with the Milestone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationBase) GetMilestoneOk() (*int32, bool) {
	if o == nil || IsNil(o.Milestone) {
		return nil, false
	}
	return o.Milestone, true
}

// HasMilestone returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationBase) HasMilestone() bool {
	if o != nil && !IsNil(o.Milestone) {
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
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationBase) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationBase) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
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
	if o == nil || IsNil(o.Assignees) {
		var ret []string
		return ret
	}
	return o.Assignees
}

// GetAssigneesOk returns a tuple with the Assignees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGitHubEndpointConfigurationBase) GetAssigneesOk() ([]string, bool) {
	if o == nil || IsNil(o.Assignees) {
		return nil, false
	}
	return o.Assignees, true
}

// HasAssignees returns a boolean if a field has been set.
func (o *NotificationGitHubEndpointConfigurationBase) HasAssignees() bool {
	if o != nil && !IsNil(o.Assignees) {
		return true
	}

	return false
}

// SetAssignees gets a reference to the given []string and assigns it to the Assignees field.
func (o *NotificationGitHubEndpointConfigurationBase) SetAssignees(v []string) {
	o.Assignees = v
}

func (o NotificationGitHubEndpointConfigurationBase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationGitHubEndpointConfigurationBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.VerifyTls) {
		toSerialize["verify_tls"] = o.VerifyTls
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.AccessToken) {
		toSerialize["access_token"] = o.AccessToken
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}
	if !IsNil(o.Milestone) {
		toSerialize["milestone"] = o.Milestone
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Assignees) {
		toSerialize["assignees"] = o.Assignees
	}
	return toSerialize, nil
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


