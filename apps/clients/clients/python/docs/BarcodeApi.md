# openapi_client.BarcodeApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_technology_barcode_base64_get**](BarcodeApi.md#v1_technology_barcode_base64_get) | **GET** /v1/technology/barcode/base64 | Generate Barcode (Base64 JSON)
[**v1_technology_barcode_get**](BarcodeApi.md#v1_technology_barcode_get) | **GET** /v1/technology/barcode | Generate Barcode (PNG)


# **v1_technology_barcode_base64_get**
> V1TechnologyBarcodeBase64Get200Response v1_technology_barcode_base64_get(data, type)

Generate Barcode (Base64 JSON)

Returns a JSON envelope containing the barcode as a base64-encoded PNG string, along with its type and dimensions.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_technology_barcode_base64_get200_response import V1TechnologyBarcodeBase64Get200Response
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
    api_instance = openapi_client.BarcodeApi(api_client)
    data = '123456789' # str | The text or numeric string to encode in the barcode
    type = 'code128' # str | Barcode format: code128, code93, code39, ean8, ean13

    try:
        # Generate Barcode (Base64 JSON)
        api_response = api_instance.v1_technology_barcode_base64_get(data, type)
        print("The response of BarcodeApi->v1_technology_barcode_base64_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling BarcodeApi->v1_technology_barcode_base64_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | **str**| The text or numeric string to encode in the barcode | 
 **type** | **str**| Barcode format: code128, code93, code39, ean8, ean13 | 

### Return type

[**V1TechnologyBarcodeBase64Get200Response**](V1TechnologyBarcodeBase64Get200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | Missing or invalid parameters (e.g. data not provided, unsupported type) |  -  |
**422** | Data is invalid for the chosen barcode type (e.g. wrong digit count for EAN-8/EAN-13, non-numeric EAN data) |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_technology_barcode_get**
> v1_technology_barcode_get(data, type)

Generate Barcode (PNG)

Returns a raw PNG image of the barcode. Ideal for direct embedding or file download.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
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
    api_instance = openapi_client.BarcodeApi(api_client)
    data = '123456789' # str | The text or numeric string to encode in the barcode
    type = 'code128' # str | Barcode format: code128, code93, code39, ean8, ean13

    try:
        # Generate Barcode (PNG)
        api_instance.v1_technology_barcode_get(data, type)
    except Exception as e:
        print("Exception when calling BarcodeApi->v1_technology_barcode_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | **str**| The text or numeric string to encode in the barcode | 
 **type** | **str**| Barcode format: code128, code93, code39, ean8, ean13 | 

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
**400** | Missing or invalid parameters (e.g. data not provided, unsupported type) |  -  |
**422** | Data is invalid for the chosen barcode type (e.g. wrong digit count for EAN-8/EAN-13, non-numeric EAN data) |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

