# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The account identifier, not updatable after creation | 
**Type** | Pointer to **string** | The user type (admin vs user). If not specified in a POST request, &#39;user&#39; is default | [optional] 
**State** | Pointer to **string** | State of the account. Disabled accounts prevent member users from logging in, deleting accounts are disabled and pending deletion and will be removed once all owned resources are garbage collected by the system | [optional] 
**Email** | Pointer to **string** | Optional email address associated with the account | [optional] 
**CreatedAt** | Pointer to **time.Time** | The timestamp when the account was created | [optional] 
**LastUpdated** | Pointer to **time.Time** | The timestamp of the last update to the account metadata itself (not users or creds) | [optional] 

## Methods

### NewAccount

`func NewAccount(name string, ) *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Account) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Account) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Account) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *Account) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Account) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Account) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Account) HasType() bool`

HasType returns a boolean if a field has been set.

### GetState

`func (o *Account) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Account) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Account) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Account) HasState() bool`

HasState returns a boolean if a field has been set.

### GetEmail

`func (o *Account) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Account) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Account) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Account) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Account) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Account) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Account) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Account) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Account) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Account) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Account) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Account) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


