# OpenapiClient::UseragentApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_technology_useragent_get**](UseragentApi.md#v1_technology_useragent_get) | **GET** /v1/technology/useragent | Parse User Agent |


## v1_technology_useragent_get

> <V1TechnologyUseragentGet200Response> v1_technology_useragent_get(ua)

Parse User Agent

Parses a user agent string and returns structured information about the browser, OS, device, and bot status.

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

api_instance = OpenapiClient::UseragentApi.new
ua = 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36' # String | The user agent string to parse.

begin
  # Parse User Agent
  result = api_instance.v1_technology_useragent_get(ua)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling UseragentApi->v1_technology_useragent_get: #{e}"
end
```

#### Using the v1_technology_useragent_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1TechnologyUseragentGet200Response>, Integer, Hash)> v1_technology_useragent_get_with_http_info(ua)

```ruby
begin
  # Parse User Agent
  data, status_code, headers = api_instance.v1_technology_useragent_get_with_http_info(ua)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1TechnologyUseragentGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling UseragentApi->v1_technology_useragent_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **ua** | **String** | The user agent string to parse. |  |

### Return type

[**V1TechnologyUseragentGet200Response**](V1TechnologyUseragentGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

