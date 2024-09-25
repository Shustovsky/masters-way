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


// JobTagAPIService JobTagAPI service
type JobTagAPIService service

type ApiCreateJobTagRequest struct {
	ctx context.Context
	ApiService *JobTagAPIService
	request *SchemasCreateJobTagPayload
}

// query params
func (r ApiCreateJobTagRequest) Request(request SchemasCreateJobTagPayload) ApiCreateJobTagRequest {
	r.request = &request
	return r
}

func (r ApiCreateJobTagRequest) Execute() (*SchemasJobTagResponse, *http.Response, error) {
	return r.ApiService.CreateJobTagExecute(r)
}

/*
CreateJobTag Create a new jobTag

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateJobTagRequest
*/
func (a *JobTagAPIService) CreateJobTag(ctx context.Context) ApiCreateJobTagRequest {
	return ApiCreateJobTagRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SchemasJobTagResponse
func (a *JobTagAPIService) CreateJobTagExecute(r ApiCreateJobTagRequest) (*SchemasJobTagResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SchemasJobTagResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobTagAPIService.CreateJobTag")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/jobTags"

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
//  @return SchemasJobTagResponseStream
func (a *JobTagAPIService) CreateJobTagStreamExecute(r ApiCreateJobTagRequest, request *http.Request, GoogleAccessToken string) (*SchemasJobTagResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarReturnValue  *SchemasJobTagResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobTagAPIService.CreateJobTag")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/jobTags"

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

type ApiDeleteJobTagRequest struct {
	ctx context.Context
	ApiService *JobTagAPIService
	jobTagId string
}

func (r ApiDeleteJobTagRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteJobTagExecute(r)
}

/*
DeleteJobTag Delete jobTag by UUID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param jobTagId jobTag ID
 @return ApiDeleteJobTagRequest
*/
func (a *JobTagAPIService) DeleteJobTag(ctx context.Context, jobTagId string) ApiDeleteJobTagRequest {
	return ApiDeleteJobTagRequest{
		ApiService: a,
		ctx: ctx,
		jobTagId: jobTagId,
	}
}

// Execute executes the request
func (a *JobTagAPIService) DeleteJobTagExecute(r ApiDeleteJobTagRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobTagAPIService.DeleteJobTag")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/jobTags/{jobTagId}"
	localVarPath = strings.Replace(localVarPath, "{"+"jobTagId"+"}", url.PathEscape(parameterValueToString(r.jobTagId, "jobTagId")), -1)

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
func (a *JobTagAPIService) DeleteJobTagStreamExecute(r ApiDeleteJobTagRequest, request *http.Request, GoogleAccessToken string) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobTagAPIService.DeleteJobTag")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/jobTags/{jobTagId}"
	localVarPath = strings.Replace(localVarPath, "{"+"jobTagId"+"}", url.PathEscape(parameterValueToString(r.jobTagId, "jobTagId")), -1)

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

type ApiUpdateJobTagRequest struct {
	ctx context.Context
	ApiService *JobTagAPIService
	jobTagId string
	request *SchemasUpdateJobTagPayload
}

// query params
func (r ApiUpdateJobTagRequest) Request(request SchemasUpdateJobTagPayload) ApiUpdateJobTagRequest {
	r.request = &request
	return r
}

func (r ApiUpdateJobTagRequest) Execute() (*SchemasJobTagResponse, *http.Response, error) {
	return r.ApiService.UpdateJobTagExecute(r)
}

/*
UpdateJobTag Update jobTag by UUID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param jobTagId jobTag UUID
 @return ApiUpdateJobTagRequest
*/
func (a *JobTagAPIService) UpdateJobTag(ctx context.Context, jobTagId string) ApiUpdateJobTagRequest {
	return ApiUpdateJobTagRequest{
		ApiService: a,
		ctx: ctx,
		jobTagId: jobTagId,
	}
}

// Execute executes the request
//  @return SchemasJobTagResponse
func (a *JobTagAPIService) UpdateJobTagExecute(r ApiUpdateJobTagRequest) (*SchemasJobTagResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SchemasJobTagResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobTagAPIService.UpdateJobTag")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/jobTags/{jobTagId}"
	localVarPath = strings.Replace(localVarPath, "{"+"jobTagId"+"}", url.PathEscape(parameterValueToString(r.jobTagId, "jobTagId")), -1)

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
//  @return SchemasJobTagResponseStream
func (a *JobTagAPIService) UpdateJobTagStreamExecute(r ApiUpdateJobTagRequest, request *http.Request, GoogleAccessToken string) (*SchemasJobTagResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarReturnValue  *SchemasJobTagResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobTagAPIService.UpdateJobTag")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/jobTags/{jobTagId}"
	localVarPath = strings.Replace(localVarPath, "{"+"jobTagId"+"}", url.PathEscape(parameterValueToString(r.jobTagId, "jobTagId")), -1)

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
