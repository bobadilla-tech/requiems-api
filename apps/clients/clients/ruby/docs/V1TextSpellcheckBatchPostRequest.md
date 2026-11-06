# OpenapiClient::V1TextSpellcheckBatchPostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **texts** | **Array&lt;Object&gt;** | The list of texts to spell-check. Between 1 and 50 items. |  |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TextSpellcheckBatchPostRequest.new(
  texts: [&quot;Ths is a tset&quot;, &quot;Smiple example&quot;]
)
```

