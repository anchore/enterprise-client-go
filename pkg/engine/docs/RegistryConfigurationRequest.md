# RegistryConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegistryUser** | Pointer to **string** | Username portion of credential to use for this registry | [optional] 
**RegistryPass** | Pointer to **string** | Password portion of credential to use for this registry | [optional] 
**RegistryType** | Pointer to **string** | Type of registry | [optional] 
**Registry** | Pointer to **string** | hostname:port string for accessing the registry, as would be used in a docker pull operation. May include some or all of a repository and wildcards (e.g. docker.io/library/_* or gcr.io/myproject/myrepository) | [optional] 
**RegistryName** | Pointer to **string** | human readable name associated with registry record | [optional] 
**RegistryVerify** | Pointer to **bool** | Use TLS/SSL verification for the registry URL | [optional] 

## Methods

### NewRegistryConfigurationRequest

`func NewRegistryConfigurationRequest() *RegistryConfigurationRequest`

NewRegistryConfigurationRequest instantiates a new RegistryConfigurationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryConfigurationRequestWithDefaults

`func NewRegistryConfigurationRequestWithDefaults() *RegistryConfigurationRequest`

NewRegistryConfigurationRequestWithDefaults instantiates a new RegistryConfigurationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistryUser

`func (o *RegistryConfigurationRequest) GetRegistryUser() string`

GetRegistryUser returns the RegistryUser field if non-nil, zero value otherwise.

### GetRegistryUserOk

`func (o *RegistryConfigurationRequest) GetRegistryUserOk() (*string, bool)`

GetRegistryUserOk returns a tuple with the RegistryUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryUser

`func (o *RegistryConfigurationRequest) SetRegistryUser(v string)`

SetRegistryUser sets RegistryUser field to given value.

### HasRegistryUser

`func (o *RegistryConfigurationRequest) HasRegistryUser() bool`

HasRegistryUser returns a boolean if a field has been set.

### GetRegistryPass

`func (o *RegistryConfigurationRequest) GetRegistryPass() string`

GetRegistryPass returns the RegistryPass field if non-nil, zero value otherwise.

### GetRegistryPassOk

`func (o *RegistryConfigurationRequest) GetRegistryPassOk() (*string, bool)`

GetRegistryPassOk returns a tuple with the RegistryPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryPass

`func (o *RegistryConfigurationRequest) SetRegistryPass(v string)`

SetRegistryPass sets RegistryPass field to given value.

### HasRegistryPass

`func (o *RegistryConfigurationRequest) HasRegistryPass() bool`

HasRegistryPass returns a boolean if a field has been set.

### GetRegistryType

`func (o *RegistryConfigurationRequest) GetRegistryType() string`

GetRegistryType returns the RegistryType field if non-nil, zero value otherwise.

### GetRegistryTypeOk

`func (o *RegistryConfigurationRequest) GetRegistryTypeOk() (*string, bool)`

GetRegistryTypeOk returns a tuple with the RegistryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryType

`func (o *RegistryConfigurationRequest) SetRegistryType(v string)`

SetRegistryType sets RegistryType field to given value.

### HasRegistryType

`func (o *RegistryConfigurationRequest) HasRegistryType() bool`

HasRegistryType returns a boolean if a field has been set.

### GetRegistry

`func (o *RegistryConfigurationRequest) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *RegistryConfigurationRequest) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *RegistryConfigurationRequest) SetRegistry(v string)`

SetRegistry sets Registry field to given value.

### HasRegistry

`func (o *RegistryConfigurationRequest) HasRegistry() bool`

HasRegistry returns a boolean if a field has been set.

### GetRegistryName

`func (o *RegistryConfigurationRequest) GetRegistryName() string`

GetRegistryName returns the RegistryName field if non-nil, zero value otherwise.

### GetRegistryNameOk

`func (o *RegistryConfigurationRequest) GetRegistryNameOk() (*string, bool)`

GetRegistryNameOk returns a tuple with the RegistryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryName

`func (o *RegistryConfigurationRequest) SetRegistryName(v string)`

SetRegistryName sets RegistryName field to given value.

### HasRegistryName

`func (o *RegistryConfigurationRequest) HasRegistryName() bool`

HasRegistryName returns a boolean if a field has been set.

### GetRegistryVerify

`func (o *RegistryConfigurationRequest) GetRegistryVerify() bool`

GetRegistryVerify returns the RegistryVerify field if non-nil, zero value otherwise.

### GetRegistryVerifyOk

`func (o *RegistryConfigurationRequest) GetRegistryVerifyOk() (*bool, bool)`

GetRegistryVerifyOk returns a tuple with the RegistryVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryVerify

`func (o *RegistryConfigurationRequest) SetRegistryVerify(v bool)`

SetRegistryVerify sets RegistryVerify field to given value.

### HasRegistryVerify

`func (o *RegistryConfigurationRequest) HasRegistryVerify() bool`

HasRegistryVerify returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


