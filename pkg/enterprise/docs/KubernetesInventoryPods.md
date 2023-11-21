# KubernetesInventoryPods

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** |  | 
**Name** | **string** |  | 
**NamespaceUid** | **string** |  | 
**NodeUid** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **map[string]string** |  | [optional] 
**Annotations** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewKubernetesInventoryPods

`func NewKubernetesInventoryPods(uid string, name string, namespaceUid string, ) *KubernetesInventoryPods`

NewKubernetesInventoryPods instantiates a new KubernetesInventoryPods object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesInventoryPodsWithDefaults

`func NewKubernetesInventoryPodsWithDefaults() *KubernetesInventoryPods`

NewKubernetesInventoryPodsWithDefaults instantiates a new KubernetesInventoryPods object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *KubernetesInventoryPods) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *KubernetesInventoryPods) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *KubernetesInventoryPods) SetUid(v string)`

SetUid sets Uid field to given value.


### GetName

`func (o *KubernetesInventoryPods) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesInventoryPods) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesInventoryPods) SetName(v string)`

SetName sets Name field to given value.


### GetNamespaceUid

`func (o *KubernetesInventoryPods) GetNamespaceUid() string`

GetNamespaceUid returns the NamespaceUid field if non-nil, zero value otherwise.

### GetNamespaceUidOk

`func (o *KubernetesInventoryPods) GetNamespaceUidOk() (*string, bool)`

GetNamespaceUidOk returns a tuple with the NamespaceUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceUid

`func (o *KubernetesInventoryPods) SetNamespaceUid(v string)`

SetNamespaceUid sets NamespaceUid field to given value.


### GetNodeUid

`func (o *KubernetesInventoryPods) GetNodeUid() string`

GetNodeUid returns the NodeUid field if non-nil, zero value otherwise.

### GetNodeUidOk

`func (o *KubernetesInventoryPods) GetNodeUidOk() (*string, bool)`

GetNodeUidOk returns a tuple with the NodeUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeUid

`func (o *KubernetesInventoryPods) SetNodeUid(v string)`

SetNodeUid sets NodeUid field to given value.

### HasNodeUid

`func (o *KubernetesInventoryPods) HasNodeUid() bool`

HasNodeUid returns a boolean if a field has been set.

### GetLabels

`func (o *KubernetesInventoryPods) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *KubernetesInventoryPods) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *KubernetesInventoryPods) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *KubernetesInventoryPods) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAnnotations

`func (o *KubernetesInventoryPods) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *KubernetesInventoryPods) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *KubernetesInventoryPods) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *KubernetesInventoryPods) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


