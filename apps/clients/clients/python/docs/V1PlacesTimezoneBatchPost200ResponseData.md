# V1PlacesTimezoneBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **List[object]** | Per-item results in input order. | [optional] 
**results_info_timezone** | **str** | IANA timezone identifier. | [optional] 
**results_info_offset** | **str** | UTC offset in +HH:MM or -HH:MM format. | [optional] 
**results_info_current_time** | **str** | Current time in UTC (RFC 3339). | [optional] 
**results_info_is_dst** | **bool** | Whether daylight saving time is currently active. | [optional] 
**results_error** | **str** | Error message if this item failed (omitted on success). | [optional] 
**total** | **int** | Total number of items processed. | [optional] 

## Example

```python
from openapi_client.models.v1_places_timezone_batch_post200_response_data import V1PlacesTimezoneBatchPost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1PlacesTimezoneBatchPost200ResponseData from a JSON string
v1_places_timezone_batch_post200_response_data_instance = V1PlacesTimezoneBatchPost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1PlacesTimezoneBatchPost200ResponseData.to_json())

# convert the object into a dict
v1_places_timezone_batch_post200_response_data_dict = v1_places_timezone_batch_post200_response_data_instance.to_dict()
# create an instance of V1PlacesTimezoneBatchPost200ResponseData from a dict
v1_places_timezone_batch_post200_response_data_from_dict = V1PlacesTimezoneBatchPost200ResponseData.from_dict(v1_places_timezone_batch_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


