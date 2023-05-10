# ECSInventoryContainers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arn** | **string** |  | 
**TaskArn** | Pointer to **string** | When a task is provided, this must reference one of the task&#39;s arn in this payload | [optional] 
**ImageTag** | **string** |  | 
**ImageDigest** | Pointer to **string** |  | [optional] 

## Methods

### NewECSInventoryContainers

`func NewECSInventoryContainers(arn string, imageTag string, ) *ECSInventoryContainers`

NewECSInventoryContainers instantiates a new ECSInventoryContainers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewECSInventoryContainersWithDefaults

`func NewECSInventoryContainersWithDefaults() *ECSInventoryContainers`

NewECSInventoryContainersWithDefaults instantiates a new ECSInventoryContainers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArn

`func (o *ECSInventoryContainers) GetArn() string`

GetArn returns the Arn field if non-nil, zero value otherwise.

### GetArnOk

`func (o *ECSInventoryContainers) GetArnOk() (*string, bool)`

GetArnOk returns a tuple with the Arn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArn

`func (o *ECSInventoryContainers) SetArn(v string)`

SetArn sets Arn field to given value.


### GetTaskArn

`func (o *ECSInventoryContainers) GetTaskArn() string`

GetTaskArn returns the TaskArn field if non-nil, zero value otherwise.

### GetTaskArnOk

`func (o *ECSInventoryContainers) GetTaskArnOk() (*string, bool)`

GetTaskArnOk returns a tuple with the TaskArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskArn

`func (o *ECSInventoryContainers) SetTaskArn(v string)`

SetTaskArn sets TaskArn field to given value.

### HasTaskArn

`func (o *ECSInventoryContainers) HasTaskArn() bool`

HasTaskArn returns a boolean if a field has been set.

### GetImageTag

`func (o *ECSInventoryContainers) GetImageTag() string`

GetImageTag returns the ImageTag field if non-nil, zero value otherwise.

### GetImageTagOk

`func (o *ECSInventoryContainers) GetImageTagOk() (*string, bool)`

GetImageTagOk returns a tuple with the ImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTag

`func (o *ECSInventoryContainers) SetImageTag(v string)`

SetImageTag sets ImageTag field to given value.


### GetImageDigest

`func (o *ECSInventoryContainers) GetImageDigest() string`

GetImageDigest returns the ImageDigest field if non-nil, zero value otherwise.

### GetImageDigestOk

`func (o *ECSInventoryContainers) GetImageDigestOk() (*string, bool)`

GetImageDigestOk returns a tuple with the ImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDigest

`func (o *ECSInventoryContainers) SetImageDigest(v string)`

SetImageDigest sets ImageDigest field to given value.

### HasImageDigest

`func (o *ECSInventoryContainers) HasImageDigest() bool`

HasImageDigest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


