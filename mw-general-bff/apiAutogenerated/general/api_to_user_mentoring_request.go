/*
Masters way general API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)


// ToUserMentoringRequestAPIService ToUserMentoringRequestAPI service
type ToUserMentoringRequestAPIService service

type ApiCreateUserMentoringRequestRequest struct {
	ctx context.Context
	ApiService *ToUserMentoringRequestAPIService
	request *SchemasCreateToUserMentoringRequestPayload
}

// query params
func (r ApiCreateUserMentoringRequestRequest) Request(request SchemasCreateToUserMentoringRequestPayload) ApiCreateUserMentoringRequestRequest {
	r.request = &request
	return r
}

func (r ApiCreateUserMentoringRequestRequest) Execute() (*SchemasToUserMentoringRequestResponse, *http.Response, error) {
	return r.ApiService.CreateUserMentoringRequestExecute(r)
}

/*
CreateUserMentoringRequest Create a new userMentoringRequest

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateUserMentoringRequestRequest
*/
func (a *ToUserMentoringRequestAPIService) CreateUserMentoringRequest(ctx context.Context) ApiCreateUserMentoringRequestRequest {
	return ApiCreateUserMentoringRequestRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SchemasToUserMentoringRequestResponse
func (a *ToUserMentoringRequestAPIService) CreateUserMentoringRequestExecute(r ApiCreateUserMentoringRequestRequest) (*SchemasToUserMentoringRequestResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SchemasToUserMentoringRequestResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ToUserMentoringRequestAPIService.CreateUserMentoringRequest")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/toUserMentoringRequests"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.request == nil {
		return localVarReturnValue, nil, reportError("request is required and must be specified")
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
	localVarPostBody = r.request
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

// Execute executes the request
//  @return SchemasToUserMentoringRequestResponseStream
func (a *ToUserMentoringRequestAPIService) CreateUserMentoringRequestStreamExecute(r ApiCreateUserMentoringRequestRequest, request *http.Request, GoogleAccessToken string) (*SchemasToUserMentoringRequestResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarReturnValue  *SchemasToUserMentoringRequestResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ToUserMentoringRequestAPIService.CreateUserMentoringRequest")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/toUserMentoringRequests"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	fmt.Println(localVarQueryParams)


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
	req, err := http.NewRequest(localVarHTTPMethod, localVarPath, request.Body)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	for key, values := range request.Header {
	   if key == "Origin" { continue }
	   for _, value := range values {
	       req.Header.Add(key, value)
	   }
	}

	req.Header.Add("GoogleAccessToken", GoogleAccessToken)

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

type ApiDeleteToUserMentoringRequestRequest struct {
	ctx context.Context
	ApiService *ToUserMentoringRequestAPIService
	userUuid string
	wayUuid string
}

func (r ApiDeleteToUserMentoringRequestRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteToUserMentoringRequestExecute(r)
}

/*
DeleteToUserMentoringRequest Delete toUserMentoringReques by UUID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userUuid user UUID
 @param wayUuid way UUID
 @return ApiDeleteToUserMentoringRequestRequest
*/
func (a *ToUserMentoringRequestAPIService) DeleteToUserMentoringRequest(ctx context.Context, userUuid string, wayUuid string) ApiDeleteToUserMentoringRequestRequest {
	return ApiDeleteToUserMentoringRequestRequest{
		ApiService: a,
		ctx: ctx,
		userUuid: userUuid,
		wayUuid: wayUuid,
	}
}

// Execute executes the request
func (a *ToUserMentoringRequestAPIService) DeleteToUserMentoringRequestExecute(r ApiDeleteToUserMentoringRequestRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ToUserMentoringRequestAPIService.DeleteToUserMentoringRequest")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/toUserMentoringRequests/{userUuid}/{wayUuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"userUuid"+"}", url.PathEscape(parameterValueToString(r.userUuid, "userUuid")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wayUuid"+"}", url.PathEscape(parameterValueToString(r.wayUuid, "wayUuid")), -1)

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

// Execute executes the request
func (a *ToUserMentoringRequestAPIService) DeleteToUserMentoringRequestStreamExecute(r ApiDeleteToUserMentoringRequestRequest, request *http.Request, GoogleAccessToken string) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ToUserMentoringRequestAPIService.DeleteToUserMentoringRequest")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/toUserMentoringRequests/{userUuid}/{wayUuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"userUuid"+"}", url.PathEscape(parameterValueToString(r.userUuid, "userUuid")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wayUuid"+"}", url.PathEscape(parameterValueToString(r.wayUuid, "wayUuid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	fmt.Println(localVarQueryParams)



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
	req, err := http.NewRequest(localVarHTTPMethod, localVarPath, request.Body)
	if err != nil {
		return nil, err
	}

	for key, values := range request.Header {
	   if key == "Origin" { continue }
	   for _, value := range values {
	       req.Header.Add(key, value)
	   }
	}

	req.Header.Add("GoogleAccessToken", GoogleAccessToken)

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