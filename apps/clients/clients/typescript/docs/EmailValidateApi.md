# EmailValidateApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1ValidationEmailBatchPost**](#v1validationemailbatchpost) | **POST** /v1/validation/email/batch | Validate Emails (Batch)|
|[**v1ValidationEmailPost**](#v1validationemailpost) | **POST** /v1/validation/email | Validate Email|

# **v1ValidationEmailBatchPost**
> V1ValidationEmailBatchPost200Response v1ValidationEmailBatchPost(v1ValidationEmailBatchPostRequest)

Validates up to 50 email addresses in a single request. Each email is processed independently and returns a full validation breakdown (syntax, MX record, disposable check, normalization, and typo suggestion). Invalid emails do not fail the request. Billing: 1 credit per email.

### Example

```typescript
import {
    EmailValidateApi,
    Configuration,
    V1ValidationEmailBatchPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new EmailValidateApi(configuration);

let v1ValidationEmailBatchPostRequest: V1ValidationEmailBatchPostRequest; //

const { status, data } = await apiInstance.v1ValidationEmailBatchPost(
    v1ValidationEmailBatchPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1ValidationEmailBatchPostRequest** | **V1ValidationEmailBatchPostRequest**|  | |


### Return type

**V1ValidationEmailBatchPost200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | Invalid JSON, malformed request body, or unexpected field types. |  -  |
|**422** | Valid JSON body that fails field validation (empty array or more than 50 emails). |  -  |
|**500** | Unexpected server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1ValidationEmailPost**
> V1ValidationEmailPost200Response v1ValidationEmailPost(v1ValidationEmailPostRequest)

Validates a single email address and returns a full breakdown of syntax validity, MX record status, disposable domain check, normalized form, and any typo suggestion.

### Example

```typescript
import {
    EmailValidateApi,
    Configuration,
    V1ValidationEmailPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new EmailValidateApi(configuration);

let v1ValidationEmailPostRequest: V1ValidationEmailPostRequest; //

const { status, data } = await apiInstance.v1ValidationEmailPost(
    v1ValidationEmailPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1ValidationEmailPostRequest** | **V1ValidationEmailPostRequest**|  | |


### Return type

**V1ValidationEmailPost200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | The request body is missing, not valid JSON, or contains unknown fields. |  -  |
|**422** | The email field is missing from the request body. |  -  |
|**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

