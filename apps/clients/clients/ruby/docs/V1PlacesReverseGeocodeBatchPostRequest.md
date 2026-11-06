# OpenapiClient::V1PlacesReverseGeocodeBatchPostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **items** | **String** | List of coordinate pairs to reverse-geocode (1–20 items). Max 20 due to Nominatim usage policy. |  |
| **items_lat** | **Float** | Latitude (-90 to 90). |  |
| **items_lon** | **Float** | Longitude (-180 to 180). |  |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1PlacesReverseGeocodeBatchPostRequest.new(
  items: null,
  items_lat: null,
  items_lon: null
)
```

