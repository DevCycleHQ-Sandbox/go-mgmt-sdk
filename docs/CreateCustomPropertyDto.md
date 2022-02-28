# CreateCustomPropertyDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name for Custom Property. | [default to null]
**Key** | **string** | Auto generated key to be used by the API to reference by &#x27;key&#x27; rather then _id for CRUD operations. Must only contain lower-case characters and &#x60;_&#x60; or &#x60;-&#x60;. | [default to null]
**Type_** | **string** | Type of the Custom Propety. Must be one of: \&quot;Boolean\&quot;, \&quot;Number\&quot;, \&quot;Semver\&quot; or \&quot;String\&quot; | [default to null]
**PropertyKey** | **string** | Custom Property key, must be unique by Project. Only to be used by the SDKs to reference the Custom Property.. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

