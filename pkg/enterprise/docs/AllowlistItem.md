# AllowlistItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Gate** | **string** |  | 
**TriggerId** | **string** |  | 
**ExpiresOn** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** | Description of the Allowlist item, human readable | [optional] 

## Methods

### NewAllowlistItem

`func NewAllowlistItem(gate string, triggerId string, ) *AllowlistItem`

NewAllowlistItem instantiates a new AllowlistItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowlistItemWithDefaults

`func NewAllowlistItemWithDefaults() *AllowlistItem`

NewAllowlistItemWithDefaults instantiates a new AllowlistItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AllowlistItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AllowlistItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AllowlistItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AllowlistItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGate

`func (o *AllowlistItem) GetGate() string`

GetGate returns the Gate field if non-nil, zero value otherwise.

### GetGateOk

`func (o *AllowlistItem) GetGateOk() (*string, bool)`

GetGateOk returns a tuple with the Gate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGate

`func (o *AllowlistItem) SetGate(v string)`

SetGate sets Gate field to given value.


### GetTriggerId

`func (o *AllowlistItem) GetTriggerId() string`

GetTriggerId returns the TriggerId field if non-nil, zero value otherwise.

### GetTriggerIdOk

`func (o *AllowlistItem) GetTriggerIdOk() (*string, bool)`

GetTriggerIdOk returns a tuple with the TriggerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerId

`func (o *AllowlistItem) SetTriggerId(v string)`

SetTriggerId sets TriggerId field to given value.


### GetExpiresOn

`func (o *AllowlistItem) GetExpiresOn() time.Time`

GetExpiresOn returns the ExpiresOn field if non-nil, zero value otherwise.

### GetExpiresOnOk

`func (o *AllowlistItem) GetExpiresOnOk() (*time.Time, bool)`

GetExpiresOnOk returns a tuple with the ExpiresOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresOn

`func (o *AllowlistItem) SetExpiresOn(v time.Time)`

SetExpiresOn sets ExpiresOn field to given value.

### HasExpiresOn

`func (o *AllowlistItem) HasExpiresOn() bool`

HasExpiresOn returns a boolean if a field has been set.

### GetDescription

`func (o *AllowlistItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AllowlistItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AllowlistItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AllowlistItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


