# OpenapiClient::V1EntertainmentHoroscopeSignGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **sign** | **String** | Normalized zodiac sign (lowercase) | [optional] |
| **date** | **String** | Today&#39;s date in YYYY-MM-DD format (UTC) | [optional] |
| **horoscope** | **String** | Daily horoscope reading | [optional] |
| **lucky_number** | **Integer** | Lucky number for the day (1-99) | [optional] |
| **mood** | **String** | Suggested mood for the day | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1EntertainmentHoroscopeSignGet200ResponseData.new(
  sign: null,
  date: null,
  horoscope: null,
  lucky_number: null,
  mood: null
)
```

