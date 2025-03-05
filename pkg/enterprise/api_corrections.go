/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.10.2
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


type CorrectionsAPI interface {

	/*
	AddCorrection Create a correction record

	Add a correction record that will be used to fix false positive matches

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddCorrectionRequest
	*/
	AddCorrection(ctx context.Context) ApiAddCorrectionRequest

	// AddCorrectionExecute executes the request
	//  @return Correction
	AddCorrectionExecute(r ApiAddCorrectionRequest) (*Correction, *http.Response, error)

	/*
	DeleteCorrectionByUuid Delete a correction by UUID

	Delete a single correction, looked up via it's uuid

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param uuid
	@return ApiDeleteCorrectionByUuidRequest
	*/
	DeleteCorrectionByUuid(ctx context.Context, uuid string) ApiDeleteCorrectionByUuidRequest

	// DeleteCorrectionByUuidExecute executes the request
	DeleteCorrectionByUuidExecute(r ApiDeleteCorrectionByUuidRequest) (*http.Response, error)

	/*
	GetCorrectionByUuid Retrieve a correction by UUID

	Returns a single correction, looked up via it's uuid

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param uuid
	@return ApiGetCorrectionByUuidRequest
	*/
	GetCorrectionByUuid(ctx context.Context, uuid string) ApiGetCorrectionByUuidRequest

	// GetCorrectionByUuidExecute executes the request
	//  @return Correction
	GetCorrectionByUuidExecute(r ApiGetCorrectionByUuidRequest) (*Correction, *http.Response, error)

	/*
	GetCorrections Retrieve a list of corrections

	Returns a list of corrections

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetCorrectionsRequest
	*/
	GetCorrections(ctx context.Context) ApiGetCorrectionsRequest

	// GetCorrectionsExecute executes the request
	//  @return []Correction
	GetCorrectionsExecute(r ApiGetCorrectionsRequest) ([]Correction, *http.Response, error)

	/*
	UpdateCorrectionByUuid Update a correction by UUID

	Updates a single correction, looked up via it's uuid

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param uuid
	@return ApiUpdateCorrectionByUuidRequest
	*/
	UpdateCorrectionByUuid(ctx context.Context, uuid string) ApiUpdateCorrectionByUuidRequest

	// UpdateCorrectionByUuidExecute executes the request
	//  @return Correction
	UpdateCorrectionByUuidExecute(r ApiUpdateCorrectionByUuidRequest) (*Correction, *http.Response, error)
}

// CorrectionsAPIService CorrectionsAPI service
type CorrectionsAPIService service

type ApiAddCorrectionRequest struct {
	ctx context.Context
	ApiService CorrectionsAPI
	correction *Correction
	xAnchoreAccount *string
}

func (r ApiAddCorrectionRequest) Correction(correction Correction) ApiAddCorrectionRequest {
	r.correction = &correction
	return r
}

// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiAddCorrectionRequest) XAnchoreAccount(xAnchoreAccount string) ApiAddCorrectionRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

func (r ApiAddCorrectionRequest) Execute() (*Correction, *http.Response, error) {
	return r.ApiService.AddCorrectionExecute(r)
}

