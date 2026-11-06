# V1TextWordsRandomGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** | Unique identifier for the word | [optional] 
**word** | **str** | The random word | [optional] 
**definition** | **str** | Dictionary definition of the word | [optional] 
**part_of_speech** | **str** | Grammatical classification (e.g., noun, verb, adjective, adverb) | [optional] 

## Example

```python
from openapi_client.models.v1_text_words_random_get200_response_data import V1TextWordsRandomGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1TextWordsRandomGet200ResponseData from a JSON string
v1_text_words_random_get200_response_data_instance = V1TextWordsRandomGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1TextWordsRandomGet200ResponseData.to_json())

# convert the object into a dict
v1_text_words_random_get200_response_data_dict = v1_text_words_random_get200_response_data_instance.to_dict()
# create an instance of V1TextWordsRandomGet200ResponseData from a dict
v1_text_words_random_get200_response_data_from_dict = V1TextWordsRandomGet200ResponseData.from_dict(v1_text_words_random_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


