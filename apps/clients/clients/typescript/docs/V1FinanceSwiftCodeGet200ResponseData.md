# V1FinanceSwiftCodeGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**swift_code** | **string** | Full 11-character SWIFT/BIC code | [optional] [default to undefined]
**bank_code** | **string** | Institution code (characters 1-4) | [optional] [default to undefined]
**country_code** | **string** | ISO 3166-1 alpha-2 country code (characters 5-6) | [optional] [default to undefined]
**location_code** | **string** | Location code (characters 7-8) | [optional] [default to undefined]
**branch_code** | **string** | Branch code (characters 9-11), XXX for primary office | [optional] [default to undefined]
**bank_name** | **string** | Bank or institution name | [optional] [default to undefined]
**city** | **string** | City of the branch or primary office | [optional] [default to undefined]
**country_name** | **string** | Full country name | [optional] [default to undefined]
**is_primary** | **boolean** | true when branch_code is XXX | [optional] [default to undefined]

## Example

```typescript
import { V1FinanceSwiftCodeGet200ResponseData } from 'api-client';

const instance: V1FinanceSwiftCodeGet200ResponseData = {
    swift_code,
    bank_code,
    country_code,
    location_code,
    branch_code,
    bank_name,
    city,
    country_name,
    is_primary,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
