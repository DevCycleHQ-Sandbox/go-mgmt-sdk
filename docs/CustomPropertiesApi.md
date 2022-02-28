# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomPropertiesControllerCreate**](CustomPropertiesApi.md#CustomPropertiesControllerCreate) | **Post** /v1/projects/{project}/customProperties | Create Custom Property
[**CustomPropertiesControllerFindAll**](CustomPropertiesApi.md#CustomPropertiesControllerFindAll) | **Get** /v1/projects/{project}/customProperties | List Custom Properties

# **CustomPropertiesControllerCreate**
> CustomProperty CustomPropertiesControllerCreate(ctx, body, project)
Create Custom Property

Create a new Custom Property

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateCustomPropertyDto**](CreateCustomPropertyDto.md)|  | 
  **project** | [**Object**](.md)| A Project key or ID | 

### Return type

[**CustomProperty**](CustomProperty.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomPropertiesControllerFindAll**
> []CustomProperty CustomPropertiesControllerFindAll(ctx, project, optional)
List Custom Properties

List Custom Properties

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **project** | [**Object**](.md)| A Project key or ID | 
 **optional** | ***CustomPropertiesApiCustomPropertiesControllerFindAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomPropertiesApiCustomPropertiesControllerFindAllOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Float64**|  | [default to 1]
 **perPage** | **optional.Float64**|  | [default to 100]
 **sortBy** | **optional.String**|  | [default to createdAt]
 **sortOrder** | **optional.String**|  | [default to desc]
 **search** | **optional.String**|  | 

### Return type

[**[]CustomProperty**](CustomProperty.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

