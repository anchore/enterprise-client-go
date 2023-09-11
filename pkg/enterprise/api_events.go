/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.0.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

type EventsApi interface {

	/*
	DeleteEvent Delete Event

	Delete an event by its event ID

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param eventId Event ID of the event to be deleted
	 @return ApiDeleteEventRequest
	*/
	DeleteEvent(ctx _context.Context, eventId string) ApiDeleteEventRequest

	// DeleteEventExecute executes the request
	DeleteEventExecute(r ApiDeleteEventRequest) (*_nethttp.Response, error)

	/*
	DeleteEvents Delete Events

	Delete all or a subset of events filtered using the optional query parameters

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiDeleteEventsRequest
	*/
	DeleteEvents(ctx _context.Context) ApiDeleteEventsRequest

	// DeleteEventsExecute executes the request
	//  @return []string
	DeleteEventsExecute(r ApiDeleteEventsRequest) ([]string, *_nethttp.Response, error)

	/*
	GetEvent Get Event

	Lookup an event by its event ID

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param eventId Event ID of the event for lookup
	 @return ApiGetEventRequest
	*/
	GetEvent(ctx _context.Context, eventId string) ApiGetEventRequest

	// GetEventExecute executes the request
	//  @return EventResponse
	GetEventExecute(r ApiGetEventRequest) (EventResponse, *_nethttp.Response, error)

	/*
	ListEventTypes List Event Types

	Returns list of event types in the category hierarchy

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiListEventTypesRequest
	*/
	ListEventTypes(ctx _context.Context) ApiListEventTypesRequest

	// ListEventTypesExecute executes the request
	//  @return []EventCategory
	ListEventTypesExecute(r ApiListEventTypesRequest) ([]EventCategory, *_nethttp.Response, error)

	/*
	ListEvents List Events

	Returns a paginated list of events in the descending order of their occurrence. Optional query parameters may be used for filtering results

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiListEventsRequest
	*/
	ListEvents(ctx _context.Context) ApiListEventsRequest

	// ListEventsExecute executes the request
	//  @return EventsList
	ListEventsExecute(r ApiListEventsRequest) (EventsList, *_nethttp.Response, error)
}

// EventsApiService EventsApi service
type EventsApiService service

type ApiDeleteEventRequest struct {
	ctx _context.Context
	ApiService EventsApi
	eventId string
	xAnchoreAccount *string
}

// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiDeleteEventRequest) XAnchoreAccount(xAnchoreAccount string) ApiDeleteEventRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

func (r ApiDeleteEventRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteEventExecute(r)
}

/*
DeleteEvent Delete Event

Delete an event by its event ID

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param eventId Event ID of the event to be deleted
 @return ApiDeleteEventRequest
*/
func (a *EventsApiService) DeleteEvent(ctx _context.Context, eventId string) ApiDeleteEventRequest {
	return ApiDeleteEventRequest{
		ApiService: a,
		ctx: ctx,
		eventId: eventId,
	}
}

