# UserCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | The username to create | 
**Password** | Pointer to **string** | The initial password for the user, must be at least 6 characters, up to 128. This must be null when the user_type is not &#39;native&#39;. | [optional] 
**UserType** | Pointer to **string** | The user&#39;s type. A Native user authenticates using user/password log on. All other users will authenticate with an IDP. | [optional] 
**IdpName** | Pointer to **string** | If the user is authenticating via an IDP, this is the name of the IDP. A &#39;native&#39; user should have this set to null. | [optional] 

## Methods

### NewUserCreationRequest

`func NewUserCreationRequest(username string, ) *UserCreationRequest`

NewUserCreationRequest instantiates a new UserCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCreationRequestWithDefaults

`func NewUserCreationRequestWithDefaults() *UserCreationRequest`

NewUserCreationRequestWithDefaults instantiates a new UserCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UserCreationRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserCreationRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserCreationRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *UserCreationRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserCreationRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserCreationRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserCreationRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUserType

`func (o *UserCreationRequest) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *UserCreationRequest) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *UserCreationRequest) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *UserCreationRequest) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetIdpName

`func (o *UserCreationRequest) GetIdpName() string`

GetIdpName returns the IdpName field if non-nil, zero value otherwise.

### GetIdpNameOk

`func (o *UserCreationRequest) GetIdpNameOk() (*string, bool)`

GetIdpNameOk returns a tuple with the IdpName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpName

`func (o *UserCreationRequest) SetIdpName(v string)`

SetIdpName sets IdpName field to given value.

### HasIdpName

`func (o *UserCreationRequest) HasIdpName() bool`

HasIdpName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


