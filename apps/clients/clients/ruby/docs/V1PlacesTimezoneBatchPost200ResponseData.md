# OpenapiClient::V1PlacesTimezoneBatchPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **results** | **Array&lt;Object&gt;** | Per-item results in input order. | [optional] |
| **results_info_timezone** | **String** | IANA timezone identifier. | [optional] |
| **results_info_offset** | **String** | UTC offset in +HH:MM or -HH:MM format. | [optional] |
| **results_info_current_time** | **String** | Current time in UTC (RFC 3339). | [optional] |
| **results_info_is_dst** | **Boolean** | Whether daylight saving time is currently active. | [optional] |
| **results_error** | **String** | Error message if this item failed (omitted on success). | [optional] |
| **total** | **Integer** | Total number of items processed. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1PlacesTimezoneBatchPost200ResponseData.new(
  results: null,
  results_info_timezone: null,
  results_info_offset: null,
  results_info_current_time: null,
  results_info_is_dst: null,
  results_error: null,
  total: null
)
```

