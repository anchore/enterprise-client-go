# RbacManagerSamlConfigurationGetAllOfUserGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserGroup** | Pointer to **string** | The name of the user group | [optional] 
**AddedOn** | Pointer to **time.Time** | The date and time the user group was associated with the IDP | [optional] 

## Methods

### NewRbacManagerSamlConfigurationGetAllOfUserGroups

`func NewRbacManagerSamlConfigurationGetAllOfUserGroups() *RbacManagerSamlConfigurationGetAllOfUserGroups`

NewRbacManagerSamlConfigurationGetAllOfUserGroups instantiates a new RbacManagerSamlConfigurationGetAllOfUserGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRbacManagerSamlConfigurationGetAllOfUserGroupsWithDefaults

`func NewRbacManagerSamlConfigurationGetAllOfUserGroupsWithDefaults() *RbacManagerSamlConfigurationGetAllOfUserGroups`

NewRbacManagerSamlConfigurationGetAllOfUserGroupsWithDefaults instantiates a new RbacManagerSamlConfigurationGetAllOfUserGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserGroup

`func (o *RbacManagerSamlConfigurationGetAllOfUserGroups) GetUserGroup() string`

GetUserGroup returns the UserGroup field if non-nil, zero value otherwise.

### GetUserGroupOk

`func (o *RbacManagerSamlConfigurationGetAllOfUserGroups) GetUserGroupOk() (*string, bool)`

GetUserGroupOk returns a tuple with the UserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroup

`func (o *RbacManagerSamlConfigurationGetAllOfUserGroups) SetUserGroup(v string)`

SetUserGroup sets UserGroup field to given value.

### HasUserGroup

`func (o *RbacManagerSamlConfigurationGetAllOfUserGroups) HasUserGroup() bool`

HasUserGroup returns a boolean if a field has been set.

### GetAddedOn

`func (o *RbacManagerSamlConfigurationGetAllOfUserGroups) GetAddedOn() time.Time`

GetAddedOn returns the AddedOn field if non-nil, zero value otherwise.

### GetAddedOnOk

`func (o *RbacManagerSamlConfigurationGetAllOfUserGroups) GetAddedOnOk() (*time.Time, bool)`

GetAddedOnOk returns a tuple with the AddedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedOn

`func (o *RbacManagerSamlConfigurationGetAllOfUserGroups) SetAddedOn(v time.Time)`

SetAddedOn sets AddedOn field to given value.

### HasAddedOn

`func (o *RbacManagerSamlConfigurationGetAllOfUserGroups) HasAddedOn() bool`

HasAddedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


