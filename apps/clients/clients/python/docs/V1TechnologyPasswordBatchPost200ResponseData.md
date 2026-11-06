# V1TechnologyPasswordBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **List[object]** | Per-item results in input order. | [optional] 
**results_result_password** | **str** | The generated password (omitted on error). | [optional] 
**results_result_length** | **int** | Length of the generated password. | [optional] 
**results_result_strength** | **str** | Strength assessment (weak, medium, or strong). | [optional] 
**results_error** | **str** | Error message if generation failed (omitted on success). | [optional] 
**total** | **int** | Total number of items processed. | [optional] 

## Example

```python
from openapi_client.models.v1_technology_password_batch_post200_response_data import V1TechnologyPasswordBatchPost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1TechnologyPasswordBatchPost200ResponseData from a JSON string
v1_technology_password_batch_post200_response_data_instance = V1TechnologyPasswordBatchPost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1TechnologyPasswordBatchPost200ResponseData.to_json())

# convert the object into a dict
v1_technology_password_batch_post200_response_data_dict = v1_technology_password_batch_post200_response_data_instance.to_dict()
# create an instance of V1TechnologyPasswordBatchPost200ResponseData from a dict
v1_technology_password_batch_post200_response_data_from_dict = V1TechnologyPasswordBatchPost200ResponseData.from_dict(v1_technology_password_batch_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


