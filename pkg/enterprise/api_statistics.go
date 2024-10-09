/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.8.0
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
)


type StatisticsApi interface {

	/*
	GetSystemStatistics List System Statistics

	Returns list of system statistics with total all-time counts.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetSystemStatisticsRequest
	*/
	GetSystemStatistics(ctx context.Context) ApiGetSystemStatisticsRequest

	// GetSystemStatisticsExecute executes the request
	//  @return SystemStatisticsList
	GetSystemStatisticsExecute(r ApiGetSystemStatisticsRequest) (*SystemStatisticsList, *http.Response, error)
}

// StatisticsApiService StatisticsApi service
type StatisticsApiService service

type ApiGetSystemStatisticsRequest struct {
	ctx context.Context
	ApiService StatisticsApi
}

func (r ApiGetSystemStatisticsRequest) Execute() (*SystemStatisticsList, *http.Response, error) {
	return r.ApiService.GetSystemStatisticsExecute(r)
}

/*
GetSystemStatistics List System Statistics

Returns list of system statistics with total all-time counts.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSystemStatisticsRequest
*/
func (a *StatisticsApiService) GetSystemStatistics(ctx context.Context) ApiGetSystemStatisticsRequest {
	return ApiGetSystemStatisticsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SystemStatisticsList
func (a *StatisticsApiService) GetSystemStatisticsExecute(r ApiGetSystemStatisticsRequest) (*SystemStatisticsList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SystemStatisticsList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatisticsApiService.GetSystemStatistics")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/statistics"

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
