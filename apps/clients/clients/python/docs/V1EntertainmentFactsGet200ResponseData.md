# V1EntertainmentFactsGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**fact** | **str** | The fact text | [optional] 
**category** | **str** | The category the fact belongs to | [optional] 
**source** | **str** | The source or publication the fact is attributed to | [optional] 

## Example

```python
from openapi_client.models.v1_entertainment_facts_get200_response_data import V1EntertainmentFactsGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1EntertainmentFactsGet200ResponseData from a JSON string
v1_entertainment_facts_get200_response_data_instance = V1EntertainmentFactsGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1EntertainmentFactsGet200ResponseData.to_json())

# convert the object into a dict
v1_entertainment_facts_get200_response_data_dict = v1_entertainment_facts_get200_response_data_instance.to_dict()
# create an instance of V1EntertainmentFactsGet200ResponseData from a dict
v1_entertainment_facts_get200_response_data_from_dict = V1EntertainmentFactsGet200ResponseData.from_dict(v1_entertainment_facts_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


