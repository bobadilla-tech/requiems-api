# openapi_client.Base64Api

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_technology_base64_decode_post**](Base64Api.md#v1_technology_base64_decode_post) | **POST** /v1/technology/base64/decode | Decode
[**v1_technology_base64_encode_post**](Base64Api.md#v1_technology_base64_encode_post) | **POST** /v1/technology/base64/encode | Encode


# **v1_technology_base64_decode_post**
> V1TechnologyBase64DecodePost200Response v1_technology_base64_decode_post(v1_technology_base64_decode_post_request)

Decode

Decode a Base64-encoded string back to plain text

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_technology_base64_decode_post200_response import V1TechnologyBase64DecodePost200Response
from openapi_client.models.v1_technology_base64_decode_post_request import V1TechnologyBase64DecodePostRequest
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
    api_instance = openapi_client.Base64Api(api_client)
    v1_technology_base64_decode_post_request = openapi_client.V1TechnologyBase64DecodePostRequest() # V1TechnologyBase64DecodePostRequest | 

    try:
        # Decode
        api_response = api_instance.v1_technology_base64_decode_post(v1_technology_base64_decode_post_request)
        print("The response of Base64Api->v1_technology_base64_decode_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling Base64Api->v1_technology_base64_decode_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_technology_base64_decode_post_request** | [**V1TechnologyBase64DecodePostRequest**](V1TechnologyBase64DecodePostRequest.md)|  | 

### Return type

[**V1TechnologyBase64DecodePost200Response**](V1TechnologyBase64DecodePost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | Missing or empty value field |  -  |
**422** | The value is not valid Base64 and cannot be decoded |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_technology_base64_encode_post**
> V1TechnologyBase64EncodePost200Response v1_technology_base64_encode_post(v1_technology_base64_encode_post_request)

Encode

Encode a plain-text string to Base64

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_technology_base64_encode_post200_response import V1TechnologyBase64EncodePost200Response
from openapi_client.models.v1_technology_base64_encode_post_request import V1TechnologyBase64EncodePostRequest
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
    api_instance = openapi_client.Base64Api(api_client)
    v1_technology_base64_encode_post_request = openapi_client.V1TechnologyBase64EncodePostRequest() # V1TechnologyBase64EncodePostRequest | 

    try:
        # Encode
        api_response = api_instance.v1_technology_base64_encode_post(v1_technology_base64_encode_post_request)
        print("The response of Base64Api->v1_technology_base64_encode_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling Base64Api->v1_technology_base64_encode_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_technology_base64_encode_post_request** | [**V1TechnologyBase64EncodePostRequest**](V1TechnologyBase64EncodePostRequest.md)|  | 

### Return type

[**V1TechnologyBase64EncodePost200Response**](V1TechnologyBase64EncodePost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | Missing or empty value field |  -  |
**422** | Validation constraint on the variant field (must be standard or url) |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

