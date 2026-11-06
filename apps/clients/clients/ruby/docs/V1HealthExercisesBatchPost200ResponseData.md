# OpenapiClient::V1HealthExercisesBatchPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **results** | **Array&lt;Object&gt;** | Exercises found for the given IDs, in input order. IDs not found are omitted. | [optional] |
| **results_id** | **Integer** | Unique exercise identifier | [optional] |
| **results_name** | **String** | Exercise name | [optional] |
| **results_body_parts** | **Array&lt;Object&gt;** | Body part categories involved | [optional] |
| **results_equipment** | **Array&lt;Object&gt;** | Equipment required | [optional] |
| **results_target_muscles** | **Array&lt;Object&gt;** | Primary muscles targeted | [optional] |
| **results_secondary_muscles** | **Array&lt;Object&gt;** | Secondary muscles engaged | [optional] |
| **results_instructions** | **Array&lt;Object&gt;** | Ordered step-by-step instructions | [optional] |
| **total** | **Integer** | Number of exercises returned | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1HealthExercisesBatchPost200ResponseData.new(
  results: null,
  results_id: null,
  results_name: null,
  results_body_parts: null,
  results_equipment: null,
  results_target_muscles: null,
  results_secondary_muscles: null,
  results_instructions: null,
  total: null
)
```

