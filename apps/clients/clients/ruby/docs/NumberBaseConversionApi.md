# OpenapiClient::NumberBaseConversionApi

All URIs are relative to *https://requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_technology_base_batch_post**](NumberBaseConversionApi.md#v1_technology_base_batch_post) | **POST** /v1/technology/base/batch | Batch Convert Base |
| [**v1_technology_base_get**](NumberBaseConversionApi.md#v1_technology_base_get) | **GET** /v1/technology/base | Convert Base |


## v1_technology_base_batch_post

> <V1TechnologyBaseBatchPost200Response> v1_technology_base_batch_post(v1_technology_base_batch_post_request)

Batch Convert Base

Convert up to 50 numbers between bases in a single request. Results are returned in input order. Per-item errors are reported inline.

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

api_instance = OpenapiClient::NumberBaseConversionApi.new
v1_technology_base_batch_post_request = OpenapiClient::V1TechnologyBaseBatchPostRequest.new({items: 'items_example', items_from: 37, items_to: 37, items_value: 'items_value_example'}) # V1TechnologyBaseBatchPostRequest | 

begin
  # Batch Convert Base
  result = api_instance.v1_technology_base_batch_post(v1_technology_base_batch_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling NumberBaseConversionApi->v1_technology_base_batch_post: #{e}"
end
```

#### Using the v1_technology_base_batch_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1TechnologyBaseBatchPost200Response>, Integer, Hash)> v1_technology_base_batch_post_with_http_info(v1_technology_base_batch_post_request)

```ruby
begin
  # Batch Convert Base
  data, status_code, headers = api_instance.v1_technology_base_batch_post_with_http_info(v1_technology_base_batch_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1TechnologyBaseBatchPost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling NumberBaseConversionApi->v1_technology_base_batch_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_technology_base_batch_post_request** | [**V1TechnologyBaseBatchPostRequest**](V1TechnologyBaseBatchPostRequest.md) |  |  |

### Return type

[**V1TechnologyBaseBatchPost200Response**](V1TechnologyBaseBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## v1_technology_base_get

> <V1TechnologyBaseGet200Response> v1_technology_base_get(from, to, value)

Convert Base

Convert an integer from one number base to another.

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

api_instance = OpenapiClient::NumberBaseConversionApi.new
from = 10 # Integer | Source base (2, 8, 10, or 16)
to = 16 # Integer | Target base (2, 8, 10, or 16)
value = '255' # String | The number as a string. Accepts optional prefixes: 0x (hex), 0b (binary), 0o (octal).

begin
  # Convert Base
  result = api_instance.v1_technology_base_get(from, to, value)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling NumberBaseConversionApi->v1_technology_base_get: #{e}"
end
```

#### Using the v1_technology_base_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1TechnologyBaseGet200Response>, Integer, Hash)> v1_technology_base_get_with_http_info(from, to, value)

```ruby
begin
  # Convert Base
  data, status_code, headers = api_instance.v1_technology_base_get_with_http_info(from, to, value)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1TechnologyBaseGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling NumberBaseConversionApi->v1_technology_base_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **from** | **Integer** | Source base (2, 8, 10, or 16) |  |
| **to** | **Integer** | Target base (2, 8, 10, or 16) |  |
| **value** | **String** | The number as a string. Accepts optional prefixes: 0x (hex), 0b (binary), 0o (octal). |  |

### Return type

[**V1TechnologyBaseGet200Response**](V1TechnologyBaseGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

