# V1PlacesPostalBatchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**items** | **str** | List of postal code lookups (1–50 items). | 
**items_code** | **str** | The postal code to look up. | 
**items_country** | **str** | ISO 3166-1 alpha-2 country code (default: US). | [optional] 

## Example

```python
from openapi_client.models.v1_places_postal_batch_post_request import V1PlacesPostalBatchPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1PlacesPostalBatchPostRequest from a JSON string
v1_places_postal_batch_post_request_instance = V1PlacesPostalBatchPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1PlacesPostalBatchPostRequest.to_json())

# convert the object into a dict
v1_places_postal_batch_post_request_dict = v1_places_postal_batch_post_request_instance.to_dict()
# create an instance of V1PlacesPostalBatchPostRequest from a dict
v1_places_postal_batch_post_request_from_dict = V1PlacesPostalBatchPostRequest.from_dict(v1_places_postal_batch_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


