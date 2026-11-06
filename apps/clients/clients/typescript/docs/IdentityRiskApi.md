# IdentityRiskApi

All URIs are relative to *https://requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1SystemsRiskScorePost**](#v1systemsriskscorepost) | **POST** /v1/systems/risk/score | Score Risk|
|[**v1SystemsSignupProtectPost**](#v1systemssignupprotectpost) | **POST** /v1/systems/signup/protect | Protect Signup|
|[**v1SystemsUserVerifyPost**](#v1systemsuserverifypost) | **POST** /v1/systems/user/verify | Verify User|

# **v1SystemsRiskScorePost**
> V1SystemsRiskScorePost200Response v1SystemsRiskScorePost(v1SystemsRiskScorePostRequest)

Score a user for risk without the full signal breakdown. Lower latency than /signup/protect, suited for background re-scoring and rate limiting.

### Example

```typescript
import {
    IdentityRiskApi,
    Configuration,
    V1SystemsRiskScorePostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new IdentityRiskApi(configuration);

let v1SystemsRiskScorePostRequest: V1SystemsRiskScorePostRequest; //

const { status, data } = await apiInstance.v1SystemsRiskScorePost(
    v1SystemsRiskScorePostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1SystemsRiskScorePostRequest** | **V1SystemsRiskScorePostRequest**|  | |


### Return type

**V1SystemsRiskScorePost200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**401** | Missing or invalid API key |  -  |
|**422** | All of email, phone, and ip_address were omitted |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1SystemsSignupProtectPost**
> V1SystemsSignupProtectPost200Response v1SystemsSignupProtectPost(v1SystemsSignupProtectPostRequest)

Evaluate a new user at signup. Returns a full risk decision with per-signal breakdown across email, phone, and IP.

### Example

```typescript
import {
    IdentityRiskApi,
    Configuration,
    V1SystemsSignupProtectPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new IdentityRiskApi(configuration);

let v1SystemsSignupProtectPostRequest: V1SystemsSignupProtectPostRequest; //

const { status, data } = await apiInstance.v1SystemsSignupProtectPost(
    v1SystemsSignupProtectPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1SystemsSignupProtectPostRequest** | **V1SystemsSignupProtectPostRequest**|  | |


### Return type

**V1SystemsSignupProtectPost200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**401** | Missing or invalid API key in the requiems-api-key header |  -  |
|**422** | All of email, phone, and ip_address were omitted. At least one is required. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1SystemsUserVerifyPost**
> V1SystemsUserVerifyPost200Response v1SystemsUserVerifyPost(v1SystemsUserVerifyPostRequest)

Deep-verify an email address using domain-level signals including WHOIS age, MX records, and domain availability. Optional IP check. Use for high-value or suspicious accounts.

### Example

```typescript
import {
    IdentityRiskApi,
    Configuration,
    V1SystemsUserVerifyPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new IdentityRiskApi(configuration);

let v1SystemsUserVerifyPostRequest: V1SystemsUserVerifyPostRequest; //

const { status, data } = await apiInstance.v1SystemsUserVerifyPost(
    v1SystemsUserVerifyPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1SystemsUserVerifyPostRequest** | **V1SystemsUserVerifyPostRequest**|  | |


### Return type

**V1SystemsUserVerifyPost200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**401** | Missing or invalid API key |  -  |
|**422** | email field is missing or empty |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

