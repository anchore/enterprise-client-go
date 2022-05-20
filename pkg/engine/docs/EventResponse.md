# EventResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeneratedUuid** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Event** | Pointer to [**EventResponseEvent**](EventResponseEvent.md) |  | [optional] 

## Methods

### NewEventResponse

`func NewEventResponse() *EventResponse`

NewEventResponse instantiates a new EventResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventResponseWithDefaults

`func NewEventResponseWithDefaults() *EventResponse`

NewEventResponseWithDefaults instantiates a new EventResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneratedUuid

`func (o *EventResponse) GetGeneratedUuid() string`

GetGeneratedUuid returns the GeneratedUuid field if non-nil, zero value otherwise.

### GetGeneratedUuidOk

`func (o *EventResponse) GetGeneratedUuidOk() (*string, bool)`

GetGeneratedUuidOk returns a tuple with the GeneratedUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedUuid

`func (o *EventResponse) SetGeneratedUuid(v string)`

SetGeneratedUuid sets GeneratedUuid field to given value.

### HasGeneratedUuid

`func (o *EventResponse) HasGeneratedUuid() bool`

HasGeneratedUuid returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EventResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EventResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EventResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EventResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEvent

`func (o *EventResponse) GetEvent() EventResponseEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventResponse) GetEventOk() (*EventResponseEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventResponse) SetEvent(v EventResponseEvent)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *EventResponse) HasEvent() bool`

HasEvent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


