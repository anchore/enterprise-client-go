# KubernetesInventoryNodes

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

### NewKubernetesInventoryNodes

`func NewKubernetesInventoryNodes(uid string, name string, ) *KubernetesInventoryNodes`

NewKubernetesInventoryNodes instantiates a new KubernetesInventoryNodes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesInventoryNodesWithDefaults

`func NewKubernetesInventoryNodesWithDefaults() *KubernetesInventoryNodes`

NewKubernetesInventoryNodesWithDefaults instantiates a new KubernetesInventoryNodes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *KubernetesInventoryNodes) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *KubernetesInventoryNodes) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *KubernetesInventoryNodes) SetUid(v string)`

SetUid sets Uid field to given value.


### GetName

`func (o *KubernetesInventoryNodes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesInventoryNodes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesInventoryNodes) SetName(v string)`

SetName sets Name field to given value.


### GetLabels

`func (o *KubernetesInventoryNodes) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *KubernetesInventoryNodes) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *KubernetesInventoryNodes) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *KubernetesInventoryNodes) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAnnotations

`func (o *KubernetesInventoryNodes) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *KubernetesInventoryNodes) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *KubernetesInventoryNodes) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *KubernetesInventoryNodes) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetKernelVersion

`func (o *KubernetesInventoryNodes) GetKernelVersion() string`

GetKernelVersion returns the KernelVersion field if non-nil, zero value otherwise.

### GetKernelVersionOk

`func (o *KubernetesInventoryNodes) GetKernelVersionOk() (*string, bool)`

GetKernelVersionOk returns a tuple with the KernelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelVersion

`func (o *KubernetesInventoryNodes) SetKernelVersion(v string)`

SetKernelVersion sets KernelVersion field to given value.

### HasKernelVersion

`func (o *KubernetesInventoryNodes) HasKernelVersion() bool`

HasKernelVersion returns a boolean if a field has been set.

### GetKubernetesVersion

`func (o *KubernetesInventoryNodes) GetKubernetesVersion() string`

GetKubernetesVersion returns the KubernetesVersion field if non-nil, zero value otherwise.

### GetKubernetesVersionOk

`func (o *KubernetesInventoryNodes) GetKubernetesVersionOk() (*string, bool)`

GetKubernetesVersionOk returns a tuple with the KubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesVersion

`func (o *KubernetesInventoryNodes) SetKubernetesVersion(v string)`

SetKubernetesVersion sets KubernetesVersion field to given value.

### HasKubernetesVersion

`func (o *KubernetesInventoryNodes) HasKubernetesVersion() bool`

HasKubernetesVersion returns a boolean if a field has been set.

### GetArch

`func (o *KubernetesInventoryNodes) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *KubernetesInventoryNodes) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *KubernetesInventoryNodes) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *KubernetesInventoryNodes) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetContainerRuntimeVersion

`func (o *KubernetesInventoryNodes) GetContainerRuntimeVersion() string`

GetContainerRuntimeVersion returns the ContainerRuntimeVersion field if non-nil, zero value otherwise.

### GetContainerRuntimeVersionOk

`func (o *KubernetesInventoryNodes) GetContainerRuntimeVersionOk() (*string, bool)`

GetContainerRuntimeVersionOk returns a tuple with the ContainerRuntimeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerRuntimeVersion

`func (o *KubernetesInventoryNodes) SetContainerRuntimeVersion(v string)`

SetContainerRuntimeVersion sets ContainerRuntimeVersion field to given value.

### HasContainerRuntimeVersion

`func (o *KubernetesInventoryNodes) HasContainerRuntimeVersion() bool`

HasContainerRuntimeVersion returns a boolean if a field has been set.

### GetKubeProxyVersion

`func (o *KubernetesInventoryNodes) GetKubeProxyVersion() string`

GetKubeProxyVersion returns the KubeProxyVersion field if non-nil, zero value otherwise.

### GetKubeProxyVersionOk

`func (o *KubernetesInventoryNodes) GetKubeProxyVersionOk() (*string, bool)`

GetKubeProxyVersionOk returns a tuple with the KubeProxyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeProxyVersion

`func (o *KubernetesInventoryNodes) SetKubeProxyVersion(v string)`

SetKubeProxyVersion sets KubeProxyVersion field to given value.

### HasKubeProxyVersion

`func (o *KubernetesInventoryNodes) HasKubeProxyVersion() bool`

HasKubeProxyVersion returns a boolean if a field has been set.

### GetKubeletVersion

`func (o *KubernetesInventoryNodes) GetKubeletVersion() string`

GetKubeletVersion returns the KubeletVersion field if non-nil, zero value otherwise.

### GetKubeletVersionOk

`func (o *KubernetesInventoryNodes) GetKubeletVersionOk() (*string, bool)`

GetKubeletVersionOk returns a tuple with the KubeletVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeletVersion

`func (o *KubernetesInventoryNodes) SetKubeletVersion(v string)`

SetKubeletVersion sets KubeletVersion field to given value.

### HasKubeletVersion

`func (o *KubernetesInventoryNodes) HasKubeletVersion() bool`

HasKubeletVersion returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *KubernetesInventoryNodes) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *KubernetesInventoryNodes) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *KubernetesInventoryNodes) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *KubernetesInventoryNodes) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


