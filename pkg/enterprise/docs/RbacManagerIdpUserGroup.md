# RbacManagerIdpUserGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserGroupUuid** | **string** | The UUID of the user group | 
**UserGroupName** | Pointer to **string** | The name of the user group | [optional] 
**MappedOn** | Pointer to **time.Time** | The timestamp when the user group was mapped to the IdP | [optional] 

## Methods

### NewRbacManagerIdpUserGroup

`func NewRbacManagerIdpUserGroup(userGroupUuid string, ) *RbacManagerIdpUserGroup`

NewRbacManagerIdpUserGroup instantiates a new RbacManagerIdpUserGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRbacManagerIdpUserGroupWithDefaults

`func NewRbacManagerIdpUserGroupWithDefaults() *RbacManagerIdpUserGroup`

NewRbacManagerIdpUserGroupWithDefaults instantiates a new RbacManagerIdpUserGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserGroupUuid

`func (o *RbacManagerIdpUserGroup) GetUserGroupUuid() string`

GetUserGroupUuid returns the UserGroupUuid field if non-nil, zero value otherwise.

### GetUserGroupUuidOk

`func (o *RbacManagerIdpUserGroup) GetUserGroupUuidOk() (*string, bool)`

GetUserGroupUuidOk returns a tuple with the UserGroupUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupUuid

`func (o *RbacManagerIdpUserGroup) SetUserGroupUuid(v string)`

SetUserGroupUuid sets UserGroupUuid field to given value.


### GetUserGroupName

`func (o *RbacManagerIdpUserGroup) GetUserGroupName() string`

GetUserGroupName returns the UserGroupName field if non-nil, zero value otherwise.

### GetUserGroupNameOk

`func (o *RbacManagerIdpUserGroup) GetUserGroupNameOk() (*string, bool)`

GetUserGroupNameOk returns a tuple with the UserGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupName

`func (o *RbacManagerIdpUserGroup) SetUserGroupName(v string)`

SetUserGroupName sets UserGroupName field to given value.

### HasUserGroupName

`func (o *RbacManagerIdpUserGroup) HasUserGroupName() bool`

HasUserGroupName returns a boolean if a field has been set.

### GetMappedOn

`func (o *RbacManagerIdpUserGroup) GetMappedOn() time.Time`

GetMappedOn returns the MappedOn field if non-nil, zero value otherwise.

### GetMappedOnOk

`func (o *RbacManagerIdpUserGroup) GetMappedOnOk() (*time.Time, bool)`

GetMappedOnOk returns a tuple with the MappedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedOn

`func (o *RbacManagerIdpUserGroup) SetMappedOn(v time.Time)`

SetMappedOn sets MappedOn field to given value.

### HasMappedOn

`func (o *RbacManagerIdpUserGroup) HasMappedOn() bool`

HasMappedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


