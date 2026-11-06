# V1ValidationEmailPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**email** | **str** | The email address exactly as supplied in the request body; null when syntax is invalid | [optional] 
**valid** | **bool** | Overall validity. True only when the address passes syntax validation and the domain has at least one MX record | [optional] 
**syntax_valid** | **bool** | Whether the address is syntactically valid according to RFC 5322 | [optional] 
**mx_valid** | **bool** | Whether the domain has at least one MX record, meaning it can receive email | [optional] 
**disposable** | **bool** | Whether the address uses a known disposable or temporary email domain | [optional] 
**normalized** | **str** | The canonical form of the address after normalization (lowercase, plus-tag removal, alias-domain resolution). Null when syntax is invalid | [optional] 
**domain** | **str** | The domain part of the address (after @). Null when syntax is invalid | [optional] 
**suggestion** | **str** | A corrected domain name when the supplied domain looks like a typo of a well-known provider (e.g. gmial.com → gmail.com). Null when no close match is found or the domain is already correct | [optional] 

## Example

```python
from openapi_client.models.v1_validation_email_post200_response_data import V1ValidationEmailPost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1ValidationEmailPost200ResponseData from a JSON string
v1_validation_email_post200_response_data_instance = V1ValidationEmailPost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1ValidationEmailPost200ResponseData.to_json())

# convert the object into a dict
v1_validation_email_post200_response_data_dict = v1_validation_email_post200_response_data_instance.to_dict()
# create an instance of V1ValidationEmailPost200ResponseData from a dict
v1_validation_email_post200_response_data_from_dict = V1ValidationEmailPost200ResponseData.from_dict(v1_validation_email_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


