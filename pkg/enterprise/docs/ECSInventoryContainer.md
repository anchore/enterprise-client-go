# ECSInventoryContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arn** | **string** |  | 
**TaskArn** | Pointer to **string** |  | [optional] 
**ImageTag** | **string** |  | 
**ImageDigest** | Pointer to **string** |  | [optional] 

## Methods

### NewECSInventoryContainer

`func NewECSInventoryContainer(arn string, imageTag string, ) *ECSInventoryContainer`

NewECSInventoryContainer instantiates a new ECSInventoryContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewECSInventoryContainerWithDefaults

`func NewECSInventoryContainerWithDefaults() *ECSInventoryContainer`

NewECSInventoryContainerWithDefaults instantiates a new ECSInventoryContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArn

`func (o *ECSInventoryContainer) GetArn() string`

GetArn returns the Arn field if non-nil, zero value otherwise.

### GetArnOk

`func (o *ECSInventoryContainer) GetArnOk() (*string, bool)`

GetArnOk returns a tuple with the Arn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArn

`func (o *ECSInventoryContainer) SetArn(v string)`

SetArn sets Arn field to given value.


### GetTaskArn

`func (o *ECSInventoryContainer) GetTaskArn() string`

GetTaskArn returns the TaskArn field if non-nil, zero value otherwise.

### GetTaskArnOk

`func (o *ECSInventoryContainer) GetTaskArnOk() (*string, bool)`

GetTaskArnOk returns a tuple with the TaskArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskArn

`func (o *ECSInventoryContainer) SetTaskArn(v string)`

SetTaskArn sets TaskArn field to given value.

### HasTaskArn

`func (o *ECSInventoryContainer) HasTaskArn() bool`

HasTaskArn returns a boolean if a field has been set.

### GetImageTag

`func (o *ECSInventoryContainer) GetImageTag() string`

GetImageTag returns the ImageTag field if non-nil, zero value otherwise.

### GetImageTagOk

`func (o *ECSInventoryContainer) GetImageTagOk() (*string, bool)`

GetImageTagOk returns a tuple with the ImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTag

`func (o *ECSInventoryContainer) SetImageTag(v string)`

SetImageTag sets ImageTag field to given value.


### GetImageDigest

`func (o *ECSInventoryContainer) GetImageDigest() string`

GetImageDigest returns the ImageDigest field if non-nil, zero value otherwise.

### GetImageDigestOk

`func (o *ECSInventoryContainer) GetImageDigestOk() (*string, bool)`

GetImageDigestOk returns a tuple with the ImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDigest

`func (o *ECSInventoryContainer) SetImageDigest(v string)`

SetImageDigest sets ImageDigest field to given value.

### HasImageDigest

`func (o *ECSInventoryContainer) HasImageDigest() bool`

HasImageDigest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


