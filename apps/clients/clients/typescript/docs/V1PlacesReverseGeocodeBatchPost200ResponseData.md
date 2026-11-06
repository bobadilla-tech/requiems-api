# V1PlacesReverseGeocodeBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **Array&lt;any&gt;** | Per-item results in input order. | [optional] [default to undefined]
**results_lat** | **number** | Latitude as provided in the request. | [optional] [default to undefined]
**results_lon** | **number** | Longitude as provided in the request. | [optional] [default to undefined]
**results_result** | **object** | Reverse geocoding result (omitted on error). Same fields as the single-item endpoint. | [optional] [default to undefined]
**results_error** | **string** | Error message if reverse geocoding failed (omitted on success). | [optional] [default to undefined]
**total** | **number** | Total number of items processed. | [optional] [default to undefined]

## Example

```typescript
import { V1PlacesReverseGeocodeBatchPost200ResponseData } from 'api-client';

const instance: V1PlacesReverseGeocodeBatchPost200ResponseData = {
    results,
    results_lat,
    results_lon,
    results_result,
    results_error,
    total,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
