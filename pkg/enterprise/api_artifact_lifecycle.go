/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.10.1
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


type ArtifactLifecycleAPI interface {

	/*
	CreateArtifactLifecyclePolicy Create new artifact lifecycle policy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateArtifactLifecyclePolicyRequest
	*/
	CreateArtifactLifecyclePolicy(ctx context.Context) ApiCreateArtifactLifecyclePolicyRequest

	// CreateArtifactLifecyclePolicyExecute executes the request
	//  @return ArtifactLifecyclePolicyResponse
	CreateArtifactLifecyclePolicyExecute(r ApiCreateArtifactLifecyclePolicyRequest) (*ArtifactLifecyclePolicyResponse, *http.Response, error)

	/*
	DeleteArtifactLifecyclePolicy Delete lifecycle policy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param policyUuid
	@return ApiDeleteArtifactLifecyclePolicyRequest
	*/
	DeleteArtifactLifecyclePolicy(ctx context.Context, policyUuid string) ApiDeleteArtifactLifecyclePolicyRequest

	// DeleteArtifactLifecyclePolicyExecute executes the request
	DeleteArtifactLifecyclePolicyExecute(r ApiDeleteArtifactLifecyclePolicyRequest) (*http.Response, error)

	/*
	GetArtifactLifecyclePolicy Get single artifact lifecycle policy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param policyUuid
	@return ApiGetArtifactLifecyclePolicyRequest
	*/
	GetArtifactLifecyclePolicy(ctx context.Context, policyUuid string) ApiGetArtifactLifecyclePolicyRequest

	// GetArtifactLifecyclePolicyExecute executes the request
	//  @return ArtifactLifecyclePolicyResponse
	GetArtifactLifecyclePolicyExecute(r ApiGetArtifactLifecyclePolicyRequest) (*ArtifactLifecyclePolicyResponse, *http.Response, error)

	/*
	GetArtifactLifecyclePolicyByVersion Get single artifact lifecycle policy by its version

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param policyUuid
	@return ApiGetArtifactLifecyclePolicyByVersionRequest
	*/
	GetArtifactLifecyclePolicyByVersion(ctx context.Context, policyUuid string) ApiGetArtifactLifecyclePolicyByVersionRequest

	// GetArtifactLifecyclePolicyByVersionExecute executes the request
	//  @return ArtifactLifecyclePolicyResponse
	GetArtifactLifecyclePolicyByVersionExecute(r ApiGetArtifactLifecyclePolicyByVersionRequest) (*ArtifactLifecyclePolicyResponse, *http.Response, error)

	/*
	ListArtifactLifecyclePolicies List all artifact lifecycle policies

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListArtifactLifecyclePoliciesRequest
	*/
	ListArtifactLifecyclePolicies(ctx context.Context) ApiListArtifactLifecyclePoliciesRequest

	// ListArtifactLifecyclePoliciesExecute executes the request
	//  @return ArtifactLifecyclePolicyList
	ListArtifactLifecyclePoliciesExecute(r ApiListArtifactLifecyclePoliciesRequest) (*ArtifactLifecyclePolicyList, *http.Response, error)

	/*
	UpdateArtifactLifecyclePolicy Update a single artifact lifecycle policy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param policyUuid
	@return ApiUpdateArtifactLifecyclePolicyRequest
	*/
	UpdateArtifactLifecyclePolicy(ctx context.Context, policyUuid string) ApiUpdateArtifactLifecyclePolicyRequest

	// UpdateArtifactLifecyclePolicyExecute executes the request
	//  @return ArtifactLifecyclePolicyResponse
	UpdateArtifactLifecyclePolicyExecute(r ApiUpdateArtifactLifecyclePolicyRequest) (*ArtifactLifecyclePolicyResponse, *http.Response, error)
}

// ArtifactLifecycleAPIService ArtifactLifecycleAPI service
type ArtifactLifecycleAPIService service

type ApiCreateArtifactLifecyclePolicyRequest struct {
	ctx context.Context
	ApiService ArtifactLifecycleAPI
	artifactLifecyclePolicy *ArtifactLifecyclePolicy
}

func (r ApiCreateArtifactLifecyclePolicyRequest) ArtifactLifecyclePolicy(artifactLifecyclePolicy ArtifactLifecyclePolicy) ApiCreateArtifactLifecyclePolicyRequest {
	r.artifactLifecyclePolicy = &artifactLifecyclePolicy
	return r
}

