# HealthData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**IntegrationType**](IntegrationType.md) |  | 
**Version** | **int64** | Version of the health data schema | 
**Errors** | Pointer to **[]string** | List of errors | [optional] 
**AccountK8sInventoryReports** | Pointer to [**map[string]AccountK8sInventoryReportInfo**](AccountK8sInventoryReportInfo.md) | Information about sent inventory report for accounts | [optional] 

## Methods

### NewHealthData

`func NewHealthData(type_ IntegrationType, version int64, ) *HealthData`

NewHealthData instantiates a new HealthData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthDataWithDefaults

`func NewHealthDataWithDefaults() *HealthData`

NewHealthDataWithDefaults instantiates a new HealthData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *HealthData) GetType() IntegrationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HealthData) GetTypeOk() (*IntegrationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HealthData) SetType(v IntegrationType)`

SetType sets Type field to given value.


### GetVersion

`func (o *HealthData) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HealthData) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HealthData) SetVersion(v int64)`

SetVersion sets Version field to given value.


### GetErrors

`func (o *HealthData) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *HealthData) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *HealthData) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *HealthData) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetAccountK8sInventoryReports

`func (o *HealthData) GetAccountK8sInventoryReports() map[string]AccountK8sInventoryReportInfo`

GetAccountK8sInventoryReports returns the AccountK8sInventoryReports field if non-nil, zero value otherwise.

### GetAccountK8sInventoryReportsOk

`func (o *HealthData) GetAccountK8sInventoryReportsOk() (*map[string]AccountK8sInventoryReportInfo, bool)`

GetAccountK8sInventoryReportsOk returns a tuple with the AccountK8sInventoryReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountK8sInventoryReports

`func (o *HealthData) SetAccountK8sInventoryReports(v map[string]AccountK8sInventoryReportInfo)`

SetAccountK8sInventoryReports sets AccountK8sInventoryReports field to given value.

### HasAccountK8sInventoryReports

`func (o *HealthData) HasAccountK8sInventoryReports() bool`

HasAccountK8sInventoryReports returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


