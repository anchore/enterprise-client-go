# Integration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | Globally unique identifier for the integration instance | 
**Type** | [**IntegrationType**](IntegrationType.md) |  | 
**Name** | Pointer to **string** | Name of the integration instance | [optional] 
**Version** | Pointer to **string** | Version of the integration instance | [optional] 
**ReportedStatus** | Pointer to [**IntegrationReportedStatus**](IntegrationReportedStatus.md) |  | [optional] 
**IntegrationStatus** | Pointer to [**IntegrationStatus**](IntegrationStatus.md) |  | [optional] 
**LastSeen** | Pointer to **time.Time** | Timestamp of last received health report from integration instance | [optional] [readonly] 
**Uptime** | Pointer to **float32** | Running time of integration instance in seconds | [optional] 
**AccountName** | **string** | Account that the integration instance reports to by default | 
**Username** | **string** | User that the integration instance authenticates as by default | 
**ExplicitlyAccountBound** | Pointer to **[]string** | List of accounts that the integration instance is statically configured to handle | [optional] 
**Accounts** | Pointer to **[]string** | List of accounts that the integration instance handles | [optional] 
**ClusterName** | Pointer to **string** | Name of cluster where the integration instance runs | [optional] 
**Namespace** | Pointer to **string** | Namespace in which the integration instance runs | [optional] 
**HealthReportInterval** | Pointer to **int32** | Interval (in seconds) between health reports | [optional] 
**RegistrationId** | Pointer to **string** | identifier for the integration during registration until it has received its integration_id | [optional] 
**RegistrationInstanceId** | Pointer to **string** | Unique identifier for the integration instance (among its replicas) during registration | [optional] 
**Description** | Pointer to **string** | Short description of the integration instance | [optional] 
**StartedAt** | Pointer to **time.Time** | Timestamp when integration instance was started | [optional] 
**Namespaces** | Pointer to **[]string** | List of namespaces handled by the integration instance | [optional] 
**Configuration** | Pointer to **interface{}** | Configuration of the integration instance | [optional] 

## Methods

### NewIntegration

`func NewIntegration(uuid string, type_ IntegrationType, accountName string, username string, ) *Integration`

NewIntegration instantiates a new Integration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationWithDefaults

`func NewIntegrationWithDefaults() *Integration`

NewIntegrationWithDefaults instantiates a new Integration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *Integration) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Integration) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Integration) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetType

`func (o *Integration) GetType() IntegrationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Integration) GetTypeOk() (*IntegrationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Integration) SetType(v IntegrationType)`

SetType sets Type field to given value.


### GetName

`func (o *Integration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Integration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Integration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Integration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *Integration) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Integration) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Integration) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Integration) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetReportedStatus

`func (o *Integration) GetReportedStatus() IntegrationReportedStatus`

GetReportedStatus returns the ReportedStatus field if non-nil, zero value otherwise.

### GetReportedStatusOk

`func (o *Integration) GetReportedStatusOk() (*IntegrationReportedStatus, bool)`

GetReportedStatusOk returns a tuple with the ReportedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedStatus

`func (o *Integration) SetReportedStatus(v IntegrationReportedStatus)`

SetReportedStatus sets ReportedStatus field to given value.

### HasReportedStatus

`func (o *Integration) HasReportedStatus() bool`

HasReportedStatus returns a boolean if a field has been set.

### GetIntegrationStatus

`func (o *Integration) GetIntegrationStatus() IntegrationStatus`

GetIntegrationStatus returns the IntegrationStatus field if non-nil, zero value otherwise.

### GetIntegrationStatusOk

`func (o *Integration) GetIntegrationStatusOk() (*IntegrationStatus, bool)`

GetIntegrationStatusOk returns a tuple with the IntegrationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationStatus

`func (o *Integration) SetIntegrationStatus(v IntegrationStatus)`

SetIntegrationStatus sets IntegrationStatus field to given value.

### HasIntegrationStatus

`func (o *Integration) HasIntegrationStatus() bool`

HasIntegrationStatus returns a boolean if a field has been set.

### GetLastSeen

`func (o *Integration) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *Integration) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *Integration) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *Integration) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetUptime

