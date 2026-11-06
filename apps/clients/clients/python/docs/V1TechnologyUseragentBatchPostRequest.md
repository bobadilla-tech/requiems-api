# V1TechnologyUseragentBatchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**user_agents** | **List[object]** | Array of user agents to parse (min: 1, max: 50). | 

## Example

```python
from openapi_client.models.v1_technology_useragent_batch_post_request import V1TechnologyUseragentBatchPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1TechnologyUseragentBatchPostRequest from a JSON string
v1_technology_useragent_batch_post_request_instance = V1TechnologyUseragentBatchPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1TechnologyUseragentBatchPostRequest.to_json())

# convert the object into a dict
v1_technology_useragent_batch_post_request_dict = v1_technology_useragent_batch_post_request_instance.to_dict()
# create an instance of V1TechnologyUseragentBatchPostRequest from a dict
v1_technology_useragent_batch_post_request_from_dict = V1TechnologyUseragentBatchPostRequest.from_dict(v1_technology_useragent_batch_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


