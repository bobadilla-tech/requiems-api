# OpenapiClient::V1TextLoremBatchPostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **items** | **Array&lt;Object&gt;** | Array of text generation requests. Minimum 1 item, maximum 50 items. |  |
| **items_paragraphs** | **Integer** | Number of paragraphs to generate (1-20) | [optional] |
| **items_sentences** | **Integer** | Number of sentences per paragraph (1-20) | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TextLoremBatchPostRequest.new(
  items: null,
  items_paragraphs: 3,
  items_sentences: 5
)
```

