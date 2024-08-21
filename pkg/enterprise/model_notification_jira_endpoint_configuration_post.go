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
	"time"
)

// NotificationJiraEndpointConfigurationPost Configuration for jira endpoint
type NotificationJiraEndpointConfigurationPost struct {
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
	// Jira endpoint URL including host and port, should begin with 'http://' or 'https://'
	Url string `json:"url"`
	// Jira username for creating issues
	Username string `json:"username"`
	// Jira access token for creating issues
	Password string `json:"password"`
	// Key of the Jira project for creating issues
	ProjectKey string `json:"project_key"`
	// Type associated with the issue
	IssueType string `json:"issue_type"`
	// Priority assigned to the issue
	Priority *string `json:"priority,omitempty"`
	// Jira user to associate with the issue
	Assignee *string `json:"assignee,omitempty"`
	// List of labels to associate with the issue
	Labels []string `json:"labels,omitempty"`
}

// NewNotificationJiraEndpointConfigurationPost instantiates a new NotificationJiraEndpointConfigurationPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationJiraEndpointConfigurationPost(url string, username string, password string, projectKey string, issueType string) *NotificationJiraEndpointConfigurationPost {
	this := NotificationJiraEndpointConfigurationPost{}
	this.Url = url
	this.Username = username
	this.Password = password
	this.ProjectKey = projectKey
	this.IssueType = issueType
	return &this
}

// NewNotificationJiraEndpointConfigurationPostWithDefaults instantiates a new NotificationJiraEndpointConfigurationPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationJiraEndpointConfigurationPostWithDefaults() *NotificationJiraEndpointConfigurationPost {
	this := NotificationJiraEndpointConfigurationPost{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *NotificationJiraEndpointConfigurationPost) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationJiraEndpointConfigurationPost) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *NotificationJiraEndpointConfigurationPost) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *NotificationJiraEndpointConfigurationPost) SetUuid(v string) {
	o.Uuid = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NotificationJiraEndpointConfigurationPost) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationJiraEndpointConfigurationPost) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NotificationJiraEndpointConfigurationPost) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NotificationJiraEndpointConfigurationPost) SetDescription(v string) {
	o.Description = &v
}

// GetVerifyTls returns the VerifyTls field value if set, zero value otherwise.
func (o *NotificationJiraEndpointConfigurationPost) GetVerifyTls() bool {
	if o == nil || o.VerifyTls == nil {
		var ret bool
		return ret
	}
	return *o.VerifyTls
}

// GetVerifyTlsOk returns a tuple with the VerifyTls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationJiraEndpointConfigurationPost) GetVerifyTlsOk() (*bool, bool) {
	if o == nil || o.VerifyTls == nil {
		return nil, false
	}
	return o.VerifyTls, true
}

// HasVerifyTls returns a boolean if a field has been set.
func (o *NotificationJiraEndpointConfigurationPost) HasVerifyTls() bool {
	if o != nil && o.VerifyTls != nil {
		return true
	}

	return false
}

// SetVerifyTls gets a reference to the given bool and assigns it to the VerifyTls field.
func (o *NotificationJiraEndpointConfigurationPost) SetVerifyTls(v bool) {
	o.VerifyTls = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *NotificationJiraEndpointConfigurationPost) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationJiraEndpointConfigurationPost) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *NotificationJiraEndpointConfigurationPost) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *NotificationJiraEndpointConfigurationPost) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *NotificationJiraEndpointConfigurationPost) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationJiraEndpointConfigurationPost) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *NotificationJiraEndpointConfigurationPost) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *NotificationJiraEndpointConfigurationPost) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetUrl returns the Url field value
func (o *NotificationJiraEndpointConfigurationPost) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NotificationJiraEndpointConfigurationPost) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NotificationJiraEndpointConfigurationPost) SetUrl(v string) {
	o.Url = v
}

// GetUsername returns the Username field value
func (o *NotificationJiraEndpointConfigurationPost) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *NotificationJiraEndpointConfigurationPost) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *NotificationJiraEndpointConfigurationPost) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *NotificationJiraEndpointConfigurationPost) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *NotificationJiraEndpointConfigurationPost) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *NotificationJiraEndpointConfigurationPost) SetPassword(v string) {
	o.Password = v
}

// GetProjectKey returns the ProjectKey field value
func (o *NotificationJiraEndpointConfigurationPost) GetProjectKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectKey
}

// GetProjectKeyOk returns a tuple with the ProjectKey field value
// and a boolean to check if the value has been set.
func (o *NotificationJiraEndpointConfigurationPost) GetProjectKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectKey, true
}

// SetProjectKey sets field value
func (o *NotificationJiraEndpointConfigurationPost) SetProjectKey(v string) {
	o.ProjectKey = v
}

// GetIssueType returns the IssueType field value
func (o *NotificationJiraEndpointConfigurationPost) GetIssueType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssueType
}

// GetIssueTypeOk returns a tuple with the IssueType field value
// and a boolean to check if the value has been set.
func (o *NotificationJiraEndpointConfigurationPost) GetIssueTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueType, true
}

// SetIssueType sets field value
func (o *NotificationJiraEndpointConfigurationPost) SetIssueType(v string) {
	o.IssueType = v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *NotificationJiraEndpointConfigurationPost) GetPriority() string {
	if o == nil || o.Priority == nil {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationJiraEndpointConfigurationPost) GetPriorityOk() (*string, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *NotificationJiraEndpointConfigurationPost) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *NotificationJiraEndpointConfigurationPost) SetPriority(v string) {
	o.Priority = &v
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *NotificationJiraEndpointConfigurationPost) GetAssignee() string {
	if o == nil || o.Assignee == nil {
		var ret string
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationJiraEndpointConfigurationPost) GetAssigneeOk() (*string, bool) {
	if o == nil || o.Assignee == nil {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *NotificationJiraEndpointConfigurationPost) HasAssignee() bool {
	if o != nil && o.Assignee != nil {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given string and assigns it to the Assignee field.
func (o *NotificationJiraEndpointConfigurationPost) SetAssignee(v string) {
	o.Assignee = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *NotificationJiraEndpointConfigurationPost) GetLabels() []string {
	if o == nil || o.Labels == nil {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationJiraEndpointConfigurationPost) GetLabelsOk() ([]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *NotificationJiraEndpointConfigurationPost) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *NotificationJiraEndpointConfigurationPost) SetLabels(v []string) {
	o.Labels = v
}

func (o NotificationJiraEndpointConfigurationPost) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["project_key"] = o.ProjectKey
	}
	if true {
		toSerialize["issue_type"] = o.IssueType
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.Assignee != nil {
		toSerialize["assignee"] = o.Assignee
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationJiraEndpointConfigurationPost struct {
	value *NotificationJiraEndpointConfigurationPost
	isSet bool
}

func (v NullableNotificationJiraEndpointConfigurationPost) Get() *NotificationJiraEndpointConfigurationPost {
	return v.value
}

func (v *NullableNotificationJiraEndpointConfigurationPost) Set(val *NotificationJiraEndpointConfigurationPost) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationJiraEndpointConfigurationPost) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationJiraEndpointConfigurationPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationJiraEndpointConfigurationPost(val *NotificationJiraEndpointConfigurationPost) *NullableNotificationJiraEndpointConfigurationPost {
	return &NullableNotificationJiraEndpointConfigurationPost{value: val, isSet: true}
}

func (v NullableNotificationJiraEndpointConfigurationPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationJiraEndpointConfigurationPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


