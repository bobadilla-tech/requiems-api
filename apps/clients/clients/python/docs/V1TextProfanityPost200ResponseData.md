# V1TextProfanityPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**has_profanity** | **bool** | Whether any profanity was detected in the text | [optional] 
**censored** | **str** | The input text with profane words replaced by asterisks | [optional] 
**flagged_words** | **str** | Deduplicated list of profane words found (lowercase) | [optional] 

## Example

```python
from openapi_client.models.v1_text_profanity_post200_response_data import V1TextProfanityPost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1TextProfanityPost200ResponseData from a JSON string
v1_text_profanity_post200_response_data_instance = V1TextProfanityPost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1TextProfanityPost200ResponseData.to_json())

# convert the object into a dict
v1_text_profanity_post200_response_data_dict = v1_text_profanity_post200_response_data_instance.to_dict()
# create an instance of V1TextProfanityPost200ResponseData from a dict
v1_text_profanity_post200_response_data_from_dict = V1TextProfanityPost200ResponseData.from_dict(v1_text_profanity_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


