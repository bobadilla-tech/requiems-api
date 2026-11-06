# OpenapiClient::V1TechnologyPasswordGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **password** | **String** | The generated password | [optional] |
| **length** | **Integer** | Length of the generated password | [optional] |
| **strength** | **String** | Password strength assessment (weak, medium, or strong) | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TechnologyPasswordGet200ResponseData.new(
  password: null,
  length: null,
  strength: null
)
```

