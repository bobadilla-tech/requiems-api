# V1TechnologyRandomUserGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Full name of the generated user | [optional] 
**email** | **str** | Email address of the generated user | [optional] 
**phone** | **str** | Phone number of the generated user | [optional] 
**address_street** | **str** | Street address | [optional] 
**address_city** | **str** | City name | [optional] 
**address_state** | **str** | State or region | [optional] 
**address_zip** | **str** | Postal / ZIP code | [optional] 
**address_country** | **str** | Country name | [optional] 
**avatar** | **str** | URL to a unique identicon avatar for the generated user (DiceBear) | [optional] 

## Example

```python
from openapi_client.models.v1_technology_random_user_get200_response_data import V1TechnologyRandomUserGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1TechnologyRandomUserGet200ResponseData from a JSON string
v1_technology_random_user_get200_response_data_instance = V1TechnologyRandomUserGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1TechnologyRandomUserGet200ResponseData.to_json())

# convert the object into a dict
v1_technology_random_user_get200_response_data_dict = v1_technology_random_user_get200_response_data_instance.to_dict()
# create an instance of V1TechnologyRandomUserGet200ResponseData from a dict
v1_technology_random_user_get200_response_data_from_dict = V1TechnologyRandomUserGet200ResponseData.from_dict(v1_technology_random_user_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


