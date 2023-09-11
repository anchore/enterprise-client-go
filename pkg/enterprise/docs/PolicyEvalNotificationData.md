# PolicyEvalNotificationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationUser** | Pointer to **string** |  | [optional] 
**NotificationUserEmail** | Pointer to **string** |  | [optional] 
**NotificationType** | Pointer to **string** |  | [optional] 
**NotificationPayload** | Pointer to [**PolicyEvalNotificationPayload**](PolicyEvalNotificationPayload.md) |  | [optional] 

## Methods

### NewPolicyEvalNotificationData

`func NewPolicyEvalNotificationData() *PolicyEvalNotificationData`

NewPolicyEvalNotificationData instantiates a new PolicyEvalNotificationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyEvalNotificationDataWithDefaults

`func NewPolicyEvalNotificationDataWithDefaults() *PolicyEvalNotificationData`

NewPolicyEvalNotificationDataWithDefaults instantiates a new PolicyEvalNotificationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationUser

`func (o *PolicyEvalNotificationData) GetNotificationUser() string`

GetNotificationUser returns the NotificationUser field if non-nil, zero value otherwise.

### GetNotificationUserOk

`func (o *PolicyEvalNotificationData) GetNotificationUserOk() (*string, bool)`

GetNotificationUserOk returns a tuple with the NotificationUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationUser

`func (o *PolicyEvalNotificationData) SetNotificationUser(v string)`

SetNotificationUser sets NotificationUser field to given value.

### HasNotificationUser

`func (o *PolicyEvalNotificationData) HasNotificationUser() bool`

HasNotificationUser returns a boolean if a field has been set.

### GetNotificationUserEmail

`func (o *PolicyEvalNotificationData) GetNotificationUserEmail() string`

GetNotificationUserEmail returns the NotificationUserEmail field if non-nil, zero value otherwise.

### GetNotificationUserEmailOk

`func (o *PolicyEvalNotificationData) GetNotificationUserEmailOk() (*string, bool)`

GetNotificationUserEmailOk returns a tuple with the NotificationUserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationUserEmail

`func (o *PolicyEvalNotificationData) SetNotificationUserEmail(v string)`

SetNotificationUserEmail sets NotificationUserEmail field to given value.

### HasNotificationUserEmail

`func (o *PolicyEvalNotificationData) HasNotificationUserEmail() bool`

HasNotificationUserEmail returns a boolean if a field has been set.

### GetNotificationType

`func (o *PolicyEvalNotificationData) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *PolicyEvalNotificationData) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *PolicyEvalNotificationData) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.

### HasNotificationType

`func (o *PolicyEvalNotificationData) HasNotificationType() bool`

HasNotificationType returns a boolean if a field has been set.

### GetNotificationPayload

`func (o *PolicyEvalNotificationData) GetNotificationPayload() PolicyEvalNotificationPayload`

GetNotificationPayload returns the NotificationPayload field if non-nil, zero value otherwise.

### GetNotificationPayloadOk

`func (o *PolicyEvalNotificationData) GetNotificationPayloadOk() (*PolicyEvalNotificationPayload, bool)`

GetNotificationPayloadOk returns a tuple with the NotificationPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationPayload

`func (o *PolicyEvalNotificationData) SetNotificationPayload(v PolicyEvalNotificationPayload)`

SetNotificationPayload sets NotificationPayload field to given value.

### HasNotificationPayload

`func (o *PolicyEvalNotificationData) HasNotificationPayload() bool`

HasNotificationPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


