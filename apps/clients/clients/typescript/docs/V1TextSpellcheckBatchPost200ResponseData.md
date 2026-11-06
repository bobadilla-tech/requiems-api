# V1TextSpellcheckBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **string** | One entry per input text, in the same order as the input array. Each item contains corrected (the fixed text) and corrections (list of individual corrections with original, suggested, suggestions, and position). | [optional] [default to undefined]
**total** | **number** | Number of texts processed. Equals the length of the input array. | [optional] [default to undefined]

## Example

```typescript
import { V1TextSpellcheckBatchPost200ResponseData } from 'api-client';

const instance: V1TextSpellcheckBatchPost200ResponseData = {
    results,
    total,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
