# V1PlacesGeocodeBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **List[object]** | Per-item results in input order. | [optional] 
**results_address** | **str** | The original address as provided in the request. | [optional] 
**results_result** | **object** | Geocoding result (omitted on error). Same fields as the single-item endpoint. | [optional] 
**results_error** | **str** | Error message if geocoding failed (omitted on success). | [optional] 
**total** | **int** | Total number of items processed. | [optional] 

## Example

```python
from openapi_client.models.v1_places_geocode_batch_post200_response_data import V1PlacesGeocodeBatchPost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1PlacesGeocodeBatchPost200ResponseData from a JSON string
v1_places_geocode_batch_post200_response_data_instance = V1PlacesGeocodeBatchPost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1PlacesGeocodeBatchPost200ResponseData.to_json())

# convert the object into a dict
v1_places_geocode_batch_post200_response_data_dict = v1_places_geocode_batch_post200_response_data_instance.to_dict()
# create an instance of V1PlacesGeocodeBatchPost200ResponseData from a dict
v1_places_geocode_batch_post200_response_data_from_dict = V1PlacesGeocodeBatchPost200ResponseData.from_dict(v1_places_geocode_batch_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