func (r ApiCreateArtifactLifecyclePolicyRequest) Execute() (*ArtifactLifecyclePolicyResponse, *http.Response, error) {
	return r.ApiService.CreateArtifactLifecyclePolicyExecute(r)
}

/*
CreateArtifactLifecyclePolicy Create new artifact lifecycle policy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateArtifactLifecyclePolicyRequest
*/
func (a *ArtifactLifecycleAPIService) CreateArtifactLifecyclePolicy(ctx context.Context) ApiCreateArtifactLifecyclePolicyRequest {
	return ApiCreateArtifactLifecyclePolicyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ArtifactLifecyclePolicyResponse
func (a *ArtifactLifecycleAPIService) CreateArtifactLifecyclePolicyExecute(r ApiCreateArtifactLifecyclePolicyRequest) (*ArtifactLifecyclePolicyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ArtifactLifecyclePolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArtifactLifecycleAPIService.CreateArtifactLifecyclePolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/artifact-lifecycle-policies"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.artifactLifecyclePolicy
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

type ApiDeleteArtifactLifecyclePolicyRequest struct {
	ctx context.Context
	ApiService ArtifactLifecycleAPI
	policyUuid string
}

func (r ApiDeleteArtifactLifecyclePolicyRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteArtifactLifecyclePolicyExecute(r)
}

/*
DeleteArtifactLifecyclePolicy Delete lifecycle policy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param policyUuid
 @return ApiDeleteArtifactLifecyclePolicyRequest
*/
func (a *ArtifactLifecycleAPIService) DeleteArtifactLifecyclePolicy(ctx context.Context, policyUuid string) ApiDeleteArtifactLifecyclePolicyRequest {
	return ApiDeleteArtifactLifecyclePolicyRequest{
		ApiService: a,
		ctx: ctx,
		policyUuid: policyUuid,
	}
}

// Execute executes the request
func (a *ArtifactLifecycleAPIService) DeleteArtifactLifecyclePolicyExecute(r ApiDeleteArtifactLifecyclePolicyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArtifactLifecycleAPIService.DeleteArtifactLifecyclePolicy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/artifact-lifecycle-policies/{policy_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"policy_uuid"+"}", url.PathEscape(parameterValueToString(r.policyUuid, "policyUuid")), -1)

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

type ApiGetArtifactLifecyclePolicyRequest struct {
	ctx context.Context
	ApiService ArtifactLifecycleAPI
	policyUuid string
}

func (r ApiGetArtifactLifecyclePolicyRequest) Execute() (*ArtifactLifecyclePolicyResponse, *http.Response, error) {
	return r.ApiService.GetArtifactLifecyclePolicyExecute(r)
}

/*
GetArtifactLifecyclePolicy Get single artifact lifecycle policy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param policyUuid
 @return ApiGetArtifactLifecyclePolicyRequest
*/
func (a *ArtifactLifecycleAPIService) GetArtifactLifecyclePolicy(ctx context.Context, policyUuid string) ApiGetArtifactLifecyclePolicyRequest {
	return ApiGetArtifactLifecyclePolicyRequest{
		ApiService: a,
		ctx: ctx,
		policyUuid: policyUuid,
	}
}

// Execute executes the request
//  @return ArtifactLifecyclePolicyResponse
func (a *ArtifactLifecycleAPIService) GetArtifactLifecyclePolicyExecute(r ApiGetArtifactLifecyclePolicyRequest) (*ArtifactLifecyclePolicyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ArtifactLifecyclePolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArtifactLifecycleAPIService.GetArtifactLifecyclePolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/artifact-lifecycle-policies/{policy_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"policy_uuid"+"}", url.PathEscape(parameterValueToString(r.policyUuid, "policyUuid")), -1)

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

type ApiGetArtifactLifecyclePolicyByVersionRequest struct {
	ctx context.Context
	ApiService ArtifactLifecycleAPI
	policyUuid string
	version *int32
	latest *bool
}

// Request a specific version number
func (r ApiGetArtifactLifecyclePolicyByVersionRequest) Version(version int32) ApiGetArtifactLifecyclePolicyByVersionRequest {
	r.version = &version
	return r
}

// Request the latest version
func (r ApiGetArtifactLifecyclePolicyByVersionRequest) Latest(latest bool) ApiGetArtifactLifecyclePolicyByVersionRequest {
	r.latest = &latest
	return r
}

func (r ApiGetArtifactLifecyclePolicyByVersionRequest) Execute() (*ArtifactLifecyclePolicyResponse, *http.Response, error) {
	return r.ApiService.GetArtifactLifecyclePolicyByVersionExecute(r)
}

/*
GetArtifactLifecyclePolicyByVersion Get single artifact lifecycle policy by its version

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param policyUuid
 @return ApiGetArtifactLifecyclePolicyByVersionRequest
*/
func (a *ArtifactLifecycleAPIService) GetArtifactLifecyclePolicyByVersion(ctx context.Context, policyUuid string) ApiGetArtifactLifecyclePolicyByVersionRequest {
	return ApiGetArtifactLifecyclePolicyByVersionRequest{
		ApiService: a,
		ctx: ctx,
		policyUuid: policyUuid,
	}
}

// Execute executes the request
//  @return ArtifactLifecyclePolicyResponse
func (a *ArtifactLifecycleAPIService) GetArtifactLifecyclePolicyByVersionExecute(r ApiGetArtifactLifecyclePolicyByVersionRequest) (*ArtifactLifecyclePolicyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ArtifactLifecyclePolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArtifactLifecycleAPIService.GetArtifactLifecyclePolicyByVersion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/artifact-lifecycle-policies/{policy_uuid}/versions"
	localVarPath = strings.Replace(localVarPath, "{"+"policy_uuid"+"}", url.PathEscape(parameterValueToString(r.policyUuid, "policyUuid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.version != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version", r.version, "form", "")
	}
	if r.latest != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "latest", r.latest, "form", "")
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

type ApiListArtifactLifecyclePoliciesRequest struct {
	ctx context.Context
	ApiService ArtifactLifecycleAPI
}

func (r ApiListArtifactLifecyclePoliciesRequest) Execute() (*ArtifactLifecyclePolicyList, *http.Response, error) {
	return r.ApiService.ListArtifactLifecyclePoliciesExecute(r)
}

/*
ListArtifactLifecyclePolicies List all artifact lifecycle policies

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListArtifactLifecyclePoliciesRequest
*/
func (a *ArtifactLifecycleAPIService) ListArtifactLifecyclePolicies(ctx context.Context) ApiListArtifactLifecyclePoliciesRequest {
	return ApiListArtifactLifecyclePoliciesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ArtifactLifecyclePolicyList
func (a *ArtifactLifecycleAPIService) ListArtifactLifecyclePoliciesExecute(r ApiListArtifactLifecyclePoliciesRequest) (*ArtifactLifecyclePolicyList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ArtifactLifecyclePolicyList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArtifactLifecycleAPIService.ListArtifactLifecyclePolicies")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/artifact-lifecycle-policies"

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

type ApiUpdateArtifactLifecyclePolicyRequest struct {
	ctx context.Context
	ApiService ArtifactLifecycleAPI
	policyUuid string
	artifactLifecyclePolicy *ArtifactLifecyclePolicy
}

func (r ApiUpdateArtifactLifecyclePolicyRequest) ArtifactLifecyclePolicy(artifactLifecyclePolicy ArtifactLifecyclePolicy) ApiUpdateArtifactLifecyclePolicyRequest {
	r.artifactLifecyclePolicy = &artifactLifecyclePolicy
	return r
}

func (r ApiUpdateArtifactLifecyclePolicyRequest) Execute() (*ArtifactLifecyclePolicyResponse, *http.Response, error) {
	return r.ApiService.UpdateArtifactLifecyclePolicyExecute(r)
}

/*
UpdateArtifactLifecyclePolicy Update a single artifact lifecycle policy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param policyUuid
 @return ApiUpdateArtifactLifecyclePolicyRequest
*/
func (a *ArtifactLifecycleAPIService) UpdateArtifactLifecyclePolicy(ctx context.Context, policyUuid string) ApiUpdateArtifactLifecyclePolicyRequest {
	return ApiUpdateArtifactLifecyclePolicyRequest{
		ApiService: a,
		ctx: ctx,
		policyUuid: policyUuid,
	}
}

// Execute executes the request
//  @return ArtifactLifecyclePolicyResponse
func (a *ArtifactLifecycleAPIService) UpdateArtifactLifecyclePolicyExecute(r ApiUpdateArtifactLifecyclePolicyRequest) (*ArtifactLifecyclePolicyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ArtifactLifecyclePolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArtifactLifecycleAPIService.UpdateArtifactLifecyclePolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/artifact-lifecycle-policies/{policy_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"policy_uuid"+"}", url.PathEscape(parameterValueToString(r.policyUuid, "policyUuid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.artifactLifecyclePolicy
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
