# Base64Api

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1TechnologyBase64DecodePost**](#v1technologybase64decodepost) | **POST** /v1/technology/base64/decode | Decode|
|[**v1TechnologyBase64EncodePost**](#v1technologybase64encodepost) | **POST** /v1/technology/base64/encode | Encode|

# **v1TechnologyBase64DecodePost**
> V1TechnologyBase64DecodePost200Response v1TechnologyBase64DecodePost(v1TechnologyBase64DecodePostRequest)

Decode a Base64-encoded string back to plain text

### Example

```typescript
import {
    Base64Api,
    Configuration,
    V1TechnologyBase64DecodePostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new Base64Api(configuration);

let v1TechnologyBase64DecodePostRequest: V1TechnologyBase64DecodePostRequest; //

const { status, data } = await apiInstance.v1TechnologyBase64DecodePost(
    v1TechnologyBase64DecodePostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1TechnologyBase64DecodePostRequest** | **V1TechnologyBase64DecodePostRequest**|  | |


### Return type

**V1TechnologyBase64DecodePost200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | Missing or empty value field |  -  |
|**422** | The value is not valid Base64 and cannot be decoded |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1TechnologyBase64EncodePost**
> V1TechnologyBase64EncodePost200Response v1TechnologyBase64EncodePost(v1TechnologyBase64EncodePostRequest)

Encode a plain-text string to Base64

### Example

```typescript
import {
    Base64Api,
    Configuration,
    V1TechnologyBase64EncodePostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new Base64Api(configuration);

let v1TechnologyBase64EncodePostRequest: V1TechnologyBase64EncodePostRequest; //

const { status, data } = await apiInstance.v1TechnologyBase64EncodePost(
    v1TechnologyBase64EncodePostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1TechnologyBase64EncodePostRequest** | **V1TechnologyBase64EncodePostRequest**|  | |


### Return type

**V1TechnologyBase64EncodePost200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | Missing or empty value field |  -  |
|**422** | Validation constraint on the variant field (must be standard or url) |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

