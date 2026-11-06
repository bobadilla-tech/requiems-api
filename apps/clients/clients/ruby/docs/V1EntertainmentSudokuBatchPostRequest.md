# OpenapiClient::V1EntertainmentSudokuBatchPostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **puzzles** | **Array&lt;Object&gt;** | Array of difficulty levels to generate (min: 1, max: 20). Each must be one of: easy, medium, hard. |  |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1EntertainmentSudokuBatchPostRequest.new(
  puzzles: [&quot;easy&quot;, &quot;medium&quot;, &quot;hard&quot;]
)
```

