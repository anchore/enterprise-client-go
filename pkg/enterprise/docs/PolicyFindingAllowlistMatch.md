# PolicyFindingAllowlistMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the allowlist that matched this finding | [optional] 
**Name** | Pointer to **string** | Name of the allowlist that matched this finding | [optional] 
**MatchedRuleId** | Pointer to **string** | ID of the rule within the allowlist that matched this finding | [optional] 

## Methods

### NewPolicyFindingAllowlistMatch

`func NewPolicyFindingAllowlistMatch() *PolicyFindingAllowlistMatch`

NewPolicyFindingAllowlistMatch instantiates a new PolicyFindingAllowlistMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyFindingAllowlistMatchWithDefaults

`func NewPolicyFindingAllowlistMatchWithDefaults() *PolicyFindingAllowlistMatch`

NewPolicyFindingAllowlistMatchWithDefaults instantiates a new PolicyFindingAllowlistMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PolicyFindingAllowlistMatch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyFindingAllowlistMatch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyFindingAllowlistMatch) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PolicyFindingAllowlistMatch) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PolicyFindingAllowlistMatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyFindingAllowlistMatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyFindingAllowlistMatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyFindingAllowlistMatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMatchedRuleId

`func (o *PolicyFindingAllowlistMatch) GetMatchedRuleId() string`

GetMatchedRuleId returns the MatchedRuleId field if non-nil, zero value otherwise.

### GetMatchedRuleIdOk

`func (o *PolicyFindingAllowlistMatch) GetMatchedRuleIdOk() (*string, bool)`

GetMatchedRuleIdOk returns a tuple with the MatchedRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedRuleId

`func (o *PolicyFindingAllowlistMatch) SetMatchedRuleId(v string)`

SetMatchedRuleId sets MatchedRuleId field to given value.

### HasMatchedRuleId

`func (o *PolicyFindingAllowlistMatch) HasMatchedRuleId() bool`

HasMatchedRuleId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


