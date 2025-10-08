# AnchoreSummaryKubernetesNamespacesItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**TotalImages** | Pointer to **int32** |  | [optional] 
**CompliantImages** | Pointer to **int32** |  | [optional] 
**NonCompliantImages** | Pointer to **int32** |  | [optional] 
**NotEvaluatedImages** | Pointer to **int32** |  | [optional] 
**Watched** | Pointer to **bool** |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 

## Methods

### NewAnchoreSummaryKubernetesNamespacesItem

`func NewAnchoreSummaryKubernetesNamespacesItem() *AnchoreSummaryKubernetesNamespacesItem`

NewAnchoreSummaryKubernetesNamespacesItem instantiates a new AnchoreSummaryKubernetesNamespacesItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnchoreSummaryKubernetesNamespacesItemWithDefaults

`func NewAnchoreSummaryKubernetesNamespacesItemWithDefaults() *AnchoreSummaryKubernetesNamespacesItem`

NewAnchoreSummaryKubernetesNamespacesItemWithDefaults instantiates a new AnchoreSummaryKubernetesNamespacesItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AnchoreSummaryKubernetesNamespacesItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AnchoreSummaryKubernetesNamespacesItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AnchoreSummaryKubernetesNamespacesItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AnchoreSummaryKubernetesNamespacesItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTotalImages

`func (o *AnchoreSummaryKubernetesNamespacesItem) GetTotalImages() int32`

GetTotalImages returns the TotalImages field if non-nil, zero value otherwise.

### GetTotalImagesOk

`func (o *AnchoreSummaryKubernetesNamespacesItem) GetTotalImagesOk() (*int32, bool)`

GetTotalImagesOk returns a tuple with the TotalImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalImages

`func (o *AnchoreSummaryKubernetesNamespacesItem) SetTotalImages(v int32)`

SetTotalImages sets TotalImages field to given value.

### HasTotalImages

`func (o *AnchoreSummaryKubernetesNamespacesItem) HasTotalImages() bool`

HasTotalImages returns a boolean if a field has been set.

### GetCompliantImages

`func (o *AnchoreSummaryKubernetesNamespacesItem) GetCompliantImages() int32`

GetCompliantImages returns the CompliantImages field if non-nil, zero value otherwise.

### GetCompliantImagesOk

`func (o *AnchoreSummaryKubernetesNamespacesItem) GetCompliantImagesOk() (*int32, bool)`

GetCompliantImagesOk returns a tuple with the CompliantImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliantImages

`func (o *AnchoreSummaryKubernetesNamespacesItem) SetCompliantImages(v int32)`

SetCompliantImages sets CompliantImages field to given value.

### HasCompliantImages

`func (o *AnchoreSummaryKubernetesNamespacesItem) HasCompliantImages() bool`

HasCompliantImages returns a boolean if a field has been set.

### GetNonCompliantImages

`func (o *AnchoreSummaryKubernetesNamespacesItem) GetNonCompliantImages() int32`

GetNonCompliantImages returns the NonCompliantImages field if non-nil, zero value otherwise.

### GetNonCompliantImagesOk

`func (o *AnchoreSummaryKubernetesNamespacesItem) GetNonCompliantImagesOk() (*int32, bool)`

GetNonCompliantImagesOk returns a tuple with the NonCompliantImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonCompliantImages

`func (o *AnchoreSummaryKubernetesNamespacesItem) SetNonCompliantImages(v int32)`

SetNonCompliantImages sets NonCompliantImages field to given value.

### HasNonCompliantImages

`func (o *AnchoreSummaryKubernetesNamespacesItem) HasNonCompliantImages() bool`

HasNonCompliantImages returns a boolean if a field has been set.

### GetNotEvaluatedImages

`func (o *AnchoreSummaryKubernetesNamespacesItem) GetNotEvaluatedImages() int32`

GetNotEvaluatedImages returns the NotEvaluatedImages field if non-nil, zero value otherwise.

### GetNotEvaluatedImagesOk

`func (o *AnchoreSummaryKubernetesNamespacesItem) GetNotEvaluatedImagesOk() (*int32, bool)`

GetNotEvaluatedImagesOk returns a tuple with the NotEvaluatedImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotEvaluatedImages

`func (o *AnchoreSummaryKubernetesNamespacesItem) SetNotEvaluatedImages(v int32)`

SetNotEvaluatedImages sets NotEvaluatedImages field to given value.

### HasNotEvaluatedImages

`func (o *AnchoreSummaryKubernetesNamespacesItem) HasNotEvaluatedImages() bool`

HasNotEvaluatedImages returns a boolean if a field has been set.

### GetWatched

`func (o *AnchoreSummaryKubernetesNamespacesItem) GetWatched() bool`

GetWatched returns the Watched field if non-nil, zero value otherwise.

### GetWatchedOk

`func (o *AnchoreSummaryKubernetesNamespacesItem) GetWatchedOk() (*bool, bool)`

GetWatchedOk returns a tuple with the Watched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatched

`func (o *AnchoreSummaryKubernetesNamespacesItem) SetWatched(v bool)`

SetWatched sets Watched field to given value.

### HasWatched

`func (o *AnchoreSummaryKubernetesNamespacesItem) HasWatched() bool`

HasWatched returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *AnchoreSummaryKubernetesNamespacesItem) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AnchoreSummaryKubernetesNamespacesItem) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AnchoreSummaryKubernetesNamespacesItem) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *AnchoreSummaryKubernetesNamespacesItem) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


