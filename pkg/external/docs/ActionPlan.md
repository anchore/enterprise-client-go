# ActionPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**ImageTag** | Pointer to **string** |  | [optional] 
**ImageDigest** | Pointer to **string** |  | [optional] 
**PolicyId** | Pointer to **string** |  | [optional] 
**Resolutions** | Pointer to [**[]ActionPlanResolution**](ActionPlanResolution.md) |  | [optional] 
**Endpoint** | Pointer to **string** |  | [optional] 
**ConfigurationId** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewActionPlan

`func NewActionPlan() *ActionPlan`

NewActionPlan instantiates a new ActionPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionPlanWithDefaults

`func NewActionPlanWithDefaults() *ActionPlan`

NewActionPlanWithDefaults instantiates a new ActionPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ActionPlan) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActionPlan) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActionPlan) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ActionPlan) HasType() bool`

HasType returns a boolean if a field has been set.

### GetImageTag

`func (o *ActionPlan) GetImageTag() string`

GetImageTag returns the ImageTag field if non-nil, zero value otherwise.

### GetImageTagOk

`func (o *ActionPlan) GetImageTagOk() (*string, bool)`

GetImageTagOk returns a tuple with the ImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTag

`func (o *ActionPlan) SetImageTag(v string)`

SetImageTag sets ImageTag field to given value.

### HasImageTag

`func (o *ActionPlan) HasImageTag() bool`

HasImageTag returns a boolean if a field has been set.

### GetImageDigest

`func (o *ActionPlan) GetImageDigest() string`

GetImageDigest returns the ImageDigest field if non-nil, zero value otherwise.

### GetImageDigestOk

`func (o *ActionPlan) GetImageDigestOk() (*string, bool)`

GetImageDigestOk returns a tuple with the ImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDigest

`func (o *ActionPlan) SetImageDigest(v string)`

SetImageDigest sets ImageDigest field to given value.

### HasImageDigest

`func (o *ActionPlan) HasImageDigest() bool`

HasImageDigest returns a boolean if a field has been set.

### GetPolicyId

`func (o *ActionPlan) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *ActionPlan) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *ActionPlan) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *ActionPlan) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetResolutions

`func (o *ActionPlan) GetResolutions() []ActionPlanResolution`

GetResolutions returns the Resolutions field if non-nil, zero value otherwise.

### GetResolutionsOk

`func (o *ActionPlan) GetResolutionsOk() (*[]ActionPlanResolution, bool)`

GetResolutionsOk returns a tuple with the Resolutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutions

`func (o *ActionPlan) SetResolutions(v []ActionPlanResolution)`

SetResolutions sets Resolutions field to given value.

### HasResolutions

`func (o *ActionPlan) HasResolutions() bool`

HasResolutions returns a boolean if a field has been set.

### GetEndpoint

`func (o *ActionPlan) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ActionPlan) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ActionPlan) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *ActionPlan) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetConfigurationId

`func (o *ActionPlan) GetConfigurationId() string`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *ActionPlan) GetConfigurationIdOk() (*string, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *ActionPlan) SetConfigurationId(v string)`

SetConfigurationId sets ConfigurationId field to given value.

### HasConfigurationId

`func (o *ActionPlan) HasConfigurationId() bool`

HasConfigurationId returns a boolean if a field has been set.

### GetSubject

`func (o *ActionPlan) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ActionPlan) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ActionPlan) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ActionPlan) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetMessage

`func (o *ActionPlan) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ActionPlan) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ActionPlan) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ActionPlan) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetUuid

`func (o *ActionPlan) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ActionPlan) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ActionPlan) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ActionPlan) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ActionPlan) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ActionPlan) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ActionPlan) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ActionPlan) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ActionPlan) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ActionPlan) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ActionPlan) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ActionPlan) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


