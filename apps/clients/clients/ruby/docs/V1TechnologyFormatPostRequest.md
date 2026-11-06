# OpenapiClient::V1TechnologyFormatPostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **from** | **String** | Source format. One of: json, yaml, csv, xml, toml |  |
| **to** | **String** | Target format. One of: json, yaml, csv, xml, toml |  |
| **content** | **String** | The content to convert, serialized as a string in the source format. |  |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TechnologyFormatPostRequest.new(
  from: json,
  to: yaml,
  content: {&quot;name&quot;:&quot;Alice&quot;,&quot;age&quot;:30}
)
```

