# AnchoreSummaryRuntimeImagesItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalysisStatus** | Pointer to **string** |  | [optional] 
**Compliant** | Pointer to **bool** |  | [optional] 
**TotalVulns** | Pointer to **int32** |  | [optional] 
**CriticalVulns** | Pointer to **int32** |  | [optional] 
**HighVulns** | Pointer to **int32** |  | [optional] 
**MediumVulns** | Pointer to **int32** |  | [optional] 
**LowVulns** | Pointer to **int32** |  | [optional] 
**NegligibleVulns** | Pointer to **int32** |  | [optional] 
**UnknownVulns** | Pointer to **int32** |  | [optional] 
**Context** | Pointer to **string** |  | [optional] 
**ImageTag** | Pointer to **string** |  | [optional] 
**ReportedDigest** | Pointer to **string** |  | [optional] 
**ImageDigest** | Pointer to **string** |  | [optional] 
**ParentDigest** | Pointer to **string** |  | [optional] 
**LastAnalyzed** | Pointer to **time.Time** |  | [optional] 
**LastSeen** | Pointer to **time.Time** |  | [optional] 
**PolicyId** | Pointer to **string** |  | [optional] 
**PolicyDigest** | Pointer to **string** |  | [optional] 
**PolicyReason** | Pointer to **string** |  | [optional] 
**EvaluatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAnchoreSummaryRuntimeImagesItem

`func NewAnchoreSummaryRuntimeImagesItem() *AnchoreSummaryRuntimeImagesItem`

NewAnchoreSummaryRuntimeImagesItem instantiates a new AnchoreSummaryRuntimeImagesItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnchoreSummaryRuntimeImagesItemWithDefaults

`func NewAnchoreSummaryRuntimeImagesItemWithDefaults() *AnchoreSummaryRuntimeImagesItem`

NewAnchoreSummaryRuntimeImagesItemWithDefaults instantiates a new AnchoreSummaryRuntimeImagesItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysisStatus

`func (o *AnchoreSummaryRuntimeImagesItem) GetAnalysisStatus() string`

GetAnalysisStatus returns the AnalysisStatus field if non-nil, zero value otherwise.

### GetAnalysisStatusOk

`func (o *AnchoreSummaryRuntimeImagesItem) GetAnalysisStatusOk() (*string, bool)`

GetAnalysisStatusOk returns a tuple with the AnalysisStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisStatus

`func (o *AnchoreSummaryRuntimeImagesItem) SetAnalysisStatus(v string)`

SetAnalysisStatus sets AnalysisStatus field to given value.

### HasAnalysisStatus

`func (o *AnchoreSummaryRuntimeImagesItem) HasAnalysisStatus() bool`

HasAnalysisStatus returns a boolean if a field has been set.

### GetCompliant

`func (o *AnchoreSummaryRuntimeImagesItem) GetCompliant() bool`

GetCompliant returns the Compliant field if non-nil, zero value otherwise.

### GetCompliantOk

`func (o *AnchoreSummaryRuntimeImagesItem) GetCompliantOk() (*bool, bool)`

GetCompliantOk returns a tuple with the Compliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliant

`func (o *AnchoreSummaryRuntimeImagesItem) SetCompliant(v bool)`

SetCompliant sets Compliant field to given value.

### HasCompliant

`func (o *AnchoreSummaryRuntimeImagesItem) HasCompliant() bool`

HasCompliant returns a boolean if a field has been set.

### GetTotalVulns

`func (o *AnchoreSummaryRuntimeImagesItem) GetTotalVulns() int32`

GetTotalVulns returns the TotalVulns field if non-nil, zero value otherwise.

### GetTotalVulnsOk

`func (o *AnchoreSummaryRuntimeImagesItem) GetTotalVulnsOk() (*int32, bool)`

GetTotalVulnsOk returns a tuple with the TotalVulns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVulns

