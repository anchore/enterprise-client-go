# EventResponseEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to [**EventResponseEventSource**](EventResponseEventSource.md) |  | [optional] 
**Resource** | Pointer to [**EventResponseEventResource**](EventResponseEventResource.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Level** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **interface{}** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewEventResponseEvent

`func NewEventResponseEvent() *EventResponseEvent`

NewEventResponseEvent instantiates a new EventResponseEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventResponseEventWithDefaults

`func NewEventResponseEventWithDefaults() *EventResponseEvent`

NewEventResponseEventWithDefaults instantiates a new EventResponseEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *EventResponseEvent) GetSource() EventResponseEventSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EventResponseEvent) GetSourceOk() (*EventResponseEventSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EventResponseEvent) SetSource(v EventResponseEventSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *EventResponseEvent) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetResource

`func (o *EventResponseEvent) GetResource() EventResponseEventResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *EventResponseEvent) GetResourceOk() (*EventResponseEventResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *EventResponseEvent) SetResource(v EventResponseEventResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *EventResponseEvent) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetType

`func (o *EventResponseEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventResponseEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventResponseEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EventResponseEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCategory

`func (o *EventResponseEvent) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *EventResponseEvent) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *EventResponseEvent) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *EventResponseEvent) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetLevel

`func (o *EventResponseEvent) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *EventResponseEvent) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *EventResponseEvent) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *EventResponseEvent) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetMessage

`func (o *EventResponseEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EventResponseEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EventResponseEvent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *EventResponseEvent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDetails

`func (o *EventResponseEvent) GetDetails() interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *EventResponseEvent) GetDetailsOk() (*interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *EventResponseEvent) SetDetails(v interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *EventResponseEvent) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetTimestamp

`func (o *EventResponseEvent) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EventResponseEvent) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EventResponseEvent) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *EventResponseEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


