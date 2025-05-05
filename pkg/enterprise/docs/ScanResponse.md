# ScanResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VulnScan** | Pointer to [**SBOMVulnerabilitiesResponse**](SBOMVulnerabilitiesResponse.md) |  | [optional] 
**PolicyScan** | Pointer to [**PolicyEvaluation**](PolicyEvaluation.md) |  | [optional] 

## Methods

### NewScanResponse

`func NewScanResponse() *ScanResponse`

NewScanResponse instantiates a new ScanResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanResponseWithDefaults

`func NewScanResponseWithDefaults() *ScanResponse`

NewScanResponseWithDefaults instantiates a new ScanResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVulnScan

`func (o *ScanResponse) GetVulnScan() SBOMVulnerabilitiesResponse`

GetVulnScan returns the VulnScan field if non-nil, zero value otherwise.

### GetVulnScanOk

`func (o *ScanResponse) GetVulnScanOk() (*SBOMVulnerabilitiesResponse, bool)`

GetVulnScanOk returns a tuple with the VulnScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnScan

`func (o *ScanResponse) SetVulnScan(v SBOMVulnerabilitiesResponse)`

SetVulnScan sets VulnScan field to given value.

### HasVulnScan

`func (o *ScanResponse) HasVulnScan() bool`

HasVulnScan returns a boolean if a field has been set.

### GetPolicyScan

`func (o *ScanResponse) GetPolicyScan() PolicyEvaluation`

GetPolicyScan returns the PolicyScan field if non-nil, zero value otherwise.

### GetPolicyScanOk

`func (o *ScanResponse) GetPolicyScanOk() (*PolicyEvaluation, bool)`

GetPolicyScanOk returns a tuple with the PolicyScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyScan

`func (o *ScanResponse) SetPolicyScan(v PolicyEvaluation)`

SetPolicyScan sets PolicyScan field to given value.

### HasPolicyScan

`func (o *ScanResponse) HasPolicyScan() bool`

HasPolicyScan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


