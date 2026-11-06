# V1PlacesPostalCodeGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**postal_code** | **str** | The postal code as stored in the dataset | [optional] 
**city** | **str** | Primary city or place name for the postal code | [optional] 
**state** | **str** | State, province, or administrative region name | [optional] 
**country** | **str** | ISO 3166-1 alpha-2 country code (uppercase) | [optional] 
**lat** | **float** | Latitude of the postal code centroid | [optional] 
**lon** | **float** | Longitude of the postal code centroid | [optional] 

## Example

```python
from openapi_client.models.v1_places_postal_code_get200_response_data import V1PlacesPostalCodeGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1PlacesPostalCodeGet200ResponseData from a JSON string
v1_places_postal_code_get200_response_data_instance = V1PlacesPostalCodeGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1PlacesPostalCodeGet200ResponseData.to_json())

# convert the object into a dict
v1_places_postal_code_get200_response_data_dict = v1_places_postal_code_get200_response_data_instance.to_dict()
# create an instance of V1PlacesPostalCodeGet200ResponseData from a dict
v1_places_postal_code_get200_response_data_from_dict = V1PlacesPostalCodeGet200ResponseData.from_dict(v1_places_postal_code_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


