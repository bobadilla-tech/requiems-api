# OpenapiClient::V1TechnologyMarkdownBatchPostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **markdowns** | **Array&lt;Object&gt;** | Array of Markdown strings. Min: 1, Max: 50. |  |
| **sanitize** | **Boolean** | When true, sanitizes HTML output to remove unsafe tags and attributes. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TechnologyMarkdownBatchPostRequest.new(
  markdowns: [&quot;# Hello&quot;, &quot;**bold**&quot;, &quot;- item one&quot;],
  sanitize: null
)
```

