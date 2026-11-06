# V1TechnologyConvertBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **Array&lt;any&gt;** | List of unit conversion results returned in the same order as the input operations. | [optional] [default to undefined]
**results_from** | **string** | Source unit used for the conversion. | [optional] [default to undefined]
**results_to** | **string** | Target unit used for the conversion. | [optional] [default to undefined]
**results_success** | **boolean** | Indicates whether the conversion operation was successful. | [optional] [default to undefined]
**results_error** | **string** | Error message if the conversion failed. Empty when success is true. | [optional] [default to undefined]
**results_data** | **object** | Conversion result details. Contains empty values when the conversion fails. | [optional] [default to undefined]
**results_data_from** | **string** | Source unit. | [optional] [default to undefined]
**results_data_to** | **string** | Target unit. | [optional] [default to undefined]
**results_data_input** | **number** | Original numeric value provided for conversion. | [optional] [default to undefined]
**results_data_result** | **number** | Converted numeric result. | [optional] [default to undefined]
**results_data_formula** | **string** | Formula or conversion factor used to calculate the result. | [optional] [default to undefined]
**total** | **number** | Total number of conversion results returned. | [optional] [default to undefined]

## Example

```typescript
import { V1TechnologyConvertBatchPost200ResponseData } from 'api-client';

const instance: V1TechnologyConvertBatchPost200ResponseData = {
    results,
    results_from,
    results_to,
    results_success,
    results_error,
    results_data,
    results_data_from,
    results_data_to,
    results_data_input,
    results_data_result,
    results_data_formula,
    total,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
