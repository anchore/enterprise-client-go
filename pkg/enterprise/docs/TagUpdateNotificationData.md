# TagUpdateNotificationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationUser** | Pointer to **string** |  | [optional] 
**NotificationUserEmail** | Pointer to **string** |  | [optional] 
**NotificationType** | Pointer to **string** |  | [optional] 
**NotificationPayload** | Pointer to [**TagUpdateNotificationPayload**](TagUpdateNotificationPayload.md) |  | [optional] 

## Methods

### NewTagUpdateNotificationData

`func NewTagUpdateNotificationData() *TagUpdateNotificationData`

NewTagUpdateNotificationData instantiates a new TagUpdateNotificationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagUpdateNotificationDataWithDefaults

`func NewTagUpdateNotificationDataWithDefaults() *TagUpdateNotificationData`

NewTagUpdateNotificationDataWithDefaults instantiates a new TagUpdateNotificationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationUser

`func (o *TagUpdateNotificationData) GetNotificationUser() string`

GetNotificationUser returns the NotificationUser field if non-nil, zero value otherwise.

### GetNotificationUserOk

`func (o *TagUpdateNotificationData) GetNotificationUserOk() (*string, bool)`

GetNotificationUserOk returns a tuple with the NotificationUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationUser

`func (o *TagUpdateNotificationData) SetNotificationUser(v string)`

SetNotificationUser sets NotificationUser field to given value.

### HasNotificationUser

`func (o *TagUpdateNotificationData) HasNotificationUser() bool`

HasNotificationUser returns a boolean if a field has been set.

### GetNotificationUserEmail

`func (o *TagUpdateNotificationData) GetNotificationUserEmail() string`

GetNotificationUserEmail returns the NotificationUserEmail field if non-nil, zero value otherwise.

### GetNotificationUserEmailOk

`func (o *TagUpdateNotificationData) GetNotificationUserEmailOk() (*string, bool)`

GetNotificationUserEmailOk returns a tuple with the NotificationUserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationUserEmail

`func (o *TagUpdateNotificationData) SetNotificationUserEmail(v string)`

SetNotificationUserEmail sets NotificationUserEmail field to given value.

### HasNotificationUserEmail

`func (o *TagUpdateNotificationData) HasNotificationUserEmail() bool`

HasNotificationUserEmail returns a boolean if a field has been set.

### GetNotificationType

`func (o *TagUpdateNotificationData) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *TagUpdateNotificationData) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *TagUpdateNotificationData) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.

### HasNotificationType

`func (o *TagUpdateNotificationData) HasNotificationType() bool`

HasNotificationType returns a boolean if a field has been set.

### GetNotificationPayload

`func (o *TagUpdateNotificationData) GetNotificationPayload() TagUpdateNotificationPayload`

GetNotificationPayload returns the NotificationPayload field if non-nil, zero value otherwise.

### GetNotificationPayloadOk

`func (o *TagUpdateNotificationData) GetNotificationPayloadOk() (*TagUpdateNotificationPayload, bool)`

GetNotificationPayloadOk returns a tuple with the NotificationPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationPayload

`func (o *TagUpdateNotificationData) SetNotificationPayload(v TagUpdateNotificationPayload)`

SetNotificationPayload sets NotificationPayload field to given value.

### HasNotificationPayload

`func (o *TagUpdateNotificationData) HasNotificationPayload() bool`

HasNotificationPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


