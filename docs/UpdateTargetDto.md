# UpdateTargetDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | [***AllOfUpdateTargetDtoAudience**](AllOfUpdateTargetDtoAudience.md) | Audience object describing target segmentation. | [default to null]
**Name** | **string** | Target name | [optional] [default to null]
**Rollout** | [***AllOfUpdateTargetDtoRollout**](AllOfUpdateTargetDtoRollout.md) | Rollout sub-document describing how a Target&#x27;s audience is rolled out | [optional] [default to null]
**Distribution** | [**[]TargetDistribution**](TargetDistribution.md) | Specifies variation distribution percentages for features | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

