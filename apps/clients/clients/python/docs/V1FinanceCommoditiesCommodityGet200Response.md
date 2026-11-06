# V1FinanceCommoditiesCommodityGet200Response


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**data** | [**V1FinanceCommoditiesCommodityGet200ResponseData**](V1FinanceCommoditiesCommodityGet200ResponseData.md) |  | [optional] 
**metadata** | [**V1TextAdviceGet200ResponseMetadata**](V1TextAdviceGet200ResponseMetadata.md) |  | [optional] 

## Example

```python
from openapi_client.models.v1_finance_commodities_commodity_get200_response import V1FinanceCommoditiesCommodityGet200Response

# TODO update the JSON string below
json = "{}"
# create an instance of V1FinanceCommoditiesCommodityGet200Response from a JSON string
v1_finance_commodities_commodity_get200_response_instance = V1FinanceCommoditiesCommodityGet200Response.from_json(json)
# print the JSON string representation of the object
print(V1FinanceCommoditiesCommodityGet200Response.to_json())

# convert the object into a dict
v1_finance_commodities_commodity_get200_response_dict = v1_finance_commodities_commodity_get200_response_instance.to_dict()
# create an instance of V1FinanceCommoditiesCommodityGet200Response from a dict
v1_finance_commodities_commodity_get200_response_from_dict = V1FinanceCommoditiesCommodityGet200Response.from_dict(v1_finance_commodities_commodity_get200_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


