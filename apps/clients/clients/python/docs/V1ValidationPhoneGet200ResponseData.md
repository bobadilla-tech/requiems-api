# V1ValidationPhoneGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**number** | **str** | The original number as supplied in the request | [optional] 
**valid** | **bool** | Whether the number is a valid, dialable phone number | [optional] 
**country** | **str** | ISO 3166-1 alpha-2 country code (omitted when valid is false) | [optional] 
**type** | **str** | Number type: mobile, landline, landline_or_mobile, toll_free, voip, premium_rate, shared_cost, personal_number, pager, uan, voicemail, or unknown (omitted when valid is false) | [optional] 
**formatted** | **str** | International format of the number, e.g. +44 7400 123456 (omitted when valid is false) | [optional] 
**carrier_name** | **str** | Carrier name from phone prefix metadata (omitted when carrier cannot be determined) | [optional] 
**carrier_source** | **str** | How the carrier was determined. Always \&quot;metadata\&quot; when present | [optional] 
**risk_is_voip** | **bool** | true when the number type is voip | [optional] 
**risk_is_virtual** | **bool** | true when the number is not tied to a physical SIM or fixed line: voip, personal_number, uan, pager, or voicemail | [optional] 

## Example

```python
from openapi_client.models.v1_validation_phone_get200_response_data import V1ValidationPhoneGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1ValidationPhoneGet200ResponseData from a JSON string
v1_validation_phone_get200_response_data_instance = V1ValidationPhoneGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1ValidationPhoneGet200ResponseData.to_json())

# convert the object into a dict
v1_validation_phone_get200_response_data_dict = v1_validation_phone_get200_response_data_instance.to_dict()
# create an instance of V1ValidationPhoneGet200ResponseData from a dict
v1_validation_phone_get200_response_data_from_dict = V1ValidationPhoneGet200ResponseData.from_dict(v1_validation_phone_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


