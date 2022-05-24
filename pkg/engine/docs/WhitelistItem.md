# WhitelistItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Gate** | **string** |  | 
**TriggerId** | **string** |  | 
**ExpiresOn** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewWhitelistItem

`func NewWhitelistItem(gate string, triggerId string, ) *WhitelistItem`

NewWhitelistItem instantiates a new WhitelistItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhitelistItemWithDefaults

`func NewWhitelistItemWithDefaults() *WhitelistItem`

NewWhitelistItemWithDefaults instantiates a new WhitelistItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WhitelistItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WhitelistItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WhitelistItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WhitelistItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGate

`func (o *WhitelistItem) GetGate() string`

GetGate returns the Gate field if non-nil, zero value otherwise.

### GetGateOk

`func (o *WhitelistItem) GetGateOk() (*string, bool)`

GetGateOk returns a tuple with the Gate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGate

`func (o *WhitelistItem) SetGate(v string)`

SetGate sets Gate field to given value.


### GetTriggerId

`func (o *WhitelistItem) GetTriggerId() string`

GetTriggerId returns the TriggerId field if non-nil, zero value otherwise.

### GetTriggerIdOk

`func (o *WhitelistItem) GetTriggerIdOk() (*string, bool)`

GetTriggerIdOk returns a tuple with the TriggerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerId

`func (o *WhitelistItem) SetTriggerId(v string)`

SetTriggerId sets TriggerId field to given value.


### GetExpiresOn

`func (o *WhitelistItem) GetExpiresOn() time.Time`

GetExpiresOn returns the ExpiresOn field if non-nil, zero value otherwise.

### GetExpiresOnOk

`func (o *WhitelistItem) GetExpiresOnOk() (*time.Time, bool)`

GetExpiresOnOk returns a tuple with the ExpiresOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresOn

`func (o *WhitelistItem) SetExpiresOn(v time.Time)`

SetExpiresOn sets ExpiresOn field to given value.

### HasExpiresOn

`func (o *WhitelistItem) HasExpiresOn() bool`

HasExpiresOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


