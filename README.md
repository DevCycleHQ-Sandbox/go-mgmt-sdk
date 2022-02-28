# Go API for DevCycle Management Api

No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to */*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CustomPropertiesApi* | [**CustomPropertiesControllerCreate**](docs/CustomPropertiesApi.md#custompropertiescontrollercreate) | **Post** /v1/projects/{project}/customProperties | Create Custom Property
*CustomPropertiesApi* | [**CustomPropertiesControllerFindAll**](docs/CustomPropertiesApi.md#custompropertiescontrollerfindall) | **Get** /v1/projects/{project}/customProperties | List Custom Properties
*EnvironmentsApi* | [**EnvironmentsControllerCreate**](docs/EnvironmentsApi.md#environmentscontrollercreate) | **Post** /v1/projects/{project}/environments | Create Environment
*EnvironmentsApi* | [**EnvironmentsControllerFindAll**](docs/EnvironmentsApi.md#environmentscontrollerfindall) | **Get** /v1/projects/{project}/environments | List Environments
*EnvironmentsApi* | [**EnvironmentsControllerFindOne**](docs/EnvironmentsApi.md#environmentscontrollerfindone) | **Get** /v1/projects/{project}/environments/{key} | Get an Environment
*EnvironmentsApi* | [**EnvironmentsControllerRemove**](docs/EnvironmentsApi.md#environmentscontrollerremove) | **Delete** /v1/projects/{project}/environments/{key} | Delete an Environment
*EnvironmentsApi* | [**EnvironmentsControllerUpdate**](docs/EnvironmentsApi.md#environmentscontrollerupdate) | **Patch** /v1/projects/{project}/environments/{key} | Update an Environment
*FeaturesApi* | [**FeatureConfigsControllerFindAll**](docs/FeaturesApi.md#featureconfigscontrollerfindall) | **Get** /v1/projects/{project}/features/{feature}/configurations | List Feature configurations
*FeaturesApi* | [**FeatureConfigsControllerUpdate**](docs/FeaturesApi.md#featureconfigscontrollerupdate) | **Patch** /v1/projects/{project}/features/{feature}/configurations | Update a Feature configuration
*FeaturesApi* | [**FeaturesControllerCreate**](docs/FeaturesApi.md#featurescontrollercreate) | **Post** /v1/projects/{project}/features | Create Feature
*FeaturesApi* | [**FeaturesControllerFindAll**](docs/FeaturesApi.md#featurescontrollerfindall) | **Get** /v1/projects/{project}/features | List Features
*FeaturesApi* | [**FeaturesControllerFindOne**](docs/FeaturesApi.md#featurescontrollerfindone) | **Get** /v1/projects/{project}/features/{key} | Get a Feature
*FeaturesApi* | [**FeaturesControllerRemove**](docs/FeaturesApi.md#featurescontrollerremove) | **Delete** /v1/projects/{project}/features/{key} | Delete a Feature
*FeaturesApi* | [**FeaturesControllerUpdate**](docs/FeaturesApi.md#featurescontrollerupdate) | **Patch** /v1/projects/{project}/features/{key} | Update a Feature
*FeaturesApi* | [**JiraFeaturesControllerCreate**](docs/FeaturesApi.md#jirafeaturescontrollercreate) | **Post** /v1/projects/{project}/features/{key}/integrations/jira/issues | Link feature to Jira issue
*FeaturesApi* | [**JiraFeaturesControllerFindAll**](docs/FeaturesApi.md#jirafeaturescontrollerfindall) | **Get** /v1/projects/{project}/features/{key}/integrations/jira/issues | List linked Jira Issues
*FeaturesApi* | [**JiraFeaturesControllerRemove**](docs/FeaturesApi.md#jirafeaturescontrollerremove) | **Delete** /v1/projects/{project}/features/{key}/integrations/jira/issues/{issue_id} | Unlink feature from Jira issue
*JiraIntegrationApi* | [**JiraFeaturesControllerCreate**](docs/JiraIntegrationApi.md#jirafeaturescontrollercreate) | **Post** /v1/projects/{project}/features/{key}/integrations/jira/issues | Link feature to Jira issue
*JiraIntegrationApi* | [**JiraFeaturesControllerFindAll**](docs/JiraIntegrationApi.md#jirafeaturescontrollerfindall) | **Get** /v1/projects/{project}/features/{key}/integrations/jira/issues | List linked Jira Issues
*JiraIntegrationApi* | [**JiraFeaturesControllerRemove**](docs/JiraIntegrationApi.md#jirafeaturescontrollerremove) | **Delete** /v1/projects/{project}/features/{key}/integrations/jira/issues/{issue_id} | Unlink feature from Jira issue
*ProjectsApi* | [**ProjectsControllerCreate**](docs/ProjectsApi.md#projectscontrollercreate) | **Post** /v1/projects | Create Project
*ProjectsApi* | [**ProjectsControllerFindAll**](docs/ProjectsApi.md#projectscontrollerfindall) | **Get** /v1/projects | List Projects
*ProjectsApi* | [**ProjectsControllerFindOne**](docs/ProjectsApi.md#projectscontrollerfindone) | **Get** /v1/projects/{key} | Get a Project
*ProjectsApi* | [**ProjectsControllerRemove**](docs/ProjectsApi.md#projectscontrollerremove) | **Delete** /v1/projects/{key} | Delete a Project
*ProjectsApi* | [**ProjectsControllerUpdate**](docs/ProjectsApi.md#projectscontrollerupdate) | **Patch** /v1/projects/{key} | Update a Project
*ResultsApi* | [**ResultsControllerGetEvaluationsPerHourByFeature**](docs/ResultsApi.md#resultscontrollergetevaluationsperhourbyfeature) | **Get** /v1/projects/{project}/features/{feature}/results/evaluations | 
*ResultsApi* | [**ResultsControllerGetEvaluationsPerHourByProject**](docs/ResultsApi.md#resultscontrollergetevaluationsperhourbyproject) | **Get** /v1/projects/{project}/results/evaluations | 
*ResultsApi* | [**ResultsControllerGetResultsSummary**](docs/ResultsApi.md#resultscontrollergetresultssummary) | **Get** /v1/projects/{project}/features/{feature}/results/summary | 
*VariablesApi* | [**VariablesControllerCreate**](docs/VariablesApi.md#variablescontrollercreate) | **Post** /v1/projects/{project}/variables | Create a Variable
*VariablesApi* | [**VariablesControllerFindAll**](docs/VariablesApi.md#variablescontrollerfindall) | **Get** /v1/projects/{project}/variables | List Variables
*VariablesApi* | [**VariablesControllerFindOne**](docs/VariablesApi.md#variablescontrollerfindone) | **Get** /v1/projects/{project}/variables/{key} | Get a Variable
*VariablesApi* | [**VariablesControllerRemove**](docs/VariablesApi.md#variablescontrollerremove) | **Delete** /v1/projects/{project}/variables/{key} | Delete a Variable
*VariablesApi* | [**VariablesControllerUpdate**](docs/VariablesApi.md#variablescontrollerupdate) | **Patch** /v1/projects/{project}/variables/{key} | Update a Variable

## Documentation For Models

 - [AllFilter](docs/AllFilter.md)
 - [AllOfCreateEnvironmentDtoSettings](docs/AllOfCreateEnvironmentDtoSettings.md)
 - [AllOfEnvironmentSdkKeys](docs/AllOfEnvironmentSdkKeys.md)
 - [AllOfEnvironmentSettings](docs/AllOfEnvironmentSettings.md)
 - [AllOfTargetAudience](docs/AllOfTargetAudience.md)
 - [AllOfTargetAudienceFilters](docs/AllOfTargetAudienceFilters.md)
 - [AllOfTargetRollout](docs/AllOfTargetRollout.md)
 - [AllOfUpdateEnvironmentDtoSettings](docs/AllOfUpdateEnvironmentDtoSettings.md)
 - [AllOfUpdateTargetDtoAudience](docs/AllOfUpdateTargetDtoAudience.md)
 - [AllOfUpdateTargetDtoRollout](docs/AllOfUpdateTargetDtoRollout.md)
 - [AnyOfAudienceOperatorFiltersItems](docs/AnyOfAudienceOperatorFiltersItems.md)
 - [ApiKey](docs/ApiKey.md)
 - [AudienceOperator](docs/AudienceOperator.md)
 - [BadRequestErrorResponse](docs/BadRequestErrorResponse.md)
 - [CannotDeleteLastItemErrorResponse](docs/CannotDeleteLastItemErrorResponse.md)
 - [ConflictErrorResponse](docs/ConflictErrorResponse.md)
 - [CreateCustomPropertyDto](docs/CreateCustomPropertyDto.md)
 - [CreateEnvironmentDto](docs/CreateEnvironmentDto.md)
 - [CreateFeatureDto](docs/CreateFeatureDto.md)
 - [CreateProjectDto](docs/CreateProjectDto.md)
 - [CreateVariableDto](docs/CreateVariableDto.md)
 - [CustomProperty](docs/CustomProperty.md)
 - [Environment](docs/Environment.md)
 - [EnvironmentSettings](docs/EnvironmentSettings.md)
 - [Feature](docs/Feature.md)
 - [FeatureConfig](docs/FeatureConfig.md)
 - [FeatureDataPoint](docs/FeatureDataPoint.md)
 - [FeatureVariationDto](docs/FeatureVariationDto.md)
 - [JiraIssueLink](docs/JiraIssueLink.md)
 - [LinkJiraIssueDto](docs/LinkJiraIssueDto.md)
 - [NotFoundErrorResponse](docs/NotFoundErrorResponse.md)
 - [PreconditionFailedErrorResponse](docs/PreconditionFailedErrorResponse.md)
 - [Project](docs/Project.md)
 - [ProjectDataPoint](docs/ProjectDataPoint.md)
 - [ResultEvaluationsByHourDto](docs/ResultEvaluationsByHourDto.md)
 - [ResultEvaluationsByHourDtoResult](docs/ResultEvaluationsByHourDtoResult.md)
 - [ResultProjectEvaluationsByHourDto](docs/ResultProjectEvaluationsByHourDto.md)
 - [ResultProjectEvaluationsByHourDtoResult](docs/ResultProjectEvaluationsByHourDtoResult.md)
 - [ResultSummaryDto](docs/ResultSummaryDto.md)
 - [ResultSummaryDtoResult](docs/ResultSummaryDtoResult.md)
 - [ResultSummaryDtoResultCounts](docs/ResultSummaryDtoResultCounts.md)
 - [Rollout](docs/Rollout.md)
 - [RolloutStage](docs/RolloutStage.md)
 - [SdkKeys](docs/SdkKeys.md)
 - [Target](docs/Target.md)
 - [TargetAudience](docs/TargetAudience.md)
 - [TargetDistribution](docs/TargetDistribution.md)
 - [UpdateEnvironmentDto](docs/UpdateEnvironmentDto.md)
 - [UpdateFeatureConfigDto](docs/UpdateFeatureConfigDto.md)
 - [UpdateFeatureDto](docs/UpdateFeatureDto.md)
 - [UpdateProjectDto](docs/UpdateProjectDto.md)
 - [UpdateTargetDto](docs/UpdateTargetDto.md)
 - [UpdateVariableDto](docs/UpdateVariableDto.md)
 - [UserCountryFilter](docs/UserCountryFilter.md)
 - [UserCustomFilter](docs/UserCustomFilter.md)
 - [UserFilter](docs/UserFilter.md)
 - [Variable](docs/Variable.md)
 - [Variation](docs/Variation.md)

## Documentation For Authorization
 Endpoints do not require authorization.


## Author

