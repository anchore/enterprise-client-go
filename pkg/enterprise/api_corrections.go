/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 0.1.0
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

type CorrectionsApi interface {

	/*
	AddCorrection Create a correction record

	Add a correction record that will be used to fix false positive matches

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiAddCorrectionRequest
	*/
	AddCorrection(ctx _context.Context) ApiAddCorrectionRequest

	// AddCorrectionExecute executes the request
	//  @return Correction
	AddCorrectionExecute(r ApiAddCorrectionRequest) (Correction, *_nethttp.Response, error)

	/*
	DeleteCorrectionByUuid Delete a correction by UUID

	Delete a single correction, looked up via it's uuid

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param uuid
	 @return ApiDeleteCorrectionByUuidRequest
	*/
	DeleteCorrectionByUuid(ctx _context.Context, uuid string) ApiDeleteCorrectionByUuidRequest

	// DeleteCorrectionByUuidExecute executes the request
	DeleteCorrectionByUuidExecute(r ApiDeleteCorrectionByUuidRequest) (*_nethttp.Response, error)

	/*
	GetCorrectionByUuid Retrieve a correction by UUID

	Returns a single correction, looked up via it's uuid

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param uuid
	 @return ApiGetCorrectionByUuidRequest
	*/
	GetCorrectionByUuid(ctx _context.Context, uuid string) ApiGetCorrectionByUuidRequest

	// GetCorrectionByUuidExecute executes the request
	//  @return Correction
	GetCorrectionByUuidExecute(r ApiGetCorrectionByUuidRequest) (Correction, *_nethttp.Response, error)

	/*
	GetCorrections Retrieve a list of corrections

	Returns a list of corrections

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiGetCorrectionsRequest
	*/
	GetCorrections(ctx _context.Context) ApiGetCorrectionsRequest

	// GetCorrectionsExecute executes the request
	//  @return []Correction
	GetCorrectionsExecute(r ApiGetCorrectionsRequest) ([]Correction, *_nethttp.Response, error)

	/*
	UpdateCorrectionByUuid Update a correction by UUID

	Updates a single correction, looked up via it's uuid

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param uuid
	 @return ApiUpdateCorrectionByUuidRequest
	*/
	UpdateCorrectionByUuid(ctx _context.Context, uuid string) ApiUpdateCorrectionByUuidRequest

	// UpdateCorrectionByUuidExecute executes the request
	//  @return Correction
	UpdateCorrectionByUuidExecute(r ApiUpdateCorrectionByUuidRequest) (Correction, *_nethttp.Response, error)
}

// CorrectionsApiService CorrectionsApi service
type CorrectionsApiService service

