# OpenapiClient::BinLookupApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_finance_bin_bin_get**](BinLookupApi.md#v1_finance_bin_bin_get) | **GET** /v1/finance/bin/{bin} | BIN Lookup |


## v1_finance_bin_bin_get

> <V1FinanceBinBinGet200Response> v1_finance_bin_bin_get(bin)

BIN Lookup

Returns card metadata for the given 6–8 digit BIN prefix.

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

api_instance = OpenapiClient::BinLookupApi.new
bin = '424242' # String | 6–8 digit Bank Identification Number. Dashes and spaces are stripped automatically.

begin
  # BIN Lookup
  result = api_instance.v1_finance_bin_bin_get(bin)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling BinLookupApi->v1_finance_bin_bin_get: #{e}"
end
```

#### Using the v1_finance_bin_bin_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1FinanceBinBinGet200Response>, Integer, Hash)> v1_finance_bin_bin_get_with_http_info(bin)

```ruby
begin
  # BIN Lookup
  data, status_code, headers = api_instance.v1_finance_bin_bin_get_with_http_info(bin)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1FinanceBinBinGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling BinLookupApi->v1_finance_bin_bin_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **bin** | **String** | 6–8 digit Bank Identification Number. Dashes and spaces are stripped automatically. |  |

### Return type

[**V1FinanceBinBinGet200Response**](V1FinanceBinBinGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

