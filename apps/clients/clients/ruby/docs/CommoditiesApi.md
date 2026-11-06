# OpenapiClient::CommoditiesApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_finance_commodities_commodity_get**](CommoditiesApi.md#v1_finance_commodities_commodity_get) | **GET** /v1/finance/commodities/{commodity} | Get Commodity Price |


## v1_finance_commodities_commodity_get

> <V1FinanceCommoditiesCommodityGet200Response> v1_finance_commodities_commodity_get(commodity)

Get Commodity Price

Returns the latest annual average price and up to 10 years of historical data for the requested commodity slug.

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

api_instance = OpenapiClient::CommoditiesApi.new
commodity = 'gold' # String | Commodity slug (e.g. gold, silver, oil). See supported slugs below.

begin
  # Get Commodity Price
  result = api_instance.v1_finance_commodities_commodity_get(commodity)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling CommoditiesApi->v1_finance_commodities_commodity_get: #{e}"
end
```

#### Using the v1_finance_commodities_commodity_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1FinanceCommoditiesCommodityGet200Response>, Integer, Hash)> v1_finance_commodities_commodity_get_with_http_info(commodity)

```ruby
begin
  # Get Commodity Price
  data, status_code, headers = api_instance.v1_finance_commodities_commodity_get_with_http_info(commodity)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1FinanceCommoditiesCommodityGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling CommoditiesApi->v1_finance_commodities_commodity_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **commodity** | **String** | Commodity slug (e.g. gold, silver, oil). See supported slugs below. |  |

### Return type

[**V1FinanceCommoditiesCommodityGet200Response**](V1FinanceCommoditiesCommodityGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