/*
AddCorrection Create a correction record

Add a correction record that will be used to fix false positive matches

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAddCorrectionRequest
*/
func (a *CorrectionsAPIService) AddCorrection(ctx context.Context) ApiAddCorrectionRequest {
	return ApiAddCorrectionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Correction
func (a *CorrectionsAPIService) AddCorrectionExecute(r ApiAddCorrectionRequest) (*Correction, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Correction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CorrectionsAPIService.AddCorrection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/corrections"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.correction == nil {
		return localVarReturnValue, nil, reportError("correction is required and must be specified")
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
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-anchore-account", r.xAnchoreAccount, "simple", "")
	}
	// body params
	localVarPostBody = r.correction
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

type ApiDeleteCorrectionByUuidRequest struct {
	ctx context.Context
	ApiService CorrectionsAPI
	uuid string
	xAnchoreAccount *string
}

// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiDeleteCorrectionByUuidRequest) XAnchoreAccount(xAnchoreAccount string) ApiDeleteCorrectionByUuidRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

func (r ApiDeleteCorrectionByUuidRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteCorrectionByUuidExecute(r)
}

/*
DeleteCorrectionByUuid Delete a correction by UUID

Delete a single correction, looked up via it's uuid

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param uuid
 @return ApiDeleteCorrectionByUuidRequest
*/
func (a *CorrectionsAPIService) DeleteCorrectionByUuid(ctx context.Context, uuid string) ApiDeleteCorrectionByUuidRequest {
	return ApiDeleteCorrectionByUuidRequest{
		ApiService: a,
		ctx: ctx,
		uuid: uuid,
	}
}

// Execute executes the request
func (a *CorrectionsAPIService) DeleteCorrectionByUuidExecute(r ApiDeleteCorrectionByUuidRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CorrectionsAPIService.DeleteCorrectionByUuid")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/corrections/{uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"uuid"+"}", url.PathEscape(parameterValueToString(r.uuid, "uuid")), -1)

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

type ApiGetCorrectionByUuidRequest struct {
	ctx context.Context
	ApiService CorrectionsAPI
	uuid string
	xAnchoreAccount *string
}

// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiGetCorrectionByUuidRequest) XAnchoreAccount(xAnchoreAccount string) ApiGetCorrectionByUuidRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

func (r ApiGetCorrectionByUuidRequest) Execute() (*Correction, *http.Response, error) {
	return r.ApiService.GetCorrectionByUuidExecute(r)
}

/*
GetCorrectionByUuid Retrieve a correction by UUID

Returns a single correction, looked up via it's uuid

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param uuid
 @return ApiGetCorrectionByUuidRequest
*/
func (a *CorrectionsAPIService) GetCorrectionByUuid(ctx context.Context, uuid string) ApiGetCorrectionByUuidRequest {
	return ApiGetCorrectionByUuidRequest{
		ApiService: a,
		ctx: ctx,
		uuid: uuid,
	}
}

// Execute executes the request
//  @return Correction
func (a *CorrectionsAPIService) GetCorrectionByUuidExecute(r ApiGetCorrectionByUuidRequest) (*Correction, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Correction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CorrectionsAPIService.GetCorrectionByUuid")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/corrections/{uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"uuid"+"}", url.PathEscape(parameterValueToString(r.uuid, "uuid")), -1)

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

type ApiGetCorrectionsRequest struct {
	ctx context.Context
	ApiService CorrectionsAPI
	correctionType *string
	xAnchoreAccount *string
}

func (r ApiGetCorrectionsRequest) CorrectionType(correctionType string) ApiGetCorrectionsRequest {
	r.correctionType = &correctionType
	return r
}

// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiGetCorrectionsRequest) XAnchoreAccount(xAnchoreAccount string) ApiGetCorrectionsRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

func (r ApiGetCorrectionsRequest) Execute() ([]Correction, *http.Response, error) {
	return r.ApiService.GetCorrectionsExecute(r)
}

/*
GetCorrections Retrieve a list of corrections

Returns a list of corrections

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCorrectionsRequest
*/
func (a *CorrectionsAPIService) GetCorrections(ctx context.Context) ApiGetCorrectionsRequest {
	return ApiGetCorrectionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Correction
func (a *CorrectionsAPIService) GetCorrectionsExecute(r ApiGetCorrectionsRequest) ([]Correction, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Correction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CorrectionsAPIService.GetCorrections")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/corrections"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.correctionType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "correction_type", r.correctionType, "form", "")
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

type ApiUpdateCorrectionByUuidRequest struct {
	ctx context.Context
	ApiService CorrectionsAPI
	uuid string
	correction *Correction
	xAnchoreAccount *string
}

func (r ApiUpdateCorrectionByUuidRequest) Correction(correction Correction) ApiUpdateCorrectionByUuidRequest {
	r.correction = &correction
	return r
}

// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiUpdateCorrectionByUuidRequest) XAnchoreAccount(xAnchoreAccount string) ApiUpdateCorrectionByUuidRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

func (r ApiUpdateCorrectionByUuidRequest) Execute() (*Correction, *http.Response, error) {
	return r.ApiService.UpdateCorrectionByUuidExecute(r)
}

/*
UpdateCorrectionByUuid Update a correction by UUID

Updates a single correction, looked up via it's uuid

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param uuid
 @return ApiUpdateCorrectionByUuidRequest
*/
func (a *CorrectionsAPIService) UpdateCorrectionByUuid(ctx context.Context, uuid string) ApiUpdateCorrectionByUuidRequest {
	return ApiUpdateCorrectionByUuidRequest{
		ApiService: a,
		ctx: ctx,
		uuid: uuid,
	}
}

// Execute executes the request
//  @return Correction
func (a *CorrectionsAPIService) UpdateCorrectionByUuidExecute(r ApiUpdateCorrectionByUuidRequest) (*Correction, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Correction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CorrectionsAPIService.UpdateCorrectionByUuid")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/corrections/{uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"uuid"+"}", url.PathEscape(parameterValueToString(r.uuid, "uuid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.correction == nil {
		return localVarReturnValue, nil, reportError("correction is required and must be specified")
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
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-anchore-account", r.xAnchoreAccount, "simple", "")
	}
	// body params
	localVarPostBody = r.correction
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
