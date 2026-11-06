# V1FinanceCryptoSymbolGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**symbol** | **string** | Uppercase ticker symbol | [optional] [default to undefined]
**name** | **string** | Full coin name | [optional] [default to undefined]
**price_usd** | **number** | Current price in USD | [optional] [default to undefined]
**change_24h** | **number** | Price change over the last 24 hours as a percentage | [optional] [default to undefined]
**market_cap** | **number** | Total market capitalisation in USD | [optional] [default to undefined]
**volume_24h** | **number** | Total trading volume over the last 24 hours in USD | [optional] [default to undefined]

## Example

```typescript
import { V1FinanceCryptoSymbolGet200ResponseData } from 'api-client';

const instance: V1FinanceCryptoSymbolGet200ResponseData = {
    symbol,
    name,
    price_usd,
    change_24h,
    market_cap,
    volume_24h,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
