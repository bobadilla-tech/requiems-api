# V1PlacesTimeTimezoneGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**timezone** | **str** | IANA timezone identifier (e.g. \&quot;America/New_York\&quot;) | [optional] 
**offset** | **str** | UTC offset in +HH:MM or -HH:MM format (e.g. &#39;-05:00&#39;, &#39;+05:30&#39;) | [optional] 
**current_time** | **str** | Current time in UTC, formatted as RFC 3339 (e.g. \&quot;2024-12-15T14:30:00Z\&quot;) | [optional] 
**is_dst** | **bool** | Whether the timezone is currently observing daylight saving time | [optional] 

## Example

```python
from openapi_client.models.v1_places_time_timezone_get200_response_data import V1PlacesTimeTimezoneGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1PlacesTimeTimezoneGet200ResponseData from a JSON string
v1_places_time_timezone_get200_response_data_instance = V1PlacesTimeTimezoneGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1PlacesTimeTimezoneGet200ResponseData.to_json())

# convert the object into a dict
v1_places_time_timezone_get200_response_data_dict = v1_places_time_timezone_get200_response_data_instance.to_dict()
# create an instance of V1PlacesTimeTimezoneGet200ResponseData from a dict
v1_places_time_timezone_get200_response_data_from_dict = V1PlacesTimeTimezoneGet200ResponseData.from_dict(v1_places_time_timezone_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


