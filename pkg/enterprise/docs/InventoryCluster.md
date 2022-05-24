# InventoryCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | Pointer to **string** |  | [optional] 
**InventoryType** | Pointer to **string** |  | [optional] 
**ClusterConfig** | Pointer to [**InventoryClusterConfig**](InventoryClusterConfig.md) |  | [optional] 

## Methods

### NewInventoryCluster

`func NewInventoryCluster() *InventoryCluster`

NewInventoryCluster instantiates a new InventoryCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryClusterWithDefaults

`func NewInventoryClusterWithDefaults() *InventoryCluster`

NewInventoryClusterWithDefaults instantiates a new InventoryCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *InventoryCluster) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *InventoryCluster) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *InventoryCluster) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *InventoryCluster) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetInventoryType

`func (o *InventoryCluster) GetInventoryType() string`

GetInventoryType returns the InventoryType field if non-nil, zero value otherwise.

### GetInventoryTypeOk

`func (o *InventoryCluster) GetInventoryTypeOk() (*string, bool)`

GetInventoryTypeOk returns a tuple with the InventoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryType

`func (o *InventoryCluster) SetInventoryType(v string)`

SetInventoryType sets InventoryType field to given value.

### HasInventoryType

`func (o *InventoryCluster) HasInventoryType() bool`

HasInventoryType returns a boolean if a field has been set.

### GetClusterConfig

`func (o *InventoryCluster) GetClusterConfig() InventoryClusterConfig`

GetClusterConfig returns the ClusterConfig field if non-nil, zero value otherwise.

### GetClusterConfigOk

`func (o *InventoryCluster) GetClusterConfigOk() (*InventoryClusterConfig, bool)`

GetClusterConfigOk returns a tuple with the ClusterConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterConfig

`func (o *InventoryCluster) SetClusterConfig(v InventoryClusterConfig)`

SetClusterConfig sets ClusterConfig field to given value.

### HasClusterConfig

`func (o *InventoryCluster) HasClusterConfig() bool`

HasClusterConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


