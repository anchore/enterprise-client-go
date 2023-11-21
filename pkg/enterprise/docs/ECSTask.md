# ECSTask

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

### NewECSTask

`func NewECSTask(arn string, clusterArn string, serviceArn string, taskDefinitionArn string, tags map[string]string, accountName string, ) *ECSTask`

NewECSTask instantiates a new ECSTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewECSTaskWithDefaults

`func NewECSTaskWithDefaults() *ECSTask`

NewECSTaskWithDefaults instantiates a new ECSTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArn

`func (o *ECSTask) GetArn() string`

GetArn returns the Arn field if non-nil, zero value otherwise.

### GetArnOk

`func (o *ECSTask) GetArnOk() (*string, bool)`

GetArnOk returns a tuple with the Arn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArn

`func (o *ECSTask) SetArn(v string)`

SetArn sets Arn field to given value.


### GetClusterArn

`func (o *ECSTask) GetClusterArn() string`

GetClusterArn returns the ClusterArn field if non-nil, zero value otherwise.

### GetClusterArnOk

`func (o *ECSTask) GetClusterArnOk() (*string, bool)`

GetClusterArnOk returns a tuple with the ClusterArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterArn

`func (o *ECSTask) SetClusterArn(v string)`

SetClusterArn sets ClusterArn field to given value.


### GetServiceArn

`func (o *ECSTask) GetServiceArn() string`

GetServiceArn returns the ServiceArn field if non-nil, zero value otherwise.

### GetServiceArnOk

`func (o *ECSTask) GetServiceArnOk() (*string, bool)`

GetServiceArnOk returns a tuple with the ServiceArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceArn

`func (o *ECSTask) SetServiceArn(v string)`

SetServiceArn sets ServiceArn field to given value.


### GetTaskDefinitionArn

`func (o *ECSTask) GetTaskDefinitionArn() string`

GetTaskDefinitionArn returns the TaskDefinitionArn field if non-nil, zero value otherwise.

### GetTaskDefinitionArnOk

`func (o *ECSTask) GetTaskDefinitionArnOk() (*string, bool)`

GetTaskDefinitionArnOk returns a tuple with the TaskDefinitionArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDefinitionArn

`func (o *ECSTask) SetTaskDefinitionArn(v string)`

SetTaskDefinitionArn sets TaskDefinitionArn field to given value.


### GetTags

`func (o *ECSTask) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ECSTask) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ECSTask) SetTags(v map[string]string)`

SetTags sets Tags field to given value.


### GetAccountName

`func (o *ECSTask) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *ECSTask) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *ECSTask) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


