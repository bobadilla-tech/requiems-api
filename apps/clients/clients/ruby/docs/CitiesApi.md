# OpenapiClient::CitiesApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_places_cities_city_get**](CitiesApi.md#v1_places_cities_city_get) | **GET** /v1/places/cities/{city} | Get City Info |


## v1_places_cities_city_get

> <V1PlacesCitiesCityGet200Response> v1_places_cities_city_get(city)

Get City Info

Returns metadata for a city by name. Lookup is case-insensitive.

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

api_instance = OpenapiClient::CitiesApi.new
city = 'london' # String | City name to look up (e.g. london, tokyo, new york city)

begin
  # Get City Info
  result = api_instance.v1_places_cities_city_get(city)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling CitiesApi->v1_places_cities_city_get: #{e}"
end
```

#### Using the v1_places_cities_city_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1PlacesCitiesCityGet200Response>, Integer, Hash)> v1_places_cities_city_get_with_http_info(city)

```ruby
begin
  # Get City Info
  data, status_code, headers = api_instance.v1_places_cities_city_get_with_http_info(city)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1PlacesCitiesCityGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling CitiesApi->v1_places_cities_city_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **city** | **String** | City name to look up (e.g. london, tokyo, new york city) |  |

### Return type

[**V1PlacesCitiesCityGet200Response**](V1PlacesCitiesCityGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

