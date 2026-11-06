# V1TextLoremBatchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**items** | **Array&lt;any&gt;** | Array of text generation requests. Minimum 1 item, maximum 50 items. | [default to undefined]
**items_paragraphs** | **number** | Number of paragraphs to generate (1-20) | [optional] [default to undefined]
**items_sentences** | **number** | Number of sentences per paragraph (1-20) | [optional] [default to undefined]

## Example

```typescript
import { V1TextLoremBatchPostRequest } from 'api-client';

const instance: V1TextLoremBatchPostRequest = {
    items,
    items_paragraphs,
    items_sentences,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
