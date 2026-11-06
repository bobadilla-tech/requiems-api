# V1PlacesCitiesCityGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Official city name as listed in the GeoNames dataset | [optional] 
**country** | **str** | ISO 3166-1 alpha-2 country code (uppercase) | [optional] 
**population** | **int** | City population from the GeoNames dataset | [optional] 
**timezone** | **str** | IANA timezone identifier for the city (e.g. \&quot;America/New_York\&quot;) | [optional] 
**lat** | **float** | Latitude of the city centre | [optional] 
**lon** | **float** | Longitude of the city centre | [optional] 

## Example

```python
from openapi_client.models.v1_places_cities_city_get200_response_data import V1PlacesCitiesCityGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1PlacesCitiesCityGet200ResponseData from a JSON string
v1_places_cities_city_get200_response_data_instance = V1PlacesCitiesCityGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1PlacesCitiesCityGet200ResponseData.to_json())

# convert the object into a dict
v1_places_cities_city_get200_response_data_dict = v1_places_cities_city_get200_response_data_instance.to_dict()
# create an instance of V1PlacesCitiesCityGet200ResponseData from a dict
v1_places_cities_city_get200_response_data_from_dict = V1PlacesCitiesCityGet200ResponseData.from_dict(v1_places_cities_city_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


