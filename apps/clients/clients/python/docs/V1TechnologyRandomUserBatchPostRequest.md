# V1TechnologyRandomUserBatchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**count** | **int** | Number of users to generate (min: 1, max: 50). | 

## Example

```python
from openapi_client.models.v1_technology_random_user_batch_post_request import V1TechnologyRandomUserBatchPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1TechnologyRandomUserBatchPostRequest from a JSON string
v1_technology_random_user_batch_post_request_instance = V1TechnologyRandomUserBatchPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1TechnologyRandomUserBatchPostRequest.to_json())

# convert the object into a dict
v1_technology_random_user_batch_post_request_dict = v1_technology_random_user_batch_post_request_instance.to_dict()
# create an instance of V1TechnologyRandomUserBatchPostRequest from a dict
v1_technology_random_user_batch_post_request_from_dict = V1TechnologyRandomUserBatchPostRequest.from_dict(v1_technology_random_user_batch_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


