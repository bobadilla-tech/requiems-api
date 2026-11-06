# openapi_client.ExchangeRateApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_finance_convert_get**](ExchangeRateApi.md#v1_finance_convert_get) | **GET** /v1/finance/convert | Convert Currency
[**v1_finance_exchange_rate_get**](ExchangeRateApi.md#v1_finance_exchange_rate_get) | **GET** /v1/finance/exchange-rate | Get Exchange Rate


# **v1_finance_convert_get**
> V1FinanceConvertGet200Response v1_finance_convert_get(var_from, to, amount)

Convert Currency

Converts an amount from one currency to another and returns the rate alongside the converted value.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_finance_convert_get200_response import V1FinanceConvertGet200Response
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
    api_instance = openapi_client.ExchangeRateApi(api_client)
    var_from = 'USD' # str | ISO 4217 source currency code (3 letters, e.g. USD)
    to = 'EUR' # str | ISO 4217 target currency code (3 letters, e.g. EUR)
    amount = 100 # float | Amount to convert. Must be greater than 0.

    try:
        # Convert Currency
        api_response = api_instance.v1_finance_convert_get(var_from, to, amount)
        print("The response of ExchangeRateApi->v1_finance_convert_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ExchangeRateApi->v1_finance_convert_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **var_from** | **str**| ISO 4217 source currency code (3 letters, e.g. USD) | 
 **to** | **str**| ISO 4217 target currency code (3 letters, e.g. EUR) | 
 **amount** | **float**| Amount to convert. Must be greater than 0. | 

### Return type

[**V1FinanceConvertGet200Response**](V1FinanceConvertGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | A required parameter is missing, the currency code is not 3 alphabetic characters, or the amount is 0 or negative. |  -  |
**422** | One or both currency codes are not recognised by the upstream data source. |  -  |
**503** | The exchange rate data source is temporarily unavailable. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_finance_exchange_rate_get**
> V1FinanceExchangeRateGet200Response v1_finance_exchange_rate_get(var_from, to)

Get Exchange Rate

Returns the current exchange rate between two currencies.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_finance_exchange_rate_get200_response import V1FinanceExchangeRateGet200Response
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
    api_instance = openapi_client.ExchangeRateApi(api_client)
    var_from = 'USD' # str | ISO 4217 source currency code (3 letters, e.g. USD)
    to = 'EUR' # str | ISO 4217 target currency code (3 letters, e.g. EUR)

    try:
        # Get Exchange Rate
        api_response = api_instance.v1_finance_exchange_rate_get(var_from, to)
        print("The response of ExchangeRateApi->v1_finance_exchange_rate_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ExchangeRateApi->v1_finance_exchange_rate_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **var_from** | **str**| ISO 4217 source currency code (3 letters, e.g. USD) | 
 **to** | **str**| ISO 4217 target currency code (3 letters, e.g. EUR) | 

### Return type

[**V1FinanceExchangeRateGet200Response**](V1FinanceExchangeRateGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | A required parameter is missing or the currency code is not exactly 3 alphabetic characters. |  -  |
**422** | One or both currency codes are not recognised by the upstream data source. |  -  |
**503** | The exchange rate data source is temporarily unavailable. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

