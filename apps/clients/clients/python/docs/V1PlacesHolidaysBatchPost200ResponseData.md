# V1PlacesHolidaysBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **List[object]** | One result per query, in the same order as the request | [optional] 
**results_country** | **str** | ISO 3166-1 alpha-2 country code | [optional] 
**results_year** | **int** | Year queried | [optional] 
**results_found** | **bool** | false when no holidays exist for that country/year combination | [optional] 
**results_holidays** | **List[object]** | List of holidays. Omitted when found is false. | [optional] 
**results_total** | **int** | Number of holidays. Omitted when found is false. | [optional] 
**total** | **int** | Total number of results (equals the number of queries sent) | [optional] 

## Example

```python
from openapi_client.models.v1_places_holidays_batch_post200_response_data import V1PlacesHolidaysBatchPost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1PlacesHolidaysBatchPost200ResponseData from a JSON string
v1_places_holidays_batch_post200_response_data_instance = V1PlacesHolidaysBatchPost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1PlacesHolidaysBatchPost200ResponseData.to_json())

# convert the object into a dict
v1_places_holidays_batch_post200_response_data_dict = v1_places_holidays_batch_post200_response_data_instance.to_dict()
# create an instance of V1PlacesHolidaysBatchPost200ResponseData from a dict
v1_places_holidays_batch_post200_response_data_from_dict = V1PlacesHolidaysBatchPost200ResponseData.from_dict(v1_places_holidays_batch_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


