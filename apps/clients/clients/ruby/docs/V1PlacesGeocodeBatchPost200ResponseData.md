# OpenapiClient::V1PlacesGeocodeBatchPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **results** | **Array&lt;Object&gt;** | Per-item results in input order. | [optional] |
| **results_address** | **String** | The original address as provided in the request. | [optional] |
| **results_result** | **Object** | Geocoding result (omitted on error). Same fields as the single-item endpoint. | [optional] |
| **results_error** | **String** | Error message if geocoding failed (omitted on success). | [optional] |
| **total** | **Integer** | Total number of items processed. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1PlacesGeocodeBatchPost200ResponseData.new(
  results: null,
  results_address: null,
  results_result: null,
  results_error: null,
  total: null
)
```

