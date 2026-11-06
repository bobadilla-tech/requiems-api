# V1PlacesPostalBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **List[object]** | Per-item results in input order. | [optional] 
**results_code** | **str** | The postal code as provided in the request. | [optional] 
**results_country** | **str** | The country code used for the lookup (defaulted to US when omitted). | [optional] 
**results_found** | **bool** | Whether the postal code was found in the dataset. | [optional] 
**results_result** | **object** | Postal code details (omitted when found is false). | [optional] 
**total** | **int** | Total number of items processed. | [optional] 

## Example

```python
from openapi_client.models.v1_places_postal_batch_post200_response_data import V1PlacesPostalBatchPost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1PlacesPostalBatchPost200ResponseData from a JSON string
v1_places_postal_batch_post200_response_data_instance = V1PlacesPostalBatchPost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1PlacesPostalBatchPost200ResponseData.to_json())

# convert the object into a dict
v1_places_postal_batch_post200_response_data_dict = v1_places_postal_batch_post200_response_data_instance.to_dict()
# create an instance of V1PlacesPostalBatchPost200ResponseData from a dict
v1_places_postal_batch_post200_response_data_from_dict = V1PlacesPostalBatchPost200ResponseData.from_dict(v1_places_postal_batch_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


