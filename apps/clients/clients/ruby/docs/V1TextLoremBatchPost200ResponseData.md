# OpenapiClient::V1TextLoremBatchPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **data_data** | **Object** | Contains the generated text if the specific item was successful | [optional] |
| **data_error** | **String** | Error message if the specific item failed to generate | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TextLoremBatchPost200ResponseData.new(
  data_data: null,
  data_error: null
)
```

