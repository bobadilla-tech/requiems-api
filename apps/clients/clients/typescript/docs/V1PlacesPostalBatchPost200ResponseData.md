# V1PlacesPostalBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **Array&lt;any&gt;** | Per-item results in input order. | [optional] [default to undefined]
**results_code** | **string** | The postal code as provided in the request. | [optional] [default to undefined]
**results_country** | **string** | The country code used for the lookup (defaulted to US when omitted). | [optional] [default to undefined]
**results_found** | **boolean** | Whether the postal code was found in the dataset. | [optional] [default to undefined]
**results_result** | **object** | Postal code details (omitted when found is false). | [optional] [default to undefined]
**total** | **number** | Total number of items processed. | [optional] [default to undefined]

## Example

```typescript
import { V1PlacesPostalBatchPost200ResponseData } from 'api-client';

const instance: V1PlacesPostalBatchPost200ResponseData = {
    results,
    results_code,
    results_country,
    results_found,
    results_result,
    total,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
