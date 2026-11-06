# V1TextSpellcheckPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**corrected** | **str** | The full input text with all misspelled words replaced by their suggested corrections | [optional] 
**corrections** | **str** | List of individual corrections. Each item contains: original (the misspelled word), suggested (the correction), and position (0-based character offset in the original text) | [optional] 

## Example

```python
from openapi_client.models.v1_text_spellcheck_post200_response_data import V1TextSpellcheckPost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1TextSpellcheckPost200ResponseData from a JSON string
v1_text_spellcheck_post200_response_data_instance = V1TextSpellcheckPost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1TextSpellcheckPost200ResponseData.to_json())

# convert the object into a dict
v1_text_spellcheck_post200_response_data_dict = v1_text_spellcheck_post200_response_data_instance.to_dict()
# create an instance of V1TextSpellcheckPost200ResponseData from a dict
v1_text_spellcheck_post200_response_data_from_dict = V1TextSpellcheckPost200ResponseData.from_dict(v1_text_spellcheck_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


