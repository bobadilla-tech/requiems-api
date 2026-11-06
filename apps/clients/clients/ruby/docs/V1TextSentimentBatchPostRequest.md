# OpenapiClient::V1TextSentimentBatchPostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **texts** | **Array&lt;Object&gt;** | The list of texts to analyze. Between 1 and 50 items. |  |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TextSentimentBatchPostRequest.new(
  texts: [&quot;I love this!&quot;, &quot;This is terrible.&quot;, &quot;The table is there.&quot;]
)
```

