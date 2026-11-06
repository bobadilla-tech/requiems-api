# OpenapiClient::V1TechnologyBaseBatchPostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **items** | **String** | List of conversion requests (1–50 items). |  |
| **items_from** | **Integer** | Source base (2, 8, 10, or 16). |  |
| **items_to** | **Integer** | Target base (2, 8, 10, or 16). |  |
| **items_value** | **String** | The number as a string. Accepts optional prefixes: 0x (hex), 0b (binary), 0o (octal). |  |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TechnologyBaseBatchPostRequest.new(
  items: null,
  items_from: null,
  items_to: null,
  items_value: null
)
```

