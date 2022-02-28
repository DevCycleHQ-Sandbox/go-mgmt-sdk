# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectsControllerCreate**](ProjectsApi.md#ProjectsControllerCreate) | **Post** /v1/projects | Create Project
[**ProjectsControllerFindAll**](ProjectsApi.md#ProjectsControllerFindAll) | **Get** /v1/projects | List Projects
[**ProjectsControllerFindOne**](ProjectsApi.md#ProjectsControllerFindOne) | **Get** /v1/projects/{key} | Get a Project
[**ProjectsControllerRemove**](ProjectsApi.md#ProjectsControllerRemove) | **Delete** /v1/projects/{key} | Delete a Project
[**ProjectsControllerUpdate**](ProjectsApi.md#ProjectsControllerUpdate) | **Patch** /v1/projects/{key} | Update a Project

# **ProjectsControllerCreate**
> Project ProjectsControllerCreate(ctx, body)
Create Project

Create a new Project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateProjectDto**](CreateProjectDto.md)|  | 

### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsControllerFindAll**
> []Project ProjectsControllerFindAll(ctx, optional)
List Projects

List Projects

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProjectsApiProjectsControllerFindAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiProjectsControllerFindAllOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Float64**|  | [default to 1]
 **perPage** | **optional.Float64**|  | [default to 100]
 **sortBy** | **optional.String**|  | [default to createdAt]
 **sortOrder** | **optional.String**|  | [default to desc]
 **search** | **optional.String**|  | 
 **createdBy** | **optional.String**|  | 

### Return type

[**[]Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsControllerFindOne**
> Project ProjectsControllerFindOne(ctx, key)
Get a Project

Get a Project by ID or key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| A Project key or ID | 

### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsControllerRemove**
> ProjectsControllerRemove(ctx, key)
Delete a Project

Delete a Project by ID or key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| A Project key or ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsControllerUpdate**
> Project ProjectsControllerUpdate(ctx, body, key)
Update a Project

Update a Project by ID or key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateProjectDto**](UpdateProjectDto.md)|  | 
  **key** | **string**| A Project key or ID | 

### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

