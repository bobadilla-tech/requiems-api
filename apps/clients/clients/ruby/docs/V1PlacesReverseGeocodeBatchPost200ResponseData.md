# OpenapiClient::V1PlacesReverseGeocodeBatchPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **results** | **Array&lt;Object&gt;** | Per-item results in input order. | [optional] |
| **results_lat** | **Float** | Latitude as provided in the request. | [optional] |
| **results_lon** | **Float** | Longitude as provided in the request. | [optional] |
| **results_result** | **Object** | Reverse geocoding result (omitted on error). Same fields as the single-item endpoint. | [optional] |
| **results_error** | **String** | Error message if reverse geocoding failed (omitted on success). | [optional] |
| **total** | **Integer** | Total number of items processed. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1PlacesReverseGeocodeBatchPost200ResponseData.new(
  results: null,
  results_lat: null,
  results_lon: null,
  results_result: null,
  results_error: null,
  total: null
)
```

