# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResultsControllerGetEvaluationsPerHourByFeature**](ResultsApi.md#ResultsControllerGetEvaluationsPerHourByFeature) | **Get** /v1/projects/{project}/features/{feature}/results/evaluations | 
[**ResultsControllerGetEvaluationsPerHourByProject**](ResultsApi.md#ResultsControllerGetEvaluationsPerHourByProject) | **Get** /v1/projects/{project}/results/evaluations | 
[**ResultsControllerGetResultsSummary**](ResultsApi.md#ResultsControllerGetResultsSummary) | **Get** /v1/projects/{project}/features/{feature}/results/summary | 

# **ResultsControllerGetEvaluationsPerHourByFeature**
> ResultEvaluationsByHourDto ResultsControllerGetEvaluationsPerHourByFeature(ctx, feature, project, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **feature** | [**Object**](.md)| A Feature key or ID | 
  **project** | [**Object**](.md)| A Project key or ID | 
 **optional** | ***ResultsApiResultsControllerGetEvaluationsPerHourByFeatureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ResultsApiResultsControllerGetEvaluationsPerHourByFeatureOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startDate** | **optional.Float64**|  | 
 **endDate** | **optional.Float64**|  | 
 **platform** | **optional.String**|  | 
 **variable** | **optional.String**|  | 
 **environment** | **optional.String**| A Environment key or ID | 
 **period** | **optional.String**|  | 

### Return type

[**ResultEvaluationsByHourDto**](ResultEvaluationsByHourDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResultsControllerGetEvaluationsPerHourByProject**
> ResultProjectEvaluationsByHourDto ResultsControllerGetEvaluationsPerHourByProject(ctx, project, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **project** | [**Object**](.md)| A Project key or ID | 
 **optional** | ***ResultsApiResultsControllerGetEvaluationsPerHourByProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ResultsApiResultsControllerGetEvaluationsPerHourByProjectOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **optional.Float64**|  | 
 **endDate** | **optional.Float64**|  | 
 **environment** | **optional.String**| A Environment key or ID | 
 **period** | **optional.String**|  | 

### Return type

[**ResultProjectEvaluationsByHourDto**](ResultProjectEvaluationsByHourDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResultsControllerGetResultsSummary**
> ResultSummaryDto ResultsControllerGetResultsSummary(ctx, feature, project, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **feature** | [**Object**](.md)| A Feature key or ID | 
  **project** | [**Object**](.md)| A Project key or ID | 
 **optional** | ***ResultsApiResultsControllerGetResultsSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ResultsApiResultsControllerGetResultsSummaryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startDate** | **optional.Float64**|  | 
 **endDate** | **optional.Float64**|  | 
 **platform** | **optional.String**|  | 
 **variable** | **optional.String**|  | 
 **environment** | **optional.String**| A Environment key or ID | 
 **period** | **optional.String**|  | 

### Return type

[**ResultSummaryDto**](ResultSummaryDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

