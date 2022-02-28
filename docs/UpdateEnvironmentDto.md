# UpdateEnvironmentDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A unique display name | [optional] [default to null]
**Key** | **string** | Unique Environment identifier, can be used in the SDK / API to reference by key rather than ID. Must only contain lower-case characters and &#x60;_&#x60; or &#x60;-&#x60;. | [optional] [default to null]
**Description** | **string** | Environment description. | [optional] [default to null]
**Color** | **string** | Environment display color, used to highlight different environments on the dashboard. Must use Hex color code. | [optional] [default to null]
**Type_** | **string** | The environment type | [optional] [default to null]
**Settings** | [***AllOfUpdateEnvironmentDtoSettings**](AllOfUpdateEnvironmentDtoSettings.md) | Environment based settings | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

