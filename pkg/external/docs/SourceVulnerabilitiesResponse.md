# SourceVulnerabilitiesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceId** | Pointer to **string** |  | [optional] 
**VulnerabilityType** | Pointer to **string** |  | [optional] 
**Vulnerabilities** | Pointer to [**[]Vulnerability**](Vulnerability.md) | List of Vulnerability objects | [optional] 

## Methods

### NewSourceVulnerabilitiesResponse

`func NewSourceVulnerabilitiesResponse() *SourceVulnerabilitiesResponse`

NewSourceVulnerabilitiesResponse instantiates a new SourceVulnerabilitiesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceVulnerabilitiesResponseWithDefaults

`func NewSourceVulnerabilitiesResponseWithDefaults() *SourceVulnerabilitiesResponse`

NewSourceVulnerabilitiesResponseWithDefaults instantiates a new SourceVulnerabilitiesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceId

`func (o *SourceVulnerabilitiesResponse) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *SourceVulnerabilitiesResponse) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *SourceVulnerabilitiesResponse) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *SourceVulnerabilitiesResponse) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetVulnerabilityType

`func (o *SourceVulnerabilitiesResponse) GetVulnerabilityType() string`

GetVulnerabilityType returns the VulnerabilityType field if non-nil, zero value otherwise.

### GetVulnerabilityTypeOk

`func (o *SourceVulnerabilitiesResponse) GetVulnerabilityTypeOk() (*string, bool)`

GetVulnerabilityTypeOk returns a tuple with the VulnerabilityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityType

`func (o *SourceVulnerabilitiesResponse) SetVulnerabilityType(v string)`

SetVulnerabilityType sets VulnerabilityType field to given value.

### HasVulnerabilityType

`func (o *SourceVulnerabilitiesResponse) HasVulnerabilityType() bool`

HasVulnerabilityType returns a boolean if a field has been set.

### GetVulnerabilities

`func (o *SourceVulnerabilitiesResponse) GetVulnerabilities() []Vulnerability`

GetVulnerabilities returns the Vulnerabilities field if non-nil, zero value otherwise.

### GetVulnerabilitiesOk

`func (o *SourceVulnerabilitiesResponse) GetVulnerabilitiesOk() (*[]Vulnerability, bool)`

GetVulnerabilitiesOk returns a tuple with the Vulnerabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilities

`func (o *SourceVulnerabilitiesResponse) SetVulnerabilities(v []Vulnerability)`

SetVulnerabilities sets Vulnerabilities field to given value.

### HasVulnerabilities

`func (o *SourceVulnerabilitiesResponse) HasVulnerabilities() bool`

HasVulnerabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


