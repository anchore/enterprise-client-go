# Allowlist

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Version** | **string** |  | 
**Description** | Pointer to **string** | Description of the Allowlist, human readable | [optional] 
**Items** | Pointer to [**[]AllowlistItem**](AllowlistItem.md) |  | [optional] 

## Methods

### NewAllowlist

`func NewAllowlist(id string, version string, ) *Allowlist`

NewAllowlist instantiates a new Allowlist object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowlistWithDefaults

`func NewAllowlistWithDefaults() *Allowlist`

NewAllowlistWithDefaults instantiates a new Allowlist object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Allowlist) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Allowlist) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Allowlist) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Allowlist) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Allowlist) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Allowlist) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Allowlist) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *Allowlist) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Allowlist) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Allowlist) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetDescription

`func (o *Allowlist) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Allowlist) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Allowlist) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Allowlist) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetItems

`func (o *Allowlist) GetItems() []AllowlistItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Allowlist) GetItemsOk() (*[]AllowlistItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Allowlist) SetItems(v []AllowlistItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *Allowlist) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


