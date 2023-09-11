# NotificationJiraEndpointConfigurationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | Jira endpoint URL including host and port, should begin with &#39;http://&#39; or &#39;https://&#39; | 
**Username** | **string** | Jira username for creating issues | 
**Password** | **string** | Jira password for creating issues | 
**ProjectKey** | **string** | Key of the Jira project for creating issues | 
**IssueType** | **string** | Type associated with the issue | 
**Priority** | Pointer to **string** | Priority assigned to the issue | [optional] 
**Assignee** | Pointer to **string** | Jira user to associate with the issue | [optional] 
**Labels** | Pointer to **[]string** | List of labels to associate with the issue | [optional] 

## Methods

### NewNotificationJiraEndpointConfigurationAllOf

`func NewNotificationJiraEndpointConfigurationAllOf(url string, username string, password string, projectKey string, issueType string, ) *NotificationJiraEndpointConfigurationAllOf`

NewNotificationJiraEndpointConfigurationAllOf instantiates a new NotificationJiraEndpointConfigurationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationJiraEndpointConfigurationAllOfWithDefaults

`func NewNotificationJiraEndpointConfigurationAllOfWithDefaults() *NotificationJiraEndpointConfigurationAllOf`

NewNotificationJiraEndpointConfigurationAllOfWithDefaults instantiates a new NotificationJiraEndpointConfigurationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *NotificationJiraEndpointConfigurationAllOf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NotificationJiraEndpointConfigurationAllOf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NotificationJiraEndpointConfigurationAllOf) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUsername

`func (o *NotificationJiraEndpointConfigurationAllOf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *NotificationJiraEndpointConfigurationAllOf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *NotificationJiraEndpointConfigurationAllOf) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *NotificationJiraEndpointConfigurationAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *NotificationJiraEndpointConfigurationAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *NotificationJiraEndpointConfigurationAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetProjectKey

`func (o *NotificationJiraEndpointConfigurationAllOf) GetProjectKey() string`

GetProjectKey returns the ProjectKey field if non-nil, zero value otherwise.

### GetProjectKeyOk

`func (o *NotificationJiraEndpointConfigurationAllOf) GetProjectKeyOk() (*string, bool)`

GetProjectKeyOk returns a tuple with the ProjectKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKey

`func (o *NotificationJiraEndpointConfigurationAllOf) SetProjectKey(v string)`

SetProjectKey sets ProjectKey field to given value.


### GetIssueType

`func (o *NotificationJiraEndpointConfigurationAllOf) GetIssueType() string`

GetIssueType returns the IssueType field if non-nil, zero value otherwise.

### GetIssueTypeOk

`func (o *NotificationJiraEndpointConfigurationAllOf) GetIssueTypeOk() (*string, bool)`

GetIssueTypeOk returns a tuple with the IssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueType

`func (o *NotificationJiraEndpointConfigurationAllOf) SetIssueType(v string)`

SetIssueType sets IssueType field to given value.


### GetPriority

`func (o *NotificationJiraEndpointConfigurationAllOf) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *NotificationJiraEndpointConfigurationAllOf) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *NotificationJiraEndpointConfigurationAllOf) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *NotificationJiraEndpointConfigurationAllOf) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetAssignee

`func (o *NotificationJiraEndpointConfigurationAllOf) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *NotificationJiraEndpointConfigurationAllOf) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *NotificationJiraEndpointConfigurationAllOf) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *NotificationJiraEndpointConfigurationAllOf) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetLabels

`func (o *NotificationJiraEndpointConfigurationAllOf) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *NotificationJiraEndpointConfigurationAllOf) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *NotificationJiraEndpointConfigurationAllOf) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *NotificationJiraEndpointConfigurationAllOf) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


