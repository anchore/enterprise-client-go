# KubernetesInventoryNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** |  | 
**Name** | **string** |  | 
**Labels** | Pointer to **map[string]string** |  | [optional] 
**Annotations** | Pointer to **map[string]string** |  | [optional] 
**KernelVersion** | Pointer to **string** |  | [optional] 
**KubernetesVersion** | Pointer to **string** |  | [optional] 
**Arch** | Pointer to **string** |  | [optional] 
**ContainerRuntimeVersion** | Pointer to **string** |  | [optional] 
**KubeProxyVersion** | Pointer to **string** |  | [optional] 
**KubeletVersion** | Pointer to **string** |  | [optional] 
**OperatingSystem** | Pointer to **string** |  | [optional] 

## Methods

### NewKubernetesInventoryNode

`func NewKubernetesInventoryNode(uid string, name string, ) *KubernetesInventoryNode`

NewKubernetesInventoryNode instantiates a new KubernetesInventoryNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesInventoryNodeWithDefaults

`func NewKubernetesInventoryNodeWithDefaults() *KubernetesInventoryNode`

NewKubernetesInventoryNodeWithDefaults instantiates a new KubernetesInventoryNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *KubernetesInventoryNode) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *KubernetesInventoryNode) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *KubernetesInventoryNode) SetUid(v string)`

SetUid sets Uid field to given value.


### GetName

`func (o *KubernetesInventoryNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesInventoryNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesInventoryNode) SetName(v string)`

SetName sets Name field to given value.


### GetLabels

`func (o *KubernetesInventoryNode) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *KubernetesInventoryNode) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *KubernetesInventoryNode) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *KubernetesInventoryNode) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAnnotations

`func (o *KubernetesInventoryNode) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *KubernetesInventoryNode) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *KubernetesInventoryNode) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *KubernetesInventoryNode) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetKernelVersion

`func (o *KubernetesInventoryNode) GetKernelVersion() string`

GetKernelVersion returns the KernelVersion field if non-nil, zero value otherwise.

### GetKernelVersionOk

`func (o *KubernetesInventoryNode) GetKernelVersionOk() (*string, bool)`

GetKernelVersionOk returns a tuple with the KernelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelVersion

`func (o *KubernetesInventoryNode) SetKernelVersion(v string)`

SetKernelVersion sets KernelVersion field to given value.

### HasKernelVersion

`func (o *KubernetesInventoryNode) HasKernelVersion() bool`

HasKernelVersion returns a boolean if a field has been set.

### GetKubernetesVersion

`func (o *KubernetesInventoryNode) GetKubernetesVersion() string`

GetKubernetesVersion returns the KubernetesVersion field if non-nil, zero value otherwise.

### GetKubernetesVersionOk

`func (o *KubernetesInventoryNode) GetKubernetesVersionOk() (*string, bool)`

GetKubernetesVersionOk returns a tuple with the KubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesVersion

`func (o *KubernetesInventoryNode) SetKubernetesVersion(v string)`

SetKubernetesVersion sets KubernetesVersion field to given value.

### HasKubernetesVersion

`func (o *KubernetesInventoryNode) HasKubernetesVersion() bool`

HasKubernetesVersion returns a boolean if a field has been set.

### GetArch

`func (o *KubernetesInventoryNode) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *KubernetesInventoryNode) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *KubernetesInventoryNode) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *KubernetesInventoryNode) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetContainerRuntimeVersion

`func (o *KubernetesInventoryNode) GetContainerRuntimeVersion() string`

GetContainerRuntimeVersion returns the ContainerRuntimeVersion field if non-nil, zero value otherwise.

### GetContainerRuntimeVersionOk

`func (o *KubernetesInventoryNode) GetContainerRuntimeVersionOk() (*string, bool)`

GetContainerRuntimeVersionOk returns a tuple with the ContainerRuntimeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerRuntimeVersion

`func (o *KubernetesInventoryNode) SetContainerRuntimeVersion(v string)`

SetContainerRuntimeVersion sets ContainerRuntimeVersion field to given value.

### HasContainerRuntimeVersion

`func (o *KubernetesInventoryNode) HasContainerRuntimeVersion() bool`

HasContainerRuntimeVersion returns a boolean if a field has been set.

### GetKubeProxyVersion

`func (o *KubernetesInventoryNode) GetKubeProxyVersion() string`

GetKubeProxyVersion returns the KubeProxyVersion field if non-nil, zero value otherwise.

### GetKubeProxyVersionOk

`func (o *KubernetesInventoryNode) GetKubeProxyVersionOk() (*string, bool)`

GetKubeProxyVersionOk returns a tuple with the KubeProxyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeProxyVersion

`func (o *KubernetesInventoryNode) SetKubeProxyVersion(v string)`

SetKubeProxyVersion sets KubeProxyVersion field to given value.

### HasKubeProxyVersion

`func (o *KubernetesInventoryNode) HasKubeProxyVersion() bool`

HasKubeProxyVersion returns a boolean if a field has been set.

### GetKubeletVersion

`func (o *KubernetesInventoryNode) GetKubeletVersion() string`

GetKubeletVersion returns the KubeletVersion field if non-nil, zero value otherwise.

### GetKubeletVersionOk

`func (o *KubernetesInventoryNode) GetKubeletVersionOk() (*string, bool)`

GetKubeletVersionOk returns a tuple with the KubeletVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeletVersion

`func (o *KubernetesInventoryNode) SetKubeletVersion(v string)`

SetKubeletVersion sets KubeletVersion field to given value.

### HasKubeletVersion

`func (o *KubernetesInventoryNode) HasKubeletVersion() bool`

HasKubeletVersion returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *KubernetesInventoryNode) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *KubernetesInventoryNode) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *KubernetesInventoryNode) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *KubernetesInventoryNode) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


