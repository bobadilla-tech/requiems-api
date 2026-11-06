# OpenapiClient::V1SystemsLocationResolvePost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **address** | **String** | Normalized full address returned by the geocoder | [optional] |
| **city** | **String** | City name | [optional] |
| **country** | **String** | Full country name | [optional] |
| **country_code** | **String** | ISO 3166-1 alpha-2 country code | [optional] |
| **coordinates** | **Object** | Resolved lat/lng coordinates | [optional] |
| **timezone** | **String** | IANA timezone identifier (e.g. Europe/Berlin). Null when timezone resolution fails. | [optional] |
| **utc_offset** | **String** | UTC offset string (e.g. +02:00). Null when timezone resolution fails. | [optional] |
| **current_time** | **String** | Current local time in RFC3339 format. Null when timezone resolution fails. | [optional] |
| **is_holiday_today** | **Boolean** | True when today is a public holiday in the resolved country | [optional] |
| **working_days_this_month** | **Integer** | Number of working days (excluding weekends and public holidays) in the current calendar month | [optional] |
| **next_holiday** | **String** | Next public holiday within 90 days. Shape: {date: YYYY-MM-DD, name: string, type: national} | [optional] |
| **flags** | **Array&lt;Object&gt;** | Non-fatal warnings. Possible values: timezone_unavailable, calendar_unavailable | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1SystemsLocationResolvePost200ResponseData.new(
  address: null,
  city: null,
  country: null,
  country_code: null,
  coordinates: null,
  timezone: null,
  utc_offset: null,
  current_time: null,
  is_holiday_today: null,
  working_days_this_month: null,
  next_holiday: null,
  flags: null
)
```

