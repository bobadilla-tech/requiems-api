# OpenapiClient::HoroscopeApi

All URIs are relative to *https://requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_entertainment_horoscope_batch_post**](HoroscopeApi.md#v1_entertainment_horoscope_batch_post) | **POST** /v1/entertainment/horoscope/batch | Batch Daily Horoscopes |
| [**v1_entertainment_horoscope_sign_get**](HoroscopeApi.md#v1_entertainment_horoscope_sign_get) | **GET** /v1/entertainment/horoscope/{sign} | Get Daily Horoscope |


## v1_entertainment_horoscope_batch_post

> <V1EntertainmentHoroscopeBatchPost200Response> v1_entertainment_horoscope_batch_post(v1_entertainment_horoscope_batch_post_request)

Batch Daily Horoscopes

Returns up to 12 daily horoscopes in a single request.

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
v1_entertainment_horoscope_batch_post_request = OpenapiClient::V1EntertainmentHoroscopeBatchPostRequest.new({signs: ["taurus", "cancer", "libra"]
}) # V1EntertainmentHoroscopeBatchPostRequest | 

begin
  # Batch Daily Horoscopes
  result = api_instance.v1_entertainment_horoscope_batch_post(v1_entertainment_horoscope_batch_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling HoroscopeApi->v1_entertainment_horoscope_batch_post: #{e}"
end
```

#### Using the v1_entertainment_horoscope_batch_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1EntertainmentHoroscopeBatchPost200Response>, Integer, Hash)> v1_entertainment_horoscope_batch_post_with_http_info(v1_entertainment_horoscope_batch_post_request)

```ruby
begin
  # Batch Daily Horoscopes
  data, status_code, headers = api_instance.v1_entertainment_horoscope_batch_post_with_http_info(v1_entertainment_horoscope_batch_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1EntertainmentHoroscopeBatchPost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling HoroscopeApi->v1_entertainment_horoscope_batch_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_entertainment_horoscope_batch_post_request** | [**V1EntertainmentHoroscopeBatchPostRequest**](V1EntertainmentHoroscopeBatchPostRequest.md) |  |  |

### Return type

[**V1EntertainmentHoroscopeBatchPost200Response**](V1EntertainmentHoroscopeBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


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

