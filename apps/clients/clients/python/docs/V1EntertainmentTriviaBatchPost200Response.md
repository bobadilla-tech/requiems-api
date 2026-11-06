# V1EntertainmentTriviaBatchPost200Response


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**data** | [**V1EntertainmentTriviaBatchPost200ResponseData**](V1EntertainmentTriviaBatchPost200ResponseData.md) |  | [optional] 
**metadata** | [**V1TextAdviceGet200ResponseMetadata**](V1TextAdviceGet200ResponseMetadata.md) |  | [optional] 

## Example

```python
from openapi_client.models.v1_entertainment_trivia_batch_post200_response import V1EntertainmentTriviaBatchPost200Response

# TODO update the JSON string below
json = "{}"
# create an instance of V1EntertainmentTriviaBatchPost200Response from a JSON string
v1_entertainment_trivia_batch_post200_response_instance = V1EntertainmentTriviaBatchPost200Response.from_json(json)
# print the JSON string representation of the object
print(V1EntertainmentTriviaBatchPost200Response.to_json())

# convert the object into a dict
v1_entertainment_trivia_batch_post200_response_dict = v1_entertainment_trivia_batch_post200_response_instance.to_dict()
# create an instance of V1EntertainmentTriviaBatchPost200Response from a dict
v1_entertainment_trivia_batch_post200_response_from_dict = V1EntertainmentTriviaBatchPost200Response.from_dict(v1_entertainment_trivia_batch_post200_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


