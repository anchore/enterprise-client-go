# InventoryReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | Pointer to **string** |  | [optional] 
**InventoryType** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Results** | Pointer to [**[]InventoryReportItem**](InventoryReportItem.md) |  | [optional] 

## Methods

### NewInventoryReport

`func NewInventoryReport() *InventoryReport`

NewInventoryReport instantiates a new InventoryReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryReportWithDefaults

`func NewInventoryReportWithDefaults() *InventoryReport`

NewInventoryReportWithDefaults instantiates a new InventoryReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *InventoryReport) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *InventoryReport) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *InventoryReport) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *InventoryReport) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetInventoryType

`func (o *InventoryReport) GetInventoryType() string`

GetInventoryType returns the InventoryType field if non-nil, zero value otherwise.

### GetInventoryTypeOk

`func (o *InventoryReport) GetInventoryTypeOk() (*string, bool)`

GetInventoryTypeOk returns a tuple with the InventoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryType

`func (o *InventoryReport) SetInventoryType(v string)`

SetInventoryType sets InventoryType field to given value.

### HasInventoryType

`func (o *InventoryReport) HasInventoryType() bool`

HasInventoryType returns a boolean if a field has been set.

### GetTimestamp

`func (o *InventoryReport) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InventoryReport) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InventoryReport) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *InventoryReport) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetResults

`func (o *InventoryReport) GetResults() []InventoryReportItem`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *InventoryReport) GetResultsOk() (*[]InventoryReportItem, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *InventoryReport) SetResults(v []InventoryReportItem)`

SetResults sets Results field to given value.

### HasResults

`func (o *InventoryReport) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


