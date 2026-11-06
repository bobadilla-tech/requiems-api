# V1SystemsLocationResolvePostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**address** | **string** | Human-readable address string to geocode. Either address or coordinates is required. | [optional] [default to undefined]
**coordinates** | **object** | Coordinates object for reverse geocoding. Either address or coordinates is required. Shape: {lat: float, lng: float} | [optional] [default to undefined]
**country_code** | **string** | ISO 3166-1 alpha-2 country code hint. Overrides the geocoder-inferred country code when provided. | [optional] [default to undefined]

## Example

```typescript
import { V1SystemsLocationResolvePostRequest } from 'api-client';

const instance: V1SystemsLocationResolvePostRequest = {
    address,
    coordinates,
    country_code,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
