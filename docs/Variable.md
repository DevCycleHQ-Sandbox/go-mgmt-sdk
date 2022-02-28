# Variable

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Variable name | [optional] [default to null]
**Description** | **string** | A description of the Variable | [optional] [default to null]
**Key** | **string** | Unique Variable identifier, can be used in the SDK / API to reference by key rather then ID. Must only contain lower-case characters and &#x60;_&#x60; or &#x60;-&#x60;. | [default to null]
**Id** | **string** | A unique Variable ID | [default to null]
**Project** | **string** | The ID of the Project this Variable belongs to | [default to null]
**Feature** | **string** | The ID of the Feature this Variable belongs to | [optional] [default to null]
**Type_** | **string** | The type of Variable. Must be one of [String | Boolean | Number | JSON] | [default to null]
**DefaultValue** | [***interface{}**](interface{}.md) | A default value for the Variable | [optional] [default to null]
**Source** | **string** | The system that was used for the creation of the Variable. | [default to null]
**CreatedBy** | **string** | ID of the User who created the Variable. | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The date the Variable was created | [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The date the Variable was last updated | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

