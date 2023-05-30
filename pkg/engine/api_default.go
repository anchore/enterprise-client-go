/*
Anchore Engine API Server

This is the Anchore Engine API. Provides the primary external API for users of the service.

API version: 0.7.0
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

type DefaultApi interface {

	/*
	GetOauthToken Method for GetOauthToken

	Request a jwt token for subsequent operations, this request is authenticated with normal HTTP auth

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiGetOauthTokenRequest
	*/
	GetOauthToken(ctx _context.Context) ApiGetOauthTokenRequest

	// GetOauthTokenExecute executes the request
	//  @return TokenResponse
	GetOauthTokenExecute(r ApiGetOauthTokenRequest) (TokenResponse, *_nethttp.Response, error)

	/*
	HealthCheck Method for HealthCheck

	Health check, returns 200 and no body if service is running

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiHealthCheckRequest
	*/
	HealthCheck(ctx _context.Context) ApiHealthCheckRequest

	// HealthCheckExecute executes the request
	HealthCheckExecute(r ApiHealthCheckRequest) (*_nethttp.Response, error)

	/*
	ListFileContentSearchResults Return a list of analyzer artifacts of the specified type

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param imageDigest
	 @return ApiListFileContentSearchResultsRequest
	*/
	ListFileContentSearchResults(ctx _context.Context, imageDigest string) ApiListFileContentSearchResultsRequest

	// ListFileContentSearchResultsExecute executes the request
	//  @return []FileContentSearchResult
	ListFileContentSearchResultsExecute(r ApiListFileContentSearchResultsRequest) ([]FileContentSearchResult, *_nethttp.Response, error)

	/*
	ListRetrievedFiles Return a list of analyzer artifacts of the specified type

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param imageDigest
	 @return ApiListRetrievedFilesRequest
	*/
	ListRetrievedFiles(ctx _context.Context, imageDigest string) ApiListRetrievedFilesRequest

	// ListRetrievedFilesExecute executes the request
	//  @return []RetrievedFile
	ListRetrievedFilesExecute(r ApiListRetrievedFilesRequest) ([]RetrievedFile, *_nethttp.Response, error)

	/*
	ListSecretSearchResults Return a list of analyzer artifacts of the specified type

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param imageDigest
	 @return ApiListSecretSearchResultsRequest
	*/
	ListSecretSearchResults(ctx _context.Context, imageDigest string) ApiListSecretSearchResultsRequest

	// ListSecretSearchResultsExecute executes the request
	//  @return []SecretSearchResult
	ListSecretSearchResultsExecute(r ApiListSecretSearchResultsRequest) ([]SecretSearchResult, *_nethttp.Response, error)

	/*
	Ping Method for Ping

	Simple status check

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiPingRequest
	*/
	Ping(ctx _context.Context) ApiPingRequest

	// PingExecute executes the request
	//  @return string
	PingExecute(r ApiPingRequest) (string, *_nethttp.Response, error)

	/*
	RevokeOauthToken Method for RevokeOauthToken

	Revoke a refresh token previously requested from /oauth/token

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiRevokeOauthTokenRequest
	*/
	RevokeOauthToken(ctx _context.Context) ApiRevokeOauthTokenRequest

	// RevokeOauthTokenExecute executes the request
	RevokeOauthTokenExecute(r ApiRevokeOauthTokenRequest) (*_nethttp.Response, error)

	/*
	VersionCheck Method for VersionCheck

	Returns the version object for the service, including db schema version info

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiVersionCheckRequest
	*/
	VersionCheck(ctx _context.Context) ApiVersionCheckRequest

	// VersionCheckExecute executes the request
	//  @return ServiceVersion
	VersionCheckExecute(r ApiVersionCheckRequest) (ServiceVersion, *_nethttp.Response, error)
}

// DefaultApiService DefaultApi service
type DefaultApiService service

type ApiGetOauthTokenRequest struct {
	ctx _context.Context
	ApiService DefaultApi
	grantType *string
	username *string
	password *string
	clientId *string
	refreshToken *string
}

