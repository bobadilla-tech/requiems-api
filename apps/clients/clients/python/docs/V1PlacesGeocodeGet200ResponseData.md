# V1PlacesGeocodeGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**address** | **str** | Full display name of the matched location | [optional] 
**city** | **str** | City or town of the matched location | [optional] 
**country** | **str** | ISO 3166-1 alpha-2 country code (uppercase) | [optional] 
**lat** | **float** | Latitude of the matched location | [optional] 
**lon** | **float** | Longitude of the matched location | [optional] 

## Example

```python
from openapi_client.models.v1_places_geocode_get200_response_data import V1PlacesGeocodeGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1PlacesGeocodeGet200ResponseData from a JSON string
v1_places_geocode_get200_response_data_instance = V1PlacesGeocodeGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1PlacesGeocodeGet200ResponseData.to_json())

# convert the object into a dict
v1_places_geocode_get200_response_data_dict = v1_places_geocode_get200_response_data_instance.to_dict()
# create an instance of V1PlacesGeocodeGet200ResponseData from a dict
v1_places_geocode_get200_response_data_from_dict = V1PlacesGeocodeGet200ResponseData.from_dict(v1_places_geocode_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


