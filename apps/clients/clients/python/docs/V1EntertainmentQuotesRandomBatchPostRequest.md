# V1EntertainmentQuotesRandomBatchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**count** | **int** | Number of random quotes to return. Between 1 and 50. | 

## Example

```python
from openapi_client.models.v1_entertainment_quotes_random_batch_post_request import V1EntertainmentQuotesRandomBatchPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1EntertainmentQuotesRandomBatchPostRequest from a JSON string
v1_entertainment_quotes_random_batch_post_request_instance = V1EntertainmentQuotesRandomBatchPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1EntertainmentQuotesRandomBatchPostRequest.to_json())

# convert the object into a dict
v1_entertainment_quotes_random_batch_post_request_dict = v1_entertainment_quotes_random_batch_post_request_instance.to_dict()
# create an instance of V1EntertainmentQuotesRandomBatchPostRequest from a dict
v1_entertainment_quotes_random_batch_post_request_from_dict = V1EntertainmentQuotesRandomBatchPostRequest.from_dict(v1_entertainment_quotes_random_batch_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


