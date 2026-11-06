# OpenapiClient::V1TechnologyConvertGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **from** | **String** | Source unit key | [optional] |
| **to** | **String** | Target unit key | [optional] |
| **input** | **Float** | The original input value | [optional] |
| **result** | **Float** | The converted value (rounded to 6 decimal places) | [optional] |
| **formula** | **String** | Human-readable conversion formula | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TechnologyConvertGet200ResponseData.new(
  from: null,
  to: null,
  input: null,
  result: null,
  formula: null
)
```

