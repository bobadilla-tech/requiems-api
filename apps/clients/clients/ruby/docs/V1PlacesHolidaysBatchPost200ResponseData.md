# OpenapiClient::V1PlacesHolidaysBatchPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **results** | **Array&lt;Object&gt;** | One result per query, in the same order as the request | [optional] |
| **results_country** | **String** | ISO 3166-1 alpha-2 country code | [optional] |
| **results_year** | **Integer** | Year queried | [optional] |
| **results_found** | **Boolean** | false when no holidays exist for that country/year combination | [optional] |
| **results_holidays** | **Array&lt;Object&gt;** | List of holidays. Omitted when found is false. | [optional] |
| **results_total** | **Integer** | Number of holidays. Omitted when found is false. | [optional] |
| **total** | **Integer** | Total number of results (equals the number of queries sent) | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1PlacesHolidaysBatchPost200ResponseData.new(
  results: null,
  results_country: null,
  results_year: null,
  results_found: null,
  results_holidays: null,
  results_total: null,
  total: null
)
```

