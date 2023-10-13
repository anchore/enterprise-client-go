# KubernetesInventoryNamespacesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** |  | 
**Name** | **string** |  | 
**Labels** | Pointer to **map[string]string** |  | [optional] 
**Annotations** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewKubernetesInventoryNamespacesInner

`func NewKubernetesInventoryNamespacesInner(uid string, name string, ) *KubernetesInventoryNamespacesInner`

NewKubernetesInventoryNamespacesInner instantiates a new KubernetesInventoryNamespacesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesInventoryNamespacesInnerWithDefaults

`func NewKubernetesInventoryNamespacesInnerWithDefaults() *KubernetesInventoryNamespacesInner`

NewKubernetesInventoryNamespacesInnerWithDefaults instantiates a new KubernetesInventoryNamespacesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *KubernetesInventoryNamespacesInner) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *KubernetesInventoryNamespacesInner) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *KubernetesInventoryNamespacesInner) SetUid(v string)`

SetUid sets Uid field to given value.


### GetName

`func (o *KubernetesInventoryNamespacesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesInventoryNamespacesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesInventoryNamespacesInner) SetName(v string)`

SetName sets Name field to given value.


### GetLabels

`func (o *KubernetesInventoryNamespacesInner) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *KubernetesInventoryNamespacesInner) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *KubernetesInventoryNamespacesInner) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *KubernetesInventoryNamespacesInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAnnotations

`func (o *KubernetesInventoryNamespacesInner) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *KubernetesInventoryNamespacesInner) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *KubernetesInventoryNamespacesInner) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *KubernetesInventoryNamespacesInner) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


