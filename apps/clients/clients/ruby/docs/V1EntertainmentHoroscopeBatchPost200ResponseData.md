# OpenapiClient::V1EntertainmentHoroscopeBatchPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **results** | **Array&lt;Object&gt;** | List of daily horoscopes returned in the same order as the input signs. | [optional] |
| **results_sign** | **String** | Zodiac sign the horoscope corresponds to. | [optional] |
| **results_date** | **String** | Date the horoscope is for, in YYYY-MM-DD format. | [optional] |
| **results_horoscope** | **String** | Daily horoscope reading for the given sign. | [optional] |
| **results_lucky_number** | **Integer** | Lucky number for the given sign on that date. | [optional] |
| **results_mood** | **String** | Predicted mood for the given sign on that date. | [optional] |
| **total** | **Integer** | Total number of horoscopes returned. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1EntertainmentHoroscopeBatchPost200ResponseData.new(
  results: null,
  results_sign: null,
  results_date: null,
  results_horoscope: null,
  results_lucky_number: null,
  results_mood: null,
  total: null
)
```

