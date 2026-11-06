# V1PlacesReverseGeocodeBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **List[object]** | Per-item results in input order. | [optional] 
**results_lat** | **float** | Latitude as provided in the request. | [optional] 
**results_lon** | **float** | Longitude as provided in the request. | [optional] 
**results_result** | **object** | Reverse geocoding result (omitted on error). Same fields as the single-item endpoint. | [optional] 
**results_error** | **str** | Error message if reverse geocoding failed (omitted on success). | [optional] 
**total** | **int** | Total number of items processed. | [optional] 

## Example

```python
from openapi_client.models.v1_places_reverse_geocode_batch_post200_response_data import V1PlacesReverseGeocodeBatchPost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1PlacesReverseGeocodeBatchPost200ResponseData from a JSON string
v1_places_reverse_geocode_batch_post200_response_data_instance = V1PlacesReverseGeocodeBatchPost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1PlacesReverseGeocodeBatchPost200ResponseData.to_json())

# convert the object into a dict
v1_places_reverse_geocode_batch_post200_response_data_dict = v1_places_reverse_geocode_batch_post200_response_data_instance.to_dict()
# create an instance of V1PlacesReverseGeocodeBatchPost200ResponseData from a dict
v1_places_reverse_geocode_batch_post200_response_data_from_dict = V1PlacesReverseGeocodeBatchPost200ResponseData.from_dict(v1_places_reverse_geocode_batch_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


