# V1EntertainmentEmojiSearchGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**items** | **List[object]** | List of matching emoji objects | [optional] 
**total** | **int** | Total number of matches | [optional] 

## Example

```python
from openapi_client.models.v1_entertainment_emoji_search_get200_response_data import V1EntertainmentEmojiSearchGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1EntertainmentEmojiSearchGet200ResponseData from a JSON string
v1_entertainment_emoji_search_get200_response_data_instance = V1EntertainmentEmojiSearchGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1EntertainmentEmojiSearchGet200ResponseData.to_json())

# convert the object into a dict
v1_entertainment_emoji_search_get200_response_data_dict = v1_entertainment_emoji_search_get200_response_data_instance.to_dict()
# create an instance of V1EntertainmentEmojiSearchGet200ResponseData from a dict
v1_entertainment_emoji_search_get200_response_data_from_dict = V1EntertainmentEmojiSearchGet200ResponseData.from_dict(v1_entertainment_emoji_search_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


