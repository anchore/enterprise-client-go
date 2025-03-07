# IntegrationSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Globally unique identifier for the integration instance | [optional] 
**Type** | Pointer to [**IntegrationType**](IntegrationType.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the integration instance | [optional] 
**Version** | Pointer to **string** | Version of the integration instance | [optional] 
**ReportedStatus** | Pointer to [**IntegrationReportedStatus**](IntegrationReportedStatus.md) |  | [optional] 
**IntegrationStatus** | Pointer to [**IntegrationStatus**](IntegrationStatus.md) |  | [optional] 
**LastSeen** | Pointer to **time.Time** | Timestamp of last received health report from integration instance | [optional] [readonly] 
**Uptime** | Pointer to **float32** | Running time of integration instance in seconds | [optional] 
**AccountName** | Pointer to **string** | Account that the integration instance reports to by default | [optional] 
**Username** | Pointer to **string** | User that the integration instance authenticates as by default | [optional] 
**ExplicitlyAccountBound** | Pointer to **[]string** | List of accounts that the integration instance is statically configured to handle | [optional] 
**Accounts** | Pointer to **[]string** | List of accounts that the integration instance handles | [optional] 
**ClusterName** | Pointer to **string** | Name of cluster where the integration instance runs | [optional] 
**Namespace** | Pointer to **string** | Namespace in which the integration instance runs | [optional] 
**HealthReportInterval** | Pointer to **int64** | Interval (in seconds) between health reports | [optional] 
**RegistrationId** | Pointer to **string** | identifier for the integration during registration until it has received its integration_id | [optional] 
**RegistrationInstanceId** | Pointer to **string** | Unique identifier for the integration instance (among its replicas) during registration | [optional] 

## Methods

### NewIntegrationSummary

`func NewIntegrationSummary() *IntegrationSummary`

NewIntegrationSummary instantiates a new IntegrationSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationSummaryWithDefaults

`func NewIntegrationSummaryWithDefaults() *IntegrationSummary`

NewIntegrationSummaryWithDefaults instantiates a new IntegrationSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *IntegrationSummary) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *IntegrationSummary) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *IntegrationSummary) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *IntegrationSummary) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *IntegrationSummary) GetType() IntegrationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntegrationSummary) GetTypeOk() (*IntegrationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntegrationSummary) SetType(v IntegrationType)`

SetType sets Type field to given value.

### HasType

`func (o *IntegrationSummary) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *IntegrationSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IntegrationSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *IntegrationSummary) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IntegrationSummary) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IntegrationSummary) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *IntegrationSummary) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetReportedStatus

`func (o *IntegrationSummary) GetReportedStatus() IntegrationReportedStatus`

GetReportedStatus returns the ReportedStatus field if non-nil, zero value otherwise.

### GetReportedStatusOk

`func (o *IntegrationSummary) GetReportedStatusOk() (*IntegrationReportedStatus, bool)`

GetReportedStatusOk returns a tuple with the ReportedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedStatus

`func (o *IntegrationSummary) SetReportedStatus(v IntegrationReportedStatus)`

SetReportedStatus sets ReportedStatus field to given value.

### HasReportedStatus

`func (o *IntegrationSummary) HasReportedStatus() bool`

HasReportedStatus returns a boolean if a field has been set.

### GetIntegrationStatus

`func (o *IntegrationSummary) GetIntegrationStatus() IntegrationStatus`

GetIntegrationStatus returns the IntegrationStatus field if non-nil, zero value otherwise.

### GetIntegrationStatusOk

`func (o *IntegrationSummary) GetIntegrationStatusOk() (*IntegrationStatus, bool)`

GetIntegrationStatusOk returns a tuple with the IntegrationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationStatus

`func (o *IntegrationSummary) SetIntegrationStatus(v IntegrationStatus)`

SetIntegrationStatus sets IntegrationStatus field to given value.

### HasIntegrationStatus

`func (o *IntegrationSummary) HasIntegrationStatus() bool`

HasIntegrationStatus returns a boolean if a field has been set.

### GetLastSeen

