# ImageSelectionRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Registry** | **string** |  | 
**Repository** | **string** |  | 
**Image** | [**ImageRef**](ImageRef.md) |  | 
**Description** | Pointer to **string** | Description of the Allowlist or Denylist image match, human readable | [optional] 

## Methods

### NewImageSelectionRule

`func NewImageSelectionRule(name string, registry string, repository string, image ImageRef, ) *ImageSelectionRule`

NewImageSelectionRule instantiates a new ImageSelectionRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageSelectionRuleWithDefaults

`func NewImageSelectionRuleWithDefaults() *ImageSelectionRule`

NewImageSelectionRuleWithDefaults instantiates a new ImageSelectionRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImageSelectionRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageSelectionRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageSelectionRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ImageSelectionRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ImageSelectionRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageSelectionRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageSelectionRule) SetName(v string)`

SetName sets Name field to given value.


### GetRegistry

`func (o *ImageSelectionRule) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *ImageSelectionRule) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *ImageSelectionRule) SetRegistry(v string)`

SetRegistry sets Registry field to given value.


### GetRepository

`func (o *ImageSelectionRule) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *ImageSelectionRule) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *ImageSelectionRule) SetRepository(v string)`

SetRepository sets Repository field to given value.


### GetImage

`func (o *ImageSelectionRule) GetImage() ImageRef`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ImageSelectionRule) GetImageOk() (*ImageRef, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ImageSelectionRule) SetImage(v ImageRef)`

SetImage sets Image field to given value.


### GetDescription

`func (o *ImageSelectionRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ImageSelectionRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ImageSelectionRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ImageSelectionRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


