# OpenapiClient::CryptoApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_finance_crypto_symbol_get**](CryptoApi.md#v1_finance_crypto_symbol_get) | **GET** /v1/finance/crypto/{symbol} | Get Crypto Price |


## v1_finance_crypto_symbol_get

> <V1FinanceCryptoSymbolGet200Response> v1_finance_crypto_symbol_get(symbol)

Get Crypto Price

Returns current price data for the given cryptocurrency symbol.

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

api_instance = OpenapiClient::CryptoApi.new
symbol = 'BTC' # String | Uppercase ticker symbol (e.g. BTC, ETH, SOL)

begin
  # Get Crypto Price
  result = api_instance.v1_finance_crypto_symbol_get(symbol)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling CryptoApi->v1_finance_crypto_symbol_get: #{e}"
end
```

#### Using the v1_finance_crypto_symbol_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1FinanceCryptoSymbolGet200Response>, Integer, Hash)> v1_finance_crypto_symbol_get_with_http_info(symbol)

```ruby
begin
  # Get Crypto Price
  data, status_code, headers = api_instance.v1_finance_crypto_symbol_get_with_http_info(symbol)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1FinanceCryptoSymbolGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling CryptoApi->v1_finance_crypto_symbol_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **symbol** | **String** | Uppercase ticker symbol (e.g. BTC, ETH, SOL) |  |

### Return type

[**V1FinanceCryptoSymbolGet200Response**](V1FinanceCryptoSymbolGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

