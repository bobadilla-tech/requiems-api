# V1ValidationProfanityBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **List[object]** | Per-item results in input order. | [optional] 
**results_text** | **str** | The original text that was checked. | [optional] 
**results_result_has_profanity** | **bool** | Whether any profanity was detected. | [optional] 
**results_result_censored** | **str** | Text with profane words replaced by asterisks. | [optional] 
**results_result_flagged_words** | **str** | Deduplicated list of detected profane words (lowercase). | [optional] 
**total** | **int** | Total number of items processed. | [optional] 

## Example

```python
from openapi_client.models.v1_validation_profanity_batch_post200_response_data import V1ValidationProfanityBatchPost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1ValidationProfanityBatchPost200ResponseData from a JSON string
v1_validation_profanity_batch_post200_response_data_instance = V1ValidationProfanityBatchPost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1ValidationProfanityBatchPost200ResponseData.to_json())

# convert the object into a dict
v1_validation_profanity_batch_post200_response_data_dict = v1_validation_profanity_batch_post200_response_data_instance.to_dict()
# create an instance of V1ValidationProfanityBatchPost200ResponseData from a dict
v1_validation_profanity_batch_post200_response_data_from_dict = V1ValidationProfanityBatchPost200ResponseData.from_dict(v1_validation_profanity_batch_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


