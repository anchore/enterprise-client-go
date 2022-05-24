# RegistryConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**LastUpated** | Pointer to **time.Time** |  | [optional] 
**RegistryUser** | Pointer to **string** | Username portion of credential to use for this registry | [optional] 
**RegistryType** | Pointer to **string** | Type of registry | [optional] 
**UserId** | Pointer to **string** | Engine user that owns this registry entry | [optional] 
**Registry** | Pointer to **string** | hostname:port string for accessing the registry, as would be used in a docker pull operation | [optional] 
**RegistryName** | Pointer to **string** | human readable name associated with registry record | [optional] 
**RegistryVerify** | Pointer to **bool** | Use TLS/SSL verification for the registry URL | [optional] 

## Methods

### NewRegistryConfiguration

`func NewRegistryConfiguration() *RegistryConfiguration`

NewRegistryConfiguration instantiates a new RegistryConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryConfigurationWithDefaults

`func NewRegistryConfigurationWithDefaults() *RegistryConfiguration`

NewRegistryConfigurationWithDefaults instantiates a new RegistryConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *RegistryConfiguration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RegistryConfiguration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RegistryConfiguration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RegistryConfiguration) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpated

`func (o *RegistryConfiguration) GetLastUpated() time.Time`

GetLastUpated returns the LastUpated field if non-nil, zero value otherwise.

### GetLastUpatedOk

`func (o *RegistryConfiguration) GetLastUpatedOk() (*time.Time, bool)`

GetLastUpatedOk returns a tuple with the LastUpated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpated

`func (o *RegistryConfiguration) SetLastUpated(v time.Time)`

SetLastUpated sets LastUpated field to given value.

### HasLastUpated

`func (o *RegistryConfiguration) HasLastUpated() bool`

HasLastUpated returns a boolean if a field has been set.

### GetRegistryUser

`func (o *RegistryConfiguration) GetRegistryUser() string`

GetRegistryUser returns the RegistryUser field if non-nil, zero value otherwise.

### GetRegistryUserOk

`func (o *RegistryConfiguration) GetRegistryUserOk() (*string, bool)`

GetRegistryUserOk returns a tuple with the RegistryUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryUser

`func (o *RegistryConfiguration) SetRegistryUser(v string)`

SetRegistryUser sets RegistryUser field to given value.

### HasRegistryUser

`func (o *RegistryConfiguration) HasRegistryUser() bool`

HasRegistryUser returns a boolean if a field has been set.

### GetRegistryType

`func (o *RegistryConfiguration) GetRegistryType() string`

GetRegistryType returns the RegistryType field if non-nil, zero value otherwise.

### GetRegistryTypeOk

`func (o *RegistryConfiguration) GetRegistryTypeOk() (*string, bool)`

GetRegistryTypeOk returns a tuple with the RegistryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryType

`func (o *RegistryConfiguration) SetRegistryType(v string)`

SetRegistryType sets RegistryType field to given value.

### HasRegistryType

`func (o *RegistryConfiguration) HasRegistryType() bool`

HasRegistryType returns a boolean if a field has been set.

### GetUserId

`func (o *RegistryConfiguration) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *RegistryConfiguration) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *RegistryConfiguration) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *RegistryConfiguration) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetRegistry

`func (o *RegistryConfiguration) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *RegistryConfiguration) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *RegistryConfiguration) SetRegistry(v string)`

SetRegistry sets Registry field to given value.

### HasRegistry

`func (o *RegistryConfiguration) HasRegistry() bool`

HasRegistry returns a boolean if a field has been set.

### GetRegistryName

`func (o *RegistryConfiguration) GetRegistryName() string`

GetRegistryName returns the RegistryName field if non-nil, zero value otherwise.

### GetRegistryNameOk

`func (o *RegistryConfiguration) GetRegistryNameOk() (*string, bool)`

GetRegistryNameOk returns a tuple with the RegistryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryName

`func (o *RegistryConfiguration) SetRegistryName(v string)`

SetRegistryName sets RegistryName field to given value.

### HasRegistryName

`func (o *RegistryConfiguration) HasRegistryName() bool`

HasRegistryName returns a boolean if a field has been set.

### GetRegistryVerify

`func (o *RegistryConfiguration) GetRegistryVerify() bool`

GetRegistryVerify returns the RegistryVerify field if non-nil, zero value otherwise.

### GetRegistryVerifyOk

`func (o *RegistryConfiguration) GetRegistryVerifyOk() (*bool, bool)`

GetRegistryVerifyOk returns a tuple with the RegistryVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryVerify

`func (o *RegistryConfiguration) SetRegistryVerify(v bool)`

SetRegistryVerify sets RegistryVerify field to given value.

### HasRegistryVerify

`func (o *RegistryConfiguration) HasRegistryVerify() bool`

HasRegistryVerify returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