type ApiAddCorrectionRequest struct {
	ctx _context.Context
	ApiService CorrectionsApi
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

func (r ApiAddCorrectionRequest) Execute() (Correction, *_nethttp.Response, error) {
	return r.ApiService.AddCorrectionExecute(r)
}

/*
AddCorrection Create a correction record

Add a correction record that will be used to fix false positive matches

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAddCorrectionRequest
*/
func (a *CorrectionsApiService) AddCorrection(ctx _context.Context) ApiAddCorrectionRequest {
	return ApiAddCorrectionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Correction
func (a *CorrectionsApiService) AddCorrectionExecute(r ApiAddCorrectionRequest) (Correction, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Correction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CorrectionsApiService.AddCorrection")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/corrections"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
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
		localVarHeaderParams["x-anchore-account"] = parameterToString(*r.xAnchoreAccount, "")
	}
	// body params
	localVarPostBody = r.correction
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

type ApiDeleteCorrectionByUuidRequest struct {
	ctx _context.Context
	ApiService CorrectionsApi
	uuid string
	xAnchoreAccount *string
}

// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiDeleteCorrectionByUuidRequest) XAnchoreAccount(xAnchoreAccount string) ApiDeleteCorrectionByUuidRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

func (r ApiDeleteCorrectionByUuidRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteCorrectionByUuidExecute(r)
}

/*
DeleteCorrectionByUuid Delete a correction by UUID

Delete a single correction, looked up via it's uuid

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param uuid
 @return ApiDeleteCorrectionByUuidRequest
*/
func (a *CorrectionsApiService) DeleteCorrectionByUuid(ctx _context.Context, uuid string) ApiDeleteCorrectionByUuidRequest {
	return ApiDeleteCorrectionByUuidRequest{
		ApiService: a,
		ctx: ctx,
		uuid: uuid,
	}
}

// Execute executes the request
func (a *CorrectionsApiService) DeleteCorrectionByUuidExecute(r ApiDeleteCorrectionByUuidRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CorrectionsApiService.DeleteCorrectionByUuid")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/corrections/{uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"uuid"+"}", _neturl.PathEscape(parameterToString(r.uuid, "")), -1)

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

type ApiGetCorrectionByUuidRequest struct {
	ctx _context.Context
	ApiService CorrectionsApi
	uuid string
	xAnchoreAccount *string
}

// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiGetCorrectionByUuidRequest) XAnchoreAccount(xAnchoreAccount string) ApiGetCorrectionByUuidRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

func (r ApiGetCorrectionByUuidRequest) Execute() (Correction, *_nethttp.Response, error) {
	return r.ApiService.GetCorrectionByUuidExecute(r)
}

/*
GetCorrectionByUuid Retrieve a correction by UUID

Returns a single correction, looked up via it's uuid

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param uuid
 @return ApiGetCorrectionByUuidRequest
*/
func (a *CorrectionsApiService) GetCorrectionByUuid(ctx _context.Context, uuid string) ApiGetCorrectionByUuidRequest {
	return ApiGetCorrectionByUuidRequest{
		ApiService: a,
		ctx: ctx,
		uuid: uuid,
	}
}

// Execute executes the request
//  @return Correction
func (a *CorrectionsApiService) GetCorrectionByUuidExecute(r ApiGetCorrectionByUuidRequest) (Correction, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Correction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CorrectionsApiService.GetCorrectionByUuid")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/corrections/{uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"uuid"+"}", _neturl.PathEscape(parameterToString(r.uuid, "")), -1)

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

type ApiGetCorrectionsRequest struct {
	ctx _context.Context
	ApiService CorrectionsApi
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

func (r ApiGetCorrectionsRequest) Execute() ([]Correction, *_nethttp.Response, error) {
	return r.ApiService.GetCorrectionsExecute(r)
}

/*
GetCorrections Retrieve a list of corrections

Returns a list of corrections

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCorrectionsRequest
*/
func (a *CorrectionsApiService) GetCorrections(ctx _context.Context) ApiGetCorrectionsRequest {
	return ApiGetCorrectionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Correction
func (a *CorrectionsApiService) GetCorrectionsExecute(r ApiGetCorrectionsRequest) ([]Correction, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []Correction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CorrectionsApiService.GetCorrections")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/corrections"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.correctionType != nil {
		localVarQueryParams.Add("correction_type", parameterToString(*r.correctionType, ""))
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

type ApiUpdateCorrectionByUuidRequest struct {
	ctx _context.Context
	ApiService CorrectionsApi
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

func (r ApiUpdateCorrectionByUuidRequest) Execute() (Correction, *_nethttp.Response, error) {
	return r.ApiService.UpdateCorrectionByUuidExecute(r)
}

/*
UpdateCorrectionByUuid Update a correction by UUID

Updates a single correction, looked up via it's uuid

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param uuid
 @return ApiUpdateCorrectionByUuidRequest
*/
func (a *CorrectionsApiService) UpdateCorrectionByUuid(ctx _context.Context, uuid string) ApiUpdateCorrectionByUuidRequest {
	return ApiUpdateCorrectionByUuidRequest{
		ApiService: a,
		ctx: ctx,
		uuid: uuid,
	}
}

// Execute executes the request
//  @return Correction
func (a *CorrectionsApiService) UpdateCorrectionByUuidExecute(r ApiUpdateCorrectionByUuidRequest) (Correction, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Correction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CorrectionsApiService.UpdateCorrectionByUuid")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/corrections/{uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"uuid"+"}", _neturl.PathEscape(parameterToString(r.uuid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
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
		localVarHeaderParams["x-anchore-account"] = parameterToString(*r.xAnchoreAccount, "")
	}
	// body params
	localVarPostBody = r.correction
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
