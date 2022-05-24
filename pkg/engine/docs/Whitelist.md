# Whitelist

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Version** | **string** |  | 
**Comment** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]WhitelistItem**](WhitelistItem.md) |  | [optional] 

## Methods

### NewWhitelist

`func NewWhitelist(id string, version string, ) *Whitelist`

NewWhitelist instantiates a new Whitelist object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhitelistWithDefaults

`func NewWhitelistWithDefaults() *Whitelist`

NewWhitelistWithDefaults instantiates a new Whitelist object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Whitelist) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Whitelist) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Whitelist) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Whitelist) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Whitelist) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Whitelist) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Whitelist) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *Whitelist) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Whitelist) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Whitelist) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetComment

`func (o *Whitelist) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Whitelist) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Whitelist) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Whitelist) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetItems

`func (o *Whitelist) GetItems() []WhitelistItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Whitelist) GetItemsOk() (*[]WhitelistItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Whitelist) SetItems(v []WhitelistItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *Whitelist) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


