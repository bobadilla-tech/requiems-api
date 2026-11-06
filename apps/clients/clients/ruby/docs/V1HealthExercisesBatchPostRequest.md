# OpenapiClient::V1HealthExercisesBatchPostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **ids** | **Array&lt;Object&gt;** | Array of numeric exercise IDs to fetch. Must contain between 1 and 50 items. Each ID must be a positive integer. |  |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1HealthExercisesBatchPostRequest.new(
  ids: [1, 7, 42]
)
```

