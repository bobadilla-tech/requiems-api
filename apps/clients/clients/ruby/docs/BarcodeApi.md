# OpenapiClient::BarcodeApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_technology_barcode_base64_get**](BarcodeApi.md#v1_technology_barcode_base64_get) | **GET** /v1/technology/barcode/base64 | Generate Barcode (Base64 JSON) |
| [**v1_technology_barcode_get**](BarcodeApi.md#v1_technology_barcode_get) | **GET** /v1/technology/barcode | Generate Barcode (PNG) |


## v1_technology_barcode_base64_get

> <V1TechnologyBarcodeBase64Get200Response> v1_technology_barcode_base64_get(data, type)

Generate Barcode (Base64 JSON)

Returns a JSON envelope containing the barcode as a base64-encoded PNG string, along with its type and dimensions.

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

api_instance = OpenapiClient::BarcodeApi.new
data = '123456789' # String | The text or numeric string to encode in the barcode
type = 'code128' # String | Barcode format: code128, code93, code39, ean8, ean13

begin
  # Generate Barcode (Base64 JSON)
  result = api_instance.v1_technology_barcode_base64_get(data, type)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling BarcodeApi->v1_technology_barcode_base64_get: #{e}"
end
```

#### Using the v1_technology_barcode_base64_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1TechnologyBarcodeBase64Get200Response>, Integer, Hash)> v1_technology_barcode_base64_get_with_http_info(data, type)

```ruby
begin
  # Generate Barcode (Base64 JSON)
  data, status_code, headers = api_instance.v1_technology_barcode_base64_get_with_http_info(data, type)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1TechnologyBarcodeBase64Get200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling BarcodeApi->v1_technology_barcode_base64_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **data** | **String** | The text or numeric string to encode in the barcode |  |
| **type** | **String** | Barcode format: code128, code93, code39, ean8, ean13 |  |

### Return type

[**V1TechnologyBarcodeBase64Get200Response**](V1TechnologyBarcodeBase64Get200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## v1_technology_barcode_get

> v1_technology_barcode_get(data, type)

Generate Barcode (PNG)

Returns a raw PNG image of the barcode. Ideal for direct embedding or file download.

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

api_instance = OpenapiClient::BarcodeApi.new
data = '123456789' # String | The text or numeric string to encode in the barcode
type = 'code128' # String | Barcode format: code128, code93, code39, ean8, ean13

begin
  # Generate Barcode (PNG)
  api_instance.v1_technology_barcode_get(data, type)
rescue OpenapiClient::ApiError => e
  puts "Error when calling BarcodeApi->v1_technology_barcode_get: #{e}"
end
```

#### Using the v1_technology_barcode_get_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> v1_technology_barcode_get_with_http_info(data, type)

```ruby
begin
  # Generate Barcode (PNG)
  data, status_code, headers = api_instance.v1_technology_barcode_get_with_http_info(data, type)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling BarcodeApi->v1_technology_barcode_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **data** | **String** | The text or numeric string to encode in the barcode |  |
| **type** | **String** | Barcode format: code128, code93, code39, ean8, ean13 |  |

### Return type

nil (empty response body)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

