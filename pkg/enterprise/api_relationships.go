/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.5.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


type RelationshipsApi interface {

	/*
	AddArtifactRelationship Method for AddArtifactRelationship

	Add a new relationship for this image to another artifact (source or image)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddArtifactRelationshipRequest
	*/
	AddArtifactRelationship(ctx context.Context) ApiAddArtifactRelationshipRequest

	// AddArtifactRelationshipExecute executes the request
	//  @return interface{}
	AddArtifactRelationshipExecute(r ApiAddArtifactRelationshipRequest) (interface{}, *http.Response, error)

	/*
	DeleteArtifactRelationships Method for DeleteArtifactRelationships

	Delete one or more relationships

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiDeleteArtifactRelationshipsRequest
	*/
	DeleteArtifactRelationships(ctx context.Context) ApiDeleteArtifactRelationshipsRequest

	// DeleteArtifactRelationshipsExecute executes the request
	//  @return interface{}
	DeleteArtifactRelationshipsExecute(r ApiDeleteArtifactRelationshipsRequest) (interface{}, *http.Response, error)

	/*
	GetArtifactRelationship Method for GetArtifactRelationship

	Get the relationship between software supply chain artifacts (images, source revisions, etc)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param relationshipId Id of record to retrieve
	@return ApiGetArtifactRelationshipRequest
	*/
	GetArtifactRelationship(ctx context.Context, relationshipId string) ApiGetArtifactRelationshipRequest

	// GetArtifactRelationshipExecute executes the request
	//  @return ArtifactRelationship
	GetArtifactRelationshipExecute(r ApiGetArtifactRelationshipRequest) (*ArtifactRelationship, *http.Response, error)

	/*
	GetRelationshipSbomDiff Method for GetRelationshipSbomDiff

	Return the context-aware diff of the sboms for the relationship

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param relationshipId
	@return ApiGetRelationshipSbomDiffRequest
	*/
	GetRelationshipSbomDiff(ctx context.Context, relationshipId string) ApiGetRelationshipSbomDiffRequest

	// GetRelationshipSbomDiffExecute executes the request
	//  @return RelationshipSbomDiff
	GetRelationshipSbomDiffExecute(r ApiGetRelationshipSbomDiffRequest) (*RelationshipSbomDiff, *http.Response, error)

	/*
	ListArtifactRelationships Method for ListArtifactRelationships

	List the relationships between software supply chain artifacts (images, source revisions, etc)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListArtifactRelationshipsRequest
	*/
	ListArtifactRelationships(ctx context.Context) ApiListArtifactRelationshipsRequest

	// ListArtifactRelationshipsExecute executes the request
	//  @return []ArtifactRelationship
	ListArtifactRelationshipsExecute(r ApiListArtifactRelationshipsRequest) ([]ArtifactRelationship, *http.Response, error)
}

// RelationshipsApiService RelationshipsApi service
type RelationshipsApiService service

type ApiAddArtifactRelationshipRequest struct {
	ctx context.Context
	ApiService RelationshipsApi
	relationship *ArtifactRelationship
}

func (r ApiAddArtifactRelationshipRequest) Relationship(relationship ArtifactRelationship) ApiAddArtifactRelationshipRequest {
	r.relationship = &relationship
	return r
}

func (r ApiAddArtifactRelationshipRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.AddArtifactRelationshipExecute(r)
}

