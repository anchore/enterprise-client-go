# RbacManagerPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The allowed action. e.g. getImage | [optional] 
**Target** | Pointer to **string** | The target to which the action may be applied. Either a &#39;*&#39; for all or a specific target id | [optional] 

## Methods

### NewRbacManagerPermission

`func NewRbacManagerPermission() *RbacManagerPermission`

NewRbacManagerPermission instantiates a new RbacManagerPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRbacManagerPermissionWithDefaults

`func NewRbacManagerPermissionWithDefaults() *RbacManagerPermission`

NewRbacManagerPermissionWithDefaults instantiates a new RbacManagerPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *RbacManagerPermission) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RbacManagerPermission) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RbacManagerPermission) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *RbacManagerPermission) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetTarget

`func (o *RbacManagerPermission) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *RbacManagerPermission) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *RbacManagerPermission) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *RbacManagerPermission) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


