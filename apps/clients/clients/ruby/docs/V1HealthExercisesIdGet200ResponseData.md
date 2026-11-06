# OpenapiClient::V1HealthExercisesIdGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** | Unique exercise identifier | [optional] |
| **name** | **String** | Exercise name | [optional] |
| **body_parts** | **Array&lt;Object&gt;** | Body part categories involved | [optional] |
| **equipment** | **Array&lt;Object&gt;** | Equipment required | [optional] |
| **target_muscles** | **Array&lt;Object&gt;** | Primary muscles targeted | [optional] |
| **secondary_muscles** | **Array&lt;Object&gt;** | Secondary muscles engaged | [optional] |
| **instructions** | **Array&lt;Object&gt;** | Ordered step-by-step instructions | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1HealthExercisesIdGet200ResponseData.new(
  id: null,
  name: null,
  body_parts: null,
  equipment: null,
  target_muscles: null,
  secondary_muscles: null,
  instructions: null
)
```

