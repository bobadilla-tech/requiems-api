# OpenapiClient::V1TechnologyMarkdownPostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **markdown** | **String** | The Markdown text to convert. |  |
| **sanitize** | **Boolean** | When true, sanitizes the HTML output to remove unsafe tags and attributes. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TechnologyMarkdownPostRequest.new(
  markdown: # Hello

This is **bold** and _italic_ text.,
  sanitize: null
)
```