`func (o *AnchoreSummaryRuntimeImagesItem) SetTotalVulns(v int32)`

SetTotalVulns sets TotalVulns field to given value.

### HasTotalVulns

`func (o *AnchoreSummaryRuntimeImagesItem) HasTotalVulns() bool`

HasTotalVulns returns a boolean if a field has been set.

### GetCriticalVulns

`func (o *AnchoreSummaryRuntimeImagesItem) GetCriticalVulns() int32`

GetCriticalVulns returns the CriticalVulns field if non-nil, zero value otherwise.

### GetCriticalVulnsOk

`func (o *AnchoreSummaryRuntimeImagesItem) GetCriticalVulnsOk() (*int32, bool)`

GetCriticalVulnsOk returns a tuple with the CriticalVulns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalVulns

`func (o *AnchoreSummaryRuntimeImagesItem) SetCriticalVulns(v int32)`

SetCriticalVulns sets CriticalVulns field to given value.

### HasCriticalVulns

`func (o *AnchoreSummaryRuntimeImagesItem) HasCriticalVulns() bool`

HasCriticalVulns returns a boolean if a field has been set.

### GetHighVulns

`func (o *AnchoreSummaryRuntimeImagesItem) GetHighVulns() int32`

GetHighVulns returns the HighVulns field if non-nil, zero value otherwise.

### GetHighVulnsOk

`func (o *AnchoreSummaryRuntimeImagesItem) GetHighVulnsOk() (*int32, bool)`

GetHighVulnsOk returns a tuple with the HighVulns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighVulns

`func (o *AnchoreSummaryRuntimeImagesItem) SetHighVulns(v int32)`

SetHighVulns sets HighVulns field to given value.

### HasHighVulns

`func (o *AnchoreSummaryRuntimeImagesItem) HasHighVulns() bool`

HasHighVulns returns a boolean if a field has been set.

### GetMediumVulns

`func (o *AnchoreSummaryRuntimeImagesItem) GetMediumVulns() int32`

GetMediumVulns returns the MediumVulns field if non-nil, zero value otherwise.

### GetMediumVulnsOk

`func (o *AnchoreSummaryRuntimeImagesItem) GetMediumVulnsOk() (*int32, bool)`

GetMediumVulnsOk returns a tuple with the MediumVulns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediumVulns

`func (o *AnchoreSummaryRuntimeImagesItem) SetMediumVulns(v int32)`

SetMediumVulns sets MediumVulns field to given value.

### HasMediumVulns

`func (o *AnchoreSummaryRuntimeImagesItem) HasMediumVulns() bool`

HasMediumVulns returns a boolean if a field has been set.

### GetLowVulns

`func (o *AnchoreSummaryRuntimeImagesItem) GetLowVulns() int32`

GetLowVulns returns the LowVulns field if non-nil, zero value otherwise.

### GetLowVulnsOk

`func (o *AnchoreSummaryRuntimeImagesItem) GetLowVulnsOk() (*int32, bool)`

GetLowVulnsOk returns a tuple with the LowVulns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowVulns

`func (o *AnchoreSummaryRuntimeImagesItem) SetLowVulns(v int32)`

SetLowVulns sets LowVulns field to given value.

### HasLowVulns

`func (o *AnchoreSummaryRuntimeImagesItem) HasLowVulns() bool`

HasLowVulns returns a boolean if a field has been set.

### GetNegligibleVulns

`func (o *AnchoreSummaryRuntimeImagesItem) GetNegligibleVulns() int32`

GetNegligibleVulns returns the NegligibleVulns field if non-nil, zero value otherwise.

### GetNegligibleVulnsOk

`func (o *AnchoreSummaryRuntimeImagesItem) GetNegligibleVulnsOk() (*int32, bool)`

GetNegligibleVulnsOk returns a tuple with the NegligibleVulns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegligibleVulns

`func (o *AnchoreSummaryRuntimeImagesItem) SetNegligibleVulns(v int32)`

