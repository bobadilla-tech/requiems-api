# OpenapiClient::V1PlacesPostalCodeGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **postal_code** | **String** | The postal code as stored in the dataset | [optional] |
| **city** | **String** | Primary city or place name for the postal code | [optional] |
| **state** | **String** | State, province, or administrative region name | [optional] |
| **country** | **String** | ISO 3166-1 alpha-2 country code (uppercase) | [optional] |
| **lat** | **Float** | Latitude of the postal code centroid | [optional] |
| **lon** | **Float** | Longitude of the postal code centroid | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1PlacesPostalCodeGet200ResponseData.new(
  postal_code: null,
  city: null,
  state: null,
  country: null,
  lat: null,
  lon: null
)
```

