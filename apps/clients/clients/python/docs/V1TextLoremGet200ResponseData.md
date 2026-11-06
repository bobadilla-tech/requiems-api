# V1TextLoremGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**text** | **str** | Generated Lorem Ipsum text | [optional] 
**paragraphs** | **int** | Number of paragraphs generated | [optional] 
**word_count** | **int** | Total number of words in generated text | [optional] 

## Example

```python
from openapi_client.models.v1_text_lorem_get200_response_data import V1TextLoremGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1TextLoremGet200ResponseData from a JSON string
v1_text_lorem_get200_response_data_instance = V1TextLoremGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1TextLoremGet200ResponseData.to_json())

# convert the object into a dict
v1_text_lorem_get200_response_data_dict = v1_text_lorem_get200_response_data_instance.to_dict()
# create an instance of V1TextLoremGet200ResponseData from a dict
v1_text_lorem_get200_response_data_from_dict = V1TextLoremGet200ResponseData.from_dict(v1_text_lorem_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


