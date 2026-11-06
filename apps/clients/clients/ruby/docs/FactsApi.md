# OpenapiClient::FactsApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_entertainment_facts_get**](FactsApi.md#v1_entertainment_facts_get) | **GET** /v1/entertainment/facts | Get Random Fact |


## v1_entertainment_facts_get

> <V1EntertainmentFactsGet200Response> v1_entertainment_facts_get(opts)

Get Random Fact

Returns a randomly selected fact, optionally filtered by category.

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

api_instance = OpenapiClient::FactsApi.new
opts = {
  category: 'science' # String | Filter by category. Valid values: science, history, technology, nature, space, food
}

begin
  # Get Random Fact
  result = api_instance.v1_entertainment_facts_get(opts)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling FactsApi->v1_entertainment_facts_get: #{e}"
end
```

#### Using the v1_entertainment_facts_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1EntertainmentFactsGet200Response>, Integer, Hash)> v1_entertainment_facts_get_with_http_info(opts)

```ruby
begin
  # Get Random Fact
  data, status_code, headers = api_instance.v1_entertainment_facts_get_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1EntertainmentFactsGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling FactsApi->v1_entertainment_facts_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **category** | **String** | Filter by category. Valid values: science, history, technology, nature, space, food | [optional] |

### Return type

[**V1EntertainmentFactsGet200Response**](V1EntertainmentFactsGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

