# V1SystemsContentModeratePostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**text** | **string** | The text to moderate. Any language. Required. | [default to undefined]
**language** | **string** | ISO 639-1 language code (e.g. en, de, fr). When provided, language detection is skipped and confidence is set to 1.0. | [optional] [default to undefined]

## Example

```typescript
import { V1SystemsContentModeratePostRequest } from 'api-client';

const instance: V1SystemsContentModeratePostRequest = {
    text,
    language,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
