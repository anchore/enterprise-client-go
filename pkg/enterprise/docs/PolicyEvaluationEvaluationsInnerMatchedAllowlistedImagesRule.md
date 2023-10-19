# PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Registry** | **string** |  | 
**Repository** | **string** |  | 
**Image** | [**ImageRef**](ImageRef.md) |  | 
**Description** | Pointer to **string** | Description of the Allowlist or Denylist image match, human readable | [optional] 

## Methods

### NewPolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule

`func NewPolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule(id string, name string, registry string, repository string, image ImageRef, ) *PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule`

NewPolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule instantiates a new PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRuleWithDefaults

`func NewPolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRuleWithDefaults() *PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule`

NewPolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRuleWithDefaults instantiates a new PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule) SetName(v string)`

SetName sets Name field to given value.


### GetRegistry

`func (o *PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule) SetRegistry(v string)`

SetRegistry sets Registry field to given value.


### GetRepository

`func (o *PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule) SetRepository(v string)`

SetRepository sets Repository field to given value.


### GetImage

`func (o *PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule) GetImage() ImageRef`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule) GetImageOk() (*ImageRef, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule) SetImage(v ImageRef)`

SetImage sets Image field to given value.


### GetDescription

`func (o *PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


