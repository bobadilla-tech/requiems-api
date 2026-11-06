# openapi_client.PhoneValidationApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_validation_phone_batch_post**](PhoneValidationApi.md#v1_validation_phone_batch_post) | **POST** /v1/validation/phone/batch | Batch Validate Phone Numbers
[**v1_validation_phone_get**](PhoneValidationApi.md#v1_validation_phone_get) | **GET** /v1/validation/phone | Validate Phone Number


# **v1_validation_phone_batch_post**
> V1FinanceIbanBatchPost200Response v1_validation_phone_batch_post(v1_validation_phone_batch_post_request)

Batch Validate Phone Numbers

Validates up to 50 phone numbers in a single request. Results are returned in the same order as the input.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_finance_iban_batch_post200_response import V1FinanceIbanBatchPost200Response
from openapi_client.models.v1_validation_phone_batch_post_request import V1ValidationPhoneBatchPostRequest
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
    api_instance = openapi_client.PhoneValidationApi(api_client)
    v1_validation_phone_batch_post_request = openapi_client.V1ValidationPhoneBatchPostRequest() # V1ValidationPhoneBatchPostRequest | 

    try:
        # Batch Validate Phone Numbers
        api_response = api_instance.v1_validation_phone_batch_post(v1_validation_phone_batch_post_request)
        print("The response of PhoneValidationApi->v1_validation_phone_batch_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling PhoneValidationApi->v1_validation_phone_batch_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_validation_phone_batch_post_request** | [**V1ValidationPhoneBatchPostRequest**](V1ValidationPhoneBatchPostRequest.md)|  | 

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

# **v1_validation_phone_get**
> V1ValidationPhoneGet200Response v1_validation_phone_get(number)

Validate Phone Number

Validates a single phone number and returns its country, type, formatted representation, carrier, and VOIP/virtual risk flags.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_validation_phone_get200_response import V1ValidationPhoneGet200Response
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
    api_instance = openapi_client.PhoneValidationApi(api_client)
    number = '+447400123456' # str | The phone number to validate. Must include the country calling code (e.g. +12015551234).

    try:
        # Validate Phone Number
        api_response = api_instance.v1_validation_phone_get(number)
        print("The response of PhoneValidationApi->v1_validation_phone_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling PhoneValidationApi->v1_validation_phone_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **number** | **str**| The phone number to validate. Must include the country calling code (e.g. +12015551234). | 

### Return type

[**V1ValidationPhoneGet200Response**](V1ValidationPhoneGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | The number query parameter is missing. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

