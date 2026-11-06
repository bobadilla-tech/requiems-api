# SwiftCodeApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1FinanceSwiftCodeGet**](#v1financeswiftcodeget) | **GET** /v1/finance/swift/{code} | Get SWIFT Code|
|[**v1FinanceSwiftGet**](#v1financeswiftget) | **GET** /v1/finance/swift | List SWIFT Codes|

# **v1FinanceSwiftCodeGet**
> V1FinanceSwiftCodeGet200Response v1FinanceSwiftCodeGet()

Look up bank metadata for a SWIFT/BIC code.

### Example

```typescript
import {
    SwiftCodeApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new SwiftCodeApi(configuration);

let code: string; //SWIFT/BIC code (8 or 11 alphanumeric characters) (default to undefined)

const { status, data } = await apiInstance.v1FinanceSwiftCodeGet(
    code
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **code** | [**string**] | SWIFT/BIC code (8 or 11 alphanumeric characters) | defaults to undefined|


### Return type

**V1FinanceSwiftCodeGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | Invalid SWIFT/BIC format (must be 8 or 11 valid characters). |  -  |
|**404** | SWIFT/BIC code not found in the dataset. |  -  |
|**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1FinanceSwiftGet**
> V1HealthExercisesRandomGet200Response v1FinanceSwiftGet()

List SWIFT records with optional filters and pagination.

### Example

```typescript
import {
    SwiftCodeApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new SwiftCodeApi(configuration);

let countryCode: string; //Optional 2-letter country code filter (e.g. DE, US) (optional) (default to undefined)
let bankCode: string; //Optional 4-letter bank code filter (e.g. DEUT) (optional) (default to undefined)
let q: string; //Optional text search across swift_code, bank_name, and city (optional) (default to undefined)
let limit: number; //Max rows to return (default 50, max 200) (optional) (default to undefined)
let offset: number; //Number of rows to skip (default 0) (optional) (default to undefined)

const { status, data } = await apiInstance.v1FinanceSwiftGet(
    countryCode,
    bankCode,
    q,
    limit,
    offset
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **countryCode** | [**string**] | Optional 2-letter country code filter (e.g. DE, US) | (optional) defaults to undefined|
| **bankCode** | [**string**] | Optional 4-letter bank code filter (e.g. DEUT) | (optional) defaults to undefined|
| **q** | [**string**] | Optional text search across swift_code, bank_name, and city | (optional) defaults to undefined|
| **limit** | [**number**] | Max rows to return (default 50, max 200) | (optional) defaults to undefined|
| **offset** | [**number**] | Number of rows to skip (default 0) | (optional) defaults to undefined|


### Return type

**V1HealthExercisesRandomGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | Invalid filter or pagination parameter. |  -  |
|**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

