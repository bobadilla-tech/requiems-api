# openapi_client.EmailValidateApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_validation_email_batch_post**](EmailValidateApi.md#v1_validation_email_batch_post) | **POST** /v1/validation/email/batch | Validate Emails (Batch)
[**v1_validation_email_post**](EmailValidateApi.md#v1_validation_email_post) | **POST** /v1/validation/email | Validate Email


# **v1_validation_email_batch_post**
> V1ValidationEmailBatchPost200Response v1_validation_email_batch_post(v1_validation_email_batch_post_request)

Validate Emails (Batch)

Validates up to 50 email addresses in a single request. Each email is processed independently and returns a full validation breakdown (syntax, MX record, disposable check, normalization, and typo suggestion). Invalid emails do not fail the request. Billing: 1 credit per email.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_validation_email_batch_post200_response import V1ValidationEmailBatchPost200Response
from openapi_client.models.v1_validation_email_batch_post_request import V1ValidationEmailBatchPostRequest
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
    api_instance = openapi_client.EmailValidateApi(api_client)
    v1_validation_email_batch_post_request = openapi_client.V1ValidationEmailBatchPostRequest() # V1ValidationEmailBatchPostRequest | 

    try:
        # Validate Emails (Batch)
        api_response = api_instance.v1_validation_email_batch_post(v1_validation_email_batch_post_request)
        print("The response of EmailValidateApi->v1_validation_email_batch_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling EmailValidateApi->v1_validation_email_batch_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_validation_email_batch_post_request** | [**V1ValidationEmailBatchPostRequest**](V1ValidationEmailBatchPostRequest.md)|  | 

### Return type

[**V1ValidationEmailBatchPost200Response**](V1ValidationEmailBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | Invalid JSON, malformed request body, or unexpected field types. |  -  |
**422** | Valid JSON body that fails field validation (empty array or more than 50 emails). |  -  |
**500** | Unexpected server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_validation_email_post**
> V1ValidationEmailPost200Response v1_validation_email_post(v1_validation_email_post_request)

Validate Email

Validates a single email address and returns a full breakdown of syntax validity, MX record status, disposable domain check, normalized form, and any typo suggestion.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_validation_email_post200_response import V1ValidationEmailPost200Response
from openapi_client.models.v1_validation_email_post_request import V1ValidationEmailPostRequest
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
    api_instance = openapi_client.EmailValidateApi(api_client)
    v1_validation_email_post_request = openapi_client.V1ValidationEmailPostRequest() # V1ValidationEmailPostRequest | 

    try:
        # Validate Email
        api_response = api_instance.v1_validation_email_post(v1_validation_email_post_request)
        print("The response of EmailValidateApi->v1_validation_email_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling EmailValidateApi->v1_validation_email_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_validation_email_post_request** | [**V1ValidationEmailPostRequest**](V1ValidationEmailPostRequest.md)|  | 

### Return type

[**V1ValidationEmailPost200Response**](V1ValidationEmailPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | The request body is missing, not valid JSON, or contains unknown fields. |  -  |
**422** | The email field is missing from the request body. |  -  |
**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

