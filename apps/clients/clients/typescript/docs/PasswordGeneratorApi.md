# PasswordGeneratorApi

All URIs are relative to *https://requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1TechnologyPasswordBatchPost**](#v1technologypasswordbatchpost) | **POST** /v1/technology/password/batch | Batch Generate Passwords|
|[**v1TechnologyPasswordGet**](#v1technologypasswordget) | **GET** /v1/technology/password | Generate Password|

# **v1TechnologyPasswordBatchPost**
> V1TechnologyPasswordBatchPost200Response v1TechnologyPasswordBatchPost(v1TechnologyPasswordBatchPostRequest)

Generate up to 50 passwords in a single request. Each item can have its own length and character set options. Results are returned in input order.

### Example

```typescript
import {
    PasswordGeneratorApi,
    Configuration,
    V1TechnologyPasswordBatchPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new PasswordGeneratorApi(configuration);

let v1TechnologyPasswordBatchPostRequest: V1TechnologyPasswordBatchPostRequest; //

const { status, data } = await apiInstance.v1TechnologyPasswordBatchPost(
    v1TechnologyPasswordBatchPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1TechnologyPasswordBatchPostRequest** | **V1TechnologyPasswordBatchPostRequest**|  | |


### Return type

**V1TechnologyPasswordBatchPost200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | The request body is missing or malformed. |  -  |
|**422** | The items array is missing, empty, or exceeds 50 items. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1TechnologyPasswordGet**
> V1TechnologyPasswordGet200Response v1TechnologyPasswordGet()

Generate a cryptographically secure random password with customizable character sets and length

### Example

```typescript
import {
    PasswordGeneratorApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new PasswordGeneratorApi(configuration);

let length: number; //Password length (8-128 characters) (optional) (default to undefined)
let uppercase: boolean; //Include uppercase letters (A-Z) (optional) (default to undefined)
let numbers: boolean; //Include numbers (0-9) (optional) (default to undefined)
let symbols: boolean; //Include special characters (!@#$%^&*()-_=+[]{}|;:,.<>?) (optional) (default to undefined)

const { status, data } = await apiInstance.v1TechnologyPasswordGet(
    length,
    uppercase,
    numbers,
    symbols
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **length** | [**number**] | Password length (8-128 characters) | (optional) defaults to undefined|
| **uppercase** | [**boolean**] | Include uppercase letters (A-Z) | (optional) defaults to undefined|
| **numbers** | [**boolean**] | Include numbers (0-9) | (optional) defaults to undefined|
| **symbols** | [**boolean**] | Include special characters (!@#$%^&amp;*()-_&#x3D;+[]{}|;:,.&lt;&gt;?) | (optional) defaults to undefined|


### Return type

**V1TechnologyPasswordGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | The length parameter is out of valid range (8-128) |  -  |
|**500** | Failed to generate password (rare cryptographic failure) |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

