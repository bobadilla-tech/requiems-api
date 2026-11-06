# V1EntertainmentChuckNorrisGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **str** | Unique fact identifier in the format cn_&lt;index&gt; (e.g. cn_0, cn_7) | [optional] 
**fact** | **str** | The Chuck Norris fact text | [optional] 

## Example

```python
from openapi_client.models.v1_entertainment_chuck_norris_get200_response_data import V1EntertainmentChuckNorrisGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1EntertainmentChuckNorrisGet200ResponseData from a JSON string
v1_entertainment_chuck_norris_get200_response_data_instance = V1EntertainmentChuckNorrisGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1EntertainmentChuckNorrisGet200ResponseData.to_json())

# convert the object into a dict
v1_entertainment_chuck_norris_get200_response_data_dict = v1_entertainment_chuck_norris_get200_response_data_instance.to_dict()
# create an instance of V1EntertainmentChuckNorrisGet200ResponseData from a dict
v1_entertainment_chuck_norris_get200_response_data_from_dict = V1EntertainmentChuckNorrisGet200ResponseData.from_dict(v1_entertainment_chuck_norris_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