/*
AddArtifactRelationship Method for AddArtifactRelationship

Add a new relationship for this image to another artifact (source or image)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAddArtifactRelationshipRequest
*/
func (a *RelationshipsApiService) AddArtifactRelationship(ctx context.Context) ApiAddArtifactRelationshipRequest {
	return ApiAddArtifactRelationshipRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return interface{}
func (a *RelationshipsApiService) AddArtifactRelationshipExecute(r ApiAddArtifactRelationshipRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RelationshipsApiService.AddArtifactRelationship")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/artifact-relationships"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.relationship == nil {
		return localVarReturnValue, nil, reportError("relationship is required and must be specified")
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
	localVarPostBody = r.relationship
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiDeleteArtifactRelationshipsRequest struct {
	ctx context.Context
	ApiService RelationshipsApi
	relationshipIds *[]string
}

// List of relationship Ids to delete
func (r ApiDeleteArtifactRelationshipsRequest) RelationshipIds(relationshipIds []string) ApiDeleteArtifactRelationshipsRequest {
	r.relationshipIds = &relationshipIds
	return r
}

func (r ApiDeleteArtifactRelationshipsRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.DeleteArtifactRelationshipsExecute(r)
}

/*
DeleteArtifactRelationships Method for DeleteArtifactRelationships

Delete one or more relationships

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeleteArtifactRelationshipsRequest
*/
func (a *RelationshipsApiService) DeleteArtifactRelationships(ctx context.Context) ApiDeleteArtifactRelationshipsRequest {
	return ApiDeleteArtifactRelationshipsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return interface{}
func (a *RelationshipsApiService) DeleteArtifactRelationshipsExecute(r ApiDeleteArtifactRelationshipsRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RelationshipsApiService.DeleteArtifactRelationships")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/artifact-relationships"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.relationshipIds == nil {
		return localVarReturnValue, nil, reportError("relationshipIds is required and must be specified")
	}

	localVarQueryParams.Add("relationship_ids", parameterToString(*r.relationshipIds, "csv"))
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiGetArtifactRelationshipRequest struct {
	ctx context.Context
	ApiService RelationshipsApi
	relationshipId string
}

func (r ApiGetArtifactRelationshipRequest) Execute() (*ArtifactRelationship, *http.Response, error) {
	return r.ApiService.GetArtifactRelationshipExecute(r)
}

/*
GetArtifactRelationship Method for GetArtifactRelationship

Get the relationship between software supply chain artifacts (images, source revisions, etc)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param relationshipId Id of record to retrieve
 @return ApiGetArtifactRelationshipRequest
*/
func (a *RelationshipsApiService) GetArtifactRelationship(ctx context.Context, relationshipId string) ApiGetArtifactRelationshipRequest {
	return ApiGetArtifactRelationshipRequest{
		ApiService: a,
		ctx: ctx,
		relationshipId: relationshipId,
	}
}

// Execute executes the request
//  @return ArtifactRelationship
func (a *RelationshipsApiService) GetArtifactRelationshipExecute(r ApiGetArtifactRelationshipRequest) (*ArtifactRelationship, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ArtifactRelationship
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RelationshipsApiService.GetArtifactRelationship")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/artifact-relationships/{relationship_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"relationship_id"+"}", url.PathEscape(parameterToString(r.relationshipId, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiGetRelationshipSbomDiffRequest struct {
	ctx context.Context
	ApiService RelationshipsApi
	relationshipId string
}

func (r ApiGetRelationshipSbomDiffRequest) Execute() (*RelationshipSbomDiff, *http.Response, error) {
	return r.ApiService.GetRelationshipSbomDiffExecute(r)
}

/*
GetRelationshipSbomDiff Method for GetRelationshipSbomDiff

Return the context-aware diff of the sboms for the relationship

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param relationshipId
 @return ApiGetRelationshipSbomDiffRequest
*/
func (a *RelationshipsApiService) GetRelationshipSbomDiff(ctx context.Context, relationshipId string) ApiGetRelationshipSbomDiffRequest {
	return ApiGetRelationshipSbomDiffRequest{
		ApiService: a,
		ctx: ctx,
		relationshipId: relationshipId,
	}
}

// Execute executes the request
//  @return RelationshipSbomDiff
func (a *RelationshipsApiService) GetRelationshipSbomDiffExecute(r ApiGetRelationshipSbomDiffRequest) (*RelationshipSbomDiff, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RelationshipSbomDiff
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RelationshipsApiService.GetRelationshipSbomDiff")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/artifact-relationships/{relationship_id}/diffs/sbom"
	localVarPath = strings.Replace(localVarPath, "{"+"relationship_id"+"}", url.PathEscape(parameterToString(r.relationshipId, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiListArtifactRelationshipsRequest struct {
	ctx context.Context
	ApiService RelationshipsApi
	artifactType *string
	artifactId *string
}

// Filter for artifact type as either source or target
func (r ApiListArtifactRelationshipsRequest) ArtifactType(artifactType string) ApiListArtifactRelationshipsRequest {
	r.artifactType = &artifactType
	return r
}

// Filter for artifact id as either source or target
func (r ApiListArtifactRelationshipsRequest) ArtifactId(artifactId string) ApiListArtifactRelationshipsRequest {
	r.artifactId = &artifactId
	return r
}

func (r ApiListArtifactRelationshipsRequest) Execute() ([]ArtifactRelationship, *http.Response, error) {
	return r.ApiService.ListArtifactRelationshipsExecute(r)
}

/*
ListArtifactRelationships Method for ListArtifactRelationships

List the relationships between software supply chain artifacts (images, source revisions, etc)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListArtifactRelationshipsRequest
*/
func (a *RelationshipsApiService) ListArtifactRelationships(ctx context.Context) ApiListArtifactRelationshipsRequest {
	return ApiListArtifactRelationshipsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ArtifactRelationship
func (a *RelationshipsApiService) ListArtifactRelationshipsExecute(r ApiListArtifactRelationshipsRequest) ([]ArtifactRelationship, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ArtifactRelationship
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RelationshipsApiService.ListArtifactRelationships")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/artifact-relationships"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.artifactType != nil {
		localVarQueryParams.Add("artifact_type", parameterToString(*r.artifactType, ""))
	}
	if r.artifactId != nil {
		localVarQueryParams.Add("artifact_id", parameterToString(*r.artifactId, ""))
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
