# NotificationWebhookEndpointConfigurationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | url to POST to, including any query parameters, should begin with &#39;http://&#39; or &#39;https://&#39; | 
**Username** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**VerifySsl** | Pointer to **bool** | Verify SSL certificates for HTTPS requests, disabled by default | [optional] 

## Methods

### NewNotificationWebhookEndpointConfigurationAllOf

`func NewNotificationWebhookEndpointConfigurationAllOf(url string, ) *NotificationWebhookEndpointConfigurationAllOf`

NewNotificationWebhookEndpointConfigurationAllOf instantiates a new NotificationWebhookEndpointConfigurationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationWebhookEndpointConfigurationAllOfWithDefaults

`func NewNotificationWebhookEndpointConfigurationAllOfWithDefaults() *NotificationWebhookEndpointConfigurationAllOf`

NewNotificationWebhookEndpointConfigurationAllOfWithDefaults instantiates a new NotificationWebhookEndpointConfigurationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *NotificationWebhookEndpointConfigurationAllOf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NotificationWebhookEndpointConfigurationAllOf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NotificationWebhookEndpointConfigurationAllOf) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUsername

`func (o *NotificationWebhookEndpointConfigurationAllOf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *NotificationWebhookEndpointConfigurationAllOf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *NotificationWebhookEndpointConfigurationAllOf) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *NotificationWebhookEndpointConfigurationAllOf) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *NotificationWebhookEndpointConfigurationAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *NotificationWebhookEndpointConfigurationAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *NotificationWebhookEndpointConfigurationAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *NotificationWebhookEndpointConfigurationAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetVerifySsl

`func (o *NotificationWebhookEndpointConfigurationAllOf) GetVerifySsl() bool`

GetVerifySsl returns the VerifySsl field if non-nil, zero value otherwise.

### GetVerifySslOk

`func (o *NotificationWebhookEndpointConfigurationAllOf) GetVerifySslOk() (*bool, bool)`

GetVerifySslOk returns a tuple with the VerifySsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySsl

`func (o *NotificationWebhookEndpointConfigurationAllOf) SetVerifySsl(v bool)`

SetVerifySsl sets VerifySsl field to given value.

### HasVerifySsl

`func (o *NotificationWebhookEndpointConfigurationAllOf) HasVerifySsl() bool`

HasVerifySsl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


