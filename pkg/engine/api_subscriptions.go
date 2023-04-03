/*
Anchore Engine API Server

This is the Anchore Engine API. Provides the primary external API for users of the service.

API version: 0.6.0
Contact: nurmi@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package engine

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

type SubscriptionsApi interface {

	/*
	AddSubscription Add a subscription of a specific type

	Create a new subscription to watch a tag and get notifications of changes

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiAddSubscriptionRequest
	*/
	AddSubscription(ctx _context.Context) ApiAddSubscriptionRequest

	// AddSubscriptionExecute executes the request
	//  @return []Subscription
	AddSubscriptionExecute(r ApiAddSubscriptionRequest) ([]Subscription, *_nethttp.Response, error)

	/*
	DeleteSubscription Delete subscriptions of a specific type

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param subscriptionId
	 @return ApiDeleteSubscriptionRequest
	*/
	DeleteSubscription(ctx _context.Context, subscriptionId string) ApiDeleteSubscriptionRequest

	// DeleteSubscriptionExecute executes the request
	DeleteSubscriptionExecute(r ApiDeleteSubscriptionRequest) (*_nethttp.Response, error)

	/*
	GetSubscription Get a specific subscription set

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param subscriptionId
	 @return ApiGetSubscriptionRequest
	*/
	GetSubscription(ctx _context.Context, subscriptionId string) ApiGetSubscriptionRequest

	// GetSubscriptionExecute executes the request
	//  @return []Subscription
	GetSubscriptionExecute(r ApiGetSubscriptionRequest) ([]Subscription, *_nethttp.Response, error)

	/*
	ListSubscriptions List all subscriptions

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiListSubscriptionsRequest
	*/
	ListSubscriptions(ctx _context.Context) ApiListSubscriptionsRequest

	// ListSubscriptionsExecute executes the request
	//  @return []Subscription
	ListSubscriptionsExecute(r ApiListSubscriptionsRequest) ([]Subscription, *_nethttp.Response, error)

	/*
	UpdateSubscription Update an existing and specific subscription

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param subscriptionId
	 @return ApiUpdateSubscriptionRequest
	*/
	UpdateSubscription(ctx _context.Context, subscriptionId string) ApiUpdateSubscriptionRequest

	// UpdateSubscriptionExecute executes the request
	//  @return []Subscription
	UpdateSubscriptionExecute(r ApiUpdateSubscriptionRequest) ([]Subscription, *_nethttp.Response, error)
}

// SubscriptionsApiService SubscriptionsApi service
type SubscriptionsApiService service

type ApiAddSubscriptionRequest struct {
	ctx _context.Context
	ApiService SubscriptionsApi
	subscription *SubscriptionRequest
	xAnchoreAccount *string
}

func (r ApiAddSubscriptionRequest) Subscription(subscription SubscriptionRequest) ApiAddSubscriptionRequest {
	r.subscription = &subscription
	return r
}
// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiAddSubscriptionRequest) XAnchoreAccount(xAnchoreAccount string) ApiAddSubscriptionRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

func (r ApiAddSubscriptionRequest) Execute() ([]Subscription, *_nethttp.Response, error) {
	return r.ApiService.AddSubscriptionExecute(r)
}

