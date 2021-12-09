# EventSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | **string** | The level of the event to filter. &#39;*&#39; matches events of all levels. &#39;info&#39; and &#39;error&#39; match related events respectively | 
**ResourceType** | **string** | The type of resource to filter. &#39;*&#39; matches all resource types. Some examples of resource type are &#39;image_digest&#39; or &#39;service&#39; | 
**Type** | **string** | The type of event to filter, using wildcards against type field of the event. Event types have a structured format &lt;category&gt;.&lt;subcategory&gt;.&lt;event&gt;. Thus, &#39;*&#39; matches all types of events. &#39;system.*&#39; matches all system events, &#39;user.*&#39; matches events that are relevant to individual consumption, and omitting an asterisk will do an exact match. See the GET /event_types route definition in the engine&#39;s external API for the list of event types.  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


