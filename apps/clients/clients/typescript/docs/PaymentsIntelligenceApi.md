# PaymentsIntelligenceApi

All URIs are relative to *https://requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1SystemsPaymentValidatePost**](#v1systemspaymentvalidatepost) | **POST** /v1/systems/payment/validate | Validate Payment|
|[**v1SystemsTransactionRiskPost**](#v1systemstransactionriskpost) | **POST** /v1/systems/transaction/risk | Score Transaction Risk|

# **v1SystemsPaymentValidatePost**
> V1SystemsPaymentValidatePost200Response v1SystemsPaymentValidatePost(v1SystemsPaymentValidatePostRequest)

Validate and cross-check BIN, IBAN, and SWIFT identifiers in one call. Returns per-field results and a consistency check that flags country and bank mismatches.

### Example

```typescript
import {
    PaymentsIntelligenceApi,
    Configuration,
    V1SystemsPaymentValidatePostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new PaymentsIntelligenceApi(configuration);

let v1SystemsPaymentValidatePostRequest: V1SystemsPaymentValidatePostRequest; //

const { status, data } = await apiInstance.v1SystemsPaymentValidatePost(
    v1SystemsPaymentValidatePostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1SystemsPaymentValidatePostRequest** | **V1SystemsPaymentValidatePostRequest**|  | |


### Return type

**V1SystemsPaymentValidatePost200Response**

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
|**422** | All of bin, iban, and swift were omitted. At least one is required. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1SystemsTransactionRiskPost**
> V1SystemsTransactionRiskPost200Response v1SystemsTransactionRiskPost(v1SystemsTransactionRiskPostRequest)

Score a transaction for fraud risk by cross-checking the card BIN country against IP geolocation and billing country. Detects VPN, proxy, and TOR use.

### Example

```typescript
import {
    PaymentsIntelligenceApi,
    Configuration,
    V1SystemsTransactionRiskPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new PaymentsIntelligenceApi(configuration);

let v1SystemsTransactionRiskPostRequest: V1SystemsTransactionRiskPostRequest; //

const { status, data } = await apiInstance.v1SystemsTransactionRiskPost(
    v1SystemsTransactionRiskPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1SystemsTransactionRiskPostRequest** | **V1SystemsTransactionRiskPostRequest**|  | |


### Return type

**V1SystemsTransactionRiskPost200Response**

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
|**422** | ip_address is not a valid IPv4 or IPv6 address, or card_bin is missing |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

