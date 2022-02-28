# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnvironmentsControllerCreate**](EnvironmentsApi.md#EnvironmentsControllerCreate) | **Post** /v1/projects/{project}/environments | Create Environment
[**EnvironmentsControllerFindAll**](EnvironmentsApi.md#EnvironmentsControllerFindAll) | **Get** /v1/projects/{project}/environments | List Environments
[**EnvironmentsControllerFindOne**](EnvironmentsApi.md#EnvironmentsControllerFindOne) | **Get** /v1/projects/{project}/environments/{key} | Get an Environment
[**EnvironmentsControllerRemove**](EnvironmentsApi.md#EnvironmentsControllerRemove) | **Delete** /v1/projects/{project}/environments/{key} | Delete an Environment
[**EnvironmentsControllerUpdate**](EnvironmentsApi.md#EnvironmentsControllerUpdate) | **Patch** /v1/projects/{project}/environments/{key} | Update an Environment

# **EnvironmentsControllerCreate**
> Environment EnvironmentsControllerCreate(ctx, body, project)
Create Environment

Create a new Environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateEnvironmentDto**](CreateEnvironmentDto.md)|  | 
  **project** | [**Object**](.md)| A Project key or ID | 

### Return type

[**Environment**](Environment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnvironmentsControllerFindAll**
> []Environment EnvironmentsControllerFindAll(ctx, project, optional)
List Environments

List Environments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **project** | [**Object**](.md)| A Project key or ID | 
 **optional** | ***EnvironmentsApiEnvironmentsControllerFindAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnvironmentsApiEnvironmentsControllerFindAllOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Float64**|  | [default to 1]
 **perPage** | **optional.Float64**|  | [default to 100]
 **sortBy** | **optional.String**|  | [default to createdAt]
 **sortOrder** | **optional.String**|  | [default to desc]
 **search** | **optional.String**|  | 
 **createdBy** | **optional.String**|  | 

### Return type

[**[]Environment**](Environment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnvironmentsControllerFindOne**
> Environment EnvironmentsControllerFindOne(ctx, key, project)
Get an Environment

Get an Environment by ID or key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| A Environment key or ID | 
  **project** | [**Object**](.md)| A Project key or ID | 

### Return type

[**Environment**](Environment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnvironmentsControllerRemove**
> EnvironmentsControllerRemove(ctx, key, project)
Delete an Environment

Delete an Environment by ID or key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| A Environment key or ID | 
  **project** | [**Object**](.md)| A Project key or ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnvironmentsControllerUpdate**
> Environment EnvironmentsControllerUpdate(ctx, body, key, project)
Update an Environment

Update an Environment by ID or key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateEnvironmentDto**](UpdateEnvironmentDto.md)|  | 
  **key** | **string**| A Environment key or ID | 
  **project** | [**Object**](.md)| A Project key or ID | 

### Return type

[**Environment**](Environment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

