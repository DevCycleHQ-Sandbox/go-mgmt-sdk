# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VariablesControllerCreate**](VariablesApi.md#VariablesControllerCreate) | **Post** /v1/projects/{project}/variables | Create a Variable
[**VariablesControllerFindAll**](VariablesApi.md#VariablesControllerFindAll) | **Get** /v1/projects/{project}/variables | List Variables
[**VariablesControllerFindOne**](VariablesApi.md#VariablesControllerFindOne) | **Get** /v1/projects/{project}/variables/{key} | Get a Variable
[**VariablesControllerRemove**](VariablesApi.md#VariablesControllerRemove) | **Delete** /v1/projects/{project}/variables/{key} | Delete a Variable
[**VariablesControllerUpdate**](VariablesApi.md#VariablesControllerUpdate) | **Patch** /v1/projects/{project}/variables/{key} | Update a Variable

# **VariablesControllerCreate**
> Variable VariablesControllerCreate(ctx, body, project)
Create a Variable

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateVariableDto**](CreateVariableDto.md)|  | 
  **project** | [**Object**](.md)| A Project key or ID | 

### Return type

[**Variable**](Variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VariablesControllerFindAll**
> []Variable VariablesControllerFindAll(ctx, project, optional)
List Variables

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **project** | [**Object**](.md)| A Project key or ID | 
 **optional** | ***VariablesApiVariablesControllerFindAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VariablesApiVariablesControllerFindAllOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Float64**|  | [default to 1]
 **perPage** | **optional.Float64**|  | [default to 100]
 **sortBy** | **optional.String**|  | [default to createdAt]
 **sortOrder** | **optional.String**|  | [default to desc]
 **search** | **optional.String**|  | 
 **feature** | **optional.String**|  | 
 **type_** | **optional.String**|  | 

### Return type

[**[]Variable**](Variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VariablesControllerFindOne**
> Variable VariablesControllerFindOne(ctx, key, project)
Get a Variable

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| A Variable key or ID | 
  **project** | [**Object**](.md)| A Project key or ID | 

### Return type

[**Variable**](Variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VariablesControllerRemove**
> VariablesControllerRemove(ctx, key, project)
Delete a Variable

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| A Variable key or ID | 
  **project** | [**Object**](.md)| A Project key or ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VariablesControllerUpdate**
> Variable VariablesControllerUpdate(ctx, body, key, project)
Update a Variable

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateVariableDto**](UpdateVariableDto.md)|  | 
  **key** | **string**| A Variable key or ID | 
  **project** | [**Object**](.md)| A Project key or ID | 

### Return type

[**Variable**](Variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

