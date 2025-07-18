/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.12.0
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

// checks if the NotificationJiraEndpointConfigurationPut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationJiraEndpointConfigurationPut{}

// NotificationJiraEndpointConfigurationPut Configuration for jira endpoint
type NotificationJiraEndpointConfigurationPut struct {
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
	Url string `json:"url" validate:"regexp=https?:\\/\\/.*"`
	// Jira username for creating issues
	Username string `json:"username"`
	// Jira access token for creating issues
	Password *string `json:"password,omitempty"`
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

type _NotificationJiraEndpointConfigurationPut NotificationJiraEndpointConfigurationPut

// NewNotificationJiraEndpointConfigurationPut instantiates a new NotificationJiraEndpointConfigurationPut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationJiraEndpointConfigurationPut(url string, username string, projectKey string, issueType string) *NotificationJiraEndpointConfigurationPut {
	this := NotificationJiraEndpointConfigurationPut{}
	this.Url = url
	this.Username = username
	this.ProjectKey = projectKey
	this.IssueType = issueType
	return &this
}

// NewNotificationJiraEndpointConfigurationPutWithDefaults instantiates a new NotificationJiraEndpointConfigurationPut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationJiraEndpointConfigurationPutWithDefaults() *NotificationJiraEndpointConfigurationPut {
	this := NotificationJiraEndpointConfigurationPut{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *NotificationJiraEndpointConfigurationPut) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationJiraEndpointConfigurationPut) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *NotificationJiraEndpointConfigurationPut) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *NotificationJiraEndpointConfigurationPut) SetUuid(v string) {
	o.Uuid = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NotificationJiraEndpointConfigurationPut) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationJiraEndpointConfigurationPut) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NotificationJiraEndpointConfigurationPut) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NotificationJiraEndpointConfigurationPut) SetDescription(v string) {
	o.Description = &v
}

// GetVerifyTls returns the VerifyTls field value if set, zero value otherwise.
func (o *NotificationJiraEndpointConfigurationPut) GetVerifyTls() bool {
	if o == nil || IsNil(o.VerifyTls) {
		var ret bool
		return ret
	}
	return *o.VerifyTls
}

// GetVerifyTlsOk returns a tuple with the VerifyTls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationJiraEndpointConfigurationPut) GetVerifyTlsOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifyTls) {
		return nil, false
	}
	return o.VerifyTls, true
}

// HasVerifyTls returns a boolean if a field has been set.
func (o *NotificationJiraEndpointConfigurationPut) HasVerifyTls() bool {
	if o != nil && !IsNil(o.VerifyTls) {
		return true
	}

	return false
}

// SetVerifyTls gets a reference to the given bool and assigns it to the VerifyTls field.
func (o *NotificationJiraEndpointConfigurationPut) SetVerifyTls(v bool) {
	o.VerifyTls = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *NotificationJiraEndpointConfigurationPut) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationJiraEndpointConfigurationPut) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *NotificationJiraEndpointConfigurationPut) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *NotificationJiraEndpointConfigurationPut) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *NotificationJiraEndpointConfigurationPut) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationJiraEndpointConfigurationPut) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *NotificationJiraEndpointConfigurationPut) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *NotificationJiraEndpointConfigurationPut) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetUrl returns the Url field value
func (o *NotificationJiraEndpointConfigurationPut) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NotificationJiraEndpointConfigurationPut) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NotificationJiraEndpointConfigurationPut) SetUrl(v string) {
	o.Url = v
}

// GetUsername returns the Username field value
func (o *NotificationJiraEndpointConfigurationPut) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *NotificationJiraEndpointConfigurationPut) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *NotificationJiraEndpointConfigurationPut) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *NotificationJiraEndpointConfigurationPut) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationJiraEndpointConfigurationPut) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *NotificationJiraEndpointConfigurationPut) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *NotificationJiraEndpointConfigurationPut) SetPassword(v string) {
	o.Password = &v
}

