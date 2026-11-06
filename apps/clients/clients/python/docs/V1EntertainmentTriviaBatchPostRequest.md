# V1EntertainmentTriviaBatchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**filters** | **List[object]** | Array of filters to get (min: 1, max: 50). | 

## Example

```python
from openapi_client.models.v1_entertainment_trivia_batch_post_request import V1EntertainmentTriviaBatchPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1EntertainmentTriviaBatchPostRequest from a JSON string
v1_entertainment_trivia_batch_post_request_instance = V1EntertainmentTriviaBatchPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1EntertainmentTriviaBatchPostRequest.to_json())

# convert the object into a dict
v1_entertainment_trivia_batch_post_request_dict = v1_entertainment_trivia_batch_post_request_instance.to_dict()
# create an instance of V1EntertainmentTriviaBatchPostRequest from a dict
v1_entertainment_trivia_batch_post_request_from_dict = V1EntertainmentTriviaBatchPostRequest.from_dict(v1_entertainment_trivia_batch_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


