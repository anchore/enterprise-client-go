# NotificationDispatchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Notification** | Pointer to **interface{}** | The notification object that was sent. Not documented in more detail at present due to different notification schemas exposed | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **string** | Additional delivery result details | [optional] 

## Methods

### NewNotificationDispatchResult

`func NewNotificationDispatchResult() *NotificationDispatchResult`

NewNotificationDispatchResult instantiates a new NotificationDispatchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationDispatchResultWithDefaults

`func NewNotificationDispatchResultWithDefaults() *NotificationDispatchResult`

NewNotificationDispatchResultWithDefaults instantiates a new NotificationDispatchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotification

`func (o *NotificationDispatchResult) GetNotification() interface{}`

GetNotification returns the Notification field if non-nil, zero value otherwise.

### GetNotificationOk

`func (o *NotificationDispatchResult) GetNotificationOk() (*interface{}, bool)`

GetNotificationOk returns a tuple with the Notification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotification

`func (o *NotificationDispatchResult) SetNotification(v interface{})`

SetNotification sets Notification field to given value.

### HasNotification

`func (o *NotificationDispatchResult) HasNotification() bool`

HasNotification returns a boolean if a field has been set.

### GetStatus

`func (o *NotificationDispatchResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NotificationDispatchResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NotificationDispatchResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NotificationDispatchResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDetails

`func (o *NotificationDispatchResult) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *NotificationDispatchResult) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *NotificationDispatchResult) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *NotificationDispatchResult) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


