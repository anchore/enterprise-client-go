# NotificationGitHubEndpointConfigurationPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | The instance identifier for the configuration | [optional] 
**Description** | Pointer to **string** | User friendly name or description for the configuration | [optional] 
**VerifyTls** | Pointer to **bool** | Verify the cert if using tls for connecting externally. Defaults to true if not specified | [optional] 
**CreatedAt** | Pointer to **time.Time** | Timestamp for last modification to the record | [optional] 
**LastUpdated** | Pointer to **time.Time** | Timestamp for last modification to the record | [optional] 
**Url** | Pointer to **string** | Github API endpoint, defaults to https://api.github.com if not specified | [optional] 
**Username** | **string** | GitHub username for creating issues | 
**AccessToken** | **string** | Personal access token for the GitHub account | 
**Owner** | **string** | Owner of the repository to create issues against | 
**Repository** | **string** | Name of the repository to create issues against | 
**Milestone** | Pointer to **int64** | Number of the milestone to associate with the issue | [optional] 
**Labels** | Pointer to **[]string** | List of labels to associate with the issue | [optional] 
**Assignees** | Pointer to **[]string** | List of user logins to assign to the issue. | [optional] 

## Methods

### NewNotificationGitHubEndpointConfigurationPost

`func NewNotificationGitHubEndpointConfigurationPost(username string, accessToken string, owner string, repository string, ) *NotificationGitHubEndpointConfigurationPost`

NewNotificationGitHubEndpointConfigurationPost instantiates a new NotificationGitHubEndpointConfigurationPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationGitHubEndpointConfigurationPostWithDefaults

`func NewNotificationGitHubEndpointConfigurationPostWithDefaults() *NotificationGitHubEndpointConfigurationPost`

NewNotificationGitHubEndpointConfigurationPostWithDefaults instantiates a new NotificationGitHubEndpointConfigurationPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *NotificationGitHubEndpointConfigurationPost) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *NotificationGitHubEndpointConfigurationPost) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *NotificationGitHubEndpointConfigurationPost) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *NotificationGitHubEndpointConfigurationPost) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetDescription

`func (o *NotificationGitHubEndpointConfigurationPost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NotificationGitHubEndpointConfigurationPost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NotificationGitHubEndpointConfigurationPost) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NotificationGitHubEndpointConfigurationPost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVerifyTls

`func (o *NotificationGitHubEndpointConfigurationPost) GetVerifyTls() bool`

GetVerifyTls returns the VerifyTls field if non-nil, zero value otherwise.

### GetVerifyTlsOk

`func (o *NotificationGitHubEndpointConfigurationPost) GetVerifyTlsOk() (*bool, bool)`

GetVerifyTlsOk returns a tuple with the VerifyTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyTls

`func (o *NotificationGitHubEndpointConfigurationPost) SetVerifyTls(v bool)`

SetVerifyTls sets VerifyTls field to given value.

### HasVerifyTls

`func (o *NotificationGitHubEndpointConfigurationPost) HasVerifyTls() bool`

HasVerifyTls returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NotificationGitHubEndpointConfigurationPost) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationGitHubEndpointConfigurationPost) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationGitHubEndpointConfigurationPost) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NotificationGitHubEndpointConfigurationPost) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdated

`func (o *NotificationGitHubEndpointConfigurationPost) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *NotificationGitHubEndpointConfigurationPost) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *NotificationGitHubEndpointConfigurationPost) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *NotificationGitHubEndpointConfigurationPost) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetUrl

`func (o *NotificationGitHubEndpointConfigurationPost) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NotificationGitHubEndpointConfigurationPost) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NotificationGitHubEndpointConfigurationPost) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NotificationGitHubEndpointConfigurationPost) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsername

`func (o *NotificationGitHubEndpointConfigurationPost) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *NotificationGitHubEndpointConfigurationPost) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *NotificationGitHubEndpointConfigurationPost) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetAccessToken

`func (o *NotificationGitHubEndpointConfigurationPost) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *NotificationGitHubEndpointConfigurationPost) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *NotificationGitHubEndpointConfigurationPost) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetOwner

`func (o *NotificationGitHubEndpointConfigurationPost) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *NotificationGitHubEndpointConfigurationPost) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *NotificationGitHubEndpointConfigurationPost) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetRepository

`func (o *NotificationGitHubEndpointConfigurationPost) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *NotificationGitHubEndpointConfigurationPost) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *NotificationGitHubEndpointConfigurationPost) SetRepository(v string)`

SetRepository sets Repository field to given value.


### GetMilestone

`func (o *NotificationGitHubEndpointConfigurationPost) GetMilestone() int64`

GetMilestone returns the Milestone field if non-nil, zero value otherwise.

### GetMilestoneOk

`func (o *NotificationGitHubEndpointConfigurationPost) GetMilestoneOk() (*int64, bool)`

GetMilestoneOk returns a tuple with the Milestone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestone

`func (o *NotificationGitHubEndpointConfigurationPost) SetMilestone(v int64)`

SetMilestone sets Milestone field to given value.

### HasMilestone

`func (o *NotificationGitHubEndpointConfigurationPost) HasMilestone() bool`

HasMilestone returns a boolean if a field has been set.

### GetLabels

`func (o *NotificationGitHubEndpointConfigurationPost) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *NotificationGitHubEndpointConfigurationPost) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *NotificationGitHubEndpointConfigurationPost) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *NotificationGitHubEndpointConfigurationPost) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAssignees

`func (o *NotificationGitHubEndpointConfigurationPost) GetAssignees() []string`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *NotificationGitHubEndpointConfigurationPost) GetAssigneesOk() (*[]string, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *NotificationGitHubEndpointConfigurationPost) SetAssignees(v []string)`

SetAssignees sets Assignees field to given value.

### HasAssignees

`func (o *NotificationGitHubEndpointConfigurationPost) HasAssignees() bool`

HasAssignees returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


