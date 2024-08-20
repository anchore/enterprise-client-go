/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.6.2
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


type ReportsApi interface {

	/*
	GetGlobalQueryResult Method for GetGlobalQueryResult

	Get a single saved global query result

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param resultUuid
	@return ApiGetGlobalQueryResultRequest
	*/
	GetGlobalQueryResult(ctx context.Context, resultUuid string) ApiGetGlobalQueryResultRequest

	// GetGlobalQueryResultExecute executes the request
	GetGlobalQueryResultExecute(r ApiGetGlobalQueryResultRequest) (*http.Response, error)

	/*
	GetQueryResult Method for GetQueryResult

	Get a single saved query result

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param resultUuid
	@return ApiGetQueryResultRequest
	*/
	GetQueryResult(ctx context.Context, resultUuid string) ApiGetQueryResultRequest

	// GetQueryResultExecute executes the request
	GetQueryResultExecute(r ApiGetQueryResultRequest) (*http.Response, error)
}

// ReportsApiService ReportsApi service
type ReportsApiService service

type ApiGetGlobalQueryResultRequest struct {
	ctx context.Context
	ApiService ReportsApi
	resultUuid string
	page *int32
}

// Page number to fetch. If omitted, &#39;1&#39; is default. Page numbers start at 1
func (r ApiGetGlobalQueryResultRequest) Page(page int32) ApiGetGlobalQueryResultRequest {
	r.page = &page
	return r
}

func (r ApiGetGlobalQueryResultRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetGlobalQueryResultExecute(r)
}

/*
GetGlobalQueryResult Method for GetGlobalQueryResult

Get a single saved global query result

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param resultUuid
 @return ApiGetGlobalQueryResultRequest
*/
func (a *ReportsApiService) GetGlobalQueryResult(ctx context.Context, resultUuid string) ApiGetGlobalQueryResultRequest {
	return ApiGetGlobalQueryResultRequest{
		ApiService: a,
		ctx: ctx,
		resultUuid: resultUuid,
	}
}

// Execute executes the request
func (a *ReportsApiService) GetGlobalQueryResultExecute(r ApiGetGlobalQueryResultRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportsApiService.GetGlobalQueryResult")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reporting/reports/global/scheduled-query-results/{result_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"result_uuid"+"}", url.PathEscape(parameterToString(r.resultUuid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiGetQueryResultRequest struct {
	ctx context.Context
	ApiService ReportsApi
	resultUuid string
	page *int32
}

// Page number to fetch. If omitted, &#39;1&#39; is default. Page numbers start at 1
func (r ApiGetQueryResultRequest) Page(page int32) ApiGetQueryResultRequest {
	r.page = &page
	return r
}

func (r ApiGetQueryResultRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetQueryResultExecute(r)
}

/*
GetQueryResult Method for GetQueryResult

Get a single saved query result

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param resultUuid
 @return ApiGetQueryResultRequest
*/
func (a *ReportsApiService) GetQueryResult(ctx context.Context, resultUuid string) ApiGetQueryResultRequest {
	return ApiGetQueryResultRequest{
		ApiService: a,
		ctx: ctx,
		resultUuid: resultUuid,
	}
}

// Execute executes the request
func (a *ReportsApiService) GetQueryResultExecute(r ApiGetQueryResultRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportsApiService.GetQueryResult")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reporting/scheduled-query-results/{result_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"result_uuid"+"}", url.PathEscape(parameterToString(r.resultUuid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
