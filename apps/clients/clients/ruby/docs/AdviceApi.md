# OpenapiClient::AdviceApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_text_advice_get**](AdviceApi.md#v1_text_advice_get) | **GET** /v1/text/advice | Get Random Advice |


## v1_text_advice_get

> <V1TextAdviceGet200Response> v1_text_advice_get

Get Random Advice

Returns a random piece of advice

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

api_instance = OpenapiClient::AdviceApi.new

begin
  # Get Random Advice
  result = api_instance.v1_text_advice_get
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling AdviceApi->v1_text_advice_get: #{e}"
end
```

#### Using the v1_text_advice_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1TextAdviceGet200Response>, Integer, Hash)> v1_text_advice_get_with_http_info

```ruby
begin
  # Get Random Advice
  data, status_code, headers = api_instance.v1_text_advice_get_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1TextAdviceGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling AdviceApi->v1_text_advice_get_with_http_info: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**V1TextAdviceGet200Response**](V1TextAdviceGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

