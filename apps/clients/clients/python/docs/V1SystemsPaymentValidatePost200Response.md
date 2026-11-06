# V1SystemsPaymentValidatePost200Response


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**data** | [**V1SystemsPaymentValidatePost200ResponseData**](V1SystemsPaymentValidatePost200ResponseData.md) |  | [optional] 
**metadata** | [**V1TextAdviceGet200ResponseMetadata**](V1TextAdviceGet200ResponseMetadata.md) |  | [optional] 

## Example

```python
from openapi_client.models.v1_systems_payment_validate_post200_response import V1SystemsPaymentValidatePost200Response

# TODO update the JSON string below
json = "{}"
# create an instance of V1SystemsPaymentValidatePost200Response from a JSON string
v1_systems_payment_validate_post200_response_instance = V1SystemsPaymentValidatePost200Response.from_json(json)
# print the JSON string representation of the object
print(V1SystemsPaymentValidatePost200Response.to_json())

# convert the object into a dict
v1_systems_payment_validate_post200_response_dict = v1_systems_payment_validate_post200_response_instance.to_dict()
# create an instance of V1SystemsPaymentValidatePost200Response from a dict
v1_systems_payment_validate_post200_response_from_dict = V1SystemsPaymentValidatePost200Response.from_dict(v1_systems_payment_validate_post200_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


