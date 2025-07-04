/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.12.0
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


type ArchivesAPI interface {

	/*
	ArchiveImageAnalysis Method for ArchiveImageAnalysis

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiArchiveImageAnalysisRequest
	*/
	ArchiveImageAnalysis(ctx context.Context) ApiArchiveImageAnalysisRequest

	// ArchiveImageAnalysisExecute executes the request
	//  @return []AnalysisArchiveAddResult
	ArchiveImageAnalysisExecute(r ApiArchiveImageAnalysisRequest) ([]AnalysisArchiveAddResult, *http.Response, error)

	/*
	CreateAnalysisArchiveRule Method for CreateAnalysisArchiveRule

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateAnalysisArchiveRuleRequest
	*/
	CreateAnalysisArchiveRule(ctx context.Context) ApiCreateAnalysisArchiveRuleRequest

	// CreateAnalysisArchiveRuleExecute executes the request
	//  @return AnalysisArchiveTransitionRule
	CreateAnalysisArchiveRuleExecute(r ApiCreateAnalysisArchiveRuleRequest) (*AnalysisArchiveTransitionRule, *http.Response, error)

	/*
	DeleteAnalysisArchiveRule Method for DeleteAnalysisArchiveRule

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ruleId
	@return ApiDeleteAnalysisArchiveRuleRequest
	*/
	DeleteAnalysisArchiveRule(ctx context.Context, ruleId string) ApiDeleteAnalysisArchiveRuleRequest

	// DeleteAnalysisArchiveRuleExecute executes the request
	DeleteAnalysisArchiveRuleExecute(r ApiDeleteAnalysisArchiveRuleRequest) (*http.Response, error)

	/*
	DeleteArchivedAnalysis Method for DeleteArchivedAnalysis

	Performs a synchronous archive deletion

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param imageDigest
	@return ApiDeleteArchivedAnalysisRequest
	*/
	DeleteArchivedAnalysis(ctx context.Context, imageDigest string) ApiDeleteArchivedAnalysisRequest

	// DeleteArchivedAnalysisExecute executes the request
	DeleteArchivedAnalysisExecute(r ApiDeleteArchivedAnalysisRequest) (*http.Response, error)

	/*
	GetAnalysisArchiveRule Method for GetAnalysisArchiveRule

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ruleId
	@return ApiGetAnalysisArchiveRuleRequest
	*/
	GetAnalysisArchiveRule(ctx context.Context, ruleId string) ApiGetAnalysisArchiveRuleRequest

	// GetAnalysisArchiveRuleExecute executes the request
	//  @return AnalysisArchiveTransitionRule
	GetAnalysisArchiveRuleExecute(r ApiGetAnalysisArchiveRuleRequest) (*AnalysisArchiveTransitionRule, *http.Response, error)

	/*
	GetArchivedAnalysis Method for GetArchivedAnalysis

	Returns the archive metadata record identifying the image and tags for the analysis in the archive.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param imageDigest The image digest to identify the image analysis
	@return ApiGetArchivedAnalysisRequest
	*/
	GetArchivedAnalysis(ctx context.Context, imageDigest string) ApiGetArchivedAnalysisRequest

	// GetArchivedAnalysisExecute executes the request
	//  @return ArchivedAnalysis
	GetArchivedAnalysisExecute(r ApiGetArchivedAnalysisRequest) (*ArchivedAnalysis, *http.Response, error)

	/*
	ListAnalysisArchive Method for ListAnalysisArchive

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListAnalysisArchiveRequest
	*/
	ListAnalysisArchive(ctx context.Context) ApiListAnalysisArchiveRequest

	// ListAnalysisArchiveExecute executes the request
	//  @return []ArchivedAnalysis
	ListAnalysisArchiveExecute(r ApiListAnalysisArchiveRequest) ([]ArchivedAnalysis, *http.Response, error)

	/*
	ListAnalysisArchiveRules Method for ListAnalysisArchiveRules

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListAnalysisArchiveRulesRequest
	*/
	ListAnalysisArchiveRules(ctx context.Context) ApiListAnalysisArchiveRulesRequest

	// ListAnalysisArchiveRulesExecute executes the request
	//  @return []AnalysisArchiveTransitionRule
	ListAnalysisArchiveRulesExecute(r ApiListAnalysisArchiveRulesRequest) ([]AnalysisArchiveTransitionRule, *http.Response, error)

	/*
	ListArchives Method for ListArchives

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListArchivesRequest
	*/
	ListArchives(ctx context.Context) ApiListArchivesRequest

	// ListArchivesExecute executes the request
	//  @return ArchiveSummary
	ListArchivesExecute(r ApiListArchivesRequest) (*ArchiveSummary, *http.Response, error)
}

