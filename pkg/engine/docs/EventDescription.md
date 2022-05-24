# EventDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The event type. The last component of the fully-qualified event_type (category.subcategory.event) | [optional] 
**Type** | Pointer to **string** | The fully qualified event type as would be seen in the event payload | [optional] 
**Message** | Pointer to **string** | The message associated with the event type | [optional] 
**ResourceType** | Pointer to **string** | The type of resource this event is generated from | [optional] 

## Methods

### NewEventDescription

`func NewEventDescription() *EventDescription`

NewEventDescription instantiates a new EventDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventDescriptionWithDefaults

`func NewEventDescriptionWithDefaults() *EventDescription`

NewEventDescriptionWithDefaults instantiates a new EventDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EventDescription) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventDescription) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventDescription) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EventDescription) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *EventDescription) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventDescription) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventDescription) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EventDescription) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMessage

`func (o *EventDescription) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EventDescription) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EventDescription) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *EventDescription) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResourceType

`func (o *EventDescription) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *EventDescription) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *EventDescription) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *EventDescription) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


