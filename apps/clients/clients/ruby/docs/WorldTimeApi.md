# OpenapiClient::WorldTimeApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_places_time_timezone_get**](WorldTimeApi.md#v1_places_time_timezone_get) | **GET** /v1/places/time/{timezone} | Get Current Time by Timezone |


## v1_places_time_timezone_get

> <V1PlacesTimeTimezoneGet200Response> v1_places_time_timezone_get(timezone)

Get Current Time by Timezone

Returns the current time for the given IANA timezone identifier. The timezone is supplied as a path parameter (e.g. `America/New_York`, `Europe/London`, `UTC`).

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

api_instance = OpenapiClient::WorldTimeApi.new
timezone = 'America/New_York' # String | IANA timezone identifier (e.g. 'America/New_York', 'Europe/London', 'Asia/Kolkata')

begin
  # Get Current Time by Timezone
  result = api_instance.v1_places_time_timezone_get(timezone)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling WorldTimeApi->v1_places_time_timezone_get: #{e}"
end
```

#### Using the v1_places_time_timezone_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1PlacesTimeTimezoneGet200Response>, Integer, Hash)> v1_places_time_timezone_get_with_http_info(timezone)

```ruby
begin
  # Get Current Time by Timezone
  data, status_code, headers = api_instance.v1_places_time_timezone_get_with_http_info(timezone)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1PlacesTimeTimezoneGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling WorldTimeApi->v1_places_time_timezone_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **timezone** | **String** | IANA timezone identifier (e.g. &#39;America/New_York&#39;, &#39;Europe/London&#39;, &#39;Asia/Kolkata&#39;) |  |

### Return type

[**V1PlacesTimeTimezoneGet200Response**](V1PlacesTimeTimezoneGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

