# OpenapiClient::ExchangeRateApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_finance_convert_get**](ExchangeRateApi.md#v1_finance_convert_get) | **GET** /v1/finance/convert | Convert Currency |
| [**v1_finance_exchange_rate_get**](ExchangeRateApi.md#v1_finance_exchange_rate_get) | **GET** /v1/finance/exchange-rate | Get Exchange Rate |


## v1_finance_convert_get

> <V1FinanceConvertGet200Response> v1_finance_convert_get(from, to, amount)

Convert Currency

Converts an amount from one currency to another and returns the rate alongside the converted value.

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

api_instance = OpenapiClient::ExchangeRateApi.new
from = 'USD' # String | ISO 4217 source currency code (3 letters, e.g. USD)
to = 'EUR' # String | ISO 4217 target currency code (3 letters, e.g. EUR)
amount = 100 # Float | Amount to convert. Must be greater than 0.

begin
  # Convert Currency
  result = api_instance.v1_finance_convert_get(from, to, amount)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling ExchangeRateApi->v1_finance_convert_get: #{e}"
end
```

#### Using the v1_finance_convert_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1FinanceConvertGet200Response>, Integer, Hash)> v1_finance_convert_get_with_http_info(from, to, amount)

```ruby
begin
  # Convert Currency
  data, status_code, headers = api_instance.v1_finance_convert_get_with_http_info(from, to, amount)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1FinanceConvertGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling ExchangeRateApi->v1_finance_convert_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **from** | **String** | ISO 4217 source currency code (3 letters, e.g. USD) |  |
| **to** | **String** | ISO 4217 target currency code (3 letters, e.g. EUR) |  |
| **amount** | **Float** | Amount to convert. Must be greater than 0. |  |

### Return type

[**V1FinanceConvertGet200Response**](V1FinanceConvertGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## v1_finance_exchange_rate_get

> <V1FinanceExchangeRateGet200Response> v1_finance_exchange_rate_get(from, to)

Get Exchange Rate

Returns the current exchange rate between two currencies.

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

api_instance = OpenapiClient::ExchangeRateApi.new
from = 'USD' # String | ISO 4217 source currency code (3 letters, e.g. USD)
to = 'EUR' # String | ISO 4217 target currency code (3 letters, e.g. EUR)

begin
  # Get Exchange Rate
  result = api_instance.v1_finance_exchange_rate_get(from, to)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling ExchangeRateApi->v1_finance_exchange_rate_get: #{e}"
end
```

#### Using the v1_finance_exchange_rate_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1FinanceExchangeRateGet200Response>, Integer, Hash)> v1_finance_exchange_rate_get_with_http_info(from, to)

```ruby
begin
  # Get Exchange Rate
  data, status_code, headers = api_instance.v1_finance_exchange_rate_get_with_http_info(from, to)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1FinanceExchangeRateGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling ExchangeRateApi->v1_finance_exchange_rate_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **from** | **String** | ISO 4217 source currency code (3 letters, e.g. USD) |  |
| **to** | **String** | ISO 4217 target currency code (3 letters, e.g. EUR) |  |

### Return type

[**V1FinanceExchangeRateGet200Response**](V1FinanceExchangeRateGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

