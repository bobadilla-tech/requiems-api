# openapi_client.CommoditiesApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_finance_commodities_commodity_get**](CommoditiesApi.md#v1_finance_commodities_commodity_get) | **GET** /v1/finance/commodities/{commodity} | Get Commodity Price


# **v1_finance_commodities_commodity_get**
> V1FinanceCommoditiesCommodityGet200Response v1_finance_commodities_commodity_get(commodity)

Get Commodity Price

Returns the latest annual average price and up to 10 years of historical data for the requested commodity slug.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_finance_commodities_commodity_get200_response import V1FinanceCommoditiesCommodityGet200Response
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.requiems.xyz
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://api.requiems.xyz"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: requiems-api-key
configuration.api_key['requiems-api-key'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['requiems-api-key'] = 'Bearer'

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.CommoditiesApi(api_client)
    commodity = 'gold' # str | Commodity slug (e.g. gold, silver, oil). See supported slugs below.

    try:
        # Get Commodity Price
        api_response = api_instance.v1_finance_commodities_commodity_get(commodity)
        print("The response of CommoditiesApi->v1_finance_commodities_commodity_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CommoditiesApi->v1_finance_commodities_commodity_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commodity** | **str**| Commodity slug (e.g. gold, silver, oil). See supported slugs below. | 

### Return type

[**V1FinanceCommoditiesCommodityGet200Response**](V1FinanceCommoditiesCommodityGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**404** | No data found for the given commodity slug. Check the supported slugs list. |  -  |
**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

