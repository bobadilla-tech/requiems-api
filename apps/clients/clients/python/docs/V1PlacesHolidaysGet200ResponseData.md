# V1PlacesHolidaysGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**country** | **str** | ISO 3166-1 alpha-2 country code | [optional] 
**year** | **int** | Year for which holidays are returned | [optional] 
**holidays** | **List[object]** | Array of holiday objects | [optional] 
**holidays_date** | **str** | Holiday date in YYYY-MM-DD format | [optional] 
**holidays_name** | **str** | Name of the holiday | [optional] 
**total** | **int** | Total number of holidays for the country/year | [optional] 

## Example

```python
from openapi_client.models.v1_places_holidays_get200_response_data import V1PlacesHolidaysGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1PlacesHolidaysGet200ResponseData from a JSON string
v1_places_holidays_get200_response_data_instance = V1PlacesHolidaysGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1PlacesHolidaysGet200ResponseData.to_json())

# convert the object into a dict
v1_places_holidays_get200_response_data_dict = v1_places_holidays_get200_response_data_instance.to_dict()
# create an instance of V1PlacesHolidaysGet200ResponseData from a dict
v1_places_holidays_get200_response_data_from_dict = V1PlacesHolidaysGet200ResponseData.from_dict(v1_places_holidays_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


