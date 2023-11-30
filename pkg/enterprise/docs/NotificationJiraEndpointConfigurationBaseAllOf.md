# NotificationJiraEndpointConfigurationBaseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | Jira endpoint URL including host and port, should begin with &#39;http://&#39; or &#39;https://&#39; | [optional] 
**Username** | Pointer to **string** | Jira username for creating issues | [optional] 
**Password** | Pointer to **string** | Jira access token for creating issues | [optional] 
**ProjectKey** | Pointer to **string** | Key of the Jira project for creating issues | [optional] 
**IssueType** | Pointer to **string** | Type associated with the issue | [optional] 
**Priority** | Pointer to **string** | Priority assigned to the issue | [optional] 
**Assignee** | Pointer to **string** | Jira user to associate with the issue | [optional] 
**Labels** | Pointer to **[]string** | List of labels to associate with the issue | [optional] 

## Methods

### NewNotificationJiraEndpointConfigurationBaseAllOf

`func NewNotificationJiraEndpointConfigurationBaseAllOf() *NotificationJiraEndpointConfigurationBaseAllOf`

NewNotificationJiraEndpointConfigurationBaseAllOf instantiates a new NotificationJiraEndpointConfigurationBaseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationJiraEndpointConfigurationBaseAllOfWithDefaults

`func NewNotificationJiraEndpointConfigurationBaseAllOfWithDefaults() *NotificationJiraEndpointConfigurationBaseAllOf`

NewNotificationJiraEndpointConfigurationBaseAllOfWithDefaults instantiates a new NotificationJiraEndpointConfigurationBaseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *NotificationJiraEndpointConfigurationBaseAllOf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NotificationJiraEndpointConfigurationBaseAllOf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NotificationJiraEndpointConfigurationBaseAllOf) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NotificationJiraEndpointConfigurationBaseAllOf) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsername

`func (o *NotificationJiraEndpointConfigurationBaseAllOf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *NotificationJiraEndpointConfigurationBaseAllOf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *NotificationJiraEndpointConfigurationBaseAllOf) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *NotificationJiraEndpointConfigurationBaseAllOf) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *NotificationJiraEndpointConfigurationBaseAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *NotificationJiraEndpointConfigurationBaseAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *NotificationJiraEndpointConfigurationBaseAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *NotificationJiraEndpointConfigurationBaseAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetProjectKey

`func (o *NotificationJiraEndpointConfigurationBaseAllOf) GetProjectKey() string`

GetProjectKey returns the ProjectKey field if non-nil, zero value otherwise.

### GetProjectKeyOk

`func (o *NotificationJiraEndpointConfigurationBaseAllOf) GetProjectKeyOk() (*string, bool)`

GetProjectKeyOk returns a tuple with the ProjectKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKey

`func (o *NotificationJiraEndpointConfigurationBaseAllOf) SetProjectKey(v string)`

SetProjectKey sets ProjectKey field to given value.

### HasProjectKey

`func (o *NotificationJiraEndpointConfigurationBaseAllOf) HasProjectKey() bool`

HasProjectKey returns a boolean if a field has been set.

### GetIssueType

`func (o *NotificationJiraEndpointConfigurationBaseAllOf) GetIssueType() string`

GetIssueType returns the IssueType field if non-nil, zero value otherwise.

### GetIssueTypeOk

`func (o *NotificationJiraEndpointConfigurationBaseAllOf) GetIssueTypeOk() (*string, bool)`

GetIssueTypeOk returns a tuple with the IssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueType

`func (o *NotificationJiraEndpointConfigurationBaseAllOf) SetIssueType(v string)`

SetIssueType sets IssueType field to given value.

### HasIssueType

`func (o *NotificationJiraEndpointConfigurationBaseAllOf) HasIssueType() bool`

HasIssueType returns a boolean if a field has been set.

### GetPriority

`func (o *NotificationJiraEndpointConfigurationBaseAllOf) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *NotificationJiraEndpointConfigurationBaseAllOf) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *NotificationJiraEndpointConfigurationBaseAllOf) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *NotificationJiraEndpointConfigurationBaseAllOf) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetAssignee

`func (o *NotificationJiraEndpointConfigurationBaseAllOf) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *NotificationJiraEndpointConfigurationBaseAllOf) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *NotificationJiraEndpointConfigurationBaseAllOf) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *NotificationJiraEndpointConfigurationBaseAllOf) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetLabels

`func (o *NotificationJiraEndpointConfigurationBaseAllOf) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *NotificationJiraEndpointConfigurationBaseAllOf) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *NotificationJiraEndpointConfigurationBaseAllOf) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *NotificationJiraEndpointConfigurationBaseAllOf) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


