# NotificationEndpointConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | The instance identifier for the configuration | [optional] 
**Description** | Pointer to **string** | User friendly name or description for the configuration | [optional] 
**VerifyTls** | Pointer to **bool** | Verify the cert if using tls for connecting externally. Defaults to true if not specified | [optional] 
**CreatedAt** | Pointer to **time.Time** | Timestamp for last modification to the record | [optional] 
**LastUpdated** | Pointer to **time.Time** | Timestamp for last modification to the record | [optional] 

## Methods

### NewNotificationEndpointConfiguration

`func NewNotificationEndpointConfiguration() *NotificationEndpointConfiguration`

NewNotificationEndpointConfiguration instantiates a new NotificationEndpointConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationEndpointConfigurationWithDefaults

`func NewNotificationEndpointConfigurationWithDefaults() *NotificationEndpointConfiguration`

NewNotificationEndpointConfigurationWithDefaults instantiates a new NotificationEndpointConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *NotificationEndpointConfiguration) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *NotificationEndpointConfiguration) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *NotificationEndpointConfiguration) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *NotificationEndpointConfiguration) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetDescription

`func (o *NotificationEndpointConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NotificationEndpointConfiguration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NotificationEndpointConfiguration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NotificationEndpointConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVerifyTls

`func (o *NotificationEndpointConfiguration) GetVerifyTls() bool`

GetVerifyTls returns the VerifyTls field if non-nil, zero value otherwise.

### GetVerifyTlsOk

`func (o *NotificationEndpointConfiguration) GetVerifyTlsOk() (*bool, bool)`

GetVerifyTlsOk returns a tuple with the VerifyTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyTls

`func (o *NotificationEndpointConfiguration) SetVerifyTls(v bool)`

SetVerifyTls sets VerifyTls field to given value.

### HasVerifyTls

`func (o *NotificationEndpointConfiguration) HasVerifyTls() bool`

HasVerifyTls returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NotificationEndpointConfiguration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationEndpointConfiguration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationEndpointConfiguration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NotificationEndpointConfiguration) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdated

`func (o *NotificationEndpointConfiguration) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *NotificationEndpointConfiguration) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *NotificationEndpointConfiguration) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *NotificationEndpointConfiguration) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


