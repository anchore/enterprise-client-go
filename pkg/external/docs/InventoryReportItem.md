# InventoryReportItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | Pointer to **string** |  | [optional] 
**Images** | Pointer to [**[]InventoryReportImage**](InventoryReportImage.md) |  | [optional] 

## Methods

### NewInventoryReportItem

`func NewInventoryReportItem() *InventoryReportItem`

NewInventoryReportItem instantiates a new InventoryReportItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryReportItemWithDefaults

`func NewInventoryReportItemWithDefaults() *InventoryReportItem`

NewInventoryReportItemWithDefaults instantiates a new InventoryReportItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *InventoryReportItem) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *InventoryReportItem) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *InventoryReportItem) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *InventoryReportItem) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetImages

`func (o *InventoryReportItem) GetImages() []InventoryReportImage`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *InventoryReportItem) GetImagesOk() (*[]InventoryReportImage, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *InventoryReportItem) SetImages(v []InventoryReportImage)`

SetImages sets Images field to given value.

### HasImages

`func (o *InventoryReportItem) HasImages() bool`

HasImages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


