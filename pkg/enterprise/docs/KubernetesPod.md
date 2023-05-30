# KubernetesPod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**AccountName** | **string** |  | 
**Labels** | **map[string]string** |  | 
**Annotations** | **map[string]string** |  | 
**NodeId** | Pointer to **string** |  | [optional] 
**NamespaceId** | Pointer to **string** |  | [optional] 
**LastSeen** | Pointer to **string** |  | [optional] 

## Methods

### NewKubernetesPod

`func NewKubernetesPod(id string, name string, accountName string, labels map[string]string, annotations map[string]string, ) *KubernetesPod`

NewKubernetesPod instantiates a new KubernetesPod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesPodWithDefaults

`func NewKubernetesPodWithDefaults() *KubernetesPod`

NewKubernetesPodWithDefaults instantiates a new KubernetesPod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesPod) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesPod) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesPod) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *KubernetesPod) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesPod) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesPod) SetName(v string)`

SetName sets Name field to given value.


### GetAccountName

`func (o *KubernetesPod) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *KubernetesPod) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *KubernetesPod) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.


### GetLabels

`func (o *KubernetesPod) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *KubernetesPod) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *KubernetesPod) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.


### GetAnnotations

`func (o *KubernetesPod) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *KubernetesPod) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *KubernetesPod) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.


### GetNodeId

`func (o *KubernetesPod) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *KubernetesPod) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *KubernetesPod) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *KubernetesPod) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetNamespaceId

`func (o *KubernetesPod) GetNamespaceId() string`

GetNamespaceId returns the NamespaceId field if non-nil, zero value otherwise.

### GetNamespaceIdOk

`func (o *KubernetesPod) GetNamespaceIdOk() (*string, bool)`

GetNamespaceIdOk returns a tuple with the NamespaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceId

`func (o *KubernetesPod) SetNamespaceId(v string)`

SetNamespaceId sets NamespaceId field to given value.

### HasNamespaceId

`func (o *KubernetesPod) HasNamespaceId() bool`

HasNamespaceId returns a boolean if a field has been set.

### GetLastSeen

`func (o *KubernetesPod) GetLastSeen() string`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *KubernetesPod) GetLastSeenOk() (*string, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *KubernetesPod) SetLastSeen(v string)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *KubernetesPod) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


