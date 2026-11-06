# OpenapiClient::V1TextNormalizeBatchPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **results** | **Array&lt;Object&gt;** | One normalization result per input email, in order | [optional] |
| **total** | **Integer** | Number of emails in the batch (same as results length) | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TextNormalizeBatchPost200ResponseData.new(
  results: null,
  total: null
)
```

