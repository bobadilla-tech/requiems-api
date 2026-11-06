# V1TextLoremBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**data_data** | **object** | Contains the generated text if the specific item was successful | [optional] 
**data_error** | **str** | Error message if the specific item failed to generate | [optional] 

## Example

```python
from openapi_client.models.v1_text_lorem_batch_post200_response_data import V1TextLoremBatchPost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1TextLoremBatchPost200ResponseData from a JSON string
v1_text_lorem_batch_post200_response_data_instance = V1TextLoremBatchPost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1TextLoremBatchPost200ResponseData.to_json())

# convert the object into a dict
v1_text_lorem_batch_post200_response_data_dict = v1_text_lorem_batch_post200_response_data_instance.to_dict()
# create an instance of V1TextLoremBatchPost200ResponseData from a dict
v1_text_lorem_batch_post200_response_data_from_dict = V1TextLoremBatchPost200ResponseData.from_dict(v1_text_lorem_batch_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


