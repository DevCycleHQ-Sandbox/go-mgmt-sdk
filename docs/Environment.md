# Environment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A unique display name | [default to null]
**Key** | **string** | Unique Environment identifier, can be used in the SDK / API to reference by key rather than ID. Must only contain lower-case characters and &#x60;_&#x60; or &#x60;-&#x60;. | [default to null]
**Description** | **string** | Environment description. | [optional] [default to null]
**Color** | **string** | Environment display color, used to highlight different environments on the dashboard. Must use Hex color code. | [optional] [default to null]
**Id** | **string** | A unique Environment ID | [default to null]
**Project** | **string** | The Project owning the Environment | [default to null]
**Type_** | **string** | The environment type | [default to null]
**CreatedBy** | **string** | ID of the User who created the Environment. | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The date the Environment was created | [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The date the Environment was last updated | [default to null]
**SdkKeys** | [***AllOfEnvironmentSdkKeys**](AllOfEnvironmentSdkKeys.md) | SDK Keys for mobile SDKs / client SDKs / server SDKs. Multiple keys can be created to allow for key rotation. | [default to null]
**Settings** | [***AllOfEnvironmentSettings**](AllOfEnvironmentSettings.md) | Environment based settings | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

