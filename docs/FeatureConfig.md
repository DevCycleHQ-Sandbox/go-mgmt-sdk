# FeatureConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Feature** | **string** | ID of the Feature owning the Configuration | [default to null]
**Environment** | **string** | ID of the Environment owning the Configuration | [default to null]
**CreatedBy** | **string** | User who created the Feature Configuration | [default to null]
**Status** | **string** | Status of the Feature Configuration | [default to null]
**StartedAt** | [**time.Time**](time.Time.md) | Date the Feature Configuration was started | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The date the Feature Configuration was last updated | [default to null]
**Targets** | [**[]Target**](Target.md) | The targets to evaluate what variation a user should be delivered | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

