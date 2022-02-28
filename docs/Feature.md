# Feature

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the Feature | [default to null]
**Key** | **string** | Unique key by Project, can be used in the SDK / API to reference by &#x27;key&#x27; rather than _id. Must only contain lower-case characters and &#x60;_&#x60; or &#x60;-&#x60;. | [default to null]
**Description** | **string** | Feature description. | [optional] [default to null]
**Id** | **string** | A unique Feature ID | [default to null]
**Project** | **string** | The Project owning the Feature | [default to null]
**Source** | **string** | Source that created the Feature | [default to null]
**Type_** | **string** | The Feature type | [optional] [default to null]
**CreatedBy** | **string** | ID of the User who created the Feature | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The date the Feature was created | [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The date the Feature was last updated | [default to null]
**Variations** | [**[]Variation**](Variation.md) | Variation configurations to be used by feature configurations. | [optional] [default to null]
**Variables** | [**[]Variable**](Variable.md) | Variable definitions to be referenced in variations | [optional] [default to null]
**Tags** | **[]string** | Tags to organize Features on the dashboard | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

