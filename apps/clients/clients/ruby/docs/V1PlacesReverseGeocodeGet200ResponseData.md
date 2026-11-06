# OpenapiClient::V1PlacesReverseGeocodeGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **lat** | **Float** | Latitude as provided in the request | [optional] |
| **lon** | **Float** | Longitude as provided in the request | [optional] |
| **address** | **String** | Full display name of the location at the given coordinates | [optional] |
| **city** | **String** | City or town at the given coordinates | [optional] |
| **country** | **String** | ISO 3166-1 alpha-2 country code (uppercase) | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1PlacesReverseGeocodeGet200ResponseData.new(
  lat: null,
  lon: null,
  address: null,
  city: null,
  country: null
)
```

