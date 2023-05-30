# KubernetesNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** |  | 
**Name** | **string** |  | 
**Labels** | **map[string]string** |  | 
**Annotations** | **map[string]string** |  | 
**LastSeen** | Pointer to **string** |  | [optional] 

## Methods

### NewKubernetesNode

`func NewKubernetesNode(uid string, name string, labels map[string]string, annotations map[string]string, ) *KubernetesNode`

NewKubernetesNode instantiates a new KubernetesNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNodeWithDefaults

`func NewKubernetesNodeWithDefaults() *KubernetesNode`

NewKubernetesNodeWithDefaults instantiates a new KubernetesNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *KubernetesNode) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *KubernetesNode) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *KubernetesNode) SetUid(v string)`

SetUid sets Uid field to given value.


### GetName

`func (o *KubernetesNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesNode) SetName(v string)`

SetName sets Name field to given value.


### GetLabels

`func (o *KubernetesNode) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *KubernetesNode) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *KubernetesNode) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.


### GetAnnotations

`func (o *KubernetesNode) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *KubernetesNode) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *KubernetesNode) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.


### GetLastSeen

`func (o *KubernetesNode) GetLastSeen() string`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *KubernetesNode) GetLastSeenOk() (*string, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *KubernetesNode) SetLastSeen(v string)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *KubernetesNode) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


