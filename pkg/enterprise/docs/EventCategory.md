# EventCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
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

### GetName

`func (o *EventCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventCategory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EventCategory) HasName() bool`

HasName returns a boolean if a field has been set.

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


