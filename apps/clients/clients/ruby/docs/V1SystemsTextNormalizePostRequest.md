# OpenapiClient::V1SystemsTextNormalizePostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **text** | **String** | The text to normalize. Required. |  |
| **operations** | **Array&lt;Object&gt;** | Ordered list of normalization operations to apply. At least one required. Valid values: trim, lowercase, uppercase, normalize_unicode, collapse_whitespace, strip_html, remove_punctuation |  |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1SystemsTextNormalizePostRequest.new(
  text:   Hello,   World!  ,
  operations: [&quot;trim&quot;, &quot;lowercase&quot;, &quot;collapse_whitespace&quot;]
)
```

