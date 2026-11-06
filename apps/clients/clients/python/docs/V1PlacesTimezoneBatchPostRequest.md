# V1PlacesTimezoneBatchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**items** | **str** | List of timezone lookup requests (1–50 items). | 
**items_timezone** | **str** | IANA timezone name (e.g. &#39;America/New_York&#39;). Highest priority. | [optional] 
**items_city** | **str** | City name for city-based lookup. Used when timezone is not set. | [optional] 
**items_lat** | **float** | Latitude (-90 to 90). Used when neither timezone nor city is set. | [optional] 
**items_lon** | **float** | Longitude (-180 to 180). Used alongside lat for coordinate-based lookup. | [optional] 

## Example

```python
from openapi_client.models.v1_places_timezone_batch_post_request import V1PlacesTimezoneBatchPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1PlacesTimezoneBatchPostRequest from a JSON string
v1_places_timezone_batch_post_request_instance = V1PlacesTimezoneBatchPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1PlacesTimezoneBatchPostRequest.to_json())

# convert the object into a dict
v1_places_timezone_batch_post_request_dict = v1_places_timezone_batch_post_request_instance.to_dict()
# create an instance of V1PlacesTimezoneBatchPostRequest from a dict
v1_places_timezone_batch_post_request_from_dict = V1PlacesTimezoneBatchPostRequest.from_dict(v1_places_timezone_batch_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


