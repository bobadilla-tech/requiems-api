# OpenapiClient::V1PlacesHolidaysBatchPostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **queries** | **Array&lt;Object&gt;** | Array of (country, year) pairs. Min: 1, Max: 50. |  |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1PlacesHolidaysBatchPostRequest.new(
  queries: [{&quot;country&quot;:&quot;US&quot;,&quot;year&quot;:2025},{&quot;country&quot;:&quot;AR&quot;,&quot;year&quot;:2024}]
)
```

