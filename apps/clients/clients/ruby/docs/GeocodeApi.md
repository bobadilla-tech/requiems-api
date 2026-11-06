# OpenapiClient::GeocodeApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_places_geocode_get**](GeocodeApi.md#v1_places_geocode_get) | **GET** /v1/places/geocode | Geocode Address |
| [**v1_places_reverse_geocode_get**](GeocodeApi.md#v1_places_reverse_geocode_get) | **GET** /v1/places/reverse-geocode | Reverse Geocode |


## v1_places_geocode_get

> <V1PlacesGeocodeGet200Response> v1_places_geocode_get(address)

Geocode Address

Converts a free-text address into latitude and longitude coordinates.

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

api_instance = OpenapiClient::GeocodeApi.new
address = '1600 Pennsylvania Ave NW, Washington DC' # String | The address to geocode (street, city, country, or any combination)

begin
  # Geocode Address
  result = api_instance.v1_places_geocode_get(address)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling GeocodeApi->v1_places_geocode_get: #{e}"
end
```

#### Using the v1_places_geocode_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1PlacesGeocodeGet200Response>, Integer, Hash)> v1_places_geocode_get_with_http_info(address)

```ruby
begin
  # Geocode Address
  data, status_code, headers = api_instance.v1_places_geocode_get_with_http_info(address)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1PlacesGeocodeGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling GeocodeApi->v1_places_geocode_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **address** | **String** | The address to geocode (street, city, country, or any combination) |  |

### Return type

[**V1PlacesGeocodeGet200Response**](V1PlacesGeocodeGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## v1_places_reverse_geocode_get

> <V1PlacesReverseGeocodeGet200Response> v1_places_reverse_geocode_get(lat, lon)

Reverse Geocode

Converts geographic coordinates into a human-readable address.

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

api_instance = OpenapiClient::GeocodeApi.new
lat = 38.8977 # Float | Latitude of the location (-90 to 90)
lon = -77.0365 # Float | Longitude of the location (-180 to 180)

begin
  # Reverse Geocode
  result = api_instance.v1_places_reverse_geocode_get(lat, lon)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling GeocodeApi->v1_places_reverse_geocode_get: #{e}"
end
```

#### Using the v1_places_reverse_geocode_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1PlacesReverseGeocodeGet200Response>, Integer, Hash)> v1_places_reverse_geocode_get_with_http_info(lat, lon)

```ruby
begin
  # Reverse Geocode
  data, status_code, headers = api_instance.v1_places_reverse_geocode_get_with_http_info(lat, lon)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1PlacesReverseGeocodeGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling GeocodeApi->v1_places_reverse_geocode_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **lat** | **Float** | Latitude of the location (-90 to 90) |  |
| **lon** | **Float** | Longitude of the location (-180 to 180) |  |

### Return type

[**V1PlacesReverseGeocodeGet200Response**](V1PlacesReverseGeocodeGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

