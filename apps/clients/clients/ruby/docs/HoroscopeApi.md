# OpenapiClient::HoroscopeApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_entertainment_horoscope_sign_get**](HoroscopeApi.md#v1_entertainment_horoscope_sign_get) | **GET** /v1/entertainment/horoscope/{sign} | Get Daily Horoscope |


## v1_entertainment_horoscope_sign_get

> <V1EntertainmentHoroscopeSignGet200Response> v1_entertainment_horoscope_sign_get(sign)

Get Daily Horoscope

Returns a daily horoscope reading for the specified zodiac sign.

### Examples

```ruby
require 'time'
require 'openapi_client'
# setup authorization
OpenapiClient.configure do |config|
  # Configure API key authorization: requiems-api-key
  config.api_key['requiems-api-key'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['requiems-api-key'] = 'Bearer'
end

api_instance = OpenapiClient::HoroscopeApi.new
sign = 'sign_example' # String | Zodiac sign (case-insensitive). Supported values: aries, taurus, gemini, cancer, leo, virgo, libra, scorpio, sagittarius, capricorn, aquarius, pisces

begin
  # Get Daily Horoscope
  result = api_instance.v1_entertainment_horoscope_sign_get(sign)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling HoroscopeApi->v1_entertainment_horoscope_sign_get: #{e}"
end
```

#### Using the v1_entertainment_horoscope_sign_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1EntertainmentHoroscopeSignGet200Response>, Integer, Hash)> v1_entertainment_horoscope_sign_get_with_http_info(sign)

```ruby
begin
  # Get Daily Horoscope
  data, status_code, headers = api_instance.v1_entertainment_horoscope_sign_get_with_http_info(sign)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1EntertainmentHoroscopeSignGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling HoroscopeApi->v1_entertainment_horoscope_sign_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **sign** | **String** | Zodiac sign (case-insensitive). Supported values: aries, taurus, gemini, cancer, leo, virgo, libra, scorpio, sagittarius, capricorn, aquarius, pisces |  |

### Return type

[**V1EntertainmentHoroscopeSignGet200Response**](V1EntertainmentHoroscopeSignGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

