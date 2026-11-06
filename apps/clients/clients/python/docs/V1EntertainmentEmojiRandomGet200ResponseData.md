# V1EntertainmentEmojiRandomGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**emoji** | **str** | The rendered emoji glyph | [optional] 
**name** | **str** | CLDR short name in snake_case (e.g. grinning_face) | [optional] 
**category** | **str** | Unicode category (e.g. Smileys &amp; Emotion, Animals &amp; Nature) | [optional] 
**unicode** | **str** | Unicode code-point in U+XXXX notation (e.g. U+1F600) | [optional] 

## Example

```python
from openapi_client.models.v1_entertainment_emoji_random_get200_response_data import V1EntertainmentEmojiRandomGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1EntertainmentEmojiRandomGet200ResponseData from a JSON string
v1_entertainment_emoji_random_get200_response_data_instance = V1EntertainmentEmojiRandomGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1EntertainmentEmojiRandomGet200ResponseData.to_json())

# convert the object into a dict
v1_entertainment_emoji_random_get200_response_data_dict = v1_entertainment_emoji_random_get200_response_data_instance.to_dict()
# create an instance of V1EntertainmentEmojiRandomGet200ResponseData from a dict
v1_entertainment_emoji_random_get200_response_data_from_dict = V1EntertainmentEmojiRandomGet200ResponseData.from_dict(v1_entertainment_emoji_random_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


