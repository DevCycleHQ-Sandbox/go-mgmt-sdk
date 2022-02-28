# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**JiraFeaturesControllerCreate**](JiraIntegrationApi.md#JiraFeaturesControllerCreate) | **Post** /v1/projects/{project}/features/{key}/integrations/jira/issues | Link feature to Jira issue
[**JiraFeaturesControllerFindAll**](JiraIntegrationApi.md#JiraFeaturesControllerFindAll) | **Get** /v1/projects/{project}/features/{key}/integrations/jira/issues | List linked Jira Issues
[**JiraFeaturesControllerRemove**](JiraIntegrationApi.md#JiraFeaturesControllerRemove) | **Delete** /v1/projects/{project}/features/{key}/integrations/jira/issues/{issue_id} | Unlink feature from Jira issue

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

