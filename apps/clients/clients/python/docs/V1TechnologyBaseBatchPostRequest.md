# V1TechnologyBaseBatchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**items** | **str** | List of conversion requests (1–50 items). | 
**items_from** | **int** | Source base (2, 8, 10, or 16). | 
**items_to** | **int** | Target base (2, 8, 10, or 16). | 
**items_value** | **str** | The number as a string. Accepts optional prefixes: 0x (hex), 0b (binary), 0o (octal). | 

## Example

```python
from openapi_client.models.v1_technology_base_batch_post_request import V1TechnologyBaseBatchPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1TechnologyBaseBatchPostRequest from a JSON string
v1_technology_base_batch_post_request_instance = V1TechnologyBaseBatchPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1TechnologyBaseBatchPostRequest.to_json())

# convert the object into a dict
v1_technology_base_batch_post_request_dict = v1_technology_base_batch_post_request_instance.to_dict()
# create an instance of V1TechnologyBaseBatchPostRequest from a dict
v1_technology_base_batch_post_request_from_dict = V1TechnologyBaseBatchPostRequest.from_dict(v1_technology_base_batch_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


