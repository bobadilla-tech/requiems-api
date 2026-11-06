# OpenapiClient::V1SystemsBusinessCalendarCountryGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **country_code** | **String** | ISO 3166-1 alpha-2 country code echoed from the path parameter | [optional] |
| **year** | **Integer** | Year the calendar covers | [optional] |
| **month** | **String** | Month scope. Present only when the month query parameter was provided. | [optional] |
| **working_days** | **Integer** | Number of working days (excluding weekends and public holidays) in the scope | [optional] |
| **total_days** | **String** | Total calendar days in the month. Present only when month is provided. | [optional] |
| **weekend_days** | **String** | Number of weekend days in the month. Present only when month is provided. | [optional] |
| **holidays** | **Array&lt;Object&gt;** | List of public holidays in the scope. Each item: {date: YYYY-MM-DD, name: string, type: national} | [optional] |
| **holiday_count** | **Integer** | Count of holidays in the holidays array | [optional] |
| **next_holiday** | **String** | Next public holiday within 90 days from today. Shape: {date, name, type} | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1SystemsBusinessCalendarCountryGet200ResponseData.new(
  country_code: null,
  year: null,
  month: null,
  working_days: null,
  total_days: null,
  weekend_days: null,
  holidays: null,
  holiday_count: null,
  next_holiday: null
)
```

