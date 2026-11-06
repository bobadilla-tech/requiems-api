# V1TextNormalizePost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**original** | **string** | The email address exactly as supplied in the request body | [optional] [default to undefined]
**normalized** | **string** | The canonical form of the address after all transformations | [optional] [default to undefined]
**local** | **string** | The local part (before @) of the normalized address | [optional] [default to undefined]
**domain** | **string** | The domain part (after @) of the normalized address | [optional] [default to undefined]
**changes** | **Array&lt;any&gt;** | Ordered list of transformations applied. Possible values: lowercased, trimmed_whitespace, removed_dots, removed_plus_tag, canonicalised_domain. Empty array when no changes were needed. | [optional] [default to undefined]

## Example

```typescript
import { V1TextNormalizePost200ResponseData } from 'api-client';

const instance: V1TextNormalizePost200ResponseData = {
    original,
    normalized,
    local,
    domain,
    changes,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
