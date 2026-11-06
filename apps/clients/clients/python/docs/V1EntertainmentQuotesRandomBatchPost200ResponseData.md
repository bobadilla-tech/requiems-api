# V1EntertainmentQuotesRandomBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **List[object]** | Array of random quotes, one per requested count. | [optional] 
**results_id** | **int** | Unique identifier for the quote. | [optional] 
**results_text** | **str** | The quote text. | [optional] 
**results_author** | **str** | Name of the person who said or wrote the quote. May be empty. | [optional] 
**total** | **int** | Total number of quotes returned. Always equals the requested count. | [optional] 

## Example

```python
from openapi_client.models.v1_entertainment_quotes_random_batch_post200_response_data import V1EntertainmentQuotesRandomBatchPost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1EntertainmentQuotesRandomBatchPost200ResponseData from a JSON string
v1_entertainment_quotes_random_batch_post200_response_data_instance = V1EntertainmentQuotesRandomBatchPost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1EntertainmentQuotesRandomBatchPost200ResponseData.to_json())

# convert the object into a dict
v1_entertainment_quotes_random_batch_post200_response_data_dict = v1_entertainment_quotes_random_batch_post200_response_data_instance.to_dict()
# create an instance of V1EntertainmentQuotesRandomBatchPost200ResponseData from a dict
v1_entertainment_quotes_random_batch_post200_response_data_from_dict = V1EntertainmentQuotesRandomBatchPost200ResponseData.from_dict(v1_entertainment_quotes_random_batch_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


