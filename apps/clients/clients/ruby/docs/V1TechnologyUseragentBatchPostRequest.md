# OpenapiClient::V1TechnologyUseragentBatchPostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **user_agents** | **Array&lt;Object&gt;** | Array of user agents to parse (min: 1, max: 50). |  |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TechnologyUseragentBatchPostRequest.new(
  user_agents: [
  &quot;Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36&quot;,
  &quot;Mozilla/5.0 (iPhone; CPU iPhone OS 17_0 like Mac OS X) AppleWebKit/605.1.15&quot;
]

)
```

