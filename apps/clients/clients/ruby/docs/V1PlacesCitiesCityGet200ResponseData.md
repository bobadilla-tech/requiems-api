# OpenapiClient::V1PlacesCitiesCityGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **name** | **String** | Official city name as listed in the GeoNames dataset | [optional] |
| **country** | **String** | ISO 3166-1 alpha-2 country code (uppercase) | [optional] |
| **population** | **Integer** | City population from the GeoNames dataset | [optional] |
| **timezone** | **String** | IANA timezone identifier for the city (e.g. \&quot;America/New_York\&quot;) | [optional] |
| **lat** | **Float** | Latitude of the city centre | [optional] |
| **lon** | **Float** | Longitude of the city centre | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1PlacesCitiesCityGet200ResponseData.new(
  name: null,
  country: null,
  population: null,
  timezone: null,
  lat: null,
  lon: null
)
```

