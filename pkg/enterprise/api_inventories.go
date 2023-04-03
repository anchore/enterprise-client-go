/*
Anchore Enterprise API Server

This is the Anchore Enterprise API. It provides additional external API routes and functionality for enterprise users.

API version: 0.7.0
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
)

// Linger please
var (
	_ _context.Context
)

type InventoriesApi interface {

	/*
	DeleteInventory Delete runtime inventory by type and context

	Delete runtime inventory by type and context

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiDeleteInventoryRequest
	*/
	DeleteInventory(ctx _context.Context) ApiDeleteInventoryRequest

	// DeleteInventoryExecute executes the request
	DeleteInventoryExecute(r ApiDeleteInventoryRequest) (*_nethttp.Response, error)

	/*
	GetImageInventory Return a list of the images in inventories for this account

	Returns a list of the images that are in use

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiGetImageInventoryRequest
	*/
	GetImageInventory(ctx _context.Context) ApiGetImageInventoryRequest

	// GetImageInventoryExecute executes the request
	//  @return []InventoryItem
	GetImageInventoryExecute(r ApiGetImageInventoryRequest) ([]InventoryItem, *_nethttp.Response, error)

	/*
	SyncImageInventory synchronizes the list of the images in a given cluster for the inventory

	synchronizes the list of the images that are in use

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiSyncImageInventoryRequest
	*/
	SyncImageInventory(ctx _context.Context) ApiSyncImageInventoryRequest

	// SyncImageInventoryExecute executes the request
	//  @return []InventoryItem
	SyncImageInventoryExecute(r ApiSyncImageInventoryRequest) ([]InventoryItem, *_nethttp.Response, error)
}

// InventoriesApiService InventoriesApi service
type InventoriesApiService service

type ApiDeleteInventoryRequest struct {
	ctx _context.Context
	ApiService InventoriesApi
	inventoryType *string
	context *string
	imageDigest *string
	xAnchoreAccount *string
}

func (r ApiDeleteInventoryRequest) InventoryType(inventoryType string) ApiDeleteInventoryRequest {
	r.inventoryType = &inventoryType
	return r
}
func (r ApiDeleteInventoryRequest) Context(context string) ApiDeleteInventoryRequest {
	r.context = &context
	return r
}
func (r ApiDeleteInventoryRequest) ImageDigest(imageDigest string) ApiDeleteInventoryRequest {
	r.imageDigest = &imageDigest
	return r
}
// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiDeleteInventoryRequest) XAnchoreAccount(xAnchoreAccount string) ApiDeleteInventoryRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

func (r ApiDeleteInventoryRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteInventoryExecute(r)
}

/*
DeleteInventory Delete runtime inventory by type and context

Delete runtime inventory by type and context

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeleteInventoryRequest
*/
func (a *InventoriesApiService) DeleteInventory(ctx _context.Context) ApiDeleteInventoryRequest {
	return ApiDeleteInventoryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *InventoriesApiService) DeleteInventoryExecute(r ApiDeleteInventoryRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoriesApiService.DeleteInventory")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventories"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inventoryType == nil {
		return nil, reportError("inventoryType is required and must be specified")
	}
	if r.context == nil {
		return nil, reportError("context is required and must be specified")
	}

	localVarQueryParams.Add("inventory_type", parameterToString(*r.inventoryType, ""))
	localVarQueryParams.Add("context", parameterToString(*r.context, ""))
	if r.imageDigest != nil {
		localVarQueryParams.Add("image_digest", parameterToString(*r.imageDigest, ""))
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

type ApiGetImageInventoryRequest struct {
	ctx _context.Context
	ApiService InventoriesApi
	inventoryType *string
	imageDigest *string
	context *string
	xAnchoreAccount *string
}

func (r ApiGetImageInventoryRequest) InventoryType(inventoryType string) ApiGetImageInventoryRequest {
	r.inventoryType = &inventoryType
	return r
}
func (r ApiGetImageInventoryRequest) ImageDigest(imageDigest string) ApiGetImageInventoryRequest {
	r.imageDigest = &imageDigest
	return r
}
func (r ApiGetImageInventoryRequest) Context(context string) ApiGetImageInventoryRequest {
	r.context = &context
	return r
}
// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiGetImageInventoryRequest) XAnchoreAccount(xAnchoreAccount string) ApiGetImageInventoryRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

func (r ApiGetImageInventoryRequest) Execute() ([]InventoryItem, *_nethttp.Response, error) {
	return r.ApiService.GetImageInventoryExecute(r)
}

/*
GetImageInventory Return a list of the images in inventories for this account

Returns a list of the images that are in use

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetImageInventoryRequest
*/
func (a *InventoriesApiService) GetImageInventory(ctx _context.Context) ApiGetImageInventoryRequest {
	return ApiGetImageInventoryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []InventoryItem
func (a *InventoriesApiService) GetImageInventoryExecute(r ApiGetImageInventoryRequest) ([]InventoryItem, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []InventoryItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoriesApiService.GetImageInventory")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventories"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.inventoryType != nil {
		localVarQueryParams.Add("inventory_type", parameterToString(*r.inventoryType, ""))
	}
	if r.imageDigest != nil {
		localVarQueryParams.Add("image_digest", parameterToString(*r.imageDigest, ""))
	}
	if r.context != nil {
		localVarQueryParams.Add("context", parameterToString(*r.context, ""))
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

type ApiSyncImageInventoryRequest struct {
	ctx _context.Context
	ApiService InventoriesApi
	inventory *InventoryReport
	xAnchoreAccount *string
}

func (r ApiSyncImageInventoryRequest) Inventory(inventory InventoryReport) ApiSyncImageInventoryRequest {
	r.inventory = &inventory
	return r
}
// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiSyncImageInventoryRequest) XAnchoreAccount(xAnchoreAccount string) ApiSyncImageInventoryRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

func (r ApiSyncImageInventoryRequest) Execute() ([]InventoryItem, *_nethttp.Response, error) {
	return r.ApiService.SyncImageInventoryExecute(r)
}

/*
SyncImageInventory synchronizes the list of the images in a given cluster for the inventory

synchronizes the list of the images that are in use

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSyncImageInventoryRequest
*/
func (a *InventoriesApiService) SyncImageInventory(ctx _context.Context) ApiSyncImageInventoryRequest {
	return ApiSyncImageInventoryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []InventoryItem
func (a *InventoriesApiService) SyncImageInventoryExecute(r ApiSyncImageInventoryRequest) ([]InventoryItem, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []InventoryItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoriesApiService.SyncImageInventory")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventories"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inventory == nil {
		return localVarReturnValue, nil, reportError("inventory is required and must be specified")
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
	localVarPostBody = r.inventory
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
