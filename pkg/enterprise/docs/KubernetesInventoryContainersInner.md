# KubernetesInventoryContainersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**ImageTag** | **string** |  | 
**ImageDigest** | Pointer to **string** |  | [optional] 
**PodUid** | **string** |  | 

## Methods

### NewKubernetesInventoryContainersInner

`func NewKubernetesInventoryContainersInner(id string, name string, imageTag string, podUid string, ) *KubernetesInventoryContainersInner`

NewKubernetesInventoryContainersInner instantiates a new KubernetesInventoryContainersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesInventoryContainersInnerWithDefaults

`func NewKubernetesInventoryContainersInnerWithDefaults() *KubernetesInventoryContainersInner`

NewKubernetesInventoryContainersInnerWithDefaults instantiates a new KubernetesInventoryContainersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesInventoryContainersInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesInventoryContainersInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesInventoryContainersInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *KubernetesInventoryContainersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesInventoryContainersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesInventoryContainersInner) SetName(v string)`

SetName sets Name field to given value.


### GetImageTag

`func (o *KubernetesInventoryContainersInner) GetImageTag() string`

GetImageTag returns the ImageTag field if non-nil, zero value otherwise.

### GetImageTagOk

`func (o *KubernetesInventoryContainersInner) GetImageTagOk() (*string, bool)`

GetImageTagOk returns a tuple with the ImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTag

`func (o *KubernetesInventoryContainersInner) SetImageTag(v string)`

SetImageTag sets ImageTag field to given value.


### GetImageDigest

`func (o *KubernetesInventoryContainersInner) GetImageDigest() string`

GetImageDigest returns the ImageDigest field if non-nil, zero value otherwise.

### GetImageDigestOk

`func (o *KubernetesInventoryContainersInner) GetImageDigestOk() (*string, bool)`

GetImageDigestOk returns a tuple with the ImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDigest

`func (o *KubernetesInventoryContainersInner) SetImageDigest(v string)`

SetImageDigest sets ImageDigest field to given value.

### HasImageDigest

`func (o *KubernetesInventoryContainersInner) HasImageDigest() bool`

HasImageDigest returns a boolean if a field has been set.

### GetPodUid

`func (o *KubernetesInventoryContainersInner) GetPodUid() string`

GetPodUid returns the PodUid field if non-nil, zero value otherwise.

### GetPodUidOk

`func (o *KubernetesInventoryContainersInner) GetPodUidOk() (*string, bool)`

GetPodUidOk returns a tuple with the PodUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodUid

`func (o *KubernetesInventoryContainersInner) SetPodUid(v string)`

SetPodUid sets PodUid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


