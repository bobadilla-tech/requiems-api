# PasswordGeneratorApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1TechnologyPasswordGet**](#v1technologypasswordget) | **GET** /v1/technology/password | Generate Password|

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

