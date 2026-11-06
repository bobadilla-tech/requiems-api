# V1PlacesHolidaysBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **Array&lt;any&gt;** | One result per query, in the same order as the request | [optional] [default to undefined]
**results_country** | **string** | ISO 3166-1 alpha-2 country code | [optional] [default to undefined]
**results_year** | **number** | Year queried | [optional] [default to undefined]
**results_found** | **boolean** | false when no holidays exist for that country/year combination | [optional] [default to undefined]
**results_holidays** | **Array&lt;any&gt;** | List of holidays. Omitted when found is false. | [optional] [default to undefined]
**results_total** | **number** | Number of holidays. Omitted when found is false. | [optional] [default to undefined]
**total** | **number** | Total number of results (equals the number of queries sent) | [optional] [default to undefined]

## Example

```typescript
import { V1PlacesHolidaysBatchPost200ResponseData } from 'api-client';

const instance: V1PlacesHolidaysBatchPost200ResponseData = {
    results,
    results_country,
    results_year,
    results_found,
    results_holidays,
    results_total,
    total,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
