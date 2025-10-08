# AnchoreSummaryImageTagItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalysisStatus** | Pointer to **string** |  | [optional] 
**ImageDigest** | Pointer to **string** |  | [optional] 
**AnalyzedAt** | Pointer to **NullableTime** |  | [optional] 
**ImageId** | Pointer to **string** |  | [optional] 
**TagDetectedAt** | Pointer to **time.Time** |  | [optional] 
**Inventory** | Pointer to [**[]AnchoreSummaryImageTagItemInventoryInner**](AnchoreSummaryImageTagItemInventoryInner.md) |  | [optional] 

## Methods

### NewAnchoreSummaryImageTagItem

`func NewAnchoreSummaryImageTagItem() *AnchoreSummaryImageTagItem`

NewAnchoreSummaryImageTagItem instantiates a new AnchoreSummaryImageTagItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnchoreSummaryImageTagItemWithDefaults

`func NewAnchoreSummaryImageTagItemWithDefaults() *AnchoreSummaryImageTagItem`

NewAnchoreSummaryImageTagItemWithDefaults instantiates a new AnchoreSummaryImageTagItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysisStatus

`func (o *AnchoreSummaryImageTagItem) GetAnalysisStatus() string`

GetAnalysisStatus returns the AnalysisStatus field if non-nil, zero value otherwise.

### GetAnalysisStatusOk

`func (o *AnchoreSummaryImageTagItem) GetAnalysisStatusOk() (*string, bool)`

GetAnalysisStatusOk returns a tuple with the AnalysisStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisStatus

`func (o *AnchoreSummaryImageTagItem) SetAnalysisStatus(v string)`

SetAnalysisStatus sets AnalysisStatus field to given value.

### HasAnalysisStatus

`func (o *AnchoreSummaryImageTagItem) HasAnalysisStatus() bool`

HasAnalysisStatus returns a boolean if a field has been set.

### GetImageDigest

`func (o *AnchoreSummaryImageTagItem) GetImageDigest() string`

GetImageDigest returns the ImageDigest field if non-nil, zero value otherwise.

### GetImageDigestOk

`func (o *AnchoreSummaryImageTagItem) GetImageDigestOk() (*string, bool)`

GetImageDigestOk returns a tuple with the ImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDigest

`func (o *AnchoreSummaryImageTagItem) SetImageDigest(v string)`

SetImageDigest sets ImageDigest field to given value.

### HasImageDigest

`func (o *AnchoreSummaryImageTagItem) HasImageDigest() bool`

HasImageDigest returns a boolean if a field has been set.

### GetAnalyzedAt

`func (o *AnchoreSummaryImageTagItem) GetAnalyzedAt() time.Time`

GetAnalyzedAt returns the AnalyzedAt field if non-nil, zero value otherwise.

### GetAnalyzedAtOk

`func (o *AnchoreSummaryImageTagItem) GetAnalyzedAtOk() (*time.Time, bool)`

GetAnalyzedAtOk returns a tuple with the AnalyzedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzedAt

`func (o *AnchoreSummaryImageTagItem) SetAnalyzedAt(v time.Time)`

SetAnalyzedAt sets AnalyzedAt field to given value.

### HasAnalyzedAt

`func (o *AnchoreSummaryImageTagItem) HasAnalyzedAt() bool`

HasAnalyzedAt returns a boolean if a field has been set.

### SetAnalyzedAtNil

`func (o *AnchoreSummaryImageTagItem) SetAnalyzedAtNil(b bool)`

 SetAnalyzedAtNil sets the value for AnalyzedAt to be an explicit nil

### UnsetAnalyzedAt
`func (o *AnchoreSummaryImageTagItem) UnsetAnalyzedAt()`

UnsetAnalyzedAt ensures that no value is present for AnalyzedAt, not even an explicit nil
### GetImageId

`func (o *AnchoreSummaryImageTagItem) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *AnchoreSummaryImageTagItem) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *AnchoreSummaryImageTagItem) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *AnchoreSummaryImageTagItem) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetTagDetectedAt

`func (o *AnchoreSummaryImageTagItem) GetTagDetectedAt() time.Time`

GetTagDetectedAt returns the TagDetectedAt field if non-nil, zero value otherwise.

### GetTagDetectedAtOk

`func (o *AnchoreSummaryImageTagItem) GetTagDetectedAtOk() (*time.Time, bool)`

GetTagDetectedAtOk returns a tuple with the TagDetectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagDetectedAt

`func (o *AnchoreSummaryImageTagItem) SetTagDetectedAt(v time.Time)`

SetTagDetectedAt sets TagDetectedAt field to given value.

### HasTagDetectedAt

`func (o *AnchoreSummaryImageTagItem) HasTagDetectedAt() bool`

HasTagDetectedAt returns a boolean if a field has been set.

### GetInventory

`func (o *AnchoreSummaryImageTagItem) GetInventory() []AnchoreSummaryImageTagItemInventoryInner`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *AnchoreSummaryImageTagItem) GetInventoryOk() (*[]AnchoreSummaryImageTagItemInventoryInner, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *AnchoreSummaryImageTagItem) SetInventory(v []AnchoreSummaryImageTagItemInventoryInner)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *AnchoreSummaryImageTagItem) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### SetInventoryNil

`func (o *AnchoreSummaryImageTagItem) SetInventoryNil(b bool)`

 SetInventoryNil sets the value for Inventory to be an explicit nil

### UnsetInventory
`func (o *AnchoreSummaryImageTagItem) UnsetInventory()`

UnsetInventory ensures that no value is present for Inventory, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


