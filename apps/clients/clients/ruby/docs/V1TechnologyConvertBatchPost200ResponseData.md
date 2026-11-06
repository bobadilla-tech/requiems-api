# OpenapiClient::V1TechnologyConvertBatchPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **results** | **Array&lt;Object&gt;** | List of unit conversion results returned in the same order as the input operations. | [optional] |
| **results_from** | **String** | Source unit used for the conversion. | [optional] |
| **results_to** | **String** | Target unit used for the conversion. | [optional] |
| **results_success** | **Boolean** | Indicates whether the conversion operation was successful. | [optional] |
| **results_error** | **String** | Error message if the conversion failed. Empty when success is true. | [optional] |
| **results_data** | **Object** | Conversion result details. Contains empty values when the conversion fails. | [optional] |
| **results_data_from** | **String** | Source unit. | [optional] |
| **results_data_to** | **String** | Target unit. | [optional] |
| **results_data_input** | **Float** | Original numeric value provided for conversion. | [optional] |
| **results_data_result** | **Float** | Converted numeric result. | [optional] |
| **results_data_formula** | **String** | Formula or conversion factor used to calculate the result. | [optional] |
| **total** | **Integer** | Total number of conversion results returned. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TechnologyConvertBatchPost200ResponseData.new(
  results: null,
  results_from: null,
  results_to: null,
  results_success: null,
  results_error: null,
  results_data: null,
  results_data_from: null,
  results_data_to: null,
  results_data_input: null,
  results_data_result: null,
  results_data_formula: null,
  total: null
)
```

