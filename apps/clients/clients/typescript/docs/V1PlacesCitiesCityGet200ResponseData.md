# V1PlacesCitiesCityGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | Official city name as listed in the GeoNames dataset | [optional] [default to undefined]
**country** | **string** | ISO 3166-1 alpha-2 country code (uppercase) | [optional] [default to undefined]
**population** | **number** | City population from the GeoNames dataset | [optional] [default to undefined]
**timezone** | **string** | IANA timezone identifier for the city (e.g. \&quot;America/New_York\&quot;) | [optional] [default to undefined]
**lat** | **number** | Latitude of the city centre | [optional] [default to undefined]
**lon** | **number** | Longitude of the city centre | [optional] [default to undefined]

## Example

```typescript
import { V1PlacesCitiesCityGet200ResponseData } from 'api-client';

const instance: V1PlacesCitiesCityGet200ResponseData = {
    name,
    country,
    population,
    timezone,
    lat,
    lon,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
