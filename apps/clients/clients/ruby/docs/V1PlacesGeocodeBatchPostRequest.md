# OpenapiClient::V1PlacesGeocodeBatchPostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **addresses** | **String** | List of addresses to geocode (1–20 items, each non-empty). Max 20 due to Nominatim usage policy. |  |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1PlacesGeocodeBatchPostRequest.new(
  addresses: null
)
```

