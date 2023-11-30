# ECSContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arn** | **string** |  | 
**TaskArn** | **string** |  | 
**AccountName** | **string** |  | 
**Context** | **string** |  | 
**ImageTag** | **string** |  | 
**ImageDigest** | **string** |  | 

## Methods

### NewECSContainer

`func NewECSContainer(arn string, taskArn string, accountName string, context string, imageTag string, imageDigest string, ) *ECSContainer`

NewECSContainer instantiates a new ECSContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewECSContainerWithDefaults

`func NewECSContainerWithDefaults() *ECSContainer`

NewECSContainerWithDefaults instantiates a new ECSContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArn

`func (o *ECSContainer) GetArn() string`

GetArn returns the Arn field if non-nil, zero value otherwise.

### GetArnOk

`func (o *ECSContainer) GetArnOk() (*string, bool)`

GetArnOk returns a tuple with the Arn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArn

`func (o *ECSContainer) SetArn(v string)`

SetArn sets Arn field to given value.


### GetTaskArn

`func (o *ECSContainer) GetTaskArn() string`

GetTaskArn returns the TaskArn field if non-nil, zero value otherwise.

### GetTaskArnOk

`func (o *ECSContainer) GetTaskArnOk() (*string, bool)`

GetTaskArnOk returns a tuple with the TaskArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskArn

`func (o *ECSContainer) SetTaskArn(v string)`

SetTaskArn sets TaskArn field to given value.


### GetAccountName

`func (o *ECSContainer) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *ECSContainer) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *ECSContainer) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.


### GetContext

`func (o *ECSContainer) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ECSContainer) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ECSContainer) SetContext(v string)`

SetContext sets Context field to given value.


### GetImageTag

`func (o *ECSContainer) GetImageTag() string`

GetImageTag returns the ImageTag field if non-nil, zero value otherwise.

### GetImageTagOk

`func (o *ECSContainer) GetImageTagOk() (*string, bool)`

GetImageTagOk returns a tuple with the ImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTag

`func (o *ECSContainer) SetImageTag(v string)`

SetImageTag sets ImageTag field to given value.


### GetImageDigest

`func (o *ECSContainer) GetImageDigest() string`

GetImageDigest returns the ImageDigest field if non-nil, zero value otherwise.

### GetImageDigestOk

`func (o *ECSContainer) GetImageDigestOk() (*string, bool)`

GetImageDigestOk returns a tuple with the ImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDigest

`func (o *ECSContainer) SetImageDigest(v string)`

SetImageDigest sets ImageDigest field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


