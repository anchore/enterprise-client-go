/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.9.0
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
	"time"
)


type ActionsAPI interface {

	/*
	AddActionPlan Submits an Action Plan

	Submits an Action Plan and saves upon completion

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddActionPlanRequest
	*/
	AddActionPlan(ctx context.Context) ApiAddActionPlanRequest

	// AddActionPlanExecute executes the request
	//  @return ActionPlan
	AddActionPlanExecute(r ApiAddActionPlanRequest) (*ActionPlan, *http.Response, error)

	/*
	GetActionPlans Gets a list of submitted action (remediation) plans

	Retrieves a list of action plans that have been completed

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetActionPlansRequest
	*/
	GetActionPlans(ctx context.Context) ApiGetActionPlansRequest

	// GetActionPlansExecute executes the request
	//  @return []ActionPlan
	GetActionPlansExecute(r ApiGetActionPlansRequest) ([]ActionPlan, *http.Response, error)
}

// ActionsAPIService ActionsAPI service
type ActionsAPIService service

type ApiAddActionPlanRequest struct {
	ctx context.Context
	ApiService ActionsAPI
	actionPlan *ActionPlan
}

func (r ApiAddActionPlanRequest) ActionPlan(actionPlan ActionPlan) ApiAddActionPlanRequest {
	r.actionPlan = &actionPlan
	return r
}

func (r ApiAddActionPlanRequest) Execute() (*ActionPlan, *http.Response, error) {
	return r.ApiService.AddActionPlanExecute(r)
}

/*
AddActionPlan Submits an Action Plan

Submits an Action Plan and saves upon completion

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAddActionPlanRequest
*/
func (a *ActionsAPIService) AddActionPlan(ctx context.Context) ApiAddActionPlanRequest {
	return ApiAddActionPlanRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ActionPlan
func (a *ActionsAPIService) AddActionPlanExecute(r ApiAddActionPlanRequest) (*ActionPlan, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ActionPlan
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActionsAPIService.AddActionPlan")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/actions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
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

type ApiGetActionPlansRequest struct {
	ctx context.Context
	ApiService ActionsAPI
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

func (r ApiGetActionPlansRequest) Execute() ([]ActionPlan, *http.Response, error) {
	return r.ApiService.GetActionPlansExecute(r)
}

/*
GetActionPlans Gets a list of submitted action (remediation) plans

Retrieves a list of action plans that have been completed

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetActionPlansRequest
*/
func (a *ActionsAPIService) GetActionPlans(ctx context.Context) ApiGetActionPlansRequest {
	return ApiGetActionPlansRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ActionPlan
func (a *ActionsAPIService) GetActionPlansExecute(r ApiGetActionPlansRequest) ([]ActionPlan, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ActionPlan
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActionsAPIService.GetActionPlans")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/actions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.imageTag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "image_tag", r.imageTag, "form", "")
	}
	if r.imageDigest != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "image_digest", r.imageDigest, "form", "")
	}
	if r.createdAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "created_after", r.createdAfter, "form", "")
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
