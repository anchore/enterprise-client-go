# ECSTasksTasksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arn** | **string** |  | 
**ClusterArn** | **string** |  | 
**ServiceArn** | **string** |  | 
**TaskDefinitionArn** | **string** |  | 
**Tags** | **map[string]string** |  | 
**AccountName** | **string** |  | 

## Methods

### NewECSTasksTasksInner

`func NewECSTasksTasksInner(arn string, clusterArn string, serviceArn string, taskDefinitionArn string, tags map[string]string, accountName string, ) *ECSTasksTasksInner`

NewECSTasksTasksInner instantiates a new ECSTasksTasksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewECSTasksTasksInnerWithDefaults

`func NewECSTasksTasksInnerWithDefaults() *ECSTasksTasksInner`

NewECSTasksTasksInnerWithDefaults instantiates a new ECSTasksTasksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArn

`func (o *ECSTasksTasksInner) GetArn() string`

GetArn returns the Arn field if non-nil, zero value otherwise.

### GetArnOk

`func (o *ECSTasksTasksInner) GetArnOk() (*string, bool)`

GetArnOk returns a tuple with the Arn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArn

`func (o *ECSTasksTasksInner) SetArn(v string)`

SetArn sets Arn field to given value.


### GetClusterArn

`func (o *ECSTasksTasksInner) GetClusterArn() string`

GetClusterArn returns the ClusterArn field if non-nil, zero value otherwise.

### GetClusterArnOk

`func (o *ECSTasksTasksInner) GetClusterArnOk() (*string, bool)`

GetClusterArnOk returns a tuple with the ClusterArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterArn

`func (o *ECSTasksTasksInner) SetClusterArn(v string)`

SetClusterArn sets ClusterArn field to given value.


### GetServiceArn

`func (o *ECSTasksTasksInner) GetServiceArn() string`

GetServiceArn returns the ServiceArn field if non-nil, zero value otherwise.

### GetServiceArnOk

`func (o *ECSTasksTasksInner) GetServiceArnOk() (*string, bool)`

GetServiceArnOk returns a tuple with the ServiceArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceArn

`func (o *ECSTasksTasksInner) SetServiceArn(v string)`

SetServiceArn sets ServiceArn field to given value.


### GetTaskDefinitionArn

`func (o *ECSTasksTasksInner) GetTaskDefinitionArn() string`

GetTaskDefinitionArn returns the TaskDefinitionArn field if non-nil, zero value otherwise.

### GetTaskDefinitionArnOk

`func (o *ECSTasksTasksInner) GetTaskDefinitionArnOk() (*string, bool)`

GetTaskDefinitionArnOk returns a tuple with the TaskDefinitionArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDefinitionArn

`func (o *ECSTasksTasksInner) SetTaskDefinitionArn(v string)`

SetTaskDefinitionArn sets TaskDefinitionArn field to given value.


### GetTags

`func (o *ECSTasksTasksInner) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ECSTasksTasksInner) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ECSTasksTasksInner) SetTags(v map[string]string)`

SetTags sets Tags field to given value.


### GetAccountName

`func (o *ECSTasksTasksInner) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *ECSTasksTasksInner) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *ECSTasksTasksInner) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


