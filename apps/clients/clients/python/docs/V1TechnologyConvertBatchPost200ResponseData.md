# V1TechnologyConvertBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **List[object]** | List of unit conversion results returned in the same order as the input operations. | [optional] 
**results_from** | **str** | Source unit used for the conversion. | [optional] 
**results_to** | **str** | Target unit used for the conversion. | [optional] 
**results_success** | **bool** | Indicates whether the conversion operation was successful. | [optional] 
**results_error** | **str** | Error message if the conversion failed. Empty when success is true. | [optional] 
**results_data** | **object** | Conversion result details. Contains empty values when the conversion fails. | [optional] 
**results_data_from** | **str** | Source unit. | [optional] 
**results_data_to** | **str** | Target unit. | [optional] 
**results_data_input** | **float** | Original numeric value provided for conversion. | [optional] 
**results_data_result** | **float** | Converted numeric result. | [optional] 
**results_data_formula** | **str** | Formula or conversion factor used to calculate the result. | [optional] 
**total** | **int** | Total number of conversion results returned. | [optional] 

## Example

```python
from openapi_client.models.v1_technology_convert_batch_post200_response_data import V1TechnologyConvertBatchPost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1TechnologyConvertBatchPost200ResponseData from a JSON string
v1_technology_convert_batch_post200_response_data_instance = V1TechnologyConvertBatchPost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1TechnologyConvertBatchPost200ResponseData.to_json())

# convert the object into a dict
v1_technology_convert_batch_post200_response_data_dict = v1_technology_convert_batch_post200_response_data_instance.to_dict()
# create an instance of V1TechnologyConvertBatchPost200ResponseData from a dict
v1_technology_convert_batch_post200_response_data_from_dict = V1TechnologyConvertBatchPost200ResponseData.from_dict(v1_technology_convert_batch_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


