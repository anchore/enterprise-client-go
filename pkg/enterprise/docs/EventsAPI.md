# \EventsAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEvent**](EventsAPI.md#DeleteEvent) | **Delete** /events/{event_id} | Delete Event
[**DeleteEvents**](EventsAPI.md#DeleteEvents) | **Delete** /events | Delete Events
[**GetEvent**](EventsAPI.md#GetEvent) | **Get** /events/{event_id} | Get Event
[**ListEventTypes**](EventsAPI.md#ListEventTypes) | **Get** /event-types | List Event Types
[**ListEvents**](EventsAPI.md#ListEvents) | **Get** /events | List Events



## DeleteEvent

> DeleteEvent(ctx, eventId).XAnchoreAccount(xAnchoreAccount).Execute()

Delete Event



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	eventId := "eventId_example" // string | Event ID of the event to be deleted
	xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.DeleteEvent(context.Background(), eventId).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.DeleteEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** | Event ID of the event to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEvents

> []string DeleteEvents(ctx).Before(before).Since(since).Level(level).XAnchoreAccount(xAnchoreAccount).Execute()

Delete Events



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	before := "before_example" // string | Delete events that occurred before the timestamp (optional)
	since := "since_example" // string | Delete events that occurred after the timestamp (optional)
	level := "level_example" // string | Delete events that match the level - INFO or ERROR (optional)
	xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.DeleteEvents(context.Background()).Before(before).Since(since).Level(level).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.DeleteEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteEvents`: []string
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.DeleteEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **before** | **string** | Delete events that occurred before the timestamp | 
 **since** | **string** | Delete events that occurred after the timestamp | 
 **level** | **string** | Delete events that match the level - INFO or ERROR | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvent

> EventResponse GetEvent(ctx, eventId).XAnchoreAccount(xAnchoreAccount).Execute()

Get Event



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	eventId := "eventId_example" // string | Event ID of the event for lookup
	xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.GetEvent(context.Background(), eventId).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvent`: EventResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** | Event ID of the event for lookup | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**EventResponse**](EventResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEventTypes

> []EventCategory ListEventTypes(ctx).Execute()

List Event Types



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.ListEventTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.ListEventTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEventTypes`: []EventCategory
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.ListEventTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListEventTypesRequest struct via the builder pattern


### Return type

[**[]EventCategory**](EventCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEvents

> EventsList ListEvents(ctx).SourceServiceName(sourceServiceName).SourceHostId(sourceHostId).EventType(eventType).ResourceType(resourceType).ResourceId(resourceId).Level(level).Since(since).Before(before).Page(page).Limit(limit).XAnchoreAccount(xAnchoreAccount).Execute()

List Events



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	sourceServiceName := "sourceServiceName_example" // string | Filter events by the originating service (optional)
	sourceHostId := "sourceHostId_example" // string | Filter events by the originating host ID (optional)
	eventType := "eventType_example" // string | Filter events by a prefix match on the event type (e.g. \"user.image.\") (optional)
	resourceType := "resourceType_example" // string | Filter events by the type of resource - tag, image_digest, repository etc (optional)
	resourceId := "resourceId_example" // string | Filter events by the id of the resource (optional)
	level := "level_example" // string | Filter events by the level - INFO or ERROR (optional)
	since := "since_example" // string | Return events that occurred after the timestamp (optional)
	before := "before_example" // string | Return events that occurred before the timestamp (optional)
	page := int32(56) // int32 | Pagination controls - return the nth page of results. Defaults to first page if left empty (optional) (default to 1)
	limit := int32(56) // int32 | Number of events in the result set. Defaults to 100 if left empty (optional) (default to 100)
	xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.ListEvents(context.Background()).SourceServiceName(sourceServiceName).SourceHostId(sourceHostId).EventType(eventType).ResourceType(resourceType).ResourceId(resourceId).Level(level).Since(since).Before(before).Page(page).Limit(limit).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.ListEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEvents`: EventsList
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.ListEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceServiceName** | **string** | Filter events by the originating service | 
 **sourceHostId** | **string** | Filter events by the originating host ID | 
 **eventType** | **string** | Filter events by a prefix match on the event type (e.g. \&quot;user.image.\&quot;) | 
 **resourceType** | **string** | Filter events by the type of resource - tag, image_digest, repository etc | 
 **resourceId** | **string** | Filter events by the id of the resource | 
 **level** | **string** | Filter events by the level - INFO or ERROR | 
 **since** | **string** | Return events that occurred after the timestamp | 
 **before** | **string** | Return events that occurred before the timestamp | 
 **page** | **int32** | Pagination controls - return the nth page of results. Defaults to first page if left empty | [default to 1]
 **limit** | **int32** | Number of events in the result set. Defaults to 100 if left empty | [default to 100]
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**EventsList**](EventsList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

