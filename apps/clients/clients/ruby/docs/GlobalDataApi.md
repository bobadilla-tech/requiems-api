# OpenapiClient::GlobalDataApi

All URIs are relative to *https://requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_systems_business_calendar_country_get**](GlobalDataApi.md#v1_systems_business_calendar_country_get) | **GET** /v1/systems/business-calendar/{country} | Business Calendar |
| [**v1_systems_location_resolve_post**](GlobalDataApi.md#v1_systems_location_resolve_post) | **POST** /v1/systems/location/resolve | Resolve Location |
| [**v1_systems_timezone_from_ip_ip_get**](GlobalDataApi.md#v1_systems_timezone_from_ip_ip_get) | **GET** /v1/systems/timezone/from-ip/{ip} | Timezone from IP |


## v1_systems_business_calendar_country_get

> <V1SystemsBusinessCalendarCountryGet200Response> v1_systems_business_calendar_country_get(country, opts)

Business Calendar

Retrieve working days, public holidays, and next upcoming holiday for any country and optional year/month. Covers 100+ countries.

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

api_instance = OpenapiClient::GlobalDataApi.new
country = 'DE' # String | ISO 3166-1 alpha-2 country code (e.g. DE, US, FR)
opts = {
  year: 2026, # Integer | Year to retrieve the calendar for. Defaults to the current year.
  month: 6 # Integer | Month (1–12). When provided, returns month-scoped working day count, total days, weekend days, and filtered holidays.
}

begin
  # Business Calendar
  result = api_instance.v1_systems_business_calendar_country_get(country, opts)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling GlobalDataApi->v1_systems_business_calendar_country_get: #{e}"
end
```

#### Using the v1_systems_business_calendar_country_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1SystemsBusinessCalendarCountryGet200Response>, Integer, Hash)> v1_systems_business_calendar_country_get_with_http_info(country, opts)

```ruby
begin
  # Business Calendar
  data, status_code, headers = api_instance.v1_systems_business_calendar_country_get_with_http_info(country, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1SystemsBusinessCalendarCountryGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling GlobalDataApi->v1_systems_business_calendar_country_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **country** | **String** | ISO 3166-1 alpha-2 country code (e.g. DE, US, FR) |  |
| **year** | **Integer** | Year to retrieve the calendar for. Defaults to the current year. | [optional] |
| **month** | **Integer** | Month (1–12). When provided, returns month-scoped working day count, total days, weekend days, and filtered holidays. | [optional] |

### Return type

[**V1SystemsBusinessCalendarCountryGet200Response**](V1SystemsBusinessCalendarCountryGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## v1_systems_location_resolve_post

> <V1SystemsLocationResolvePost200Response> v1_systems_location_resolve_post(v1_systems_location_resolve_post_request)

Resolve Location

Resolve an address or coordinates into a full location profile — city, country, timezone, UTC offset, current time, working days this month, and next holiday.

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

api_instance = OpenapiClient::GlobalDataApi.new
v1_systems_location_resolve_post_request = OpenapiClient::V1SystemsLocationResolvePostRequest.new # V1SystemsLocationResolvePostRequest | 

begin
  # Resolve Location
  result = api_instance.v1_systems_location_resolve_post(v1_systems_location_resolve_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling GlobalDataApi->v1_systems_location_resolve_post: #{e}"
end
```

#### Using the v1_systems_location_resolve_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1SystemsLocationResolvePost200Response>, Integer, Hash)> v1_systems_location_resolve_post_with_http_info(v1_systems_location_resolve_post_request)

```ruby
begin
  # Resolve Location
  data, status_code, headers = api_instance.v1_systems_location_resolve_post_with_http_info(v1_systems_location_resolve_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1SystemsLocationResolvePost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling GlobalDataApi->v1_systems_location_resolve_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_systems_location_resolve_post_request** | [**V1SystemsLocationResolvePostRequest**](V1SystemsLocationResolvePostRequest.md) |  |  |

### Return type

[**V1SystemsLocationResolvePost200Response**](V1SystemsLocationResolvePost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## v1_systems_timezone_from_ip_ip_get

> <V1SystemsTimezoneFromIpIpGet200Response> v1_systems_timezone_from_ip_ip_get(ip)

Timezone from IP

Resolve the timezone, UTC offset, and DST status for any IPv4 or IPv6 address without needing an address string.

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

api_instance = OpenapiClient::GlobalDataApi.new
ip = '92.168.1.1' # String | IPv4 or IPv6 address to resolve

begin
  # Timezone from IP
  result = api_instance.v1_systems_timezone_from_ip_ip_get(ip)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling GlobalDataApi->v1_systems_timezone_from_ip_ip_get: #{e}"
end
```

#### Using the v1_systems_timezone_from_ip_ip_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1SystemsTimezoneFromIpIpGet200Response>, Integer, Hash)> v1_systems_timezone_from_ip_ip_get_with_http_info(ip)

```ruby
begin
  # Timezone from IP
  data, status_code, headers = api_instance.v1_systems_timezone_from_ip_ip_get_with_http_info(ip)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1SystemsTimezoneFromIpIpGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling GlobalDataApi->v1_systems_timezone_from_ip_ip_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **ip** | **String** | IPv4 or IPv6 address to resolve |  |

### Return type

[**V1SystemsTimezoneFromIpIpGet200Response**](V1SystemsTimezoneFromIpIpGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

