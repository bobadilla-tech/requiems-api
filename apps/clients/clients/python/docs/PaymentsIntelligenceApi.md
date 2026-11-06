# openapi_client.PaymentsIntelligenceApi

All URIs are relative to *https://requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_systems_payment_validate_post**](PaymentsIntelligenceApi.md#v1_systems_payment_validate_post) | **POST** /v1/systems/payment/validate | Validate Payment
[**v1_systems_transaction_risk_post**](PaymentsIntelligenceApi.md#v1_systems_transaction_risk_post) | **POST** /v1/systems/transaction/risk | Score Transaction Risk


# **v1_systems_payment_validate_post**
> V1SystemsPaymentValidatePost200Response v1_systems_payment_validate_post(v1_systems_payment_validate_post_request)

Validate Payment

Validate and cross-check BIN, IBAN, and SWIFT identifiers in one call. Returns per-field results and a consistency check that flags country and bank mismatches.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_systems_payment_validate_post200_response import V1SystemsPaymentValidatePost200Response
from openapi_client.models.v1_systems_payment_validate_post_request import V1SystemsPaymentValidatePostRequest
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://requiems.xyz
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://requiems.xyz"
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
    api_instance = openapi_client.PaymentsIntelligenceApi(api_client)
    v1_systems_payment_validate_post_request = openapi_client.V1SystemsPaymentValidatePostRequest() # V1SystemsPaymentValidatePostRequest | 

    try:
        # Validate Payment
        api_response = api_instance.v1_systems_payment_validate_post(v1_systems_payment_validate_post_request)
        print("The response of PaymentsIntelligenceApi->v1_systems_payment_validate_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling PaymentsIntelligenceApi->v1_systems_payment_validate_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_systems_payment_validate_post_request** | [**V1SystemsPaymentValidatePostRequest**](V1SystemsPaymentValidatePostRequest.md)|  | 

### Return type

[**V1SystemsPaymentValidatePost200Response**](V1SystemsPaymentValidatePost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**401** | Missing or invalid API key |  -  |
**422** | All of bin, iban, and swift were omitted. At least one is required. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_systems_transaction_risk_post**
> V1SystemsTransactionRiskPost200Response v1_systems_transaction_risk_post(v1_systems_transaction_risk_post_request)

Score Transaction Risk

Score a transaction for fraud risk by cross-checking the card BIN country against IP geolocation and billing country. Detects VPN, proxy, and TOR use.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_systems_transaction_risk_post200_response import V1SystemsTransactionRiskPost200Response
from openapi_client.models.v1_systems_transaction_risk_post_request import V1SystemsTransactionRiskPostRequest
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://requiems.xyz
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://requiems.xyz"
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
    api_instance = openapi_client.PaymentsIntelligenceApi(api_client)
    v1_systems_transaction_risk_post_request = openapi_client.V1SystemsTransactionRiskPostRequest() # V1SystemsTransactionRiskPostRequest | 

    try:
        # Score Transaction Risk
        api_response = api_instance.v1_systems_transaction_risk_post(v1_systems_transaction_risk_post_request)
        print("The response of PaymentsIntelligenceApi->v1_systems_transaction_risk_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling PaymentsIntelligenceApi->v1_systems_transaction_risk_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_systems_transaction_risk_post_request** | [**V1SystemsTransactionRiskPostRequest**](V1SystemsTransactionRiskPostRequest.md)|  | 

### Return type

[**V1SystemsTransactionRiskPost200Response**](V1SystemsTransactionRiskPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**401** | Missing or invalid API key |  -  |
**422** | ip_address is not a valid IPv4 or IPv6 address, or card_bin is missing |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

