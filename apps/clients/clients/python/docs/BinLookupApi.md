# openapi_client.BinLookupApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_finance_bin_bin_get**](BinLookupApi.md#v1_finance_bin_bin_get) | **GET** /v1/finance/bin/{bin} | BIN Lookup


# **v1_finance_bin_bin_get**
> V1FinanceBinBinGet200Response v1_finance_bin_bin_get(bin)

BIN Lookup

Returns card metadata for the given 6–8 digit BIN prefix.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_finance_bin_bin_get200_response import V1FinanceBinBinGet200Response
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
    api_instance = openapi_client.BinLookupApi(api_client)
    bin = '424242' # str | 6–8 digit Bank Identification Number. Dashes and spaces are stripped automatically.

    try:
        # BIN Lookup
        api_response = api_instance.v1_finance_bin_bin_get(bin)
        print("The response of BinLookupApi->v1_finance_bin_bin_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling BinLookupApi->v1_finance_bin_bin_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bin** | **str**| 6–8 digit Bank Identification Number. Dashes and spaces are stripped automatically. | 

### Return type

[**V1FinanceBinBinGet200Response**](V1FinanceBinBinGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | BIN is not 6–8 digits or contains non-digit characters. |  -  |
**404** | BIN prefix not found in the database. |  -  |
**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

