# CreateSbomGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Version** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateSbomGroupRequest

`func NewCreateSbomGroupRequest(name string, version string, ) *CreateSbomGroupRequest`

NewCreateSbomGroupRequest instantiates a new CreateSbomGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSbomGroupRequestWithDefaults

`func NewCreateSbomGroupRequestWithDefaults() *CreateSbomGroupRequest`

NewCreateSbomGroupRequestWithDefaults instantiates a new CreateSbomGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateSbomGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSbomGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSbomGroupRequest) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *CreateSbomGroupRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CreateSbomGroupRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CreateSbomGroupRequest) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetDescription

`func (o *CreateSbomGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateSbomGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateSbomGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateSbomGroupRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


