# V1SystemsTextNormalizePostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**text** | **string** | The text to normalize. Required. | [default to undefined]
**operations** | **Array&lt;any&gt;** | Ordered list of normalization operations to apply. At least one required. Valid values: trim, lowercase, uppercase, normalize_unicode, collapse_whitespace, strip_html, remove_punctuation | [default to undefined]

## Example

```typescript
import { V1SystemsTextNormalizePostRequest } from 'api-client';

const instance: V1SystemsTextNormalizePostRequest = {
    text,
    operations,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
