# V1TechnologyPasswordGet200Response


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**data** | [**V1TechnologyPasswordGet200ResponseData**](V1TechnologyPasswordGet200ResponseData.md) |  | [optional] 
**metadata** | [**V1TextAdviceGet200ResponseMetadata**](V1TextAdviceGet200ResponseMetadata.md) |  | [optional] 

## Example

```python
from openapi_client.models.v1_technology_password_get200_response import V1TechnologyPasswordGet200Response

# TODO update the JSON string below
json = "{}"
# create an instance of V1TechnologyPasswordGet200Response from a JSON string
v1_technology_password_get200_response_instance = V1TechnologyPasswordGet200Response.from_json(json)
# print the JSON string representation of the object
print(V1TechnologyPasswordGet200Response.to_json())

# convert the object into a dict
v1_technology_password_get200_response_dict = v1_technology_password_get200_response_instance.to_dict()
# create an instance of V1TechnologyPasswordGet200Response from a dict
v1_technology_password_get200_response_from_dict = V1TechnologyPasswordGet200Response.from_dict(v1_technology_password_get200_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