// ArchivesAPIService ArchivesAPI service
type ArchivesAPIService service

type ApiArchiveImageAnalysisRequest struct {
	ctx context.Context
	ApiService ArchivesAPI
	imageReferences *[]string
}

func (r ApiArchiveImageAnalysisRequest) ImageReferences(imageReferences []string) ApiArchiveImageAnalysisRequest {
	r.imageReferences = &imageReferences
	return r
}

func (r ApiArchiveImageAnalysisRequest) Execute() ([]AnalysisArchiveAddResult, *http.Response, error) {
	return r.ApiService.ArchiveImageAnalysisExecute(r)
}

/*
ArchiveImageAnalysis Method for ArchiveImageAnalysis

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiArchiveImageAnalysisRequest
*/
func (a *ArchivesAPIService) ArchiveImageAnalysis(ctx context.Context) ApiArchiveImageAnalysisRequest {
	return ApiArchiveImageAnalysisRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []AnalysisArchiveAddResult
func (a *ArchivesAPIService) ArchiveImageAnalysisExecute(r ApiArchiveImageAnalysisRequest) ([]AnalysisArchiveAddResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []AnalysisArchiveAddResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArchivesAPIService.ArchiveImageAnalysis")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/archives/images"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.imageReferences == nil {
		return localVarReturnValue, nil, reportError("imageReferences is required and must be specified")
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
	localVarPostBody = r.imageReferences
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

type ApiCreateAnalysisArchiveRuleRequest struct {
	ctx context.Context
	ApiService ArchivesAPI
	rule *AnalysisArchiveTransitionRule
}

func (r ApiCreateAnalysisArchiveRuleRequest) Rule(rule AnalysisArchiveTransitionRule) ApiCreateAnalysisArchiveRuleRequest {
	r.rule = &rule
	return r
}

func (r ApiCreateAnalysisArchiveRuleRequest) Execute() (*AnalysisArchiveTransitionRule, *http.Response, error) {
	return r.ApiService.CreateAnalysisArchiveRuleExecute(r)
}

/*
CreateAnalysisArchiveRule Method for CreateAnalysisArchiveRule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateAnalysisArchiveRuleRequest
*/
func (a *ArchivesAPIService) CreateAnalysisArchiveRule(ctx context.Context) ApiCreateAnalysisArchiveRuleRequest {
	return ApiCreateAnalysisArchiveRuleRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AnalysisArchiveTransitionRule
func (a *ArchivesAPIService) CreateAnalysisArchiveRuleExecute(r ApiCreateAnalysisArchiveRuleRequest) (*AnalysisArchiveTransitionRule, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AnalysisArchiveTransitionRule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArchivesAPIService.CreateAnalysisArchiveRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/archives/rules"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.rule == nil {
		return localVarReturnValue, nil, reportError("rule is required and must be specified")
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
	localVarPostBody = r.rule
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

type ApiDeleteAnalysisArchiveRuleRequest struct {
	ctx context.Context
	ApiService ArchivesAPI
	ruleId string
}

func (r ApiDeleteAnalysisArchiveRuleRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteAnalysisArchiveRuleExecute(r)
}

/*
DeleteAnalysisArchiveRule Method for DeleteAnalysisArchiveRule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ruleId
 @return ApiDeleteAnalysisArchiveRuleRequest
*/
func (a *ArchivesAPIService) DeleteAnalysisArchiveRule(ctx context.Context, ruleId string) ApiDeleteAnalysisArchiveRuleRequest {
	return ApiDeleteAnalysisArchiveRuleRequest{
		ApiService: a,
		ctx: ctx,
		ruleId: ruleId,
	}
}

// Execute executes the request
func (a *ArchivesAPIService) DeleteAnalysisArchiveRuleExecute(r ApiDeleteAnalysisArchiveRuleRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArchivesAPIService.DeleteAnalysisArchiveRule")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/archives/rules/{rule_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"rule_id"+"}", url.PathEscape(parameterValueToString(r.ruleId, "ruleId")), -1)

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
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeleteArchivedAnalysisRequest struct {
	ctx context.Context
	ApiService ArchivesAPI
	imageDigest string
	force *bool
}

func (r ApiDeleteArchivedAnalysisRequest) Force(force bool) ApiDeleteArchivedAnalysisRequest {
	r.force = &force
	return r
}

func (r ApiDeleteArchivedAnalysisRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteArchivedAnalysisExecute(r)
}

/*
DeleteArchivedAnalysis Method for DeleteArchivedAnalysis

Performs a synchronous archive deletion

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param imageDigest
 @return ApiDeleteArchivedAnalysisRequest
*/
func (a *ArchivesAPIService) DeleteArchivedAnalysis(ctx context.Context, imageDigest string) ApiDeleteArchivedAnalysisRequest {
	return ApiDeleteArchivedAnalysisRequest{
		ApiService: a,
		ctx: ctx,
		imageDigest: imageDigest,
	}
}

// Execute executes the request
func (a *ArchivesAPIService) DeleteArchivedAnalysisExecute(r ApiDeleteArchivedAnalysisRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArchivesAPIService.DeleteArchivedAnalysis")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/archives/images/{image_digest}"
	localVarPath = strings.Replace(localVarPath, "{"+"image_digest"+"}", url.PathEscape(parameterValueToString(r.imageDigest, "imageDigest")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.force != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "force", r.force, "form", "")
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
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetAnalysisArchiveRuleRequest struct {
	ctx context.Context
	ApiService ArchivesAPI
	ruleId string
}

func (r ApiGetAnalysisArchiveRuleRequest) Execute() (*AnalysisArchiveTransitionRule, *http.Response, error) {
	return r.ApiService.GetAnalysisArchiveRuleExecute(r)
}

/*
GetAnalysisArchiveRule Method for GetAnalysisArchiveRule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ruleId
 @return ApiGetAnalysisArchiveRuleRequest
*/
func (a *ArchivesAPIService) GetAnalysisArchiveRule(ctx context.Context, ruleId string) ApiGetAnalysisArchiveRuleRequest {
	return ApiGetAnalysisArchiveRuleRequest{
		ApiService: a,
		ctx: ctx,
		ruleId: ruleId,
	}
}

// Execute executes the request
//  @return AnalysisArchiveTransitionRule
func (a *ArchivesAPIService) GetAnalysisArchiveRuleExecute(r ApiGetAnalysisArchiveRuleRequest) (*AnalysisArchiveTransitionRule, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AnalysisArchiveTransitionRule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArchivesAPIService.GetAnalysisArchiveRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/archives/rules/{rule_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"rule_id"+"}", url.PathEscape(parameterValueToString(r.ruleId, "ruleId")), -1)

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

type ApiGetArchivedAnalysisRequest struct {
	ctx context.Context
	ApiService ArchivesAPI
	imageDigest string
}

func (r ApiGetArchivedAnalysisRequest) Execute() (*ArchivedAnalysis, *http.Response, error) {
	return r.ApiService.GetArchivedAnalysisExecute(r)
}

/*
GetArchivedAnalysis Method for GetArchivedAnalysis

Returns the archive metadata record identifying the image and tags for the analysis in the archive.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param imageDigest The image digest to identify the image analysis
 @return ApiGetArchivedAnalysisRequest
*/
func (a *ArchivesAPIService) GetArchivedAnalysis(ctx context.Context, imageDigest string) ApiGetArchivedAnalysisRequest {
	return ApiGetArchivedAnalysisRequest{
		ApiService: a,
		ctx: ctx,
		imageDigest: imageDigest,
	}
}

// Execute executes the request
//  @return ArchivedAnalysis
func (a *ArchivesAPIService) GetArchivedAnalysisExecute(r ApiGetArchivedAnalysisRequest) (*ArchivedAnalysis, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ArchivedAnalysis
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArchivesAPIService.GetArchivedAnalysis")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/archives/images/{image_digest}"
	localVarPath = strings.Replace(localVarPath, "{"+"image_digest"+"}", url.PathEscape(parameterValueToString(r.imageDigest, "imageDigest")), -1)

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

type ApiListAnalysisArchiveRequest struct {
	ctx context.Context
	ApiService ArchivesAPI
}

func (r ApiListAnalysisArchiveRequest) Execute() ([]ArchivedAnalysis, *http.Response, error) {
	return r.ApiService.ListAnalysisArchiveExecute(r)
}

/*
ListAnalysisArchive Method for ListAnalysisArchive

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListAnalysisArchiveRequest
*/
func (a *ArchivesAPIService) ListAnalysisArchive(ctx context.Context) ApiListAnalysisArchiveRequest {
	return ApiListAnalysisArchiveRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ArchivedAnalysis
func (a *ArchivesAPIService) ListAnalysisArchiveExecute(r ApiListAnalysisArchiveRequest) ([]ArchivedAnalysis, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ArchivedAnalysis
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArchivesAPIService.ListAnalysisArchive")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/archives/images"

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

type ApiListAnalysisArchiveRulesRequest struct {
	ctx context.Context
	ApiService ArchivesAPI
	systemGlobal *bool
}

// If true include system global rules (owned by admin) even for non-admin users. Defaults to true if not set. Can be set to false to exclude globals
func (r ApiListAnalysisArchiveRulesRequest) SystemGlobal(systemGlobal bool) ApiListAnalysisArchiveRulesRequest {
	r.systemGlobal = &systemGlobal
	return r
}

func (r ApiListAnalysisArchiveRulesRequest) Execute() ([]AnalysisArchiveTransitionRule, *http.Response, error) {
	return r.ApiService.ListAnalysisArchiveRulesExecute(r)
}

/*
ListAnalysisArchiveRules Method for ListAnalysisArchiveRules

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListAnalysisArchiveRulesRequest
*/
func (a *ArchivesAPIService) ListAnalysisArchiveRules(ctx context.Context) ApiListAnalysisArchiveRulesRequest {
	return ApiListAnalysisArchiveRulesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []AnalysisArchiveTransitionRule
func (a *ArchivesAPIService) ListAnalysisArchiveRulesExecute(r ApiListAnalysisArchiveRulesRequest) ([]AnalysisArchiveTransitionRule, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []AnalysisArchiveTransitionRule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArchivesAPIService.ListAnalysisArchiveRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/archives/rules"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.systemGlobal != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "system_global", r.systemGlobal, "form", "")
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

type ApiListArchivesRequest struct {
	ctx context.Context
	ApiService ArchivesAPI
}

func (r ApiListArchivesRequest) Execute() (*ArchiveSummary, *http.Response, error) {
	return r.ApiService.ListArchivesExecute(r)
}

/*
ListArchives Method for ListArchives

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListArchivesRequest
*/
func (a *ArchivesAPIService) ListArchives(ctx context.Context) ApiListArchivesRequest {
	return ApiListArchivesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ArchiveSummary
func (a *ArchivesAPIService) ListArchivesExecute(r ApiListArchivesRequest) (*ArchiveSummary, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ArchiveSummary
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArchivesAPIService.ListArchives")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/archives"

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
