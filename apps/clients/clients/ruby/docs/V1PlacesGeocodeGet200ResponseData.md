# OpenapiClient::V1PlacesGeocodeGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **address** | **String** | Full display name of the matched location | [optional] |
| **city** | **String** | City or town of the matched location | [optional] |
| **country** | **String** | ISO 3166-1 alpha-2 country code (uppercase) | [optional] |
| **lat** | **Float** | Latitude of the matched location | [optional] |
| **lon** | **Float** | Longitude of the matched location | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1PlacesGeocodeGet200ResponseData.new(
  address: null,
  city: null,
  country: null,
  lat: null,
  lon: null
)
```

