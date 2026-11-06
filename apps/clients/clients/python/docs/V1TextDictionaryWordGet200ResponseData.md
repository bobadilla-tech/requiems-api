# V1TextDictionaryWordGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**word** | **str** | The normalized (lowercased) word that was looked up | [optional] 
**phonetic** | **str** | IPA phonetic transcription of the word (may be omitted if unavailable) | [optional] 
**definitions** | **str** | One or more definitions for the word, each with partOfSpeech, definition, and an optional example | [optional] 
**definitions_part_of_speech** | **str** | Grammatical category (e.g. noun, verb, adjective) | [optional] 
**definitions_definition** | **str** | Plain-text definition of the word | [optional] 
**definitions_example** | **str** | Example sentence using the word (may be omitted) | [optional] 
**synonyms** | **str** | List of words with similar meaning | [optional] 

## Example

```python
from openapi_client.models.v1_text_dictionary_word_get200_response_data import V1TextDictionaryWordGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1TextDictionaryWordGet200ResponseData from a JSON string
v1_text_dictionary_word_get200_response_data_instance = V1TextDictionaryWordGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1TextDictionaryWordGet200ResponseData.to_json())

# convert the object into a dict
v1_text_dictionary_word_get200_response_data_dict = v1_text_dictionary_word_get200_response_data_instance.to_dict()
# create an instance of V1TextDictionaryWordGet200ResponseData from a dict
v1_text_dictionary_word_get200_response_data_from_dict = V1TextDictionaryWordGet200ResponseData.from_dict(v1_text_dictionary_word_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


