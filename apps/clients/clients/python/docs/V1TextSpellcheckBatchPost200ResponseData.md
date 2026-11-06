# V1TextSpellcheckBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **str** | One entry per input text, in the same order as the input array. Each item contains corrected (the fixed text) and corrections (list of individual corrections with original, suggested, suggestions, and position). | [optional] 
**total** | **int** | Number of texts processed. Equals the length of the input array. | [optional] 

## Example

```python
from openapi_client.models.v1_text_spellcheck_batch_post200_response_data import V1TextSpellcheckBatchPost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1TextSpellcheckBatchPost200ResponseData from a JSON string
v1_text_spellcheck_batch_post200_response_data_instance = V1TextSpellcheckBatchPost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1TextSpellcheckBatchPost200ResponseData.to_json())

# convert the object into a dict
v1_text_spellcheck_batch_post200_response_data_dict = v1_text_spellcheck_batch_post200_response_data_instance.to_dict()
# create an instance of V1TextSpellcheckBatchPost200ResponseData from a dict
v1_text_spellcheck_batch_post200_response_data_from_dict = V1TextSpellcheckBatchPost200ResponseData.from_dict(v1_text_spellcheck_batch_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


