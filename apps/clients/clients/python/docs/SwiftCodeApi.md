# openapi_client.SwiftCodeApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_finance_swift_code_get**](SwiftCodeApi.md#v1_finance_swift_code_get) | **GET** /v1/finance/swift/{code} | Get SWIFT Code
[**v1_finance_swift_get**](SwiftCodeApi.md#v1_finance_swift_get) | **GET** /v1/finance/swift | List SWIFT Codes


# **v1_finance_swift_code_get**
> V1FinanceSwiftCodeGet200Response v1_finance_swift_code_get(code)

Get SWIFT Code

Look up bank metadata for a SWIFT/BIC code.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_finance_swift_code_get200_response import V1FinanceSwiftCodeGet200Response
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
    api_instance = openapi_client.SwiftCodeApi(api_client)
    code = 'DEUTDEDB' # str | SWIFT/BIC code (8 or 11 alphanumeric characters)

    try:
        # Get SWIFT Code
        api_response = api_instance.v1_finance_swift_code_get(code)
        print("The response of SwiftCodeApi->v1_finance_swift_code_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SwiftCodeApi->v1_finance_swift_code_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **str**| SWIFT/BIC code (8 or 11 alphanumeric characters) | 

### Return type

[**V1FinanceSwiftCodeGet200Response**](V1FinanceSwiftCodeGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | Invalid SWIFT/BIC format (must be 8 or 11 valid characters). |  -  |
**404** | SWIFT/BIC code not found in the dataset. |  -  |
**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_finance_swift_get**
> V1HealthExercisesRandomGet200Response v1_finance_swift_get(country_code=country_code, bank_code=bank_code, q=q, limit=limit, offset=offset)

List SWIFT Codes

List SWIFT records with optional filters and pagination.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_health_exercises_random_get200_response import V1HealthExercisesRandomGet200Response
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
    api_instance = openapi_client.SwiftCodeApi(api_client)
    country_code = 'DE' # str | Optional 2-letter country code filter (e.g. DE, US) (optional)
    bank_code = 'DEUT' # str | Optional 4-letter bank code filter (e.g. DEUT) (optional)
    q = 'deutsche' # str | Optional text search across swift_code, bank_name, and city (optional)
    limit = 50 # int | Max rows to return (default 50, max 200) (optional)
    offset = 0 # int | Number of rows to skip (default 0) (optional)

    try:
        # List SWIFT Codes
        api_response = api_instance.v1_finance_swift_get(country_code=country_code, bank_code=bank_code, q=q, limit=limit, offset=offset)
        print("The response of SwiftCodeApi->v1_finance_swift_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SwiftCodeApi->v1_finance_swift_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country_code** | **str**| Optional 2-letter country code filter (e.g. DE, US) | [optional] 
 **bank_code** | **str**| Optional 4-letter bank code filter (e.g. DEUT) | [optional] 
 **q** | **str**| Optional text search across swift_code, bank_name, and city | [optional] 
 **limit** | **int**| Max rows to return (default 50, max 200) | [optional] 
 **offset** | **int**| Number of rows to skip (default 0) | [optional] 

### Return type

[**V1HealthExercisesRandomGet200Response**](V1HealthExercisesRandomGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | Invalid filter or pagination parameter. |  -  |
**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

