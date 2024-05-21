/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.6.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// NotificationJiraEndpointConfigurationBaseAllOf struct for NotificationJiraEndpointConfigurationBaseAllOf
type NotificationJiraEndpointConfigurationBaseAllOf struct {
	// Jira endpoint URL including host and port, should begin with 'http://' or 'https://'
	Url *string `json:"url,omitempty"`
	// Jira username for creating issues
	Username *string `json:"username,omitempty"`
	// Jira access token for creating issues
	Password *string `json:"password,omitempty"`
	// Key of the Jira project for creating issues
	ProjectKey *string `json:"project_key,omitempty"`
	// Type associated with the issue
	IssueType *string `json:"issue_type,omitempty"`
	// Priority assigned to the issue
	Priority *string `json:"priority,omitempty"`
	// Jira user to associate with the issue
	Assignee *string `json:"assignee,omitempty"`
	// List of labels to associate with the issue
	Labels []string `json:"labels,omitempty"`
}

// NewNotificationJiraEndpointConfigurationBaseAllOf instantiates a new NotificationJiraEndpointConfigurationBaseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationJiraEndpointConfigurationBaseAllOf() *NotificationJiraEndpointConfigurationBaseAllOf {
	this := NotificationJiraEndpointConfigurationBaseAllOf{}
	return &this
}

// NewNotificationJiraEndpointConfigurationBaseAllOfWithDefaults instantiates a new NotificationJiraEndpointConfigurationBaseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationJiraEndpointConfigurationBaseAllOfWithDefaults() *NotificationJiraEndpointConfigurationBaseAllOf {
	this := NotificationJiraEndpointConfigurationBaseAllOf{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *NotificationJiraEndpointConfigurationBaseAllOf) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationJiraEndpointConfigurationBaseAllOf) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *NotificationJiraEndpointConfigurationBaseAllOf) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *NotificationJiraEndpointConfigurationBaseAllOf) SetUrl(v string) {
	o.Url = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *NotificationJiraEndpointConfigurationBaseAllOf) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationJiraEndpointConfigurationBaseAllOf) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *NotificationJiraEndpointConfigurationBaseAllOf) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *NotificationJiraEndpointConfigurationBaseAllOf) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *NotificationJiraEndpointConfigurationBaseAllOf) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationJiraEndpointConfigurationBaseAllOf) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *NotificationJiraEndpointConfigurationBaseAllOf) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *NotificationJiraEndpointConfigurationBaseAllOf) SetPassword(v string) {
	o.Password = &v
}

// GetProjectKey returns the ProjectKey field value if set, zero value otherwise.
func (o *NotificationJiraEndpointConfigurationBaseAllOf) GetProjectKey() string {
	if o == nil || o.ProjectKey == nil {
		var ret string
		return ret
	}
	return *o.ProjectKey
}

// GetProjectKeyOk returns a tuple with the ProjectKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationJiraEndpointConfigurationBaseAllOf) GetProjectKeyOk() (*string, bool) {
	if o == nil || o.ProjectKey == nil {
		return nil, false
	}
	return o.ProjectKey, true
}

// HasProjectKey returns a boolean if a field has been set.
func (o *NotificationJiraEndpointConfigurationBaseAllOf) HasProjectKey() bool {
	if o != nil && o.ProjectKey != nil {
		return true
	}

	return false
}

// SetProjectKey gets a reference to the given string and assigns it to the ProjectKey field.
func (o *NotificationJiraEndpointConfigurationBaseAllOf) SetProjectKey(v string) {
	o.ProjectKey = &v
}

// GetIssueType returns the IssueType field value if set, zero value otherwise.
func (o *NotificationJiraEndpointConfigurationBaseAllOf) GetIssueType() string {
	if o == nil || o.IssueType == nil {
		var ret string
		return ret
	}
	return *o.IssueType
}

// GetIssueTypeOk returns a tuple with the IssueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationJiraEndpointConfigurationBaseAllOf) GetIssueTypeOk() (*string, bool) {
	if o == nil || o.IssueType == nil {
		return nil, false
	}
	return o.IssueType, true
}

// HasIssueType returns a boolean if a field has been set.
func (o *NotificationJiraEndpointConfigurationBaseAllOf) HasIssueType() bool {
	if o != nil && o.IssueType != nil {
		return true
	}

	return false
}

// SetIssueType gets a reference to the given string and assigns it to the IssueType field.
func (o *NotificationJiraEndpointConfigurationBaseAllOf) SetIssueType(v string) {
	o.IssueType = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *NotificationJiraEndpointConfigurationBaseAllOf) GetPriority() string {
	if o == nil || o.Priority == nil {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationJiraEndpointConfigurationBaseAllOf) GetPriorityOk() (*string, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *NotificationJiraEndpointConfigurationBaseAllOf) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *NotificationJiraEndpointConfigurationBaseAllOf) SetPriority(v string) {
	o.Priority = &v
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *NotificationJiraEndpointConfigurationBaseAllOf) GetAssignee() string {
	if o == nil || o.Assignee == nil {
		var ret string
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationJiraEndpointConfigurationBaseAllOf) GetAssigneeOk() (*string, bool) {
	if o == nil || o.Assignee == nil {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *NotificationJiraEndpointConfigurationBaseAllOf) HasAssignee() bool {
	if o != nil && o.Assignee != nil {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given string and assigns it to the Assignee field.
func (o *NotificationJiraEndpointConfigurationBaseAllOf) SetAssignee(v string) {
	o.Assignee = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *NotificationJiraEndpointConfigurationBaseAllOf) GetLabels() []string {
	if o == nil || o.Labels == nil {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationJiraEndpointConfigurationBaseAllOf) GetLabelsOk() ([]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *NotificationJiraEndpointConfigurationBaseAllOf) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *NotificationJiraEndpointConfigurationBaseAllOf) SetLabels(v []string) {
	o.Labels = v
}

func (o NotificationJiraEndpointConfigurationBaseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.ProjectKey != nil {
		toSerialize["project_key"] = o.ProjectKey
	}
	if o.IssueType != nil {
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

type NullableNotificationJiraEndpointConfigurationBaseAllOf struct {
	value *NotificationJiraEndpointConfigurationBaseAllOf
	isSet bool
}

func (v NullableNotificationJiraEndpointConfigurationBaseAllOf) Get() *NotificationJiraEndpointConfigurationBaseAllOf {
	return v.value
}

func (v *NullableNotificationJiraEndpointConfigurationBaseAllOf) Set(val *NotificationJiraEndpointConfigurationBaseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationJiraEndpointConfigurationBaseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationJiraEndpointConfigurationBaseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationJiraEndpointConfigurationBaseAllOf(val *NotificationJiraEndpointConfigurationBaseAllOf) *NullableNotificationJiraEndpointConfigurationBaseAllOf {
	return &NullableNotificationJiraEndpointConfigurationBaseAllOf{value: val, isSet: true}
}

func (v NullableNotificationJiraEndpointConfigurationBaseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationJiraEndpointConfigurationBaseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


