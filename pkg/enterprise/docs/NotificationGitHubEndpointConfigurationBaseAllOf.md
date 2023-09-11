# NotificationGitHubEndpointConfigurationBaseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | Github API endpoint, defaults to https://api.github.com if not specified | [optional] 
**Username** | Pointer to **string** | GitHub username for creating issues | [optional] 
**AccessToken** | Pointer to **string** | Personal access token for the GitHub account | [optional] 
**Owner** | Pointer to **string** | Owner of the repository to create issues against | [optional] 
**Repository** | Pointer to **string** | Name of the repository to create issues against | [optional] 
**Milestone** | Pointer to **int32** | Number of the milestone to associate with the issue | [optional] 
**Labels** | Pointer to **[]string** | List of labels to associate with the issue | [optional] 
**Assignees** | Pointer to **[]string** | List of user logins to assign to the issue. | [optional] 

## Methods

### NewNotificationGitHubEndpointConfigurationBaseAllOf

`func NewNotificationGitHubEndpointConfigurationBaseAllOf() *NotificationGitHubEndpointConfigurationBaseAllOf`

NewNotificationGitHubEndpointConfigurationBaseAllOf instantiates a new NotificationGitHubEndpointConfigurationBaseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationGitHubEndpointConfigurationBaseAllOfWithDefaults

`func NewNotificationGitHubEndpointConfigurationBaseAllOfWithDefaults() *NotificationGitHubEndpointConfigurationBaseAllOf`

NewNotificationGitHubEndpointConfigurationBaseAllOfWithDefaults instantiates a new NotificationGitHubEndpointConfigurationBaseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *NotificationGitHubEndpointConfigurationBaseAllOf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NotificationGitHubEndpointConfigurationBaseAllOf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NotificationGitHubEndpointConfigurationBaseAllOf) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NotificationGitHubEndpointConfigurationBaseAllOf) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsername

`func (o *NotificationGitHubEndpointConfigurationBaseAllOf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *NotificationGitHubEndpointConfigurationBaseAllOf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *NotificationGitHubEndpointConfigurationBaseAllOf) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *NotificationGitHubEndpointConfigurationBaseAllOf) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetAccessToken

`func (o *NotificationGitHubEndpointConfigurationBaseAllOf) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *NotificationGitHubEndpointConfigurationBaseAllOf) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *NotificationGitHubEndpointConfigurationBaseAllOf) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *NotificationGitHubEndpointConfigurationBaseAllOf) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetOwner

`func (o *NotificationGitHubEndpointConfigurationBaseAllOf) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *NotificationGitHubEndpointConfigurationBaseAllOf) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *NotificationGitHubEndpointConfigurationBaseAllOf) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *NotificationGitHubEndpointConfigurationBaseAllOf) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetRepository

`func (o *NotificationGitHubEndpointConfigurationBaseAllOf) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *NotificationGitHubEndpointConfigurationBaseAllOf) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *NotificationGitHubEndpointConfigurationBaseAllOf) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *NotificationGitHubEndpointConfigurationBaseAllOf) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetMilestone

`func (o *NotificationGitHubEndpointConfigurationBaseAllOf) GetMilestone() int32`

GetMilestone returns the Milestone field if non-nil, zero value otherwise.

### GetMilestoneOk

`func (o *NotificationGitHubEndpointConfigurationBaseAllOf) GetMilestoneOk() (*int32, bool)`

GetMilestoneOk returns a tuple with the Milestone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestone

`func (o *NotificationGitHubEndpointConfigurationBaseAllOf) SetMilestone(v int32)`

SetMilestone sets Milestone field to given value.

### HasMilestone

`func (o *NotificationGitHubEndpointConfigurationBaseAllOf) HasMilestone() bool`

HasMilestone returns a boolean if a field has been set.

### GetLabels

`func (o *NotificationGitHubEndpointConfigurationBaseAllOf) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *NotificationGitHubEndpointConfigurationBaseAllOf) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *NotificationGitHubEndpointConfigurationBaseAllOf) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *NotificationGitHubEndpointConfigurationBaseAllOf) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAssignees

`func (o *NotificationGitHubEndpointConfigurationBaseAllOf) GetAssignees() []string`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *NotificationGitHubEndpointConfigurationBaseAllOf) GetAssigneesOk() (*[]string, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *NotificationGitHubEndpointConfigurationBaseAllOf) SetAssignees(v []string)`

SetAssignees sets Assignees field to given value.

### HasAssignees

`func (o *NotificationGitHubEndpointConfigurationBaseAllOf) HasAssignees() bool`

HasAssignees returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


