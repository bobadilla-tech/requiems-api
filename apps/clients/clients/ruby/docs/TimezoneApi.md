# OpenapiClient::TimezoneApi

All URIs are relative to *https://requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_places_timezone_batch_post**](TimezoneApi.md#v1_places_timezone_batch_post) | **POST** /v1/places/timezone/batch | Batch Timezone Lookup |
| [**v1_places_timezone_get**](TimezoneApi.md#v1_places_timezone_get) | **GET** /v1/places/timezone | Get Timezone |


## v1_places_timezone_batch_post

> <V1PlacesTimezoneBatchPost200Response> v1_places_timezone_batch_post(v1_places_timezone_batch_post_request)

Batch Timezone Lookup

Look up timezone information for up to 50 locations in a single request. Each item can specify a timezone name, city, or lat/lon coordinates. Priority is timezone name > city > coordinates. Results are returned in input order.

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

api_instance = OpenapiClient::TimezoneApi.new
v1_places_timezone_batch_post_request = OpenapiClient::V1PlacesTimezoneBatchPostRequest.new({items: 'items_example'}) # V1PlacesTimezoneBatchPostRequest | 

begin
  # Batch Timezone Lookup
  result = api_instance.v1_places_timezone_batch_post(v1_places_timezone_batch_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling TimezoneApi->v1_places_timezone_batch_post: #{e}"
end
```

#### Using the v1_places_timezone_batch_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1PlacesTimezoneBatchPost200Response>, Integer, Hash)> v1_places_timezone_batch_post_with_http_info(v1_places_timezone_batch_post_request)

```ruby
begin
  # Batch Timezone Lookup
  data, status_code, headers = api_instance.v1_places_timezone_batch_post_with_http_info(v1_places_timezone_batch_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1PlacesTimezoneBatchPost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling TimezoneApi->v1_places_timezone_batch_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_places_timezone_batch_post_request** | [**V1PlacesTimezoneBatchPostRequest**](V1PlacesTimezoneBatchPostRequest.md) |  |  |

### Return type

[**V1PlacesTimezoneBatchPost200Response**](V1PlacesTimezoneBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## v1_places_timezone_get

> <V1PlacesTimezoneGet200Response> v1_places_timezone_get(opts)

Get Timezone

Returns timezone information for the given coordinates or city name. Provide either `city` or both `lat` and `lon`.

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

api_instance = OpenapiClient::TimezoneApi.new
opts = {
  lat: 'lat_example', # String | Latitude of the location (-90 to 90). Required when using coordinate-based lookup.
  lon: 'lon_example', # String | Longitude of the location (-180 to 180). Required when using coordinate-based lookup.
  city: 'city_example' # String | City name for city-based lookup (e.g. 'Tokyo', 'London'). Required when not using coordinates.
}

begin
  # Get Timezone
  result = api_instance.v1_places_timezone_get(opts)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling TimezoneApi->v1_places_timezone_get: #{e}"
end
```

#### Using the v1_places_timezone_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1PlacesTimezoneGet200Response>, Integer, Hash)> v1_places_timezone_get_with_http_info(opts)

```ruby
begin
  # Get Timezone
  data, status_code, headers = api_instance.v1_places_timezone_get_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1PlacesTimezoneGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling TimezoneApi->v1_places_timezone_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **lat** | **String** | Latitude of the location (-90 to 90). Required when using coordinate-based lookup. | [optional] |
| **lon** | **String** | Longitude of the location (-180 to 180). Required when using coordinate-based lookup. | [optional] |
| **city** | **String** | City name for city-based lookup (e.g. &#39;Tokyo&#39;, &#39;London&#39;). Required when not using coordinates. | [optional] |

### Return type

[**V1PlacesTimezoneGet200Response**](V1PlacesTimezoneGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

