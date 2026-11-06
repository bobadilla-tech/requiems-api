# PhoneValidationApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1ValidationPhoneBatchPost**](#v1validationphonebatchpost) | **POST** /v1/validation/phone/batch | Batch Validate Phone Numbers|
|[**v1ValidationPhoneGet**](#v1validationphoneget) | **GET** /v1/validation/phone | Validate Phone Number|

# **v1ValidationPhoneBatchPost**
> V1FinanceIbanBatchPost200Response v1ValidationPhoneBatchPost(v1ValidationPhoneBatchPostRequest)

Validates up to 50 phone numbers in a single request. Results are returned in the same order as the input.

### Example

```typescript
import {
    PhoneValidationApi,
    Configuration,
    V1ValidationPhoneBatchPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new PhoneValidationApi(configuration);

let v1ValidationPhoneBatchPostRequest: V1ValidationPhoneBatchPostRequest; //

const { status, data } = await apiInstance.v1ValidationPhoneBatchPost(
    v1ValidationPhoneBatchPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1ValidationPhoneBatchPostRequest** | **V1ValidationPhoneBatchPostRequest**|  | |


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

# **v1ValidationPhoneGet**
> V1ValidationPhoneGet200Response v1ValidationPhoneGet()

Validates a single phone number and returns its country, type, formatted representation, carrier, and VOIP/virtual risk flags.

### Example

```typescript
import {
    PhoneValidationApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new PhoneValidationApi(configuration);

let number: string; //The phone number to validate. Must include the country calling code (e.g. +12015551234). (default to undefined)

const { status, data } = await apiInstance.v1ValidationPhoneGet(
    number
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **number** | [**string**] | The phone number to validate. Must include the country calling code (e.g. +12015551234). | defaults to undefined|


### Return type

**V1ValidationPhoneGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | The number query parameter is missing. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