// Execute executes the request
func (a *EventsApiService) DeleteEventExecute(r ApiDeleteEventRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.DeleteEvent")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/events/{event_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"event_id"+"}", _neturl.PathEscape(parameterToString(r.eventId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xAnchoreAccount != nil {
		localVarHeaderParams["x-anchore-account"] = parameterToString(*r.xAnchoreAccount, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeleteEventsRequest struct {
	ctx _context.Context
	ApiService EventsApi
	before *string
	since *string
	level *string
	xAnchoreAccount *string
}

// Delete events that occurred before the timestamp
func (r ApiDeleteEventsRequest) Before(before string) ApiDeleteEventsRequest {
	r.before = &before
	return r
}
// Delete events that occurred after the timestamp
func (r ApiDeleteEventsRequest) Since(since string) ApiDeleteEventsRequest {
	r.since = &since
	return r
}
// Delete events that match the level - INFO or ERROR
func (r ApiDeleteEventsRequest) Level(level string) ApiDeleteEventsRequest {
	r.level = &level
	return r
}
// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiDeleteEventsRequest) XAnchoreAccount(xAnchoreAccount string) ApiDeleteEventsRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

func (r ApiDeleteEventsRequest) Execute() ([]string, *_nethttp.Response, error) {
	return r.ApiService.DeleteEventsExecute(r)
}

/*
DeleteEvents Delete Events

Delete all or a subset of events filtered using the optional query parameters

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeleteEventsRequest
*/
func (a *EventsApiService) DeleteEvents(ctx _context.Context) ApiDeleteEventsRequest {
	return ApiDeleteEventsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []string
func (a *EventsApiService) DeleteEventsExecute(r ApiDeleteEventsRequest) ([]string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.DeleteEvents")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/events"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.before != nil {
		localVarQueryParams.Add("before", parameterToString(*r.before, ""))
	}
	if r.since != nil {
		localVarQueryParams.Add("since", parameterToString(*r.since, ""))
	}
	if r.level != nil {
		localVarQueryParams.Add("level", parameterToString(*r.level, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xAnchoreAccount != nil {
		localVarHeaderParams["x-anchore-account"] = parameterToString(*r.xAnchoreAccount, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetEventRequest struct {
	ctx _context.Context
	ApiService EventsApi
	eventId string
	xAnchoreAccount *string
}

// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiGetEventRequest) XAnchoreAccount(xAnchoreAccount string) ApiGetEventRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

func (r ApiGetEventRequest) Execute() (EventResponse, *_nethttp.Response, error) {
	return r.ApiService.GetEventExecute(r)
}

/*
GetEvent Get Event

Lookup an event by its event ID

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param eventId Event ID of the event for lookup
 @return ApiGetEventRequest
*/
func (a *EventsApiService) GetEvent(ctx _context.Context, eventId string) ApiGetEventRequest {
	return ApiGetEventRequest{
		ApiService: a,
		ctx: ctx,
		eventId: eventId,
	}
}

// Execute executes the request
//  @return EventResponse
func (a *EventsApiService) GetEventExecute(r ApiGetEventRequest) (EventResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EventResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.GetEvent")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/events/{event_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"event_id"+"}", _neturl.PathEscape(parameterToString(r.eventId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xAnchoreAccount != nil {
		localVarHeaderParams["x-anchore-account"] = parameterToString(*r.xAnchoreAccount, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListEventTypesRequest struct {
	ctx _context.Context
	ApiService EventsApi
}


func (r ApiListEventTypesRequest) Execute() ([]EventCategory, *_nethttp.Response, error) {
	return r.ApiService.ListEventTypesExecute(r)
}

/*
ListEventTypes List Event Types

Returns list of event types in the category hierarchy

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListEventTypesRequest
*/
func (a *EventsApiService) ListEventTypes(ctx _context.Context) ApiListEventTypesRequest {
	return ApiListEventTypesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []EventCategory
func (a *EventsApiService) ListEventTypesExecute(r ApiListEventTypesRequest) ([]EventCategory, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []EventCategory
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.ListEventTypes")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/event-types"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListEventsRequest struct {
	ctx _context.Context
	ApiService EventsApi
	sourceServiceName *string
	sourceHostId *string
	eventType *string
	resourceType *string
	resourceId *string
	level *string
	since *string
	before *string
	page *int32
	limit *int32
	xAnchoreAccount *string
}

// Filter events by the originating service
func (r ApiListEventsRequest) SourceServiceName(sourceServiceName string) ApiListEventsRequest {
	r.sourceServiceName = &sourceServiceName
	return r
}
// Filter events by the originating host ID
func (r ApiListEventsRequest) SourceHostId(sourceHostId string) ApiListEventsRequest {
	r.sourceHostId = &sourceHostId
	return r
}
// Filter events by a prefix match on the event type (e.g. \&quot;user.image.\&quot;)
func (r ApiListEventsRequest) EventType(eventType string) ApiListEventsRequest {
	r.eventType = &eventType
	return r
}
// Filter events by the type of resource - tag, image_digest, repository etc
func (r ApiListEventsRequest) ResourceType(resourceType string) ApiListEventsRequest {
	r.resourceType = &resourceType
	return r
}
// Filter events by the id of the resource
func (r ApiListEventsRequest) ResourceId(resourceId string) ApiListEventsRequest {
	r.resourceId = &resourceId
	return r
}
// Filter events by the level - INFO or ERROR
func (r ApiListEventsRequest) Level(level string) ApiListEventsRequest {
	r.level = &level
	return r
}
// Return events that occurred after the timestamp
func (r ApiListEventsRequest) Since(since string) ApiListEventsRequest {
	r.since = &since
	return r
}
// Return events that occurred before the timestamp
func (r ApiListEventsRequest) Before(before string) ApiListEventsRequest {
	r.before = &before
	return r
}
// Pagination controls - return the nth page of results. Defaults to first page if left empty
func (r ApiListEventsRequest) Page(page int32) ApiListEventsRequest {
	r.page = &page
	return r
}
// Number of events in the result set. Defaults to 100 if left empty
func (r ApiListEventsRequest) Limit(limit int32) ApiListEventsRequest {
	r.limit = &limit
	return r
}
// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiListEventsRequest) XAnchoreAccount(xAnchoreAccount string) ApiListEventsRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

func (r ApiListEventsRequest) Execute() (EventsList, *_nethttp.Response, error) {
	return r.ApiService.ListEventsExecute(r)
}

/*
ListEvents List Events

Returns a paginated list of events in the descending order of their occurrence. Optional query parameters may be used for filtering results

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListEventsRequest
*/
func (a *EventsApiService) ListEvents(ctx _context.Context) ApiListEventsRequest {
	return ApiListEventsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EventsList
func (a *EventsApiService) ListEventsExecute(r ApiListEventsRequest) (EventsList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EventsList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.ListEvents")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/events"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.sourceServiceName != nil {
		localVarQueryParams.Add("source_service_name", parameterToString(*r.sourceServiceName, ""))
	}
	if r.sourceHostId != nil {
		localVarQueryParams.Add("source_host_id", parameterToString(*r.sourceHostId, ""))
	}
	if r.eventType != nil {
		localVarQueryParams.Add("event_type", parameterToString(*r.eventType, ""))
	}
	if r.resourceType != nil {
		localVarQueryParams.Add("resource_type", parameterToString(*r.resourceType, ""))
	}
	if r.resourceId != nil {
		localVarQueryParams.Add("resource_id", parameterToString(*r.resourceId, ""))
	}
	if r.level != nil {
		localVarQueryParams.Add("level", parameterToString(*r.level, ""))
	}
	if r.since != nil {
		localVarQueryParams.Add("since", parameterToString(*r.since, ""))
	}
	if r.before != nil {
		localVarQueryParams.Add("before", parameterToString(*r.before, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xAnchoreAccount != nil {
		localVarHeaderParams["x-anchore-account"] = parameterToString(*r.xAnchoreAccount, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
