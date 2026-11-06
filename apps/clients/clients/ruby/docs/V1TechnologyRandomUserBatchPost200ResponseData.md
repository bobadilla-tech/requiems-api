# OpenapiClient::V1TechnologyRandomUserBatchPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **results** | **Array&lt;Object&gt;** | Ordered list of generated user profiles. Same fields as the single-user endpoint. | [optional] |
| **total** | **Integer** | Number of users returned. Matches the requested count. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TechnologyRandomUserBatchPost200ResponseData.new(
  results: null,
  total: null
)
```

