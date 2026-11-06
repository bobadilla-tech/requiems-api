# V1PlacesGeocodeBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **Array&lt;any&gt;** | Per-item results in input order. | [optional] [default to undefined]
**results_address** | **string** | The original address as provided in the request. | [optional] [default to undefined]
**results_result** | **object** | Geocoding result (omitted on error). Same fields as the single-item endpoint. | [optional] [default to undefined]
**results_error** | **string** | Error message if geocoding failed (omitted on success). | [optional] [default to undefined]
**total** | **number** | Total number of items processed. | [optional] [default to undefined]

## Example

```typescript
import { V1PlacesGeocodeBatchPost200ResponseData } from 'api-client';

const instance: V1PlacesGeocodeBatchPost200ResponseData = {
    results,
    results_address,
    results_result,
    results_error,
    total,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
