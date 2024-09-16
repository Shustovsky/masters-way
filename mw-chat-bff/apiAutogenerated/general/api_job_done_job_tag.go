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


// JobDoneJobTagAPIService JobDoneJobTagAPI service
type JobDoneJobTagAPIService service

type ApiCreateJobDoneJobTagRequest struct {
	ctx context.Context
	ApiService *JobDoneJobTagAPIService
	request *SchemasCreateJobDoneJobTagPayload
}

// query params
func (r ApiCreateJobDoneJobTagRequest) Request(request SchemasCreateJobDoneJobTagPayload) ApiCreateJobDoneJobTagRequest {
	r.request = &request
	return r
}

func (r ApiCreateJobDoneJobTagRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateJobDoneJobTagExecute(r)
}

/*
CreateJobDoneJobTag Create a new jobDoneJobTag

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateJobDoneJobTagRequest
*/
func (a *JobDoneJobTagAPIService) CreateJobDoneJobTag(ctx context.Context) ApiCreateJobDoneJobTagRequest {
	return ApiCreateJobDoneJobTagRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *JobDoneJobTagAPIService) CreateJobDoneJobTagExecute(r ApiCreateJobDoneJobTagRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobDoneJobTagAPIService.CreateJobDoneJobTag")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/jobDoneJobTags"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.request == nil {
		return nil, reportError("request is required and must be specified")
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
	// body params
	localVarPostBody = r.request
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
func (a *JobDoneJobTagAPIService) CreateJobDoneJobTagStreamExecute(r ApiCreateJobDoneJobTagRequest, request *http.Request, GoogleAccessToken string) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobDoneJobTagAPIService.CreateJobDoneJobTag")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/jobDoneJobTags"

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

type ApiDeleteJobDoneJobTagRequest struct {
	ctx context.Context
	ApiService *JobDoneJobTagAPIService
	jobDoneId string
	jobTagId string
}

func (r ApiDeleteJobDoneJobTagRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteJobDoneJobTagExecute(r)
}

/*
DeleteJobDoneJobTag Delete jobDoneJobTag by UUID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param jobDoneId jobDone ID
 @param jobTagId jobTag UUID
 @return ApiDeleteJobDoneJobTagRequest
*/
func (a *JobDoneJobTagAPIService) DeleteJobDoneJobTag(ctx context.Context, jobDoneId string, jobTagId string) ApiDeleteJobDoneJobTagRequest {
	return ApiDeleteJobDoneJobTagRequest{
		ApiService: a,
		ctx: ctx,
		jobDoneId: jobDoneId,
		jobTagId: jobTagId,
	}
}

// Execute executes the request
func (a *JobDoneJobTagAPIService) DeleteJobDoneJobTagExecute(r ApiDeleteJobDoneJobTagRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobDoneJobTagAPIService.DeleteJobDoneJobTag")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/jobDoneJobTags/{jobTagId}/{jobDoneId}"
	localVarPath = strings.Replace(localVarPath, "{"+"jobDoneId"+"}", url.PathEscape(parameterValueToString(r.jobDoneId, "jobDoneId")), -1)
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
func (a *JobDoneJobTagAPIService) DeleteJobDoneJobTagStreamExecute(r ApiDeleteJobDoneJobTagRequest, request *http.Request, GoogleAccessToken string) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobDoneJobTagAPIService.DeleteJobDoneJobTag")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/jobDoneJobTags/{jobTagId}/{jobDoneId}"
	localVarPath = strings.Replace(localVarPath, "{"+"jobDoneId"+"}", url.PathEscape(parameterValueToString(r.jobDoneId, "jobDoneId")), -1)
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
