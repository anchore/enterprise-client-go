/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.7.2
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type EventsAPI interface {

	/*
	DeleteEvent Delete Event

	Delete an event by its event ID

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param eventId Event ID of the event to be deleted
	@return ApiDeleteEventRequest
	*/
	DeleteEvent(ctx context.Context, eventId string) ApiDeleteEventRequest

	// DeleteEventExecute executes the request
	DeleteEventExecute(r ApiDeleteEventRequest) (*http.Response, error)

	/*
	DeleteEvents Delete Events

	Delete all or a subset of events filtered using the optional query parameters

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiDeleteEventsRequest
	*/
	DeleteEvents(ctx context.Context) ApiDeleteEventsRequest

	// DeleteEventsExecute executes the request
	//  @return []string
	DeleteEventsExecute(r ApiDeleteEventsRequest) ([]string, *http.Response, error)

	/*
	GetEvent Get Event

	Lookup an event by its event ID

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param eventId Event ID of the event for lookup
	@return ApiGetEventRequest
	*/
	GetEvent(ctx context.Context, eventId string) ApiGetEventRequest

	// GetEventExecute executes the request
	//  @return EventResponse
	GetEventExecute(r ApiGetEventRequest) (*EventResponse, *http.Response, error)

	/*
	ListEventTypes List Event Types

	Returns list of event types in the category hierarchy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListEventTypesRequest
	*/
	ListEventTypes(ctx context.Context) ApiListEventTypesRequest

	// ListEventTypesExecute executes the request
	//  @return []EventCategory
	ListEventTypesExecute(r ApiListEventTypesRequest) ([]EventCategory, *http.Response, error)

	/*
	ListEvents List Events

	Returns a paginated list of events in the descending order of their occurrence. Optional query parameters may be used for filtering results

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListEventsRequest
	*/
	ListEvents(ctx context.Context) ApiListEventsRequest

	// ListEventsExecute executes the request
	//  @return EventsList
	ListEventsExecute(r ApiListEventsRequest) (*EventsList, *http.Response, error)
}

// EventsAPIService EventsAPI service
type EventsAPIService service

type ApiDeleteEventRequest struct {
	ctx context.Context
	ApiService EventsAPI
	eventId string
	xAnchoreAccount *string
}

// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiDeleteEventRequest) XAnchoreAccount(xAnchoreAccount string) ApiDeleteEventRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

func (r ApiDeleteEventRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteEventExecute(r)
}

/*
DeleteEvent Delete Event

Delete an event by its event ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param eventId Event ID of the event to be deleted
 @return ApiDeleteEventRequest
*/
func (a *EventsAPIService) DeleteEvent(ctx context.Context, eventId string) ApiDeleteEventRequest {
	return ApiDeleteEventRequest{
		ApiService: a,
		ctx: ctx,
		eventId: eventId,
	}
}

// Execute executes the request
func (a *EventsAPIService) DeleteEventExecute(r ApiDeleteEventRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsAPIService.DeleteEvent")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/events/{event_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"event_id"+"}", url.PathEscape(parameterValueToString(r.eventId, "eventId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-anchore-account", r.xAnchoreAccount, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeleteEventsRequest struct {
	ctx context.Context
	ApiService EventsAPI
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

func (r ApiDeleteEventsRequest) Execute() ([]string, *http.Response, error) {
	return r.ApiService.DeleteEventsExecute(r)
}

/*
DeleteEvents Delete Events

Delete all or a subset of events filtered using the optional query parameters

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeleteEventsRequest
*/
func (a *EventsAPIService) DeleteEvents(ctx context.Context) ApiDeleteEventsRequest {
	return ApiDeleteEventsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []string
func (a *EventsAPIService) DeleteEventsExecute(r ApiDeleteEventsRequest) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsAPIService.DeleteEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/events"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.before != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "before", r.before, "form", "")
	}
	if r.since != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "since", r.since, "form", "")
	}
	if r.level != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "level", r.level, "form", "")
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
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-anchore-account", r.xAnchoreAccount, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetEventRequest struct {
	ctx context.Context
	ApiService EventsAPI
	eventId string
	xAnchoreAccount *string
}

// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiGetEventRequest) XAnchoreAccount(xAnchoreAccount string) ApiGetEventRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

func (r ApiGetEventRequest) Execute() (*EventResponse, *http.Response, error) {
	return r.ApiService.GetEventExecute(r)
}

/*
GetEvent Get Event

Lookup an event by its event ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param eventId Event ID of the event for lookup
 @return ApiGetEventRequest
*/
func (a *EventsAPIService) GetEvent(ctx context.Context, eventId string) ApiGetEventRequest {
	return ApiGetEventRequest{
		ApiService: a,
		ctx: ctx,
		eventId: eventId,
	}
}

// Execute executes the request
//  @return EventResponse
func (a *EventsAPIService) GetEventExecute(r ApiGetEventRequest) (*EventResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EventResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsAPIService.GetEvent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/events/{event_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"event_id"+"}", url.PathEscape(parameterValueToString(r.eventId, "eventId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-anchore-account", r.xAnchoreAccount, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListEventTypesRequest struct {
	ctx context.Context
	ApiService EventsAPI
}

func (r ApiListEventTypesRequest) Execute() ([]EventCategory, *http.Response, error) {
	return r.ApiService.ListEventTypesExecute(r)
}

/*
ListEventTypes List Event Types

Returns list of event types in the category hierarchy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListEventTypesRequest
*/
func (a *EventsAPIService) ListEventTypes(ctx context.Context) ApiListEventTypesRequest {
	return ApiListEventTypesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []EventCategory
func (a *EventsAPIService) ListEventTypesExecute(r ApiListEventTypesRequest) ([]EventCategory, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []EventCategory
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsAPIService.ListEventTypes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/event-types"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListEventsRequest struct {
	ctx context.Context
	ApiService EventsAPI
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

func (r ApiListEventsRequest) Execute() (*EventsList, *http.Response, error) {
	return r.ApiService.ListEventsExecute(r)
}

/*
ListEvents List Events

Returns a paginated list of events in the descending order of their occurrence. Optional query parameters may be used for filtering results

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListEventsRequest
*/
func (a *EventsAPIService) ListEvents(ctx context.Context) ApiListEventsRequest {
	return ApiListEventsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EventsList
func (a *EventsAPIService) ListEventsExecute(r ApiListEventsRequest) (*EventsList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EventsList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsAPIService.ListEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/events"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sourceServiceName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "source_service_name", r.sourceServiceName, "form", "")
	}
	if r.sourceHostId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "source_host_id", r.sourceHostId, "form", "")
	}
	if r.eventType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "event_type", r.eventType, "form", "")
	}
	if r.resourceType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "resource_type", r.resourceType, "form", "")
	}
	if r.resourceId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "resource_id", r.resourceId, "form", "")
	}
	if r.level != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "level", r.level, "form", "")
	}
	if r.since != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "since", r.since, "form", "")
	}
	if r.before != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "before", r.before, "form", "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	} else {
		var defaultValue int32 = 1
		r.page = &defaultValue
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	} else {
		var defaultValue int32 = 100
		r.limit = &defaultValue
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
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-anchore-account", r.xAnchoreAccount, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
