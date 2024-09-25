/*
Masters way chat API

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
	"os"
)


// FileAPIService FileAPI service
type FileAPIService service

type ApiDeleteFilesRequest struct {
	ctx context.Context
	ApiService *FileAPIService
	fileIDs *[]string
}

// List of file IDs to delete
func (r ApiDeleteFilesRequest) FileIDs(fileIDs []string) ApiDeleteFilesRequest {
	r.fileIDs = &fileIDs
	return r
}

func (r ApiDeleteFilesRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteFilesExecute(r)
}

/*
DeleteFiles Delete files by IDs

Delete multiple files from the server storage using their IDs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeleteFilesRequest
*/
func (a *FileAPIService) DeleteFiles(ctx context.Context) ApiDeleteFilesRequest {
	return ApiDeleteFilesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *FileAPIService) DeleteFilesExecute(r ApiDeleteFilesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FileAPIService.DeleteFiles")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/files"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.fileIDs == nil {
		return nil, reportError("fileIDs is required and must be specified")
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
	localVarPostBody = r.fileIDs
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
func (a *FileAPIService) DeleteFilesStreamExecute(r ApiDeleteFilesRequest, request *http.Request, GoogleAccessToken string) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FileAPIService.DeleteFiles")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/files"

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

type ApiUploadFileRequest struct {
	ctx context.Context
	ApiService *FileAPIService
	roomId *string
	file *os.File
}

// Room id
func (r ApiUploadFileRequest) RoomId(roomId string) ApiUploadFileRequest {
	r.roomId = &roomId
	return r
}

// File to upload
func (r ApiUploadFileRequest) File(file *os.File) ApiUploadFileRequest {
	r.file = file
	return r
}

func (r ApiUploadFileRequest) Execute() (*SchemasUploadFileResponse, *http.Response, error) {
	return r.ApiService.UploadFileExecute(r)
}

/*
UploadFile Upload file to storage

Uploads a file to the server and stores it in the designated storage path

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUploadFileRequest
*/
func (a *FileAPIService) UploadFile(ctx context.Context) ApiUploadFileRequest {
	return ApiUploadFileRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SchemasUploadFileResponse
func (a *FileAPIService) UploadFileExecute(r ApiUploadFileRequest) (*SchemasUploadFileResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SchemasUploadFileResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FileAPIService.UploadFile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/files"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.roomId == nil {
		return localVarReturnValue, nil, reportError("roomId is required and must be specified")
	}
	if r.file == nil {
		return localVarReturnValue, nil, reportError("file is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "roomId", r.roomId, "", "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"
	fileLocalVarFile := r.file

	if fileLocalVarFile != nil {
		fbs, _ := io.ReadAll(fileLocalVarFile)

		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
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

// Execute executes the request
//  @return SchemasUploadFileResponseStream
func (a *FileAPIService) UploadFileStreamExecute(r ApiUploadFileRequest, request *http.Request, GoogleAccessToken string) (*SchemasUploadFileResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarReturnValue  *SchemasUploadFileResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FileAPIService.UploadFile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/files"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	fmt.Println(localVarQueryParams)



		parameterAddToHeaderOrQuery(localVarQueryParams, "roomId", r.roomId, "", "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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