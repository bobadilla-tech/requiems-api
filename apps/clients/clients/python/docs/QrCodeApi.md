# openapi_client.QrCodeApi

All URIs are relative to *https://requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_technology_qr_base64_batch_post**](QrCodeApi.md#v1_technology_qr_base64_batch_post) | **POST** /v1/technology/qr/base64/batch | Batch Generate QR Codes (Base64)
[**v1_technology_qr_base64_get**](QrCodeApi.md#v1_technology_qr_base64_get) | **GET** /v1/technology/qr/base64 | Generate QR Code (Base64 JSON)
[**v1_technology_qr_get**](QrCodeApi.md#v1_technology_qr_get) | **GET** /v1/technology/qr | Generate QR Code (PNG)


# **v1_technology_qr_base64_batch_post**
> V1TechnologyQrBase64BatchPost200Response v1_technology_qr_base64_batch_post(v1_technology_qr_base64_batch_post_request)

Batch Generate QR Codes (Base64)

Generate up to 50 base64-encoded QR codes in a single request. Results are returned in input order. The PNG endpoint has no batch variant — use this endpoint for batch generation.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_technology_qr_base64_batch_post200_response import V1TechnologyQrBase64BatchPost200Response
from openapi_client.models.v1_technology_qr_base64_batch_post_request import V1TechnologyQrBase64BatchPostRequest
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
    api_instance = openapi_client.QrCodeApi(api_client)
    v1_technology_qr_base64_batch_post_request = openapi_client.V1TechnologyQrBase64BatchPostRequest() # V1TechnologyQrBase64BatchPostRequest | 

    try:
        # Batch Generate QR Codes (Base64)
        api_response = api_instance.v1_technology_qr_base64_batch_post(v1_technology_qr_base64_batch_post_request)
        print("The response of QrCodeApi->v1_technology_qr_base64_batch_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling QrCodeApi->v1_technology_qr_base64_batch_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_technology_qr_base64_batch_post_request** | [**V1TechnologyQrBase64BatchPostRequest**](V1TechnologyQrBase64BatchPostRequest.md)|  | 

### Return type

[**V1TechnologyQrBase64BatchPost200Response**](V1TechnologyQrBase64BatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | The request body is missing or malformed. |  -  |
**422** | The items array is missing, empty, or exceeds 50 items. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_technology_qr_base64_get**
> V1TechnologyQrBase64Get200Response v1_technology_qr_base64_get(data, size=size, recovery=recovery)

Generate QR Code (Base64 JSON)

Returns a JSON envelope containing the QR code as a base64-encoded PNG string, along with its dimensions.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_technology_qr_base64_get200_response import V1TechnologyQrBase64Get200Response
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
    api_instance = openapi_client.QrCodeApi(api_client)
    data = 'https://example.com' # str | The text or URL to encode in the QR code
    size = 200 # int | Image size in pixels (default: 256, min: 50, max: 1000) (optional)
    recovery = 'highest' # str | Error-correction level: low (7%), medium (15%), high (25%), highest (30%). Default: medium (optional)

    try:
        # Generate QR Code (Base64 JSON)
        api_response = api_instance.v1_technology_qr_base64_get(data, size=size, recovery=recovery)
        print("The response of QrCodeApi->v1_technology_qr_base64_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling QrCodeApi->v1_technology_qr_base64_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | **str**| The text or URL to encode in the QR code | 
 **size** | **int**| Image size in pixels (default: 256, min: 50, max: 1000) | [optional] 
 **recovery** | **str**| Error-correction level: low (7%), medium (15%), high (25%), highest (30%). Default: medium | [optional] 

### Return type

[**V1TechnologyQrBase64Get200Response**](V1TechnologyQrBase64Get200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | Missing or invalid parameters (e.g. data not provided, size out of range, unknown recovery level) |  -  |
**500** | Failed to generate QR code |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_technology_qr_get**
> v1_technology_qr_get(data, size=size, recovery=recovery)

Generate QR Code (PNG)

Returns a raw PNG image of the QR code. Ideal for direct embedding or file download.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
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
    api_instance = openapi_client.QrCodeApi(api_client)
    data = 'https://example.com' # str | The text or URL to encode in the QR code
    size = 200 # int | Image size in pixels (default: 256, min: 50, max: 1000) (optional)
    recovery = 'high' # str | Error-correction level: low (7%), medium (15%), high (25%), highest (30%). Higher levels are more robust to physical damage but produce larger images. Default: medium (optional)

    try:
        # Generate QR Code (PNG)
        api_instance.v1_technology_qr_get(data, size=size, recovery=recovery)
    except Exception as e:
        print("Exception when calling QrCodeApi->v1_technology_qr_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | **str**| The text or URL to encode in the QR code | 
 **size** | **int**| Image size in pixels (default: 256, min: 50, max: 1000) | [optional] 
 **recovery** | **str**| Error-correction level: low (7%), medium (15%), high (25%), highest (30%). Higher levels are more robust to physical damage but produce larger images. Default: medium | [optional] 

### Return type

void (empty response body)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | Missing or invalid parameters (e.g. data not provided, size out of range, unknown recovery level) |  -  |
**500** | Failed to generate QR code |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

