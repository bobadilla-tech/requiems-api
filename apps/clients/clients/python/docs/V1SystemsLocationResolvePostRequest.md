# V1SystemsLocationResolvePostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**address** | **str** | Human-readable address string to geocode. Either address or coordinates is required. | [optional] 
**coordinates** | **object** | Coordinates object for reverse geocoding. Either address or coordinates is required. Shape: {lat: float, lng: float} | [optional] 
**country_code** | **str** | ISO 3166-1 alpha-2 country code hint. Overrides the geocoder-inferred country code when provided. | [optional] 

## Example

```python
from openapi_client.models.v1_systems_location_resolve_post_request import V1SystemsLocationResolvePostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1SystemsLocationResolvePostRequest from a JSON string
v1_systems_location_resolve_post_request_instance = V1SystemsLocationResolvePostRequest.from_json(json)
# print the JSON string representation of the object
print(V1SystemsLocationResolvePostRequest.to_json())

# convert the object into a dict
v1_systems_location_resolve_post_request_dict = v1_systems_location_resolve_post_request_instance.to_dict()
# create an instance of V1SystemsLocationResolvePostRequest from a dict
v1_systems_location_resolve_post_request_from_dict = V1SystemsLocationResolvePostRequest.from_dict(v1_systems_location_resolve_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


