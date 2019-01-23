/*
 * Strava API v3
 *
 * Strava API
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type UploadsApiService service

/* UploadsApiService Upload Activity
Uploads a new data file to create an activity from. Requires activity:write scope.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "file" (*os.File) The uploaded file.
    @param "name" (string) The desired name of the resulting activity.
    @param "description" (string) The desired description of the resulting activity.
    @param "trainer" (string) Whether the resulting activity should be marked as having been performed on a trainer.
    @param "commute" (string) Whether the resulting activity should be tagged as a commute.
    @param "dataType" (string) The format of the uploaded file.
    @param "externalId" (string) The desired external identifier of the resulting activity.
@return Upload*/
func (a *UploadsApiService) CreateUpload(ctx context.Context, localVarOptionals map[string]interface{}) (Upload, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     Upload
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/uploads"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["name"], "string", "name"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["description"], "string", "description"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["trainer"], "string", "trainer"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["commute"], "string", "commute"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["dataType"], "string", "dataType"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["externalId"], "string", "externalId"); err != nil {
		return successPayload, nil, err
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var localVarFile (*os.File)
	if localVarTempParam, localVarOk := localVarOptionals["file"].(*os.File); localVarOk {
		localVarFile = localVarTempParam
	}
	if localVarFile != nil {
		fbs, _ := ioutil.ReadAll(localVarFile)
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
		localVarFile.Close()
	}
	if localVarTempParam, localVarOk := localVarOptionals["name"].(string); localVarOk {
		localVarFormParams.Add("name", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["description"].(string); localVarOk {
		localVarFormParams.Add("description", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["trainer"].(string); localVarOk {
		localVarFormParams.Add("trainer", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["commute"].(string); localVarOk {
		localVarFormParams.Add("commute", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["dataType"].(string); localVarOk {
		localVarFormParams.Add("data_type", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["externalId"].(string); localVarOk {
		localVarFormParams.Add("external_id", parameterToString(localVarTempParam, ""))
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* UploadsApiService Get Upload
Returns an upload for a given identifier. Requires activity:write scope.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param uploadId The identifier of the upload.
@return Upload*/
func (a *UploadsApiService) GetUploadById(ctx context.Context, uploadId int64) (Upload, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     Upload
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/uploads/{uploadId}"
	localVarPath = strings.Replace(localVarPath, "{"+"uploadId"+"}", fmt.Sprintf("%v", uploadId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}
