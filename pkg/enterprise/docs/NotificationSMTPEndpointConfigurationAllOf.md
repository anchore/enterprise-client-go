# NotificationSMTPEndpointConfigurationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** |  | 
**Port** | **int32** |  | 
**Username** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**UseTls** | Pointer to **bool** | Encrypt the SMTP connection with TLS. Defaults to true | [optional] 
**From** | **string** | The from address to use for emails send by this configuration | 
**To** | **string** | The address to which the emails are sent | 

## Methods

### NewNotificationSMTPEndpointConfigurationAllOf

`func NewNotificationSMTPEndpointConfigurationAllOf(host string, port int32, from string, to string, ) *NotificationSMTPEndpointConfigurationAllOf`

NewNotificationSMTPEndpointConfigurationAllOf instantiates a new NotificationSMTPEndpointConfigurationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSMTPEndpointConfigurationAllOfWithDefaults

`func NewNotificationSMTPEndpointConfigurationAllOfWithDefaults() *NotificationSMTPEndpointConfigurationAllOf`

NewNotificationSMTPEndpointConfigurationAllOfWithDefaults instantiates a new NotificationSMTPEndpointConfigurationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *NotificationSMTPEndpointConfigurationAllOf) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *NotificationSMTPEndpointConfigurationAllOf) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *NotificationSMTPEndpointConfigurationAllOf) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *NotificationSMTPEndpointConfigurationAllOf) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *NotificationSMTPEndpointConfigurationAllOf) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *NotificationSMTPEndpointConfigurationAllOf) SetPort(v int32)`

SetPort sets Port field to given value.


### GetUsername

`func (o *NotificationSMTPEndpointConfigurationAllOf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *NotificationSMTPEndpointConfigurationAllOf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *NotificationSMTPEndpointConfigurationAllOf) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *NotificationSMTPEndpointConfigurationAllOf) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *NotificationSMTPEndpointConfigurationAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *NotificationSMTPEndpointConfigurationAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *NotificationSMTPEndpointConfigurationAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *NotificationSMTPEndpointConfigurationAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUseTls

`func (o *NotificationSMTPEndpointConfigurationAllOf) GetUseTls() bool`

GetUseTls returns the UseTls field if non-nil, zero value otherwise.

### GetUseTlsOk

`func (o *NotificationSMTPEndpointConfigurationAllOf) GetUseTlsOk() (*bool, bool)`

GetUseTlsOk returns a tuple with the UseTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTls

`func (o *NotificationSMTPEndpointConfigurationAllOf) SetUseTls(v bool)`

SetUseTls sets UseTls field to given value.

### HasUseTls

`func (o *NotificationSMTPEndpointConfigurationAllOf) HasUseTls() bool`

HasUseTls returns a boolean if a field has been set.

### GetFrom

`func (o *NotificationSMTPEndpointConfigurationAllOf) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *NotificationSMTPEndpointConfigurationAllOf) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *NotificationSMTPEndpointConfigurationAllOf) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTo

`func (o *NotificationSMTPEndpointConfigurationAllOf) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *NotificationSMTPEndpointConfigurationAllOf) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *NotificationSMTPEndpointConfigurationAllOf) SetTo(v string)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


