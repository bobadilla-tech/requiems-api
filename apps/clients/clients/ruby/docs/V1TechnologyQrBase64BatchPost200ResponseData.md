# OpenapiClient::V1TechnologyQrBase64BatchPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **results** | **Array&lt;Object&gt;** | Per-item results in input order. | [optional] |
| **results_data** | **String** | The input text or URL that was encoded. | [optional] |
| **results_image** | **String** | Base64-encoded PNG image data (omitted on error). | [optional] |
| **results_width** | **Integer** | Width of the generated image in pixels (omitted on error). | [optional] |
| **results_height** | **Integer** | Height of the generated image in pixels (omitted on error). | [optional] |
| **results_error** | **String** | Error message if generation failed (omitted on success). | [optional] |
| **total** | **Integer** | Total number of items processed. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TechnologyQrBase64BatchPost200ResponseData.new(
  results: null,
  results_data: null,
  results_image: null,
  results_width: null,
  results_height: null,
  results_error: null,
  total: null
)
```

