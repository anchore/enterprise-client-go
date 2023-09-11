# NotificationJiraEndpointConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | The instance identifier for the configuration | [optional] 
**Description** | Pointer to **string** | User friendly name or description for the configuration | [optional] 
**VerifyTls** | Pointer to **bool** | Verify the cert if using tls for connecting externally. Defaults to true if not specified | [optional] 
**CreatedAt** | Pointer to **time.Time** | Timestamp for last modification to the record | [optional] 
**LastUpdated** | Pointer to **time.Time** | Timestamp for last modification to the record | [optional] 
**Url** | **string** | Jira endpoint URL including host and port, should begin with &#39;http://&#39; or &#39;https://&#39; | 
**Username** | **string** | Jira username for creating issues | 
**Password** | **string** | Jira password for creating issues | 
**ProjectKey** | **string** | Key of the Jira project for creating issues | 
**IssueType** | **string** | Type associated with the issue | 
**Priority** | Pointer to **string** | Priority assigned to the issue | [optional] 
**Assignee** | Pointer to **string** | Jira user to associate with the issue | [optional] 
**Labels** | Pointer to **[]string** | List of labels to associate with the issue | [optional] 

## Methods

### NewNotificationJiraEndpointConfiguration

`func NewNotificationJiraEndpointConfiguration(url string, username string, password string, projectKey string, issueType string, ) *NotificationJiraEndpointConfiguration`

NewNotificationJiraEndpointConfiguration instantiates a new NotificationJiraEndpointConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationJiraEndpointConfigurationWithDefaults

`func NewNotificationJiraEndpointConfigurationWithDefaults() *NotificationJiraEndpointConfiguration`

NewNotificationJiraEndpointConfigurationWithDefaults instantiates a new NotificationJiraEndpointConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *NotificationJiraEndpointConfiguration) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *NotificationJiraEndpointConfiguration) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *NotificationJiraEndpointConfiguration) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *NotificationJiraEndpointConfiguration) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetDescription

`func (o *NotificationJiraEndpointConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NotificationJiraEndpointConfiguration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NotificationJiraEndpointConfiguration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NotificationJiraEndpointConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVerifyTls

`func (o *NotificationJiraEndpointConfiguration) GetVerifyTls() bool`

GetVerifyTls returns the VerifyTls field if non-nil, zero value otherwise.

### GetVerifyTlsOk

`func (o *NotificationJiraEndpointConfiguration) GetVerifyTlsOk() (*bool, bool)`

GetVerifyTlsOk returns a tuple with the VerifyTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyTls

`func (o *NotificationJiraEndpointConfiguration) SetVerifyTls(v bool)`

SetVerifyTls sets VerifyTls field to given value.

### HasVerifyTls

`func (o *NotificationJiraEndpointConfiguration) HasVerifyTls() bool`

HasVerifyTls returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NotificationJiraEndpointConfiguration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationJiraEndpointConfiguration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationJiraEndpointConfiguration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NotificationJiraEndpointConfiguration) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdated

`func (o *NotificationJiraEndpointConfiguration) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *NotificationJiraEndpointConfiguration) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *NotificationJiraEndpointConfiguration) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *NotificationJiraEndpointConfiguration) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetUrl

`func (o *NotificationJiraEndpointConfiguration) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NotificationJiraEndpointConfiguration) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NotificationJiraEndpointConfiguration) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUsername

`func (o *NotificationJiraEndpointConfiguration) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *NotificationJiraEndpointConfiguration) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *NotificationJiraEndpointConfiguration) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *NotificationJiraEndpointConfiguration) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *NotificationJiraEndpointConfiguration) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *NotificationJiraEndpointConfiguration) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetProjectKey

`func (o *NotificationJiraEndpointConfiguration) GetProjectKey() string`

GetProjectKey returns the ProjectKey field if non-nil, zero value otherwise.

### GetProjectKeyOk

`func (o *NotificationJiraEndpointConfiguration) GetProjectKeyOk() (*string, bool)`

GetProjectKeyOk returns a tuple with the ProjectKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKey

`func (o *NotificationJiraEndpointConfiguration) SetProjectKey(v string)`

SetProjectKey sets ProjectKey field to given value.


### GetIssueType

`func (o *NotificationJiraEndpointConfiguration) GetIssueType() string`

GetIssueType returns the IssueType field if non-nil, zero value otherwise.

### GetIssueTypeOk

`func (o *NotificationJiraEndpointConfiguration) GetIssueTypeOk() (*string, bool)`

GetIssueTypeOk returns a tuple with the IssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueType

`func (o *NotificationJiraEndpointConfiguration) SetIssueType(v string)`

SetIssueType sets IssueType field to given value.


### GetPriority

`func (o *NotificationJiraEndpointConfiguration) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *NotificationJiraEndpointConfiguration) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *NotificationJiraEndpointConfiguration) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *NotificationJiraEndpointConfiguration) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetAssignee

`func (o *NotificationJiraEndpointConfiguration) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *NotificationJiraEndpointConfiguration) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *NotificationJiraEndpointConfiguration) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *NotificationJiraEndpointConfiguration) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetLabels

`func (o *NotificationJiraEndpointConfiguration) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *NotificationJiraEndpointConfiguration) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *NotificationJiraEndpointConfiguration) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *NotificationJiraEndpointConfiguration) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


