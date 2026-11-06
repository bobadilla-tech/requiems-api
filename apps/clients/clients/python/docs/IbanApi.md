# openapi_client.IbanApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_finance_iban_batch_post**](IbanApi.md#v1_finance_iban_batch_post) | **POST** /v1/finance/iban/batch | Batch Validate IBANs
[**v1_finance_iban_iban_get**](IbanApi.md#v1_finance_iban_iban_get) | **GET** /v1/finance/iban/{iban} | Validate IBAN


# **v1_finance_iban_batch_post**
> V1FinanceIbanBatchPost200Response v1_finance_iban_batch_post(v1_finance_iban_batch_post_request)

Batch Validate IBANs

Validates up to 50 iban numbers in a single request. Results are returned in the same order as the input.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_finance_iban_batch_post200_response import V1FinanceIbanBatchPost200Response
from openapi_client.models.v1_finance_iban_batch_post_request import V1FinanceIbanBatchPostRequest
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
    api_instance = openapi_client.IbanApi(api_client)
    v1_finance_iban_batch_post_request = openapi_client.V1FinanceIbanBatchPostRequest() # V1FinanceIbanBatchPostRequest | 

    try:
        # Batch Validate IBANs
        api_response = api_instance.v1_finance_iban_batch_post(v1_finance_iban_batch_post_request)
        print("The response of IbanApi->v1_finance_iban_batch_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling IbanApi->v1_finance_iban_batch_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_finance_iban_batch_post_request** | [**V1FinanceIbanBatchPostRequest**](V1FinanceIbanBatchPostRequest.md)|  | 

### Return type

[**V1FinanceIbanBatchPost200Response**](V1FinanceIbanBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**422** | The numbers array is missing, empty, or contains more than 50 items. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_finance_iban_iban_get**
> V1FinanceIbanIbanGet200Response v1_finance_iban_iban_get(iban)

Validate IBAN

Validates an IBAN and returns the country, bank code, and account number. Spaces in the input are stripped automatically. Always returns HTTP 200 — check the valid field to determine whether the IBAN is valid.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_finance_iban_iban_get200_response import V1FinanceIbanIbanGet200Response
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
    api_instance = openapi_client.IbanApi(api_client)
    iban = 'DE89370400440532013000' # str | The IBAN to validate. Spaces are stripped. Case-insensitive.

    try:
        # Validate IBAN
        api_response = api_instance.v1_finance_iban_iban_get(iban)
        print("The response of IbanApi->v1_finance_iban_iban_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling IbanApi->v1_finance_iban_iban_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iban** | **str**| The IBAN to validate. Spaces are stripped. Case-insensitive. | 

### Return type

[**V1FinanceIbanIbanGet200Response**](V1FinanceIbanIbanGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**500** | Unexpected server error (e.g. database unreachable). |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

