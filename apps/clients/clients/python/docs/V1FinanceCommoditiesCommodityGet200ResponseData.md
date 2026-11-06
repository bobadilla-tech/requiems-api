# V1FinanceCommoditiesCommodityGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**commodity** | **str** | The commodity slug as provided in the request path | [optional] 
**name** | **str** | Human-readable commodity name | [optional] 
**price** | **float** | Latest annual average price in the commodity&#39;s display unit | [optional] 
**unit** | **str** | Price unit (oz, barrel, mmbtu, lb, or metric_ton) | [optional] 
**currency** | **str** | Currency code — always USD | [optional] 
**change_24h** | **float** | Year-over-year percentage change from the prior year&#39;s annual average (positive &#x3D; price increased) | [optional] 
**historical** | **List[object]** | Up to 10 prior years of annual average prices, ordered newest to oldest | [optional] 
**historical_period** | **str** | Year of the historical data point | [optional] 
**historical_price** | **float** | Annual average price for that year | [optional] 

## Example

```python
from openapi_client.models.v1_finance_commodities_commodity_get200_response_data import V1FinanceCommoditiesCommodityGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1FinanceCommoditiesCommodityGet200ResponseData from a JSON string
v1_finance_commodities_commodity_get200_response_data_instance = V1FinanceCommoditiesCommodityGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1FinanceCommoditiesCommodityGet200ResponseData.to_json())

# convert the object into a dict
v1_finance_commodities_commodity_get200_response_data_dict = v1_finance_commodities_commodity_get200_response_data_instance.to_dict()
# create an instance of V1FinanceCommoditiesCommodityGet200ResponseData from a dict
v1_finance_commodities_commodity_get200_response_data_from_dict = V1FinanceCommoditiesCommodityGet200ResponseData.from_dict(v1_finance_commodities_commodity_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


