# NotificationTeamsEndpointConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | The instance identifier for the configuration | [optional] 
**Description** | Pointer to **string** | User friendly name or description for the configuration | [optional] 
**VerifyTls** | Pointer to **bool** | Verify the cert if using tls for connecting externally. Defaults to true if not specified | [optional] 
**CreatedAt** | Pointer to **time.Time** | Timestamp for last modification to the record | [optional] 
**LastUpdated** | Pointer to **time.Time** | Timestamp for last modification to the record | [optional] 
**Url** | **string** | url to POST to, including any query parameters, should begin with &#39;http://&#39; or &#39;https://&#39; | 

## Methods

### NewNotificationTeamsEndpointConfiguration

`func NewNotificationTeamsEndpointConfiguration(url string, ) *NotificationTeamsEndpointConfiguration`

NewNotificationTeamsEndpointConfiguration instantiates a new NotificationTeamsEndpointConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationTeamsEndpointConfigurationWithDefaults

`func NewNotificationTeamsEndpointConfigurationWithDefaults() *NotificationTeamsEndpointConfiguration`

NewNotificationTeamsEndpointConfigurationWithDefaults instantiates a new NotificationTeamsEndpointConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *NotificationTeamsEndpointConfiguration) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *NotificationTeamsEndpointConfiguration) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *NotificationTeamsEndpointConfiguration) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *NotificationTeamsEndpointConfiguration) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetDescription

`func (o *NotificationTeamsEndpointConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NotificationTeamsEndpointConfiguration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NotificationTeamsEndpointConfiguration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NotificationTeamsEndpointConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVerifyTls

`func (o *NotificationTeamsEndpointConfiguration) GetVerifyTls() bool`

GetVerifyTls returns the VerifyTls field if non-nil, zero value otherwise.

### GetVerifyTlsOk

`func (o *NotificationTeamsEndpointConfiguration) GetVerifyTlsOk() (*bool, bool)`

GetVerifyTlsOk returns a tuple with the VerifyTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyTls

`func (o *NotificationTeamsEndpointConfiguration) SetVerifyTls(v bool)`

SetVerifyTls sets VerifyTls field to given value.

### HasVerifyTls

`func (o *NotificationTeamsEndpointConfiguration) HasVerifyTls() bool`

HasVerifyTls returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NotificationTeamsEndpointConfiguration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationTeamsEndpointConfiguration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationTeamsEndpointConfiguration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NotificationTeamsEndpointConfiguration) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdated

`func (o *NotificationTeamsEndpointConfiguration) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *NotificationTeamsEndpointConfiguration) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *NotificationTeamsEndpointConfiguration) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *NotificationTeamsEndpointConfiguration) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetUrl

`func (o *NotificationTeamsEndpointConfiguration) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NotificationTeamsEndpointConfiguration) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NotificationTeamsEndpointConfiguration) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


