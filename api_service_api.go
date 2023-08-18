/*
OpenAPI definition

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// ServiceApiApiService ServiceApiApi service
type ServiceApiApiService service

type ApiAdoptServiceRequest struct {
	ctx context.Context
	ApiService *ServiceApiApiService
	adoptServiceRequest *AdoptServiceRequest
	authorization *string
	dPoP *string
}

func (r ApiAdoptServiceRequest) AdoptServiceRequest(adoptServiceRequest AdoptServiceRequest) ApiAdoptServiceRequest {
	r.adoptServiceRequest = &adoptServiceRequest
	return r
}

func (r ApiAdoptServiceRequest) Authorization(authorization string) ApiAdoptServiceRequest {
	r.authorization = &authorization
	return r
}

func (r ApiAdoptServiceRequest) DPoP(dPoP string) ApiAdoptServiceRequest {
	r.dPoP = &dPoP
	return r
}

func (r ApiAdoptServiceRequest) Execute() (*ServiceInstanceManagementResponse, *http.Response, error) {
	return r.ApiService.AdoptServiceExecute(r)
}

/*
AdoptService Method for AdoptService

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAdoptServiceRequest
*/
func (a *ServiceApiApiService) AdoptService(ctx context.Context) ApiAdoptServiceRequest {
	return ApiAdoptServiceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ServiceInstanceManagementResponse
func (a *ServiceApiApiService) AdoptServiceExecute(r ApiAdoptServiceRequest) (*ServiceInstanceManagementResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServiceInstanceManagementResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceApiApiService.AdoptService")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/service/adopt"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.adoptServiceRequest == nil {
		return localVarReturnValue, nil, reportError("adoptServiceRequest is required and must be specified")
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
	if r.authorization != nil {
		localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
	}
	if r.dPoP != nil {
		localVarHeaderParams["DPoP"] = parameterToString(*r.dPoP, "")
	}
	// body params
	localVarPostBody = r.adoptServiceRequest
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

type ApiCreateServiceRequest struct {
	ctx context.Context
	ApiService *ServiceApiApiService
	createServiceRequest *CreateServiceRequest
	authorization *string
	dPoP *string
}

func (r ApiCreateServiceRequest) CreateServiceRequest(createServiceRequest CreateServiceRequest) ApiCreateServiceRequest {
	r.createServiceRequest = &createServiceRequest
	return r
}

func (r ApiCreateServiceRequest) Authorization(authorization string) ApiCreateServiceRequest {
	r.authorization = &authorization
	return r
}

func (r ApiCreateServiceRequest) DPoP(dPoP string) ApiCreateServiceRequest {
	r.dPoP = &dPoP
	return r
}

func (r ApiCreateServiceRequest) Execute() (*Service, *http.Response, error) {
	return r.ApiService.CreateServiceExecute(r)
}

/*
CreateService Method for CreateService

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateServiceRequest
*/
func (a *ServiceApiApiService) CreateService(ctx context.Context) ApiCreateServiceRequest {
	return ApiCreateServiceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Service
func (a *ServiceApiApiService) CreateServiceExecute(r ApiCreateServiceRequest) (*Service, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Service
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceApiApiService.CreateService")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/service"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createServiceRequest == nil {
		return localVarReturnValue, nil, reportError("createServiceRequest is required and must be specified")
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
	if r.authorization != nil {
		localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
	}
	if r.dPoP != nil {
		localVarHeaderParams["DPoP"] = parameterToString(*r.dPoP, "")
	}
	// body params
	localVarPostBody = r.createServiceRequest
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

type ApiDeleteServiceRequest struct {
	ctx context.Context
	ApiService *ServiceApiApiService
	deleteServiceRequest *DeleteServiceRequest
	authorization *string
	dPoP *string
}

func (r ApiDeleteServiceRequest) DeleteServiceRequest(deleteServiceRequest DeleteServiceRequest) ApiDeleteServiceRequest {
	r.deleteServiceRequest = &deleteServiceRequest
	return r
}

func (r ApiDeleteServiceRequest) Authorization(authorization string) ApiDeleteServiceRequest {
	r.authorization = &authorization
	return r
}

func (r ApiDeleteServiceRequest) DPoP(dPoP string) ApiDeleteServiceRequest {
	r.dPoP = &dPoP
	return r
}

func (r ApiDeleteServiceRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteServiceExecute(r)
}

/*
DeleteService Method for DeleteService

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeleteServiceRequest
*/
func (a *ServiceApiApiService) DeleteService(ctx context.Context) ApiDeleteServiceRequest {
	return ApiDeleteServiceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ServiceApiApiService) DeleteServiceExecute(r ApiDeleteServiceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceApiApiService.DeleteService")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/service/remove"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.deleteServiceRequest == nil {
		return nil, reportError("deleteServiceRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	if r.authorization != nil {
		localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
	}
	if r.dPoP != nil {
		localVarHeaderParams["DPoP"] = parameterToString(*r.dPoP, "")
	}
	// body params
	localVarPostBody = r.deleteServiceRequest
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

type ApiFindServiceRequest struct {
	ctx context.Context
	ApiService *ServiceApiApiService
	id int64
	authorization *string
	dPoP *string
}

func (r ApiFindServiceRequest) Authorization(authorization string) ApiFindServiceRequest {
	r.authorization = &authorization
	return r
}

func (r ApiFindServiceRequest) DPoP(dPoP string) ApiFindServiceRequest {
	r.dPoP = &dPoP
	return r
}

func (r ApiFindServiceRequest) Execute() ([]ServiceInstanceManagementResponse, *http.Response, error) {
	return r.ApiService.FindServiceExecute(r)
}

/*
FindService Method for FindService

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiFindServiceRequest
*/
func (a *ServiceApiApiService) FindService(ctx context.Context, id int64) ApiFindServiceRequest {
	return ApiFindServiceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return []ServiceInstanceManagementResponse
func (a *ServiceApiApiService) FindServiceExecute(r ApiFindServiceRequest) ([]ServiceInstanceManagementResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ServiceInstanceManagementResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceApiApiService.FindService")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/service/find/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

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
	if r.authorization != nil {
		localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
	}
	if r.dPoP != nil {
		localVarHeaderParams["DPoP"] = parameterToString(*r.dPoP, "")
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

type ApiGetOrphansRequest struct {
	ctx context.Context
	ApiService *ServiceApiApiService
	authorization *string
	dPoP *string
}

func (r ApiGetOrphansRequest) Authorization(authorization string) ApiGetOrphansRequest {
	r.authorization = &authorization
	return r
}

func (r ApiGetOrphansRequest) DPoP(dPoP string) ApiGetOrphansRequest {
	r.dPoP = &dPoP
	return r
}

func (r ApiGetOrphansRequest) Execute() (*OrphanServiceResponse, *http.Response, error) {
	return r.ApiService.GetOrphansExecute(r)
}

/*
GetOrphans Method for GetOrphans

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetOrphansRequest
*/
func (a *ServiceApiApiService) GetOrphans(ctx context.Context) ApiGetOrphansRequest {
	return ApiGetOrphansRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OrphanServiceResponse
func (a *ServiceApiApiService) GetOrphansExecute(r ApiGetOrphansRequest) (*OrphanServiceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OrphanServiceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceApiApiService.GetOrphans")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/service/orphans"

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
	if r.authorization != nil {
		localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
	}
	if r.dPoP != nil {
		localVarHeaderParams["DPoP"] = parameterToString(*r.dPoP, "")
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

type ApiMoveServiceRequest struct {
	ctx context.Context
	ApiService *ServiceApiApiService
	moveServiceRequest *MoveServiceRequest
	authorization *string
	dPoP *string
}

func (r ApiMoveServiceRequest) MoveServiceRequest(moveServiceRequest MoveServiceRequest) ApiMoveServiceRequest {
	r.moveServiceRequest = &moveServiceRequest
	return r
}

func (r ApiMoveServiceRequest) Authorization(authorization string) ApiMoveServiceRequest {
	r.authorization = &authorization
	return r
}

func (r ApiMoveServiceRequest) DPoP(dPoP string) ApiMoveServiceRequest {
	r.dPoP = &dPoP
	return r
}

func (r ApiMoveServiceRequest) Execute() (*ServiceInstanceManagementResponse, *http.Response, error) {
	return r.ApiService.MoveServiceExecute(r)
}

/*
MoveService Method for MoveService

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMoveServiceRequest
*/
func (a *ServiceApiApiService) MoveService(ctx context.Context) ApiMoveServiceRequest {
	return ApiMoveServiceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ServiceInstanceManagementResponse
func (a *ServiceApiApiService) MoveServiceExecute(r ApiMoveServiceRequest) (*ServiceInstanceManagementResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServiceInstanceManagementResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceApiApiService.MoveService")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/service/move"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.moveServiceRequest == nil {
		return localVarReturnValue, nil, reportError("moveServiceRequest is required and must be specified")
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
	if r.authorization != nil {
		localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
	}
	if r.dPoP != nil {
		localVarHeaderParams["DPoP"] = parameterToString(*r.dPoP, "")
	}
	// body params
	localVarPostBody = r.moveServiceRequest
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

type ApiRemoveOrphanServiceRequest struct {
	ctx context.Context
	ApiService *ServiceApiApiService
	removeOrphanRequest *RemoveOrphanRequest
	authorization *string
	dPoP *string
}

func (r ApiRemoveOrphanServiceRequest) RemoveOrphanRequest(removeOrphanRequest RemoveOrphanRequest) ApiRemoveOrphanServiceRequest {
	r.removeOrphanRequest = &removeOrphanRequest
	return r
}

func (r ApiRemoveOrphanServiceRequest) Authorization(authorization string) ApiRemoveOrphanServiceRequest {
	r.authorization = &authorization
	return r
}

func (r ApiRemoveOrphanServiceRequest) DPoP(dPoP string) ApiRemoveOrphanServiceRequest {
	r.dPoP = &dPoP
	return r
}

func (r ApiRemoveOrphanServiceRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveOrphanServiceExecute(r)
}

/*
RemoveOrphanService Method for RemoveOrphanService

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRemoveOrphanServiceRequest
*/
func (a *ServiceApiApiService) RemoveOrphanService(ctx context.Context) ApiRemoveOrphanServiceRequest {
	return ApiRemoveOrphanServiceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ServiceApiApiService) RemoveOrphanServiceExecute(r ApiRemoveOrphanServiceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceApiApiService.RemoveOrphanService")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/service/orphans/remove"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.removeOrphanRequest == nil {
		return nil, reportError("removeOrphanRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	if r.authorization != nil {
		localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
	}
	if r.dPoP != nil {
		localVarHeaderParams["DPoP"] = parameterToString(*r.dPoP, "")
	}
	// body params
	localVarPostBody = r.removeOrphanRequest
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
