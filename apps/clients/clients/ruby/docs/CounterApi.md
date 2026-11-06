# OpenapiClient::CounterApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_technology_counter_namespace_get**](CounterApi.md#v1_technology_counter_namespace_get) | **GET** /v1/technology/counter/{namespace} | Get Counter Value |
| [**v1_technology_counter_namespace_post**](CounterApi.md#v1_technology_counter_namespace_post) | **POST** /v1/technology/counter/{namespace} | Increment Counter |


## v1_technology_counter_namespace_get

> <V1TechnologyCounterNamespaceGet200Response> v1_technology_counter_namespace_get(namespace)

Get Counter Value

Get the current value of a counter without incrementing it

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

api_instance = OpenapiClient::CounterApi.new
namespace = 'page-views' # String | Counter namespace (1-64 chars: alphanumeric, hyphen, underscore)

begin
  # Get Counter Value
  result = api_instance.v1_technology_counter_namespace_get(namespace)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling CounterApi->v1_technology_counter_namespace_get: #{e}"
end
```

#### Using the v1_technology_counter_namespace_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1TechnologyCounterNamespaceGet200Response>, Integer, Hash)> v1_technology_counter_namespace_get_with_http_info(namespace)

```ruby
begin
  # Get Counter Value
  data, status_code, headers = api_instance.v1_technology_counter_namespace_get_with_http_info(namespace)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1TechnologyCounterNamespaceGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling CounterApi->v1_technology_counter_namespace_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **namespace** | **String** | Counter namespace (1-64 chars: alphanumeric, hyphen, underscore) |  |

### Return type

[**V1TechnologyCounterNamespaceGet200Response**](V1TechnologyCounterNamespaceGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## v1_technology_counter_namespace_post

> <V1TechnologyCounterNamespacePost200Response> v1_technology_counter_namespace_post(namespace)

Increment Counter

Atomically increment a counter in the specified namespace and return the new value

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

api_instance = OpenapiClient::CounterApi.new
namespace = 'page-views' # String | Counter namespace (1-64 chars: alphanumeric, hyphen, underscore)

begin
  # Increment Counter
  result = api_instance.v1_technology_counter_namespace_post(namespace)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling CounterApi->v1_technology_counter_namespace_post: #{e}"
end
```

#### Using the v1_technology_counter_namespace_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1TechnologyCounterNamespacePost200Response>, Integer, Hash)> v1_technology_counter_namespace_post_with_http_info(namespace)

```ruby
begin
  # Increment Counter
  data, status_code, headers = api_instance.v1_technology_counter_namespace_post_with_http_info(namespace)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1TechnologyCounterNamespacePost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling CounterApi->v1_technology_counter_namespace_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **namespace** | **String** | Counter namespace (1-64 chars: alphanumeric, hyphen, underscore) |  |

### Return type

[**V1TechnologyCounterNamespacePost200Response**](V1TechnologyCounterNamespacePost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