`func (o *Integration) GetUptime() float32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *Integration) GetUptimeOk() (*float32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *Integration) SetUptime(v float32)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *Integration) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetAccountName

`func (o *Integration) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *Integration) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *Integration) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.


### GetUsername

`func (o *Integration) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Integration) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Integration) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetExplicitlyAccountBound

`func (o *Integration) GetExplicitlyAccountBound() []string`

GetExplicitlyAccountBound returns the ExplicitlyAccountBound field if non-nil, zero value otherwise.

### GetExplicitlyAccountBoundOk

`func (o *Integration) GetExplicitlyAccountBoundOk() (*[]string, bool)`

GetExplicitlyAccountBoundOk returns a tuple with the ExplicitlyAccountBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitlyAccountBound

`func (o *Integration) SetExplicitlyAccountBound(v []string)`

SetExplicitlyAccountBound sets ExplicitlyAccountBound field to given value.

### HasExplicitlyAccountBound

`func (o *Integration) HasExplicitlyAccountBound() bool`

HasExplicitlyAccountBound returns a boolean if a field has been set.

### GetAccounts

`func (o *Integration) GetAccounts() []string`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *Integration) GetAccountsOk() (*[]string, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *Integration) SetAccounts(v []string)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *Integration) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetClusterName

`func (o *Integration) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *Integration) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *Integration) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *Integration) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetNamespace

`func (o *Integration) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *Integration) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *Integration) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *Integration) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetHealthReportInterval

`func (o *Integration) GetHealthReportInterval() int32`

GetHealthReportInterval returns the HealthReportInterval field if non-nil, zero value otherwise.

### GetHealthReportIntervalOk

`func (o *Integration) GetHealthReportIntervalOk() (*int32, bool)`

GetHealthReportIntervalOk returns a tuple with the HealthReportInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthReportInterval

`func (o *Integration) SetHealthReportInterval(v int32)`

SetHealthReportInterval sets HealthReportInterval field to given value.

### HasHealthReportInterval

`func (o *Integration) HasHealthReportInterval() bool`

HasHealthReportInterval returns a boolean if a field has been set.

### GetRegistrationId

`func (o *Integration) GetRegistrationId() string`

GetRegistrationId returns the RegistrationId field if non-nil, zero value otherwise.

### GetRegistrationIdOk

`func (o *Integration) GetRegistrationIdOk() (*string, bool)`

GetRegistrationIdOk returns a tuple with the RegistrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationId

`func (o *Integration) SetRegistrationId(v string)`

SetRegistrationId sets RegistrationId field to given value.

### HasRegistrationId

`func (o *Integration) HasRegistrationId() bool`

HasRegistrationId returns a boolean if a field has been set.

### GetRegistrationInstanceId

`func (o *Integration) GetRegistrationInstanceId() string`

GetRegistrationInstanceId returns the RegistrationInstanceId field if non-nil, zero value otherwise.

### GetRegistrationInstanceIdOk

`func (o *Integration) GetRegistrationInstanceIdOk() (*string, bool)`

GetRegistrationInstanceIdOk returns a tuple with the RegistrationInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationInstanceId

`func (o *Integration) SetRegistrationInstanceId(v string)`

SetRegistrationInstanceId sets RegistrationInstanceId field to given value.

### HasRegistrationInstanceId

`func (o *Integration) HasRegistrationInstanceId() bool`

HasRegistrationInstanceId returns a boolean if a field has been set.

### GetDescription

`func (o *Integration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Integration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Integration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Integration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStartedAt

`func (o *Integration) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Integration) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Integration) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *Integration) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetNamespaces

`func (o *Integration) GetNamespaces() []string`

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

`func (o *Integration) GetNamespacesOk() (*[]string, bool)`

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

`func (o *Integration) SetNamespaces(v []string)`

SetNamespaces sets Namespaces field to given value.

### HasNamespaces

`func (o *Integration) HasNamespaces() bool`

HasNamespaces returns a boolean if a field has been set.

### GetConfiguration

`func (o *Integration) GetConfiguration() interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *Integration) GetConfigurationOk() (*interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *Integration) SetConfiguration(v interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *Integration) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


