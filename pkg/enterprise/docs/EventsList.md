# EventsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]EventResponse**](EventResponse.md) | List of events | [optional] 
**NextPage** | Pointer to **bool** | Boolean flag, True indicates there are more events and False otherwise | [optional] 
**ItemCount** | Pointer to **int64** | Number of events in this page | [optional] 
**Page** | Pointer to **int64** | Page number of this result set | [optional] 

## Methods

### NewEventsList

`func NewEventsList() *EventsList`

NewEventsList instantiates a new EventsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsListWithDefaults

`func NewEventsListWithDefaults() *EventsList`

NewEventsListWithDefaults instantiates a new EventsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *EventsList) GetResults() []EventResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *EventsList) GetResultsOk() (*[]EventResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *EventsList) SetResults(v []EventResponse)`

SetResults sets Results field to given value.

### HasResults

`func (o *EventsList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetNextPage

`func (o *EventsList) GetNextPage() bool`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *EventsList) GetNextPageOk() (*bool, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *EventsList) SetNextPage(v bool)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *EventsList) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.

### GetItemCount

`func (o *EventsList) GetItemCount() int64`

GetItemCount returns the ItemCount field if non-nil, zero value otherwise.

### GetItemCountOk

`func (o *EventsList) GetItemCountOk() (*int64, bool)`

GetItemCountOk returns a tuple with the ItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCount

`func (o *EventsList) SetItemCount(v int64)`

SetItemCount sets ItemCount field to given value.

### HasItemCount

`func (o *EventsList) HasItemCount() bool`

HasItemCount returns a boolean if a field has been set.

### GetPage

`func (o *EventsList) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *EventsList) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *EventsList) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *EventsList) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


