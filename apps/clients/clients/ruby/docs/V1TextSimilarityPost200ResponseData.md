# OpenapiClient::V1TextSimilarityPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **similarity** | **Float** | Cosine similarity score between the two texts, in the range [0, 1]. | [optional] |
| **method** | **String** | The algorithm used. Currently always &#39;cosine&#39;. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TextSimilarityPost200ResponseData.new(
  similarity: null,
  method: null
)
```