SetNegligibleVulns sets NegligibleVulns field to given value.

### HasNegligibleVulns

`func (o *AnchoreSummaryRuntimeImagesItem) HasNegligibleVulns() bool`

HasNegligibleVulns returns a boolean if a field has been set.

### GetUnknownVulns

`func (o *AnchoreSummaryRuntimeImagesItem) GetUnknownVulns() int32`

GetUnknownVulns returns the UnknownVulns field if non-nil, zero value otherwise.

### GetUnknownVulnsOk

`func (o *AnchoreSummaryRuntimeImagesItem) GetUnknownVulnsOk() (*int32, bool)`

GetUnknownVulnsOk returns a tuple with the UnknownVulns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownVulns

`func (o *AnchoreSummaryRuntimeImagesItem) SetUnknownVulns(v int32)`

SetUnknownVulns sets UnknownVulns field to given value.

### HasUnknownVulns

`func (o *AnchoreSummaryRuntimeImagesItem) HasUnknownVulns() bool`

HasUnknownVulns returns a boolean if a field has been set.

### GetContext

`func (o *AnchoreSummaryRuntimeImagesItem) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *AnchoreSummaryRuntimeImagesItem) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *AnchoreSummaryRuntimeImagesItem) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *AnchoreSummaryRuntimeImagesItem) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetImageTag

`func (o *AnchoreSummaryRuntimeImagesItem) GetImageTag() string`

GetImageTag returns the ImageTag field if non-nil, zero value otherwise.

### GetImageTagOk

`func (o *AnchoreSummaryRuntimeImagesItem) GetImageTagOk() (*string, bool)`

GetImageTagOk returns a tuple with the ImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTag

`func (o *AnchoreSummaryRuntimeImagesItem) SetImageTag(v string)`

SetImageTag sets ImageTag field to given value.

### HasImageTag

`func (o *AnchoreSummaryRuntimeImagesItem) HasImageTag() bool`

HasImageTag returns a boolean if a field has been set.

### GetReportedDigest

`func (o *AnchoreSummaryRuntimeImagesItem) GetReportedDigest() string`

GetReportedDigest returns the ReportedDigest field if non-nil, zero value otherwise.

### GetReportedDigestOk

`func (o *AnchoreSummaryRuntimeImagesItem) GetReportedDigestOk() (*string, bool)`

GetReportedDigestOk returns a tuple with the ReportedDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedDigest

`func (o *AnchoreSummaryRuntimeImagesItem) SetReportedDigest(v string)`

SetReportedDigest sets ReportedDigest field to given value.

### HasReportedDigest

`func (o *AnchoreSummaryRuntimeImagesItem) HasReportedDigest() bool`

HasReportedDigest returns a boolean if a field has been set.

### GetImageDigest

`func (o *AnchoreSummaryRuntimeImagesItem) GetImageDigest() string`

GetImageDigest returns the ImageDigest field if non-nil, zero value otherwise.

### GetImageDigestOk

`func (o *AnchoreSummaryRuntimeImagesItem) GetImageDigestOk() (*string, bool)`

GetImageDigestOk returns a tuple with the ImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDigest

`func (o *AnchoreSummaryRuntimeImagesItem) SetImageDigest(v string)`

SetImageDigest sets ImageDigest field to given value.

### HasImageDigest

`func (o *AnchoreSummaryRuntimeImagesItem) HasImageDigest() bool`

HasImageDigest returns a boolean if a field has been set.

### GetParentDigest

`func (o *AnchoreSummaryRuntimeImagesItem) GetParentDigest() string`

GetParentDigest returns the ParentDigest field if non-nil, zero value otherwise.

### GetParentDigestOk

`func (o *AnchoreSummaryRuntimeImagesItem) GetParentDigestOk() (*string, bool)`

GetParentDigestOk returns a tuple with the ParentDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDigest

`func (o *AnchoreSummaryRuntimeImagesItem) SetParentDigest(v string)`

SetParentDigest sets ParentDigest field to given value.

### HasParentDigest

`func (o *AnchoreSummaryRuntimeImagesItem) HasParentDigest() bool`

HasParentDigest returns a boolean if a field has been set.

### GetLastAnalyzed

`func (o *AnchoreSummaryRuntimeImagesItem) GetLastAnalyzed() time.Time`

GetLastAnalyzed returns the LastAnalyzed field if non-nil, zero value otherwise.

### GetLastAnalyzedOk

`func (o *AnchoreSummaryRuntimeImagesItem) GetLastAnalyzedOk() (*time.Time, bool)`

GetLastAnalyzedOk returns a tuple with the LastAnalyzed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAnalyzed

`func (o *AnchoreSummaryRuntimeImagesItem) SetLastAnalyzed(v time.Time)`

SetLastAnalyzed sets LastAnalyzed field to given value.

### HasLastAnalyzed

`func (o *AnchoreSummaryRuntimeImagesItem) HasLastAnalyzed() bool`

HasLastAnalyzed returns a boolean if a field has been set.

### GetLastSeen

`func (o *AnchoreSummaryRuntimeImagesItem) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *AnchoreSummaryRuntimeImagesItem) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *AnchoreSummaryRuntimeImagesItem) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *AnchoreSummaryRuntimeImagesItem) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetPolicyId

