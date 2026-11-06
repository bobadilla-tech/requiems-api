# V1TechnologyBase64DecodePostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**value** | **str** | The Base64-encoded string to decode | 
**variant** | **str** | Encoding variant: standard (default) or url (URL-safe base64url) | [optional] 

## Example

```python
from openapi_client.models.v1_technology_base64_decode_post_request import V1TechnologyBase64DecodePostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1TechnologyBase64DecodePostRequest from a JSON string
v1_technology_base64_decode_post_request_instance = V1TechnologyBase64DecodePostRequest.from_json(json)
# print the JSON string representation of the object
print(V1TechnologyBase64DecodePostRequest.to_json())

# convert the object into a dict
v1_technology_base64_decode_post_request_dict = v1_technology_base64_decode_post_request_instance.to_dict()
# create an instance of V1TechnologyBase64DecodePostRequest from a dict
v1_technology_base64_decode_post_request_from_dict = V1TechnologyBase64DecodePostRequest.from_dict(v1_technology_base64_decode_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