/*
AddSubscription Add a subscription of a specific type

Create a new subscription to watch a tag and get notifications of changes

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAddSubscriptionRequest
*/
func (a *SubscriptionsApiService) AddSubscription(ctx _context.Context) ApiAddSubscriptionRequest {
	return ApiAddSubscriptionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Subscription
func (a *SubscriptionsApiService) AddSubscriptionExecute(r ApiAddSubscriptionRequest) ([]Subscription, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []Subscription
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionsApiService.AddSubscription")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscriptions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.subscription == nil {
		return localVarReturnValue, nil, reportError("subscription is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.subscription
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

type ApiDeleteSubscriptionRequest struct {
	ctx _context.Context
	ApiService SubscriptionsApi
	subscriptionId string
	xAnchoreAccount *string
}

// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiDeleteSubscriptionRequest) XAnchoreAccount(xAnchoreAccount string) ApiDeleteSubscriptionRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

func (r ApiDeleteSubscriptionRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteSubscriptionExecute(r)
}

/*
DeleteSubscription Delete subscriptions of a specific type

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param subscriptionId
 @return ApiDeleteSubscriptionRequest
*/
func (a *SubscriptionsApiService) DeleteSubscription(ctx _context.Context, subscriptionId string) ApiDeleteSubscriptionRequest {
	return ApiDeleteSubscriptionRequest{
		ApiService: a,
		ctx: ctx,
		subscriptionId: subscriptionId,
	}
}

// Execute executes the request
func (a *SubscriptionsApiService) DeleteSubscriptionExecute(r ApiDeleteSubscriptionRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionsApiService.DeleteSubscription")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscriptions/{subscriptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionId"+"}", _neturl.PathEscape(parameterToString(r.subscriptionId, "")), -1)

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
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetSubscriptionRequest struct {
	ctx _context.Context
	ApiService SubscriptionsApi
	subscriptionId string
	xAnchoreAccount *string
}

// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiGetSubscriptionRequest) XAnchoreAccount(xAnchoreAccount string) ApiGetSubscriptionRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

func (r ApiGetSubscriptionRequest) Execute() ([]Subscription, *_nethttp.Response, error) {
	return r.ApiService.GetSubscriptionExecute(r)
}

/*
GetSubscription Get a specific subscription set

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param subscriptionId
 @return ApiGetSubscriptionRequest
*/
func (a *SubscriptionsApiService) GetSubscription(ctx _context.Context, subscriptionId string) ApiGetSubscriptionRequest {
	return ApiGetSubscriptionRequest{
		ApiService: a,
		ctx: ctx,
		subscriptionId: subscriptionId,
	}
}

// Execute executes the request
//  @return []Subscription
func (a *SubscriptionsApiService) GetSubscriptionExecute(r ApiGetSubscriptionRequest) ([]Subscription, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []Subscription
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionsApiService.GetSubscription")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscriptions/{subscriptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionId"+"}", _neturl.PathEscape(parameterToString(r.subscriptionId, "")), -1)

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

type ApiListSubscriptionsRequest struct {
	ctx _context.Context
	ApiService SubscriptionsApi
	subscriptionKey *string
	subscriptionType *string
	xAnchoreAccount *string
}

// filter only subscriptions matching key
func (r ApiListSubscriptionsRequest) SubscriptionKey(subscriptionKey string) ApiListSubscriptionsRequest {
	r.subscriptionKey = &subscriptionKey
	return r
}
// filter only subscriptions matching type
func (r ApiListSubscriptionsRequest) SubscriptionType(subscriptionType string) ApiListSubscriptionsRequest {
	r.subscriptionType = &subscriptionType
	return r
}
// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiListSubscriptionsRequest) XAnchoreAccount(xAnchoreAccount string) ApiListSubscriptionsRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

func (r ApiListSubscriptionsRequest) Execute() ([]Subscription, *_nethttp.Response, error) {
	return r.ApiService.ListSubscriptionsExecute(r)
}

/*
ListSubscriptions List all subscriptions

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListSubscriptionsRequest
*/
func (a *SubscriptionsApiService) ListSubscriptions(ctx _context.Context) ApiListSubscriptionsRequest {
	return ApiListSubscriptionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Subscription
func (a *SubscriptionsApiService) ListSubscriptionsExecute(r ApiListSubscriptionsRequest) ([]Subscription, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []Subscription
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionsApiService.ListSubscriptions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscriptions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.subscriptionKey != nil {
		localVarQueryParams.Add("subscription_key", parameterToString(*r.subscriptionKey, ""))
	}
	if r.subscriptionType != nil {
		localVarQueryParams.Add("subscription_type", parameterToString(*r.subscriptionType, ""))
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

type ApiUpdateSubscriptionRequest struct {
	ctx _context.Context
	ApiService SubscriptionsApi
	subscriptionId string
	subscription *SubscriptionUpdate
	xAnchoreAccount *string
}

func (r ApiUpdateSubscriptionRequest) Subscription(subscription SubscriptionUpdate) ApiUpdateSubscriptionRequest {
	r.subscription = &subscription
	return r
}
// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiUpdateSubscriptionRequest) XAnchoreAccount(xAnchoreAccount string) ApiUpdateSubscriptionRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

func (r ApiUpdateSubscriptionRequest) Execute() ([]Subscription, *_nethttp.Response, error) {
	return r.ApiService.UpdateSubscriptionExecute(r)
}

/*
UpdateSubscription Update an existing and specific subscription

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param subscriptionId
 @return ApiUpdateSubscriptionRequest
*/
func (a *SubscriptionsApiService) UpdateSubscription(ctx _context.Context, subscriptionId string) ApiUpdateSubscriptionRequest {
	return ApiUpdateSubscriptionRequest{
		ApiService: a,
		ctx: ctx,
		subscriptionId: subscriptionId,
	}
}

// Execute executes the request
//  @return []Subscription
func (a *SubscriptionsApiService) UpdateSubscriptionExecute(r ApiUpdateSubscriptionRequest) ([]Subscription, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []Subscription
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionsApiService.UpdateSubscription")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscriptions/{subscriptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionId"+"}", _neturl.PathEscape(parameterToString(r.subscriptionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.subscription == nil {
		return localVarReturnValue, nil, reportError("subscription is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.subscription
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
