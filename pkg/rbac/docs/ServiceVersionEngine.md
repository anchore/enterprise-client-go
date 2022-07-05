# ServiceVersionEngine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | Version of the installed engine library | [optional] 
**Db** | Pointer to **string** | Version of the installed engine db schema | [optional] 

## Methods

### NewServiceVersionEngine

`func NewServiceVersionEngine() *ServiceVersionEngine`

NewServiceVersionEngine instantiates a new ServiceVersionEngine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceVersionEngineWithDefaults

`func NewServiceVersionEngineWithDefaults() *ServiceVersionEngine`

NewServiceVersionEngineWithDefaults instantiates a new ServiceVersionEngine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ServiceVersionEngine) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ServiceVersionEngine) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ServiceVersionEngine) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ServiceVersionEngine) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDb

`func (o *ServiceVersionEngine) GetDb() string`

GetDb returns the Db field if non-nil, zero value otherwise.

### GetDbOk

`func (o *ServiceVersionEngine) GetDbOk() (*string, bool)`

GetDbOk returns a tuple with the Db field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDb

`func (o *ServiceVersionEngine) SetDb(v string)`

SetDb sets Db field to given value.

### HasDb

`func (o *ServiceVersionEngine) HasDb() bool`

HasDb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