// GetProjectKey returns the ProjectKey field value
func (o *NotificationJiraEndpointConfigurationPut) GetProjectKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectKey
}

// GetProjectKeyOk returns a tuple with the ProjectKey field value
// and a boolean to check if the value has been set.
func (o *NotificationJiraEndpointConfigurationPut) GetProjectKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectKey, true
}

// SetProjectKey sets field value
func (o *NotificationJiraEndpointConfigurationPut) SetProjectKey(v string) {
	o.ProjectKey = v
}

// GetIssueType returns the IssueType field value
func (o *NotificationJiraEndpointConfigurationPut) GetIssueType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssueType
}

// GetIssueTypeOk returns a tuple with the IssueType field value
// and a boolean to check if the value has been set.
func (o *NotificationJiraEndpointConfigurationPut) GetIssueTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueType, true
}

// SetIssueType sets field value
func (o *NotificationJiraEndpointConfigurationPut) SetIssueType(v string) {
	o.IssueType = v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *NotificationJiraEndpointConfigurationPut) GetPriority() string {
	if o == nil || IsNil(o.Priority) {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationJiraEndpointConfigurationPut) GetPriorityOk() (*string, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *NotificationJiraEndpointConfigurationPut) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *NotificationJiraEndpointConfigurationPut) SetPriority(v string) {
	o.Priority = &v
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *NotificationJiraEndpointConfigurationPut) GetAssignee() string {
	if o == nil || IsNil(o.Assignee) {
		var ret string
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationJiraEndpointConfigurationPut) GetAssigneeOk() (*string, bool) {
	if o == nil || IsNil(o.Assignee) {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *NotificationJiraEndpointConfigurationPut) HasAssignee() bool {
	if o != nil && !IsNil(o.Assignee) {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given string and assigns it to the Assignee field.
func (o *NotificationJiraEndpointConfigurationPut) SetAssignee(v string) {
	o.Assignee = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *NotificationJiraEndpointConfigurationPut) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationJiraEndpointConfigurationPut) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *NotificationJiraEndpointConfigurationPut) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *NotificationJiraEndpointConfigurationPut) SetLabels(v []string) {
	o.Labels = v
}

func (o NotificationJiraEndpointConfigurationPut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationJiraEndpointConfigurationPut) ToMap() (map[string]interface{}, error) {
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
	toSerialize["url"] = o.Url
	toSerialize["username"] = o.Username
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	toSerialize["project_key"] = o.ProjectKey
	toSerialize["issue_type"] = o.IssueType
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.Assignee) {
		toSerialize["assignee"] = o.Assignee
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	return toSerialize, nil
}

func (o *NotificationJiraEndpointConfigurationPut) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
		"username",
		"project_key",
		"issue_type",
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

	varNotificationJiraEndpointConfigurationPut := _NotificationJiraEndpointConfigurationPut{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNotificationJiraEndpointConfigurationPut)

	if err != nil {
		return err
	}

	*o = NotificationJiraEndpointConfigurationPut(varNotificationJiraEndpointConfigurationPut)

	return err
}

type NullableNotificationJiraEndpointConfigurationPut struct {
	value *NotificationJiraEndpointConfigurationPut
	isSet bool
}

func (v NullableNotificationJiraEndpointConfigurationPut) Get() *NotificationJiraEndpointConfigurationPut {
	return v.value
}

func (v *NullableNotificationJiraEndpointConfigurationPut) Set(val *NotificationJiraEndpointConfigurationPut) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationJiraEndpointConfigurationPut) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationJiraEndpointConfigurationPut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationJiraEndpointConfigurationPut(val *NotificationJiraEndpointConfigurationPut) *NullableNotificationJiraEndpointConfigurationPut {
	return &NullableNotificationJiraEndpointConfigurationPut{value: val, isSet: true}
}

func (v NullableNotificationJiraEndpointConfigurationPut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationJiraEndpointConfigurationPut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


