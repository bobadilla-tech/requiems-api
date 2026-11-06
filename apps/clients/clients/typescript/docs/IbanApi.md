# IbanApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1FinanceIbanBatchPost**](#v1financeibanbatchpost) | **POST** /v1/finance/iban/batch | Batch Validate IBANs|
|[**v1FinanceIbanIbanGet**](#v1financeibanibanget) | **GET** /v1/finance/iban/{iban} | Validate IBAN|

# **v1FinanceIbanBatchPost**
> V1FinanceIbanBatchPost200Response v1FinanceIbanBatchPost(v1FinanceIbanBatchPostRequest)

Validates up to 50 iban numbers in a single request. Results are returned in the same order as the input.

### Example

```typescript
import {
    IbanApi,
    Configuration,
    V1FinanceIbanBatchPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new IbanApi(configuration);

let v1FinanceIbanBatchPostRequest: V1FinanceIbanBatchPostRequest; //

const { status, data } = await apiInstance.v1FinanceIbanBatchPost(
    v1FinanceIbanBatchPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1FinanceIbanBatchPostRequest** | **V1FinanceIbanBatchPostRequest**|  | |


### Return type

**V1FinanceIbanBatchPost200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**422** | The numbers array is missing, empty, or contains more than 50 items. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1FinanceIbanIbanGet**
> V1FinanceIbanIbanGet200Response v1FinanceIbanIbanGet()

Validates an IBAN and returns the country, bank code, and account number. Spaces in the input are stripped automatically. Always returns HTTP 200 — check the valid field to determine whether the IBAN is valid.

### Example

```typescript
import {
    IbanApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new IbanApi(configuration);

let iban: string; //The IBAN to validate. Spaces are stripped. Case-insensitive. (default to undefined)

const { status, data } = await apiInstance.v1FinanceIbanIbanGet(
    iban
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **iban** | [**string**] | The IBAN to validate. Spaces are stripped. Case-insensitive. | defaults to undefined|


### Return type

**V1FinanceIbanIbanGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**500** | Unexpected server error (e.g. database unreachable). |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

