# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FeatureConfigsControllerFindAll**](FeaturesApi.md#FeatureConfigsControllerFindAll) | **Get** /v1/projects/{project}/features/{feature}/configurations | List Feature configurations
[**FeatureConfigsControllerUpdate**](FeaturesApi.md#FeatureConfigsControllerUpdate) | **Patch** /v1/projects/{project}/features/{feature}/configurations | Update a Feature configuration
[**FeaturesControllerCreate**](FeaturesApi.md#FeaturesControllerCreate) | **Post** /v1/projects/{project}/features | Create Feature
[**FeaturesControllerFindAll**](FeaturesApi.md#FeaturesControllerFindAll) | **Get** /v1/projects/{project}/features | List Features
[**FeaturesControllerFindOne**](FeaturesApi.md#FeaturesControllerFindOne) | **Get** /v1/projects/{project}/features/{key} | Get a Feature
[**FeaturesControllerRemove**](FeaturesApi.md#FeaturesControllerRemove) | **Delete** /v1/projects/{project}/features/{key} | Delete a Feature
[**FeaturesControllerUpdate**](FeaturesApi.md#FeaturesControllerUpdate) | **Patch** /v1/projects/{project}/features/{key} | Update a Feature
[**JiraFeaturesControllerCreate**](FeaturesApi.md#JiraFeaturesControllerCreate) | **Post** /v1/projects/{project}/features/{key}/integrations/jira/issues | Link feature to Jira issue
[**JiraFeaturesControllerFindAll**](FeaturesApi.md#JiraFeaturesControllerFindAll) | **Get** /v1/projects/{project}/features/{key}/integrations/jira/issues | List linked Jira Issues
[**JiraFeaturesControllerRemove**](FeaturesApi.md#JiraFeaturesControllerRemove) | **Delete** /v1/projects/{project}/features/{key}/integrations/jira/issues/{issue_id} | Unlink feature from Jira issue

# **FeatureConfigsControllerFindAll**
> []FeatureConfig FeatureConfigsControllerFindAll(ctx, feature, project, optional)
List Feature configurations

List Feature configurations for all environments or by environment key or ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **feature** | [**Object**](.md)| A Feature key or ID | 
  **project** | [**Object**](.md)| A Project key or ID | 
 **optional** | ***FeaturesApiFeatureConfigsControllerFindAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeaturesApiFeatureConfigsControllerFindAllOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **environment** | [**optional.Interface of Object**](.md)| A Environment key or ID | 

### Return type

[**[]FeatureConfig**](FeatureConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FeatureConfigsControllerUpdate**
> FeatureConfig FeatureConfigsControllerUpdate(ctx, body, environment, feature, project)
Update a Feature configuration

Update a Feature configuration by environment key or ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateFeatureConfigDto**](UpdateFeatureConfigDto.md)|  | 
  **environment** | [**Object**](.md)| A Environment key or ID | 
  **feature** | [**Object**](.md)| A Feature key or ID | 
  **project** | [**Object**](.md)| A Project key or ID | 

### Return type

[**FeatureConfig**](FeatureConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FeaturesControllerCreate**
> Feature FeaturesControllerCreate(ctx, body, project)
Create Feature

Create a new Feature

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateFeatureDto**](CreateFeatureDto.md)|  | 
  **project** | [**Object**](.md)| A Project key or ID | 

### Return type

[**Feature**](Feature.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FeaturesControllerFindAll**
> []Feature FeaturesControllerFindAll(ctx, project, optional)
List Features

List Features

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **project** | [**Object**](.md)| A Project key or ID | 
 **optional** | ***FeaturesApiFeaturesControllerFindAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeaturesApiFeaturesControllerFindAllOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Float64**|  | [default to 1]
 **perPage** | **optional.Float64**|  | [default to 100]
 **sortBy** | **optional.String**|  | [default to createdAt]
 **sortOrder** | **optional.String**|  | [default to desc]
 **search** | **optional.String**|  | 
 **createdBy** | **optional.String**|  | 
 **type_** | **optional.String**|  | 

### Return type

[**[]Feature**](Feature.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FeaturesControllerFindOne**
> Feature FeaturesControllerFindOne(ctx, key, project)
Get a Feature

Get a Feature by ID or key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| A Feature key or ID | 
  **project** | [**Object**](.md)| A Project key or ID | 

### Return type

[**Feature**](Feature.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FeaturesControllerRemove**
> FeaturesControllerRemove(ctx, key, project)
Delete a Feature

Delete a Feature by ID or key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| A Feature key or ID | 
  **project** | [**Object**](.md)| A Project key or ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FeaturesControllerUpdate**
> Feature FeaturesControllerUpdate(ctx, body, key, project)
Update a Feature

Update a Feature by ID or key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateFeatureDto**](UpdateFeatureDto.md)|  | 
  **key** | **string**| A Feature key or ID | 
  **project** | [**Object**](.md)| A Project key or ID | 

### Return type

[**Feature**](Feature.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JiraFeaturesControllerCreate**
> JiraIssueLink JiraFeaturesControllerCreate(ctx, body, key, project)
Link feature to Jira issue

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LinkJiraIssueDto**](LinkJiraIssueDto.md)|  | 
  **key** | **string**|  | 
  **project** | [**Object**](.md)| A Project key or ID | 

### Return type

[**JiraIssueLink**](JiraIssueLink.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JiraFeaturesControllerFindAll**
> []JiraIssueLink JiraFeaturesControllerFindAll(ctx, key, project)
List linked Jira Issues

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**|  | 
  **project** | [**Object**](.md)| A Project key or ID | 

### Return type

[**[]JiraIssueLink**](JiraIssueLink.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JiraFeaturesControllerRemove**
> JiraFeaturesControllerRemove(ctx, key, issueId, project)
Unlink feature from Jira issue

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| A Feature key or ID | 
  **issueId** | **string**|  | 
  **project** | [**Object**](.md)| A Project key or ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

