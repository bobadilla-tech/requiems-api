# OpenapiClient::InflationApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_finance_inflation_batch_post**](InflationApi.md#v1_finance_inflation_batch_post) | **POST** /v1/finance/inflation/batch | Batch Inflation Rates |
| [**v1_finance_inflation_get**](InflationApi.md#v1_finance_inflation_get) | **GET** /v1/finance/inflation | Get Inflation Rate |


## v1_finance_inflation_batch_post

> <V1FinanceInflationBatchPost200Response> v1_finance_inflation_batch_post(v1_finance_inflation_batch_post_request)

Batch Inflation Rates

Returns inflation data for up to 50 countries in a single request. Results are in the same order as the input. Countries with no data return found: false instead of failing the whole request. Billing: 1 credit per country (not per HTTP request).

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

api_instance = OpenapiClient::InflationApi.new
v1_finance_inflation_batch_post_request = OpenapiClient::V1FinanceInflationBatchPostRequest.new({countries: ["US",  "AR",  "DE"]}) # V1FinanceInflationBatchPostRequest | 

begin
  # Batch Inflation Rates
  result = api_instance.v1_finance_inflation_batch_post(v1_finance_inflation_batch_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling InflationApi->v1_finance_inflation_batch_post: #{e}"
end
```

#### Using the v1_finance_inflation_batch_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1FinanceInflationBatchPost200Response>, Integer, Hash)> v1_finance_inflation_batch_post_with_http_info(v1_finance_inflation_batch_post_request)

```ruby
begin
  # Batch Inflation Rates
  data, status_code, headers = api_instance.v1_finance_inflation_batch_post_with_http_info(v1_finance_inflation_batch_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1FinanceInflationBatchPost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling InflationApi->v1_finance_inflation_batch_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_finance_inflation_batch_post_request** | [**V1FinanceInflationBatchPostRequest**](V1FinanceInflationBatchPostRequest.md) |  |  |

### Return type

[**V1FinanceInflationBatchPost200Response**](V1FinanceInflationBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## v1_finance_inflation_get

> <V1FinanceInflationGet200Response> v1_finance_inflation_get(country)

Get Inflation Rate

Returns the latest annual CPI inflation rate for a country plus the previous 10 years of historical data.

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

api_instance = OpenapiClient::InflationApi.new
country = 'US' # String | ISO 3166-1 alpha-2 country code (e.g. US, GB, DE). Case-insensitive.

begin
  # Get Inflation Rate
  result = api_instance.v1_finance_inflation_get(country)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling InflationApi->v1_finance_inflation_get: #{e}"
end
```

#### Using the v1_finance_inflation_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1FinanceInflationGet200Response>, Integer, Hash)> v1_finance_inflation_get_with_http_info(country)

```ruby
begin
  # Get Inflation Rate
  data, status_code, headers = api_instance.v1_finance_inflation_get_with_http_info(country)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1FinanceInflationGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling InflationApi->v1_finance_inflation_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **country** | **String** | ISO 3166-1 alpha-2 country code (e.g. US, GB, DE). Case-insensitive. |  |

### Return type

[**V1FinanceInflationGet200Response**](V1FinanceInflationGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

