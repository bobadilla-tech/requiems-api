# V1TechnologyRandomUserBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **List[object]** | Ordered list of generated user profiles. Same fields as the single-user endpoint. | [optional] 
**total** | **int** | Number of users returned. Matches the requested count. | [optional] 

## Example

```python
from openapi_client.models.v1_technology_random_user_batch_post200_response_data import V1TechnologyRandomUserBatchPost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1TechnologyRandomUserBatchPost200ResponseData from a JSON string
v1_technology_random_user_batch_post200_response_data_instance = V1TechnologyRandomUserBatchPost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1TechnologyRandomUserBatchPost200ResponseData.to_json())

# convert the object into a dict
v1_technology_random_user_batch_post200_response_data_dict = v1_technology_random_user_batch_post200_response_data_instance.to_dict()
# create an instance of V1TechnologyRandomUserBatchPost200ResponseData from a dict
v1_technology_random_user_batch_post200_response_data_from_dict = V1TechnologyRandomUserBatchPost200ResponseData.from_dict(v1_technology_random_user_batch_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


