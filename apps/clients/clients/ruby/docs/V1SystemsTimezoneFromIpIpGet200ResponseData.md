# OpenapiClient::V1SystemsTimezoneFromIpIpGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **ip** | **String** | The IP address as provided | [optional] |
| **city** | **String** | City resolved from the IP | [optional] |
| **country_code** | **String** | ISO 3166-1 alpha-2 country code | [optional] |
| **timezone** | **String** | IANA timezone identifier. Null when timezone resolution fails. | [optional] |
| **utc_offset** | **String** | UTC offset string. Null when timezone resolution fails. | [optional] |
| **dst_active** | **String** | True when daylight saving time is currently active. Null when timezone resolution fails. | [optional] |
| **current_time** | **String** | Current local time in RFC3339 format | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1SystemsTimezoneFromIpIpGet200ResponseData.new(
  ip: null,
  city: null,
  country_code: null,
  timezone: null,
  utc_offset: null,
  dst_active: null,
  current_time: null
)
```

