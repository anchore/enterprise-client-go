# KubernetesInventoryPodsInner

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

### NewKubernetesInventoryPodsInner

`func NewKubernetesInventoryPodsInner(uid string, name string, namespaceUid string, ) *KubernetesInventoryPodsInner`

NewKubernetesInventoryPodsInner instantiates a new KubernetesInventoryPodsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesInventoryPodsInnerWithDefaults

`func NewKubernetesInventoryPodsInnerWithDefaults() *KubernetesInventoryPodsInner`

NewKubernetesInventoryPodsInnerWithDefaults instantiates a new KubernetesInventoryPodsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *KubernetesInventoryPodsInner) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *KubernetesInventoryPodsInner) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *KubernetesInventoryPodsInner) SetUid(v string)`

SetUid sets Uid field to given value.


### GetName

`func (o *KubernetesInventoryPodsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesInventoryPodsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesInventoryPodsInner) SetName(v string)`

SetName sets Name field to given value.


### GetNamespaceUid

`func (o *KubernetesInventoryPodsInner) GetNamespaceUid() string`

GetNamespaceUid returns the NamespaceUid field if non-nil, zero value otherwise.

### GetNamespaceUidOk

`func (o *KubernetesInventoryPodsInner) GetNamespaceUidOk() (*string, bool)`

GetNamespaceUidOk returns a tuple with the NamespaceUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceUid

`func (o *KubernetesInventoryPodsInner) SetNamespaceUid(v string)`

SetNamespaceUid sets NamespaceUid field to given value.


### GetNodeUid

`func (o *KubernetesInventoryPodsInner) GetNodeUid() string`

GetNodeUid returns the NodeUid field if non-nil, zero value otherwise.

### GetNodeUidOk

`func (o *KubernetesInventoryPodsInner) GetNodeUidOk() (*string, bool)`

GetNodeUidOk returns a tuple with the NodeUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeUid

`func (o *KubernetesInventoryPodsInner) SetNodeUid(v string)`

SetNodeUid sets NodeUid field to given value.

### HasNodeUid

`func (o *KubernetesInventoryPodsInner) HasNodeUid() bool`

HasNodeUid returns a boolean if a field has been set.

### GetLabels

`func (o *KubernetesInventoryPodsInner) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *KubernetesInventoryPodsInner) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *KubernetesInventoryPodsInner) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *KubernetesInventoryPodsInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAnnotations

`func (o *KubernetesInventoryPodsInner) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *KubernetesInventoryPodsInner) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *KubernetesInventoryPodsInner) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *KubernetesInventoryPodsInner) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


