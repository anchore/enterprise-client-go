# KubernetesInventoryContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Corresponds to ContainerStatus.containerID in the Kubernetes Spec | 
**Name** | **string** |  | 
**ImageTag** | **string** |  | 
**ImageDigest** | Pointer to **string** |  | [optional] 
**PodUid** | **string** |  | 

## Methods

### NewKubernetesInventoryContainer

`func NewKubernetesInventoryContainer(id string, name string, imageTag string, podUid string, ) *KubernetesInventoryContainer`

NewKubernetesInventoryContainer instantiates a new KubernetesInventoryContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesInventoryContainerWithDefaults

`func NewKubernetesInventoryContainerWithDefaults() *KubernetesInventoryContainer`

NewKubernetesInventoryContainerWithDefaults instantiates a new KubernetesInventoryContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesInventoryContainer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesInventoryContainer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesInventoryContainer) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *KubernetesInventoryContainer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesInventoryContainer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesInventoryContainer) SetName(v string)`

SetName sets Name field to given value.


### GetImageTag

`func (o *KubernetesInventoryContainer) GetImageTag() string`

GetImageTag returns the ImageTag field if non-nil, zero value otherwise.

### GetImageTagOk

`func (o *KubernetesInventoryContainer) GetImageTagOk() (*string, bool)`

GetImageTagOk returns a tuple with the ImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTag

`func (o *KubernetesInventoryContainer) SetImageTag(v string)`

SetImageTag sets ImageTag field to given value.


### GetImageDigest

`func (o *KubernetesInventoryContainer) GetImageDigest() string`

GetImageDigest returns the ImageDigest field if non-nil, zero value otherwise.

### GetImageDigestOk

`func (o *KubernetesInventoryContainer) GetImageDigestOk() (*string, bool)`

GetImageDigestOk returns a tuple with the ImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDigest

`func (o *KubernetesInventoryContainer) SetImageDigest(v string)`

SetImageDigest sets ImageDigest field to given value.

### HasImageDigest

`func (o *KubernetesInventoryContainer) HasImageDigest() bool`

HasImageDigest returns a boolean if a field has been set.

### GetPodUid

`func (o *KubernetesInventoryContainer) GetPodUid() string`

GetPodUid returns the PodUid field if non-nil, zero value otherwise.

### GetPodUidOk

`func (o *KubernetesInventoryContainer) GetPodUidOk() (*string, bool)`

GetPodUidOk returns a tuple with the PodUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodUid

`func (o *KubernetesInventoryContainer) SetPodUid(v string)`

SetPodUid sets PodUid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


