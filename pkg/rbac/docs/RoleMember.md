# RoleMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**ForAccount** | **string** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRoleMember

`func NewRoleMember(username string, forAccount string, ) *RoleMember`

NewRoleMember instantiates a new RoleMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleMemberWithDefaults

`func NewRoleMemberWithDefaults() *RoleMember`

NewRoleMemberWithDefaults instantiates a new RoleMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *RoleMember) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RoleMember) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RoleMember) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetForAccount

`func (o *RoleMember) GetForAccount() string`

GetForAccount returns the ForAccount field if non-nil, zero value otherwise.

### GetForAccountOk

`func (o *RoleMember) GetForAccountOk() (*string, bool)`

GetForAccountOk returns a tuple with the ForAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForAccount

`func (o *RoleMember) SetForAccount(v string)`

SetForAccount sets ForAccount field to given value.


### GetCreatedAt

`func (o *RoleMember) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RoleMember) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RoleMember) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RoleMember) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


