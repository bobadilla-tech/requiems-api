# V1PlacesPostalCodeGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**postal_code** | **string** | The postal code as stored in the dataset | [optional] [default to undefined]
**city** | **string** | Primary city or place name for the postal code | [optional] [default to undefined]
**state** | **string** | State, province, or administrative region name | [optional] [default to undefined]
**country** | **string** | ISO 3166-1 alpha-2 country code (uppercase) | [optional] [default to undefined]
**lat** | **number** | Latitude of the postal code centroid | [optional] [default to undefined]
**lon** | **number** | Longitude of the postal code centroid | [optional] [default to undefined]

## Example

```typescript
import { V1PlacesPostalCodeGet200ResponseData } from 'api-client';

const instance: V1PlacesPostalCodeGet200ResponseData = {
    postal_code,
    city,
    state,
    country,
    lat,
    lon,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
