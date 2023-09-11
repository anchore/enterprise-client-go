# AnalysisArchiveSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalImageCount** | Pointer to **int32** | The number of unique images (digests) in the archive | [optional] 
**TotalTagCount** | Pointer to **int32** | The number of tag records (registry/repo:tag pull strings) in the archive. This may include repeated tags but will always have a unique tag-&gt;digest mapping per record. | [optional] 
**TotalDataBytes** | Pointer to **int32** | The total sum of all the bytes stored to the backing storage. Accounts for anchore-applied compression, but not compression by the underlying storage system. | [optional] 
**LastUpdated** | Pointer to **time.Time** | The timestamp of the most recent archived image | [optional] 

## Methods

### NewAnalysisArchiveSummary

`func NewAnalysisArchiveSummary() *AnalysisArchiveSummary`

NewAnalysisArchiveSummary instantiates a new AnalysisArchiveSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisArchiveSummaryWithDefaults

`func NewAnalysisArchiveSummaryWithDefaults() *AnalysisArchiveSummary`

NewAnalysisArchiveSummaryWithDefaults instantiates a new AnalysisArchiveSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalImageCount

`func (o *AnalysisArchiveSummary) GetTotalImageCount() int32`

GetTotalImageCount returns the TotalImageCount field if non-nil, zero value otherwise.

### GetTotalImageCountOk

`func (o *AnalysisArchiveSummary) GetTotalImageCountOk() (*int32, bool)`

GetTotalImageCountOk returns a tuple with the TotalImageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalImageCount

`func (o *AnalysisArchiveSummary) SetTotalImageCount(v int32)`

SetTotalImageCount sets TotalImageCount field to given value.

### HasTotalImageCount

`func (o *AnalysisArchiveSummary) HasTotalImageCount() bool`

HasTotalImageCount returns a boolean if a field has been set.

### GetTotalTagCount

`func (o *AnalysisArchiveSummary) GetTotalTagCount() int32`

GetTotalTagCount returns the TotalTagCount field if non-nil, zero value otherwise.

### GetTotalTagCountOk

`func (o *AnalysisArchiveSummary) GetTotalTagCountOk() (*int32, bool)`

GetTotalTagCountOk returns a tuple with the TotalTagCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTagCount

`func (o *AnalysisArchiveSummary) SetTotalTagCount(v int32)`

SetTotalTagCount sets TotalTagCount field to given value.

### HasTotalTagCount

`func (o *AnalysisArchiveSummary) HasTotalTagCount() bool`

HasTotalTagCount returns a boolean if a field has been set.

### GetTotalDataBytes

`func (o *AnalysisArchiveSummary) GetTotalDataBytes() int32`

GetTotalDataBytes returns the TotalDataBytes field if non-nil, zero value otherwise.

### GetTotalDataBytesOk

`func (o *AnalysisArchiveSummary) GetTotalDataBytesOk() (*int32, bool)`

GetTotalDataBytesOk returns a tuple with the TotalDataBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDataBytes

`func (o *AnalysisArchiveSummary) SetTotalDataBytes(v int32)`

SetTotalDataBytes sets TotalDataBytes field to given value.

### HasTotalDataBytes

`func (o *AnalysisArchiveSummary) HasTotalDataBytes() bool`

HasTotalDataBytes returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AnalysisArchiveSummary) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AnalysisArchiveSummary) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AnalysisArchiveSummary) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AnalysisArchiveSummary) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


