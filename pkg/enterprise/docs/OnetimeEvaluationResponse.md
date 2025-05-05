# OnetimeEvaluationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vulnerabilities** | Pointer to [**[]PackageVulnerability**](PackageVulnerability.md) |  | [optional] 
**Evaluation** | Pointer to [**PolicyEvaluationResult**](PolicyEvaluationResult.md) |  | [optional] 

## Methods

### NewOnetimeEvaluationResponse

`func NewOnetimeEvaluationResponse() *OnetimeEvaluationResponse`

NewOnetimeEvaluationResponse instantiates a new OnetimeEvaluationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnetimeEvaluationResponseWithDefaults

`func NewOnetimeEvaluationResponseWithDefaults() *OnetimeEvaluationResponse`

NewOnetimeEvaluationResponseWithDefaults instantiates a new OnetimeEvaluationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVulnerabilities

`func (o *OnetimeEvaluationResponse) GetVulnerabilities() []PackageVulnerability`

GetVulnerabilities returns the Vulnerabilities field if non-nil, zero value otherwise.

### GetVulnerabilitiesOk

`func (o *OnetimeEvaluationResponse) GetVulnerabilitiesOk() (*[]PackageVulnerability, bool)`

GetVulnerabilitiesOk returns a tuple with the Vulnerabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilities

`func (o *OnetimeEvaluationResponse) SetVulnerabilities(v []PackageVulnerability)`

SetVulnerabilities sets Vulnerabilities field to given value.

### HasVulnerabilities

`func (o *OnetimeEvaluationResponse) HasVulnerabilities() bool`

HasVulnerabilities returns a boolean if a field has been set.

### GetEvaluation

`func (o *OnetimeEvaluationResponse) GetEvaluation() PolicyEvaluationResult`

GetEvaluation returns the Evaluation field if non-nil, zero value otherwise.

### GetEvaluationOk

`func (o *OnetimeEvaluationResponse) GetEvaluationOk() (*PolicyEvaluationResult, bool)`

GetEvaluationOk returns a tuple with the Evaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluation

`func (o *OnetimeEvaluationResponse) SetEvaluation(v PolicyEvaluationResult)`

SetEvaluation sets Evaluation field to given value.

### HasEvaluation

`func (o *OnetimeEvaluationResponse) HasEvaluation() bool`

HasEvaluation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


