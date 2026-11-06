# V1PlacesWorkingDaysGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**working_days** | **int** | Number of working days between the two dates (excluding weekends and optionally holidays) | [optional] 
**var_from** | **str** | Start date (echoed from request) | [optional] 
**to** | **str** | End date (echoed from request) | [optional] 
**country** | **str** | Country code (echoed from request, empty string if not provided) | [optional] 
**subdivision** | **str** | Subdivision code (echoed from request, empty string if not provided) | [optional] 

## Example

```python
from openapi_client.models.v1_places_working_days_get200_response_data import V1PlacesWorkingDaysGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1PlacesWorkingDaysGet200ResponseData from a JSON string
v1_places_working_days_get200_response_data_instance = V1PlacesWorkingDaysGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1PlacesWorkingDaysGet200ResponseData.to_json())

# convert the object into a dict
v1_places_working_days_get200_response_data_dict = v1_places_working_days_get200_response_data_instance.to_dict()
# create an instance of V1PlacesWorkingDaysGet200ResponseData from a dict
v1_places_working_days_get200_response_data_from_dict = V1PlacesWorkingDaysGet200ResponseData.from_dict(v1_places_working_days_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


