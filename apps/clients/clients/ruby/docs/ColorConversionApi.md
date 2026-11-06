# OpenapiClient::ColorConversionApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_technology_color_get**](ColorConversionApi.md#v1_technology_color_get) | **GET** /v1/technology/color | Convert Color |


## v1_technology_color_get

> <V1TechnologyColorGet200Response> v1_technology_color_get(from, to, value)

Convert Color

Convert a color value from one format to another. The response always includes all four formats.

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

api_instance = OpenapiClient::ColorConversionApi.new
from = 'hex' # String | Source color format: hex, rgb, hsl, or cmyk
to = 'hsl' # String | Target color format: hex, rgb, hsl, or cmyk
value = '#ff5733' # String | Color value in the source format (e.g. #ff5733, rgb(255,87,51), hsl(11,100%,60%), cmyk(0%,66%,80%,0%))

begin
  # Convert Color
  result = api_instance.v1_technology_color_get(from, to, value)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling ColorConversionApi->v1_technology_color_get: #{e}"
end
```

#### Using the v1_technology_color_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1TechnologyColorGet200Response>, Integer, Hash)> v1_technology_color_get_with_http_info(from, to, value)

```ruby
begin
  # Convert Color
  data, status_code, headers = api_instance.v1_technology_color_get_with_http_info(from, to, value)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1TechnologyColorGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling ColorConversionApi->v1_technology_color_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **from** | **String** | Source color format: hex, rgb, hsl, or cmyk |  |
| **to** | **String** | Target color format: hex, rgb, hsl, or cmyk |  |
| **value** | **String** | Color value in the source format (e.g. #ff5733, rgb(255,87,51), hsl(11,100%,60%), cmyk(0%,66%,80%,0%)) |  |

### Return type

[**V1TechnologyColorGet200Response**](V1TechnologyColorGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

