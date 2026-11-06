# V1TextDetectLanguagePost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**language** | **str** | Full name of the detected language (e.g. French, English, Spanish) | [optional] 
**code** | **str** | ISO 639-1 two-letter language code (e.g. fr, en, es). Empty string when detection is unreliable. | [optional] 
**confidence** | **str** | Confidence score between 0.0 and 1.0. 0.0 is returned when the language cannot be reliably detected. | [optional] 

## Example

```python
from openapi_client.models.v1_text_detect_language_post200_response_data import V1TextDetectLanguagePost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1TextDetectLanguagePost200ResponseData from a JSON string
v1_text_detect_language_post200_response_data_instance = V1TextDetectLanguagePost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1TextDetectLanguagePost200ResponseData.to_json())

# convert the object into a dict
v1_text_detect_language_post200_response_data_dict = v1_text_detect_language_post200_response_data_instance.to_dict()
# create an instance of V1TextDetectLanguagePost200ResponseData from a dict
v1_text_detect_language_post200_response_data_from_dict = V1TextDetectLanguagePost200ResponseData.from_dict(v1_text_detect_language_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


