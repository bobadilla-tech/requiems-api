# openapi_client.InflationApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_finance_inflation_batch_post**](InflationApi.md#v1_finance_inflation_batch_post) | **POST** /v1/finance/inflation/batch | Batch Inflation Rates
[**v1_finance_inflation_get**](InflationApi.md#v1_finance_inflation_get) | **GET** /v1/finance/inflation | Get Inflation Rate


# **v1_finance_inflation_batch_post**
> V1FinanceInflationBatchPost200Response v1_finance_inflation_batch_post(v1_finance_inflation_batch_post_request)

Batch Inflation Rates

Returns inflation data for up to 50 countries in a single request. Results are in the same order as the input. Countries with no data return found: false instead of failing the whole request. Billing: 1 credit per country (not per HTTP request).

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_finance_inflation_batch_post200_response import V1FinanceInflationBatchPost200Response
from openapi_client.models.v1_finance_inflation_batch_post_request import V1FinanceInflationBatchPostRequest
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
    api_instance = openapi_client.InflationApi(api_client)
    v1_finance_inflation_batch_post_request = openapi_client.V1FinanceInflationBatchPostRequest() # V1FinanceInflationBatchPostRequest | 

    try:
        # Batch Inflation Rates
        api_response = api_instance.v1_finance_inflation_batch_post(v1_finance_inflation_batch_post_request)
        print("The response of InflationApi->v1_finance_inflation_batch_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InflationApi->v1_finance_inflation_batch_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_finance_inflation_batch_post_request** | [**V1FinanceInflationBatchPostRequest**](V1FinanceInflationBatchPostRequest.md)|  | 

### Return type

[**V1FinanceInflationBatchPost200Response**](V1FinanceInflationBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**422** | Body is invalid: empty array, more than 50 items, or a bad country code. |  -  |
**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_finance_inflation_get**
> V1FinanceInflationGet200Response v1_finance_inflation_get(country)

Get Inflation Rate

Returns the latest annual CPI inflation rate for a country plus the previous 10 years of historical data.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_finance_inflation_get200_response import V1FinanceInflationGet200Response
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
    api_instance = openapi_client.InflationApi(api_client)
    country = 'US' # str | ISO 3166-1 alpha-2 country code (e.g. US, GB, DE). Case-insensitive.

    try:
        # Get Inflation Rate
        api_response = api_instance.v1_finance_inflation_get(country)
        print("The response of InflationApi->v1_finance_inflation_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InflationApi->v1_finance_inflation_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **str**| ISO 3166-1 alpha-2 country code (e.g. US, GB, DE). Case-insensitive. | 

### Return type

[**V1FinanceInflationGet200Response**](V1FinanceInflationGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | The country parameter is missing or is not a valid ISO 3166-1 alpha-2 code. |  -  |
**404** | No inflation data found for the given country code. |  -  |
**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

