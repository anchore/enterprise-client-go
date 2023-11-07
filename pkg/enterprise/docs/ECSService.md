# ECSService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arn** | **string** |  | 
**ClusterArn** | **string** |  | 
**Tags** | **map[string]string** |  | 
**AccountName** | **string** |  | 

## Methods

### NewECSService

`func NewECSService(arn string, clusterArn string, tags map[string]string, accountName string, ) *ECSService`

NewECSService instantiates a new ECSService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewECSServiceWithDefaults

`func NewECSServiceWithDefaults() *ECSService`

NewECSServiceWithDefaults instantiates a new ECSService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArn

`func (o *ECSService) GetArn() string`

GetArn returns the Arn field if non-nil, zero value otherwise.

### GetArnOk

`func (o *ECSService) GetArnOk() (*string, bool)`

GetArnOk returns a tuple with the Arn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArn

`func (o *ECSService) SetArn(v string)`

SetArn sets Arn field to given value.


### GetClusterArn

`func (o *ECSService) GetClusterArn() string`

GetClusterArn returns the ClusterArn field if non-nil, zero value otherwise.

### GetClusterArnOk

`func (o *ECSService) GetClusterArnOk() (*string, bool)`

GetClusterArnOk returns a tuple with the ClusterArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterArn

`func (o *ECSService) SetClusterArn(v string)`

SetClusterArn sets ClusterArn field to given value.


### GetTags

`func (o *ECSService) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ECSService) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ECSService) SetTags(v map[string]string)`

SetTags sets Tags field to given value.


### GetAccountName

`func (o *ECSService) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *ECSService) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *ECSService) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


