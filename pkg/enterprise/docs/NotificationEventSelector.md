# NotificationEventSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | **string** | The level of the event to filter. &#39;*&#39; matches events of all levels. &#39;info&#39; and &#39;error&#39; match related events respectively | 
**ResourceType** | **string** | The type of resource to filter. &#39;*&#39; matches all resource types. Some examples of resource type are &#39;image_digest&#39; or &#39;service&#39; | 
**Type** | **string** | The type of event to filter, using wildcards against type field of the event. Event types have a structured format &lt;category&gt;.&lt;subcategory&gt;.&lt;event&gt;. Thus, &#39;*&#39; matches all types of events. &#39;system.*&#39; matches all system events, &#39;user.*&#39; matches events that are relevant to individual consumption, and omitting an asterisk will do an exact match. See the GET /event_types route definition in the engine&#39;s external API for the list of event types.  | 

## Methods

### NewNotificationEventSelector

`func NewNotificationEventSelector(level string, resourceType string, type_ string, ) *NotificationEventSelector`

NewNotificationEventSelector instantiates a new NotificationEventSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationEventSelectorWithDefaults

`func NewNotificationEventSelectorWithDefaults() *NotificationEventSelector`

NewNotificationEventSelectorWithDefaults instantiates a new NotificationEventSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *NotificationEventSelector) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *NotificationEventSelector) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *NotificationEventSelector) SetLevel(v string)`

SetLevel sets Level field to given value.


### GetResourceType

`func (o *NotificationEventSelector) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *NotificationEventSelector) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *NotificationEventSelector) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetType

`func (o *NotificationEventSelector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationEventSelector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationEventSelector) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


