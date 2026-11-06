# V1SystemsContentModeratePost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**is_safe** | **bool** | False when any moderation category is triggered. True otherwise. | [optional] 
**toxicity_score** | **str** | Reserved for future use. Currently always null. | [optional] 
**sentiment** | **str** | Detected sentiment: positive, negative, or neutral | [optional] 
**language** | **str** | Detected or provided ISO 639-1 language code. Null when detection confidence is 0. | [optional] 
**language_confidence** | **str** | Confidence in the detected language from 0.0 to 1.0. 1.0 when language was explicitly provided. | [optional] 
**flags** | **List[object]** | Triggered moderation flags. Currently possible: text_profanity | [optional] 
**categories** | **object** | Per-category boolean flags: profanity, hate_speech, spam, violence | [optional] 

## Example

```python
from openapi_client.models.v1_systems_content_moderate_post200_response_data import V1SystemsContentModeratePost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1SystemsContentModeratePost200ResponseData from a JSON string
v1_systems_content_moderate_post200_response_data_instance = V1SystemsContentModeratePost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1SystemsContentModeratePost200ResponseData.to_json())

# convert the object into a dict
v1_systems_content_moderate_post200_response_data_dict = v1_systems_content_moderate_post200_response_data_instance.to_dict()
# create an instance of V1SystemsContentModeratePost200ResponseData from a dict
v1_systems_content_moderate_post200_response_data_from_dict = V1SystemsContentModeratePost200ResponseData.from_dict(v1_systems_content_moderate_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


