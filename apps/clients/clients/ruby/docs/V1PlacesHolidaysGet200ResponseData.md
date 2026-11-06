# OpenapiClient::V1PlacesHolidaysGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **country** | **String** | ISO 3166-1 alpha-2 country code | [optional] |
| **year** | **Integer** | Year for which holidays are returned | [optional] |
| **holidays** | **Array&lt;Object&gt;** | Array of holiday objects | [optional] |
| **holidays_date** | **String** | Holiday date in YYYY-MM-DD format | [optional] |
| **holidays_name** | **String** | Name of the holiday | [optional] |
| **total** | **Integer** | Total number of holidays for the country/year | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1PlacesHolidaysGet200ResponseData.new(
  country: null,
  year: null,
  holidays: null,
  holidays_date: null,
  holidays_name: null,
  total: null
)
```

