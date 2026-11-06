# OpenapiClient::V1PlacesTimezoneGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **timezone** | **String** | IANA timezone identifier (e.g. \&quot;Europe/London\&quot;, \&quot;Asia/Tokyo\&quot;) | [optional] |
| **offset** | **String** | UTC offset in +HH:MM or -HH:MM format (e.g. &#39;+05:30&#39;, &#39;-05:00&#39;) | [optional] |
| **current_time** | **String** | Current time in UTC, formatted as RFC 3339 (e.g. \&quot;2024-12-15T14:30:00Z\&quot;) | [optional] |
| **is_dst** | **Boolean** | Whether the location is currently observing daylight saving time | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1PlacesTimezoneGet200ResponseData.new(
  timezone: null,
  offset: null,
  current_time: null,
  is_dst: null
)
```

