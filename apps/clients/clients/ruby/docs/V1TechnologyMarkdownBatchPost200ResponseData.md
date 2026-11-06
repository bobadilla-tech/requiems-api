# OpenapiClient::V1TechnologyMarkdownBatchPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **results** | **Array&lt;Object&gt;** | One HTML result per Markdown input, preserving input order | [optional] |
| **results_html** | **String** | Rendered HTML output for each Markdown string | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TechnologyMarkdownBatchPost200ResponseData.new(
  results: null,
  results_html: null
)
```

