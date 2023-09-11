# EventSubcategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Events** | Pointer to [**[]EventDescription**](EventDescription.md) |  | [optional] 

## Methods

### NewEventSubcategory

`func NewEventSubcategory() *EventSubcategory`

NewEventSubcategory instantiates a new EventSubcategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventSubcategoryWithDefaults

`func NewEventSubcategoryWithDefaults() *EventSubcategory`

NewEventSubcategoryWithDefaults instantiates a new EventSubcategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EventSubcategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventSubcategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventSubcategory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EventSubcategory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *EventSubcategory) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventSubcategory) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EventSubcategory) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EventSubcategory) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEvents

`func (o *EventSubcategory) GetEvents() []EventDescription`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *EventSubcategory) GetEventsOk() (*[]EventDescription, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *EventSubcategory) SetEvents(v []EventDescription)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *EventSubcategory) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


