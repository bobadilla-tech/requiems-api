# V1PlacesReverseGeocodeBatchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**items** | **str** | List of coordinate pairs to reverse-geocode (1–20 items). Max 20 due to Nominatim usage policy. | 
**items_lat** | **float** | Latitude (-90 to 90). | 
**items_lon** | **float** | Longitude (-180 to 180). | 

## Example

```python
from openapi_client.models.v1_places_reverse_geocode_batch_post_request import V1PlacesReverseGeocodeBatchPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1PlacesReverseGeocodeBatchPostRequest from a JSON string
v1_places_reverse_geocode_batch_post_request_instance = V1PlacesReverseGeocodeBatchPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1PlacesReverseGeocodeBatchPostRequest.to_json())

# convert the object into a dict
v1_places_reverse_geocode_batch_post_request_dict = v1_places_reverse_geocode_batch_post_request_instance.to_dict()
# create an instance of V1PlacesReverseGeocodeBatchPostRequest from a dict
v1_places_reverse_geocode_batch_post_request_from_dict = V1PlacesReverseGeocodeBatchPostRequest.from_dict(v1_places_reverse_geocode_batch_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


