# V1ValidationEmailPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**email** | **string** | The email address exactly as supplied in the request body; null when syntax is invalid | [optional] [default to undefined]
**valid** | **boolean** | Overall validity. True only when the address passes syntax validation and the domain has at least one MX record | [optional] [default to undefined]
**syntax_valid** | **boolean** | Whether the address is syntactically valid according to RFC 5322 | [optional] [default to undefined]
**mx_valid** | **boolean** | Whether the domain has at least one MX record, meaning it can receive email | [optional] [default to undefined]
**disposable** | **boolean** | Whether the address uses a known disposable or temporary email domain | [optional] [default to undefined]
**normalized** | **string** | The canonical form of the address after normalization (lowercase, plus-tag removal, alias-domain resolution). Null when syntax is invalid | [optional] [default to undefined]
**domain** | **string** | The domain part of the address (after @). Null when syntax is invalid | [optional] [default to undefined]
**suggestion** | **string** | A corrected domain name when the supplied domain looks like a typo of a well-known provider (e.g. gmial.com → gmail.com). Null when no close match is found or the domain is already correct | [optional] [default to undefined]

## Example

```typescript
import { V1ValidationEmailPost200ResponseData } from 'api-client';

const instance: V1ValidationEmailPost200ResponseData = {
    email,
    valid,
    syntax_valid,
    mx_valid,
    disposable,
    normalized,
    domain,
    suggestion,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