`func (o *AnchoreSummaryRuntimeImagesItem) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *AnchoreSummaryRuntimeImagesItem) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *AnchoreSummaryRuntimeImagesItem) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *AnchoreSummaryRuntimeImagesItem) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPolicyDigest

`func (o *AnchoreSummaryRuntimeImagesItem) GetPolicyDigest() string`

GetPolicyDigest returns the PolicyDigest field if non-nil, zero value otherwise.

### GetPolicyDigestOk

`func (o *AnchoreSummaryRuntimeImagesItem) GetPolicyDigestOk() (*string, bool)`

GetPolicyDigestOk returns a tuple with the PolicyDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyDigest

`func (o *AnchoreSummaryRuntimeImagesItem) SetPolicyDigest(v string)`

SetPolicyDigest sets PolicyDigest field to given value.

### HasPolicyDigest

`func (o *AnchoreSummaryRuntimeImagesItem) HasPolicyDigest() bool`

HasPolicyDigest returns a boolean if a field has been set.

### GetPolicyReason

`func (o *AnchoreSummaryRuntimeImagesItem) GetPolicyReason() string`

GetPolicyReason returns the PolicyReason field if non-nil, zero value otherwise.

### GetPolicyReasonOk

`func (o *AnchoreSummaryRuntimeImagesItem) GetPolicyReasonOk() (*string, bool)`

GetPolicyReasonOk returns a tuple with the PolicyReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyReason

`func (o *AnchoreSummaryRuntimeImagesItem) SetPolicyReason(v string)`

SetPolicyReason sets PolicyReason field to given value.

### HasPolicyReason

`func (o *AnchoreSummaryRuntimeImagesItem) HasPolicyReason() bool`

HasPolicyReason returns a boolean if a field has been set.

### GetEvaluatedAt

`func (o *AnchoreSummaryRuntimeImagesItem) GetEvaluatedAt() time.Time`

GetEvaluatedAt returns the EvaluatedAt field if non-nil, zero value otherwise.

### GetEvaluatedAtOk

`func (o *AnchoreSummaryRuntimeImagesItem) GetEvaluatedAtOk() (*time.Time, bool)`

GetEvaluatedAtOk returns a tuple with the EvaluatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatedAt

`func (o *AnchoreSummaryRuntimeImagesItem) SetEvaluatedAt(v time.Time)`

SetEvaluatedAt sets EvaluatedAt field to given value.

### HasEvaluatedAt

`func (o *AnchoreSummaryRuntimeImagesItem) HasEvaluatedAt() bool`

HasEvaluatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


