# AccountCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The account name to use. This will identify the account and must be globally unique in the system. | 
**Email** | Pointer to **string** | An optional email to associate with the account for contact purposes | [optional] 

## Methods

### NewAccountCreationRequest

`func NewAccountCreationRequest(name string, ) *AccountCreationRequest`

NewAccountCreationRequest instantiates a new AccountCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCreationRequestWithDefaults

`func NewAccountCreationRequestWithDefaults() *AccountCreationRequest`

NewAccountCreationRequestWithDefaults instantiates a new AccountCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AccountCreationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountCreationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountCreationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *AccountCreationRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AccountCreationRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AccountCreationRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AccountCreationRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


