# EventCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Subcategories** | Pointer to [**[]EventSubcategory**](EventSubcategory.md) |  | [optional] 

## Methods

### NewEventCategory

`func NewEventCategory() *EventCategory`

NewEventCategory instantiates a new EventCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventCategoryWithDefaults

`func NewEventCategoryWithDefaults() *EventCategory`

NewEventCategoryWithDefaults instantiates a new EventCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *EventCategory) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *EventCategory) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *EventCategory) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *EventCategory) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *EventCategory) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventCategory) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EventCategory) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EventCategory) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSubcategories

`func (o *EventCategory) GetSubcategories() []EventSubcategory`

GetSubcategories returns the Subcategories field if non-nil, zero value otherwise.

### GetSubcategoriesOk

`func (o *EventCategory) GetSubcategoriesOk() (*[]EventSubcategory, bool)`

GetSubcategoriesOk returns a tuple with the Subcategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubcategories

`func (o *EventCategory) SetSubcategories(v []EventSubcategory)`

SetSubcategories sets Subcategories field to given value.

### HasSubcategories

`func (o *EventCategory) HasSubcategories() bool`

HasSubcategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


