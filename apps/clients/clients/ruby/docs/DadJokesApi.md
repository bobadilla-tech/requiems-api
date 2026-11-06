# OpenapiClient::DadJokesApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_entertainment_jokes_dad_get**](DadJokesApi.md#v1_entertainment_jokes_dad_get) | **GET** /v1/entertainment/jokes/dad | Get Random Dad Joke |


## v1_entertainment_jokes_dad_get

> <V1EntertainmentJokesDadGet200Response> v1_entertainment_jokes_dad_get

Get Random Dad Joke

Returns a randomly selected dad joke from the collection.

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

api_instance = OpenapiClient::DadJokesApi.new

begin
  # Get Random Dad Joke
  result = api_instance.v1_entertainment_jokes_dad_get
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling DadJokesApi->v1_entertainment_jokes_dad_get: #{e}"
end
```

#### Using the v1_entertainment_jokes_dad_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1EntertainmentJokesDadGet200Response>, Integer, Hash)> v1_entertainment_jokes_dad_get_with_http_info

```ruby
begin
  # Get Random Dad Joke
  data, status_code, headers = api_instance.v1_entertainment_jokes_dad_get_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1EntertainmentJokesDadGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling DadJokesApi->v1_entertainment_jokes_dad_get_with_http_info: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**V1EntertainmentJokesDadGet200Response**](V1EntertainmentJokesDadGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

