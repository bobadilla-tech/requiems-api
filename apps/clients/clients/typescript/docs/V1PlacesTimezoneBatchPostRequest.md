# V1PlacesTimezoneBatchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**items** | **string** | List of timezone lookup requests (1–50 items). | [default to undefined]
**items_timezone** | **string** | IANA timezone name (e.g. \&#39;America/New_York\&#39;). Highest priority. | [optional] [default to undefined]
**items_city** | **string** | City name for city-based lookup. Used when timezone is not set. | [optional] [default to undefined]
**items_lat** | **number** | Latitude (-90 to 90). Used when neither timezone nor city is set. | [optional] [default to undefined]
**items_lon** | **number** | Longitude (-180 to 180). Used alongside lat for coordinate-based lookup. | [optional] [default to undefined]

## Example

```typescript
import { V1PlacesTimezoneBatchPostRequest } from 'api-client';

const instance: V1PlacesTimezoneBatchPostRequest = {
    items,
    items_timezone,
    items_city,
    items_lat,
    items_lon,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
