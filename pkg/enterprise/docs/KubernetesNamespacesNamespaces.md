# KubernetesNamespacesNamespaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** |  | 
**Name** | **string** |  | 
**Labels** | **map[string]string** |  | 
**Annotations** | **map[string]string** |  | 

## Methods

### NewKubernetesNamespacesNamespaces

`func NewKubernetesNamespacesNamespaces(uid string, name string, labels map[string]string, annotations map[string]string, ) *KubernetesNamespacesNamespaces`

NewKubernetesNamespacesNamespaces instantiates a new KubernetesNamespacesNamespaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNamespacesNamespacesWithDefaults

`func NewKubernetesNamespacesNamespacesWithDefaults() *KubernetesNamespacesNamespaces`

NewKubernetesNamespacesNamespacesWithDefaults instantiates a new KubernetesNamespacesNamespaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *KubernetesNamespacesNamespaces) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *KubernetesNamespacesNamespaces) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *KubernetesNamespacesNamespaces) SetUid(v string)`

SetUid sets Uid field to given value.


### GetName

`func (o *KubernetesNamespacesNamespaces) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesNamespacesNamespaces) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesNamespacesNamespaces) SetName(v string)`

SetName sets Name field to given value.


### GetLabels

`func (o *KubernetesNamespacesNamespaces) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *KubernetesNamespacesNamespaces) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *KubernetesNamespacesNamespaces) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.


### GetAnnotations

`func (o *KubernetesNamespacesNamespaces) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *KubernetesNamespacesNamespaces) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *KubernetesNamespacesNamespaces) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


