# V1PlacesHolidaysBatchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**queries** | **List[object]** | Array of (country, year) pairs. Min: 1, Max: 50. | 

## Example

```python
from openapi_client.models.v1_places_holidays_batch_post_request import V1PlacesHolidaysBatchPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1PlacesHolidaysBatchPostRequest from a JSON string
v1_places_holidays_batch_post_request_instance = V1PlacesHolidaysBatchPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1PlacesHolidaysBatchPostRequest.to_json())

# convert the object into a dict
v1_places_holidays_batch_post_request_dict = v1_places_holidays_batch_post_request_instance.to_dict()
# create an instance of V1PlacesHolidaysBatchPostRequest from a dict
v1_places_holidays_batch_post_request_from_dict = V1PlacesHolidaysBatchPostRequest.from_dict(v1_places_holidays_batch_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


