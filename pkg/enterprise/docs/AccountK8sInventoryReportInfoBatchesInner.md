# AccountK8sInventoryReportInfoBatchesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchIndex** | **int64** | Index of the batch | 
**SendTimestamp** | **time.Time** | Timestamp when the batch was sent | 
**Error** | Pointer to **string** | Error message if the sending was unsuccessful | [optional] 

## Methods

### NewAccountK8sInventoryReportInfoBatchesInner

`func NewAccountK8sInventoryReportInfoBatchesInner(batchIndex int64, sendTimestamp time.Time, ) *AccountK8sInventoryReportInfoBatchesInner`

NewAccountK8sInventoryReportInfoBatchesInner instantiates a new AccountK8sInventoryReportInfoBatchesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountK8sInventoryReportInfoBatchesInnerWithDefaults

`func NewAccountK8sInventoryReportInfoBatchesInnerWithDefaults() *AccountK8sInventoryReportInfoBatchesInner`

NewAccountK8sInventoryReportInfoBatchesInnerWithDefaults instantiates a new AccountK8sInventoryReportInfoBatchesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchIndex

`func (o *AccountK8sInventoryReportInfoBatchesInner) GetBatchIndex() int64`

GetBatchIndex returns the BatchIndex field if non-nil, zero value otherwise.

### GetBatchIndexOk

`func (o *AccountK8sInventoryReportInfoBatchesInner) GetBatchIndexOk() (*int64, bool)`

GetBatchIndexOk returns a tuple with the BatchIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchIndex

`func (o *AccountK8sInventoryReportInfoBatchesInner) SetBatchIndex(v int64)`

SetBatchIndex sets BatchIndex field to given value.


### GetSendTimestamp

`func (o *AccountK8sInventoryReportInfoBatchesInner) GetSendTimestamp() time.Time`

GetSendTimestamp returns the SendTimestamp field if non-nil, zero value otherwise.

### GetSendTimestampOk

`func (o *AccountK8sInventoryReportInfoBatchesInner) GetSendTimestampOk() (*time.Time, bool)`

GetSendTimestampOk returns a tuple with the SendTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendTimestamp

`func (o *AccountK8sInventoryReportInfoBatchesInner) SetSendTimestamp(v time.Time)`

SetSendTimestamp sets SendTimestamp field to given value.


### GetError

`func (o *AccountK8sInventoryReportInfoBatchesInner) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AccountK8sInventoryReportInfoBatchesInner) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AccountK8sInventoryReportInfoBatchesInner) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AccountK8sInventoryReportInfoBatchesInner) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


