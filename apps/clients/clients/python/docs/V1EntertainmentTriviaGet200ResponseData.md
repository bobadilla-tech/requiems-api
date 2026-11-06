# V1EntertainmentTriviaGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**question** | **str** | The trivia question text | [optional] 
**options** | **str** | Four multiple-choice answer options | [optional] 
**answer** | **str** | The correct answer — always one of the values in options | [optional] 
**category** | **str** | The category the question belongs to | [optional] 
**difficulty** | **str** | The difficulty level of the question (easy, medium, or hard) | [optional] 

## Example

```python
from openapi_client.models.v1_entertainment_trivia_get200_response_data import V1EntertainmentTriviaGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1EntertainmentTriviaGet200ResponseData from a JSON string
v1_entertainment_trivia_get200_response_data_instance = V1EntertainmentTriviaGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1EntertainmentTriviaGet200ResponseData.to_json())

# convert the object into a dict
v1_entertainment_trivia_get200_response_data_dict = v1_entertainment_trivia_get200_response_data_instance.to_dict()
# create an instance of V1EntertainmentTriviaGet200ResponseData from a dict
v1_entertainment_trivia_get200_response_data_from_dict = V1EntertainmentTriviaGet200ResponseData.from_dict(v1_entertainment_trivia_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


