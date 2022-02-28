# CustomProperty

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name for Custom Property. | [default to null]
**Key** | **string** | Auto generated key to be used by the API to reference by &#x27;key&#x27; rather then _id for CRUD operations. Must only contain lower-case characters and &#x60;_&#x60; or &#x60;-&#x60;. | [default to null]
**Id** | **string** | A unique CustomProperty ID | [default to null]
**Project** | **string** | The Project owning the Custom Property | [default to null]
**CreatedBy** | **string** | ID of the User who created the project | [default to null]
**PropertyKey** | **string** | Custom Property key, must be unique by Project. Only to be used by the SDKs to reference the Custom Property.. | [default to null]
**Type_** | **string** | The Custom Property type | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The date the Project was created | [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The date the Project was last updated | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

