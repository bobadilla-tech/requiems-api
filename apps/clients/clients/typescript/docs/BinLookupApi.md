# BinLookupApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1FinanceBinBinGet**](#v1financebinbinget) | **GET** /v1/finance/bin/{bin} | BIN Lookup|

# **v1FinanceBinBinGet**
> V1FinanceBinBinGet200Response v1FinanceBinBinGet()

Returns card metadata for the given 6–8 digit BIN prefix.

### Example

```typescript
import {
    BinLookupApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new BinLookupApi(configuration);

let bin: string; //6–8 digit Bank Identification Number. Dashes and spaces are stripped automatically. (default to undefined)

const { status, data } = await apiInstance.v1FinanceBinBinGet(
    bin
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **bin** | [**string**] | 6–8 digit Bank Identification Number. Dashes and spaces are stripped automatically. | defaults to undefined|


### Return type

**V1FinanceBinBinGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | BIN is not 6–8 digits or contains non-digit characters. |  -  |
|**404** | BIN prefix not found in the database. |  -  |
|**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

