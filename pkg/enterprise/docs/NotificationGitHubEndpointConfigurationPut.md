# NotificationGitHubEndpointConfigurationPut

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
**AccessToken** | Pointer to **string** | Personal access token for the GitHub account | [optional] 
**Owner** | **string** | Owner of the repository to create issues against | 
**Repository** | **string** | Name of the repository to create issues against | 
**Milestone** | Pointer to **int32** | Number of the milestone to associate with the issue | [optional] 
**Labels** | Pointer to **[]string** | List of labels to associate with the issue | [optional] 
**Assignees** | Pointer to **[]string** | List of user logins to assign to the issue. | [optional] 

## Methods

### NewNotificationGitHubEndpointConfigurationPut

`func NewNotificationGitHubEndpointConfigurationPut(username string, owner string, repository string, ) *NotificationGitHubEndpointConfigurationPut`

NewNotificationGitHubEndpointConfigurationPut instantiates a new NotificationGitHubEndpointConfigurationPut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationGitHubEndpointConfigurationPutWithDefaults

`func NewNotificationGitHubEndpointConfigurationPutWithDefaults() *NotificationGitHubEndpointConfigurationPut`

NewNotificationGitHubEndpointConfigurationPutWithDefaults instantiates a new NotificationGitHubEndpointConfigurationPut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *NotificationGitHubEndpointConfigurationPut) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *NotificationGitHubEndpointConfigurationPut) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *NotificationGitHubEndpointConfigurationPut) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *NotificationGitHubEndpointConfigurationPut) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetDescription

`func (o *NotificationGitHubEndpointConfigurationPut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NotificationGitHubEndpointConfigurationPut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NotificationGitHubEndpointConfigurationPut) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NotificationGitHubEndpointConfigurationPut) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVerifyTls

`func (o *NotificationGitHubEndpointConfigurationPut) GetVerifyTls() bool`

GetVerifyTls returns the VerifyTls field if non-nil, zero value otherwise.

### GetVerifyTlsOk

`func (o *NotificationGitHubEndpointConfigurationPut) GetVerifyTlsOk() (*bool, bool)`

GetVerifyTlsOk returns a tuple with the VerifyTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyTls

`func (o *NotificationGitHubEndpointConfigurationPut) SetVerifyTls(v bool)`

SetVerifyTls sets VerifyTls field to given value.

### HasVerifyTls

`func (o *NotificationGitHubEndpointConfigurationPut) HasVerifyTls() bool`

HasVerifyTls returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NotificationGitHubEndpointConfigurationPut) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationGitHubEndpointConfigurationPut) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationGitHubEndpointConfigurationPut) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NotificationGitHubEndpointConfigurationPut) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdated

`func (o *NotificationGitHubEndpointConfigurationPut) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *NotificationGitHubEndpointConfigurationPut) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *NotificationGitHubEndpointConfigurationPut) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *NotificationGitHubEndpointConfigurationPut) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetUrl

`func (o *NotificationGitHubEndpointConfigurationPut) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NotificationGitHubEndpointConfigurationPut) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NotificationGitHubEndpointConfigurationPut) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NotificationGitHubEndpointConfigurationPut) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsername

`func (o *NotificationGitHubEndpointConfigurationPut) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *NotificationGitHubEndpointConfigurationPut) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *NotificationGitHubEndpointConfigurationPut) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetAccessToken

`func (o *NotificationGitHubEndpointConfigurationPut) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *NotificationGitHubEndpointConfigurationPut) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *NotificationGitHubEndpointConfigurationPut) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *NotificationGitHubEndpointConfigurationPut) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetOwner

`func (o *NotificationGitHubEndpointConfigurationPut) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *NotificationGitHubEndpointConfigurationPut) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *NotificationGitHubEndpointConfigurationPut) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetRepository

`func (o *NotificationGitHubEndpointConfigurationPut) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *NotificationGitHubEndpointConfigurationPut) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *NotificationGitHubEndpointConfigurationPut) SetRepository(v string)`

SetRepository sets Repository field to given value.


### GetMilestone

`func (o *NotificationGitHubEndpointConfigurationPut) GetMilestone() int32`

GetMilestone returns the Milestone field if non-nil, zero value otherwise.

### GetMilestoneOk

`func (o *NotificationGitHubEndpointConfigurationPut) GetMilestoneOk() (*int32, bool)`

GetMilestoneOk returns a tuple with the Milestone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestone

`func (o *NotificationGitHubEndpointConfigurationPut) SetMilestone(v int32)`

SetMilestone sets Milestone field to given value.

### HasMilestone

`func (o *NotificationGitHubEndpointConfigurationPut) HasMilestone() bool`

HasMilestone returns a boolean if a field has been set.

### GetLabels

`func (o *NotificationGitHubEndpointConfigurationPut) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *NotificationGitHubEndpointConfigurationPut) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *NotificationGitHubEndpointConfigurationPut) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *NotificationGitHubEndpointConfigurationPut) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAssignees

`func (o *NotificationGitHubEndpointConfigurationPut) GetAssignees() []string`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *NotificationGitHubEndpointConfigurationPut) GetAssigneesOk() (*[]string, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *NotificationGitHubEndpointConfigurationPut) SetAssignees(v []string)`

SetAssignees sets Assignees field to given value.

### HasAssignees

`func (o *NotificationGitHubEndpointConfigurationPut) HasAssignees() bool`

HasAssignees returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


