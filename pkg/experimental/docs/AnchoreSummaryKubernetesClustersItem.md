# AnchoreSummaryKubernetesClustersItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**NamespacesCount** | Pointer to **int32** |  | [optional] 
**TotalImages** | Pointer to **int32** |  | [optional] 
**CompliantImages** | Pointer to **int32** |  | [optional] 
**NonCompliantImages** | Pointer to **int32** |  | [optional] 
**NotEvaluatedImages** | Pointer to **int32** |  | [optional] 
**Watched** | Pointer to **bool** |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**Namespaces** | Pointer to [**[]AnchoreSummaryKubernetesNamespacesItem**](AnchoreSummaryKubernetesNamespacesItem.md) |  | [optional] 

## Methods

### NewAnchoreSummaryKubernetesClustersItem

`func NewAnchoreSummaryKubernetesClustersItem() *AnchoreSummaryKubernetesClustersItem`

NewAnchoreSummaryKubernetesClustersItem instantiates a new AnchoreSummaryKubernetesClustersItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnchoreSummaryKubernetesClustersItemWithDefaults

`func NewAnchoreSummaryKubernetesClustersItemWithDefaults() *AnchoreSummaryKubernetesClustersItem`

NewAnchoreSummaryKubernetesClustersItemWithDefaults instantiates a new AnchoreSummaryKubernetesClustersItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AnchoreSummaryKubernetesClustersItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AnchoreSummaryKubernetesClustersItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AnchoreSummaryKubernetesClustersItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AnchoreSummaryKubernetesClustersItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespacesCount

`func (o *AnchoreSummaryKubernetesClustersItem) GetNamespacesCount() int32`

GetNamespacesCount returns the NamespacesCount field if non-nil, zero value otherwise.

### GetNamespacesCountOk

`func (o *AnchoreSummaryKubernetesClustersItem) GetNamespacesCountOk() (*int32, bool)`

GetNamespacesCountOk returns a tuple with the NamespacesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespacesCount

`func (o *AnchoreSummaryKubernetesClustersItem) SetNamespacesCount(v int32)`

SetNamespacesCount sets NamespacesCount field to given value.

### HasNamespacesCount

`func (o *AnchoreSummaryKubernetesClustersItem) HasNamespacesCount() bool`

HasNamespacesCount returns a boolean if a field has been set.

### GetTotalImages

`func (o *AnchoreSummaryKubernetesClustersItem) GetTotalImages() int32`

GetTotalImages returns the TotalImages field if non-nil, zero value otherwise.

### GetTotalImagesOk

`func (o *AnchoreSummaryKubernetesClustersItem) GetTotalImagesOk() (*int32, bool)`

GetTotalImagesOk returns a tuple with the TotalImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalImages

`func (o *AnchoreSummaryKubernetesClustersItem) SetTotalImages(v int32)`

SetTotalImages sets TotalImages field to given value.

### HasTotalImages

`func (o *AnchoreSummaryKubernetesClustersItem) HasTotalImages() bool`

HasTotalImages returns a boolean if a field has been set.

### GetCompliantImages

`func (o *AnchoreSummaryKubernetesClustersItem) GetCompliantImages() int32`

GetCompliantImages returns the CompliantImages field if non-nil, zero value otherwise.

### GetCompliantImagesOk

`func (o *AnchoreSummaryKubernetesClustersItem) GetCompliantImagesOk() (*int32, bool)`

GetCompliantImagesOk returns a tuple with the CompliantImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliantImages

`func (o *AnchoreSummaryKubernetesClustersItem) SetCompliantImages(v int32)`

SetCompliantImages sets CompliantImages field to given value.

### HasCompliantImages

`func (o *AnchoreSummaryKubernetesClustersItem) HasCompliantImages() bool`

HasCompliantImages returns a boolean if a field has been set.

### GetNonCompliantImages

`func (o *AnchoreSummaryKubernetesClustersItem) GetNonCompliantImages() int32`

GetNonCompliantImages returns the NonCompliantImages field if non-nil, zero value otherwise.

### GetNonCompliantImagesOk

`func (o *AnchoreSummaryKubernetesClustersItem) GetNonCompliantImagesOk() (*int32, bool)`

GetNonCompliantImagesOk returns a tuple with the NonCompliantImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonCompliantImages

`func (o *AnchoreSummaryKubernetesClustersItem) SetNonCompliantImages(v int32)`

SetNonCompliantImages sets NonCompliantImages field to given value.

### HasNonCompliantImages

`func (o *AnchoreSummaryKubernetesClustersItem) HasNonCompliantImages() bool`

HasNonCompliantImages returns a boolean if a field has been set.

### GetNotEvaluatedImages

`func (o *AnchoreSummaryKubernetesClustersItem) GetNotEvaluatedImages() int32`

GetNotEvaluatedImages returns the NotEvaluatedImages field if non-nil, zero value otherwise.

### GetNotEvaluatedImagesOk

`func (o *AnchoreSummaryKubernetesClustersItem) GetNotEvaluatedImagesOk() (*int32, bool)`

GetNotEvaluatedImagesOk returns a tuple with the NotEvaluatedImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotEvaluatedImages

`func (o *AnchoreSummaryKubernetesClustersItem) SetNotEvaluatedImages(v int32)`

SetNotEvaluatedImages sets NotEvaluatedImages field to given value.

### HasNotEvaluatedImages

`func (o *AnchoreSummaryKubernetesClustersItem) HasNotEvaluatedImages() bool`

HasNotEvaluatedImages returns a boolean if a field has been set.

### GetWatched

`func (o *AnchoreSummaryKubernetesClustersItem) GetWatched() bool`

GetWatched returns the Watched field if non-nil, zero value otherwise.

### GetWatchedOk

`func (o *AnchoreSummaryKubernetesClustersItem) GetWatchedOk() (*bool, bool)`

GetWatchedOk returns a tuple with the Watched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatched

`func (o *AnchoreSummaryKubernetesClustersItem) SetWatched(v bool)`

SetWatched sets Watched field to given value.

### HasWatched

`func (o *AnchoreSummaryKubernetesClustersItem) HasWatched() bool`

HasWatched returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *AnchoreSummaryKubernetesClustersItem) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AnchoreSummaryKubernetesClustersItem) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AnchoreSummaryKubernetesClustersItem) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *AnchoreSummaryKubernetesClustersItem) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetNamespaces

`func (o *AnchoreSummaryKubernetesClustersItem) GetNamespaces() []AnchoreSummaryKubernetesNamespacesItem`

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

`func (o *AnchoreSummaryKubernetesClustersItem) GetNamespacesOk() (*[]AnchoreSummaryKubernetesNamespacesItem, bool)`

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

`func (o *AnchoreSummaryKubernetesClustersItem) SetNamespaces(v []AnchoreSummaryKubernetesNamespacesItem)`

SetNamespaces sets Namespaces field to given value.

### HasNamespaces

`func (o *AnchoreSummaryKubernetesClustersItem) HasNamespaces() bool`

HasNamespaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


