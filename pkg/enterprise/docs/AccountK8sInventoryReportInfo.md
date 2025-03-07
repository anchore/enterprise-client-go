# AccountK8sInventoryReportInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportTimestamp** | **string** | Timestamp for the inventory report | 
**AccountName** | **string** | Account to which the inventory reports belong | 
**SentAsUser** | **string** | User that the integration instance sent the inventory report as | 
**BatchSize** | **int64** | Number of batches that the inventory report was sent in | 
**LastSuccessfulIndex** | **int64** | Index of last successfully sent batch, -1 if none were successful | 
**HasErrors** | **bool** | true if one or more of the batches coult not be sent, false otherwise | 
**Batches** | [**[]AccountK8sInventoryReportInfoBatchesInner**](AccountK8sInventoryReportInfoBatchesInner.md) | List of information about inventory report batches | 

## Methods

### NewAccountK8sInventoryReportInfo

`func NewAccountK8sInventoryReportInfo(reportTimestamp string, accountName string, sentAsUser string, batchSize int64, lastSuccessfulIndex int64, hasErrors bool, batches []AccountK8sInventoryReportInfoBatchesInner, ) *AccountK8sInventoryReportInfo`

NewAccountK8sInventoryReportInfo instantiates a new AccountK8sInventoryReportInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountK8sInventoryReportInfoWithDefaults

`func NewAccountK8sInventoryReportInfoWithDefaults() *AccountK8sInventoryReportInfo`

NewAccountK8sInventoryReportInfoWithDefaults instantiates a new AccountK8sInventoryReportInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportTimestamp

`func (o *AccountK8sInventoryReportInfo) GetReportTimestamp() string`

GetReportTimestamp returns the ReportTimestamp field if non-nil, zero value otherwise.

### GetReportTimestampOk

`func (o *AccountK8sInventoryReportInfo) GetReportTimestampOk() (*string, bool)`

GetReportTimestampOk returns a tuple with the ReportTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTimestamp

`func (o *AccountK8sInventoryReportInfo) SetReportTimestamp(v string)`

SetReportTimestamp sets ReportTimestamp field to given value.


### GetAccountName

`func (o *AccountK8sInventoryReportInfo) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *AccountK8sInventoryReportInfo) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *AccountK8sInventoryReportInfo) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.


### GetSentAsUser

`func (o *AccountK8sInventoryReportInfo) GetSentAsUser() string`

GetSentAsUser returns the SentAsUser field if non-nil, zero value otherwise.

### GetSentAsUserOk

`func (o *AccountK8sInventoryReportInfo) GetSentAsUserOk() (*string, bool)`

GetSentAsUserOk returns a tuple with the SentAsUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAsUser

`func (o *AccountK8sInventoryReportInfo) SetSentAsUser(v string)`

SetSentAsUser sets SentAsUser field to given value.


### GetBatchSize

`func (o *AccountK8sInventoryReportInfo) GetBatchSize() int64`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *AccountK8sInventoryReportInfo) GetBatchSizeOk() (*int64, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *AccountK8sInventoryReportInfo) SetBatchSize(v int64)`

SetBatchSize sets BatchSize field to given value.


### GetLastSuccessfulIndex

`func (o *AccountK8sInventoryReportInfo) GetLastSuccessfulIndex() int64`

GetLastSuccessfulIndex returns the LastSuccessfulIndex field if non-nil, zero value otherwise.

### GetLastSuccessfulIndexOk

`func (o *AccountK8sInventoryReportInfo) GetLastSuccessfulIndexOk() (*int64, bool)`

GetLastSuccessfulIndexOk returns a tuple with the LastSuccessfulIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessfulIndex

`func (o *AccountK8sInventoryReportInfo) SetLastSuccessfulIndex(v int64)`

SetLastSuccessfulIndex sets LastSuccessfulIndex field to given value.


### GetHasErrors

`func (o *AccountK8sInventoryReportInfo) GetHasErrors() bool`

GetHasErrors returns the HasErrors field if non-nil, zero value otherwise.

### GetHasErrorsOk

`func (o *AccountK8sInventoryReportInfo) GetHasErrorsOk() (*bool, bool)`

GetHasErrorsOk returns a tuple with the HasErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasErrors

`func (o *AccountK8sInventoryReportInfo) SetHasErrors(v bool)`

SetHasErrors sets HasErrors field to given value.


### GetBatches

`func (o *AccountK8sInventoryReportInfo) GetBatches() []AccountK8sInventoryReportInfoBatchesInner`

GetBatches returns the Batches field if non-nil, zero value otherwise.

### GetBatchesOk

`func (o *AccountK8sInventoryReportInfo) GetBatchesOk() (*[]AccountK8sInventoryReportInfoBatchesInner, bool)`

GetBatchesOk returns a tuple with the Batches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatches

`func (o *AccountK8sInventoryReportInfo) SetBatches(v []AccountK8sInventoryReportInfoBatchesInner)`

SetBatches sets Batches field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


