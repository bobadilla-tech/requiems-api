# OpenapiClient::V1PlacesTimezoneBatchPostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **items** | **String** | List of timezone lookup requests (1–50 items). |  |
| **items_timezone** | **String** | IANA timezone name (e.g. &#39;America/New_York&#39;). Highest priority. | [optional] |
| **items_city** | **String** | City name for city-based lookup. Used when timezone is not set. | [optional] |
| **items_lat** | **Float** | Latitude (-90 to 90). Used when neither timezone nor city is set. | [optional] |
| **items_lon** | **Float** | Longitude (-180 to 180). Used alongside lat for coordinate-based lookup. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1PlacesTimezoneBatchPostRequest.new(
  items: null,
  items_timezone: null,
  items_city: null,
  items_lat: null,
  items_lon: null
)
```

