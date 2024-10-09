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


type CatalogApi interface {

	/*
	GetDeploymentHistory List Deployment History

	Returns list of deployment history entries

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetDeploymentHistoryRequest
	*/
	GetDeploymentHistory(ctx context.Context) ApiGetDeploymentHistoryRequest

	// GetDeploymentHistoryExecute executes the request
	//  @return DeploymentHistoryList
	GetDeploymentHistoryExecute(r ApiGetDeploymentHistoryRequest) (*DeploymentHistoryList, *http.Response, error)
}

// CatalogApiService CatalogApi service
type CatalogApiService service

type ApiGetDeploymentHistoryRequest struct {
	ctx context.Context
	ApiService CatalogApi
}

func (r ApiGetDeploymentHistoryRequest) Execute() (*DeploymentHistoryList, *http.Response, error) {
	return r.ApiService.GetDeploymentHistoryExecute(r)
}

/*
GetDeploymentHistory List Deployment History

Returns list of deployment history entries

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetDeploymentHistoryRequest
*/
func (a *CatalogApiService) GetDeploymentHistory(ctx context.Context) ApiGetDeploymentHistoryRequest {
	return ApiGetDeploymentHistoryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DeploymentHistoryList
func (a *CatalogApiService) GetDeploymentHistoryExecute(r ApiGetDeploymentHistoryRequest) (*DeploymentHistoryList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DeploymentHistoryList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CatalogApiService.GetDeploymentHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/deployment-history"

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
