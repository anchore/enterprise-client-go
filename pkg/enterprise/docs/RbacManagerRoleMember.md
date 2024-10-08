# RbacManagerRoleMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**ForAccount** | Pointer to **string** | Deprecated. Please use domain_name instead. | [optional] 
**DomainName** | Pointer to **string** | The domain scope that applies to the set of roles. This will be the account name if the domain scope is an account. | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRbacManagerRoleMember

`func NewRbacManagerRoleMember(username string, ) *RbacManagerRoleMember`

NewRbacManagerRoleMember instantiates a new RbacManagerRoleMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRbacManagerRoleMemberWithDefaults

`func NewRbacManagerRoleMemberWithDefaults() *RbacManagerRoleMember`

NewRbacManagerRoleMemberWithDefaults instantiates a new RbacManagerRoleMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *RbacManagerRoleMember) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RbacManagerRoleMember) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RbacManagerRoleMember) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetForAccount

`func (o *RbacManagerRoleMember) GetForAccount() string`

GetForAccount returns the ForAccount field if non-nil, zero value otherwise.

### GetForAccountOk

`func (o *RbacManagerRoleMember) GetForAccountOk() (*string, bool)`

GetForAccountOk returns a tuple with the ForAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForAccount

`func (o *RbacManagerRoleMember) SetForAccount(v string)`

SetForAccount sets ForAccount field to given value.

### HasForAccount

`func (o *RbacManagerRoleMember) HasForAccount() bool`

HasForAccount returns a boolean if a field has been set.

### GetDomainName

`func (o *RbacManagerRoleMember) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *RbacManagerRoleMember) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *RbacManagerRoleMember) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *RbacManagerRoleMember) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RbacManagerRoleMember) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RbacManagerRoleMember) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RbacManagerRoleMember) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RbacManagerRoleMember) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


