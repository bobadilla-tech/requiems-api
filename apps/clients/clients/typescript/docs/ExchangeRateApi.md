# ExchangeRateApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1FinanceConvertGet**](#v1financeconvertget) | **GET** /v1/finance/convert | Convert Currency|
|[**v1FinanceExchangeRateGet**](#v1financeexchangerateget) | **GET** /v1/finance/exchange-rate | Get Exchange Rate|

# **v1FinanceConvertGet**
> V1FinanceConvertGet200Response v1FinanceConvertGet()

Converts an amount from one currency to another and returns the rate alongside the converted value.

### Example

```typescript
import {
    ExchangeRateApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new ExchangeRateApi(configuration);

let from: string; //ISO 4217 source currency code (3 letters, e.g. USD) (default to undefined)
let to: string; //ISO 4217 target currency code (3 letters, e.g. EUR) (default to undefined)
let amount: number; //Amount to convert. Must be greater than 0. (default to undefined)

const { status, data } = await apiInstance.v1FinanceConvertGet(
    from,
    to,
    amount
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **from** | [**string**] | ISO 4217 source currency code (3 letters, e.g. USD) | defaults to undefined|
| **to** | [**string**] | ISO 4217 target currency code (3 letters, e.g. EUR) | defaults to undefined|
| **amount** | [**number**] | Amount to convert. Must be greater than 0. | defaults to undefined|


### Return type

**V1FinanceConvertGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | A required parameter is missing, the currency code is not 3 alphabetic characters, or the amount is 0 or negative. |  -  |
|**422** | One or both currency codes are not recognised by the upstream data source. |  -  |
|**503** | The exchange rate data source is temporarily unavailable. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1FinanceExchangeRateGet**
> V1FinanceExchangeRateGet200Response v1FinanceExchangeRateGet()

Returns the current exchange rate between two currencies.

### Example

```typescript
import {
    ExchangeRateApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new ExchangeRateApi(configuration);

let from: string; //ISO 4217 source currency code (3 letters, e.g. USD) (default to undefined)
let to: string; //ISO 4217 target currency code (3 letters, e.g. EUR) (default to undefined)

const { status, data } = await apiInstance.v1FinanceExchangeRateGet(
    from,
    to
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **from** | [**string**] | ISO 4217 source currency code (3 letters, e.g. USD) | defaults to undefined|
| **to** | [**string**] | ISO 4217 target currency code (3 letters, e.g. EUR) | defaults to undefined|


### Return type

**V1FinanceExchangeRateGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | A required parameter is missing or the currency code is not exactly 3 alphabetic characters. |  -  |
|**422** | One or both currency codes are not recognised by the upstream data source. |  -  |
|**503** | The exchange rate data source is temporarily unavailable. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

