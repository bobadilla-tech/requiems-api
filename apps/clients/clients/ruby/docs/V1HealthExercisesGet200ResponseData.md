# OpenapiClient::V1HealthExercisesGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **items** | **Array&lt;Object&gt;** | Array of exercise objects for the current page | [optional] |
| **items_id** | **Integer** | Unique exercise identifier | [optional] |
| **items_name** | **String** | Exercise name | [optional] |
| **items_body_parts** | **Array&lt;Object&gt;** | Body part categories involved | [optional] |
| **items_equipment** | **Array&lt;Object&gt;** | Equipment required | [optional] |
| **items_target_muscles** | **Array&lt;Object&gt;** | Primary muscles targeted | [optional] |
| **items_secondary_muscles** | **Array&lt;Object&gt;** | Secondary muscles engaged | [optional] |
| **items_instructions** | **Array&lt;Object&gt;** | Ordered step-by-step instructions | [optional] |
| **total** | **Integer** | Total number of exercises matching the filters | [optional] |
| **page** | **Integer** | Current page number | [optional] |
| **per_page** | **Integer** | Number of results per page | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1HealthExercisesGet200ResponseData.new(
  items: null,
  items_id: null,
  items_name: null,
  items_body_parts: null,
  items_equipment: null,
  items_target_muscles: null,
  items_secondary_muscles: null,
  items_instructions: null,
  total: null,
  page: null,
  per_page: null
)
```

