# NotificationGitHubEndpointConfigurationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | Github API endpoint, defaults to https://api.github.com if not specified | [optional] 
**Username** | **string** | GitHub username for creating issues | 
**AccessToken** | **string** | Personal access token for the GitHub account | 
**Owner** | **string** | Owner of the repository to create issues against | 
**Repository** | **string** | Name of the repository to create issues against | 
**Milestone** | Pointer to **int32** | Number of the milestone to associate with the issue | [optional] 
**Labels** | Pointer to **[]string** | List of labels to associate with the issue | [optional] 
**Assignees** | Pointer to **[]string** | List of user logins to assign to the issue. | [optional] 

## Methods

### NewNotificationGitHubEndpointConfigurationAllOf

`func NewNotificationGitHubEndpointConfigurationAllOf(username string, accessToken string, owner string, repository string, ) *NotificationGitHubEndpointConfigurationAllOf`

NewNotificationGitHubEndpointConfigurationAllOf instantiates a new NotificationGitHubEndpointConfigurationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationGitHubEndpointConfigurationAllOfWithDefaults

`func NewNotificationGitHubEndpointConfigurationAllOfWithDefaults() *NotificationGitHubEndpointConfigurationAllOf`

NewNotificationGitHubEndpointConfigurationAllOfWithDefaults instantiates a new NotificationGitHubEndpointConfigurationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *NotificationGitHubEndpointConfigurationAllOf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NotificationGitHubEndpointConfigurationAllOf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NotificationGitHubEndpointConfigurationAllOf) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NotificationGitHubEndpointConfigurationAllOf) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsername

`func (o *NotificationGitHubEndpointConfigurationAllOf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *NotificationGitHubEndpointConfigurationAllOf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *NotificationGitHubEndpointConfigurationAllOf) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetAccessToken

`func (o *NotificationGitHubEndpointConfigurationAllOf) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *NotificationGitHubEndpointConfigurationAllOf) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *NotificationGitHubEndpointConfigurationAllOf) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetOwner

`func (o *NotificationGitHubEndpointConfigurationAllOf) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *NotificationGitHubEndpointConfigurationAllOf) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *NotificationGitHubEndpointConfigurationAllOf) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetRepository

`func (o *NotificationGitHubEndpointConfigurationAllOf) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *NotificationGitHubEndpointConfigurationAllOf) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *NotificationGitHubEndpointConfigurationAllOf) SetRepository(v string)`

SetRepository sets Repository field to given value.


### GetMilestone

`func (o *NotificationGitHubEndpointConfigurationAllOf) GetMilestone() int32`

GetMilestone returns the Milestone field if non-nil, zero value otherwise.

### GetMilestoneOk

`func (o *NotificationGitHubEndpointConfigurationAllOf) GetMilestoneOk() (*int32, bool)`

GetMilestoneOk returns a tuple with the Milestone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestone

`func (o *NotificationGitHubEndpointConfigurationAllOf) SetMilestone(v int32)`

SetMilestone sets Milestone field to given value.

### HasMilestone

`func (o *NotificationGitHubEndpointConfigurationAllOf) HasMilestone() bool`

HasMilestone returns a boolean if a field has been set.

### GetLabels

`func (o *NotificationGitHubEndpointConfigurationAllOf) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *NotificationGitHubEndpointConfigurationAllOf) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *NotificationGitHubEndpointConfigurationAllOf) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *NotificationGitHubEndpointConfigurationAllOf) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAssignees

`func (o *NotificationGitHubEndpointConfigurationAllOf) GetAssignees() []string`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *NotificationGitHubEndpointConfigurationAllOf) GetAssigneesOk() (*[]string, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *NotificationGitHubEndpointConfigurationAllOf) SetAssignees(v []string)`

SetAssignees sets Assignees field to given value.

### HasAssignees

`func (o *NotificationGitHubEndpointConfigurationAllOf) HasAssignees() bool`

HasAssignees returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


