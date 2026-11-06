# OpenapiClient::V1SystemsTextNormalizePost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **original** | **String** | The input text as provided | [optional] |
| **normalized** | **String** | The text after all operations have been applied in order | [optional] |
| **operations** | **Array&lt;Object&gt;** | The operations list echoed from the request | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1SystemsTextNormalizePost200ResponseData.new(
  original: null,
  normalized: null,
  operations: null
)
```

