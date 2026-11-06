# V1TechnologyBaseBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **List[object]** | Per-item results in input order. | [optional] 
**results_from** | **int** | The source base used for conversion. | [optional] 
**results_to** | **int** | The target base used for conversion. | [optional] 
**results_input** | **str** | The original value as provided in the request. | [optional] 
**results_result** | **str** | The converted value in the target base (omitted on error). | [optional] 
**results_error** | **str** | Error message if this item failed (omitted on success). | [optional] 
**total** | **int** | Total number of items processed. | [optional] 

## Example

```python
from openapi_client.models.v1_technology_base_batch_post200_response_data import V1TechnologyBaseBatchPost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1TechnologyBaseBatchPost200ResponseData from a JSON string
v1_technology_base_batch_post200_response_data_instance = V1TechnologyBaseBatchPost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1TechnologyBaseBatchPost200ResponseData.to_json())

# convert the object into a dict
v1_technology_base_batch_post200_response_data_dict = v1_technology_base_batch_post200_response_data_instance.to_dict()
# create an instance of V1TechnologyBaseBatchPost200ResponseData from a dict
v1_technology_base_batch_post200_response_data_from_dict = V1TechnologyBaseBatchPost200ResponseData.from_dict(v1_technology_base_batch_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


