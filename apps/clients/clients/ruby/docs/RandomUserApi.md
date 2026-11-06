# OpenapiClient::RandomUserApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_technology_random_user_get**](RandomUserApi.md#v1_technology_random_user_get) | **GET** /v1/technology/random-user | Get Random User |


## v1_technology_random_user_get

> <V1TechnologyRandomUserGet200Response> v1_technology_random_user_get

Get Random User

Returns a randomly generated fake user profile.

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

api_instance = OpenapiClient::RandomUserApi.new

begin
  # Get Random User
  result = api_instance.v1_technology_random_user_get
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling RandomUserApi->v1_technology_random_user_get: #{e}"
end
```

#### Using the v1_technology_random_user_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1TechnologyRandomUserGet200Response>, Integer, Hash)> v1_technology_random_user_get_with_http_info

```ruby
begin
  # Get Random User
  data, status_code, headers = api_instance.v1_technology_random_user_get_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1TechnologyRandomUserGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling RandomUserApi->v1_technology_random_user_get_with_http_info: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**V1TechnologyRandomUserGet200Response**](V1TechnologyRandomUserGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

