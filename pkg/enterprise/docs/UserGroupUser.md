# UserGroupUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | The username of the user | 
**AddedOn** | Pointer to **time.Time** | The timestamp of when the user was added to the group | [optional] 

## Methods

### NewUserGroupUser

`func NewUserGroupUser(username string, ) *UserGroupUser`

NewUserGroupUser instantiates a new UserGroupUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupUserWithDefaults

`func NewUserGroupUserWithDefaults() *UserGroupUser`

NewUserGroupUserWithDefaults instantiates a new UserGroupUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UserGroupUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserGroupUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserGroupUser) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetAddedOn

`func (o *UserGroupUser) GetAddedOn() time.Time`

GetAddedOn returns the AddedOn field if non-nil, zero value otherwise.

### GetAddedOnOk

`func (o *UserGroupUser) GetAddedOnOk() (*time.Time, bool)`

GetAddedOnOk returns a tuple with the AddedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedOn

`func (o *UserGroupUser) SetAddedOn(v time.Time)`

SetAddedOn sets AddedOn field to given value.

### HasAddedOn

`func (o *UserGroupUser) HasAddedOn() bool`

HasAddedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


