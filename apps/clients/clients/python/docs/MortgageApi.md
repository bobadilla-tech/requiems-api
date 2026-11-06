# openapi_client.MortgageApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_finance_mortgage_get**](MortgageApi.md#v1_finance_mortgage_get) | **GET** /v1/finance/mortgage | Calculate Mortgage


# **v1_finance_mortgage_get**
> V1FinanceMortgageGet200Response v1_finance_mortgage_get(principal, rate, years)

Calculate Mortgage

Returns the monthly payment, total cost, and full amortization schedule for a fixed-rate mortgage.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_finance_mortgage_get200_response import V1FinanceMortgageGet200Response
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
    api_instance = openapi_client.MortgageApi(api_client)
    principal = 300000 # float | Loan amount in your chosen currency (e.g. 300000 for $300,000)
    rate = 6.5 # float | Annual interest rate as a percentage (e.g. 6.5 for 6.5%). Must be greater than 0.
    years = 30 # int | Loan term in years (1–50)

    try:
        # Calculate Mortgage
        api_response = api_instance.v1_finance_mortgage_get(principal, rate, years)
        print("The response of MortgageApi->v1_finance_mortgage_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MortgageApi->v1_finance_mortgage_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **principal** | **float**| Loan amount in your chosen currency (e.g. 300000 for $300,000) | 
 **rate** | **float**| Annual interest rate as a percentage (e.g. 6.5 for 6.5%). Must be greater than 0. | 
 **years** | **int**| Loan term in years (1–50) | 

### Return type

[**V1FinanceMortgageGet200Response**](V1FinanceMortgageGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | A required parameter is missing, not a valid number, or out of range (e.g. years &gt; 50 or rate &lt;&#x3D; 0). |  -  |
**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

