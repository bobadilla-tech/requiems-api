# CryptoApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1FinanceCryptoSymbolGet**](#v1financecryptosymbolget) | **GET** /v1/finance/crypto/{symbol} | Get Crypto Price|

# **v1FinanceCryptoSymbolGet**
> V1FinanceCryptoSymbolGet200Response v1FinanceCryptoSymbolGet()

Returns current price data for the given cryptocurrency symbol.

### Example

```typescript
import {
    CryptoApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new CryptoApi(configuration);

let symbol: string; //Uppercase ticker symbol (e.g. BTC, ETH, SOL) (default to undefined)

const { status, data } = await apiInstance.v1FinanceCryptoSymbolGet(
    symbol
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **symbol** | [**string**] | Uppercase ticker symbol (e.g. BTC, ETH, SOL) | defaults to undefined|


### Return type

**V1FinanceCryptoSymbolGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**422** | The symbol is not in the supported coin list. |  -  |
|**503** | CoinGecko is unavailable or returned an unexpected response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

