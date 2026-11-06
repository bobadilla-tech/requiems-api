# OpenapiClient::PostalCodeApi

All URIs are relative to *https://requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_places_postal_batch_post**](PostalCodeApi.md#v1_places_postal_batch_post) | **POST** /v1/places/postal/batch | Batch Lookup Postal Codes |
| [**v1_places_postal_code_get**](PostalCodeApi.md#v1_places_postal_code_get) | **GET** /v1/places/postal/{code} | Lookup Postal Code |


## v1_places_postal_batch_post

> <V1PlacesPostalBatchPost200Response> v1_places_postal_batch_post(v1_places_postal_batch_post_request)

Batch Lookup Postal Codes

Look up city, state, and coordinates for up to 50 postal codes in a single request. Results are returned in input order.

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

api_instance = OpenapiClient::PostalCodeApi.new
v1_places_postal_batch_post_request = OpenapiClient::V1PlacesPostalBatchPostRequest.new({items: 'items_example', items_code: 'items_code_example'}) # V1PlacesPostalBatchPostRequest | 

begin
  # Batch Lookup Postal Codes
  result = api_instance.v1_places_postal_batch_post(v1_places_postal_batch_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling PostalCodeApi->v1_places_postal_batch_post: #{e}"
end
```

#### Using the v1_places_postal_batch_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1PlacesPostalBatchPost200Response>, Integer, Hash)> v1_places_postal_batch_post_with_http_info(v1_places_postal_batch_post_request)

```ruby
begin
  # Batch Lookup Postal Codes
  data, status_code, headers = api_instance.v1_places_postal_batch_post_with_http_info(v1_places_postal_batch_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1PlacesPostalBatchPost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling PostalCodeApi->v1_places_postal_batch_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_places_postal_batch_post_request** | [**V1PlacesPostalBatchPostRequest**](V1PlacesPostalBatchPostRequest.md) |  |  |

### Return type

[**V1PlacesPostalBatchPost200Response**](V1PlacesPostalBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## v1_places_postal_code_get

> <V1PlacesPostalCodeGet200Response> v1_places_postal_code_get(code, opts)

Lookup Postal Code

Returns city, state, country, and coordinates for the given postal code.

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

api_instance = OpenapiClient::PostalCodeApi.new
code = '10001' # String | The postal code to look up (e.g. 10001 for New York, SW1A 1AA for London)
opts = {
  country: 'US' # String | ISO 3166-1 alpha-2 country code (default: US)
}

begin
  # Lookup Postal Code
  result = api_instance.v1_places_postal_code_get(code, opts)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling PostalCodeApi->v1_places_postal_code_get: #{e}"
end
```

#### Using the v1_places_postal_code_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1PlacesPostalCodeGet200Response>, Integer, Hash)> v1_places_postal_code_get_with_http_info(code, opts)

```ruby
begin
  # Lookup Postal Code
  data, status_code, headers = api_instance.v1_places_postal_code_get_with_http_info(code, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1PlacesPostalCodeGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling PostalCodeApi->v1_places_postal_code_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **code** | **String** | The postal code to look up (e.g. 10001 for New York, SW1A 1AA for London) |  |
| **country** | **String** | ISO 3166-1 alpha-2 country code (default: US) | [optional] |

### Return type

[**V1PlacesPostalCodeGet200Response**](V1PlacesPostalCodeGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

