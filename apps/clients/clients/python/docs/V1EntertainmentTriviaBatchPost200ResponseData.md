# V1EntertainmentTriviaBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **List[object]** | List of trivia questions returned in the same order as the input filters. | [optional] 
**results_category** | **str** | The topic filter used to retrieve the trivia question. | [optional] 
**results_difficulty** | **str** | The complexity level filter used to retrieve the trivia question. | [optional] 
**results_error** | **str** | Describes why the question could not be retrieved for this filter. Omitted if the query was successful. | [optional] 
**results_data** | **object** | The trivia question returned for this filter. Empty if the query failed. | [optional] 
**results_data_question** | **str** | The trivia question text. | [optional] 
**results_data_options** | **List[object]** | List of possible answers for the question. | [optional] 
**results_data_answer** | **str** | The correct answer to the trivia question. | [optional] 
**results_data_category** | **str** | The category the trivia question belongs to. | [optional] 
**results_data_difficulty** | **str** | The difficulty level of the trivia question. | [optional] 
**total** | **int** | Total number of trivia results returned. | [optional] 

## Example

```python
from openapi_client.models.v1_entertainment_trivia_batch_post200_response_data import V1EntertainmentTriviaBatchPost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1EntertainmentTriviaBatchPost200ResponseData from a JSON string
v1_entertainment_trivia_batch_post200_response_data_instance = V1EntertainmentTriviaBatchPost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1EntertainmentTriviaBatchPost200ResponseData.to_json())

# convert the object into a dict
v1_entertainment_trivia_batch_post200_response_data_dict = v1_entertainment_trivia_batch_post200_response_data_instance.to_dict()
# create an instance of V1EntertainmentTriviaBatchPost200ResponseData from a dict
v1_entertainment_trivia_batch_post200_response_data_from_dict = V1EntertainmentTriviaBatchPost200ResponseData.from_dict(v1_entertainment_trivia_batch_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


