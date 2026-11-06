# openapi_client.PasswordGeneratorApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_technology_password_get**](PasswordGeneratorApi.md#v1_technology_password_get) | **GET** /v1/technology/password | Generate Password


# **v1_technology_password_get**
> V1TechnologyPasswordGet200Response v1_technology_password_get(length=length, uppercase=uppercase, numbers=numbers, symbols=symbols)

Generate Password

Generate a cryptographically secure random password with customizable character sets and length

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_technology_password_get200_response import V1TechnologyPasswordGet200Response
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
    api_instance = openapi_client.PasswordGeneratorApi(api_client)
    length = 16 # int | Password length (8-128 characters) (optional)
    uppercase = true # bool | Include uppercase letters (A-Z) (optional)
    numbers = true # bool | Include numbers (0-9) (optional)
    symbols = true # bool | Include special characters (!@#$%^&*()-_=+[]{}|;:,.<>?) (optional)

    try:
        # Generate Password
        api_response = api_instance.v1_technology_password_get(length=length, uppercase=uppercase, numbers=numbers, symbols=symbols)
        print("The response of PasswordGeneratorApi->v1_technology_password_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling PasswordGeneratorApi->v1_technology_password_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **length** | **int**| Password length (8-128 characters) | [optional] 
 **uppercase** | **bool**| Include uppercase letters (A-Z) | [optional] 
 **numbers** | **bool**| Include numbers (0-9) | [optional] 
 **symbols** | **bool**| Include special characters (!@#$%^&amp;*()-_&#x3D;+[]{}|;:,.&lt;&gt;?) | [optional] 

### Return type

[**V1TechnologyPasswordGet200Response**](V1TechnologyPasswordGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | The length parameter is out of valid range (8-128) |  -  |
**500** | Failed to generate password (rare cryptographic failure) |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

