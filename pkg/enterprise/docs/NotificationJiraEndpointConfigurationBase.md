# NotificationJiraEndpointConfigurationBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | The instance identifier for the configuration | [optional] 
**Description** | Pointer to **string** | User friendly name or description for the configuration | [optional] 
**VerifyTls** | Pointer to **bool** | Verify the cert if using tls for connecting externally. Defaults to true if not specified | [optional] 
**CreatedAt** | Pointer to **time.Time** | Timestamp for last modification to the record | [optional] 
**LastUpdated** | Pointer to **time.Time** | Timestamp for last modification to the record | [optional] 
**Url** | Pointer to **string** | Jira endpoint URL including host and port, should begin with &#39;http://&#39; or &#39;https://&#39; | [optional] 
**Username** | Pointer to **string** | Jira username for creating issues | [optional] 
**Password** | Pointer to **string** | Jira access token for creating issues | [optional] 
**ProjectKey** | Pointer to **string** | Key of the Jira project for creating issues | [optional] 
**IssueType** | Pointer to **string** | Type associated with the issue | [optional] 
**Priority** | Pointer to **string** | Priority assigned to the issue | [optional] 
**Assignee** | Pointer to **string** | Jira user to associate with the issue | [optional] 
**Labels** | Pointer to **[]string** | List of labels to associate with the issue | [optional] 

## Methods

### NewNotificationJiraEndpointConfigurationBase

`func NewNotificationJiraEndpointConfigurationBase() *NotificationJiraEndpointConfigurationBase`

NewNotificationJiraEndpointConfigurationBase instantiates a new NotificationJiraEndpointConfigurationBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationJiraEndpointConfigurationBaseWithDefaults

`func NewNotificationJiraEndpointConfigurationBaseWithDefaults() *NotificationJiraEndpointConfigurationBase`

NewNotificationJiraEndpointConfigurationBaseWithDefaults instantiates a new NotificationJiraEndpointConfigurationBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *NotificationJiraEndpointConfigurationBase) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *NotificationJiraEndpointConfigurationBase) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *NotificationJiraEndpointConfigurationBase) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *NotificationJiraEndpointConfigurationBase) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetDescription

`func (o *NotificationJiraEndpointConfigurationBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NotificationJiraEndpointConfigurationBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NotificationJiraEndpointConfigurationBase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NotificationJiraEndpointConfigurationBase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVerifyTls

`func (o *NotificationJiraEndpointConfigurationBase) GetVerifyTls() bool`

GetVerifyTls returns the VerifyTls field if non-nil, zero value otherwise.

### GetVerifyTlsOk

`func (o *NotificationJiraEndpointConfigurationBase) GetVerifyTlsOk() (*bool, bool)`

GetVerifyTlsOk returns a tuple with the VerifyTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyTls

`func (o *NotificationJiraEndpointConfigurationBase) SetVerifyTls(v bool)`

SetVerifyTls sets VerifyTls field to given value.

### HasVerifyTls

`func (o *NotificationJiraEndpointConfigurationBase) HasVerifyTls() bool`

HasVerifyTls returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NotificationJiraEndpointConfigurationBase) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationJiraEndpointConfigurationBase) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationJiraEndpointConfigurationBase) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NotificationJiraEndpointConfigurationBase) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdated

`func (o *NotificationJiraEndpointConfigurationBase) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *NotificationJiraEndpointConfigurationBase) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *NotificationJiraEndpointConfigurationBase) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *NotificationJiraEndpointConfigurationBase) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetUrl

`func (o *NotificationJiraEndpointConfigurationBase) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NotificationJiraEndpointConfigurationBase) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NotificationJiraEndpointConfigurationBase) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NotificationJiraEndpointConfigurationBase) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsername

`func (o *NotificationJiraEndpointConfigurationBase) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *NotificationJiraEndpointConfigurationBase) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *NotificationJiraEndpointConfigurationBase) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *NotificationJiraEndpointConfigurationBase) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *NotificationJiraEndpointConfigurationBase) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *NotificationJiraEndpointConfigurationBase) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *NotificationJiraEndpointConfigurationBase) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *NotificationJiraEndpointConfigurationBase) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetProjectKey

`func (o *NotificationJiraEndpointConfigurationBase) GetProjectKey() string`

GetProjectKey returns the ProjectKey field if non-nil, zero value otherwise.

### GetProjectKeyOk

`func (o *NotificationJiraEndpointConfigurationBase) GetProjectKeyOk() (*string, bool)`

GetProjectKeyOk returns a tuple with the ProjectKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKey

`func (o *NotificationJiraEndpointConfigurationBase) SetProjectKey(v string)`

SetProjectKey sets ProjectKey field to given value.

### HasProjectKey

`func (o *NotificationJiraEndpointConfigurationBase) HasProjectKey() bool`

HasProjectKey returns a boolean if a field has been set.

### GetIssueType

`func (o *NotificationJiraEndpointConfigurationBase) GetIssueType() string`

GetIssueType returns the IssueType field if non-nil, zero value otherwise.

### GetIssueTypeOk

`func (o *NotificationJiraEndpointConfigurationBase) GetIssueTypeOk() (*string, bool)`

GetIssueTypeOk returns a tuple with the IssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueType

`func (o *NotificationJiraEndpointConfigurationBase) SetIssueType(v string)`

SetIssueType sets IssueType field to given value.

### HasIssueType

`func (o *NotificationJiraEndpointConfigurationBase) HasIssueType() bool`

HasIssueType returns a boolean if a field has been set.

### GetPriority

`func (o *NotificationJiraEndpointConfigurationBase) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *NotificationJiraEndpointConfigurationBase) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *NotificationJiraEndpointConfigurationBase) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *NotificationJiraEndpointConfigurationBase) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetAssignee

`func (o *NotificationJiraEndpointConfigurationBase) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *NotificationJiraEndpointConfigurationBase) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *NotificationJiraEndpointConfigurationBase) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *NotificationJiraEndpointConfigurationBase) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetLabels

`func (o *NotificationJiraEndpointConfigurationBase) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *NotificationJiraEndpointConfigurationBase) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *NotificationJiraEndpointConfigurationBase) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *NotificationJiraEndpointConfigurationBase) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


