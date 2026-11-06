# OpenapiClient::HolidaysApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_places_holidays_batch_post**](HolidaysApi.md#v1_places_holidays_batch_post) | **POST** /v1/places/holidays/batch | Batch Get Holidays |
| [**v1_places_holidays_get**](HolidaysApi.md#v1_places_holidays_get) | **GET** /v1/places/holidays | Get Holidays |


## v1_places_holidays_batch_post

> <V1PlacesHolidaysBatchPost200Response> v1_places_holidays_batch_post(v1_places_holidays_batch_post_request)

Batch Get Holidays

Returns holidays for up to 50 (country, year) pairs in a single request. Each pair is processed independently — if one combination has no data, it returns found:false without failing the entire batch.

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

api_instance = OpenapiClient::HolidaysApi.new
v1_places_holidays_batch_post_request = OpenapiClient::V1PlacesHolidaysBatchPostRequest.new({queries: [{"country": "US", "year": 2025}, {"country": "AR", "year": 2024}]}) # V1PlacesHolidaysBatchPostRequest | 

begin
  # Batch Get Holidays
  result = api_instance.v1_places_holidays_batch_post(v1_places_holidays_batch_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling HolidaysApi->v1_places_holidays_batch_post: #{e}"
end
```

#### Using the v1_places_holidays_batch_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1PlacesHolidaysBatchPost200Response>, Integer, Hash)> v1_places_holidays_batch_post_with_http_info(v1_places_holidays_batch_post_request)

```ruby
begin
  # Batch Get Holidays
  data, status_code, headers = api_instance.v1_places_holidays_batch_post_with_http_info(v1_places_holidays_batch_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1PlacesHolidaysBatchPost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling HolidaysApi->v1_places_holidays_batch_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_places_holidays_batch_post_request** | [**V1PlacesHolidaysBatchPostRequest**](V1PlacesHolidaysBatchPostRequest.md) |  |  |

### Return type

[**V1PlacesHolidaysBatchPost200Response**](V1PlacesHolidaysBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## v1_places_holidays_get

> <V1PlacesHolidaysGet200Response> v1_places_holidays_get(country, year)

Get Holidays

Returns a list of public holidays for the specified country and year

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

api_instance = OpenapiClient::HolidaysApi.new
country = 'US' # String | ISO 3166-1 alpha-2 country code (e.g., \"US\", \"GB\", \"DE\")
year = 2025 # Integer | Year for which to retrieve holidays (e.g., 2025)

begin
  # Get Holidays
  result = api_instance.v1_places_holidays_get(country, year)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling HolidaysApi->v1_places_holidays_get: #{e}"
end
```

#### Using the v1_places_holidays_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1PlacesHolidaysGet200Response>, Integer, Hash)> v1_places_holidays_get_with_http_info(country, year)

```ruby
begin
  # Get Holidays
  data, status_code, headers = api_instance.v1_places_holidays_get_with_http_info(country, year)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1PlacesHolidaysGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling HolidaysApi->v1_places_holidays_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **country** | **String** | ISO 3166-1 alpha-2 country code (e.g., \&quot;US\&quot;, \&quot;GB\&quot;, \&quot;DE\&quot;) |  |
| **year** | **Integer** | Year for which to retrieve holidays (e.g., 2025) |  |

### Return type

[**V1PlacesHolidaysGet200Response**](V1PlacesHolidaysGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

