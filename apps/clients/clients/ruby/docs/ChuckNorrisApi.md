# OpenapiClient::ChuckNorrisApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_entertainment_chuck_norris_get**](ChuckNorrisApi.md#v1_entertainment_chuck_norris_get) | **GET** /v1/entertainment/chuck-norris | Get Random Chuck Norris Fact |


## v1_entertainment_chuck_norris_get

> <V1EntertainmentChuckNorrisGet200Response> v1_entertainment_chuck_norris_get

Get Random Chuck Norris Fact

Returns a randomly selected Chuck Norris fact from the built-in database.

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

api_instance = OpenapiClient::ChuckNorrisApi.new

begin
  # Get Random Chuck Norris Fact
  result = api_instance.v1_entertainment_chuck_norris_get
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling ChuckNorrisApi->v1_entertainment_chuck_norris_get: #{e}"
end
```

#### Using the v1_entertainment_chuck_norris_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1EntertainmentChuckNorrisGet200Response>, Integer, Hash)> v1_entertainment_chuck_norris_get_with_http_info

```ruby
begin
  # Get Random Chuck Norris Fact
  data, status_code, headers = api_instance.v1_entertainment_chuck_norris_get_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1EntertainmentChuckNorrisGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling ChuckNorrisApi->v1_entertainment_chuck_norris_get_with_http_info: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**V1EntertainmentChuckNorrisGet200Response**](V1EntertainmentChuckNorrisGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

