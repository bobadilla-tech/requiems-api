# V1TextNormalizeBatchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**emails** | **List[object]** | Array of addresses to normalize (min 1, max 100; each entry non-empty) | 

## Example

```python
from openapi_client.models.v1_text_normalize_batch_post_request import V1TextNormalizeBatchPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1TextNormalizeBatchPostRequest from a JSON string
v1_text_normalize_batch_post_request_instance = V1TextNormalizeBatchPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1TextNormalizeBatchPostRequest.to_json())

# convert the object into a dict
v1_text_normalize_batch_post_request_dict = v1_text_normalize_batch_post_request_instance.to_dict()
# create an instance of V1TextNormalizeBatchPostRequest from a dict
v1_text_normalize_batch_post_request_from_dict = V1TextNormalizeBatchPostRequest.from_dict(v1_text_normalize_batch_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


