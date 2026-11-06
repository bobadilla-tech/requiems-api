# V1FinanceSwiftCodeGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**swift_code** | **str** | Full 11-character SWIFT/BIC code | [optional] 
**bank_code** | **str** | Institution code (characters 1-4) | [optional] 
**country_code** | **str** | ISO 3166-1 alpha-2 country code (characters 5-6) | [optional] 
**location_code** | **str** | Location code (characters 7-8) | [optional] 
**branch_code** | **str** | Branch code (characters 9-11), XXX for primary office | [optional] 
**bank_name** | **str** | Bank or institution name | [optional] 
**city** | **str** | City of the branch or primary office | [optional] 
**country_name** | **str** | Full country name | [optional] 
**is_primary** | **bool** | true when branch_code is XXX | [optional] 

## Example

```python
from openapi_client.models.v1_finance_swift_code_get200_response_data import V1FinanceSwiftCodeGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1FinanceSwiftCodeGet200ResponseData from a JSON string
v1_finance_swift_code_get200_response_data_instance = V1FinanceSwiftCodeGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1FinanceSwiftCodeGet200ResponseData.to_json())

# convert the object into a dict
v1_finance_swift_code_get200_response_data_dict = v1_finance_swift_code_get200_response_data_instance.to_dict()
# create an instance of V1FinanceSwiftCodeGet200ResponseData from a dict
v1_finance_swift_code_get200_response_data_from_dict = V1FinanceSwiftCodeGet200ResponseData.from_dict(v1_finance_swift_code_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


