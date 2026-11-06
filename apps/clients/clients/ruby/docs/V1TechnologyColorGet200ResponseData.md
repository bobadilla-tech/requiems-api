# OpenapiClient::V1TechnologyColorGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **input** | **String** | The original value passed in the value parameter | [optional] |
| **result** | **String** | The color expressed in the requested to format | [optional] |
| **formats_hex** | **String** | HEX representation (#rrggbb) | [optional] |
| **formats_rgb** | **String** | RGB representation (rgb(r, g, b)) | [optional] |
| **formats_hsl** | **String** | HSL representation (hsl(h, s%, l%)) | [optional] |
| **formats_cmyk** | **String** | CMYK representation (cmyk(c%, m%, y%, k%)) | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TechnologyColorGet200ResponseData.new(
  input: null,
  result: null,
  formats_hex: null,
  formats_rgb: null,
  formats_hsl: null,
  formats_cmyk: null
)
```