// OAuth Grant type for token
func (r ApiGetOauthTokenRequest) GrantType(grantType string) ApiGetOauthTokenRequest {
	r.grantType = &grantType
	return r
}
// User to assign OAuth token to
func (r ApiGetOauthTokenRequest) Username(username string) ApiGetOauthTokenRequest {
	r.username = &username
	return r
}
// Password for corresponding user
func (r ApiGetOauthTokenRequest) Password(password string) ApiGetOauthTokenRequest {
	r.password = &password
	return r
}
// The type of client used for the OAuth token
func (r ApiGetOauthTokenRequest) ClientId(clientId string) ApiGetOauthTokenRequest {
	r.clientId = &clientId
	return r
}
// The refresh token from a previous password grant request, used to get a new access_token
func (r ApiGetOauthTokenRequest) RefreshToken(refreshToken string) ApiGetOauthTokenRequest {
	r.refreshToken = &refreshToken
	return r
}

func (r ApiGetOauthTokenRequest) Execute() (TokenResponse, *_nethttp.Response, error) {
	return r.ApiService.GetOauthTokenExecute(r)
}

/*
GetOauthToken Method for GetOauthToken

Request a jwt token for subsequent operations, this request is authenticated with normal HTTP auth

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetOauthTokenRequest
*/
func (a *DefaultApiService) GetOauthToken(ctx _context.Context) ApiGetOauthTokenRequest {
	return ApiGetOauthTokenRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TokenResponse
func (a *DefaultApiService) GetOauthTokenExecute(r ApiGetOauthTokenRequest) (TokenResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TokenResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GetOauthToken")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/token"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

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
	if r.grantType != nil {
		localVarFormParams.Add("grant_type", parameterToString(*r.grantType, ""))
	}
	if r.username != nil {
		localVarFormParams.Add("username", parameterToString(*r.username, ""))
	}
	if r.password != nil {
		localVarFormParams.Add("password", parameterToString(*r.password, ""))
	}
	if r.clientId != nil {
		localVarFormParams.Add("client_id", parameterToString(*r.clientId, ""))
	}
	if r.refreshToken != nil {
		localVarFormParams.Add("refresh_token", parameterToString(*r.refreshToken, ""))
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

type ApiHealthCheckRequest struct {
	ctx _context.Context
	ApiService DefaultApi
}


func (r ApiHealthCheckRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.HealthCheckExecute(r)
}

/*
HealthCheck Method for HealthCheck

Health check, returns 200 and no body if service is running

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiHealthCheckRequest
*/
func (a *DefaultApiService) HealthCheck(ctx _context.Context) ApiHealthCheckRequest {
	return ApiHealthCheckRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *DefaultApiService) HealthCheckExecute(r ApiHealthCheckRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.HealthCheck")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/health"

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

type ApiListFileContentSearchResultsRequest struct {
	ctx _context.Context
	ApiService DefaultApi
	imageDigest string
}


func (r ApiListFileContentSearchResultsRequest) Execute() ([]FileContentSearchResult, *_nethttp.Response, error) {
	return r.ApiService.ListFileContentSearchResultsExecute(r)
}

/*
ListFileContentSearchResults Return a list of analyzer artifacts of the specified type

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param imageDigest
 @return ApiListFileContentSearchResultsRequest
*/
func (a *DefaultApiService) ListFileContentSearchResults(ctx _context.Context, imageDigest string) ApiListFileContentSearchResultsRequest {
	return ApiListFileContentSearchResultsRequest{
		ApiService: a,
		ctx: ctx,
		imageDigest: imageDigest,
	}
}

// Execute executes the request
//  @return []FileContentSearchResult
func (a *DefaultApiService) ListFileContentSearchResultsExecute(r ApiListFileContentSearchResultsRequest) ([]FileContentSearchResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []FileContentSearchResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.ListFileContentSearchResults")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/images/{imageDigest}/artifacts/file_content_search"
	localVarPath = strings.Replace(localVarPath, "{"+"imageDigest"+"}", _neturl.PathEscape(parameterToString(r.imageDigest, "")), -1)

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

type ApiListRetrievedFilesRequest struct {
	ctx _context.Context
	ApiService DefaultApi
	imageDigest string
}


func (r ApiListRetrievedFilesRequest) Execute() ([]RetrievedFile, *_nethttp.Response, error) {
	return r.ApiService.ListRetrievedFilesExecute(r)
}

/*
ListRetrievedFiles Return a list of analyzer artifacts of the specified type

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param imageDigest
 @return ApiListRetrievedFilesRequest
*/
func (a *DefaultApiService) ListRetrievedFiles(ctx _context.Context, imageDigest string) ApiListRetrievedFilesRequest {
	return ApiListRetrievedFilesRequest{
		ApiService: a,
		ctx: ctx,
		imageDigest: imageDigest,
	}
}

// Execute executes the request
//  @return []RetrievedFile
func (a *DefaultApiService) ListRetrievedFilesExecute(r ApiListRetrievedFilesRequest) ([]RetrievedFile, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []RetrievedFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.ListRetrievedFiles")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/images/{imageDigest}/artifacts/retrieved_files"
	localVarPath = strings.Replace(localVarPath, "{"+"imageDigest"+"}", _neturl.PathEscape(parameterToString(r.imageDigest, "")), -1)

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

type ApiListSecretSearchResultsRequest struct {
	ctx _context.Context
	ApiService DefaultApi
	imageDigest string
}


func (r ApiListSecretSearchResultsRequest) Execute() ([]SecretSearchResult, *_nethttp.Response, error) {
	return r.ApiService.ListSecretSearchResultsExecute(r)
}

/*
ListSecretSearchResults Return a list of analyzer artifacts of the specified type

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param imageDigest
 @return ApiListSecretSearchResultsRequest
*/
func (a *DefaultApiService) ListSecretSearchResults(ctx _context.Context, imageDigest string) ApiListSecretSearchResultsRequest {
	return ApiListSecretSearchResultsRequest{
		ApiService: a,
		ctx: ctx,
		imageDigest: imageDigest,
	}
}

// Execute executes the request
//  @return []SecretSearchResult
func (a *DefaultApiService) ListSecretSearchResultsExecute(r ApiListSecretSearchResultsRequest) ([]SecretSearchResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []SecretSearchResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.ListSecretSearchResults")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/images/{imageDigest}/artifacts/secret_search"
	localVarPath = strings.Replace(localVarPath, "{"+"imageDigest"+"}", _neturl.PathEscape(parameterToString(r.imageDigest, "")), -1)

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

type ApiPingRequest struct {
	ctx _context.Context
	ApiService DefaultApi
}


func (r ApiPingRequest) Execute() (string, *_nethttp.Response, error) {
	return r.ApiService.PingExecute(r)
}

/*
Ping Method for Ping

Simple status check

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPingRequest
*/
func (a *DefaultApiService) Ping(ctx _context.Context) ApiPingRequest {
	return ApiPingRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return string
func (a *DefaultApiService) PingExecute(r ApiPingRequest) (string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.Ping")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/"

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

type ApiRevokeOauthTokenRequest struct {
	ctx _context.Context
	ApiService DefaultApi
	token *string
	tokenTypeHint *string
}

// The token to be revoked
func (r ApiRevokeOauthTokenRequest) Token(token string) ApiRevokeOauthTokenRequest {
	r.token = &token
	return r
}
// A hint about the type of token to be revoked
func (r ApiRevokeOauthTokenRequest) TokenTypeHint(tokenTypeHint string) ApiRevokeOauthTokenRequest {
	r.tokenTypeHint = &tokenTypeHint
	return r
}

func (r ApiRevokeOauthTokenRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.RevokeOauthTokenExecute(r)
}

/*
RevokeOauthToken Method for RevokeOauthToken

Revoke a refresh token previously requested from /oauth/token

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRevokeOauthTokenRequest
*/
func (a *DefaultApiService) RevokeOauthToken(ctx _context.Context) ApiRevokeOauthTokenRequest {
	return ApiRevokeOauthTokenRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *DefaultApiService) RevokeOauthTokenExecute(r ApiRevokeOauthTokenRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.RevokeOauthToken")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/revoke"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

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
	if r.token != nil {
		localVarFormParams.Add("token", parameterToString(*r.token, ""))
	}
	if r.tokenTypeHint != nil {
		localVarFormParams.Add("token_type_hint", parameterToString(*r.tokenTypeHint, ""))
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v InlineResponse400
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiVersionCheckRequest struct {
	ctx _context.Context
	ApiService DefaultApi
}


func (r ApiVersionCheckRequest) Execute() (ServiceVersion, *_nethttp.Response, error) {
	return r.ApiService.VersionCheckExecute(r)
}

/*
VersionCheck Method for VersionCheck

Returns the version object for the service, including db schema version info

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiVersionCheckRequest
*/
func (a *DefaultApiService) VersionCheck(ctx _context.Context) ApiVersionCheckRequest {
	return ApiVersionCheckRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ServiceVersion
func (a *DefaultApiService) VersionCheckExecute(r ApiVersionCheckRequest) (ServiceVersion, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ServiceVersion
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.VersionCheck")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/version"

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
