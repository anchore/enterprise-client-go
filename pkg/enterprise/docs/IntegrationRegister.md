# IntegrationRegister

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegistrationId** | **string** | identifier for the integration during registration until it has received its integration_id | 
**RegistrationInstanceId** | **string** | Unique identifier for the integration instance (among its replicas) during registration | 
**Type** | [**IntegrationType**](IntegrationType.md) |  | 
**Name** | Pointer to **string** | Name of the integration instance | [optional] 
**Description** | Pointer to **string** | Short description of the integration instance | [optional] 
**Version** | Pointer to **string** | Version of the integration instance | [optional] 
**StartedAt** | Pointer to **time.Time** | Timestamp when integration instance was started | [optional] 
**Uptime** | Pointer to **float32** | Running time of integration instance in seconds | [optional] 
**Username** | **string** | User that the integration instance authenticates as by default | 
**ExplicitlyAccountBound** | Pointer to **[]string** | List of accounts that the integration instance is statically configured to handle | [optional] 
**Namespaces** | Pointer to **[]string** | List of namespaces that the integration instance is statically configured to handle | [optional] 
**Configuration** | Pointer to **interface{}** | Configuration of the integration instance | [optional] 
**ClusterName** | Pointer to **string** | Name of cluster where the integration instance runs | [optional] 
**Namespace** | Pointer to **string** | Namespace in which the integration instance runs | [optional] 
**HealthReportInterval** | **int64** | Interval (in seconds) between health reports | 

## Methods

### NewIntegrationRegister

`func NewIntegrationRegister(registrationId string, registrationInstanceId string, type_ IntegrationType, username string, healthReportInterval int64, ) *IntegrationRegister`

NewIntegrationRegister instantiates a new IntegrationRegister object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationRegisterWithDefaults

`func NewIntegrationRegisterWithDefaults() *IntegrationRegister`

NewIntegrationRegisterWithDefaults instantiates a new IntegrationRegister object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistrationId

`func (o *IntegrationRegister) GetRegistrationId() string`

GetRegistrationId returns the RegistrationId field if non-nil, zero value otherwise.

### GetRegistrationIdOk

`func (o *IntegrationRegister) GetRegistrationIdOk() (*string, bool)`

GetRegistrationIdOk returns a tuple with the RegistrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationId

`func (o *IntegrationRegister) SetRegistrationId(v string)`

SetRegistrationId sets RegistrationId field to given value.


### GetRegistrationInstanceId

`func (o *IntegrationRegister) GetRegistrationInstanceId() string`

GetRegistrationInstanceId returns the RegistrationInstanceId field if non-nil, zero value otherwise.

### GetRegistrationInstanceIdOk

`func (o *IntegrationRegister) GetRegistrationInstanceIdOk() (*string, bool)`

GetRegistrationInstanceIdOk returns a tuple with the RegistrationInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationInstanceId

`func (o *IntegrationRegister) SetRegistrationInstanceId(v string)`

SetRegistrationInstanceId sets RegistrationInstanceId field to given value.


### GetType

`func (o *IntegrationRegister) GetType() IntegrationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntegrationRegister) GetTypeOk() (*IntegrationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntegrationRegister) SetType(v IntegrationType)`

SetType sets Type field to given value.


### GetName

`func (o *IntegrationRegister) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationRegister) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationRegister) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IntegrationRegister) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *IntegrationRegister) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IntegrationRegister) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IntegrationRegister) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IntegrationRegister) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersion

`func (o *IntegrationRegister) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IntegrationRegister) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IntegrationRegister) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *IntegrationRegister) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetStartedAt

`func (o *IntegrationRegister) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *IntegrationRegister) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *IntegrationRegister) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *IntegrationRegister) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetUptime

`func (o *IntegrationRegister) GetUptime() float32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *IntegrationRegister) GetUptimeOk() (*float32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *IntegrationRegister) SetUptime(v float32)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *IntegrationRegister) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetUsername

`func (o *IntegrationRegister) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *IntegrationRegister) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *IntegrationRegister) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetExplicitlyAccountBound

`func (o *IntegrationRegister) GetExplicitlyAccountBound() []string`

GetExplicitlyAccountBound returns the ExplicitlyAccountBound field if non-nil, zero value otherwise.

### GetExplicitlyAccountBoundOk

`func (o *IntegrationRegister) GetExplicitlyAccountBoundOk() (*[]string, bool)`

GetExplicitlyAccountBoundOk returns a tuple with the ExplicitlyAccountBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitlyAccountBound

`func (o *IntegrationRegister) SetExplicitlyAccountBound(v []string)`

SetExplicitlyAccountBound sets ExplicitlyAccountBound field to given value.

### HasExplicitlyAccountBound

`func (o *IntegrationRegister) HasExplicitlyAccountBound() bool`

HasExplicitlyAccountBound returns a boolean if a field has been set.

### GetNamespaces

`func (o *IntegrationRegister) GetNamespaces() []string`

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

`func (o *IntegrationRegister) GetNamespacesOk() (*[]string, bool)`

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

`func (o *IntegrationRegister) SetNamespaces(v []string)`

SetNamespaces sets Namespaces field to given value.

### HasNamespaces

`func (o *IntegrationRegister) HasNamespaces() bool`

HasNamespaces returns a boolean if a field has been set.

### GetConfiguration

`func (o *IntegrationRegister) GetConfiguration() interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *IntegrationRegister) GetConfigurationOk() (*interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *IntegrationRegister) SetConfiguration(v interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *IntegrationRegister) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetClusterName

`func (o *IntegrationRegister) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *IntegrationRegister) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *IntegrationRegister) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *IntegrationRegister) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetNamespace

`func (o *IntegrationRegister) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *IntegrationRegister) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *IntegrationRegister) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *IntegrationRegister) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetHealthReportInterval

`func (o *IntegrationRegister) GetHealthReportInterval() int64`

GetHealthReportInterval returns the HealthReportInterval field if non-nil, zero value otherwise.

### GetHealthReportIntervalOk

`func (o *IntegrationRegister) GetHealthReportIntervalOk() (*int64, bool)`

GetHealthReportIntervalOk returns a tuple with the HealthReportInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthReportInterval

`func (o *IntegrationRegister) SetHealthReportInterval(v int64)`

SetHealthReportInterval sets HealthReportInterval field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


