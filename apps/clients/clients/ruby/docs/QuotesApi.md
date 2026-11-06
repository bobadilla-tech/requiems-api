# OpenapiClient::QuotesApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_text_quotes_random_get**](QuotesApi.md#v1_text_quotes_random_get) | **GET** /v1/text/quotes/random | Get Random Quote |


## v1_text_quotes_random_get

> <V1TextQuotesRandomGet200Response> v1_text_quotes_random_get

Get Random Quote

Returns a random inspirational quote with author attribution

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

api_instance = OpenapiClient::QuotesApi.new

begin
  # Get Random Quote
  result = api_instance.v1_text_quotes_random_get
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling QuotesApi->v1_text_quotes_random_get: #{e}"
end
```

#### Using the v1_text_quotes_random_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1TextQuotesRandomGet200Response>, Integer, Hash)> v1_text_quotes_random_get_with_http_info

```ruby
begin
  # Get Random Quote
  data, status_code, headers = api_instance.v1_text_quotes_random_get_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1TextQuotesRandomGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling QuotesApi->v1_text_quotes_random_get_with_http_info: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**V1TextQuotesRandomGet200Response**](V1TextQuotesRandomGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

