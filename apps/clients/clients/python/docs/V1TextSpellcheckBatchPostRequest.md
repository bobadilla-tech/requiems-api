# V1TextSpellcheckBatchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**texts** | **List[object]** | The list of texts to spell-check. Between 1 and 50 items. | 

## Example

```python
from openapi_client.models.v1_text_spellcheck_batch_post_request import V1TextSpellcheckBatchPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1TextSpellcheckBatchPostRequest from a JSON string
v1_text_spellcheck_batch_post_request_instance = V1TextSpellcheckBatchPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1TextSpellcheckBatchPostRequest.to_json())

# convert the object into a dict
v1_text_spellcheck_batch_post_request_dict = v1_text_spellcheck_batch_post_request_instance.to_dict()
# create an instance of V1TextSpellcheckBatchPostRequest from a dict
v1_text_spellcheck_batch_post_request_from_dict = V1TextSpellcheckBatchPostRequest.from_dict(v1_text_spellcheck_batch_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


