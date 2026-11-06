# OpenapiClient::SwiftCodeApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_finance_swift_code_get**](SwiftCodeApi.md#v1_finance_swift_code_get) | **GET** /v1/finance/swift/{code} | Get SWIFT Code |
| [**v1_finance_swift_get**](SwiftCodeApi.md#v1_finance_swift_get) | **GET** /v1/finance/swift | List SWIFT Codes |


## v1_finance_swift_code_get

> <V1FinanceSwiftCodeGet200Response> v1_finance_swift_code_get(code)

Get SWIFT Code

Look up bank metadata for a SWIFT/BIC code.

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

api_instance = OpenapiClient::SwiftCodeApi.new
code = 'DEUTDEDB' # String | SWIFT/BIC code (8 or 11 alphanumeric characters)

begin
  # Get SWIFT Code
  result = api_instance.v1_finance_swift_code_get(code)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling SwiftCodeApi->v1_finance_swift_code_get: #{e}"
end
```

#### Using the v1_finance_swift_code_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1FinanceSwiftCodeGet200Response>, Integer, Hash)> v1_finance_swift_code_get_with_http_info(code)

```ruby
begin
  # Get SWIFT Code
  data, status_code, headers = api_instance.v1_finance_swift_code_get_with_http_info(code)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1FinanceSwiftCodeGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling SwiftCodeApi->v1_finance_swift_code_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **code** | **String** | SWIFT/BIC code (8 or 11 alphanumeric characters) |  |

### Return type

[**V1FinanceSwiftCodeGet200Response**](V1FinanceSwiftCodeGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## v1_finance_swift_get

> <V1HealthExercisesRandomGet200Response> v1_finance_swift_get(opts)

List SWIFT Codes

List SWIFT records with optional filters and pagination.

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

api_instance = OpenapiClient::SwiftCodeApi.new
opts = {
  country_code: 'DE', # String | Optional 2-letter country code filter (e.g. DE, US)
  bank_code: 'DEUT', # String | Optional 4-letter bank code filter (e.g. DEUT)
  q: 'deutsche', # String | Optional text search across swift_code, bank_name, and city
  limit: 50, # Integer | Max rows to return (default 50, max 200)
  offset: 0 # Integer | Number of rows to skip (default 0)
}

begin
  # List SWIFT Codes
  result = api_instance.v1_finance_swift_get(opts)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling SwiftCodeApi->v1_finance_swift_get: #{e}"
end
```

#### Using the v1_finance_swift_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1HealthExercisesRandomGet200Response>, Integer, Hash)> v1_finance_swift_get_with_http_info(opts)

```ruby
begin
  # List SWIFT Codes
  data, status_code, headers = api_instance.v1_finance_swift_get_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1HealthExercisesRandomGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling SwiftCodeApi->v1_finance_swift_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **country_code** | **String** | Optional 2-letter country code filter (e.g. DE, US) | [optional] |
| **bank_code** | **String** | Optional 4-letter bank code filter (e.g. DEUT) | [optional] |
| **q** | **String** | Optional text search across swift_code, bank_name, and city | [optional] |
| **limit** | **Integer** | Max rows to return (default 50, max 200) | [optional] |
| **offset** | **Integer** | Number of rows to skip (default 0) | [optional] |

### Return type

[**V1HealthExercisesRandomGet200Response**](V1HealthExercisesRandomGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