`func (o *IntegrationSummary) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *IntegrationSummary) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *IntegrationSummary) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *IntegrationSummary) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetUptime

`func (o *IntegrationSummary) GetUptime() float32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *IntegrationSummary) GetUptimeOk() (*float32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *IntegrationSummary) SetUptime(v float32)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *IntegrationSummary) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetAccountName

`func (o *IntegrationSummary) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *IntegrationSummary) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *IntegrationSummary) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *IntegrationSummary) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetUsername

`func (o *IntegrationSummary) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *IntegrationSummary) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *IntegrationSummary) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *IntegrationSummary) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetExplicitlyAccountBound

`func (o *IntegrationSummary) GetExplicitlyAccountBound() []string`

GetExplicitlyAccountBound returns the ExplicitlyAccountBound field if non-nil, zero value otherwise.

### GetExplicitlyAccountBoundOk

`func (o *IntegrationSummary) GetExplicitlyAccountBoundOk() (*[]string, bool)`

GetExplicitlyAccountBoundOk returns a tuple with the ExplicitlyAccountBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitlyAccountBound

`func (o *IntegrationSummary) SetExplicitlyAccountBound(v []string)`

SetExplicitlyAccountBound sets ExplicitlyAccountBound field to given value.

### HasExplicitlyAccountBound

`func (o *IntegrationSummary) HasExplicitlyAccountBound() bool`

HasExplicitlyAccountBound returns a boolean if a field has been set.

### GetAccounts

`func (o *IntegrationSummary) GetAccounts() []string`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *IntegrationSummary) GetAccountsOk() (*[]string, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *IntegrationSummary) SetAccounts(v []string)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *IntegrationSummary) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetClusterName

`func (o *IntegrationSummary) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *IntegrationSummary) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *IntegrationSummary) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *IntegrationSummary) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetNamespace

`func (o *IntegrationSummary) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *IntegrationSummary) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *IntegrationSummary) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *IntegrationSummary) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetHealthReportInterval

`func (o *IntegrationSummary) GetHealthReportInterval() int64`

GetHealthReportInterval returns the HealthReportInterval field if non-nil, zero value otherwise.

### GetHealthReportIntervalOk

`func (o *IntegrationSummary) GetHealthReportIntervalOk() (*int64, bool)`

GetHealthReportIntervalOk returns a tuple with the HealthReportInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthReportInterval

`func (o *IntegrationSummary) SetHealthReportInterval(v int64)`

SetHealthReportInterval sets HealthReportInterval field to given value.

### HasHealthReportInterval

`func (o *IntegrationSummary) HasHealthReportInterval() bool`

HasHealthReportInterval returns a boolean if a field has been set.

### GetRegistrationId

`func (o *IntegrationSummary) GetRegistrationId() string`

GetRegistrationId returns the RegistrationId field if non-nil, zero value otherwise.

### GetRegistrationIdOk

`func (o *IntegrationSummary) GetRegistrationIdOk() (*string, bool)`

GetRegistrationIdOk returns a tuple with the RegistrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationId

`func (o *IntegrationSummary) SetRegistrationId(v string)`

SetRegistrationId sets RegistrationId field to given value.

### HasRegistrationId

`func (o *IntegrationSummary) HasRegistrationId() bool`

HasRegistrationId returns a boolean if a field has been set.

### GetRegistrationInstanceId

`func (o *IntegrationSummary) GetRegistrationInstanceId() string`

GetRegistrationInstanceId returns the RegistrationInstanceId field if non-nil, zero value otherwise.

### GetRegistrationInstanceIdOk

`func (o *IntegrationSummary) GetRegistrationInstanceIdOk() (*string, bool)`

GetRegistrationInstanceIdOk returns a tuple with the RegistrationInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationInstanceId

`func (o *IntegrationSummary) SetRegistrationInstanceId(v string)`

SetRegistrationInstanceId sets RegistrationInstanceId field to given value.

### HasRegistrationInstanceId

`func (o *IntegrationSummary) HasRegistrationInstanceId() bool`

HasRegistrationInstanceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


