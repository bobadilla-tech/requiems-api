# OpenapiClient::V1TechnologyQrBase64BatchPostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **items** | **String** | List of QR generation requests (1–50 items). |  |
| **items_data** | **String** | The text or URL to encode. |  |
| **items_size** | **Integer** | Image size in pixels (50–1000). Defaults to 256. | [optional] |
| **items_recovery** | **String** | Error-correction level: low, medium, high, highest. Defaults to medium. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TechnologyQrBase64BatchPostRequest.new(
  items: null,
  items_data: null,
  items_size: null,
  items_recovery: null
)
```

