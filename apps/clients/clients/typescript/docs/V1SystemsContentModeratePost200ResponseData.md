# V1SystemsContentModeratePost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**is_safe** | **boolean** | False when any moderation category is triggered. True otherwise. | [optional] [default to undefined]
**toxicity_score** | **string** | Reserved for future use. Currently always null. | [optional] [default to undefined]
**sentiment** | **string** | Detected sentiment: positive, negative, or neutral | [optional] [default to undefined]
**language** | **string** | Detected or provided ISO 639-1 language code. Null when detection confidence is 0. | [optional] [default to undefined]
**language_confidence** | **string** | Confidence in the detected language from 0.0 to 1.0. 1.0 when language was explicitly provided. | [optional] [default to undefined]
**flags** | **Array&lt;any&gt;** | Triggered moderation flags. Currently possible: text_profanity | [optional] [default to undefined]
**categories** | **object** | Per-category boolean flags: profanity, hate_speech, spam, violence | [optional] [default to undefined]

## Example

```typescript
import { V1SystemsContentModeratePost200ResponseData } from 'api-client';

const instance: V1SystemsContentModeratePost200ResponseData = {
    is_safe,
    toxicity_score,
    sentiment,
    language,
    language_confidence,
    flags,
    categories,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
