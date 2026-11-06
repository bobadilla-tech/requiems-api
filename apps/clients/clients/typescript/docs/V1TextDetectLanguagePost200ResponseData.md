# V1TextDetectLanguagePost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**language** | **string** | Full name of the detected language (e.g. French, English, Spanish) | [optional] [default to undefined]
**code** | **string** | ISO 639-1 two-letter language code (e.g. fr, en, es). Empty string when detection is unreliable. | [optional] [default to undefined]
**confidence** | **string** | Confidence score between 0.0 and 1.0. 0.0 is returned when the language cannot be reliably detected. | [optional] [default to undefined]

## Example

```typescript
import { V1TextDetectLanguagePost200ResponseData } from 'api-client';

const instance: V1TextDetectLanguagePost200ResponseData = {
    language,
    code,
    confidence,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
