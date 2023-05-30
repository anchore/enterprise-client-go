/*
Anchore Enterprise API Server

This is the Anchore Enterprise API. It provides additional external API routes and functionality for enterprise users.

API version: 0.8.0
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
	"time"
)

// Linger please
var (
	_ _context.Context
)

type ActionsApi interface {

	/*
	AddActionPlan Submits an Action Plan

	Submits an Action Plan and saves upon completion

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiAddActionPlanRequest
	*/
	AddActionPlan(ctx _context.Context) ApiAddActionPlanRequest

	// AddActionPlanExecute executes the request
	//  @return ActionPlan
	AddActionPlanExecute(r ApiAddActionPlanRequest) (ActionPlan, *_nethttp.Response, error)

	/*
	GetActionPlans Gets a list of submitted action (remediation) plans

	Retrieves a list of action plans that have been completed

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiGetActionPlansRequest
	*/
	GetActionPlans(ctx _context.Context) ApiGetActionPlansRequest

	// GetActionPlansExecute executes the request
	//  @return []ActionPlan
	GetActionPlansExecute(r ApiGetActionPlansRequest) ([]ActionPlan, *_nethttp.Response, error)
}

// ActionsApiService ActionsApi service
type ActionsApiService service

type ApiAddActionPlanRequest struct {
	ctx _context.Context
	ApiService ActionsApi
	actionPlan *ActionPlan
}

func (r ApiAddActionPlanRequest) ActionPlan(actionPlan ActionPlan) ApiAddActionPlanRequest {
	r.actionPlan = &actionPlan
	return r
}

func (r ApiAddActionPlanRequest) Execute() (ActionPlan, *_nethttp.Response, error) {
	return r.ApiService.AddActionPlanExecute(r)
}

/*
AddActionPlan Submits an Action Plan

Submits an Action Plan and saves upon completion

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAddActionPlanRequest
*/
func (a *ActionsApiService) AddActionPlan(ctx _context.Context) ApiAddActionPlanRequest {
	return ApiAddActionPlanRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ActionPlan
func (a *ActionsApiService) AddActionPlanExecute(r ApiAddActionPlanRequest) (ActionPlan, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ActionPlan
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActionsApiService.AddActionPlan")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/actions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.actionPlan == nil {
		return localVarReturnValue, nil, reportError("actionPlan is required and must be specified")
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
	// body params
	localVarPostBody = r.actionPlan
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

type ApiGetActionPlansRequest struct {
	ctx _context.Context
	ApiService ActionsApi
	imageTag *string
	imageDigest *string
	createdAfter *time.Time
	xAnchoreAccount *string
}

func (r ApiGetActionPlansRequest) ImageTag(imageTag string) ApiGetActionPlansRequest {
	r.imageTag = &imageTag
	return r
}
func (r ApiGetActionPlansRequest) ImageDigest(imageDigest string) ApiGetActionPlansRequest {
	r.imageDigest = &imageDigest
	return r
}
// RFC 3339 formatted UTC timestamp to filter out action plans that were only created after this date
func (r ApiGetActionPlansRequest) CreatedAfter(createdAfter time.Time) ApiGetActionPlansRequest {
	r.createdAfter = &createdAfter
	return r
}
// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiGetActionPlansRequest) XAnchoreAccount(xAnchoreAccount string) ApiGetActionPlansRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

func (r ApiGetActionPlansRequest) Execute() ([]ActionPlan, *_nethttp.Response, error) {
	return r.ApiService.GetActionPlansExecute(r)
}

/*
GetActionPlans Gets a list of submitted action (remediation) plans

Retrieves a list of action plans that have been completed

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetActionPlansRequest
*/
func (a *ActionsApiService) GetActionPlans(ctx _context.Context) ApiGetActionPlansRequest {
	return ApiGetActionPlansRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ActionPlan
func (a *ActionsApiService) GetActionPlansExecute(r ApiGetActionPlansRequest) ([]ActionPlan, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []ActionPlan
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActionsApiService.GetActionPlans")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/actions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.imageTag != nil {
		localVarQueryParams.Add("image_tag", parameterToString(*r.imageTag, ""))
	}
	if r.imageDigest != nil {
		localVarQueryParams.Add("image_digest", parameterToString(*r.imageDigest, ""))
	}
	if r.createdAfter != nil {
		localVarQueryParams.Add("created_after", parameterToString(*r.createdAfter, ""))
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
