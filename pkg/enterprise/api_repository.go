/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.7.2
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
)


// RepositoryAPIService RepositoryAPI service
type RepositoryAPIService service

type ApiAddRepositoryRequest struct {
	ctx context.Context
	ApiService *RepositoryAPIService
	repository *string
	autoSubscribe *bool
	dryRun *bool
	excludeExistingTags *bool
	xAnchoreAccount *string
}

// Full repository to add e.g. docker.io/library/alpine
func (r ApiAddRepositoryRequest) Repository(repository string) ApiAddRepositoryRequest {
	r.repository = &repository
	return r
}

// Flag to enable/disable auto tag_update activation when new images from a repo are added. Default is false.
func (r ApiAddRepositoryRequest) AutoSubscribe(autoSubscribe bool) ApiAddRepositoryRequest {
	r.autoSubscribe = &autoSubscribe
	return r
}

// Flag to return tags in the repository without actually watching the repository. Default is false.
func (r ApiAddRepositoryRequest) DryRun(dryRun bool) ApiAddRepositoryRequest {
	r.dryRun = &dryRun
	return r
}

// Flag that indicates if the watcher will exclude existing tags from the repository during the first run.  When set to &#39;true&#39;, the watcher will only add newly detected tags to the system from this time forward. Default is false.
func (r ApiAddRepositoryRequest) ExcludeExistingTags(excludeExistingTags bool) ApiAddRepositoryRequest {
	r.excludeExistingTags = &excludeExistingTags
	return r
}

// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiAddRepositoryRequest) XAnchoreAccount(xAnchoreAccount string) ApiAddRepositoryRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

func (r ApiAddRepositoryRequest) Execute() ([]Subscription, *http.Response, error) {
	return r.ApiService.AddRepositoryExecute(r)
}

/*
AddRepository Add repository to watch

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAddRepositoryRequest
*/
func (a *RepositoryAPIService) AddRepository(ctx context.Context) ApiAddRepositoryRequest {
	return ApiAddRepositoryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Subscription
func (a *RepositoryAPIService) AddRepositoryExecute(r ApiAddRepositoryRequest) ([]Subscription, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Subscription
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoryAPIService.AddRepository")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/repositories"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.repository == nil {
		return localVarReturnValue, nil, reportError("repository is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "repository", r.repository, "form", "")
	if r.autoSubscribe != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "auto_subscribe", r.autoSubscribe, "form", "")
	}
	if r.dryRun != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dry_run", r.dryRun, "form", "")
	}
	if r.excludeExistingTags != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_existing_tags", r.excludeExistingTags, "form", "")
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
