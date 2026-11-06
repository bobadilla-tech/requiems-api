# V1TextProfanityPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**has_profanity** | **boolean** | Whether any profanity was detected in the text | [optional] [default to undefined]
**censored** | **string** | The input text with profane words replaced by asterisks | [optional] [default to undefined]
**flagged_words** | **string** | Deduplicated list of profane words found (lowercase) | [optional] [default to undefined]

## Example

```typescript
import { V1TextProfanityPost200ResponseData } from 'api-client';

const instance: V1TextProfanityPost200ResponseData = {
    has_profanity,
    censored,
    flagged_words,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
