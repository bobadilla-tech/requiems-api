# OpenapiClient::UnitConversionApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_technology_convert_get**](UnitConversionApi.md#v1_technology_convert_get) | **GET** /v1/technology/convert | Convert Units |
| [**v1_technology_convert_units_get**](UnitConversionApi.md#v1_technology_convert_units_get) | **GET** /v1/technology/convert/units | List Available Units |


## v1_technology_convert_get

> <V1TechnologyConvertGet200Response> v1_technology_convert_get(from, to, value)

Convert Units

Convert a value from one unit to another

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

api_instance = OpenapiClient::UnitConversionApi.new
from = 'miles' # String | Source unit key (e.g. miles, kg, c)
to = 'km' # String | Target unit key (e.g. km, lb, f)
value = 10 # Float | Numeric value to convert

begin
  # Convert Units
  result = api_instance.v1_technology_convert_get(from, to, value)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling UnitConversionApi->v1_technology_convert_get: #{e}"
end
```

#### Using the v1_technology_convert_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1TechnologyConvertGet200Response>, Integer, Hash)> v1_technology_convert_get_with_http_info(from, to, value)

```ruby
begin
  # Convert Units
  data, status_code, headers = api_instance.v1_technology_convert_get_with_http_info(from, to, value)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1TechnologyConvertGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling UnitConversionApi->v1_technology_convert_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **from** | **String** | Source unit key (e.g. miles, kg, c) |  |
| **to** | **String** | Target unit key (e.g. km, lb, f) |  |
| **value** | **Float** | Numeric value to convert |  |

### Return type

[**V1TechnologyConvertGet200Response**](V1TechnologyConvertGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## v1_technology_convert_units_get

> <V1TechnologyConvertUnitsGet200Response> v1_technology_convert_units_get

List Available Units

Returns all available unit conversion types grouped by measurement category

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

api_instance = OpenapiClient::UnitConversionApi.new

begin
  # List Available Units
  result = api_instance.v1_technology_convert_units_get
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling UnitConversionApi->v1_technology_convert_units_get: #{e}"
end
```

#### Using the v1_technology_convert_units_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1TechnologyConvertUnitsGet200Response>, Integer, Hash)> v1_technology_convert_units_get_with_http_info

```ruby
begin
  # List Available Units
  data, status_code, headers = api_instance.v1_technology_convert_units_get_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1TechnologyConvertUnitsGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling UnitConversionApi->v1_technology_convert_units_get_with_http_info: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**V1TechnologyConvertUnitsGet200Response**](V1TechnologyConvertUnitsGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

