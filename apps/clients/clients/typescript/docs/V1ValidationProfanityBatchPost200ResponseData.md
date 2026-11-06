# V1ValidationProfanityBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **Array&lt;any&gt;** | Per-item results in input order. | [optional] [default to undefined]
**results_text** | **string** | The original text that was checked. | [optional] [default to undefined]
**results_result_has_profanity** | **boolean** | Whether any profanity was detected. | [optional] [default to undefined]
**results_result_censored** | **string** | Text with profane words replaced by asterisks. | [optional] [default to undefined]
**results_result_flagged_words** | **string** | Deduplicated list of detected profane words (lowercase). | [optional] [default to undefined]
**total** | **number** | Total number of items processed. | [optional] [default to undefined]

## Example

```typescript
import { V1ValidationProfanityBatchPost200ResponseData } from 'api-client';

const instance: V1ValidationProfanityBatchPost200ResponseData = {
    results,
    results_text,
    results_result_has_profanity,
    results_result_censored,
    results_result_flagged_words,
    total,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
