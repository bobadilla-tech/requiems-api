# OpenapiClient::V1EntertainmentTriviaBatchPostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **filters** | **Array&lt;Object&gt;** | Array of filters to get (min: 1, max: 50). |  |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1EntertainmentTriviaBatchPostRequest.new(
  filters: [{&quot;category&quot;: &quot;sports&quot;, &quot;difficulty&quot;: &quot;easy&quot;},{&quot;category&quot;: &quot;history&quot;, &quot;difficulty&quot;: &quot;medium&quot;}]

)
```

