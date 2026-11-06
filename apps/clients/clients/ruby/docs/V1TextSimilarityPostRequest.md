# OpenapiClient::V1TextSimilarityPostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **text1** | **String** | The first text to compare. |  |
| **text2** | **String** | The second text to compare. |  |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TextSimilarityPostRequest.new(
  text1: The cat sat on the mat,
  text2: A cat was sitting on a mat
)
```

