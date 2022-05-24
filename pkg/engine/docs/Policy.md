# Policy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Version** | **string** |  | 
**ArtifactType** | Pointer to **string** |  | [optional] 
**Rules** | Pointer to [**[]PolicyRule**](PolicyRule.md) |  | [optional] 

## Methods

### NewPolicy

`func NewPolicy(id string, version string, ) *Policy`

NewPolicy instantiates a new Policy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyWithDefaults

`func NewPolicyWithDefaults() *Policy`

NewPolicyWithDefaults instantiates a new Policy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Policy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Policy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Policy) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Policy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Policy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Policy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Policy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetComment

`func (o *Policy) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Policy) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Policy) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Policy) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetVersion

`func (o *Policy) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Policy) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Policy) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetArtifactType

`func (o *Policy) GetArtifactType() string`

GetArtifactType returns the ArtifactType field if non-nil, zero value otherwise.

### GetArtifactTypeOk

`func (o *Policy) GetArtifactTypeOk() (*string, bool)`

GetArtifactTypeOk returns a tuple with the ArtifactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactType

`func (o *Policy) SetArtifactType(v string)`

SetArtifactType sets ArtifactType field to given value.

### HasArtifactType

`func (o *Policy) HasArtifactType() bool`

HasArtifactType returns a boolean if a field has been set.

### GetRules

`func (o *Policy) GetRules() []PolicyRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *Policy) GetRulesOk() (*[]PolicyRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *Policy) SetRules(v []PolicyRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *Policy) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


