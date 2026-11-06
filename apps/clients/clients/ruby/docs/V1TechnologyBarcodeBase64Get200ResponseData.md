# OpenapiClient::V1TechnologyBarcodeBase64Get200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **image** | **String** | Base64-encoded PNG image data | [optional] |
| **type** | **String** | The barcode format that was used | [optional] |
| **width** | **Integer** | Width of the generated image in pixels | [optional] |
| **height** | **Integer** | Height of the generated image in pixels | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TechnologyBarcodeBase64Get200ResponseData.new(
  image: null,
  type: null,
  width: null,
  height: null
)
```

