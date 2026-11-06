# V1PlacesReverseGeocodeBatchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**items** | **string** | List of coordinate pairs to reverse-geocode (1–20 items). Max 20 due to Nominatim usage policy. | [default to undefined]
**items_lat** | **number** | Latitude (-90 to 90). | [default to undefined]
**items_lon** | **number** | Longitude (-180 to 180). | [default to undefined]

## Example

```typescript
import { V1PlacesReverseGeocodeBatchPostRequest } from 'api-client';

const instance: V1PlacesReverseGeocodeBatchPostRequest = {
    items,
    items_lat,
    items_lon,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
